//go:build !windows

package scanner

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// Regex untuk parsing output scanimage -L
// Contoh: device `genesys:libusb:001:002' is a Canon LiDE 220 flatbed scanner
var deviceRegex = regexp.MustCompile("device [`'](.+?)['`] is a (.+)")

func ListScanners() ([]ScannerInfo, error) {
	// Pastikan scanimage terinstall
	path, err := exec.LookPath("scanimage")
	if err != nil {
		return nil, fmt.Errorf("scanimage tidak ditemukan, silakan install sane-utils: %w", err)
	}

	cmd := exec.Command(path, "-L")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("gagal menjalankan scanimage -L: %w", err)
	}

	var scanners []ScannerInfo
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		matches := deviceRegex.FindStringSubmatch(line)
		if len(matches) >= 3 {
			scanners = append(scanners, ScannerInfo{
				ID:   matches[1],
				Name: matches[2],
			})
		}
	}

	return scanners, nil
}

func Scan(req ScanRequest) (*ScanResult, error) {
	path, err := exec.LookPath("scanimage")
	if err != nil {
		return nil, fmt.Errorf("scanimage tidak ditemukan: %w", err)
	}

	args := []string{
		"--device-name", req.ScannerID,
		"--format", req.Format,
	}

	if req.DPI > 0 {
		args = append(args, "--resolution", strconv.Itoa(req.DPI))
	}

	// Mapping Color Mode
	// SANE biasanya: Color, Gray, Lineart
	switch strings.ToLower(req.ColorMode) {
	case "color":
		args = append(args, "--mode", "Color")
	case "grayscale":
		args = append(args, "--mode", "Gray")
	case "bw":
		args = append(args, "--mode", "Lineart")
	}

	cmd := exec.Command(path, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("scan gagal: %w (stderr: %s)", err, stderr.String())
	}

	data := stdout.Bytes()
	if len(data) == 0 {
		return nil, fmt.Errorf("tidak ada data yang dihasilkan dari scanner")
	}

	// Ambil dimensi gambar (opsional tapi bagus untuk UI)
	width, height := 0, 0
	imgConfig, _, err := image.DecodeConfig(bytes.NewReader(data))
	if err == nil {
		width = imgConfig.Width
		height = imgConfig.Height
	}

	return &ScanResult{
		Data:   base64.StdEncoding.EncodeToString(data),
		Format: req.Format,
		Width:  width,
		Height: height,
	}, nil
}
