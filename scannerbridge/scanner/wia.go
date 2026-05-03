//go:build windows

package scanner

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

const (
	wiaFormatBMP  = "{B96B3CAE-0728-11D3-9D7B-0000F81EF32E}"
	wiaFormatJPEG = "{B96B3CB4-0728-11D3-9D7B-0000F81EF32E}"
	wiaFormatPNG  = "{B96B3CB1-0728-11D3-9D7B-0000F81EF32E}"

	wiaIntentColor     int32 = 1
	wiaIntentGrayscale int32 = 2
	wiaIntentText      int32 = 4
)

func varInt(v *ole.VARIANT) int {
	if v == nil {
		return 0
	}
	val := v.Value()
	if val == nil {
		return 0
	}
	switch x := val.(type) {
	case int16:
		return int(x)
	case int32:
		return int(x)
	case int64:
		return int(x)
	default:
		return 0
	}
}

func newDeviceManager() (*ole.IDispatch, func(), error) {
	dmUnk, err := oleutil.CreateObject("WIA.DeviceManager")
	if err != nil {
		return nil, nil, fmt.Errorf("WIA.DeviceManager tidak tersedia: %w", err)
	}
	dmDisp, err := dmUnk.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		dmUnk.Release()
		return nil, nil, err
	}
	return dmDisp, func() {
		dmDisp.Release()
		dmUnk.Release()
	}, nil
}

func ListScanners() ([]ScannerInfo, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED)
	defer ole.CoUninitialize()

	dmDisp, cleanup, err := newDeviceManager()
	if err != nil {
		return nil, err
	}
	defer cleanup()

	diVar, err := oleutil.GetProperty(dmDisp, "DeviceInfos")
	if err != nil {
		return nil, fmt.Errorf("get DeviceInfos: %w", err)
	}
	if diVar != nil {
		defer diVar.Clear()
	} else {
		return []ScannerInfo{}, nil
	}

	diDisp := diVar.ToIDispatch()
	if diDisp == nil {
		return []ScannerInfo{}, nil
	}
	defer diDisp.Release()

	countVar, err := oleutil.GetProperty(diDisp, "Count")
	count := 0
	if err == nil && countVar != nil {
		count = varInt(countVar)
		countVar.Clear()
	}

	fmt.Printf("[Bridge] Scanning for devices... Found Count: %d\n", count)

	var result []ScannerInfo
	for i := 1; i <= count; i++ {
		itemVar, err := oleutil.GetProperty(diDisp, "Item", i)
		if err != nil {
			continue
		}
		itemDisp := itemVar.ToIDispatch()
		if itemDisp == nil {
			if itemVar != nil { itemVar.Clear() }
			continue
		}

		typeVar, errType := oleutil.GetProperty(itemDisp, "Type")
		devType := 0
		if errType == nil && typeVar != nil {
			devType = varInt(typeVar)
			typeVar.Clear()
		}

		idVar, errID := oleutil.GetProperty(itemDisp, "DeviceID")
		nameVar, errName := oleutil.GetProperty(itemDisp, "Name")
		
		deviceName := "Unknown Device"
		deviceID := ""

		if errID == nil && idVar != nil {
			deviceID = idVar.ToString()
			idVar.Clear()
		}
		if errName == nil && nameVar != nil {
			deviceName = nameVar.ToString()
			nameVar.Clear()
		}

		fmt.Printf("[Bridge] Device Found: %s (Type: %d, ID: %s)\n", deviceName, devType, deviceID)

		result = append(result, ScannerInfo{
			ID:   deviceID,
			Name: deviceName,
		})

		itemDisp.Release()
		if itemVar != nil { itemVar.Clear() }
	}

	return result, nil
}

