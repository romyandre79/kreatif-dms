import os
# Set environment variables BEFORE any other imports to ensure they are picked up
os.environ["KMP_DUPLICATE_LIB_OK"] = "TRUE"
os.environ["FLAGS_enable_mkldnn"] = "0"
os.environ["PADDLE_WITH_MKLDNN"] = "OFF"
os.environ["OMP_NUM_THREADS"] = "1"
os.environ["FLAGS_enable_pir_api"] = "0"
os.environ["FLAGS_enable_new_executor"] = "0"

from fastapi import FastAPI, UploadFile, File, HTTPException, Request, Depends, status
from fastapi.responses import HTMLResponse
from fastapi.security import HTTPBasic, HTTPBasicCredentials
import easyocr
import uvicorn
import io
import pypdfium2 as pdfium
from PIL import Image
import numpy as np
import logging
import time
import secrets
from datetime import datetime
from collections import deque
from dotenv import load_dotenv

# Load environment variables
load_dotenv()

# Configure logging
logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s] %(name)s: %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S"
)
logger = logging.getLogger("ocr-service")

app = FastAPI(title="Kreatif DMS OCR Service")
security = HTTPBasic()

# Configuration
ADMIN_USER = os.getenv("OCR_ADMIN_USER", "admin")
ADMIN_PASSWORD = os.getenv("OCR_ADMIN_PASSWORD", "admin123")

# Global State for Dashboard
request_history = deque(maxlen=100)
stats = {
    "total_processed": 0,
    "total_duration": 0,
    "failed_count": 0,
    "start_time": datetime.now().strftime("%Y-%m-%d %H:%M:%S")
}

# Initialize EasyOCR
logger.info("Initializing EasyOCR Engine (Stable for Python 3.13)...")
# Using Indonesian and English
reader = easyocr.Reader(['id', 'en'], gpu=False)
logger.info("EasyOCR Engine Ready.")

def authenticate(credentials: HTTPBasicCredentials = Depends(security)):
    current_username_bytes = credentials.username.encode("utf8")
    correct_username_bytes = ADMIN_USER.encode("utf8")
    is_correct_username = secrets.compare_digest(
        current_username_bytes, correct_username_bytes
    )
    current_password_bytes = credentials.password.encode("utf8")
    correct_password_bytes = ADMIN_PASSWORD.encode("utf8")
    is_correct_password = secrets.compare_digest(
        current_password_bytes, correct_password_bytes
    )
    if not (is_correct_username and is_correct_password):
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Incorrect username or password",
            headers={"WWW-Authenticate": "Basic"},
        )
    return credentials.username

