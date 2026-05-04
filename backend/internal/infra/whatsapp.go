package infra

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/kreatif/dms-backend/internal/repository"
)

type WhatsAppService struct {
	repo repository.Querier
}

func NewWhatsAppService(repo repository.Querier) *WhatsAppService {
	return &WhatsAppService{repo: repo}
}

type whatsappNodeConfig struct {
	Token    string `json:"token"`
	Endpoint string `json:"endpoint"`
}

func (s *WhatsAppService) getWhatsAppConfig(ctx context.Context) (string, string, string, string, error) {
	node, err := s.repo.GetIntegrationNodeByType(ctx, "WHATSAPP")
	if err != nil {
		return "", "", "", "", fmt.Errorf("whatsapp integration node not found: %v", err)
	}

	if !node.IsActive.Bool {
		return "", "", "", "", fmt.Errorf("whatsapp integration is disabled")
	}

	var nodeCfg struct {
		Token      string `json:"token"`
		AuthHeader string `json:"auth_header"`
	}
	if err := json.Unmarshal(node.ConfigJson, &nodeCfg); err != nil {
		return "", "", "", "", fmt.Errorf("failed to parse whatsapp config: %v", err)
	}

	driver := node.Driver.String
	if driver == "" {
		driver = "fonnte" // Default legacy
	}

	endpoint := node.Endpoint
	token := nodeCfg.Token
	authHeader := nodeCfg.AuthHeader

	return driver, endpoint, token, authHeader, nil
}

func (s *WhatsAppService) SendMessage(ctx context.Context, to string, message string) error {
	driver, endpoint, token, authHeader, err := s.getWhatsAppConfig(ctx)
	if err != nil {
		return err
	}

	log.Printf("[WhatsAppService] Sending message to %s using driver %s", to, driver)

	var req *http.Request
	switch driver {
	case "fonnte":
		if endpoint == "" {
			endpoint = "https://api.fonnte.com/send"
		}
		data := map[string]string{"target": to, "message": message}
		jsonData, _ := json.Marshal(data)
		req, _ = http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(jsonData))
		req.Header.Set("Authorization", token)
		req.Header.Set("Content-Type", "application/json")

	case "wablas":
		if endpoint == "" {
			endpoint = "https://api.wablas.com/api/send-message"
		}
		data := map[string]string{"phone": to, "message": message}
		jsonData, _ := json.Marshal(data)
		req, _ = http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(jsonData))
		req.Header.Set("Authorization", token)
		req.Header.Set("Content-Type", "application/json")

	case "generic":
		data := map[string]string{"to": to, "message": message}
		jsonData, _ := json.Marshal(data)
		req, _ = http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(jsonData))
		if authHeader != "" {
			req.Header.Set("Authorization", authHeader)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("ngrok-skip-browser-warning", "true")

	default:
		return fmt.Errorf("unsupported whatsapp driver: %s", driver)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("whatsapp provider error (%d): %s", resp.StatusCode, string(body))
	}

	return nil
}

func (s *WhatsAppService) SendTemplate(ctx context.Context, to string, template string, params map[string]string) error {
	// Simple template replacement
	msg := template
	for k, v := range params {
		msg = strings.ReplaceAll(msg, "{{"+k+"}}", v)
	}
	return s.SendMessage(ctx, to, msg)
}
func (s *WhatsAppService) TestConnection(ctx context.Context, endpoint, token string) (map[string]interface{}, error) {
	// We need driver context for testing too, or just test Fonnte as default
	node, _ := s.repo.GetIntegrationNodeByType(ctx, "WHATSAPP")
	driver := "fonnte"
	if node.Driver.Valid {
		driver = node.Driver.String
	}

	var testUrl string
	switch driver {
	case "fonnte":
		testUrl = "https://api.fonnte.com/device"
	case "wablas":
		testUrl = "https://api.wablas.com/api/device/info"
	case "generic":
		testUrl = endpoint
	default:
		testUrl = "https://api.fonnte.com/device"
	}

	req, err := http.NewRequestWithContext(ctx, "POST", testUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", token)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("whatsapp gateway error: %d", resp.StatusCode)
	}

	return result, nil
}
