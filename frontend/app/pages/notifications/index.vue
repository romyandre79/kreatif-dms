<template>
  <div class="space-y-6" v-motion-fade>
    <div>
      <h1 class="text-3xl font-bold text-slate-900 dark:text-white">All Notifications</h1>
      <p class="text-slate-500 mt-1">Review your recent system alerts and updates.</p>
    </div>
    
    <div class="glass rounded-3xl p-12 flex flex-col items-center justify-center text-center border-2 border-dashed border-slate-200 dark:border-slate-800">
      <div class="w-20 h-20 rounded-2xl bg-orange-50 dark:bg-orange-900/20 flex items-center justify-center text-orange-500 mb-6">
        <LucideBell class="w-10 h-10" />
      </div>
      <h2 class="text-xl font-black text-slate-800 dark:text-white">Notification Inbox</h2>
      <p class="text-slate-500 max-w-sm mt-2 leading-relaxed">No new notifications. We'll alert you when something important happens.</p>
    </div>
  </div>
</template>

<script setup>
import { LucideBell } from 'lucide-vue-next'
</script>