@app.get("/", response_class=HTMLResponse, tags=["UI"])
async def dashboard(request: Request, username: str = Depends(authenticate)):
    avg_time = stats["total_duration"] / stats["total_processed"] if stats["total_processed"] > 0 else 0
    success_rate = ((stats["total_processed"] - stats["failed_count"]) / stats["total_processed"] * 100) if stats["total_processed"] > 0 else 100

    rows = ""
    for req in reversed(request_history):
        status_color = "text-green-500 bg-green-500/10" if req["status"] == "Success" else "text-red-500 bg-red-500/10"
        rows += f"""
        <tr class="border-b border-slate-800/50 hover:bg-slate-800/30 transition-colors">
            <td class="py-4 px-6 text-slate-400 font-mono text-[10px]">{req["timestamp"]}</td>
            <td class="py-4 px-6 font-bold text-slate-200">{req["filename"]}</td>
            <td class="py-4 px-6 text-slate-400 text-xs">{req["size"]}</td>
            <td class="py-4 px-6 font-black text-xs">{req["duration"]:.2f}s</td>
            <td class="py-4 px-6">
                <span class="px-3 py-1 rounded-full text-[9px] font-black uppercase tracking-widest {status_color}">
                    {req["status"]}
                </span>
            </td>
        </tr>
        """

    if not rows:
        rows = '<tr><td colspan="5" class="py-20 text-center text-slate-500 font-bold uppercase tracking-widest text-xs">No processing history yet</td></tr>'

    return f"""
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>OCR Service Dashboard | Kreatif DMS</title>
        <script src="https://cdn.tailwindcss.com"></script>
        <link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;900&display=swap" rel="stylesheet">
        <script src="https://unpkg.com/lucide@latest"></script>
        <style>
            body {{ font-family: 'Outfit', sans-serif; background-color: #020617; color: #f8fafc; }}
            .glass {{ background: rgba(15, 23, 42, 0.6); backdrop-filter: blur(12px); border: 1px solid rgba(255,255,255,0.05); }}
            .glow {{ box-shadow: 0 0 40px -10px rgba(56, 189, 248, 0.2); }}
        </style>
    </head>
    <body class="p-8 lg:p-12 min-h-screen">
        <div class="max-w-7xl mx-auto space-y-12">
            <!-- Header -->
            <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
                <div class="space-y-1">
                    <div class="flex items-center gap-3">
                        <div class="p-2 bg-sky-500/10 rounded-lg"><i data-lucide="scan-text" class="text-sky-500 w-6 h-6"></i></div>
                        <h1 class="text-3xl font-black tracking-tight uppercase">OCR Engine <span class="text-sky-500">Service</span></h1>
                    </div>
                    <p class="text-slate-500 text-sm font-medium">PaddleOCR High-Performance Intelligence Dashboard</p>
                </div>
                <div class="flex items-center gap-4">
                    <div class="flex flex-col items-end mr-4">
                        <span class="text-[9px] font-black text-slate-500 uppercase tracking-widest">Logged in as</span>
                        <span class="text-xs font-bold text-sky-400">{username}</span>
                    </div>
                    <a href="/docs" class="px-6 py-2.5 bg-slate-900 border border-slate-800 rounded-xl text-xs font-black uppercase tracking-widest hover:bg-slate-800 transition-all flex items-center gap-2">
                        <i data-lucide="file-code" class="w-4 h-4 text-sky-500"></i> API Docs
                    </a>
                </div>
            </div>

            <!-- Stats -->
            <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
                <div class="glass p-8 rounded-[2rem] glow">
                    <p class="text-[10px] font-black text-slate-500 uppercase tracking-[0.2em] mb-4">Total Processed</p>
                    <p class="text-4xl font-black">{stats["total_processed"]}</p>
                </div>
                <div class="glass p-8 rounded-[2rem]">
                    <p class="text-[10px] font-black text-slate-500 uppercase tracking-[0.2em] mb-4">Avg Processing Time</p>
                    <p class="text-4xl font-black text-sky-500">{avg_time:.2f}s</p>
                </div>
                <div class="glass p-8 rounded-[2rem]">
                    <p class="text-[10px] font-black text-slate-500 uppercase tracking-[0.2em] mb-4">Engine Success Rate</p>
                    <p class="text-4xl font-black text-green-500">{success_rate:.1f}%</p>
                </div>
                <div class="glass p-8 rounded-[2rem]">
                    <p class="text-[10px] font-black text-slate-500 uppercase tracking-[0.2em] mb-4">Service Uptime From</p>
                    <p class="text-xs font-black uppercase tracking-tight text-slate-400 mt-2">{stats["start_time"]}</p>
                </div>
            </div>

            <div class="grid grid-cols-1 lg:grid-cols-3 gap-10">
                <!-- Playground -->
                <div class="lg:col-span-1 space-y-6">
                    <div class="glass p-10 rounded-[2.5rem] space-y-8 sticky top-12">
                        <div class="flex items-center gap-3">
                            <i data-lucide="beaker" class="text-sky-500 w-5 h-5"></i>
                            <h2 class="font-black text-sm uppercase tracking-widest">Engine Playground</h2>
                        </div>
                        
                        <div id="dropzone" class="border-2 border-dashed border-slate-800 rounded-[2rem] p-10 text-center space-y-4 hover:border-sky-500/50 hover:bg-sky-500/5 transition-all cursor-pointer group">
                            <input type="file" id="fileInput" class="hidden" accept="image/*,.pdf">
                            <div class="w-16 h-16 bg-slate-900 rounded-2xl flex items-center justify-center mx-auto group-hover:scale-110 transition-transform">
                                <i data-lucide="upload-cloud" class="text-slate-500 group-hover:text-sky-500 w-8 h-8"></i>
                            </div>
                            <div>
                                <p class="text-xs font-black uppercase tracking-widest">Drop test document</p>
                                <p class="text-[10px] text-slate-500 font-bold mt-1">PDF or Images (Max 10MB)</p>
                            </div>
                        </div>

                        <div id="loading" class="hidden">
                            <div class="flex items-center gap-3 p-4 bg-sky-500/10 rounded-2xl border border-sky-500/20">
                                <i data-lucide="refresh-cw" class="w-4 h-4 text-sky-500 animate-spin"></i>
                                <p class="text-[10px] font-black text-sky-500 uppercase tracking-widest">Extracting intelligence...</p>
                            </div>
                        </div>

                        <div id="result" class="hidden space-y-4">
                            <div class="flex items-center justify-between">
                                <p class="text-[10px] font-black text-slate-500 uppercase">Extraction Result</p>
                                <button onclick="copyText()" class="text-sky-500 hover:text-sky-400"><i data-lucide="copy" class="w-4 h-4"></i></button>
                            </div>
                            <div class="bg-slate-950/50 rounded-2xl p-6 border border-slate-800">
                                <pre id="ocrText" class="text-xs text-slate-300 font-medium whitespace-pre-wrap leading-relaxed max-h-[300px] overflow-y-auto custom-scrollbar"></pre>
                            </div>
                        </div>
                    </div>
                </div>

                <!-- Recent Requests -->
                <div class="lg:col-span-2">
                    <div class="glass rounded-[2.5rem] overflow-hidden shadow-2xl h-full">
                        <div class="p-8 border-b border-slate-800/50 flex items-center justify-between">
                            <div class="flex items-center gap-3">
                                <i data-lucide="history" class="text-sky-500 w-5 h-5"></i>
                                <h2 class="font-black text-sm uppercase tracking-widest">Recent Processing History</h2>
                            </div>
                            <button onclick="window.location.reload()" class="p-2 bg-slate-900 rounded-lg text-slate-500 hover:text-white transition-colors">
                                <i data-lucide="refresh-cw" class="w-4 h-4"></i>
                            </button>
                        </div>
                        <div class="overflow-x-auto">
                            <table class="w-full text-left">
                                <thead class="text-[10px] font-black text-slate-500 uppercase tracking-[0.2em]">
                                    <tr class="bg-slate-900/50">
                                        <th class="py-5 px-6">Timestamp</th>
                                        <th class="py-5 px-6">File Name</th>
                                        <th class="py-5 px-6">Size</th>
                                        <th class="py-5 px-6">Duration</th>
                                        <th class="py-5 px-6">Status</th>
                                    </tr>
                                </thead>
                                <tbody class="text-sm">
                                    {rows}
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <script>
            lucide.createIcons();
            
            const dropzone = document.getElementById('dropzone');
            const fileInput = document.getElementById('fileInput');
            const loading = document.getElementById('loading');
            const result = document.getElementById('result');
            const ocrText = document.getElementById('ocrText');

            dropzone.onclick = () => fileInput.click();
            
            fileInput.onchange = (e) => {{
                if (e.target.files.length > 0) handleUpload(e.target.files[0]);
            }};

            async function handleUpload(file) {{
                const formData = new FormData();
                formData.append('file', file);

                loading.classList.remove('hidden');
                result.classList.add('hidden');
                dropzone.classList.add('opacity-50', 'pointer-events-none');

                try {{
                    const resp = await fetch('/ocr/process', {{
                        method: 'POST',
                        body: formData
                    }});
                    
                    if (!resp.ok) throw new Error('Failed to process OCR');
                    
                    const data = await resp.json();
                    ocrText.innerText = data.full_text || 'No text extracted.';
                    result.classList.remove('hidden');
                }} catch (err) {{
                    alert(err.message);
                }} finally {{
                    loading.classList.add('hidden');
                    dropzone.classList.remove('opacity-50', 'pointer-events-none');
                }}
            }}

            function copyText() {{
                navigator.clipboard.writeText(ocrText.innerText);
                alert('Text copied to clipboard!');
            }}
        </script>
    </body>
    </html>
    """

