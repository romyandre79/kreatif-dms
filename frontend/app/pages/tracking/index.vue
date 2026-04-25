<template>
  <div class="space-y-6" v-motion-fade>
    <div>
      <h1 class="text-3xl font-bold text-slate-900 dark:text-white">{{ $t('layout.menu.tracking') }}</h1>
      <p class="text-slate-500 mt-1">{{ $t('placeholders.coming_soon_desc', { page: $t('layout.menu.tracking') }) }}</p>
    </div>
    
    <div class="glass rounded-3xl p-12 flex flex-col items-center justify-center text-center border-2 border-dashed border-slate-200 dark:border-slate-800">
      <div class="w-20 h-20 rounded-2xl bg-primary-50 dark:bg-primary-900/20 flex items-center justify-center text-primary-500 mb-6">
        <LucideBarChart3 class="w-10 h-10" />
      </div>
      <h2 class="text-xl font-black text-slate-800 dark:text-white">{{ $t('placeholders.coming_soon_title', { page: $t('layout.menu.tracking') }) }}</h2>
      <p class="text-slate-500 max-w-sm mt-2 leading-relaxed">{{ $t('placeholders.coming_soon_desc', { page: $t('layout.menu.tracking') }) }}</p>
    </div>
  </div>
</template>

<script setup>
import { LucideBarChart3 } from 'lucide-vue-next'
</script>
