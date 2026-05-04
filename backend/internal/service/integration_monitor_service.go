package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
	"crypto/tls"

	"github.com/go-ldap/ldap/v3"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kreatif/dms-backend/internal/repository"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type IntegrationMonitorService struct {
	repo repository.Querier
}

func NewIntegrationMonitorService(repo repository.Querier) *IntegrationMonitorService {
	return &IntegrationMonitorService{repo: repo}
}

func (s *IntegrationMonitorService) StartMonitoring(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	log.Println("[IntegrationMonitor] Monitoring worker started (interval: 30s)")

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s.CheckAllNodes(ctx)
			}
		}
	}()
}

func (s *IntegrationMonitorService) CheckAllNodes(ctx context.Context) {
	nodes, err := s.repo.ListIntegrationNodes(ctx)
	if err != nil {
		log.Printf("[IntegrationMonitor] Error listing nodes: %v", err)
		return
	}

	for _, node := range nodes {
		if !node.IsActive.Bool {
			continue
		}

		go s.checkNode(ctx, node)
	}
}

func (s *IntegrationMonitorService) checkNode(ctx context.Context, node repository.IntegrationNode) {
	var status string = "online"
	var latency int32 = 0
	var lastErr string = ""

	start := time.Now()

	switch node.ServiceType {
	case "LDAP":
		err := s.checkLDAP(node.Endpoint)
		if err != nil {
			status = "offline"
			lastErr = err.Error()
		}
	case "S3":
		err := s.checkS3(node)
		if err != nil {
			status = "offline"
			lastErr = err.Error()
		}
	case "DATABASE", "ELASTICSEARCH":
		err := s.checkTCP(node.Endpoint)
		if err != nil {
			status = "offline"
			lastErr = err.Error()
		}
	case "WHATSAPP", "AI":
		target := node.Endpoint
		if strings.HasPrefix(target, "http") {
			healthURL := strings.TrimSuffix(target, "/") + "/health"
			err := s.checkHTTP(healthURL)
			if err != nil {
				status = "offline"
				lastErr = err.Error()
			}
		} else {
			err := s.checkHTTP(target)
			if err != nil {
				status = "offline"
				lastErr = err.Error()
			}
		}
	case "SMTP", "OCR":
		target := node.Endpoint
		if strings.HasPrefix(target, "http") {
			// Ensure exactly one slash before health
			healthURL := strings.TrimSuffix(target, "/") + "/health"
			err := s.checkHTTP(healthURL)
			if err != nil {
				status = "offline"
				lastErr = err.Error()
			}
		} else {
			err := s.checkTCP(target)
			if err != nil {
				status = "offline"
				lastErr = err.Error()
			}
		}
	case "SCANNER_LOCAL", "SCANNER_NETWORK":
		// Use connection_string from config for scanners
		var cfg struct {
			ConnectionString string `json:"connection_string"`
		}
		json.Unmarshal(node.ConfigJson, &cfg)
		
		target := cfg.ConnectionString
		if target == "" {
			target = node.Endpoint
		}
		
		// If it's a URL, try /health
		if strings.HasPrefix(target, "http") {
			if !strings.HasSuffix(target, "/") {
				target += "/"
			}
			err := s.checkHTTP(target + "health")
			if err != nil {
				status = "offline"
				lastErr = err.Error()
			}
		} else {
			err := s.checkTCP(target)
			if err != nil {
				status = "offline"
				lastErr = err.Error()
			}
		}
	case "HTTP", "SCANNER", "PRINTER":
		err := s.checkHTTP(node.Endpoint)
		if err != nil {
			status = "offline"
			lastErr = err.Error()
		}
	default:
		err := s.checkTCP(node.Endpoint)
		if err != nil {
			status = "offline"
			lastErr = err.Error()
		}
	}

	latency = int32(time.Since(start).Milliseconds())

	_, err := s.repo.UpdateIntegrationNodeStatus(ctx, repository.UpdateIntegrationNodeStatusParams{
		ID:          node.ID,
		Status:      pgtype.Text{String: status, Valid: true},
		LastLatency: pgtype.Int4{Int32: latency, Valid: true},
		LastError:   pgtype.Text{String: lastErr, Valid: lastErr != ""},
	})
	if err != nil {
		log.Printf("[IntegrationMonitor] Error updating status for %s: %v", node.Name, err)
	}
}

func (s *IntegrationMonitorService) checkLDAP(endpoint string) error {
	url := endpoint
	if !strings.HasPrefix(url, "ldap://") && !strings.HasPrefix(url, "ldaps://") {
		url = "ldap://" + url
	}
	
	l, err := ldap.DialURL(url)
	if err != nil {
		return err
	}
	l.Close()
	return nil
}

func (s *IntegrationMonitorService) checkS3(node repository.IntegrationNode) error {
	var nodeCfg struct {
		AccessKey string      `json:"access_key"`
		SecretKey string      `json:"secret_key"`
		UseSSL    interface{} `json:"use_ssl"`
	}
	json.Unmarshal(node.ConfigJson, &nodeCfg)

	useSSL := false
	switch v := nodeCfg.UseSSL.(type) {
	case bool:
		useSSL = v
	case string:
		useSSL = (v == "true" || v == "1")
	}

	client, err := minio.New(node.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(nodeCfg.AccessKey, nodeCfg.SecretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.ListBuckets(ctx)
	return err
}

func (s *IntegrationMonitorService) checkTCP(endpoint string) error {
	conn, err := net.DialTimeout("tcp", endpoint, 5*time.Second)
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}

func (s *IntegrationMonitorService) checkHTTP(endpoint string) error {
	url := endpoint
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url // Default to https for external APIs
	}

	// Create custom transport to skip SSL verification (common issue on RHEL 7)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: tr,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Add ngrok bypass header
	req.Header.Set("ngrok-skip-browser-warning", "true")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[IntegrationMonitor] Error connecting to %s: %v", url, err)
		return err
	}
	defer resp.Body.Close()

	// 200-399 are considered "online"
	if resp.StatusCode >= 400 {
		return fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	return nil
}
