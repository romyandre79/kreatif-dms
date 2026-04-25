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
	if !s.cfg.AIRefinementEnabled || s.cfg.GeminiAPIKey == "" {
		if len(text) > 100 {
			return text[:100] + "...", nil
		}
		return text, nil
	}

	prompt := fmt.Sprintf("Ringkas teks dokumen berikut dalam 2-3 kalimat yang padat:\n\n%s", text)
	return s.callGemini(ctx, prompt)
}

func (s *AIService) RefineOCRText(ctx context.Context, rawText string) (string, error) {
	if !s.cfg.AIRefinementEnabled || s.cfg.GeminiAPIKey == "" {
		return rawText, nil
	}

	log.Printf("[AIService] Refining OCR text with Gemini AI...")
	prompt := fmt.Sprintf(`Anda adalah asisten digitalisasi profesional. Berikut adalah teks hasil OCR mentah dari sebuah dokumen yang mungkin mengandung typo, teks terpotong, atau noise dari stempel dan tanda tangan.
Tugas Anda:
1. Perbaiki typo dan susun kembali kalimat agar mengalir dengan baik.
2. Hilangkan karakter sampah hasil noise OCR (seperti @, |, _, dll).
3. Pertahankan struktur data penting seperti Nomor Dokumen, Tanggal, dan Pihak Terkait.
4. Kembalikan HANYA teks yang sudah bersih tanpa penjelasan tambahan.

TEKS MENTAH:
%s`, rawText)

	refined, err := s.callGemini(ctx, prompt)
	if err != nil {
		log.Printf("[AIService] Warning: AI refinement failed, falling back to raw text: %v", err)
		return rawText, nil
	}

	return refined, nil
}

func (s *AIService) ExtractMetadata(ctx context.Context, text string) (map[string]interface{}, error) {
	if !s.cfg.AIMetadataEnabled || s.cfg.GeminiAPIKey == "" {
		return map[string]interface{}{}, nil
	}

	log.Printf("[AIService] Extracting metadata with Gemini AI...")
	prompt := fmt.Sprintf(`Anda adalah asisten ekstraksi data dokumen. Ekstrak informasi penting dari teks dokumen berikut ke dalam format JSON.
Informasi yang harus dicari:
- document_number (Nomor Dokumen/Surat)
- document_date (Tanggal Dokumen dalam format YYYY-MM-DD)
- document_type (Tipe Dokumen, misal: Invoice, Perjanjian, Memo, dll)
- parties (Daftar pihak yang terlibat, misal: Nama Perusahaan, Nama Orang)
- total_amount (Jika ada nilai uang, ambil angkanya saja)
- currency (Mata uang jika ada)

Kembalikan HANYA JSON valid tanpa penjelasan tambahan.

TEKS DOKUMEN:
%s`, text)

	jsonStr, err := s.callGemini(ctx, prompt)
	if err != nil {
		return map[string]interface{}{}, err
	}

	// Clean up JSON if AI adds markdown code blocks
	jsonStr = s.cleanJSONString(jsonStr)

	var metadata map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &metadata); err != nil {
		log.Printf("[AIService] Error parsing AI metadata JSON: %v. Raw: %s", err, jsonStr)
		return map[string]interface{}{}, err
	}

	return metadata, nil
}

func (s *AIService) cleanJSONString(s_input string) string {
	// Simple cleanup for markdown blocks
	if len(s_input) > 7 && s_input[:7] == "```json" {
		s_input = s_input[7:]
		if len(s_input) > 3 && s_input[len(s_input)-3:] == "```" {
			s_input = s_input[:len(s_input)-3]
		}
	} else if len(s_input) > 3 && s_input[:3] == "```" {
		s_input = s_input[3:]
		if len(s_input) > 3 && s_input[len(s_input)-3:] == "```" {
			s_input = s_input[:len(s_input)-3]
		}
	}
	return s_input
}

func (s *AIService) callGemini(ctx context.Context, prompt string) (string, error) {
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent?key=%s", s.cfg.GeminiAPIKey)

	payload := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{"text": prompt},
				},
			},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("gemini api error: %d - %s", resp.StatusCode, string(body))
	}

	var result struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Candidates) > 0 && len(result.Candidates[0].Content.Parts) > 0 {
		return result.Candidates[0].Content.Parts[0].Text, nil
	}

	return "", fmt.Errorf("no content generated from gemini")
}
