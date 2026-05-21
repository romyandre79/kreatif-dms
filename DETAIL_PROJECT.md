File ini menjelaskan detail setiap fitur yang ada pada sistem document management 
saya, fitur tersebut ada di tabel dokumen, user, sirkulasi, pinjam dan 
documan detail.

### BP-01: User Access & Document Intake

### 1. Login Page
- **Status**: ✅ Frontend ada, ✅ Backend ada
- **Vue File**: `frontend/app/pages/login.vue`
- **Deskripsi**: Antarmuka login premium dengan fitur:
  - **Multi-Auth**: SSO (LDAP/AD) dan Akun Lokal.
  - **Security**: Password recovery, MFA ready, dan animasi glassmorphism.
- **Backend API**: ✅ `POST /api/v1/auth/login` (Status: ✅ Resolved LDAP Protocol Error & Role Mapping)

---

### 2. Multi-Role Dashboard System
- **Status**: ⚠️ Integrasi Sedang Berjalan (Frontend ✅, Backend 🚧)
- **Vue Files**: 
  - `frontend/app/pages/dashboard.vue` (Router)
  - `frontend/app/components/dashboard/*.vue`
- **Backend API**: 🚧 `GET /api/v1/dashboard/summary`
- **Logic**: Mengembalikan data ringkasan berbeda (Staff/Manager/Admin) berdasarkan JWT role.
- **Handler**: `dashboard_handler.go` (Sedang dibuat)
- **Service**: `dashboard_service.go` (Sedang dibuat)

---

### 3. Document Intake System (Manual & Bulk)
- **Status**: ⚠️ Integrasi Sedang Berjalan (Frontend ✅, Backend 🚧)
- **Vue File**: `frontend/app/pages/documents/upload.vue`
- **Backend API**:
  - ✅ `POST /api/v1/documents/` — Manual Entry (mendukung multi-file attachment via field `files[]`)
  - 🚧 `POST /api/v1/documents/bulk-import` — Excel Bulk Processing
  - ✅ `POST /api/v1/batches/` — Batch Control
- **Logic**: 
  - Manual: Menyimpan satu dokumen dengan metadata lengkap. File utama (index 0) disimpan di kolom dokumen utama; file tambahan (index 1+) disimpan di tabel `document_files` (migration 000080).
  - Bulk: Parsing file Excel, validasi baris, dan pembuatan dokumen massal dalam satu batch.
- **Handler**: `document_handler.go` & `batch_handler.go`
- **Service**: `document_service.go`
- **Database Tables**: `documents` ✅, `document_files` ✅ (migration 000080 — multi-attachment)






BP-02



BP-03



BP-04



BP-05



BP-06



BP-07



## Master Data

---

### 1. Pengaturan Single Sign On (LDAP/AD)
- **Status**: ✅ Selesai (Frontend ✅, Backend ✅)
- **Vue File**: `frontend/app/pages/settings/sso.vue`
- **Backend API**: ✅ `GET/POST /api/v1/master/settings/sso`
- **Dynamic Config**: ✅ Konfigurasi sekarang disimpan di tabel `integration_nodes` (ServiceType: LDAP).
- **Handler**: `master_handler.go`
- **Implementation**: Menggunakan `LDAPService` dinamis yang mendukung `LLDAP` dan `Active Directory`.

---

### 2. Pengaturan Role / Role Management
- **Status**: ✅ Selesai (Frontend ✅, Backend ✅)
- **Vue File**: `frontend/app/pages/admin/roles.vue`
- **Backend API**: ✅ CRUD `/api/v1/master/roles`
- **Handler**: `master_handler.go`
- **Role Baru**: `superadmin`, `manajer`, `admin doc controller`, `kepala doc controller`, `user`

---

### 3. Pendaftaran PIN Pengguna
- **Status**: ✅ Selesai (Frontend ✅, Backend ✅)
- **Vue File**: `frontend/app/pages/admin/security/pin.vue`
- **Backend API**: ✅ `POST /api/v1/auth/pin` (Set & Verify)
- **Handler**: `auth_handler.go`

---

### 4. Pengaturan Departemen / Departemen Management
- **Status**: ✅ Selesai (Frontend ✅, Backend ✅)
- **Vue File**: `frontend/app/pages/config/department.vue`
- **Backend API**: ✅ CRUD `/api/v1/master/departments`
- **Handler**: `master_handler.go`

---

### 5. Pengaturan Cabang / Branch Management
- **Status**: ✅ Selesai (Frontend ✅, Backend ✅)
- **Vue File**: `frontend/app/pages/config/branch.vue`
- **Backend API**: ✅ CRUD `/api/v1/master/branches`
- **Handler**: `master_handler.go`

---

### 6. Pengaturan Rak / Rack Management
- **Status**: ✅ Selesai (Frontend ✅, Backend ✅)
- **Vue File**: `frontend/app/pages/config/rack.vue`
- **Backend API**: ✅ CRUD `/api/v1/master/racks`
- **Handler**: `master_handler.go`

---

### 7. Pengaturan Box / Box Management
- **Status**: ✅ Selesai (Frontend ✅, Backend ✅)
- **Vue File**: `frontend/app/pages/config/box.vue`
- **Backend API**: ✅ CRUD `/api/v1/master/boxes`
- **Handler**: `master_handler.go`

---

### 8. Pengaturan Ordner / Ordner Management
- **Status**: ✅ Selesai (Frontend ✅, Backend ✅)
- **Vue File**: `frontend/app/pages/config/ordner.vue`
- **Backend API**: ✅ CRUD `/api/v1/master/ordners`
- **Handler**: `master_handler.go`

---

### 9. Kebijakan Retensi / Retention Policies
- **Status**: ✅ Selesai (Frontend ✅, Backend ✅)
- **Backend API**: ✅ `GET /api/v1/master/retention`
- **Handler**: `master_handler.go`

---

### 10. Manajemen Hardware (RFID & Labeling)
- **Status**: ✅ Selesai (Frontend ✅, Backend ✅)
- **Backend API**: ✅ `POST /api/v1/hardware/rfid/assign`, `GET /api/v1/hardware/labels/generate`
- **Handler**: `hardware_handler.go`

---

### 11. Integration Status Monitor
- **Status**: ✅ Selesai (Frontend ✅, Backend ✅)
- **Vue File**: `frontend/app/pages/config/integration.vue`
- **Backend API**: ✅ `GET /api/v1/master/integration/status` — Live Telemetry
- **Deskripsi**: Dashboard pemantauan kesehatan layanan pihak ketiga dan hardware secara real-time.
- **Backend Worker**: `IntegrationMonitorService` melakukan health-check setiap 30 detik.
- **Database Driven**: Semua konfigurasi dipindah dari `.env` ke tabel `integration_nodes`.
- **Integrasi API Eksternal**:
  - **AD/LDAP**: Integrasi LLDAP/Active Directory untuk autentikasi SSO (Status: ✅ Stable with StartTLS & Fresh Connection Support).
  - **S3 Storage (MinIO)**: Konfigurasi dinamis untuk penyimpanan dokumen terenkripsi.
  - **SMTP Relay**: Menggunakan Mailpit (dev) atau SMTP relay lainnya via database.
  - **Hardware Gateway**: Monitoring status koneksi Network Scanner, Printer ZPL, dan RFID Encoder.
- **Features**: Live Latency Trend (Database History), Alert Feed (Dynamic Status), & Health Score Telemetry.

---

### 12. User / Document Detail (Legacy Bridge)
- **Status**: ✅ Frontend ada, ⚠️ Backend parsial
- **Vue Files**:
  - `frontend/app/pages/admin/users/index.vue` — user management
  - `frontend/app/pages/admin/users/hierarchy.vue` — user hierarchy org
  - `frontend/app/pages/documents/[id].vue` — document detail view
  - `frontend/app/pages/documents/index.vue` — document list
- **Backend API**: ⚠️ Parsial — yang sudah ada:
  - `GET /api/v1/users/` ✅ (list users)
  - `GET /api/v1/users/pending` ✅ (pending approval)
  - `POST /api/v1/users/:id/approve` ✅ (approve user)
  - Perlu ditambah:
    - `GET /api/v1/users/:id` — detail user
    - `PUT /api/v1/users/:id` — update user
    - `GET /api/v1/documents/` — list documents (dengan filter)
    - `GET /api/v1/documents/:id` — detail dokumen lengkap
- **Handler yang ada**: `user_handler.go` ✅, `document_handler.go` ✅
- **Database Tables**:
  - `users` ✅, `roles` ✅, `documents` ✅, `document_versions` ✅
  - `document_role_access` ✅, `document_tracking_events` (baru)
- **Relasi Service**: PostgreSQL, MinIO (file preview), Elasticsearch (search)
- **Catatan**: "Legacy Bridge" menunjukkan ini adalah jembatan untuk memigrasikan data dari sistem lama. View `documents/[id].vue` sudah sangat detail (28KB) dengan tab versioning, metadata, ACL.

---

### 9. Split Screen Input Workspace
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/registration/migration.vue` (37KB — halaman terbesar)
- **Backend API**: ❌ Belum ada endpoint — perlu:
  - `POST /api/v1/registrations` — buat registrasi dokumen baru
  - `PUT /api/v1/registrations/:id` — update data registrasi
  - `GET /api/v1/registrations/:id/suggestions` — rekomendasi lokasi penyimpanan
  - `POST /api/v1/registrations/:id/validate` — validasi data & generate filename
  - `POST /api/v1/registrations/:id/complete` — finalisasi registrasi
- **Handler yang dibutuhkan**: `registration_handler.go`
- **Service yang dibutuhkan**: `registration_service.go`
- **Database Tables**:
  - `document_registrations` (baru dari migration 000010)
  - `staging_documents` (baru)
  - `documents` ✅, `racks` ✅, `boxes` ✅, `ordners` ✅
  - `archival_number_sequences` (baru — auto-numbering)
  - `rack_capacities` (baru — rekomendasi lokasi)
- **Relasi Service**: PostgreSQL, MinIO (upload file), OCR Service (extract text), Gemini AI (metadata)
- **Catatan**: Ini halaman paling kompleks (37KB Vue). Split screen kiri=dokumen preview, kanan=form input. Butuh integrasi OCR pipeline lengkap.

---

### 10. OCR & AI Review
- **Status**: ✅ Frontend ada, ⚠️ Backend parsial (OCR worker ada)
- **Vue Files**:
  - `frontend/app/pages/intake/ocr-processing.vue` — proses OCR sedang berjalan
  - `frontend/app/pages/intake/ocr-review.vue` — review & koreksi hasil OCR
- **Backend API**: ⚠️ Parsial — yang sudah ada:
  - Worker: `ProcessDocumentOCR()` ✅ (di `worker/processor.go`)
  - OCR Service: `POST /ocr/process` ✅ (di `ocr-service/main.py`)
  - AI Refinement: `RefineOCRText()` ✅ (di `infra/ai_service.go`)
  - Perlu ditambah:
    - `GET /api/v1/ocr/jobs` — list OCR jobs
    - `GET /api/v1/ocr/jobs/:id` — detail hasil OCR
    - `PUT /api/v1/ocr/jobs/:id/review` — submit koreksi teks
    - `POST /api/v1/ocr/jobs/:id/reprocess` — jalankan ulang OCR
- **Handler yang dibutuhkan**: `ocr_handler.go`
- **Service yang dibutuhkan**: Extend `document_service.go`
- **Database Tables**:
  - `ocr_jobs` (baru dari migration 000010) — tracking OCR invocations
  - `intake_sessions` (baru) — pipeline state
  - `async_task_queue` (baru) — link ke Asynq worker
  - `documents.extracted_text` ✅, `documents.is_ocr_processed` ✅
- **Relasi Service**: OCR Service (PaddleOCR), Gemini AI (refinement), Asynq/Redis (worker queue), MinIO (source file), Elasticsearch (re-index after correction)
- **Catatan**: Worker OCR sudah berjalan end-to-end. Yang belum ada adalah API untuk review/koreksi manual hasil OCR dari frontend.

---

### 11. Duplicate Detection Resolution
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/intake/duplicate-check.vue`
- **Backend API**: ❌ Belum ada endpoint — perlu:
  - `POST /api/v1/intake/:id/duplicate-check` — jalankan pengecekan duplikasi
  - `GET /api/v1/intake/:id/duplicates` — list potensi duplikat
  - `POST /api/v1/intake/:id/duplicate-resolve` — resolusi (keep/merge/reject)
