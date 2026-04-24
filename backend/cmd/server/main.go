package main

import (
	"log"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/hibiken/asynq"
	_ "github.com/kreatif/dms-backend/docs"
	"github.com/kreatif/dms-backend/internal/config"
	"github.com/kreatif/dms-backend/internal/handler"
	"github.com/kreatif/dms-backend/internal/infra"
	"github.com/kreatif/dms-backend/internal/middleware"
	"github.com/kreatif/dms-backend/internal/repository"
	"github.com/kreatif/dms-backend/internal/service"
	"github.com/kreatif/dms-backend/pkg/database"
)

// @title Kreatif DMS API
// @version 1.0
// @description API for Document Management System with AI and RFID support.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email romy@kreatif.id

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// Load config
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// Handle Subcommands (Artisan-like)
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			cmd := "up"
			if len(os.Args) > 2 {
				cmd = os.Args[2]
			}
			database.RunMigrations(cfg.DatabaseURL, "db/migrations", cmd)
			return
		}
	}

	// Initialize Infrastructure
	dbPool := config.InitDatabase(cfg.DatabaseURL)
	defer dbPool.Close()

	rdb := config.InitRedis(cfg.RedisURL)
	defer rdb.Close()

	asynqClient := asynq.NewClient(asynq.RedisClientOpt{Addr: cfg.RedisURL})
	defer asynqClient.Close()

	minioClient := config.InitMinIO(cfg.MinIOEndpoint, cfg.MinIOAccessKey, cfg.MinIOSecretKey, cfg.MinIOUseSSL, cfg.MinIOBucket)
	_ = minioClient // For later use

	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{cfg.ElasticsearchURL},
	})
	if err != nil {
		log.Fatalf("ES client error: %v", err)
	}

	// Repositories
	repo := repository.New(dbPool)

	// Services
	storageSvc := infra.NewStorageService(minioClient, cfg.MinIOBucket)
	ldapSvc := infra.NewLDAPService(cfg)
	emailSvc := infra.NewEmailService(cfg)
	authSvc := service.NewAuthService(repo, cfg, ldapSvc, emailSvc)
	docSvc := service.NewDocumentService(repo, storageSvc, asynqClient)
	searchSvc := infra.NewSearchService(es)
	_ = service.NewCacheService(rdb) // Initialized for performance later

	// Handlers
	authHandler := handler.NewAuthHandler(authSvc)
	userHandler := handler.NewUserHandler(repo)
	docHandler := handler.NewDocumentHandler(docSvc, searchSvc)
	batchHandler := handler.NewBatchHandler(docSvc)

	// Create Fiber App
	app := fiber.New(fiber.Config{
		AppName:      "Kreatif DMS API",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  60 * time.Second,
		BodyLimit:    100 * 1024 * 1024, // 100MB Limit
	})

	// Middlewares
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // Adjust this to your frontend URL in production
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))
	app.Use(helmet.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(limiter.New(limiter.Config{
		Max:               cfg.RateLimitMax,
		Expiration:        cfg.RateLimitExpiration,
		LimiterMiddleware: limiter.FixedWindow{},
	}))

	// Routes
	api := app.Group("/api/v1")

	// Auth Routes
	authGroup := api.Group("/auth")
	authGroup.Post("/login", authHandler.Login)
	authGroup.Post("/register", authHandler.Register)
	authGroup.Post("/forgot-password", authHandler.ForgotPassword)

	// User Routes
	userGroup := api.Group("/users")
	userGroup.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	userGroup.Get("/", middleware.RoleMiddleware("admin"), userHandler.List)
	userGroup.Get("/pending", middleware.RoleMiddleware("admin"), authHandler.ListPendingUsers)
	userGroup.Post("/:id/approve", middleware.RoleMiddleware("admin"), authHandler.ApproveUser)

	// Document Routes
	docGroup := api.Group("/documents")
	docGroup.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	docGroup.Post("/", docHandler.Upload)
	docGroup.Get("/:id/preview", docHandler.Preview)
	docGroup.Get("/search", docHandler.Search)

	// Batch Routes
	batchGroup := api.Group("/batches")
	batchGroup.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	batchGroup.Post("/", batchHandler.Create)
	batchGroup.Get("/:id", batchHandler.GetStatus)

	// Health check
	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// Root Landing Page
	app.Get("/", func(c fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.SendString(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Kreatif DMS API - System Online</title>
				<style>
					@import url('https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600&display=swap');
					body {
						margin: 0;
						padding: 0;
						font-family: 'Outfit', sans-serif;
						background-color: #0f172a;
						color: #f8fafc;
						display: flex;
						justify-content: center;
						align-items: center;
						height: 100vh;
						overflow: hidden;
					}
					.container {
						text-align: center;
						position: relative;
						z-index: 10;
						padding: 2rem;
						background: rgba(30, 41, 59, 0.5);
						backdrop-filter: blur(20px);
						border: 1px solid rgba(255, 255, 255, 0.1);
						border-radius: 2rem;
						box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5);
					}
					.glow {
						position: absolute;
						width: 300px;
						height: 300px;
						background: radial-gradient(circle, rgba(59, 130, 246, 0.3) 0%, rgba(59, 130, 246, 0) 70%);
						top: 50%;
						left: 50%;
						transform: translate(-50%, -50%);
						z-index: 1;
					}
					h1 {
						font-size: 3rem;
						font-weight: 600;
						background: linear-gradient(to right, #60a5fa, #a78bfa);
						-webkit-background-clip: text;
						-webkit-text-fill-color: transparent;
						margin-bottom: 1rem;
					}
					p {
						font-size: 1.1rem;
						color: #94a3b8;
						margin-bottom: 2rem;
					}
					.status {
						display: inline-flex;
						align-items: center;
						background: rgba(34, 197, 94, 0.1);
						color: #4ade80;
						padding: 0.5rem 1rem;
						border-radius: 9999px;
						font-size: 0.875rem;
						font-weight: 600;
						border: 1px solid rgba(34, 197, 94, 0.2);
					}
					.dot {
						width: 8px;
						height: 8px;
						background-color: #22c55e;
						border-radius: 50%;
						margin-right: 0.5rem;
						box-shadow: 0 0 10px #22c55e;
						animation: pulse 2s infinite;
					}
					@keyframes pulse {
						0% { transform: scale(1); opacity: 1; }
						50% { transform: scale(1.5); opacity: 0.5; }
						100% { transform: scale(1); opacity: 1; }
					}
					.links {
						margin-top: 2rem;
						display: flex;
						gap: 1rem;
						justify-content: center;
					}
					a {
						color: #60a5fa;
						text-decoration: none;
						font-size: 0.875rem;
						transition: color 0.2s;
					}
					a:hover {
						color: #93c5fd;
					}
				</style>
			</head>
			<body>
				<div class="glow"></div>
				<div class="container">
					<div class="status">
						<span class="dot"></span>
						SYSTEM OPERATIONAL
					</div>
					<h1>Kreatif DMS API</h1>
					<p>The core engine for premium document management.</p>
					<div class="links">
						<a href="/swagger/index.html">API Documentation</a>
						<a href="/health">Health Status</a>
					</div>
				</div>
			</body>
			</html>
		`)
	})

	// Swagger documentation
	app.Get("/swagger/*", swaggo.HandlerDefault)

	// Start server
	log.Printf("Server starting on port %s", cfg.AppPort)
	log.Fatal(app.Listen(":"+cfg.AppPort, fiber.ListenConfig{
		EnablePrefork: cfg.AppPrefork,
	}))
}
