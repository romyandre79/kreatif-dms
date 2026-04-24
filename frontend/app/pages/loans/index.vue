<template>
  <div class="space-y-6" v-motion-fade>
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-3xl font-bold text-slate-900 dark:text-white">Document Loans</h1>
        <p class="text-slate-500 mt-1">Manage and track your document requests.</p>
      </div>
    </div>
    
    <div class="glass rounded-3xl p-12 flex flex-col items-center justify-center text-center border-2 border-dashed border-slate-200 dark:border-slate-800">
      <div class="w-20 h-20 rounded-2xl bg-blue-50 dark:bg-blue-900/20 flex items-center justify-center text-blue-500 mb-6">
        <LucideBookOpen class="w-10 h-10" />
      </div>
      <h2 class="text-xl font-black text-slate-800 dark:text-white">Loans Module Under Construction</h2>
      <p class="text-slate-500 max-w-sm mt-2 leading-relaxed">We are currently building the document loan and tracking system. Check back soon!</p>
    </div>
  </div>
</template>

<script setup>
import { LucideBookOpen } from 'lucide-vue-next'
</script>
