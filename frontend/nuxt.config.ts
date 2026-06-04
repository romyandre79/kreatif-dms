// https://nuxt.com/docs/api/configuration/nuxt-config
// Triggering route refresh
import tailwindcss from "@tailwindcss/vite";

export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  future: {
  },
  ssr: false, // Disable SSR to avoid vite-node IPC issues on Windows
  devtools: { enabled: false }, // Disable devtools to save RAM
  sourcemap: { server: false, client: false }, // Disable sourcemaps to save RAM
  app: {
    head: {
      link: [
        { rel: 'preconnect', href: 'https://fonts.googleapis.com' },
        { rel: 'preconnect', href: 'https://fonts.gstatic.com', crossorigin: '' },
        { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@400;500;600;700;800&display=swap' }
      ]
    }
  },
  modules: [
    '@pinia/nuxt',
    '@vueuse/nuxt',
    '@vueuse/motion/nuxt',
    '@nuxt/eslint',
    'nuxt-security',
    '@nuxtjs/i18n'
  ],
  i18n: {
    locales: [
      { 
        code: 'en', 
        name: 'English', 
        files: [
          'en/common.json',
          'en/layout.json',
          'en/landing.json',
          'en/login.json',
          'en/dashboard.json',
          'en/documents.json',
          'en/upload.json',
          'en/manifest.json',
          'en/warehouse.json',
          'en/loans.json',
          'en/approvals.json',
          'en/admin.json',
          'en/notifications.json',
          'en/intake.json',
          'en/circulation.json',
          'en/stock.json'
        ] 
      },
      { 
        code: 'id', 
        name: 'Bahasa Indonesia', 
        files: [
          'id/common.json',
          'id/layout.json',
          'id/landing.json',
          'id/login.json',
          'id/dashboard.json',
          'id/documents.json',
          'id/upload.json',
          'id/manifest.json',
          'id/warehouse.json',
          'id/loans.json',
          'id/approvals.json',
          'id/admin.json',
          'id/notifications.json',
          'id/intake.json',
          'id/circulation.json',
          'id/stock.json'
        ] 
      }
    ],
    defaultLocale: 'id',
    lazy: false,
    langDir: 'locales',
    restructureDir: 'app',
    strategy: 'no_prefix'
  },
  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || 'http://localhost:8080/api/v1',
      ocrUrl: process.env.NUXT_PUBLIC_OCR_URL || 'http://localhost:8000',
      appName: process.env.NUXT_PUBLIC_APP_NAME || 'Kreatif DMS',
      appVersion: process.env.NUXT_PUBLIC_APP_VERSION || 'V3.1.2',
      copyright: process.env.NUXT_PUBLIC_COPYRIGHT || 'PT Prisma Data Abadi',
    }
  },
  security: {
    headers: {
      contentSecurityPolicy: {
        'img-src': ["'self'", "data:", "blob:", "https:", "http:"],
        'script-src': ["'self'", "'unsafe-inline'", "https:", "https://vercel.live"],
        'frame-src': ["'self'", "https:", "http:"],
      },
    },
  },
  devServer: {
    host: 'localhost',
    port: process.env.APP_PORT ? parseInt(process.env.APP_PORT) : 3000
  },
  vite: {
    plugins: [
      tailwindcss(),
    ],
    optimizeDeps: {
      include: [
        'lucide-vue-next',
      ]
    },
    server: {
      watch: {
        usePolling: true,
        interval: 1000, // Faster polling than 2500 for better responsiveness
        binaryInterval: 1000
      },
      hmr: {
        overlay: false, // Disable HMR overlay to save RAM
        protocol: 'ws' // Explicitly use WebSockets to avoid IPC channel issues
      }
    },
    build: {
      chunkSizeWarningLimit: 1000,
      commonjsOptions: {
        transformMixedEsModules: true
      }
    }
  },
  css: ['./app/assets/css/main.css'],
})
