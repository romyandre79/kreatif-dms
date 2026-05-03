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
	"strings"

	"github.com/kreatif/dms-backend/internal/config"
	"github.com/kreatif/dms-backend/internal/repository"
)

type AIService struct {
	repo repository.Querier
	cfg  config.Config
}

func NewAIService(repo repository.Querier, cfg config.Config) *AIService {
	return &AIService{repo: repo, cfg: cfg}
}

type AIInsight struct {
	DocType      string                 `json:"doc_type"`
	Summary      string                 `json:"summary"`
	Entities     map[string]interface{} `json:"entities"`
	CleanedText  string                 `json:"cleaned_text"`
	Confidence   float64                `json:"confidence_score"`
}

type OCRResponse struct {
	Status      string      `json:"status"`
	Filename    string      `json:"filename"`
	FullText    string      `json:"full_text"`
	Insight     *AIInsight  `json:"insight"`
	AIAnalysis  interface{} `json:"ai_analysis"`
	PreviewPath string      `json:"preview_path"`
	PreviewPaths []string   `json:"preview_paths"`
	Words       []struct {
		Text       string  `json:"text"`
		Confidence float64 `json:"confidence"`
		Page       int     `json:"page"`
		Box        struct {
			X int `json:"x"`
			Y int `json:"y"`
			W int `json:"w"`
			H int `json:"h"`
		} `json:"box"`
	} `json:"words"`
}

func (s *AIService) getOCRConfig(ctx context.Context) (string, string, string, error) {
	node, err := s.repo.GetIntegrationNodeByType(ctx, "OCR")
	if err != nil {
		return "", "", "", fmt.Errorf("OCR integration node not found in database: %v", err)
	}

	if !node.IsActive.Bool {
		return "", "", "", fmt.Errorf("OCR integration is disabled in database")
	}

	var nodeCfg struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.Unmarshal(node.ConfigJson, &nodeCfg); err != nil {
		return "", "", "", fmt.Errorf("failed to parse OCR config JSON: %v", err)
	}

	url := node.Endpoint
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	return url, nodeCfg.Username, nodeCfg.Password, nil
}

