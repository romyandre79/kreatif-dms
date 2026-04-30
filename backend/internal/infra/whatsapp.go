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

func (s *WhatsAppService) getWhatsAppConfig(ctx context.Context) (string, string, error) {
	node, err := s.repo.GetIntegrationNodeByType(ctx, "WHATSAPP")
	if err != nil {
		return "", "", fmt.Errorf("whatsapp integration node not found: %v", err)
	}

	if !node.IsActive.Bool {
		return "", "", fmt.Errorf("whatsapp integration is disabled")
	}

	var nodeCfg whatsappNodeConfig
	if err := json.Unmarshal(node.ConfigJson, &nodeCfg); err != nil {
		return "", "", fmt.Errorf("failed to parse whatsapp config: %v", err)
	}

	endpoint := node.Endpoint
	if endpoint == "" {
		endpoint = nodeCfg.Endpoint
	}

	// Default Fonnte endpoint if not specified
	if endpoint == "" {
		endpoint = "https://api.fonnte.com/send"
	}

	return endpoint, nodeCfg.Token, nil
}

func (s *WhatsAppService) SendMessage(ctx context.Context, to string, message string) error {
	endpoint, token, err := s.getWhatsAppConfig(ctx)
	if err != nil {
		return err
	}

	log.Printf("[WhatsAppService] Sending message to %s", to)

	// Fonnte expects multipart/form-data or application/x-www-form-urlencoded
	// but can also accept JSON in some versions. We'll use a simple form-encoded approach.
	
	data := map[string]string{
		"target":  to,
		"message": message,
	}
	
	jsonData, _ := json.Marshal(data)

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Printf("[WhatsAppService] ERROR: %s", string(body))
		return fmt.Errorf("fonnte API returned status %d: %s", resp.StatusCode, string(body))
	}

	log.Printf("[WhatsAppService] Message sent successfully to %s", to)
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
