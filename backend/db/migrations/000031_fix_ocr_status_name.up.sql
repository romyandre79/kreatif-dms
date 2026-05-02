-- Migration: Fix status case for existing OCR jobs
UPDATE ocr_jobs SET status = 'Success' WHERE status = 'completed';