@app.get("/health")
def health_check():
    return {"status": "ok", "stats": stats}

@app.post("/ocr/process")
async def process_ocr(file: UploadFile = File(...), username: str = Depends(authenticate)):
    start_time = time.time()
    filename = file.filename
    content = await file.read()
    file_size = len(content)
    
    logger.info(f"Processing OCR request for file: {filename} (Size: {file_size} bytes) - Auth User: {username}")
    
    results = []
    full_text_parts = []

    try:
        if filename.lower().endswith('.pdf'):
            # Process PDF
            pdf = pdfium.PdfDocument(content)
            for page_index in range(len(pdf)):
                page = pdf[page_index]
                # Render page to image (scale=2 for better OCR)
                bitmap = page.render(scale=2)
                pil_image = bitmap.to_pil()
                
                # Convert PIL to numpy array for PaddleOCR
                img_array = np.array(pil_image)
                
                # Process page with EasyOCR
                page_result = reader.readtext(img_array)
                if page_result:
                    line_count = len(page_result)
                    logger.info(f"Page {page_index + 1}: Found {line_count} lines of text.")
                    extracted = format_ocr_result(page_result, page_index + 1)
                    results.extend(extracted["words"])
                    full_text_parts.append(extracted["full_text"])
                else:
                    logger.warning(f"Page {page_index + 1}: No text detected by engine.")
        else:
            # Process Image with EasyOCR
            img = Image.open(io.BytesIO(content))
            img_array = np.array(img.convert('RGB'))
            
            page_result = reader.readtext(img_array)
            if page_result:
                line_count = len(page_result)
                logger.info(f"Image processed: Found {line_count} lines of text.")
                extracted = format_ocr_result(page_result, 1)
                results.extend(extracted["words"])
                full_text_parts.append(extracted["full_text"])
            else:
                logger.warning("Image processed: No text detected by engine.")

        duration = time.time() - start_time
        logger.info(f"OCR completed for {filename} in {duration:.2f}s")
        
        # Record Success Stats
        stats["total_processed"] += 1
        stats["total_duration"] += duration
        request_history.append({
            "timestamp": datetime.now().strftime("%H:%M:%S"),
            "filename": filename,
            "size": format_size(file_size),
            "duration": duration,
            "status": "Success"
        })

    except Exception as e:
        logger.error(f"Error processing OCR for {filename}: {str(e)}")
        stats["total_processed"] += 1
        stats["failed_count"] += 1
        request_history.append({
            "timestamp": datetime.now().strftime("%H:%M:%S"),
            "filename": filename,
            "size": format_size(file_size),
            "duration": time.time() - start_time,
            "status": "Failed"
        })
        raise HTTPException(status_code=500, detail=f"OCR Processing Error: {str(e)}")

    return {
        "filename": filename,
        "full_text": "\n".join(full_text_parts),
        "words": results
    }

def format_ocr_result(result, page_num):
    words = []
    texts = []
    for line in result:
        box = line[0]  # [[x1,y1], [x2,y2], [x3,y3], [x4,y4]]
        text = line[1]
        confidence = line[2]
        
        texts.append(text)
        
        # Simplify box to x, y, w, h
        x = min(p[0] for p in box)
        y = min(p[1] for p in box)
        w = max(p[0] for p in box) - x
        h = max(p[1] for p in box) - y
        
        words.append({
            "text": text,
            "confidence": float(confidence),
            "page": page_num,
            "box": {"x": int(x), "y": int(y), "w": int(w), "h": int(h)}
        })
    
    return {
        "words": words,
        "full_text": " ".join(texts)
    }

def format_size(bytes):
    for unit in ['B', 'KB', 'MB', 'GB']:
        if bytes < 1024:
            return f"{bytes:.1f} {unit}"
        bytes /= 1024
    return f"{bytes:.1f} TB"

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
