# Kreatif DMS — Frontend Design Style Guide

Dokumen ini mencerminkan implementasi desain aktual pada kodebase. Gunakan sebagai referensi saat membuat atau memodifikasi halaman Vue.

---

## 1. Design Concept

- **Premium Enterprise UI** — dark navy brand, clean whitespace, high contrast
- **Compact & Dense** — informasi padat tanpa terasa sesak
- **Glassmorphism Subtle** — backdrop blur pada header, overlay, modal
- **Motion Aware** — menggunakan `@vueuse/motion` untuk animasi masuk
- **Dark Mode Ready** — semua komponen support `dark:` variant Tailwind
- **Typography Dominan** — uppercase tracking-widest sebagai signature style

---

## 2. Color Palette

### Brand Colors (Hardcoded — pakai langsung di kelas Tailwind)

```
Navy Primary:   #1E3A5F  → bg-[#1E3A5F]  (button utama, heading, icon aktif)
Navy Hover:     #152943  → hover:bg-[#152943]
Navy Overlay:   #1E3A5F/40 → bg-[#1E3A5F]/40  (modal backdrop)

Dark Surface:   #0A0F1C  → dark bg body
Dark Card:      #0D121F  → dark card / sidebar
Dark Card Alt:  #1A2234  → dark icon container
```

### Tailwind Theme Colors (`--color-primary-*` = sky/cyan — dipakai untuk aksen)

```css
/* main.css @theme block */
--color-primary-500: #0ea5e9;   /* sky-500 — active menu, badge, ring, link */
--color-primary-600: #0284c7;   /* sky-600 */
```

