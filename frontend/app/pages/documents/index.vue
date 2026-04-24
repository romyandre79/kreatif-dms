<template>
  <div class="space-y-8">
    <div v-motion-fade>
      <h1 class="text-3xl font-bold text-slate-900 dark:text-white">Documents</h1>
      <p class="text-slate-500 dark:text-slate-400 mt-1">Manage and organize your digital assets.</p>
    </div>
    
    <div class="glass rounded-3xl p-12 flex flex-col items-center justify-center text-center space-y-4" v-motion-slide-visible-bottom>
      <div class="w-20 h-20 rounded-2xl bg-primary-50 dark:bg-primary-900/20 flex items-center justify-center text-primary-500">
        <LucideFiles class="w-10 h-10" />
      </div>
      <h2 class="text-xl font-bold">Document Explorer Coming Soon</h2>
      <p class="text-slate-500 max-w-md">We're building a powerful document management system for you. Stay tuned!</p>
    </div>
  </div>
</template>

<script setup>
import { LucideFiles } from 'lucide-vue-next'
</script>
