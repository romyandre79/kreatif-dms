package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/hibiken/asynq"
	"github.com/kreatif/dms-backend/internal/config"
	"github.com/kreatif/dms-backend/internal/infra"
	"github.com/kreatif/dms-backend/internal/repository"
	"github.com/kreatif/dms-backend/internal/worker"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	dbPool, err := config.InitDatabase(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer dbPool.Close()

	minioClient, err := config.InitMinIO(cfg.MinIOEndpoint, cfg.MinIOAccessKey, cfg.MinIOSecretKey, cfg.MinIOUseSSL, cfg.MinIOBucket)
	if err != nil && cfg.MinIOEndpoint != "" {
		log.Printf("WARNING: Worker starting without MinIO: %v", err)
	}
	
	// Repositories
	repo := repository.New(dbPool)

	// 4. Elasticsearch (Priority: DB -> .env)
	esURL := strings.TrimSpace(cfg.ElasticsearchURL)
	var esUser, esPass, esAPIKey string

	// Try "SEARCH" type first, then "ELASTICSEARCH"
	esNode, err := repo.GetIntegrationNodeByType(context.Background(), "SEARCH")
	if err != nil {
		esNode, err = repo.GetIntegrationNodeByType(context.Background(), "ELASTICSEARCH")
	}

	if err == nil && esNode.IsActive.Bool {
		esURL = strings.TrimSpace(esNode.Endpoint)
		log.Printf("[Worker] Search Engine: Using configuration from database node (%s): %s", esNode.ServiceType, esURL)

		var esCfg struct {
			Username string `json:"username"`
			Password string `json:"password"`
			APIKey   string `json:"api_key"`
		}
		json.Unmarshal(esNode.ConfigJson, &esCfg)
		esUser = esCfg.Username
		esPass = esCfg.Password
		esAPIKey = esCfg.APIKey
	}

	if esURL != "" && !strings.HasPrefix(esURL, "http://") && !strings.HasPrefix(esURL, "https://") {
		esURL = "http://" + esURL
	}

	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{esURL},
		Username:  esUser,
		Password:  esPass,
		APIKey:    esAPIKey,
	})
	if err != nil {
		log.Fatalf("ES client error: %v", err)
	}

	// Services
	storageSvc := infra.NewStorageService(cfg, repo, minioClient, cfg.MinIOBucket)
	aiSvc := infra.NewAIService(repo, cfg)
	searchSvc := infra.NewSearchService(es)

	// Processor
	processor := worker.NewTaskProcessor(repo, storageSvc, aiSvc, searchSvc)

	// Asynq Server (Redis Configuration Priority: DB -> .env -> Default)
	redisURL := cfg.RedisURL
	node, err := repo.GetIntegrationNodeByType(context.Background(), "REDIS")
	if err == nil && node.IsActive.Bool {
		redisURL = node.Endpoint
		log.Printf("[Worker] Redis: Using configuration from database node: %s", redisURL)
	} else if redisURL != "" {
		log.Printf("[Worker] Redis: Using configuration from .env: %s", redisURL)
	} else {
		redisURL = "127.0.0.1:6379" // Absolute fallback
		log.Printf("[Worker] Redis: No configuration found in DB or .env, using default: %s", redisURL)
	}

	concurrency := cfg.WorkerConcurrency
	if concurrency < 1 {
		concurrency = 5
	}

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisURL},
		asynq.Config{Concurrency: concurrency},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(worker.TypeDocumentOCR, processor.ProcessDocumentOCR)

	log.Println("Worker server starting...")
	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
