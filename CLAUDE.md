# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**Kreatif DMS** is a premium, microservice-based Document Management System with AI-powered search, OCR, and robust security. The system manages document workflows including intake, storage, loans, approvals, and retention across a hierarchical organizational structure (Company → Branch → Department → Rack → Box → Folder).

### Core Architecture

```
kreatif-dms/
├── backend/           # Go (Fiber v3) - Main API & business logic
├── frontend/          # Nuxt 4 + Tailwind CSS v4 - Web UI
├── ocr-service/       # Python (FastAPI + PaddleOCR) - OCR microservice
├── scannerbridge/     # Go - Windows scanner hardware bridge
├── mobile/            # Flutter - Mobile app (iOS/Android)
└── docker-compose.yml # Infrastructure (PostgreSQL, Redis, MinIO, Elasticsearch)
```

### Key Technologies

**Backend Stack:**
- **Framework**: Fiber v3 (high-performance Go web framework)
- **Database**: PostgreSQL with pgx v5 + sqlc for type-safe queries
- **Migrations**: golang-migrate
- **Authentication**: JWT (golang-jwt v5) with RBAC
- **Async Jobs**: Asynq (Redis-backed queue)
- **File Storage**: MinIO with AES256 encryption
- **Search**: Elasticsearch v8.12.0
- **Caching**: Redis v7
- **API Docs**: Swagger/Swag

**Frontend Stack:**
- **Framework**: Nuxt 4 with SSR disabled (Windows optimization)
- **Styling**: Tailwind CSS v4 with PostCSS
- **State Management**: Pinia
- **Security**: nuxt-security with CSP headers
- **Internationalization**: @nuxtjs/i18n (English, Indonesian)
- **Icons**: lucide-vue-next
- **Quality**: ESLint + Prettier

**OCR Microservice:**
- **Framework**: FastAPI
- **Engine**: PaddleOCR (Indonesian & English support)
- **File Support**: Images (JPG, PNG) and multi-page PDFs
- **Output**: Word coordinates + full text aggregation

**Infrastructure:**
- **Container Orchestration**: Docker Compose
- **Database**: PostgreSQL 16 Alpine
- **Cache**: Redis 7 Alpine
- **Object Storage**: MinIO
- **Search Engine**: Elasticsearch 8.12.0

---

## Development Workflow

### Prerequisites

- **Docker Desktop** - for running PostgreSQL, Redis, MinIO, Elasticsearch
- **Go 1.25+** - backend development
- **Node.js 20+** - frontend development
- **Python 3.10+** - OCR service development
- **Git Bash or PowerShell** - terminal

### Quick Start

#### 1. Start Infrastructure (one-time setup)

```bash
# From root directory
docker-compose up -d
```

This starts:
- **PostgreSQL**: `localhost:5432` (user: postgres, db: kreatif_dms)
- **Redis**: `localhost:6379`
- **MinIO**: `localhost:9000` (console: `localhost:9001`)
- **Elasticsearch**: `localhost:9200`

#### 2. Backend Setup

```bash
cd backend

# Copy environment (adjust values as needed)
cp .env.example .env

# Run database migrations
go run cmd/server/main.go migrate up

# Start API server
go run cmd/server/main.go
# API runs at http://localhost:8080
# Swagger docs at http://localhost:8080/swagger/
```

#### 3. Backend Worker (required for async tasks)

In a separate terminal:

```bash
cd backend

# Background jobs: OCR processing, document indexing, batch operations
go run cmd/worker/main.go
```

#### 4. Frontend Setup

```bash
cd frontend

# Install dependencies
npm install

# Development server (with HMR)
npm run dev
# Dashboard at http://localhost:3000
```

#### 5. OCR Service (optional)

```bash
cd ocr-service

# Create virtual environment
python -m venv venv

# Activate (Windows)
.\venv\Scripts\activate
# Or (Linux/Mac)
source venv/bin/activate

# Install dependencies
pip install -r requirements.txt

# Run service
python main.py
# OCR API at http://localhost:8000/docs
```

### Common Development Commands

**Using Makefile (from root):**

```bash
make up              # Start Docker containers
make down            # Stop containers
make dev-back        # Run backend in development
make dev-front       # Run frontend in development
make sqlc            # Regenerate Go code from SQL queries
make migrate-up      # Run database migrations
make migrate-down    # Rollback database migrations
make test            # Run all Go tests
make swagger-gen     # Generate Swagger documentation
```

**Backend Commands:**

