import os
import json
import asyncio
import logging
from google import genai

logger = logging.getLogger("ocr-service.ai")

# Configuration
AI_ENABLED = os.getenv("AI_ENABLED", "false").lower() == "true"
GEMINI_API_KEY = os.getenv("GEMINI_API_KEY", "")
GEMINI_MODEL = os.getenv("GEMINI_MODEL", "gemini-1.5-flash")
AI_VISION_ENABLED = os.getenv("AI_VISION_ENABLED", "false").lower() == "true"

ai_client = None
if AI_ENABLED and GEMINI_API_KEY:
    try:
        ai_client = genai.Client(api_key=GEMINI_API_KEY)
    except Exception as e:
        logger.error(f"Failed to initialize Gemini AI: {str(e)}")

async def ocr_with_ai_vision(image_bytes: bytes, mime_type: str = "image/jpeg"):
    if not ai_client:
        return None
    
    try:
        response = await asyncio.to_thread(
            ai_client.models.generate_content,
            model=GEMINI_MODEL,
            contents=[
                genai.types.Part.from_bytes(data=image_bytes, mime_type=mime_type),
                "Extract all text from this document accurately. Maintain the structure and provide the raw text output."
            ]
        )
        return response.text
    except Exception as e:
        logger.error(f"AI Vision OCR error: {str(e)}")
        return None

async def analyze_with_ai(text: str, vision_text: str = None):
    if not ai_client or (not text.strip() and not vision_text):
        return None
    
    prompt = f"""
    You are an expert document analyzer for Kreatif DMS. 
    Analyze the following document content and provide a structured JSON response.
    
    I will provide you with two inputs:
    1. RAW OCR TEXT: Teks dari engine OCR (dengan koordinat).
    2. VISION TEXT: Teks yang Anda lihat langsung dari gambar (Vision).
    
    Compare both and provide the most accurate extraction.
    
    TASKS:
    1. Identify document type (Invoice, Receipt, ID Card, Contract, etc.)
    2. Extract key entities (Document Number, Date, Total Amount, Vendor Name, etc.)
    3. Provide a very brief summary (1 sentence).
    4. Provide a 'cleaned_text' version that is human-readable.
    
    RAW OCR TEXT:
    {text}
    
    VISION TEXT:
    {vision_text if vision_text else 'Not available'}
    
    JSON FORMAT:
    {{
        "doc_type": "...",
        "summary": "...",
        "entities": {{ "key": "value", ... }},
        "cleaned_text": "...",
        "confidence_score": 0.0-1.0
    }}
    """
    try:
        response = await asyncio.to_thread(
            ai_client.models.generate_content, 
            model=GEMINI_MODEL, 
            contents=prompt
        )
        content = response.text
        if "```json" in content:
            content = content.split("```json")[1].split("```")[0]
        elif "```" in content:
            content = content.split("```")[1].split("```")[0]
        return json.loads(content.strip())
    except Exception as e:
        logger.error(f"AI Analysis error: {str(e)}")
        return None