- **Handler yang dibutuhkan**: Tambah di `intake_handler.go`
- **Service yang dibutuhkan**: `duplicate_detection_service.go`
- **Database Tables**:
  - `intake_sessions.duplicate_status` (baru)
  - `intake_sessions.duplicate_document_id` (baru)
  - `document_registrations.is_duplicate_checked` (baru)
  - `search_index_status` (baru — digunakan utk query ES similarity)
- **Relasi Service**: Elasticsearch (similarity search), PostgreSQL (metadata match)
- **Catatan**: Duplikasi bisa dicek via: (1) checksum SHA256 (exact match), (2) judul/metadata similarity via ES, (3) OCR text similarity. Butuh implementasi multi-strategy.

---

### 12. Identity & Foldering Finalization
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue Files**:
  - `frontend/app/pages/intake/path-review.vue` — review path/lokasi penyimpanan
  - `frontend/app/pages/intake/final-review.vue` — final review sebelum arsip
- **Backend API**: ❌ Belum ada endpoint — perlu:
  - `GET /api/v1/intake/:id/path-suggestion` — rekomendasi path berdasarkan departemen & kapasitas
  - `PUT /api/v1/intake/:id/path` — override/confirm path
  - `POST /api/v1/intake/:id/finalize` — finalisasi dan generate kode arsip
  - `GET /api/v1/racks/available` — rak dengan kapasitas tersedia
- **Handler yang dibutuhkan**: Tambah di `intake_handler.go`
- **Service yang dibutuhkan**: `archival_service.go`
- **Database Tables**:
  - `intake_sessions` (baru) — suggested_department_id, suggested_rack_id
  - `archival_number_sequences` (baru) — auto-generate kode
  - `rack_capacities` (baru) — cek kapasitas
  - `racks` ✅, `boxes` ✅, `ordners` ✅
  - `file_storage_objects` (baru) — finalisasi file ke MinIO
- **Relasi Service**: PostgreSQL, MinIO (final file placement)
- **Catatan**: Proses ini menghasilkan kode arsip final (contoh: `DMS-INV-2024-00892-CORP.pdf`) dan menempatkan dokumen di lokasi fisik (Rak → Box → Ordner).

---

### 13. Bulk Import Workbench
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/import/excel.vue`
- **Backend API**: ❌ Belum ada endpoint — perlu:
  - `POST /api/v1/import/upload` — upload file Excel/CSV
  - `GET /api/v1/import/jobs` — list import jobs
  - `GET /api/v1/import/jobs/:id` — detail & progress
  - `POST /api/v1/import/jobs/:id/start` — mulai proses import
  - `GET /api/v1/import/jobs/:id/errors` — download error log
  - `GET /api/v1/import/template` — download template Excel
- **Handler yang dibutuhkan**: `import_handler.go`
- **Service yang dibutuhkan**: `import_service.go`
- **Database Tables**:
  - `import_jobs` (baru dari migration 000010)
  - `staging_documents` (baru) — staging area sebelum registrasi
  - `document_registrations` (baru) — auto-create dari Excel rows
  - `async_task_queue` (baru) — background processing
- **Relasi Service**: Asynq/Redis (background job), MinIO (upload Excel), PostgreSQL
- **Catatan**: Import Excel perlu library Go seperti `excelize`. Setiap row menjadi `document_registration` yang masuk ke pipeline intake.

---

### 14. Digital Mailroom Registration
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue Files**:
  - `frontend/app/pages/intake/external-registration.vue` — registrasi surat masuk
  - `frontend/app/pages/intake/scan-web.vue` — scan via web camera/scanner
  - `frontend/app/pages/intake/management.vue` — manajemen antrian intake
- **Backend API**: ❌ Belum ada endpoint — perlu:
  - `POST /api/v1/intake/sessions` — buat session intake baru
  - `GET /api/v1/intake/sessions` — list semua session (dengan filter status)
  - `GET /api/v1/intake/sessions/:id` — detail session
  - `PUT /api/v1/intake/sessions/:id` — update data session
  - `POST /api/v1/intake/sessions/:id/upload` — upload file scan/foto
  - `POST /api/v1/intake/sessions/:id/ocr` — trigger OCR processing
- **Handler yang dibutuhkan**: `intake_handler.go`
- **Service yang dibutuhkan**: `intake_service.go`
- **Database Tables**:
  - `intake_sessions` (baru dari migration 000010) — full pipeline
  - `ocr_jobs` (baru) — OCR processing tracking
  - `async_task_queue` (baru) — background jobs
  - `file_storage_objects` (baru) — uploaded scans
- **Relasi Service**: OCR Service, Gemini AI, Asynq/Redis, MinIO, Elasticsearch
- **Catatan**: Digital Mailroom adalah entry point utama dokumen masuk. Alur: Scan/Upload → OCR → Review → Duplicate Check → Path Review → Final Review → Archived.

---

### 15. Manajemen Override Status Rak
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/warehouse/rack-override.vue`
- **Backend API**: ❌ Belum ada endpoint — perlu:
  - `GET /api/v1/warehouse/racks` — list semua rak + status kapasitas
  - `PUT /api/v1/warehouse/racks/:id/capacity` — override kapasitas
  - `PUT /api/v1/warehouse/racks/:id/status` — ubah status rak (active/maintenance/full)
  - `GET /api/v1/warehouse/racks/:id/history` — riwayat perubahan
- **Handler yang dibutuhkan**: `warehouse_handler.go`
- **Service yang dibutuhkan**: `warehouse_service.go`
- **Database Tables**:
  - `racks` ✅ (sudah ada)
  - `rack_capacities` (baru dari migration 000010) — max_capacity, current_usage, usage_pct
- **Relasi Service**: PostgreSQL
- **Catatan**: Override diperlukan ketika kondisi fisik rak berubah (rusak, maintenance). `rack_capacities` punya generated column `usage_pct` otomatis.

---

### 16. Departemen Zoning Control Center
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/warehouse/departments.vue`
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/warehouse/zones` — mapping departemen ke area gudang
  - `PUT /api/v1/warehouse/zones/:dept_id` — assign zona gudang ke departemen
  - `GET /api/v1/warehouse/zones/map` — visual map data
- **Handler yang dibutuhkan**: Tambah di `warehouse_handler.go`
- **Service yang dibutuhkan**: Tambah di `warehouse_service.go`
- **Database Tables**:
  - `departments` ✅ (sudah ada)
  - `racks` ✅ (sudah ada, FK department_id)
  - `rack_capacities` (baru) — monitoring per zona
- **Relasi Service**: PostgreSQL
- **Catatan**: Zoning mengatur area gudang mana yang dialokasikan untuk departemen tertentu. Data sudah tersedia via relasi `departments → racks`.

---

### 17. RFID / QR Label Generation
- **Status**: ✅ Selesai (Frontend ✅, Backend ✅)
- **Vue Files**:
  - `frontend/app/pages/warehouse/labels.vue` — antrian cetak label
  - `frontend/app/pages/warehouse/labels-print-assignment.vue` — penugasan boks & preview
- **Backend API**: ✅ `GET /api/v1/intake/labels/waiting`, `GET /api/v1/intake/stats`
- **Handler**: `intake_handler.go`
- **Service**: `intake_service.go`
- **Logic**: Mengelola antrean dokumen yang menunggu pelabelan fisik dan memberikan statistik dashboard.

---

### 18. Label Printing Console
- **Status**: ✅ Selesai (Frontend ✅, Backend ✅)
- **Vue File**: `frontend/app/pages/warehouse/labels-print-assignment.vue`
- **Backend API**: ✅ `POST /api/v1/intake/boxes/assign`, `GET /api/v1/intake/boxes/search`
- **Handler**: `intake_handler.go`
- **Service**: `intake_service.go`
- **Logic**: Alur penugasan dokumen ke boks fisik (archiving) disertai simulasi pencetakan label thermal ZPL/QR.
- **Database Tables**:
  - `boxes` ✅ (assigned_docs_count updated)
  - `documents` ✅ (status changed to 'archived')

---

### 19. Capacity Monitoring & Recommendation / Monitoring Kapasitas & Rekomendasi
- **Status**: ✅ Selesai (Frontend ✅, Backend ✅)
- **Vue File**: `frontend/app/pages/warehouse/labels-print-assignment.vue`
- **Backend API**: ✅ `GET /api/v1/intake/boxes/recommend`
- **Handler**: `intake_handler.go`
- **Service**: `intake_service.go`
- **Logic**: Algoritma Smart Recommendation yang mencarikan boks kosong berdasarkan Departemen, Kapasitas Rak (`max_docs_capacity`), dan Zonasi Kategori Dokumen (`allowed_category_ids`).
- **Database Tables**:
  - `racks`, `boxes`, `ordners` ✅ (dengan kolom kapasitas baru)
  - `rack_capacities` ✅ (monitoring utilitas)

---

### 20. Warehouse Topology Explorer
- **Status**: ✅ Selesai (Frontend ✅, Backend ✅)
- **Vue Files**:
  - `frontend/app/pages/admin/warehouse/topology.vue`
  - `frontend/app/pages/warehouse/structure.vue`
- **Backend API**: ✅ `GET /api/v1/master/topology` (Recursive hierarchy)
- **Handler**: `master_handler.go`

---

### 21. Parent-Child RFID Label Model (F-16)
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/rfid/model-f16.vue` (197 baris)
- **Deskripsi UI**: Layout 3 kolom — kiri: Container Preview (parent box + QR/RFID status), tengah: Child Items (drag-sortable list dokumen dalam box), kanan: Summary (capacity, hierarchy, LDAP network status).
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/rfid/containers/:id` — detail container + child items
  - `PUT /api/v1/rfid/containers/:id/items` — update urutan/isi container
  - `POST /api/v1/rfid/containers/:id/link` — link QR + RFID tag ke container
  - `GET /api/v1/rfid/containers/:id/children` — list child documents
  - `POST /api/v1/rfid/containers/:id/scan` — proses scan RFID batch
- **Handler yang dibutuhkan**: `rfid_handler.go`
- **Service yang dibutuhkan**: `rfid_service.go`
- **Database Tables**:
  - `rfid_tags` ✅ (migration 000001) — tag_id, document_id
  - `generated_labels` (baru, migration 000010) — QR label registry
  - `boxes` ✅ (container = box)
  - `documents` ✅ (child items)
- **Relasi Service**: PostgreSQL, RFID Reader (external hardware)
- **Catatan**: Model F-16 menunjukkan relasi parent-child dimana 1 box (parent) berisi banyak dokumen (child). Setiap child bisa memiliki RFID tag sendiri. UI sudah mendukung drag-and-drop reordering.

---

### 22. Location Detail & Slot Config
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/warehouse/location-detail.vue` (202 baris)
- **Deskripsi UI**: Layout 3 kolom — kiri: Location Info (ID, zone, media type, capacity limits soft/hard, utilization bar), tengah: Slot Table (kode slot, box assignment, status: filled/available/overflow/locked), kanan: AI Recommendations (relocation suggestion, retention alerts) + Incoming Items queue.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/warehouse/locations/:id` — detail lokasi + slots + kapasitas
  - `PUT /api/v1/warehouse/locations/:id/limits` — update max/soft/hard capacity limits
  - `GET /api/v1/warehouse/locations/:id/slots` — list slot + occupancy
  - `PUT /api/v1/warehouse/locations/:id/slots/:slot` — assign/release box ke slot
  - `GET /api/v1/warehouse/locations/:id/incoming` — antrian box yang akan masuk
  - `GET /api/v1/warehouse/locations/:id/recommendations` — AI rekomendasi relokasi
- **Handler yang dibutuhkan**: Tambah di `warehouse_handler.go`
- **Service yang dibutuhkan**: Tambah di `warehouse_service.go`
- **Database Tables**:
  - `racks` ✅ (lokasi fisik)
  - `rack_capacities` (baru, migration 000010) — max_capacity, current_usage, soft_limit, hard_limit
  - `boxes` ✅ (slot occupancy)
  - `documents` ✅ (isi per box)
  - `retention_policies` (baru) — untuk retention alerts
- **Relasi Service**: PostgreSQL, Gemini AI (rekomendasi relokasi)
- **Catatan**: UI menampilkan status "AT CAPACITY" dengan animasi pulse saat utilization > 90%. AI recommendation cards (relocation & retention alerts) membutuhkan engine analisis data.

---

