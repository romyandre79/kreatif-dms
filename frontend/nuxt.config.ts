import tailwindcss from "@tailwindcss/vite";

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  future: {
    compatibilityVersion: 4,
  },
  devtools: { enabled: true },
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
    '@nuxtjs/i18n',
    '@vite-pwa/nuxt'
  ],
  pwa: {
    registerType: 'autoUpdate',
    manifest: {
      name: 'Kreatif DMS',
      short_name: 'KreatifDMS',
      theme_color: '#1E3A5F',
      icons: [
        {
          src: 'icons/icon.png',
          sizes: '192x192',
          type: 'image/png'
        },
        {
          src: 'icons/icon.png',
          sizes: '512x512',
          type: 'image/png'
        },
        {
          src: 'icons/icon.png',
          sizes: '512x512',
          type: 'image/png',
          purpose: 'any maskable'
        }
      ]
    },
    workbox: {
      navigateFallback: '/',
      globPatterns: ['**/*.{js,css,html,png,svg,ico}']
    },
    client: {
      installPrompt: true,
      periodicSyncForUpdates: 3600,
    },
    devOptions: {
      enabled: true,
      type: 'classic',
    }
  },
  i18n: {
    locales: [
      { code: 'en', name: 'English', file: 'en.json' },
      { code: 'id', name: 'Bahasa Indonesia', file: 'id.json' }
    ],
    defaultLocale: 'id',
    lazy: false,
    langDir: './locales',
    strategy: 'no_prefix'
  },
  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || 'http://localhost:8080/api/v1',
      appName: process.env.NUXT_PUBLIC_APP_NAME || 'Kreatif DMS',
      copyright: process.env.NUXT_PUBLIC_COPYRIGHT || 'PT Prisma Data Abadi',
    }
  },
  security: {
    headers: {
      contentSecurityPolicy: {
        'img-src': ["'self'", "data:", "https:", "http:"],
      },
    },
  },
  vite: {
    plugins: [
      tailwindcss(),
    ],
  },
  css: ['./app/assets/css/main.css'],
})
