import os
# Set environment variables BEFORE any other imports
os.environ["KMP_DUPLICATE_LIB_OK"] = "TRUE"

import asyncio
import io
import time
import secrets
import json
import logging
import numpy as np
from datetime import datetime
from PIL import Image
import pypdfium2 as pdfium
import uvicorn
from concurrent.futures import ThreadPoolExecutor
from contextlib import asynccontextmanager

from fastapi import FastAPI, UploadFile, File, Form, HTTPException, Request, Depends, status, WebSocket, WebSocketDisconnect
from fastapi.responses import HTMLResponse
from fastapi.security import HTTPBasic, HTTPBasicCredentials
from fastapi.staticfiles import StaticFiles
from dotenv import load_dotenv

# Load environment variables BEFORE importing custom modules
load_dotenv()

# Import our custom modules
from db_logic import init_db, get_db, save_request_log, get_stats, get_recent_history, get_ocr_result_by_id
from ai_logic import AI_ENABLED, AI_VISION_ENABLED, ocr_with_ai_vision, analyze_with_ai
from ocr_engine import run_ocr, format_ocr_result, IMAGE_SCALE
from ws_manager import manager, broadcast_log, WebSocketLogHandler
from dashboard_html import DASHBOARD_HTML


# Configure logging
logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s] %(name)s: %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S"
)
logger = logging.getLogger("ocr-service")

# Setup WebSocket Logging
ws_handler = WebSocketLogHandler()
ws_handler.setFormatter(logging.Formatter("%(message)s"))
logger.addHandler(ws_handler)

@asynccontextmanager
async def lifespan(app: FastAPI):
    import ws_manager
    ws_manager.main_loop = asyncio.get_event_loop()
    logger.info("Kreatif DMS OCR Service started (Refactored Mode).")
    init_db()
    yield
    logger.info("Kreatif DMS OCR Service is shutting down.")

from fastapi.middleware.cors import CORSMiddleware

# Auth Config
ADMIN_USER = os.getenv("OCR_ADMIN_USER", "admin")
ADMIN_PASSWORD = os.getenv("OCR_ADMIN_PASSWORD", "admin123")
ALLOWED_CLIENT_HOSTS = [h.strip() for h in os.getenv("CLIENT_HOST", "http://localhost:3000").split(",")]
RATE_LIMIT = int(os.getenv("RATE_LIMITER", "100"))
MAX_PARALLEL_PAGES = int(os.getenv("MAX_PARALLEL_PAGES", "4"))

app = FastAPI(title="Kreatif DMS OCR Service", lifespan=lifespan)

