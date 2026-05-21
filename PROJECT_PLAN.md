# Kreatif DMS - Project Implementation Plan

## Tech Stack
- **Backend Framework**: Go + Fiber v3
- **Database**: PostgreSQL (pgx v5, sqlc)
- **Caching & Rate Limiting**: Redis (go-redis v9)
- **File Storage**: MinIO
- **JWT Auth**: golang-jwt v5
- **AI Integration**: Gemini, OpenAI, Ollama
- **Frontend**: Nuxt 4 + TailwindCSS
- **Role System**: Admin, Manager, Doc Controller, User, Guest
- **Data Authorization**: User based on Department & Branch
- **Document Structure**: Company -> Kantor Cabang -> Departemen -> Rak -> Boks -> Ordner
- **Approval Rules**: 
  - Standard User: Upload to their department, approved by Dept Head.
  - Doc Controller: Can upload to ANY department, approved by Doc Controller Head.
- **Access Control & Masking**: 
  - Dokumen bisa dibatasi hanya untuk Role tertentu.
  - Jika User diluar Role tersebut mencoba melihat, informasi judul/metadata akan ditampilkan sebagai "**RAHASIA**".
- **Core Features**: 
  - Document Versioning (V1, V2, V3...)
  - Peminjaman & Approval Dokumen (Workflow)
  - RFID/NFC Scanning (Check-in & Check-out dokumen fisik)
  - Notifikasi Otomatis (Email & WhatsApp)
  - OCR (Lokal, PaddleOCR via Python Microservice) dengan UI Koreksi Teks (Per-Kata)
  - Full-Text Search (Isi Dokumen & Metadata) via Elasticsearch
  - Keamanan File: Enkripsi AES256 di MinIO
  - Preview PDF Aman: Watermark otomatis (posisi bisa diatur)

## Task Checklist

### Phase 1: Infrastructure & Setup
- [ ] Initialize Root Directory (`c:\lara\www\kreatif-dms`)
- [ ] Create `docker-compose.yml` (PostgreSQL, Redis, MinIO)
- [ ] Initialize Go Backend (`go mod init backend`)
- [ ] Setup Makefile for developer commands

### Phase 2: Backend Core Structure
- [ ] Setup Fiber v3 project structure (Clean Architecture)
- [ ] Configure Viper (load `.env`)
- [ ] Setup Database Connection Pool (`pgxpool`)
- [ ] Setup Redis Client
- [ ] Setup MinIO Client

### Phase 3: Database & Migration
- [ ] Write SQL schema migrations (Roles, Users, Folders, Documents, etc)
- [ ] Write `sqlc` queries
- [ ] Generate Go code from `sqlc`

### Phase 4: Backend Security & Auth
- [ ] Implement JWT token generation & parsing
- [ ] Implement Redis Token Blacklist (Logout)
- [ ] Implement Fiber Middlewares (CORS, Limiter via Redis, Helmet)
- [ ] Implement RBAC Middleware (`RequireRole`)

### Phase 5: Backend Features (API Handlers)
- [ ] Auth & User Handlers
- [ ] Document Handlers (Upload to MinIO, Save metadata to DB)
- [ ] Folder Management Handlers
- [ ] AI Integration Handlers (Gemini, OpenAI, Ollama routers)

### Phase 6: Frontend Nuxt 4 Setup
- [ ] Initialize Nuxt 4 (`npx create-nuxt@latest frontend`)
- [ ] Setup TailwindCSS & modern UI components
- [ ] Setup Pinia for state management
- [ ] Build Auth pages (Login, Register) & Auth middleware
- [ ] Build Dashboard & Document Explorer pages
- [ ] Build AI Chat Interface

### Phase 7: Final Testing & Polish
- [ ] E2E Testing for Upload & AI Chat
- [ ] UI/UX Polish (Glassmorphism, Animations)
- [ ] Security Headers & Permissions Validation

### Phase 8: Flutter Mobile App
- [x] Initialize Flutter Project (`flutter create mobile`)
- [ ] Setup State Management (Riverpod) & API Client (Dio)
- [ ] Design System (Colors, Typography, Themes)
- [ ] Implement Auth Flow (Login, Token Persistence)
- [ ] Implement Document Browser & Search
- [ ] Implement QR/Barcode Scanner for Inventory Audit
- [ ] Implement Document Detail & Preview

### Phase 9: Circulation & Settings (Ad-hoc Features)
- [x] Implement dynamic system-wide Penalty Policy setting & display
- [x] Refactor loan timeline nodes logic & hide reschedule options when pending
- [x] Disable borrowing action & show crossed-out cart icon for active loans
- [x] Activate, localize, and connect circulation pickup preparation page (with automatic redirect to checkout)
- [x] Add multi-file (attachments) support for document upload — migration 000080 (`document_files` table), new `document_files.go` repository, `ExtraFiles` param in `UploadDocument` service
- [x] Overhaul `circulation/checkout.vue` with full handover verification flow (identity check, per-document condition & note, signature area, completion gate)
- [x] Update loan status labels in backend: `l1_approved` → "Dalam Persiapan", `l2_approved` → "Siap Diambil", new `active` status → "Active"
- [x] Add `returnDate` field to all-loans list response
- [x] Add i18n circulation locales (`frontend/app/locales/en/circulation.json`, `id/circulation.json`) covering pickup, checkout, checkin, inbox, routing slip, labels, signature flows
- [x] Update `loans/index.vue` and `loans/my.vue` with accurate status badge display aligned to new status labels

