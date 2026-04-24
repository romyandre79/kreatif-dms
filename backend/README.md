# Kreatif DMS - Backend

Backend service for Kreatif Document Management System built with Go and Fiber v3.

## 🚀 Features
- **Framework**: Fiber v3 (High performance)
- **Security**:
    - **Encryption**: AES256 (SSE-S3) for file storage in MinIO.
    - **Authentication**: JWT with Role-Based Access Control (RBAC).
    - **Security Headers**: Nuxt-like security headers via Helmet.
    - **Rate Limiting**: Protect against brute force.
    - **Input Validation**: Strict payload validation.
- **Documents**:
    - **PDF Watermarking**: Dynamic real-time watermarking.
    - **Hierarchical Storage**: Organizational structure (Company > Branch > Dept).
- **Search**: Full-text search powered by Elasticsearch.
- **Async Workers**: Background OCR and indexing using Redis (Asynq).
- **Batch Processing**: Track real-time progress for multiple document uploads.
- **Observability**: Detailed internal logging for all services and activities.
- **Mail Service**: SMTP integration for notifications (Mailpit-ready).
- **Database**: PostgreSQL with `sqlc` for type-safe queries and `golang-migrate`.

## 🛠️ Requirements
- Go 1.25+
- PostgreSQL
- Redis
- MinIO
- Elasticsearch

## 🚦 How to Run

### 1. Configuration
Copy `.env` and adjust the values:
```bash
cp .env.example .env
```

Key `.env` Toggles:
- `APP_PREFORK`: Enable high-performance multi-process mode (Linux).
- `RATE_LIMIT_MAX`: Max requests allowed within the expiration window.
- `RATE_LIMIT_EXPIRATION`: Window duration (e.g., `1m`, `30s`).
- `SMTP_*`: Configuration for local/production mail server.

### 2. Database Migration
```bash
# Run migrations up
go run cmd/server/main.go migrate up

# Run migrations down
go run cmd/server/main.go migrate down
```

### 3. Run Server
```bash
go run cmd/server/main.go
```

### 4. Run Worker
```bash
go run cmd/worker/main.go
```

## 📚 API Documentation
- **Swagger**: `http://localhost:8080/swagger/*`
- **Landing Page**: `http://localhost:8080/` (Status monitor)
- **Health Check**: `http://localhost:8080/health`