# Enable CORS using settings from .env
app.add_middleware(
    CORSMiddleware,
    allow_origins=ALLOWED_CLIENT_HOSTS,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

security = HTTPBasic()

# Ensure uploads directory exists
UPLOAD_DIR = "uploads"
if not os.path.exists(UPLOAD_DIR):
    os.makedirs(UPLOAD_DIR)
app.mount("/uploads", StaticFiles(directory=UPLOAD_DIR), name="uploads")

def authenticate(credentials: HTTPBasicCredentials = Depends(security)):
    if not (secrets.compare_digest(credentials.username.encode("utf8"), ADMIN_USER.encode("utf8")) and 
            secrets.compare_digest(credentials.password.encode("utf8"), ADMIN_PASSWORD.encode("utf8"))):
        raise HTTPException(status_code=401, detail="Unauthorized", headers={"WWW-Authenticate": "Basic"})
    return credentials.username

def check_rate_limit(request: Request):
    host = request.client.host
    now = time.time()
    with get_db() as conn:
        cursor = conn.cursor()
        cursor.execute("DELETE FROM rate_limits WHERE timestamp < ?", (now - 3600,))
        cursor.execute("SELECT COUNT(*) FROM rate_limits WHERE client_host = ?", (host,))
        if cursor.fetchone()[0] >= RATE_LIMIT:
            raise HTTPException(status_code=429, detail="Rate limit exceeded")
        cursor.execute("INSERT INTO rate_limits (client_host, timestamp) VALUES (?, ?)", (host, now))
        conn.commit()
    return True

@app.get("/", response_class=HTMLResponse, tags=["UI"])
async def dashboard(username: str = Depends(authenticate)):
    return DASHBOARD_HTML

@app.get("/ocr/stats")
async def get_stats_api(page: int = 1, page_size: int = 10, username: str = Depends(authenticate)):
    from db_logic import get_total_history_count
    total_logs, dur, failed = get_stats()
    
    # Pagination logic
    total_items = get_total_history_count()
    offset = (page - 1) * page_size
    history_raw = get_recent_history(limit=page_size, offset=offset)
    
    history = []
    for row in history_raw:
        history.append({
            "id": row[0], "start_time": row[1], "end_time": row[2],
            "filename": row[3], "size": row[4], "duration": row[5],
            "status": row[6], "accuracy": row[7], "ai_analysis": row[8]
        })
    
    return {
        "stats": {
            "total": total_logs, 
            "avg_time": dur/total_logs if total_logs > 0 else 0, 
            "success_rate": (total_logs-failed)/total_logs*100 if total_logs > 0 else 100
        },
        "history": history,
        "pagination": {
            "total_items": total_items,
            "page": page,
            "page_size": page_size,
            "total_pages": (total_items + page_size - 1) // page_size
        }
    }

@app.get("/ocr/result/{job_id}")
async def get_result(job_id: int, username: str = Depends(authenticate)):
    row = get_ocr_result_by_id(job_id)
    if not row: raise HTTPException(status_code=404, detail="Not found")
    return {
        "id": row[0], "timestamp": row[1], "filename": row[2], "status": row[3],
        "words_json": row[4], "file_path": row[5], "preview_path": row[6],
        "ai_analysis": row[7], "preview_paths": json.loads(row[8]) if row[8] else [row[6]]
    }

@app.websocket("/ws/logs")
async def websocket_endpoint(websocket: WebSocket):
    await manager.connect(websocket)
    try:
        while True: await websocket.receive_text()
    except WebSocketDisconnect:
        manager.disconnect(websocket)

@app.post("/ocr/process")
async def process_ocr(
    file: UploadFile = File(...), 
    ai_model: str = Form(None), 
    ai_api_key: str = Form(None),
    username: str = Depends(authenticate), 
    rate_ok: bool = Depends(check_rate_limit)
):
    logger.info(f"Received upload request: {file.filename} | AI Model: {ai_model}")
    await broadcast_log(f"Received upload request: {file.filename} (AI: {ai_model or 'Default'})", "info")
    start_time_str = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    start_time_val = time.time()
    filename = file.filename
    content = await file.read()
    timestamp = int(time.time())
    save_filename = f"{timestamp}_{filename}"
    file_path = os.path.join(UPLOAD_DIR, save_filename)
    with open(file_path, "wb") as f: f.write(content)
    
    results = []
    full_text_parts = []
    preview_paths = []
    preview_path = f"/uploads/{save_filename}"

    try:
        if filename.lower().endswith('.pdf'):
            pdf = pdfium.PdfDocument(content)
            for i in range(len(pdf)):
                pname = f"preview_{timestamp}_{filename}_p{i+1}.jpg"
                ppath = os.path.join(UPLOAD_DIR, pname)
                pdf[i].render(scale=IMAGE_SCALE).to_pil().save(ppath)
                preview_paths.append(f"/uploads/{pname}")
            preview_path = preview_paths[0]
            
            def p_page(idx):
                img = np.array(pdf[idx].render(scale=IMAGE_SCALE).to_pil())
                return format_ocr_result(run_ocr(img), idx + 1)

            with ThreadPoolExecutor(max_workers=MAX_PARALLEL_PAGES) as exe:
                paged_data = list(exe.map(p_page, range(len(pdf))))
            
            for i, ext in enumerate(paged_data):
                results.extend(ext["words"])
                full_text_parts.append(ext["full_text"])
                await broadcast_log(f"Page {i+1} completed: {len(ext['words'])} text blocks", "success")
        else:
            await broadcast_log(f"Processing Image: {filename}", "info")
            img = np.array(Image.open(io.BytesIO(content)).convert('RGB'))
            ext = format_ocr_result(run_ocr(img), 1)
            results.extend(ext["words"])
            full_text_parts.append(ext["full_text"])
            await broadcast_log(f"Image scan completed: {len(ext['words'])} blocks", "success")

        # AI Analysis
        all_text = " ".join(full_text_parts)
        ai_result = None
        if AI_ENABLED:
            vision_text = None
            if AI_VISION_ENABLED:
                v_img = content if not filename.lower().endswith('.pdf') else io.BytesIO()
                if filename.lower().endswith('.pdf'): 
                    pdf[0].render(scale=IMAGE_SCALE).to_pil().save(v_img, format='JPEG')
                    v_img = v_img.getvalue()
                vision_text = await ocr_with_ai_vision(v_img, model=ai_model, api_key=ai_api_key)
            ai_result = await analyze_with_ai(all_text, vision_text, model=ai_model, api_key=ai_api_key)

        duration = time.time() - start_time_val
        avg_acc = sum(w['confidence'] for w in results) / len(results) if results else 0
        
        save_request_log({
            "start_time": start_time_str, "end_time": datetime.now().strftime("%Y-%m-%d %H:%M:%S"),
            "filename": filename, "size": f"{len(content)/1024:.1f} KB", "duration": duration,
            "status": "Success", "words_json": json.dumps(results), "file_path": f"/uploads/{save_filename}",
            "preview_path": preview_path, "accuracy": avg_acc, "ai_analysis": json.dumps(ai_result),
            "preview_paths": json.dumps(preview_paths if preview_paths else [preview_path])
        })
        
        await manager.broadcast("REFRESH_HISTORY")
        return {
            "status": "Success", 
            "filename": filename, 
            "insight": ai_result, 
            "preview_path": preview_path, 
            "preview_paths": preview_paths or [preview_path], 
            "words": results
        }

    except Exception as e:
        logger.error(f"OCR Error: {str(e)}")
        raise HTTPException(status_code=500, detail=str(e))

if __name__ == "__main__":
    port = int(os.getenv("APP_PORT", 8000))
    uvicorn.run(app, host="0.0.0.0", port=port)