```bash
# Development with environment variable
cd backend
set APP_ENV=development && go run cmd/server/main.go

# Production build
go build -o bin/server cmd/server/main.go

# Run tests
go test -v ./...

# Run specific test file
go test -v ./internal/auth -run TestJWT

# Database migrations
go run cmd/server/main.go migrate up
go run cmd/server/main.go migrate down
go run cmd/server/main.go migrate goto 15  # Go to specific version

# Generate code from SQL
sqlc generate

# Regenerate Swagger docs
swag init -g cmd/server/main.go
```

**Frontend Commands:**

```bash
cd frontend

# Development (with HMR, file polling enabled for Windows)
npm run dev

# Production build
npm run build

# Preview production build
npm run preview

# Code generation (Nuxt auto-imports)
npm run postinstall

# Linting
npm run lint
```

**OCR Service Commands:**

```bash
cd ocr-service

# Development
python main.py

# Docker build
docker build -t ocr-service .
docker run -p 8000:8000 ocr-service
```

---

## Architecture Details

### Backend Structure

**cmd/** - Entry points
- `cmd/server/main.go` - API server with Fiber v3 setup, routing, middleware
- `cmd/worker/main.go` - Background job processor (Asynq)

**internal/** - Core application logic
- `auth/` - JWT token generation, parsing, validation
- `config/` - Viper-based configuration loading, database/redis/storage initialization
- `handler/` - HTTP request handlers (11 domain areas):
  - `auth_handler.go` - Login, register, token refresh, password reset
  - `document_handler.go` - Upload, download, preview, watermarking, metadata
  - `master_handler.go` - Master data CRUD (companies, branches, departments, racks, boxes, orderers)
  - `intake_handler.go` - Document intake workflow (barcode generation, approval)
  - `loan_handler.go` - Document loan/circulation workflow with tracking
  - `dashboard_handler.go` - Dashboard statistics and KPIs
  - `user_handler.go` - User profile, password, MFA settings
  - `batch_handler.go` - Batch upload progress tracking
  - `notification_handler.go` - Notification management
  - `email_template_handler.go` - Email template CRUD
  - `hardware_handler.go` - Hardware device management (RFID/NFC scanners)

- `infra/` - External service integrations
  - `ai.go` - AI service (Gemini, OpenAI, Ollama integration)
  - `storage.go` - MinIO client initialization & file operations
  - `search.go` - Elasticsearch indexing and full-text search
  - `mail.go` - SMTP/Mailpit email integration
  - `email.go` - Email template rendering
  - `ldap.go` - LDAP/Active Directory authentication
  - `whatsapp.go` - WhatsApp gateway integration

- `middleware/` - Fiber middleware
  - `auth.go` - JWT token validation, role-based access control (RBAC)

- `model/` - Data structures (JSON request/response models)

- `notification/` - Notification business logic

- `repository/` - **Database layer (generated from sqlc)**
  - 50+ generated `.sql.go` files from `db/queries/`
  - Type-safe database operations using pgx v5

- `service/` - Business logic layer
  - `auth_service.go` - Authentication, token, password management
  - `document_service.go` - Document lifecycle, versioning, watermarking
  - `intake_service.go` - Document intake workflow logic
  - `loan_service.go` - Loan approval, circulation tracking
  - `master_service.go` - Master data management
  - `dashboard_service.go` - Analytics and reporting
  - `cache_service.go` - Redis caching layer
  - `integration_monitor_service.go` - Health checks for external services

- `worker/` - Background job handlers for Asynq

**pkg/** - Shared utilities
- `database/` - Database connection pooling, migration runner
- `response/` - Standardized API response formatting
- `utils/` - Helper functions

**db/**
- `migrations/` - 22+ SQL migration files (up/down) managed by golang-migrate
  - Schema initialization, seed data, ACL setup, system modules, user fields, etc.
- `queries/` - SQL query files processed by sqlc
  - 13 domain query files covering users, documents, loans, intake, warehouse, etc.

### Frontend Structure

**app/pages/** - Route-based pages (Nuxt 4 file-based routing)
- `index.vue` - Landing page
- `login.vue` - Authentication page
- `dashboard.vue` - Main dashboard
- `admin/` - Admin panel (user management, roles, permissions)
- `approvals/` - Document approval workflows
- `circulation/` - Document circulation/loans
- `config/` - System configuration
- `documents/` - Document browser, search, preview
- `intake/` - Document intake form and workflow
- `loans/` - Loan management
- `notifications/` - Notification center
- `profile/` - User profile management
- `registration/` - New document registration
- `retention/` - Document retention policies
- `scan/` - RFID/NFC scanner interface
- `settings/` - User settings
- `stock/` - Inventory management
- `tracking/` - Document tracking
- `warehouse/` - Warehouse/physical location management

**app/components/** - Reusable Vue 3 components
- Global: Alert, Modal, Dropdown, SearchableSelect, SignaturePad, FloatingCart
- `dashboard/` - Dashboard-specific components
- `integration/` - Integration UI components

**app/stores/** - Pinia state management
- `auth.ts` - Authentication state (token, user, roles)
- `cart.ts` - Document cart for batch operations
- `notifications.ts` - Notification state
- `toast.ts` - Toast notification state

**app/composables/** - Vue 3 composition API utilities
- `useApi.ts` - Axios wrapper with auto token injection
- `useToast.ts` - Toast notification composable

**app/utils/**
- `logger.ts` - Color-coded logger for debugging
- `format.ts` - Date, number, file size formatting

**app/middleware/** - Route guards for authentication

**app/locales/** - i18n translations
- `en/` & `id/` - English and Indonesian

### Database Schema Highlights

- **Users & Security**: users, roles, permissions, user_approvals, mfa_settings
- **Organizational Structure**: companies, branches, departments
- **Physical Storage**: racks, boxes, orderers (shelving hierarchy)
- **Documents**: documents, document_versions, document_metadata, documents_acl
- **Workflows**: intakes, loans, approvals
- **OCR**: ocr_jobs, ocr_corrections
- **Notifications**: notifications, email_templates
- **Search**: Elasticsearch integration
- **System**: activity_logs, announcements, integration_nodes, system_modules, system_permissions

---

## Key Patterns & Concepts

### Role-Based Access Control (RBAC)

Five role tiers: Admin, Manager DC, Doc Controller, Manager, User, Guest with hierarchical department/branch permissions and document-level ACL (unauthorized users see "**RAHASIA**" masking).

### Document Workflow

Intake → Approval → Storage (encrypted MinIO + PostgreSQL metadata) → Indexing (Elasticsearch) → OCR (optional PaddleOCR) → Loan (circulation tracking) → Retention (scheduled deletion).

### Async Job Processing

Asynq (Redis queue) handles OCR processing, Elasticsearch indexing, email notifications, batch operations, retention cleanup with configurable concurrency (default: 5).

### File Encryption & Security

- MinIO files encrypted with AES256 (STORAGE_ENCRYPTION_KEY)
- PDF watermarking with dynamic overlay
- JWT tokens with Redis blacklist for logout

### Environment Configuration

Viper loads `.env` with support for development/staging/production variants. Key variables: database URL, Redis URL, MinIO credentials, JWT secrets, SMTP/LDAP/WhatsApp integration settings.

---

## Deployment Environments

- **Development**: `.env`, HTTP, sourcemaps enabled
- **Staging**: `.env.staging`, HTTPS, optimized builds
- **Production**: `.env.production`, hardened security, minified assets

---

## Testing

```bash
cd backend
go test -v ./...                    # All tests
go test -v ./internal/auth          # Specific package
go test -v -run TestJWT ./...       # Match pattern
```

Test files: `internal/auth/jwt_test.go`, `internal/middleware/auth_test.go`, `internal/service/auth_service_test.go`

---

## Key Files to Understand First

- `backend/cmd/server/main.go` - Server initialization & routing
- `backend/internal/config/config.go` - Configuration structure
- `backend/internal/handler/document_handler.go` - Core document flow
- `backend/internal/service/document_service.go` - Business logic
- `frontend/nuxt.config.ts` - Frontend framework setup
- `frontend/app/stores/auth.ts` - Auth state management
- `frontend/design.md` - **Frontend UI/UX style guide** (colors, typography, components, spacing, dark mode — WAJIB dibaca sebelum mengerjakan frontend)
- `docker-compose.yml` - Infrastructure definition
- `PROJECT_PLAN.md` - High-level requirements
- `DETAIL_PROJECT.md` - Per-feature implementation status (frontend, backend, database)

---

## Windows-Specific Notes

- Nuxt HMR configured with WebSocket + file polling (1000ms interval) in `nuxt.config.ts`
- Devtools & sourcemaps disabled to reduce RAM usage
- Backend scanner bridge (`scannerbridge/main.exe`) provides Windows hardware integration
- Set environment variables with `set VAR=value` in PowerShell/CMD, not `export`

---

## Code Standards

- **Backend**: Clean Architecture (handler → service → repository)
- **Frontend**: Vue 3 Composition API + TypeScript
- **Naming**: camelCase for code, snake_case for DB columns
- **Commits**: Conventional commits (feat:, fix:, docs:, etc.)
- **Linting**: ESLint in frontend, Golang conventions in backend

