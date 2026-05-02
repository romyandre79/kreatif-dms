import sqlite3
import os
import time
import json
from contextlib import contextmanager

DB_PATH = os.getenv("DATABASE_URL", "ocr_service.db")

def init_db():
    with sqlite3.connect(DB_PATH) as conn:
        cursor = conn.cursor()
        # History table
        cursor.execute("""
            CREATE TABLE IF NOT EXISTS request_logs (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                start_time TEXT,
                end_time TEXT,
                filename TEXT,
                size TEXT,
                duration REAL,
                status TEXT,
                words_json TEXT,
                file_path TEXT,
                preview_path TEXT,
                accuracy REAL,
                ai_analysis TEXT,
                preview_paths TEXT
            )
        """)
        # Rate limit table
        cursor.execute("""
            CREATE TABLE IF NOT EXISTS rate_limits (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                client_host TEXT,
                timestamp REAL
            )
        """)
        conn.commit()

@contextmanager
def get_db():
    conn = sqlite3.connect(DB_PATH)
    try:
        yield conn
    finally:
        conn.close()

def save_request_log(data):
    with get_db() as conn:
        cursor = conn.cursor()
        cursor.execute("""
            INSERT INTO request_logs (
                start_time, end_time, filename, size, duration, status, 
                words_json, file_path, preview_path, accuracy, ai_analysis, preview_paths
            ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
        """, (
            data['start_time'], data['end_time'], data['filename'], data['size'], 
            data['duration'], data['status'], data['words_json'], 
            data['file_path'], data['preview_path'], data['accuracy'], 
            data['ai_analysis'], data['preview_paths']
        ))
        conn.commit()

def get_stats():
    with get_db() as conn:
        cursor = conn.cursor()
        cursor.execute("SELECT COUNT(*), SUM(duration), SUM(CASE WHEN status='Failed' THEN 1 ELSE 0 END) FROM request_logs")
        total, total_dur, failed = cursor.fetchone()
        return total or 0, total_dur or 0, failed or 0

def get_recent_history(limit=50):
    with get_db() as conn:
        cursor = conn.cursor()
        cursor.execute("SELECT id, start_time, end_time, filename, size, duration, status, accuracy, ai_analysis FROM request_logs ORDER BY id DESC LIMIT ?", (limit,))
        return cursor.fetchall()

def get_ocr_result_by_id(job_id):
    with get_db() as conn:
        cursor = conn.cursor()
        cursor.execute("SELECT id, start_time, filename, status, words_json, file_path, preview_path, ai_analysis, preview_paths FROM request_logs WHERE id = ?", (job_id,))
        return cursor.fetchone()