func Scan(req ScanRequest) (*ScanResult, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	fmt.Println("[Bridge] Memulai proses scan...")
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED)
	defer ole.CoUninitialize()

	dmDisp, cleanup, err := newDeviceManager()
	if err != nil {
		return nil, err
	}
	defer cleanup()

	// 1. Cari scanner
	deviceInfoDisp, err := findScanner(dmDisp, req.ScannerID)
	if err != nil {
		return nil, err
	}
	defer deviceInfoDisp.Release()

	// 2. Connect
	fmt.Println("[Bridge] Connecting to hardware...")
	deviceVar, err := oleutil.CallMethod(deviceInfoDisp, "Connect")
	if err != nil {
		return nil, fmt.Errorf("gagal koneksi ke scanner: %w", err)
	}
	// Manual clear instead of defer to be safer
	deviceDisp := deviceVar.ToIDispatch()
	
	if deviceDisp == nil {
		deviceVar.Clear()
		return nil, fmt.Errorf("gagal mendapatkan device interface")
	}

	// 3. Get Items
	itemsVar, err := oleutil.GetProperty(deviceDisp, "Items")
	if err != nil {
		deviceDisp.Release()
		deviceVar.Clear()
		return nil, fmt.Errorf("gagal ambil items: %w", err)
	}
	itemsDisp := itemsVar.ToIDispatch()
	if itemsDisp == nil {
		if itemsVar != nil { itemsVar.Clear() }
		deviceDisp.Release()
		deviceVar.Clear()
		return nil, fmt.Errorf("items bukan IDispatch")
	}

	// 4. Get first item
	itemVar, err := oleutil.GetProperty(itemsDisp, "Item", 1)
	if err != nil {
		itemsDisp.Release()
		if itemsVar != nil { itemsVar.Clear() }
		deviceDisp.Release()
		deviceVar.Clear()
		return nil, fmt.Errorf("scanner item tidak ditemukan: %w", err)
	}
	itemDisp := itemVar.ToIDispatch()
	if itemDisp == nil {
		if itemVar != nil { itemVar.Clear() }
		itemsDisp.Release()
		if itemsVar != nil { itemsVar.Clear() }
		deviceDisp.Release()
		deviceVar.Clear()
		return nil, fmt.Errorf("item bukan IDispatch")
	}

	// Apply settings & Transfer
	fmt.Println("[Bridge] Applying settings & starting transfer...")
	applyScanSettings(itemDisp, req)

	outputFormat, _, ext := resolveFormat(req.Format)
	imageFileVar, err := oleutil.CallMethod(itemDisp, "Transfer", outputFormat)
	if err != nil {
		itemDisp.Release()
		if itemVar != nil { itemVar.Clear() }
		itemsDisp.Release()
		if itemsVar != nil { itemsVar.Clear() }
		deviceDisp.Release()
		deviceVar.Clear()
		return nil, fmt.Errorf("transfer gagal: %w", err)
	}
	
	imageFileDisp := imageFileVar.ToIDispatch()
	if imageFileDisp == nil {
		if imageFileVar != nil { imageFileVar.Clear() }
		itemDisp.Release()
		// ... cleanup ...
		return nil, fmt.Errorf("transfer tidak mengembalikan IDispatch")
	}

	// Baca dimensi (safe)
	width := 0
	height := 0
	wV, _ := oleutil.GetProperty(imageFileDisp, "Width")
	hV, _ := oleutil.GetProperty(imageFileDisp, "Height")
	width = varInt(wV)
	height = varInt(hV)
	if wV != nil { wV.Clear() }
	if hV != nil { hV.Clear() }

	// Simpan ke temp file
	tmpPath := filepath.Join(os.TempDir(), fmt.Sprintf("scan_%d%s", time.Now().UnixNano(), ext))
	fmt.Printf("[Bridge] Saving to %s...\n", tmpPath)
	_, err = oleutil.CallMethod(imageFileDisp, "SaveFile", tmpPath)
	
	// Cleanup OLE Objects before reading file to free memory
	imageFileDisp.Release()
	if imageFileVar != nil { imageFileVar.Clear() }
	itemDisp.Release()
	if itemVar != nil { itemVar.Clear() }
	itemsDisp.Release()
	if itemsVar != nil { itemsVar.Clear() }
	deviceDisp.Release()
	if deviceVar != nil { deviceVar.Clear() }

	if err != nil {
		return nil, fmt.Errorf("simpan hasil scan: %w", err)
	}

	data, err := os.ReadFile(tmpPath)
	os.Remove(tmpPath)
	if err != nil {
		return nil, fmt.Errorf("baca file scan: %w", err)
	}

	fmt.Printf("[Bridge] Scan SELESAI. Size: %.2f MB\n", float64(len(data))/(1024*1024))
	return &ScanResult{
		Data:   base64.StdEncoding.EncodeToString(data),
		Format: outputFormat,
		Width:  width,
		Height: height,
	}, nil
}