func (s *AIService) ProcessOCR(ctx context.Context, fileName string, content []byte) (*OCRResponse, error) {
	log.Printf("[AIService] Processing OCR for file: %s (Size: %d bytes)", fileName, len(content))
	
	// Fetch config from Database
	ocrURL, ocrUser, ocrPass, err := s.getOCRConfig(ctx)
	if err != nil {
		log.Printf("[AIService] Error getting OCR config: %v", err)
		return nil, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add AI config if available (Fields BEFORE file)
	_, aiKey, aiModel, _, err := s.getAIConfig(ctx)
	if err == nil {
		log.Printf("[AIService] Sending AI Config to OCR: Model=%s, KeyLen=%d", aiModel, len(aiKey))
		if aiModel != "" {
			writer.WriteField("ai_model", aiModel)
		}
		if aiKey != "" {
			writer.WriteField("ai_api_key", aiKey)
		}
	} else {
		log.Printf("[AIService] Warning: Could not get AI config for OCR: %v", err)
	}

	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		log.Printf("[AIService] Error creating form file: %v", err)
		return nil, err
	}
	_, err = io.Copy(part, bytes.NewReader(content))
	if err != nil {
		log.Printf("[AIService] Error copying content to form: %v", err)
		return nil, err
	}

	writer.Close()

	req, err := http.NewRequestWithContext(ctx, "POST", ocrURL+"/ocr/process", body)
	if err != nil {
		log.Printf("[AIService] Error creating OCR request: %v", err)
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	
	// Add Basic Auth
	if ocrUser != "" && ocrPass != "" {
		log.Printf("[AIService] Sending Basic Auth with User: %s (Password Length: %d)", ocrUser, len(ocrPass))
		req.SetBasicAuth(ocrUser, ocrPass)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[AIService] Error calling OCR service: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("[AIService] OCR service returned error code: %d", resp.StatusCode)
		return nil, fmt.Errorf("OCR service error: %d", resp.StatusCode)
	}

	var ocrRes OCRResponse
	if err := json.NewDecoder(resp.Body).Decode(&ocrRes); err != nil {
		log.Printf("[AIService] Error decoding OCR response: %v", err)
		return nil, err
	}

	return &ocrRes, nil
}

func (s *AIService) Summarize(ctx context.Context, text string) (string, error) {
	if !s.cfg.AIRefinementEnabled || s.cfg.GeminiAPIKey == "" {
		if len(text) > 100 {
			return text[:100] + "...", nil
		}
		return text, nil
	}

	prompt := fmt.Sprintf("Ringkas teks dokumen berikut dalam 2-3 kalimat yang padat:\n\n%s", text)
	return s.callAI(ctx, prompt)
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

	refined, err := s.callAI(ctx, prompt)
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

	jsonStr, err := s.callAI(ctx, prompt)
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

func (s *AIService) getAIConfig(ctx context.Context) (string, string, string, string, error) {
	node, err := s.repo.GetIntegrationNodeByType(ctx, "AI")
	if err != nil {
		return "", "", "", "", fmt.Errorf("AI integration node not found: %v", err)
	}

	if !node.IsActive.Bool {
		return "", "", "", "", fmt.Errorf("AI integration is disabled")
	}

	var nodeCfg struct {
		Token        string `json:"token"`
		Model        string `json:"model"`
		SystemPrompt string `json:"system_prompt"`
	}
	if err := json.Unmarshal(node.ConfigJson, &nodeCfg); err != nil {
		return "", "", "", "", fmt.Errorf("failed to parse AI config: %v", err)
	}

	driver := node.Driver.String
	apiKey := nodeCfg.Token
	model := nodeCfg.Model
	systemPrompt := nodeCfg.SystemPrompt

	if apiKey == "" {
		apiKey = s.cfg.GeminiAPIKey
		driver = "gemini"
	}

	// Final fallback for driver if still empty
	if driver == "" {
		driver = "gemini"
	}

	if model == "" {
		model = "gemini-1.5-flash"
	}
	if systemPrompt == "" {
		systemPrompt = "You are a helpful assistant for Kreatif DMS."
	}

	return driver, apiKey, model, systemPrompt, nil
}

func (s *AIService) callAI(ctx context.Context, prompt string) (string, error) {
	driver, apiKey, model, systemPrompt, err := s.getAIConfig(ctx)
	if err != nil {
		// Fallback to legacy config if node not found
		if s.cfg.GeminiAPIKey != "" {
			driver = "gemini"
			apiKey = s.cfg.GeminiAPIKey
			model = "gemini-1.5-flash"
			systemPrompt = "You are a helpful assistant for Kreatif DMS."
		} else {
			return "", err
		}
	}

	if driver == "openai" {
		return s.callOpenAI(ctx, apiKey, model, systemPrompt, prompt)
	}

	return s.callGemini(ctx, apiKey, model, systemPrompt, prompt)
}

func (s *AIService) callOpenAI(ctx context.Context, apiKey, model, systemPrompt, prompt string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"
	payload := map[string]interface{}{
		"model": model,
		"messages": []map[string]interface{}{
			{"role": "system", "content": systemPrompt},
			{"role": "user", "content": prompt},
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
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("openai api error: %d - %s", resp.StatusCode, string(body))
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Choices) > 0 {
		return result.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no content generated from openai")
}

func (s *AIService) callGemini(ctx context.Context, apiKey, model, systemPrompt, prompt string) (string, error) {
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", model, apiKey)

	payload := map[string]interface{}{
		"system_instruction": map[string]interface{}{
			"parts": []map[string]interface{}{
				{"text": systemPrompt},
			},
		},
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
func (s *AIService) FetchGeminiModels(ctx context.Context, apiKey string) ([]map[string]string, error) {
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models?key=%s", apiKey)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("gemini api error: %d - %s", resp.StatusCode, string(body))
	}

	var result struct {
		Models []struct {
			Name                       string   `json:"name"`
			DisplayName                string   `json:"displayName"`
			SupportedGenerationMethods []string `json:"supportedGenerationMethods"`
		} `json:"models"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	models := []map[string]string{}
	for _, m := range result.Models {
		isGenerative := false
		for _, method := range m.SupportedGenerationMethods {
			if method == "generateContent" {
				isGenerative = true
				break
			}
		}
		if isGenerative {
			models = append(models, map[string]string{
				"id":   m.Name[7:], // remove "models/"
				"name": m.DisplayName,
			})
		}
	}

	return models, nil
}

func (s *AIService) FetchOpenAIModels(ctx context.Context, apiKey string) ([]map[string]string, error) {
	url := "https://api.openai.com/v1/models"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("openai api error: %d - %s", resp.StatusCode, string(body))
	}

	var result struct {
		Data []struct {
			ID string `json:"id"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	models := []map[string]string{}
	for _, m := range result.Data {
		// Filter for common gpt models
		if len(m.ID) >= 3 && m.ID[:3] == "gpt" {
			models = append(models, map[string]string{
				"id":   m.ID,
				"name": m.ID,
			})
		}
	}

	return models, nil
}