> **Aturan**: `#1E3A5F` untuk elemen brand utama (tombol, heading). `primary-500` (#0ea5e9) untuk aksen interaktif (hover state, badge notifikasi, active indicator, focus ring).

### Neutral / Surface Colors

```
Background:     #F8FAFC  → bg-slate-50 / bg-[#F8FAFC]
Surface:        #FFFFFF  → bg-white
Border:         slate-200/60  → border-slate-200/60
Text Primary:   #1E3A5F  → text-[#1E3A5F]  (heading, label penting)
Text Body:      #0f172a  → text-slate-900
Text Secondary: #64748b  → text-slate-500
Text Muted:     #94a3b8  → text-slate-400  (placeholder, metadata)
```

### Status Colors

```
Success:  #16a34a  → text/bg green-600
Warning:  #f59e0b  → text/bg amber-400
Danger:   #dc2626  → text/bg red-600
Info:     #0ea5e9  → primary-500
```

---

## 3. Typography

### Font

```css
/* nuxt.config.ts — loaded dari Google Fonts */
font-family: "Plus Jakarta Sans", "Inter", ui-sans-serif, system-ui, ...

/* Font weights yang dipakai */
400 — body text
500 — medium label
600 — semibold
700 — bold
800 — extrabold (page title)
```

### Skala Teks (Tailwind classes)

```
Page Title:    text-4xl font-extrabold tracking-tight text-[#1E3A5F]   (32-36px)
Section Title: text-xl font-black uppercase tracking-tight              (20px)
Card Label:    text-[11px] font-black uppercase tracking-widest         (11px)
Body:          text-sm font-medium text-slate-600                       (14px)
Meta/Caption:  text-[10px] font-bold uppercase tracking-widest text-slate-400
Micro:         text-[9px] font-bold uppercase tracking-widest           (9px)
```

> **Signature style**: uppercase + tracking-widest + font-black pada label dan heading adalah ciri khas UI ini. Konsisten di seluruh halaman.

---

## 4. Layout Structure

```
┌─────────────────────────────────────────────────┐
│  Sidebar (280px, fixed, white)                  │
│  ┌────┐  ┌──────────────────────────────────┐   │
│  │    │  │  Header (80px, sticky, blur)     │   │
│  │    │  ├──────────────────────────────────┤   │
│  │    │  │  Main Content (p-2 wrapper)      │   │
│  │    │  │  └─ Page content (p-6 to p-8)   │   │
│  └────┘  └──────────────────────────────────┘   │
└─────────────────────────────────────────────────┘
```

```css
/* Layout grid */
.app-layout {
  display: flex;
  min-height: 100vh;
  background: #F8FAFC;
}

/* Main wrapper offset */
.main-wrapper {
  margin-left: 280px;   /* lg:ml-[280px] */
  flex: 1;
  display: flex;
  flex-direction: column;
}
```

---

## 5. Sidebar

```css
/* Aktual dari default.vue */
.sidebar {
  width: 280px;              /* w-[280px] */
  background: white;         /* bg-white dark:bg-[#0D121F] */
  border-right: 1px solid rgba(226, 232, 240, 0.6);
  position: fixed;
  height: 100%;
  z-index: 110;
  transition: all 500ms cubic-bezier(0.4, 0, 0.2, 1);
}

/* Logo area */
.sidebar-logo { padding: 32px 32px 16px; }

/* Menu item base */
.sidebar-menu {
  padding: 14px 20px;
  border-radius: 16px;       /* rounded-2xl */
  transition: all 300ms;
  font-size: 11px;
  font-weight: 900;
  text-transform: uppercase;
  letter-spacing: tight;
}

/* Active menu */
.sidebar-menu.active {
  background: #0ea5e9;       /* bg-primary-500 */
  color: white;
  box-shadow: 0 20px 25px rgba(14, 165, 233, 0.2);
}

/* Icon container */
.menu-icon {
  width: 32px; height: 32px;
  border-radius: 12px;       /* rounded-xl */
}

/* Submenu child */
.submenu-child {
  padding: 12px 36px;
  border-radius: 12px;
  font-size: 10px;
  font-weight: 900;
}

/* Active child indicator */
.submenu-child.active::before {
  content: '';
  position: absolute; left: 0; top: 0; bottom: 0;
  width: 4px;
  background: #0ea5e9;
  border-radius: 9999px;
}
```

---

## 6. Header / Navbar

```css
.header {
  height: 80px;              /* h-20 */
  background: rgba(255, 255, 255, 0.7);   /* bg-white/70 */
  backdrop-filter: blur(24px);            /* backdrop-blur-xl */
  border-bottom: 1px solid rgba(226, 232, 240, 0.6);
  padding: 0 32px;           /* px-8 */
  position: sticky;
  top: 0;
  z-index: 90;
}

/* Search input */
.header-search {
  background: transparent;
  border: none;
  width: 256px;              /* w-64 */
  font-size: 14px;
  font-weight: 500;
  color: #475569;
}

/* Icon button */
.header-icon-btn {
  padding: 10px;
  border-radius: 12px;
  background: #f8fafc;       /* bg-slate-50 */
}
```

---

## 7. Card Component

```css
.card {
  background: white;
  border-radius: 24px;       /* rounded-3xl */
  padding: 24px;             /* p-6 — standar */
  border: 1px solid rgba(226, 232, 240, 0.6);
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.04);
}

/* Variant: larger padding untuk halaman detail */
.card-lg { padding: 32px; } /* p-8 */

/* Variant: glassmorphism */
.card-glass {
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 22px;
}

/* Hover lift */
.card:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 24px rgba(0, 0, 0, 0.08);
  transition: all 0.3s ease;
}
```

> **Aturan border-radius**: `rounded-3xl` (24px) untuk card utama, `rounded-2xl` (16px) untuk card kecil/nested, `rounded-xl` (12px) untuk inner element.

---

## 8. Button Design

```css
/* Primary — brand navy */
.btn-primary {
  background: #1E3A5F;
  color: white;
  padding: 14px 24px;        /* py-3.5 px-6 */
  border-radius: 16px;       /* rounded-2xl */
  font-size: 10-11px;
  font-weight: 900;
  text-transform: uppercase;
  letter-spacing: 0.2em;
  transition: all 0.3s;
  box-shadow: 0 20px 25px rgba(30, 58, 95, 0.2);
}
.btn-primary:hover { background: #152943; }

/* Secondary */
.btn-secondary {
  background: white;
  border: 1px solid #e2e8f0;
  color: #64748b;
  border-radius: 16px;
  padding: 14px 24px;
  font-size: 10px;
  font-weight: 900;
  text-transform: uppercase;
  letter-spacing: 0.2em;
}

/* Danger */
.btn-danger { background: #dc2626; color: white; border-radius: 16px; }

/* Disabled state */
.btn:disabled {
  background: #f1f5f9;
  color: #94a3b8;
  cursor: not-allowed;
  box-shadow: none;
}

/* Loading state */
.btn-loading { opacity: 0.8; cursor: wait; }
```

---

## 9. Form / Input Design

```css
.input {
  width: 100%;
  border: 1px solid #e2e8f0;
  border-radius: 12px;       /* rounded-xl */
  padding: 12px 16px;
  font-size: 13px;
  font-weight: 600;
  color: #1e293b;
  background: white;
  outline: none;
  transition: all 0.2s;
}

.input:focus {
  border-color: #0ea5e9;     /* primary-500 */
  box-shadow: 0 0 0 4px rgba(14, 165, 233, 0.1);  /* ring-primary-500/10 */
}

.label {
  font-size: 10px;
  font-weight: 900;
  text-transform: uppercase;
  letter-spacing: 0.1em;
  color: #94a3b8;            /* text-slate-400 */
  margin-bottom: 8px;
  display: block;
}

/* Select */
.select {
  border-radius: 12px;
  padding: 12px 16px;
  font-size: 13px;
  font-weight: 600;
}

/* Textarea */
.textarea {
  border-radius: 16px;
  padding: 16px;
  font-size: 13px;
  resize: none;
}
```

---

## 10. Table Design

```css
.table-container {
  background: white;
  border-radius: 24px;       /* rounded-3xl */
  border: 1px solid rgba(226, 232, 240, 0.6);
  overflow: hidden;
}

.table {
  width: 100%;
  border-collapse: collapse;
}

.table th {
  background: #f8fafc;       /* bg-slate-50 */
  font-size: 9px;
  font-weight: 900;
  text-transform: uppercase;
  letter-spacing: 0.15em;
  color: #94a3b8;
  padding: 14px 16px;        /* px-4 py-3.5 */
  text-align: left;
}

.table td {
  padding: 14px 16px;
  border-bottom: 1px solid #f8fafc;
  font-size: 13px;
  color: #374151;
}

.table tr:hover td { background: #fafafa; }

.table tr:last-child td { border-bottom: none; }
```

---

## 11. Badge / Status Chip

```css
.badge {
  padding: 4px 12px;
  border-radius: 9999px;     /* rounded-full */
  font-size: 9px;
  font-weight: 900;
  text-transform: uppercase;
  letter-spacing: 0.15em;
  display: inline-flex;
  align-items: center;
  gap: 4px;
}

/* Variants */
.badge-active    { background: #dcfce7; color: #166534; }   /* green */
.badge-pending   { background: #fef3c7; color: #92400e; }   /* amber */
.badge-danger    { background: #fee2e2; color: #991b1b; }   /* red */
.badge-info      { background: #e0f2fe; color: #075985; }   /* sky */
.badge-navy      { background: #1E3A5F; color: white; }     /* brand */
.badge-default   { background: #f1f5f9; color: #64748b; }   /* slate */
```

---

## 12. Stat Card (Dashboard Widget)

```css
.stat-card {
  background: white;
  border-radius: 24px;
  padding: 24px;
  border: 1px solid rgba(226, 232, 240, 0.6);
  min-height: 110px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  transition: all 0.3s;
}

.stat-title {
  font-size: 9px;
  font-weight: 900;
  text-transform: uppercase;
  letter-spacing: 0.2em;
  color: #94a3b8;
}

.stat-value {
  font-size: 36px;
  font-weight: 800;
  color: #1E3A5F;
  line-height: 1;
}

.stat-change {
  font-size: 10px;
  font-weight: 700;
  color: #16a34a;
}

/* Colored variant */
.stat-card-navy {
  background: linear-gradient(135deg, #1E3A5F, #152943);
  color: white;
  border: none;
}
```

---

## 13. Upload Area

```css
.upload-area {
  border: 2px dashed #cbd5e1;  /* border-slate-300 */
  border-radius: 24px;
  padding: 48px 40px;
  text-align: center;
  background: #f8fafc;
  transition: all 0.3s;
  cursor: pointer;
}

.upload-area:hover,
.upload-area.drag-over {
  border-color: #0ea5e9;      /* primary-500 */
  background: #f0f9ff;        /* sky-50 */
}
```

---

## 14. Modal

```css
.modal-overlay {
  background: rgba(30, 58, 95, 0.4);  /* bg-[#1E3A5F]/40 */
  backdrop-filter: blur(4px);
}

.modal {
  background: white;
  border-radius: 24px;        /* rounded-3xl */
  padding: 32px;
  width: 520px;
  max-width: 90vw;
  box-shadow: 0 25px 50px rgba(0, 0, 0, 0.2);
  border: 1px solid rgba(226, 232, 240, 0.5);
}

.modal-title {
  font-size: 11px;
  font-weight: 900;
  text-transform: uppercase;
  letter-spacing: 0.2em;
  color: #1E3A5F;
}
```

---

## 15. Animation

### Library: `@vueuse/motion`

```html
<!-- Fade in dari bawah -->
<div v-motion-fade>...</div>

<!-- Slide dari kiri -->
<div v-motion-slide-left>...</div>

<!-- Slide dari kanan -->
<div v-motion-slide-right>...</div>
```

### CSS Transitions

```css
/* Standard transition */
transition: all 0.3s ease;

/* Layout transition */
transition: all 500ms cubic-bezier(0.4, 0, 0.2, 1);

/* Fade */
.fade-enter-active, .fade-leave-active { transition: opacity 0.3s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; }

/* Expand (accordion/submenu) */
.expand-enter-active, .expand-leave-active {
  transition: all 0.3s ease-out;
  max-height: 500px;
  overflow: hidden;
}
.expand-enter-from, .expand-leave-to {
  max-height: 0;
  opacity: 0;
}
```

---

## 16. Spacing Convention

```
Page wrapper:   p-2     (8px) — main content slot di layout
Section outer:  p-6     (24px) — outer card/section
Section inner:  p-8     (32px) — detail card dengan lebih banyak konten
Grid gap:       gap-6   (24px) — standar grid
Grid gap lg:    gap-8   (32px) — untuk grid jarang

Section margin: space-y-6 (24px) — jarak antar section dalam halaman
```

> Hindari `p-10` (40px) atau `gap-12` (48px) kecuali untuk hero section / halaman spesial.

### Compact & Mobile-Friendly Standards
Untuk mengoptimalkan ruang vertikal agar tidak terlalu banyak ruang kosong (whitespace) namun tetap ramah perangkat seluler (mobile friendly):
- **Paddings**: Gunakan padding yang lebih kecil (`p-4` atau `p-5`) di mobile dan batasi maksimal `p-6` di desktop. Ini menghindari pemborosan ruang pada layar kecil.
- **Gaps & Grid**: Gunakan `gap-4` atau `gap-5` untuk jarak grid. Hindari `gap-8` atau lebih besar di area kerja utama. Gunakan breakpoint `lg:` untuk memisahkan layout kolom (misal: `grid-cols-1 lg:grid-cols-3`) agar otomatis mengalir menjadi satu kolom di mobile.
- **Vertical Spacing**: Gunakan `space-y-4` atau `space-y-5` untuk menumpuk elemen. Ini merapatkan susunan komponen tanpa membuatnya terasa berhimpitan.
- **Textarea & Input**: Batasi baris input seperti textarea ke `rows="3"` dengan padding `p-4` untuk menghemat tinggi lipatan layar (viewport height).
- **Responsive Font & Text**: Gunakan kombinasi text responsif (misal: `text-2xl lg:text-3xl` untuk judul utama) agar font mengecil secara proporsional di layar handphone.

---

## 17. Responsive Design

```css
/* Sidebar collapse — mobile */
@media (max-width: 1024px) {
  .sidebar { transform: translateX(-100%); }
  .sidebar.open { transform: translateX(0); box-shadow: 0 25px 50px rgba(0,0,0,0.3); }
  .main-wrapper { margin-left: 0; }
}

/* Grid responsive */
.grid-auto { grid-template-columns: repeat(auto-fit, minmax(220px, 1fr)); }

/* Table overflow */
.table-container { overflow-x: auto; }
```

---

## 18. Dark Mode

```css
/* Body */
body          → bg-slate-50      dark:bg-[#0A0F1C]
Sidebar       → bg-white         dark:bg-[#0D121F]
Header        → bg-white/70      dark:bg-[#0A0F1C]/70
Card          → bg-white         dark:bg-[#0D121F]
Icon bg       → bg-slate-100     dark:bg-[#1A2234]
Border        → border-slate-200 dark:border-slate-800
Text primary  → text-[#1E3A5F]  dark:text-white
Text body     → text-slate-700   dark:text-slate-300
Text muted    → text-slate-400   dark:text-slate-500
```

---

## 19. Utility Classes

```css
/* Defined in main.css */
.glass {
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  /* dark: bg-slate-900/70 border-slate-800/20 */
}

/* Custom scrollbar */
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: rgba(148, 163, 184, 0.2);
  border-radius: 10px;
}
```

---

## 20. Component Checklist

Saat membuat halaman baru, pastikan:

- [ ] Font class `font-sans` sudah ada di root element (atau inherit dari layout)
- [ ] Warna heading gunakan `text-[#1E3A5F]` (bukan `text-blue-900` dll)
- [ ] Primary button gunakan `bg-[#1E3A5F] hover:bg-[#152943]`
- [ ] Card pakai `rounded-3xl` + `border border-slate-200/60`
- [ ] Label input pakai `text-[10px] font-black uppercase tracking-widest text-slate-400`
- [ ] Badge pakai `rounded-full text-[9px] font-black uppercase tracking-widest`
- [ ] Tambahkan `dark:` variant untuk semua warna background & text
- [ ] Gunakan `v-motion-fade` atau `v-motion-slide-left/right` untuk animasi masuk
- [ ] Semua section heading gunakan `text-[11px] font-black uppercase tracking-widest`

---

## 21. Ikon

Semua ikon menggunakan **lucide-vue-next**:

```vue
<script setup>
import { LucideFileText, LucideSearch, LucideCheckCircle2 } from 'lucide-vue-next'
</script>

<template>
  <!-- Ukuran standar -->
  <LucideFileText class="w-4 h-4" />      <!-- 16px — inline/menu -->
  <LucideSearch class="w-5 h-5" />        <!-- 20px — header/card -->
  <LucideCheckCircle2 class="w-6 h-6" />  <!-- 24px — feature icon -->
</template>
```

---

## 22. Stack Frontend Aktual

| Layer | Teknologi |
|-------|-----------|
| Framework | Nuxt 4 (SSR disabled) |
| Styling | Tailwind CSS v4 + PostCSS |
| Font | Plus Jakarta Sans (Google Fonts) |
| Icons | lucide-vue-next |
| Animation | @vueuse/motion |
| State | Pinia |
| i18n | @nuxtjs/i18n (id default, en) |
| Security | nuxt-security (CSP headers) |
| HTTP | Axios via `useApi` composable |
