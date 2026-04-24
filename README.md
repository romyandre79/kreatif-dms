# Kreatif DMS (Document Management System)

A premium, microservice-based Document Management System with AI-powered search, OCR, and robust security.

## 🏗️ Architecture
- **Backend**: Go (Fiber v3) - [Docs](./backend/README.md)
- **Frontend**: Nuxt 4 & Tailwind CSS v4 - [Docs](./frontend/README.md)
- **OCR Service**: Python (FastAPI + PaddleOCR) - [Docs](./ocr-service/README.md)
- **Infrastructure**: Docker Compose (PostgreSQL, Redis, MinIO, Elasticsearch)

## 🚀 Quick Start (Docker)
Ensure you have Docker and Docker Compose installed.

```bash
# Start all services
docker-compose up -d
```

## 🛠️ Components
1. **API Server**: Handles business logic, authentication, and document metadata.
2. **Task Worker**: Processes asynchronous tasks like OCR and indexing.
3. **OCR Engine**: Extracts text from images and PDFs using AI.
4. **Web UI**: Modern dashboard for managing and searching documents.

## 🔐 Key Features
- **AES256 Encryption** on MinIO.
- **Full-Text Search** across all documents.
- **Dynamic PDF Watermarking**.
- **Role-Based Access Control (RBAC)**.
- **AI-Powered Summarization** (Gemini/OpenAI Integration).
- **Batch Processing Tracking**: Real-time progress monitoring for multiple uploads.
- **System Observability**: Integrated logging across all services.
- **Rate Limiting & Security**: Configurable API protection.
