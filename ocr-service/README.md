# Kreatif DMS - OCR Service

Microservice for Optical Character Recognition (OCR) using PaddleOCR and FastAPI.

## 🚀 Features
- **Engine**: PaddleOCR (Support for ID/EN).
- **Format Support**: Image (JPG, PNG) and Multi-page PDF.
- **Output**: Detailed word coordinates and aggregated `full_text`.
- **FastAPI**: Modern, high-performance Python API.
- **Logging**: Integrated request and error tracking.
- **Observability**: Built-in landing page and health monitoring.

## 🛠️ Requirements
- Python 3.10+
- (Optional) Poppler-utils (if using pdf2image)

## 🚦 How to Run (Locally)

### 1. Setup Virtual Environment
```bash
python -m venv venv
# Windows
.\venv\Scripts\activate
# Linux/Mac
source venv/bin/activate
```

### 2. Install Dependencies
```bash
pip install -r requirements.txt
```

### 3. Run Service
```bash
python main.py
```
The service will be available at `http://localhost:8000`.

## 🚦 How to Run (Docker)
```bash
docker build -t ocr-service .
docker run -p 8000:8000 ocr-service
```

## 📡 API Documentation
- **Swagger UI**: `http://localhost:8000/docs` (Interactive testing)
- **Redoc**: `http://localhost:8000/redoc`
- **Landing Page**: `http://localhost:8000/`
- **Health Check**: `http://localhost:8000/health`

## 📡 API Endpoints
- `POST /ocr/process`: Process file (Image/PDF).
