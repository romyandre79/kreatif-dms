# Kreatif DMS (Document Management System)

A premium, microservice-based Document Management System with AI-powered search, OCR, and robust security.

## 🏗️ Architecture
- **Backend**: Go (Fiber v3) - [Docs](./backend/README.md)
- **Frontend**: Nuxt 4 & Tailwind CSS v4 - [Docs](./frontend/README.md)
- **OCR Service**: Python (FastAPI + PaddleOCR) - [Docs](./ocr-service/README.md)
- **Infrastructure**: Docker Compose (PostgreSQL, Redis, MinIO, Elasticsearch)

## 🔐 Key Features
- **AES256 Encryption** on MinIO.
- **Full-Text Search** across all documents.
- **Dynamic PDF Watermarking**.
- **Role-Based Access Control (RBAC)**.
- **AI-Powered Summarization** (Gemini/OpenAI Integration).
- **Batch Processing Tracking**: Real-time progress monitoring for multiple uploads.
- **System Observability**: Integrated logging across all services.
- **Rate Limiting & Security**: Configurable API protection.

## 🚀 How To Run (Detailed)

### 1. Prasyarat
- **Docker Desktop** (untuk Database & Storage)
- **Go 1.22+**
- **Node.js 20+**
- **Git Bash / Terminal**

### 2a. Jalankan Infrastruktur (Docker)
Gunakan perintah ini untuk menjalankan database tanpa membebani RAM dengan build aplikasi:
```bash
docker-compose up -d
```
Service yang berjalan:
- **PostgreSQL**: `localhost:5432` (User/DB: `postgres/kreatif_dms`)
- **Redis**: `localhost:6379`
- **MinIO**: `localhost:9000` (Console: `localhost:9001`)
- **Elasticsearch**: `localhost:9200`

### 2b. Download dan Jalankan Terpisah Infrastruktur

Download MinIO:
https://min.io/download?os=windows&arch=amd64

Download PostgreSQL:
https://www.postgresql.org/download/

Download Redis:
https://redis.io/download/


Download ElasticSearch:
https://www.elastic.co/downloads/


### 3. Setup Backend
```bash
cd backend
# Salin environment (sesuaikan jika perlu)
cp .env.example .env
# Jalankan migrasi database
go run cmd/server/main.go migrate up
# Jalankan server
go run cmd/server/main.go
```
API akan berjalan di `http://localhost:8080`.

### 4. Setup Worker (Background Tasks)
Worker wajib dijalankan agar proses asinkronus (seperti OCR dan Indexing) bisa berjalan:
```bash
cd backend
# Jalankan worker
go run cmd/worker/main.go
```

### 5. Setup Frontend
```bash
cd frontend
# Install dependencies
npm install
# Jalankan development server
npm run dev
```
Dashboard akan berjalan di `http://localhost:3000`.

### 6. Setup OCR Service (Opsional)
Jika membutuhkan fitur OCR, jalankan service Python:
```bash
cd ocr-service
# Gunakan virtualenv
python -m venv venv
source venv/scripts/activate
pip install -r requirements.txt
python main.py
```
OCR akan berjalan di `http://localhost:8000`.