from fastapi import FastAPI, UploadFile, File, HTTPException
from paddleocr import PaddleOCR
import uvicorn
import os
import io
import pypdfium2 as pdfium
from PIL import Image
import numpy as np
import logging
import time

# Configure logging
logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s] %(name)s: %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S"
)
logger = logging.getLogger("ocr-service")

app = FastAPI(title="Kreatif DMS OCR Service")

# Initialize PaddleOCR
# use_textline_orientation=True enables direction classification
# lang='id' for Indonesian support
ocr = PaddleOCR(use_textline_orientation=True, lang='id')

@app.get("/", tags=["Home"])
def read_root():
    from fastapi.responses import HTMLResponse
    return HTMLResponse(content="""
        <!DOCTYPE html>
        <html>
            <head>
                <title>Kreatif DMS - OCR Service</title>
                <style>
                    body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background: #0f172a; color: white; display: flex; justify-content: center; align-items: center; height: 100vh; margin: 0; }
                    .card { background: rgba(30, 41, 59, 0.7); padding: 3rem; border-radius: 1.5rem; text-align: center; backdrop-filter: blur(10px); border: 1px solid rgba(255,255,255,0.1); box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.3); }
                    h1 { color: #38bdf8; margin-bottom: 0.5rem; }
                    .status { color: #4ade80; font-weight: bold; margin-bottom: 2rem; }
                    a { color: #38bdf8; text-decoration: none; margin: 0 1rem; border: 1px solid #38bdf8; padding: 0.5rem 1rem; border-radius: 0.5rem; transition: all 0.3s; }
                    a:hover { background: #38bdf8; color: #0f172a; }
                </style>
            </head>
            <body>
                <div class="card">
                    <h1>OCR Service Engine</h1>
                    <div class="status">● PaddleOCR Engine Online</div>
                    <p>High-performance OCR service for Kreatif DMS.</p>
                    <div style="margin-top: 2rem;">
                        <a href="/docs">API Docs (Swagger)</a>
                        <a href="/health">Health Check</a>
                    </div>
                </div>
            </body>
        </html>
    """)

@app.get("/health")
def health_check():
    return {"status": "ok"}

@app.post("/ocr/process")
async def process_ocr(file: UploadFile = File(...)):
    start_time = time.time()
    filename = file.filename
    content = await file.read()
    
    logger.info(f"Processing OCR request for file: {filename} (Size: {len(content)} bytes)")
    
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
                
                # Process page
                page_result = ocr.ocr(img_array, cls=True)
                if page_result and page_result[0]:
                    extracted = format_ocr_result(page_result[0], page_index + 1)
                    results.extend(extracted["words"])
                    full_text_parts.append(extracted["full_text"])
        else:
            # Process Image
            img = Image.open(io.BytesIO(content))
            img_array = np.array(img.convert('RGB'))
            
            page_result = ocr.ocr(img_array, cls=True)
            if page_result and page_result[0]:
                extracted = format_ocr_result(page_result[0], 1)
                results.extend(extracted["words"])
                full_text_parts.append(extracted["full_text"])

        duration = time.time() - start_time
        logger.info(f"OCR completed for {filename} in {duration:.2f}s")

    except Exception as e:
        logger.error(f"Error processing OCR for {filename}: {str(e)}")
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
        text = line[1][0]
        confidence = line[1][1]
        
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

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