### 23. Input Entry Hub & Mode Selection / Manajemen Input Dokumen
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/intake/management.vue` (173 baris)
- **Deskripsi UI**: 3 mode cards (Single Input, Bulk Import, Digital Mailroom) + Recent Jobs table (ID, type, time, status: running/success/draft/failed) + Stats bar (processed today, OCR capacity utilization) + Quick Start guide.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/intake/dashboard` — statistik intake harian (total, kapasitas OCR)
  - `GET /api/v1/intake/jobs` — list recent jobs (semua tipe)
  - `GET /api/v1/intake/jobs/:id` — detail job
  - `POST /api/v1/intake/jobs/:id/retry` — retry failed job
  - `DELETE /api/v1/intake/jobs/:id` — batalkan draft job
- **Handler yang dibutuhkan**: `intake_handler.go`
- **Service yang dibutuhkan**: `intake_service.go`
- **Database Tables**:
  - `intake_sessions` (baru, migration 000010) — pipeline tracking
  - `import_jobs` (baru) — bulk import tracking
  - `async_task_queue` (baru) — background job status
  - `dashboard_stats_cache` (baru) — cached daily stats
- **Relasi Service**: PostgreSQL, Asynq/Redis (job queue), OCR Service (capacity check)
- **Catatan**: Ini adalah hub utama yang mengarahkan ke 3 alur berbeda (single → scan-web, bulk → excel import, mailroom → external-registration). Recent Jobs table menampilkan status real-time dari semua tipe intake.

---

### 21. OCR Extraction Process Monitor / Proses Ekstraksi Dokumen
- **Status**: ✅ Frontend ada, ⚠️ Backend parsial (worker OCR ada)
- **Vue File**: `frontend/app/pages/intake/ocr-processing.vue` (191 baris)
- **Deskripsi UI**: 4-step pipeline stepper (Preprocessing → OCR Extraction → AI Enrichment → Quality Assurance) dengan progress bar aktif + Quality Panel (avg confidence 92.4%, pages processed, character count) + Warnings Table (page, field, confidence %, action: Inspect) + Footer actions (Rerun OCR, Download, Continue).
- **Backend API**: ⚠️ Parsial — worker ada, tapi perlu API baru:
  - `GET /api/v1/ocr/jobs/:id/progress` — real-time progress (via SSE/WebSocket)
  - `GET /api/v1/ocr/jobs/:id/quality` — quality metrics (confidence, page count, char count)
  - `GET /api/v1/ocr/jobs/:id/warnings` — list fields with low confidence
  - `POST /api/v1/ocr/jobs/:id/rerun` — jalankan ulang OCR
  - `GET /api/v1/ocr/jobs/:id/download` — download raw/refined text
- **Handler yang dibutuhkan**: `ocr_handler.go`
- **Service yang dibutuhkan**: Extend `document_service.go`
- **Database Tables**:
  - `ocr_jobs` (baru, migration 000010) — status, confidence_avg, word_count, words_json
  - `async_task_queue` (baru) — link ke Asynq task
  - `intake_sessions` (baru) — pipeline state
- **Relasi Service**: OCR Service (PaddleOCR), Gemini AI (refinement), Asynq/Redis (progress tracking)
- **Catatan**: Pipeline 4 langkah sudah divisualisasikan dengan stepper. `ocr_jobs.words_json` berisi detail per-word: `[{text, confidence, page, box}]` yang digunakan untuk populate warnings table. Real-time progress bisa via Server-Sent Events.

---

### 22. Folder Path Review
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/intake/path-review.vue` (214 baris)
- **Deskripsi UI**: Layout 2 kolom — kiri: Interactive folder tree (ARSIP → PT → Dept → Year → Type → Document) dengan preview file placement, kanan: Path Rules (department policy, formula preview: `/ARSIP/{CORP}/{DEPT}/{YEAR}/{DOC_TYPE}`), toggle options (auto-create folder, apply retention).
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/intake/:id/path` — suggested path berdasarkan metadata
  - `PUT /api/v1/intake/:id/path` — override path manual
  - `POST /api/v1/intake/:id/path/validate` — validasi path (collision check)
  - `GET /api/v1/departments/:id/folder-policy` — ambil naming policy departemen
  - `GET /api/v1/warehouse/tree` — folder tree structure untuk preview
- **Handler yang dibutuhkan**: Tambah di `intake_handler.go`
- **Service yang dibutuhkan**: `archival_service.go`
- **Database Tables**:
  - `intake_sessions` (baru) — suggested path fields
  - `archival_number_sequences` (baru) — naming policy per dept/year
  - `departments` ✅, `racks` ✅, `boxes` ✅, `ordners` ✅
  - `file_storage_objects` (baru) — MinIO path mapping
- **Relasi Service**: PostgreSQL, MinIO (path mapping)
- **Catatan**: Formula `{CORP}/{DEPT}/{YEAR}/{DOC_TYPE}` adalah configurable per departemen. UI sudah menampilkan toggle auto-create folder dan apply retention. Health score (84.2%) dihitung dari completeness metadata.

---

### 23. Auto Numbering Studio
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/archival/numbering.vue` (187 baris)
- **Deskripsi UI**: Layout 2 kolom — kiri: Numbering Preview besar (`#AD-2024-08912`), Format preview (`[IT].[AKD].[INV].[HARDWARE].[2024].[00452]`), detail (schema, increment, type), Collision Checker (primary & secondary check) — kanan: Folder Path Tree (visual tree ARSIP → AKIRADATA → IT → 2024 → INVOICE → TAG) + Path Rules.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/archival/numbering/preview` — generate preview nomor berikutnya
  - `POST /api/v1/archival/numbering/check-collision` — cek duplikasi nomor arsip
  - `GET /api/v1/archival/numbering/schemas` — list skema penomoran
  - `POST /api/v1/archival/numbering/schemas` — buat skema baru
  - `PUT /api/v1/archival/numbering/schemas/:id` — update skema
  - `POST /api/v1/archival/numbering/commit` — commit & finalisasi kode arsip
- **Handler yang dibutuhkan**: `archival_handler.go`
- **Service yang dibutuhkan**: `archival_service.go`
- **Database Tables**:
  - `archival_number_sequences` (baru, migration 000010) — department_id, prefix, year, last_number
  - `documents` ✅ — collision check target
  - `departments` ✅ — dept-specific numbering
- **Relasi Service**: PostgreSQL
- **Catatan**: Auto-numbering menggunakan format `{DEPT}.{COMPANY}.{DOC_TYPE}.{CATEGORY}.{YEAR}.{SEQUENCE}`. Collision checker melakukan pengecekan primary (database) dan secondary (Elasticsearch) untuk memastikan tidak ada duplikat kode arsip.

---

### 24. Final Review & Commit
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/intake/final-review.vue` (152 baris)
- **Deskripsi UI**: Layout 2 kolom — kiri: Stats cards (Completeness 95%, OCR Confidence 92%), Validation Table (duplicate check: passed, archival ID, filename, folder path, intake channel), Digital Twin verification preview — kanan: Risk Checklist (PII scan, LDAP sync, hash verification, retention policy) semua ✅, Publish Controls (Commit & Publish, Send to Supervisor, Save as Draft, Purge), Auth timestamp.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/intake/:id/final-review` — ambil semua validasi result
  - `POST /api/v1/intake/:id/commit` — finalisasi & publish ke arsip
  - `POST /api/v1/intake/:id/send-supervisor` — kirim ke supervisor untuk approval
  - `POST /api/v1/intake/:id/save-draft` — simpan sebagai draft
  - `DELETE /api/v1/intake/:id/purge` — hapus permanen
  - `GET /api/v1/intake/:id/risk-check` — jalankan risk checklist (PII, hash, retention)
- **Handler yang dibutuhkan**: Tambah di `intake_handler.go`
- **Service yang dibutuhkan**: `intake_service.go`, `compliance_service.go`
- **Database Tables**:
  - `intake_sessions` (baru) — status, final validation results
  - `documents` ✅ — target insert setelah commit
  - `file_storage_objects` (baru) — finalize MinIO placement
  - `search_index_status` (baru) — trigger ES indexing
  - `async_task_queue` (baru) — background commit job
  - `approval_workflows` (baru) — supervisor approval flow
- **Relasi Service**: PostgreSQL, MinIO (commit file), Elasticsearch (index), Asynq/Redis (background commit)
- **Catatan**: Ini adalah langkah terakhir intake pipeline. Commit menghasilkan: (1) insert ke `documents`, (2) upload final ke MinIO, (3) index ke Elasticsearch, (4) generate kode arsip via `archival_number_sequences`, (5) update `rack_capacities`.

---

### 25. Dynamic Filter Sidebar
- **Status**: ✅ Frontend ada (built-in), ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/documents/index.vue` — bagian filter bar (baris 103-129)
- **Deskripsi UI**: Horizontal filter bar dengan: search within folder, dropdowns (Document Type, Year, Status: Available/On Loan), Sort button. Filters berubah kontekstual berdasarkan folder yang dipilih di sidebar kiri.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/documents?type=&year=&status=&dept=&sort=&page=&limit=` — list dengan multi-filter
  - `GET /api/v1/documents/filters` — available filter options (dynamic per folder)
  - `GET /api/v1/documents/facets?folder_id=` — faceted counts per filter value
- **Handler yang dibutuhkan**: Extend `document_handler.go`
- **Service yang dibutuhkan**: Extend `document_service.go`
- **Database Tables**:
  - `documents` ✅ — filtered query
  - `departments` ✅ — filter by dept
  - `document_role_access` ✅ — ACL filtering
- **Relasi Service**: PostgreSQL, Elasticsearch (faceted search)
- **Catatan**: Filters bersifat dinamis — opsi yang tersedia berubah berdasarkan folder yang aktif. Faceted counts (jumlah dokumen per filter value) idealnya di-query via Elasticsearch aggregations.

---

### 26. Search Results & Snippets
- **Status**: ✅ Frontend ada (built-in), ⚠️ Backend parsial (search endpoint ada)
- **Vue File**: `frontend/app/pages/documents/index.vue` — bagian search results (baris 134-180)
- **Deskripsi UI**: Saat query aktif (`?q=...`), tampilan berubah ke search mode: Header menampilkan "Showing X results for 'query'", card per result menampilkan: title, folder path breadcrumb, highlighted snippet, tags, status badge. Empty state menampilkan "did you mean" suggestions.
- **Backend API**: ⚠️ Parsial — yang ada:
  - `GET /api/v1/documents/search` ✅ (endpoint ada tapi sederhana)
  - Perlu diperkaya:
    - `GET /api/v1/documents/search?q=&highlight=true&snippet_size=200` — full-text search + highlight
    - `GET /api/v1/documents/search/suggest?q=` — autocomplete/did-you-mean
- **Handler yang ada**: `document_handler.go` ✅ (method `Search` sudah ada)
- **Service yang dibutuhkan**: Extend `document_service.go`
- **Database Tables**:
  - `documents` ✅ — search source
  - `search_index_status` (baru) — track indexed docs
- **Relasi Service**: Elasticsearch (full-text search, highlighting, suggestions), PostgreSQL (fallback)
- **Catatan**: Search results menampilkan HTML-highlighted snippets (`<mark>` tags). UI sudah handle empty state dengan "did you mean" suggestions — ini memerlukan Elasticsearch suggest API atau fuzzy matching.

---

