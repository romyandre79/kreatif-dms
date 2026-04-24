package infra

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"log"

	"github.com/kreatif/dms-backend/internal/config"
)

type AIService struct {
	cfg config.Config
}

func NewAIService(cfg config.Config) *AIService {
	return &AIService{cfg: cfg}
}

type OCRResponse struct {
	Words []struct {
		Text       string `json:"text"`
		Confidence float64 `json:"confidence"`
	} `json:"words"`
}

func (s *AIService) ProcessOCR(ctx context.Context, fileName string, content []byte) (string, error) {
	log.Printf("[AIService] Processing OCR for file: %s (Size: %d bytes)", fileName, len(content))
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		log.Printf("[AIService] Error creating form file: %v", err)
		return "", err
	}
	_, err = io.Copy(part, bytes.NewReader(content))
	if err != nil {
		log.Printf("[AIService] Error copying content to form: %v", err)
		return "", err
	}
	writer.Close()

	req, err := http.NewRequestWithContext(ctx, "POST", s.cfg.OCRServiceURL+"/ocr/process", body)
	if err != nil {
		log.Printf("[AIService] Error creating OCR request: %v", err)
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[AIService] Error calling OCR service: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("[AIService] OCR service returned error code: %d", resp.StatusCode)
		return "", fmt.Errorf("OCR service error: %d", resp.StatusCode)
	}

	var ocrRes OCRResponse
	if err := json.NewDecoder(resp.Body).Decode(&ocrRes); err != nil {
		log.Printf("[AIService] Error decoding OCR response: %v", err)
		return "", err
	}

	var fullText string
	for _, w := range ocrRes.Words {
		fullText += w.Text + " "
	}

	return fullText, nil
}

func (s *AIService) Summarize(ctx context.Context, text string) (string, error) {
	log.Printf("[AIService] Summarizing text (Length: %d)", len(text))
	// Placeholder for Gemini/OpenAI call
	// For now, return a mock summary
	if len(text) > 100 {
		return text[:100] + "...", nil
	}
	return text, nil
}
