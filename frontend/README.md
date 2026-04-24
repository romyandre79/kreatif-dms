# Kreatif DMS - Frontend

Modern frontend for Kreatif Document Management System built with Nuxt 4 and Tailwind CSS v4.

## 🚀 Features
- **Framework**: Nuxt 4 (Nuxt v4 app structure).
- **Styling**: Tailwind CSS v4 (CSS-first configuration).
- **State Management**: Pinia.
- **Security**: Nuxt Security (CSP, SRI, Helmet-like headers).
- **Quality**: ESLint + Prettier (@nuxt/eslint).
- **Environments**: Support for `.env`, `.env.staging`, and `.env.production`.
- **Observability**: Custom color-coded logger for easier debugging (`app/utils/logger.ts`).

## 🛠️ Requirements
- Node.js 18+
- npm / pnpm / yarn

## 🚦 How to Run

### 1. Install Dependencies
```bash
npm install
```

### 2. Development
```bash
npm run dev
```
The app will be available at `http://localhost:3000`.

### 3. Build for Production
```bash
# Production
npm run build

# Staging (if configured in scripts)
# nuxi build --dotenv .env.staging
```

## 📂 Structure
- `app/`: Source code (Components, Pages, Assets).
- `assets/`: Global CSS and fonts.
- `public/`: Static assets.