### 27. Universal Search Home
- **Status**: ✅ Frontend ada (built-in), ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/documents/index.vue` — bagian sidebar search (baris 8-15) + header search bar
- **Deskripsi UI**: Search box di sidebar kiri (Explorer search) + search bar di header (global search). Sidebar search filters folder tree, header search triggers full-text search dan mengubah view ke search results mode.
- **Backend API**: ❌ Belum ada endpoint universal search — perlu:
  - `GET /api/v1/search?q=&scope=all` — universal search across documents, folders, users
  - `GET /api/v1/search/recent` — recent searches per user
  - `GET /api/v1/search/popular` — trending/popular searches
- **Handler yang dibutuhkan**: `search_handler.go`
- **Service yang dibutuhkan**: `search_service.go`
- **Database Tables**:
  - `documents` ✅, `departments` ✅, `users` ✅
  - `search_index_status` (baru) — ensure docs are indexed
- **Relasi Service**: Elasticsearch (multi-index search), PostgreSQL (folder tree search), Redis (cache recent searches)
- **Catatan**: Universal Search sebaiknya menggunakan Elasticsearch multi-index query untuk mencari across: documents (full-text + metadata), folder tree (path matching), users (name matching).

---

### 28. Drill Down Search Workspace
- **Status**: ✅ Frontend ada (built-in), ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/documents/index.vue` — bagian folder tree sidebar (baris 19-70) + explorer table (baris 183-238)
- **Deskripsi UI**: Sidebar kiri: interactive folder tree (expandable, multi-level: Root → Company → Dept → Year → Month) dengan document count per folder. Main area: data table (checkbox select, name+ID, type badge, date, dept, status dot, actions: view/add-to-cart). Floating cart bar muncul saat ada item selected.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/folders/tree` — recursive folder tree structure
  - `GET /api/v1/folders/:id/documents?page=&limit=&sort=` — documents per folder with pagination
  - `GET /api/v1/folders/:id/children` — sub-folders
  - `GET /api/v1/folders/:id/stats` — document count per folder node
- **Handler yang dibutuhkan**: Tambah di `document_handler.go` atau buat `folder_handler.go`
- **Service yang dibutuhkan**: `folder_service.go`
- **Database Tables**:
  - `companies` ✅ → `branches` ✅ → `departments` ✅ (folder hierarchy)
  - `documents` ✅ — content per folder
  - `loan_cart_items` (baru) — shopping cart state
- **Relasi Service**: PostgreSQL (recursive CTE for tree), Redis (cache folder stats)
- **Catatan**: Folder tree sebenarnya dibangun dari hirarki organisasi (Company → Branch → Department → [Year] → [Month]). Year/Month level dihasilkan dari `documents.created_at`. Tree perlu recursive query (CTE) di PostgreSQL. Floating cart bar (baris 242-264) terhubung ke loan workflow.

---

### 29. Document Detail & Related Sidebar
- **Status**: ✅ Frontend ada (sangat lengkap, 450 baris), ⚠️ Backend parsial
- **Vue Files**:
  - `frontend/app/pages/documents/[id].vue` (450 baris — file kedua terbesar)
- **Deskripsi UI**: 2 mode — Normal mode: PDF viewer toolbar (page nav, zoom, print, download, fullscreen), mock PDF content, Actions bar (Add to Cart, Request Digital Copy), Metadata table, Physical Location card (building, rack, box + QR code), Related Documents list, Loan History timeline. Enlarged mode: full-screen split (kiri: PDF viewer, kanan: collapsible info panel dengan metadata, location, related docs, loan history, action buttons).
- **Backend API**: ⚠️ Parsial — yang ada:
  - `GET /api/v1/documents/:id/preview` ✅ (preview file)
  - Perlu ditambah:
    - `GET /api/v1/documents/:id` — full detail (metadata, location, tags)
    - `GET /api/v1/documents/:id/related` — related documents
    - `GET /api/v1/documents/:id/loan-history` — riwayat peminjaman
    - `GET /api/v1/documents/:id/versions` — daftar versi
    - `POST /api/v1/documents/:id/cart` — tambah ke loan cart
    - `POST /api/v1/documents/:id/request-digital` — request digital copy
    - `GET /api/v1/documents/:id/location` — lokasi fisik lengkap (building → room → rack → box → folder)
- **Handler yang ada**: `document_handler.go` ✅ (perlu extend banyak method)
- **Service yang dibutuhkan**: Extend `document_service.go`
- **Database Tables**:
  - `documents` ✅, `document_versions` ✅ — core data
  - `borrow_requests` ✅ — loan history
  - `loan_cart_items` (baru) — add to cart
  - `document_tracking_events` (baru) — audit trail
  - `racks` ✅, `boxes` ✅, `ordners` ✅ — physical location
  - `rfid_tags` ✅ — QR/RFID link
  - `file_storage_objects` (baru) — MinIO file reference
- **Relasi Service**: PostgreSQL, MinIO (PDF preview/download), Elasticsearch (related docs via similarity)
- **Catatan**: Ini adalah view paling kaya fitur (450 baris). Enlarged mode split-screen sangat mirip dengan PDF readers profesional (Adobe Acrobat style). Related docs bisa diambil dari: (1) same folder, (2) ES text similarity, (3) manual links. Loan History menampilkan timeline visual dengan avatar users. Watermark di PDF preview menampilkan: `{username} - NIK: {nik} - {timestamp}`.

---

### 30. Relationship Management Workbench
- **Status**: ✅ Frontend ada (built-in di document detail), ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/documents/[id].vue` — bagian Related Documents (baris 162-179, 365-380)
- **Deskripsi UI**: List related documents dengan icon, nama, deskripsi relasi (Attached-Linked, Reference-Digital Only), hover effect. Tersedia di normal mode dan enlarged mode sidebar.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/documents/:id/related` — list dokumen terkait
  - `POST /api/v1/documents/:id/related` — tambah relasi manual
  - `DELETE /api/v1/documents/:id/related/:relatedId` — hapus relasi
  - `GET /api/v1/documents/:id/related/suggestions` — AI-suggested relations
- **Handler**: Extend `document_handler.go`
- **Service**: Extend `document_service.go`
- **Database Tables**: `document_relations` (baru), `documents` ✅
- **Relasi Service**: PostgreSQL, Elasticsearch (similarity-based suggestions)

---

### 31. Search Analytic & Query Logs
- **Status**: ❌ Frontend belum ada (placeholder), ❌ Backend belum ada
- **Vue File**: Belum ada — perlu buat `frontend/app/pages/admin/analytics/search.vue`
- **Deskripsi UI**: Dashboard analytics: top search terms, failed queries, avg response time, search volume chart, user search patterns. Export ke CSV/PDF.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/analytics/search/top-queries` — top search terms
  - `GET /api/v1/analytics/search/failed` — failed/zero-result queries
  - `GET /api/v1/analytics/search/volume?range=7d` — volume chart data
  - `GET /api/v1/analytics/search/export` — export CSV
- **Handler**: `analytics_handler.go` (baru)
- **Service**: `analytics_service.go` (baru)
- **Database Tables**: `search_query_logs` (baru), `search_index_status` ✅
- **Relasi Service**: PostgreSQL, Elasticsearch (query logs), Redis (cached aggregations)

---

