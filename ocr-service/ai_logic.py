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

def get_ai_client(api_key: str = None):
    key = api_key or GEMINI_API_KEY
    if not key:
        return None
    try:
        return genai.Client(api_key=key)
    except Exception as e:
        logger.error(f"Failed to initialize Gemini AI: {str(e)}")
        return None

async def ocr_with_ai_vision(image_bytes: bytes, model: str = None, api_key: str = None, mime_type: str = "image/jpeg"):
    client = get_ai_client(api_key)
    if not client:
        return None
    
    model_name = model or GEMINI_MODEL
    logger.info(f"Using model: {model_name} for Vision OCR")
    
    try:
        response = await asyncio.to_thread(
            client.models.generate_content,
            model=model_name,
            contents=[
                genai.types.Part.from_bytes(data=image_bytes, mime_type=mime_type),
                "Extract all text from this document accurately. Maintain the structure and provide the raw text output."
            ]
        )
        return response.text
    except Exception as e:
        logger.error(f"AI Vision OCR error ({model_name}): {str(e)}")
        return None

async def analyze_with_ai(text: str, vision_text: str = None, model: str = None, api_key: str = None):
    client = get_ai_client(api_key)
    if not client or (not text.strip() and not vision_text):
        return None
    
    model_name = model or GEMINI_MODEL

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
            client.models.generate_content, 
            model=model_name, 
            contents=prompt
        )
        content = response.text
        if "```json" in content:
            content = content.split("```json")[1].split("```")[0]
        elif "```" in content:
            content = content.split("```")[1].split("```")[0]
        return json.loads(content.strip())
    except Exception as e:
        logger.error(f"AI Analysis error ({model_name}): {str(e)}")
        return None
