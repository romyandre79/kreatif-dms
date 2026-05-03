package scanner

type ScannerInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ScanRequest struct {
	ScannerID string `json:"scanner_id"`
	DPI       int    `json:"dpi"`         // 150, 300, 600
	ColorMode string `json:"color_mode"`  // "color", "grayscale", "bw"
	Format    string `json:"format"`      // "jpeg", "png"
}

type ScanResult struct {
	Data   string `json:"data"`
	Format string `json:"format"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