### 32. Search Context Commit
- **Status**: ✅ Frontend ada (built-in di documents), ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/documents/index.vue` — floating cart bar (baris 242-264)
- **Deskripsi UI**: Saat dokumen di-select via checkbox, floating cart bar muncul di bawah menampilkan jumlah item + tombol "View Cart" yang navigate ke `/loans/cart`.
- **Backend API**: ❌ Belum ada — perlu:
  - `POST /api/v1/cart/items` — tambah dokumen ke cart
  - `GET /api/v1/cart` — list items di cart
  - `DELETE /api/v1/cart/items/:id` — hapus dari cart
  - `POST /api/v1/cart/checkout` — submit cart sebagai loan request
- **Handler**: `cart_handler.go` (baru)
- **Service**: `cart_service.go` (baru)
- **Database Tables**: `loan_cart_items` (baru), `borrow_requests` ✅, `documents` ✅
- **Relasi Service**: PostgreSQL, Redis (session cart)

---

### 34. Intelligent Audit Mission Planner
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/stock/missions.vue` (176 baris)
- **Deskripsi UI**: Stats grid (Pending Approval: 12, Active Missions: 8, Accuracy Rate: 94%), tab filter (All/Scheduled/Completed), missions table (target zone, deadline, operator+avatar, type badge: SCHEDULED/SPOT CHECK, status, actions). Pagination + info tip box.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/stock/missions?tab=&page=&limit=` — list missions
  - `POST /api/v1/stock/missions` — buat mission baru
  - `PUT /api/v1/stock/missions/:id` — update mission
  - `GET /api/v1/stock/missions/stats` — dashboard stats
  - `POST /api/v1/stock/missions/:id/approve` — approve mission
  - `DELETE /api/v1/stock/missions/:id` — cancel mission
- **Handler**: `stock_handler.go` (baru)
- **Service**: `stock_service.go` (baru)
- **Database Tables**: `stock_opname_missions` (baru, migration 000010), `stock_opname_items` (baru), `users` ✅
- **Relasi Service**: PostgreSQL, Asynq/Redis (scheduled mission jobs)

---

### 35. RFID Mobile Blind Audit Interface
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/rfid/blind-audit.vue` (New)
- **Deskripsi UI**: Antarmuka mobile untuk audit fisik menggunakan RFID (Blind Audit Mode). Menampilkan:
  - **Audit Progress Circle**: Visualisasi jumlah item terpindai vs target, serta persentase penyelesaian misi.
  - **Auditor & Device Info**: Identitas auditor aktif dan ID perangkat RFID yang terhubung.
  - **Location Lock (QR Scan)**: Mekanisme penguncian lokasi audit melalui pemindaian QR Rack sebelum pemindaian item dimulai.
  - **Live Scan Feed (Blind Mode)**: Daftar hasil pemindaian real-time dengan kategori status (MATCH, MISSING, EXTRA) tanpa menampilkan nama dokumen (blind mode).
  - **Handheld Action Bar**: Tombol aksi cepat untuk memicu scan rack, scan item, pause, dan finish mission.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/audit/missions/:id` — get mission details and targets
  - `POST /api/v1/audit/scan` — submit scanned RFID tag
  - `POST /api/v1/audit/location` — submit scanned Rack QR
- **Handler**: `rfid_handler.go` (baru)
- **Service**: `rfid_service.go` (baru)
- **Database Tables**: `audit_missions` ✅, `audit_scans` (baru)
- **Relasi Service**: RFID Hardware Agent, WebSocket (untuk real-time feed)

---

### 36. Real Time Discrepancy Alert & Resolution
- **Status**: ✅ Frontend ada (3-state view), ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/stock/reconciliation.vue` (397 baris)
- **Deskripsi UI**: 3-state flow: (1) **Report** — summary cards (Match: 142, On Loan: 5, Missing: 3, Extra: 1), inventory table (SKU, sys vs phy qty, variance, status badge, location), tab filter, export + investigate buttons. (2) **Resolution** — item detail, resolution actions (relocate/write-off), actor info. (3) **Validation** — unresolved missing items table with status dropdown, digital signature pad + PIN input, approve/reject buttons, audit trail.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/stock/reconciliation/:missionId` — reconciliation report
  - `GET /api/v1/stock/reconciliation/:missionId/items?status=` — filtered items
  - `PUT /api/v1/stock/reconciliation/:missionId/items/:id/resolve` — resolve discrepancy
  - `POST /api/v1/stock/reconciliation/:missionId/approve` — head approval + signature
  - `GET /api/v1/stock/reconciliation/:missionId/export` — export PDF/CSV
- **Handler**: Tambah di `stock_handler.go`
- **Service**: Tambah di `stock_service.go`
- **Database Tables**: `stock_opname_missions` ✅, `stock_opname_items` (baru), `approval_workflows` (baru), `documents` ✅
- **Relasi Service**: PostgreSQL, MinIO (signature upload)

---

### 37. Security Trimming Search Policy (F-35)
- **Status**: ✅ Frontend ada (built-in role UI), ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/roles.vue` (terkait permission matrix)
- **Deskripsi UI**: Role-based access control yang menentukan dokumen mana yang visible di search results per user/role. Terkait dengan Dynamic Filter (#25) — ACL filtering diterapkan di query level.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/security/search-policies` — list kebijakan search per role
  - `PUT /api/v1/security/search-policies/:roleId` — update search policy
  - Middleware: search query harus filter berdasarkan `document_role_access`
- **Handler**: `security_handler.go` (baru)
- **Service**: `security_service.go` (baru)
- **Database Tables**: `document_role_access` ✅, `roles` ✅, `user_roles` ✅
- **Relasi Service**: PostgreSQL, Elasticsearch (filtered queries)
- **Catatan**: Security trimming = setiap search query di-filter oleh user's role permissions. Dokumen yang tidak boleh dilihat oleh role tertentu tidak muncul di search results.

---

### 38. Encryption Compliance Dashboard
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/compliance.vue` (178 baris)
- **Deskripsi UI**: Layout 3 kolom — kiri: Verification status cards (TLS 1.3: Active, SSL Certificate: Critical, AES-256: Verified, Security Patch: Warning, Audit Trail: Synced). Tengah: Compliance Score besar + History timeline + Compliance Matrix table. Kanan: Recent Findings + Quick Actions.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/compliance/status` — overall compliance status
  - `GET /api/v1/compliance/checks` — list verification checks
  - `POST /api/v1/compliance/checks/run` — trigger compliance scan
  - `GET /api/v1/compliance/history` — compliance score history
  - `GET /api/v1/compliance/export` — export report
- **Handler**: `compliance_handler.go` (baru)
- **Service**: `compliance_service.go` (baru)
- **Database Tables**: `compliance_checks` (baru), `compliance_history` (baru)
- **Relasi Service**: PostgreSQL, external TLS/SSL checker

---

### 39. Granular Access Policy Matrix
- **Status**: ✅ Frontend ada, ⚠️ Backend parsial (role CRUD ada)
- **Vue File**: `frontend/app/pages/admin/roles.vue` (terkait permission matrix)
- **Deskripsi UI**: Permission matrix per role: module-level access (read/write/delete/admin) untuk setiap fitur DMS. Drag-drop role priority, inheritance rules.
- **Backend API**: ⚠️ Parsial — role CRUD ada, perlu:
  - `GET /api/v1/roles/:id/permissions` — detailed permission matrix
  - `PUT /api/v1/roles/:id/permissions` — batch update permissions
  - `GET /api/v1/roles/:id/effective-permissions` — computed effective permissions (incl. inheritance)
- **Handler**: Extend `role_handler.go`
- **Service**: Extend `role_service.go`
- **Database Tables**: `roles` ✅, `role_permissions` ✅, `document_role_access` ✅
- **Relasi Service**: PostgreSQL

---

### 40. RFID Gate Monitor
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/rfid/gate-monitor.vue` (New)
- **Deskripsi UI**: Konsol pemantauan keamanan gerbang RFID (Gate Security) secara real-time. Menampilkan:
  - **Live Event Stream**: Daftar log aktivitas RFID (Tag ID, User/Doc, Status Authorized/Unauthorized) yang masuk secara terus-menerus.
  - **Incident Alert Module**: Panel peringatan kritis saat terjadi pelanggaran (Invalid Exit), lengkap dengan informasi subjek (User), kelas dokumen, waktu deteksi milidetik, dan Gate ID.
  - **Hardware Diagnostics**: Status kesehatan perangkat keras gerbang (Antenna Power, DB Latency) dan heatmap frekuensi lalu lintas.
  - **Command & Control Panel**: Tombol intervensi cepat untuk Acknowledge Alarm, Lock User Access, dan Notify Security.
  - **System Overrides**: Toggle untuk Emergency Open dan Hard Lockdown pada gerbang fisik.
- **Backend API**: ❌ Belum ada — perlu:
  - `WS /api/v1/rfid/gate/stream` — real-time event stream from local gate controller
  - `POST /api/v1/rfid/gate/control` — execute lockdown or emergency open
  - `POST /api/v1/rfid/gate/acknowledge` — clear current alarm
- **Handler**: `rfid_handler.go` (baru)
- **Service**: `rfid_service.go` (baru)
- **Relasi Service**: RFID Hardware Gateway, WebSocket, Notification Service (Push to Security)

---

### 41. History Log Explorer
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/logs/explorer.vue` (New)
- **Deskripsi UI**: Antarmuka eksplorasi riwayat aktivitas sistem yang komprehensif. Menampilkan:
  - **Advanced Filter Bar**: Pencarian berdasarkan tipe event, administrator, dan rentang tanggal yang presisi.
  - **System Activity Table**: Daftar log kronologis dengan informasi user (avatar), tipe event (Badge color-coded), dan dokumen terkait.
  - **Event Detail Sidebar**: Panel detail mendalam mencakup:
    - **Action Performed**: Ringkasan tindakan (misal: Document Deletion).
    - **Document Info**: Nama file, repository path, dan Object GUID.
    - **User Metadata**: Login ID, access level, departemen, dan Terminal ID.
    - **Network Details**: IP Address, MAC Address, dan Browser Agent String.
    - **Blockchain Hash Evidence**: Bukti integritas log yang terdaftar di private chain (SHA-256).
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/logs` — search and filter audit logs
  - `GET /api/v1/logs/:id/detail` — get full log metadata and blockchain hash
- **Handler**: `audit_handler.go` (baru)
- **Service**: `audit_service.go` (baru)
- **Database Tables**: `audit_logs` ✅, `blockchain_evidence` (baru)
- **Relasi Service**: PostgreSQL, Elasticsearch, Blockchain Node (Immutable Ledger)

---

### 42. Version Timeline & Checkout Workbench / Riwayat Versi & Pusat Kontrol
- **Status**: ✅ Frontend parsial (upload + document detail), ❌ Backend belum ada
- **Vue Files**: `frontend/app/pages/documents/upload.vue` (347 baris — manual + bulk upload), `frontend/app/pages/documents/[id].vue` (metadata: version v2.1 Final)
- **Deskripsi UI**: Upload page mendukung manual entry dan bulk Excel import. Document detail menampilkan version info. Perlu: version timeline view, checkout/checkin controls, diff viewer antar versi.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/documents/:id/versions` — list semua versi
  - `POST /api/v1/documents/:id/versions` — upload versi baru
  - `POST /api/v1/documents/:id/checkout` — checkout (lock for editing)
  - `POST /api/v1/documents/:id/checkin` — checkin (release lock + new version)
  - `GET /api/v1/documents/:id/versions/:versionId/download` — download versi tertentu
- **Handler**: Extend `document_handler.go`
- **Service**: Extend `document_service.go`
- **Database Tables**: `document_versions` ✅, `documents` ✅, `file_storage_objects` (baru)
- **Relasi Service**: PostgreSQL, MinIO (versioned file storage)

---

### 43. Secure Redaction & Annotation
- **Status**: ❌ Frontend belum ada, ❌ Backend belum ada
- **Vue File**: Belum ada — perlu buat `frontend/app/pages/documents/[id]/redact.vue`
- **Deskripsi UI**: PDF viewer dengan overlay tools: redaction boxes (permanent black-out), highlight annotations, text notes, stamp (CONFIDENTIAL/APPROVED). Save as new version.
- **Backend API**: ❌ Belum ada — perlu:
  - `POST /api/v1/documents/:id/redact` — apply redactions (permanent)
  - `POST /api/v1/documents/:id/annotate` — add annotations (non-destructive)
  - `GET /api/v1/documents/:id/annotations` — list annotations
  - `DELETE /api/v1/documents/:id/annotations/:annotId` — remove annotation
- **Handler**: Extend `document_handler.go`
- **Service**: `redaction_service.go` (baru)
- **Database Tables**: `document_annotations` (baru), `document_versions` ✅, `documents` ✅
- **Relasi Service**: PostgreSQL, MinIO (redacted PDF storage), PDF processing library

---

### 44. Secure Preview with Dynamic Watermark
- **Status**: ✅ Frontend parsial (built-in di document detail), ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/documents/[id].vue` — preview area (baris 70-100)
- **Deskripsi UI**: Document viewer menampilkan preview dokumen. Watermark perlu ditambahkan secara dinamis (nama user + timestamp + "CONFIDENTIAL" overlay) saat user melihat dokumen sensitif.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/documents/:id/preview?watermark=true` — preview dengan dynamic watermark
  - `GET /api/v1/documents/:id/preview/config` — watermark config per security level
- **Handler**: Extend `document_handler.go`
- **Service**: Extend `document_service.go`
- **Database Tables**: `documents` ✅, `document_role_access` ✅
- **Relasi Service**: PostgreSQL, MinIO (file streaming), PDF watermark library (Go)

---

### 45. Multi channel Notification Orchestrator
- **Status**: ✅ Frontend ada (notification center), ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/notifications/index.vue` (917 bytes — placeholder)
- **Deskripsi UI**: Notification center — list notifikasi per user. Perlu: multi-channel (email, in-app, WhatsApp), notification preferences, template management.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/notifications?page=&limit=&read=` — list notifikasi
  - `PUT /api/v1/notifications/:id/read` — mark as read
  - `PUT /api/v1/notifications/read-all` — mark all read
  - `GET /api/v1/notifications/preferences` — user notification preferences
  - `PUT /api/v1/notifications/preferences` — update preferences
- **Handler**: `notification_handler.go` (baru)
- **Service**: `notification_service.go` (baru)
- **Database Tables**: `notifications` (baru), `notification_preferences` (baru), `users` ✅
- **Relasi Service**: PostgreSQL, Redis (real-time push), SMTP (email), WhatsApp API

---

### 46. Extension Request Workbench
- **Status**: ✅ Frontend ada (2 halaman), ❌ Backend belum ada
- **Vue Files**: `frontend/app/pages/loans/my.vue` (334 baris — extension modal), `frontend/app/pages/approvals/extensions.vue` (285 baris — approval queue)
- **Deskripsi UI**: **my.vue**: User melihat loan aktif, time remaining card, tracking timeline (Approved → Ready → On Loan → Extension Requested), modal request extension (days + reason + info box). **extensions.vue**: Approval queue table (req no, requestor+avatar, current due, requested days, new due, reason, risk flag), detail panel (docs linked, return history, prior extensions, risk summary), PIN confirmation modal.
- **Backend API**: ❌ Belum ada — perlu:
  - `POST /api/v1/loans/:id/extension` — submit extension request
  - `GET /api/v1/approvals/extensions?page=&limit=` — list pending extensions
  - `PUT /api/v1/approvals/extensions/:id/approve` — approve with PIN
  - `PUT /api/v1/approvals/extensions/:id/reject` — reject with notes
  - `GET /api/v1/approvals/extensions/:id/risk` — risk assessment
- **Handler**: `loan_handler.go` (baru), extend `approval_handler.go`
- **Service**: `loan_service.go` (baru), `approval_service.go` (baru)
- **Database Tables**: `borrow_requests` ✅, `borrow_extensions` (baru), `approval_workflows` (baru), `users` ✅
- **Relasi Service**: PostgreSQL, notification_service (alert approver)

---

### 47. Loan Timer & Reminder Monitor
- **Status**: ✅ Frontend ada (built-in di my loans), ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/loans/my.vue` — time remaining card (baris 96-102), tabs (All/Due Soon/Overdue)
- **Deskripsi UI**: Big card "Time Remaining" (countdown, expires date), tab filter (All/Due Soon/Overdue), penalty warning box, auto-notification policies.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/loans/my?tab=&page=&limit=` — user's active loans
  - `GET /api/v1/loans/overdue` — overdue loans (admin view)
  - Background: Asynq job to send reminder at D-3, D-1, D+0
- **Handler**: Extend `loan_handler.go`
- **Service**: Extend `loan_service.go`
- **Database Tables**: `borrow_requests` ✅, `notifications` (baru), `users` ✅
- **Relasi Service**: PostgreSQL, Asynq/Redis (scheduled reminder jobs), notification_service

---

### 48. Pre Registration & Handover Form / Pendaftaran Dokumen Baru
- **Status**: ✅ Frontend ada (multi-state flow), ❌ Backend belum ada
- **Vue Files**: `frontend/app/pages/registration/staging.vue` (289 baris), `frontend/app/pages/registration/migration.vue` (530 baris)
- **Deskripsi UI**: **staging.vue**: Stats (Queued: 124, Processing: 12, Ready: 85), upload zone, staging table (filename, pages, OCR status + progress bar, operator+avatar), OCR engine status (ABBYY Finereader 99.2%), priority card. **migration.vue**: Multi-state wizard (initial → review → mandatory → duplicate_check → location_assignment → print_qr → success), split-screen OCR viewer + metadata form.
- **Backend API**: ❌ Belum ada — perlu:
  - `POST /api/v1/registration/upload` — upload dokumen baru
  - `GET /api/v1/registration/staging?page=&limit=` — list staging queue
  - `GET /api/v1/registration/staging/stats` — stats
  - `POST /api/v1/registration/:id/process` — start OCR processing
  - `PUT /api/v1/registration/:id/metadata` — update metadata results
  - `POST /api/v1/registration/:id/commit` — final commit
- **Handler**: `registration_handler.go` (baru)
- **Service**: `registration_service.go` (baru)
- **Database Tables**: `documents` ✅, `document_versions` ✅, `ocr_results` (baru), `racks` ✅
- **Relasi Service**: PostgreSQL, MinIO (file upload), Asynq/Redis (OCR job queue), OCR Worker

---

### 49. Distribution Center
- **Status**: ✅ Frontend ada (multi-state flow), ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/circulation/incoming.vue` (653 baris)
- **Deskripsi UI**: Multi-state: (1) **List** — incoming document queue. (2) **Validation** — split-screen: left = document viewer (zoom, rotate, OCR text view), right = metadata validation form with OCR-highlighted fields (amber dashed borders). Full MOU/contract preview with header, signatories, witness section.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/circulation/incoming?page=&limit=` — incoming queue
  - `PUT /api/v1/circulation/:id/validate` — validate metadata
  - `PUT /api/v1/circulation/:id/route` — route to department
  - `GET /api/v1/circulation/:id/ocr` — get OCR text
- **Handler**: `circulation_handler.go` (baru)
- **Service**: `circulation_service.go` (baru)
- **Database Tables**: `documents` ✅, `document_circulations` ✅, `routing_slip_recipients` ✅, `departments` ✅
- **Relasi Service**: PostgreSQL, MinIO (document viewer), OCR Worker

---

### 50. Physical Loan Desk
- **Status**: ✅ Frontend selesai dioverhaul, ⚠️ Backend parsial (queue dari loans endpoint)
- **Vue File**: `frontend/app/pages/circulation/checkout.vue`
- **Deskripsi UI**: Layout 2-panel — kiri: antrean serah terima (filter `l2_approved`), klik untuk pilih request. Kanan: verification flow — (1) checklist identitas peminjam (KTP/Badge & kode pickup), (2) tabel dokumen per baris dengan checkbox `handedOver` + kolom kondisi & catatan, (3) area tanda tangan digital, (4) konfirmasi serah terima. Tombol "Complete Handover" hanya aktif jika semua dokumen di-checklist dan konfirmasi dicentang.
- **Backend API**: 
  - ✅ `GET /api/v1/loans` — fetch queue (filter `l2_approved` di frontend)
  - ❌ Perlu ditambah:
    - `POST /api/v1/loans/:id/complete-handover` — finalisasi serah terima + ubah status ke `active`
    - `GET /api/v1/loans/:id/pickup-slots` — available time slots
- **Handler**: `loan_handler.go`
- **Service**: `loan_service.go`
- **Database Tables**: `borrow_requests` ✅, `loan_request_items` ✅
- **Relasi Service**: PostgreSQL, notification_service
- **i18n**: ✅ `frontend/app/locales/*/circulation.json` key `circulation.checkout.*`

---

### 51. Two Level Approval Console
- **Status**: ✅ Frontend ada (L1 + L2 views), ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/approvals/loans.vue` (496 baris)
- **Deskripsi UI**: Dynamic L1/L2 view — stat cards (Waiting L1: 24 / Waiting L2: 18, Approved: 12 / Release Ready: 8, Rejected, Avg Time), request table (req no, requestor+avatar, department, doc title, docs count, urgency badge, date, status), detail side panel (documents list, metadata, notes), PIN verification modal, batch approve/reject footer.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/approvals/loans?level=L1&page=&limit=` — list pending approvals
  - `GET /api/v1/approvals/loans/stats?level=` — dashboard stats
  - `PUT /api/v1/approvals/loans/:id/approve` — approve (L1 or L2)
  - `PUT /api/v1/approvals/loans/:id/reject` — reject with reason
  - `POST /api/v1/approvals/loans/batch` — batch approve/reject
- **Handler**: `approval_handler.go` (baru)
- **Service**: `approval_service.go` (baru)
- **Database Tables**: `borrow_requests` ✅, `approval_workflows` (baru), `users` ✅
- **Relasi Service**: PostgreSQL, notification_service (alert next approver)
- **Catatan**: L1 = Department Head approval, L2 = Division/Legal Head final approval.

---

### 52. Fast Track Package Selector
- **Status**: ✅ Frontend ada (multi-step wizard), ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/loans/fast-track.vue` (556 baris)
- **Deskripsi UI**: Multi-step wizard: (1) **Form** — purpose selector (Internal/External Audit, Legal Review, Annual Tax), department filter, duration input (hari kerja), notes, step progress bar. (2) **Selection** — filtered documents grid, search, select all, cart summary. (3) **Review** — selected docs table, delivery info, final submit. (4) **Success** — confirmation.
- **Backend API**: ❌ Belum ada — perlu:
  - `POST /api/v1/loans/fast-track` — submit fast-track request (skip normal queue)
  - `GET /api/v1/loans/fast-track/packages` — available fast-track packages
  - `GET /api/v1/documents/search?purpose=&dept=` — filtered doc search
- **Handler**: Extend `loan_handler.go`
- **Service**: Extend `loan_service.go`
- **Database Tables**: `borrow_requests` ✅, `fast_track_packages` (baru), `documents` ✅
- **Relasi Service**: PostgreSQL, approval_service (auto-escalate)

---

### 53. Request Cart Workspace
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/loans/cart.vue` (136 baris)
- **Deskripsi UI**: 2 kolom — kiri: cart items list (checkbox, doc icon + color, doc ID + name, dept, availability badge, remove button), select all + clear cart. Kanan: loan summary (total docs, duration dropdown 7/14/30 days, purpose textarea, info box, checkout button, continue browsing button).
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/cart` — list cart items
  - `POST /api/v1/cart/items` — add item to cart
  - `DELETE /api/v1/cart/items/:id` — remove item
  - `DELETE /api/v1/cart` — clear cart
  - `POST /api/v1/cart/checkout` — proceed to checkout
- **Handler**: `cart_handler.go` (baru)
- **Service**: `cart_service.go` (baru)
- **Database Tables**: `loan_cart_items` (baru), `documents` ✅
- **Relasi Service**: PostgreSQL, Redis (session-based cart)

---

### 54. System Configuration Panel / Konfigurasi Distribusi Data
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/settings/sso.vue` (236 baris)
- **Deskripsi UI**: 3 kolom — kiri: connection settings (LDAP/AD server, port, base DN, bind DN, test connection button). Tengah: attribute mapping (username, email, department, full name fields), group mapping table (LDAP Group → DMS Role). Kanan: sync status cards (sync schedule, last sync time, total synced users, conflict resolution log).
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/settings/sso` — get SSO config
  - `PUT /api/v1/settings/sso` — update SSO config
  - `POST /api/v1/settings/sso/test` — test LDAP connection
  - `POST /api/v1/settings/sso/sync` — trigger manual sync
  - `GET /api/v1/settings/sso/sync-status` — sync status
- **Handler**: `settings_handler.go` (baru)
- **Service**: `settings_service.go` (baru), `ldap_service.go` (baru)
- **Database Tables**: `system_settings` (baru), `ldap_sync_logs` (baru), `users` ✅
- **Relasi Service**: PostgreSQL, LDAP/AD Server, Asynq/Redis (scheduled sync)

---

### 55. Retention & Disposal Alert Center (F-53)
- **Status**: ✅ Frontend ada (10 halaman retention), ❌ Backend belum ada
- **Vue Files**: `frontend/app/pages/retention/approaching.vue` (181 baris), `decision.vue`, `batch.vue`, `history.vue`, `export.vue`, `shredding.vue`, `purge-confirm.vue`, `purge-success.vue`, `upload-bast.vue`, `log-detail.vue`
- **Deskripsi UI**: **approaching.vue**: Stats card (On Hold: 8), info banner, main table (checkbox, doc ID, title, dept, age in years, retention rule, status badge DUE FOR REVIEW/ON HOLD/OVERDUE, actions), pagination. Sub-pages: decision (approve/reject disposal), batch (bulk disposal), shredding (physical destruction tracking), export (disposal certificate PDF), upload-bast (upload Berita Acara Serah Terima).
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/retention/approaching?page=&limit=&status=` — docs approaching retention limit
  - `GET /api/v1/retention/stats` — retention stats
  - `PUT /api/v1/retention/:id/decision` — approve/reject disposal
  - `POST /api/v1/retention/batch` — batch disposal
  - `POST /api/v1/retention/:id/hold` — place on legal hold
  - `GET /api/v1/retention/history` — disposal history
  - `GET /api/v1/retention/:id/certificate` — download disposal certificate
- **Handler**: `retention_handler.go` (baru)
- **Service**: `retention_service.go` (baru)
- **Database Tables**: `retention_policies` ✅, `disposal_records` (baru), `documents` ✅
- **Relasi Service**: PostgreSQL, MinIO (BAST upload), Asynq/Redis (scheduled retention checks)

---

### 56. Backup & One Click Recovery Center
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/backup.vue` (290 baris)
- **Deskripsi UI**: Backup schedule controls, retention period settings, last backup info, encryption status (AES-256), one-click restore button, backup history table, storage usage chart.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/backup/status` — current backup status
  - `POST /api/v1/backup/trigger` — trigger manual backup
  - `POST /api/v1/backup/restore/:snapshotId` — one-click restore
  - `GET /api/v1/backup/history` — backup history
  - `PUT /api/v1/backup/schedule` — update backup schedule
- **Handler**: `backup_handler.go` (baru)
- **Service**: `backup_service.go` (baru)
- **Database Tables**: `backup_snapshots` (baru), `system_settings` (baru)
- **Relasi Service**: PostgreSQL, MinIO (backup storage), Asynq/Redis (scheduled backups)

---

### 57. User Management Panel / Manajemen Pengguna
- **Status**: ✅ Frontend ada (full CRUD + detail), ⚠️ Backend parsial
- **Vue Files**: `frontend/app/pages/admin/users/index.vue` (246 baris), `hierarchy.vue` (reporting structure)
- **Deskripsi UI**: Stats grid (Total: 1,248, Active: 1,120, Privileged: 42, Inactive: 128), user table (role badge, dept, status, last login + IP, reporting), detail sidebar (profile card, action grid: reset/role/reporting/deactivate, system info: LDAP sync/access dots/verified, reporting structure, recent activity logs timeline).
- **Backend API**: ⚠️ Parsial — user CRUD ada, perlu extend:
  - `GET /api/v1/users?page=&limit=&role=&dept=&status=` — paginated users
  - `GET /api/v1/users/:id` — user detail
  - `PUT /api/v1/users/:id/role` — change role
  - `PUT /api/v1/users/:id/deactivate` — deactivate user
  - `POST /api/v1/users/:id/reset-password` — reset password
  - `GET /api/v1/users/:id/activity-log` — recent activity
  - `GET /api/v1/users/hierarchy` — reporting hierarchy tree
- **Handler**: Extend `user_handler.go`
- **Service**: Extend `user_service.go`
- **Database Tables**: `users` ✅, `user_roles` ✅, `departments` ✅, `user_activity_logs` (baru)
- **Relasi Service**: PostgreSQL, LDAP (sync)

---

### 58. Warehouse Occupancy Analytics / Analisis Keteisian Gudang
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/warehouse/monitoring.vue` (174 baris)
- **Deskripsi UI**: Stats grid (Total docs: 1,248 +2% MoM, Capacity: 82.4%/1,028 units Optimal, Reserved: 12 Critical, Accuracy: 96.8%), heatmap grid (128 cells color-coded: empty/low/medium/high/reserved), recommendations table (rack ID + zone, dept, reason, suitability score %, apply button).
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/warehouse/occupancy/stats` — occupancy statistics
  - `GET /api/v1/warehouse/occupancy/heatmap` — heatmap data
  - `GET /api/v1/warehouse/occupancy/recommendations` — AI-based rack recommendations
  - `POST /api/v1/warehouse/occupancy/recommendations/:id/apply` — apply recommendation
- **Handler**: `warehouse_handler.go` (baru)
- **Service**: `warehouse_service.go` (baru)
- **Database Tables**: `racks` ✅, `rack_slots` ✅, `documents` ✅, `departments` ✅
- **Relasi Service**: PostgreSQL, AI/Elasticsearch (placement recommendations)




---

### 59. Comprehensive Audit Trail Explorer / Manajemen Audit (F-48)
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/compliance.vue` (178 baris)
- **Deskripsi UI**: Audit logs grid menampilkan history aktivitas user (Who, When, What, IP, Status). Filter berdasarkan rentang waktu, departemen, dan jenis aksi. Summary cards untuk verifikasi sistem (TLS, SSL, AES Encryption, Patch Status).
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/audit/logs?page=&limit=&user_id=&action=` — list audit trails
  - `GET /api/v1/audit/stats` — audit summary stats
  - `GET /api/v1/audit/export` — export audit to Excel/PDF
- **Handler**: `audit_handler.go` (baru)
- **Service**: `audit_service.go` (baru)
- **Database Tables**: `audit_logs` (baru), `users` ✅
- **Relasi Service**: PostgreSQL (Log Storage)

---

### 60. Operational Circulation Report (F-46)
- **Status**: ❌ Frontend & Backend belum ada
- **Vue File**: `frontend/app/pages/reports/circulation.vue` (Planned)
- **Deskripsi UI**: Dashboard laporan operasional sirkulasi dokumen. Visualisasi tren peminjaman, dokumen paling aktif, dan bottleneck sirkulasi.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/reports/circulation/summary` — ringkasan sirkulasi
  - `GET /api/v1/reports/circulation/trend` — tren data bulanan
- **Handler**: `report_handler.go` (baru)
- **Service**: `report_service.go` (baru)
- **Database Tables**: `borrow_requests` ✅, `documents` ✅
- **Relasi Service**: PostgreSQL, Redis (Caching Report)

---

### 61. User Scorecard & Activity Timeline
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/users/scorecard.vue` (New)
- **Deskripsi UI**: Antarmuka profil keamanan dan produktivitas pengguna. Menampilkan:
  - **User Selection Sidebar**: Daftar direktori pengguna berdasarkan departemen dan ID.
  - **Chronological Activity Stream**: Timeline aktivitas kritis (Policy violations, Sensitive access, Approvals, Overdue behavior) dengan tombol aksi cepat (Lock Account, dsb).
  - **Security Metrics Side Panel**: Statistik akumulatif akses, percobaan dokumen sensitif, dan performa SLA persetujuan.
  - **Risk Assessment Module**: Skor risiko pengguna secara dinamis (misal: 720 - Moderate) berdasarkan perilaku yang terdeteksi sistem AI.
  - **Compliance Note System**: Form untuk menyematkan catatan kepatuhan pada profil pengguna.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/users/:id/scorecard` — get user metrics and risk score
  - `GET /api/v1/users/:id/timeline` — get activity log
  - `POST /api/v1/users/:id/lock` — lock account due to violation
- **Handler**: `user_handler.go` (baru)
- **Service**: `user_service.go` (baru)
- **Database Tables**: `user_activity_logs` ✅, `user_risk_scores` (baru)
- **Relasi Service**: PostgreSQL, Redis (untuk tracking velocity threshold)

---

### 62. Department Performance Dashboard
- **Status**: ❌ Frontend & Backend belum ada
- **Vue File**: `frontend/app/pages/dashboard/dept-performance.vue` (Planned)
- **Deskripsi UI**: Dashboard perbandingan performa antar departemen dalam pengelolaan dokumen (peminjaman vs pengembalian tepat waktu).
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/reports/departments/efficiency` — efisiensi per departemen
- **Handler**: Extend `report_handler.go`
- **Service**: Extend `report_service.go`

---

### 63. Executive KPI Dashboard
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/dashboard/executive.vue` (New)
- **Deskripsi UI**: Dashboard analitik tingkat tinggi untuk manajemen puncak. Menampilkan:
  - **Executive Summary Cards**: Metrik utama seperti Total Active Archive (PB), Monthly Growth %, Active Loans count, Overdue Count (Risk), dan Compliance Score.
  - **Growth by Department Chart**: Visualisasi perbandingan pertumbuhan data antar unit bisnis utama (Operations vs Legal).
  - **Role Activity Monitor**: Distribusi aktivitas sistem berdasarkan peran pengguna (Admin, Auditor, dsb).
  - **Approval SLA Tracker**: Pemantauan rata-rata waktu respon persetujuan dan persentase keberhasilan SLA.
  - **Real-time Risk Alerts Feed**: Tabel insidensi keamanan dan operasional dengan kategori keparahan (Critical, Medium, Low).
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/dashboard/executive/summary` — aggregate metrics
  - `GET /api/v1/dashboard/executive/growth` — data for growth chart
  - `GET /api/v1/dashboard/executive/risk-alerts` — live alert feed
- **Handler**: `dashboard_handler.go` (baru)
- **Service**: `dashboard_service.go` (baru)
- **Database Tables**: `system_stats_daily` (baru), `risk_alerts` (baru)
- **Relasi Service**: PostgreSQL, Redis (untuk pre-aggregated analytics)

---

### 64. No Code Master Data Console / Manajemen PT 
- **Status**: ✅ Frontend ada, ✅ Backend ada (Full CRUD)
- **Vue File**: `frontend/app/pages/config/company.vue` (New)
- **Deskripsi UI**: Konsol administrasi master data dengan pendekatan low-code/no-code. Menampilkan:
  - **Entity Sidebar**: Navigasi antar entitas master (PT, Departemen, Tipe Dokumen, Lokasi, Retensi).
  - **Inline Editing Table**: Tabel manajemen entitas (PT) dengan dukungan penambahan baris langsung di tabel, status verifikasi NPWP, dan aksi CRUD.
  - **Validation Rules Engine**: Visualisasi aturan validasi (Naming convention, Tax compliance) yang aktif pada entitas tersebut.
  - **Pending Changes Monitor**: Daftar perubahan yang belum dipublikasi ke server produksi.
  - **Publishing Control**: Tombol sinkronisasi massal untuk menerapkan perubahan draf ke sistem inti.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/master-data/pt` — list companies
  - `POST /api/v1/master-data/pt/publish` — batch commit pending changes
  - `GET /api/v1/master-data/rules` — get validation rules
- **Handler**: `master_data_handler.go` (baru)
- **Service**: `master_data_service.go` (baru)
- **Database Tables**: `master_companies` ✅, `master_validation_rules` (baru)
- **Relasi Service**: PostgreSQL, LDAP (untuk validasi data user/dept)
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/master-data/entities` — list entities
  - `POST /api/v1/master-data/entities/:type` — create entry
- **Handler**: `master_data_handler.go` (baru)
- **Service**: `master_data_service.go` (baru)
- **Database Tables**: `master_companies`, `master_branches`

---

### 65. Scan to Web Launcher
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/scan/launcher.vue` (New)
- **Deskripsi UI**: Antarmuka pemicu pemindaian langsung dari browser. Menampilkan:
  - **Scanner Connection Panel**: Status koneksi perangkat (Connected/Disconnected), nama perangkat, alamat IP, versi driver, dan status sensor kertas (Paper Feed).
  - **Scan Configuration Form**: Pengaturan mode warna, resolusi (DPI), duplex scanning (bolak-balik), dan ukuran kertas (A4, F4, dsb).
  - **Warning & Instruction**: Panduan proses transfer data langsung ke cloud/browser.
  - **Mulai Scan Trigger**: Tombol eksekusi untuk mengirim sinyal perintah ke agent TWAIN/WIA lokal.
  - **Idle Monitor Area**: Visualisasi status standby scanner.
- **Backend API**: Bridge via WebSocket agent lokal.
- **Handler**: `scan_handler.go` (baru)
- **Service**: `scan_service.go` (baru)
- **Relasi Service**: WebSocket Local Agent (untuk komunikasi hardware)

---

### 66. Live Scan Result
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/scan/live.vue` (New)
- **Deskripsi UI**: Antarmuka preview hasil pemindaian secara real-time. Menampilkan:
  - **Thumbnails Sidebar**: Daftar halaman yang sedang dipindai dengan status (Success, Active, Error) dan progress bar scanning.
  - **Main Preview Toolbar**: Kontrol zoom, rotasi, navigasi halaman, tombol hapus/retry per halaman.
  - **Session Metadata**: Informasi teknis sesi scan (ID, waktu mulai, total ukuran file).
  - **Status Browser Intake**: Indikator koneksi aman (Secure Tunnel) ke engine scanner.
  - **Action Panel**: Tombol eksekusi hasil scan, simpan draf, atau pembatalan sesi.
- **Backend API**: ❌ Belum ada — perlu:
  - `WS /api/v1/scan/stream` — stream images from local scanner agent
  - `POST /api/v1/scan/finalize` — convert scanned images to single PDF
- **Handler**: `scan_handler.go` (baru)
- **Service**: `scan_service.go` (baru)
- **Relasi Service**: WebSocket Agent (Local), MinIO (Temporary Image Storage)

---

### 67. Document Versioning Center
- **Status**: ✅ Frontend ada (built-in in upload), ⚠️ Backend parsial
- **Vue File**: `frontend/app/pages/documents/upload.vue` (46862 bytes)
- **Deskripsi UI**: Interface untuk mengunggah revisi dokumen baru tanpa menghapus versi lama. Menampilkan history versi.
- **Backend API**: ⚠️ Parsial — perlu extend:
  - `GET /api/v1/documents/:id/versions` — list versions
  - `POST /api/v1/documents/:id/versions` — upload new version
- **Handler**: Extend `document_handler.go`
- **Service**: Extend `document_service.go`
- **Database Tables**: `document_versions` ✅
- **Relasi Service**: PostgreSQL, MinIO (Versioned Storage)

---

### 68. Version Detail & Compare
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/documents/compare.vue` (New)
- **Deskripsi UI**: Antarmuka perbandingan komprehensif antara dua versi dokumen. Menampilkan:
  - **Metadata Snapshot Comparison**: Perbandingan author, klasifikasi (INTERNAL vs CONFIDENTIAL), jumlah halaman, dan ukuran file.
  - **Textual Diff Preview**: Visualisasi perubahan teks (Additions/Deletions) dengan highlight warna dan catatan editor.
  - **Change Approval & Integrity**: Status persetujuan auditor dan persentase kecocokan integritas dokumen.
  - **Action History Log**: Rekam jejak aktivitas perbandingan (Viewed, Compared, Exported).
  - **Difference Summary**: Ringkasan statistik perubahan (Text growth, line additions/deletions, dan risk assessment).
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/documents/:id/compare/:v1/:v2` — get textual and metadata diff
  - `POST /api/v1/documents/:id/compare/export` — export comparison to PDF
- **Handler**: Extend `document_handler.go`
- **Service**: Extend `document_service.go`
- **Database Tables**: `document_versions` ✅, `approval_history` (baru)
- **Relasi Service**: PostgreSQL, MinIO (untuk mengambil konten file versi lama)

---

### 69. Distribution Scenario Switch
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/distribution/scenarios.vue` (New)
- **Deskripsi UI**: Antarmuka konfigurasi metode pengiriman (Dual Scenario). Menampilkan:
  - **Skenario A (Download Mandiri)**: Generator One-Time Link dengan masa berlaku 24 jam, proteksi Multi-Factor, dan batasan retry akses.
  - **Skenario B (Kirim Email)**: Form input email penerima, pemilihan template email, dan ringkasan lampiran dokumen.
  - **Status Validation Bar**: Indikator kesiapan sistem (enkripsi AES-256) dan tombol eksekusi distribusi.
  - **Metode Log & Security Card**: Histori draf skenario dan informasi kepatuhan hukum (audit log/IP tracking).
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/distribution/scenarios` — list available scenarios
  - `POST /api/v1/distribution/generate-link` — generate secure one-time link
  - `POST /api/v1/distribution/send-email` — execute email distribution
- **Handler**: Extend `distribution_handler.go`
- **Service**: Extend `distribution_service.go`
- **Database Tables**: `distribution_scenarios` ✅, `distribution_links` (baru)
- **Relasi Service**: PostgreSQL, SMTP Relay Service, Notification Service
- **Backend API**: ❌ Belum ada.
- **Database Tables**: `distribution_scenarios` (baru)

---

### 70. Distribution Result & Audit Trail
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/distribution/audit.vue` (New)
- **Deskripsi UI**: Dashboard verifikasi hasil distribusi dan audit trail permanen. Menampilkan:
  - **Success Card**: Status pengiriman (Success/Failed), Channel (One-Time Link, Email, dsb), dan Timestamp.
  - **Link Scenario Info**: Detail token keamanan, hitung mundur kadaluarsa link, dan status akses pertama.
  - **Audit Trail Logs**: Tabel kronologis aktivitas (Who, Action, Time, IP/Device, Outcome) dengan format immutable ledger.
  - **Footer Info**: Catatan keamanan enkripsi dan tombol unduh sertifikat/cetak jejak audit.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/distribution/:id/audit` — get distribution audit logs
  - `GET /api/v1/distribution/:id/token-info` — get security token status
  - `GET /api/v1/distribution/:id/certificate` — download PDF certificate
- **Handler**: `distribution_handler.go` (baru)
- **Service**: `distribution_service.go` (baru)
- **Database Tables**: `distribution_audit_logs` (baru), `distribution_tokens` (baru)
- **Relasi Service**: PostgreSQL, Notification Service (untuk pengiriman link)

---

### 71. On Premise Deployment Readiness
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/deployment/readiness.vue` (11892 bytes)
- **Deskripsi UI**: Checklist kesiapan infrastruktur (Server RAM, Storage Space, DB Connectivity, MinIO Health, Redis Latency).
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/system/readiness` — check system health
- **Service**: `system_service.go` (baru)

---

### 72. Backup Scheduler & Restore Wizard
- **Status**: ✅ Frontend ada (built-in in backup.vue), ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/backup.vue` (19310 bytes)
- **Deskripsi UI**: Penjadwalan backup otomatis (daily/weekly/monthly) dan wizard langkah-demi-langkah untuk pemulihan data.
- **Backend API**: ❌ Belum ada — perlu:
  - `POST /api/v1/backup/schedule` — set schedule
  - `POST /api/v1/backup/restore` — run restore wizard
- **Service**: `backup_service.go` (baru)

---

### 73. Security & Encryption Baseline Center
- **Status**: ✅ Frontend ada (built-in in compliance.vue), ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/compliance.vue`
- **Deskripsi UI**: Dashboard status enkripsi data rest dan in-transit. Monitoring login mencurigakan (brute force detection status).

---

### 74. Integration Status Monitor
- **Status**: ✅ Frontend ada, ❌ Backend belum ada
- **Vue File**: `frontend/app/pages/admin/integration/status.vue` (New)
- **Deskripsi UI**: Dashboard monitor kesehatan layanan terintegrasi secara real-time. Menampilkan:
  - **Stats Cards**: Overall Health %, Avg Latency, Active Alerts count, Sync Operations.
  - **Service Table**: List layanan (LDAP, SQL, S3, Network Scanner, Printer, RFID, SMTP) dengan Host Address, Status badge (SELESAI/LATENCY SPIKE/GAGAL), dan Sparkline Latency.
  - **Live Alert Feed**: Feed aktivitas alert (Offline, Latency, Heartbeat Restored).
  - **Latency Trend Chart**: Bar chart tren latency dalam 60 menit terakhir.
- **Backend API**: ❌ Belum ada — perlu:
  - `GET /api/v1/integration/status` — current health and stats
  - `GET /api/v1/integration/services` — list of services and metrics
  - `GET /api/v1/integration/alerts` — live alert feed
  - `POST /api/v1/integration/refresh` — trigger manual health check
- **Handler**: `integration_handler.go` (baru)
- **Service**: `integration_service.go` (baru)
- **Database Tables**: `system_integration_logs` (baru), `system_health_metrics` (baru)
- **Relasi Service**: PostgreSQL, External Services (LDAP, S3, DB, etc.)

---

### 75. Penalty Policy Settings
- **Status**: ✅ Selesai (Frontend ✅, Backend ✅)
- **Vue Files**: 
  - `frontend/app/pages/config/params.vue` — setting interface
  - `frontend/app/pages/loans/my.vue` — display policy
  - `frontend/app/pages/loans/index.vue` — display policy
- **Backend API**: ✅ `GET /api/v1/loans/penalty-policy`
- **Handler**: `loan_handler.go`
- **Service**: `document_service.go` (`GetSystemSetting`)
- **Database Tables**: `system_settings`
- **Deskripsi**: Kebijakan denda keterlambatan pengembalian dokumen yang dikonfigurasi dinamis di parameter sistem dan ditampilkan ke peminjam.

---

### 76. Circulation Pickup Desk
- **Status**: ✅ Selesai (Frontend ✅, Backend ✅)
- **Vue File**: `frontend/app/pages/circulation/pickup.vue`
- **Backend API**:
  - ✅ `GET /api/v1/loans` — fetch pickup queue (filter `l1_approved`)
  - ✅ `GET /api/v1/loans/:id` — fetch loan documents checklist
  - ✅ `POST /api/v1/loans/:id/approve` — mark ready for pickup (status → `l2_approved`)
  - ✅ `POST /api/v1/loans/:id/reject` — return request with reason
- **Handler**: `loan_handler.go` & `approval_handler.go`
- **Service**: `document_service.go`
- **Database Tables**: `borrow_requests`, `loan_request_items`
- **Deskripsi**: Loket persiapan fisik dokumen untuk pengambilan peminjaman. Controller memeriksa dokumen berdasarkan rak/box, menandai status per-dokumen (Ditemukan/Tidak Ditemukan/Rusak), dan menandai siap serah terima dengan auto-redirect ke halaman checkout sirkulasi. Setelah approve, status berubah ke `l2_approved` dan peminjam dapat dijadwalkan pickup di loket checkout.
- **i18n**: ✅ `frontend/app/locales/*/circulation.json` key `circulation.pickup.*`

---

### 77. Document Multi-Attachment (Document Files)
- **Status**: ✅ Selesai (Frontend ✅, Backend ✅)
- **Vue File**: `frontend/app/pages/documents/upload.vue`
- **Backend API**: ✅ `POST /api/v1/documents/` — mendukung field `files[]` untuk multiple file upload
- **Handler**: `document_handler.go` — menggunakan `MultipartForm()`, iterasi `files[]` + fallback ke `file`
- **Service**: `document_service.go` — `ExtraFileParam` struct, upload extra files ke MinIO, simpan di `document_files`
- **Repository**: `backend/internal/repository/document_files.go` — `CreateDocumentFile`, `ListDocumentFiles`
- **Database Tables**: `document_files` ✅ (migration 000080) — `document_id`, `file_name`, `file_path`, `file_size`, `mime_type`, `sort_order`
- **Deskripsi**: File pertama menjadi dokumen utama; file ke-2 dst disimpan sebagai lampiran di tabel `document_files` yang terhubung via `document_id`. Response upload menyertakan `document_files` array.

---

### Ringkasan Status Master Data

| # | Fitur | Frontend | Backend API | Database |
|---|-------|----------|-------------|----------|
| 1 | SSO | ✅ | ❌ | ✅ |
| 2 | Role Management | ✅ | ⚠️ Parsial | ✅ |
| 3 | PIN Pengguna | ✅ | ❌ | ✅ |
| 4 | Departemen Mgmt | ✅ | ❌ | ✅ |
| 5 | User/Doc Detail | ✅ | ⚠️ Parsial | ✅ |
| 6 | Split Screen Input | ✅ | ❌ | ✅ |
| 7 | OCR & AI Review | ✅ | ⚠️ Worker ada | ✅ |
| 8 | Duplicate Detection | ✅ | ❌ | ✅ |
| 9 | Identity & Foldering | ✅ | ❌ | ✅ |
| 10 | Bulk Import | ✅ | ❌ | ✅ |
| 11 | Digital Mailroom | ✅ | ❌ | ✅ |
| 12 | Override Rak | ✅ | ❌ | ✅ |
| 13 | Zoning Control | ✅ | ❌ | ✅ |
| 14 | RFID/QR Label | ✅ | ✅ | ✅ |
| 15 | Label Printing | ✅ | ✅ | ✅ |
| 16 | Capacity Monitor | ✅ | ✅ | ✅ |
| 17 | Warehouse Topology | ✅ | ❌ | ✅ |
| 18 | RFID Model F-16 | ✅ | ❌ | ✅ |
| 19 | Location Detail & Slot | ✅ | ❌ | ✅ |
| 20 | Input Entry Hub | ✅ | ❌ | ✅ |
| 21 | OCR Process Monitor | ✅ | ⚠️ Worker ada | ✅ |
| 22 | Folder Path Review | ✅ | ❌ | ✅ |
| 23 | Auto Numbering | ✅ | ❌ | ✅ |
| 24 | Final Review & Commit | ✅ | ❌ | ✅ |
| 25 | Dynamic Filter | ✅ | ❌ | ✅ |
| 26 | Search Results | ✅ | ⚠️ Parsial | ✅ |
| 27 | Universal Search | ✅ | ❌ | ✅ |
| 28 | Drill Down Workspace | ✅ | ❌ | ✅ |
| 29 | Document Detail | ✅ | ⚠️ Parsial | ✅ |
| 30 | Relationship Mgmt | ✅ | ❌ | ✅ |
| 31 | Search Analytics | ❌ | ❌ | ❌ |
| 32 | Search Context Commit | ✅ | ❌ | ✅ |
| 34 | Audit Mission Planner | ✅ | ❌ | ✅ |
| 35 | Blind Audit Interface | ✅ | ❌ | ✅ |
| 36 | Discrepancy Resolution | ✅ | ❌ | ✅ |
| 37 | Security Trimming | ✅ | ❌ | ✅ |
| 38 | Encryption Compliance | ✅ | ❌ | ❌ |
| 39 | Access Policy Matrix | ✅ | ⚠️ Parsial | ✅ |
| 40 | RFID Gate Monitor | ✅ | ❌ | ❌ |
| 41 | History Log Explorer | ❌ | ❌ | ✅ |
| 42 | Version & Checkout | ✅ (parsial) | ❌ | ✅ |
| 43 | Redaction & Annotation | ❌ | ❌ | ❌ |
| 44 | Secure Preview | ✅ | ❌ | ✅ |
| 45 | Notification Orchestrator | ✅ | ❌ | ❌ |
| 46 | Extension Request | ✅ | ❌ | ❌ |
| 47 | Loan Timer | ✅ | ❌ | ✅ |
| 48 | Pre Registration | ✅ | ❌ | ❌ |
| 49 | Distribution Center | ✅ | ❌ | ❌ |
| 50 | Physical Loan Desk (Checkout) | ✅ | ⚠️ Parsial | ✅ |
| 51 | Two Level Approval | ✅ | ❌ | ❌ |
| 52 | Fast Track Selector | ✅ | ❌ | ❌ |
| 53 | Request Cart | ✅ | ❌ | ❌ |
| 54 | System Config (SSO) | ✅ | ❌ | ❌ |
| 55 | Retention Disposal | ✅ | ❌ | ✅ |
| 56 | Backup Recovery | ✅ | ❌ | ❌ |
| 57 | User Mgmt Panel | ✅ | ⚠️ Parsial | ✅ |
| 58 | Warehouse Occupancy | ✅ | ❌ | ✅ |
| 59 | Audit Trail | ✅ | ❌ | ❌ |
| 60 | Circulation Report | ❌ | ❌ | ✅ |
| 61 | User Scorecard | ✅ | ❌ | ✅ |
| 62 | Dept Performance | ❌ | ❌ | ✅ |
| 63 | Executive KPI | ✅ | ❌ | ✅ |
| 64 | No Code Console | ✅ | ❌ | ❌ |
| 65 | Scan Launcher | ✅ | ❌ | ❌ |
| 66 | Live Scan | ✅ | ❌ | ❌ |
| 67 | Doc Versioning | ✅ | ⚠️ Parsial | ✅ |
| 68 | Version Compare | ✅ | ❌ | ❌ |
| 69 | Distribution Scenario | ✅ | ❌ | ❌ |
| 70 | Distribution Result & Audit Trail | ✅ | ❌ | ❌ |
| 71 | Deployment Readiness | ✅ | ❌ | ✅ |
| 72 | Backup Scheduler | ✅ | ❌ | ❌ |
| 73 | Security Baseline | ✅ | ❌ | ✅ |
| 74 | Integration Status | ✅ | ❌ | ✅ |
| 75 | Penalty Policy Settings | ✅ | ✅ | ✅ |
| 76 | Circulation Pickup Desk | ✅ | ✅ | ✅ |
| 77 | Document Multi-Attachment | ✅ | ✅ | ✅ |

> **Kesimpulan**: Dari 77 fitur yang dipetakan, **57 fitur sudah memiliki frontend**, **20 fitur belum ada frontend**. Progress backend meningkat dengan selesainya modul **Warehouse QR Labeling & Smart Storage**, **Penalty Policy**, loket **Circulation Pickup Desk**, **Circulation Checkout** (overhaul), dan **Document Multi-Attachment** (migration 000080). Status loan flow diperbarui: `l1_approved` = "Dalam Persiapan", `l2_approved` = "Siap Diambil", `active` = "Active". Locales i18n modul circulation (`en` + `id`) sudah tersedia lengkap. Backend handler yang sudah aktif: `auth_handler.go`, `user_handler.go`, `document_handler.go`, `master_handler.go`, `dashboard_handler.go`, `intake_handler.go`, `loan_handler.go` (extended), `approval_handler.go` (extended). Backend handler yang masih perlu dibuat/diperluas: `notification_handler.go`, `registration_handler.go`, `circulation_handler.go` (incoming), `cart_handler.go`, `retention_handler.go`, `audit_handler.go`.

