<template>
  <div class="space-y-8">
    <div v-motion-fade>
      <h1 class="text-3xl font-bold text-slate-900 dark:text-white">{{ $t('layout.navbar.account_settings') }}</h1>
      <p class="text-slate-500 dark:text-slate-400 mt-1">{{ $t('placeholders.coming_soon_desc', { page: $t('layout.navbar.account_settings') }) }}</p>
    </div>
    
    <div class="glass rounded-3xl p-12 flex flex-col items-center justify-center text-center space-y-4" v-motion-slide-visible-bottom>
      <div class="w-20 h-20 rounded-2xl bg-slate-50 dark:bg-slate-900/20 flex items-center justify-center text-slate-500">
        <LucideSettings class="w-10 h-10" />
      </div>
      <h2 class="text-xl font-bold">{{ $t('placeholders.coming_soon_title', { page: $t('layout.navbar.account_settings') }) }}</h2>
      <p class="text-slate-500 max-w-md">{{ $t('placeholders.coming_soon_desc', { page: $t('layout.navbar.account_settings') }) }}</p>
    </div>
  </div>
</template>

<script setup>
import { LucideSettings } from 'lucide-vue-next'
</script>
