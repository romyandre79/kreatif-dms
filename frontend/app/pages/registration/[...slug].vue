<template>
  <div class="space-y-6" v-motion-fade>
    <div>
      <h1 class="text-3xl font-bold text-slate-900 dark:text-white uppercase">{{ $route.params.slug?.[0]?.replace(/-/g, ' ') || 'Registration' }}</h1>
      <p class="text-slate-500 mt-1">Module: Registration > {{ $route.params.slug?.join(' > ') }}</p>
    </div>
    
    <div class="glass rounded-3xl p-12 flex flex-col items-center justify-center text-center border-2 border-dashed border-slate-200 dark:border-slate-800">
      <div class="w-20 h-20 rounded-2xl bg-slate-50 dark:bg-slate-900/20 flex items-center justify-center text-slate-500 mb-6">
        <LucideFileText class="w-10 h-10" />
      </div>
      <h2 class="text-xl font-black text-slate-800 dark:text-white">Feature in Development</h2>
      <p class="text-slate-500 max-w-sm mt-2 leading-relaxed">The specific feature for this path is currently being integrated into the core system.</p>
      <button @click="$router.back()" class="mt-8 px-6 py-2.5 bg-slate-100 hover:bg-slate-200 dark:bg-slate-800 rounded-xl text-xs font-black uppercase tracking-widest transition-all">Go Back</button>
    </div>
  </div>
</template>

<script setup>
import { LucideFileText } from 'lucide-vue-next'
</script>
