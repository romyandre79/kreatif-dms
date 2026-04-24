<template>
  <div class="min-h-[60vh] flex flex-col items-center justify-center p-8 text-center" v-motion-fade>
    <div class="w-24 h-24 rounded-3xl bg-primary-50 dark:bg-primary-900/20 flex items-center justify-center text-primary-500 mb-8 shadow-xl shadow-primary-500/10">
      <LucideConstruction class="w-12 h-12" />
    </div>
    
    <h1 class="text-4xl font-black text-slate-900 dark:text-white mb-4 tracking-tight">Feature in Development</h1>
    <p class="text-slate-500 dark:text-slate-400 max-w-md text-lg leading-relaxed mb-10">
      The page for <span class="font-black text-primary-600">/{{ $route.params.slug?.join('/') }}</span> is currently under construction as part of our premium DMS experience.
    </p>
    
    <div class="flex gap-4">
      <button @click="$router.back()" class="px-8 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-2xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all">
        Go Back
      </button>
      <NuxtLink to="/dashboard" class="px-8 py-3.5 bg-[#1E3A5F] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-lg shadow-blue-900/20 hover:bg-[#152943] transition-all">
        Return to Dashboard
      </NuxtLink>
    </div>
    
    <div class="mt-20 grid grid-cols-1 md:grid-cols-3 gap-6 w-full max-w-4xl">
      <div class="glass p-6 rounded-2xl border-l-4 border-blue-500 text-left">
        <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">Next Step</p>
        <p class="text-sm font-bold text-slate-700 dark:text-slate-200">Database Schema Integration</p>
      </div>
      <div class="glass p-6 rounded-2xl border-l-4 border-teal-500 text-left">
        <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">Status</p>
        <p class="text-sm font-bold text-slate-700 dark:text-slate-200">Frontend UI Blueprint Ready</p>
      </div>
      <div class="glass p-6 rounded-2xl border-l-4 border-orange-500 text-left">
        <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">Priority</p>
        <p class="text-sm font-bold text-slate-700 dark:text-slate-200">High Operations Module</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { LucideConstruction } from 'lucide-vue-next'
</script>