func findScanner(dmDisp *ole.IDispatch, scannerID string) (*ole.IDispatch, error) {
	diVar, err := oleutil.GetProperty(dmDisp, "DeviceInfos")
	if err != nil {
		return nil, err
	}
	if diVar != nil {
		defer diVar.Clear()
	}
	
	diDisp := diVar.ToIDispatch()
	if diDisp == nil {
		return nil, fmt.Errorf("DeviceInfos bukan IDispatch")
	}
	defer diDisp.Release()

	countVar, err := oleutil.GetProperty(diDisp, "Count")
	count := 0
	if err == nil && countVar != nil {
		count = varInt(countVar)
		countVar.Clear()
	}

	for i := 1; i <= count; i++ {
		itemVar, err := oleutil.GetProperty(diDisp, "Item", i)
		if err != nil {
			continue
		}
		itemDisp := itemVar.ToIDispatch()
		if itemDisp == nil {
			if itemVar != nil { itemVar.Clear() }
			continue
		}

		idVar, errID := oleutil.GetProperty(itemDisp, "DeviceID")
		found := false
		if errID == nil && idVar != nil {
			if idVar.ToString() == scannerID {
				found = true
			}
			idVar.Clear()
		}

		if found {
			if itemVar != nil { itemVar.Clear() }
			return itemDisp, nil
		}
		
		itemDisp.Release()
		if itemVar != nil { itemVar.Clear() }
	}

	return nil, fmt.Errorf("scanner tidak ditemukan: %s", scannerID)
}

func applyScanSettings(itemDisp *ole.IDispatch, req ScanRequest) {
	propsVar, err := oleutil.GetProperty(itemDisp, "Properties")
	if err != nil {
		return
	}
	if propsVar != nil {
		defer propsVar.Clear()
	}
	
	propsDisp := propsVar.ToIDispatch()
	if propsDisp == nil {
		return
	}
	defer propsDisp.Release()

	if req.DPI > 0 {
		setWIAProp(propsDisp, "Horizontal Resolution", int32(req.DPI))
		setWIAProp(propsDisp, "Vertical Resolution", int32(req.DPI))
	}

	switch req.ColorMode {
	case "color":
		setWIAProp(propsDisp, "Current Intent", wiaIntentColor)
	case "grayscale":
		setWIAProp(propsDisp, "Current Intent", wiaIntentGrayscale)
	case "bw":
		setWIAProp(propsDisp, "Current Intent", wiaIntentText)
	}
}

func setWIAProp(propsDisp *ole.IDispatch, name string, value interface{}) {
	propVar, err := oleutil.CallMethod(propsDisp, "Item", name)
	if err != nil {
		return
	}
	if propVar != nil {
		defer propVar.Clear()
	}
	
	propDisp := propVar.ToIDispatch()
	if propDisp == nil {
		return
	}
	defer propDisp.Release()
	
	oleutil.PutProperty(propDisp, "Value", value) //nolint:errcheck
}

func resolveFormat(format string) (guid, name, ext string) {
	switch format {
	case "png":
		return wiaFormatPNG, "png", ".png"
	default:
		return wiaFormatJPEG, "jpeg", ".jpg"
	}
}
