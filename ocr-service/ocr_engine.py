import easyocr
import logging
import os

logger = logging.getLogger("ocr-service.engine")

# Configuration
OCR_WORKERS = int(os.getenv("WORKERS", "0"))
BATCH_SIZE = int(os.getenv("BATCH_SIZE", "1"))
IMAGE_SCALE = float(os.getenv("IMAGE_SCALE", "2.0"))

# Global reader instance
_reader = None

def get_reader():
    global _reader
    if _reader is None:
        logger.info("Initializing Kreatif DMS OCR Engine (EasyOCR)...")
        _reader = easyocr.Reader(['id', 'en'], gpu=False, verbose=False)
        logger.info("Kreatif DMS OCR Engine Ready.")
    return _reader

def format_ocr_result(result, page_number):
    words = []
    full_text_parts = []
    for line in result:
        # line format: [[x,y],[x,y],[x,y],[x,y]], text, confidence
        coords = line[0]
        text = line[1]
        conf = line[2]
        
        # Calculate bounding box
        x_min = min(p[0] for p in coords)
        y_min = min(p[1] for p in coords)
        x_max = max(p[0] for p in coords)
        y_max = max(p[1] for p in coords)
        
        words.append({
            "text": text,
            "confidence": float(conf),
            "page": page_number,
            "box": {
                "x": int(x_min),
                "y": int(y_min),
                "w": int(x_max - x_min),
                "h": int(y_max - y_min)
            }
        })
        full_text_parts.append(text)
    
    return {
        "words": words,
        "full_text": " ".join(full_text_parts)
    }

def run_ocr(img_array):
    return get_reader().readtext(img_array, workers=OCR_WORKERS, batch_size=BATCH_SIZE)
