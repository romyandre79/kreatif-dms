package main

import (
	"context"
	"log"

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
	
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{cfg.ElasticsearchURL},
	})
	if err != nil {
		log.Fatalf("ES client error: %v", err)
	}

	// Repositories & Services
	repo := repository.New(dbPool)
	storageSvc := infra.NewStorageService(cfg, repo, minioClient, cfg.MinIOBucket)
	aiSvc := infra.NewAIService(repo, cfg)
	searchSvc := infra.NewSearchService(es)

	// Processor
	processor := worker.NewTaskProcessor(repo, storageSvc, aiSvc, searchSvc)

	// Asynq Server (Redis from DB)
	redisURL := cfg.RedisURL
	node, err := repo.GetIntegrationNodeByType(context.Background(), "REDIS")
	if err == nil && node.IsActive.Bool {
		redisURL = node.Endpoint
		log.Printf("Worker using Redis configuration from database: %s", redisURL)
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
