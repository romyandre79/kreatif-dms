<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Header Banner -->
    <div class="bg-red-50 border border-red-100 p-10 rounded-lg flex flex-col md:flex-row md:items-center gap-10 relative overflow-hidden group">
      <div class="absolute top-0 right-0 w-64 h-64 bg-red-100/50 rounded-full -translate-x-10 -translate-y-20 group-hover:scale-110 transition-transform"></div>
      <div class="w-16 h-16 rounded-lg bg-red-500 text-white flex items-center justify-center shrink-0 shadow-xl shadow-red-500/20 z-10">
        <LucideAlertTriangle class="w-8 h-8" />
      </div>
      <div class="space-y-2 flex-grow z-10">
        <div class="flex items-center gap-4">
          <h1 class="text-2xl font-black text-red-600 uppercase tracking-tight">{{ $t('intake.duplicate.banner.title') }}</h1>
          <span class="px-3 py-1 bg-red-500 text-white rounded-full text-[8px] font-black uppercase tracking-widest shadow-lg">{{ $t('intake.duplicate.banner.risk', { val: 94 }) }}</span>
        </div>
        <p class="text-xs font-bold text-red-400 leading-relaxed max-w-3xl">
          {{ $t('intake.duplicate.banner.desc') }}
        </p>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Top: Data Comparison Matrix -->
      <div class="lg:col-span-8 glass rounded-lg overflow-hidden shadow-sm">
        <div class="p-8 flex items-center justify-between border-b border-slate-50 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-800/50">
          <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('intake.duplicate.matrix.title') }}</h3>
          <span class="px-4 py-1 bg-blue-50 text-blue-500 rounded-lg text-[9px] font-black font-mono">REF-0982-X</span>
        </div>
        <table class="w-full text-left">
          <thead>
            <tr class="text-[8px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-50 dark:border-slate-800">
              <th class="p-6 pl-10">{{ $t('intake.duplicate.matrix.cols.attr') }}</th>
              <th class="p-6">{{ $t('intake.duplicate.matrix.cols.curr') }}</th>
              <th class="p-6 pr-10">{{ $t('intake.duplicate.matrix.cols.match') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
            <tr v-for="attr in [
              { id: 'id', curr: 'DOC-2024-TMP-001', match: 'AKR-ADM-2023-455', color: 'text-blue-500' },
              { id: 'file', curr: 'Laporan_Audit_Q4_Final.pdf', match: 'Laporan_Audit_Q4_2023.pdf' },
              { id: 'checksum', curr: 'e96a18c428cb38d5f268...', match: 'e99a18c428cb38d5f268...', font: 'font-mono' },
              { id: 'dept', curr: 'Finance & Audit', match: 'Finance & Internal Audit' },
              { id: 'uploader', curr: 'Admin_Jakarta', match: 'System_Auto_Archiver' }
            ]" :key="attr.id" class="group hover:bg-slate-50/30 transition-all">
              <td class="p-6 pl-10 text-[10px] font-black text-slate-400 uppercase">{{ $t(`intake.duplicate.matrix.attrs.${attr.id}`) }}</td>
              <td :class="`p-6 text-[11px] font-black text-slate-600 dark:text-slate-300 uppercase ${attr.font || ''}`">{{ attr.curr }}</td>
              <td :class="`p-6 pr-10 text-[11px] font-black ${attr.color || 'text-slate-600 dark:text-slate-300'} uppercase ${attr.font || ''}`">{{ attr.match }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Right Column: Decision & Audit -->
      <div class="lg:col-span-4 space-y-10">
        <!-- Decision Engine -->
        <div class="glass p-10 rounded-lg space-y-10 shadow-sm border border-slate-50 dark:border-slate-800">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.duplicate.decision.title') }}</h3>
          <div class="space-y-4">
            <button class="w-full py-5 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-2xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center justify-center gap-4 group">
              <LucideSave class="w-5 h-5 group-hover:scale-110 transition-transform" />
              {{ $t('intake.duplicate.decision.save') }}
            </button>
            <button class="w-full py-5 bg-slate-100 dark:bg-slate-800 text-[#1E3A5F] dark:text-white rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-200 transition-all flex items-center justify-center gap-4">
              <LucideLink class="w-5 h-5" />
              {{ $t('intake.duplicate.decision.link') }}
            </button>
            <div class="py-4 border-t border-slate-50 dark:border-slate-800 flex flex-col gap-4">
              <button class="w-full py-5 bg-white dark:bg-slate-900 border-2 border-red-50 text-red-500 rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-red-50 transition-all flex items-center justify-center gap-4">
                <LucideShieldAlert class="w-5 h-5" />
                {{ $t('intake.duplicate.decision.escalate') }}
              </button>
              <button class="text-[10px] font-black text-slate-300 uppercase tracking-widest hover:text-slate-500 transition-colors text-center">
                {{ $t('intake.duplicate.decision.cancel') }}
              </button>
            </div>
          </div>
        </div>

        <!-- Audit Context -->
        <div class="bg-[#1E3A5F] p-10 rounded-lg text-white space-y-6 relative overflow-hidden group shadow-2xl shadow-blue-900/40">
          <LucideSearch class="absolute top-0 right-0 w-32 h-32 text-white/5 -rotate-12 translate-x-10 -translate-y-10 group-hover:scale-110 transition-transform" />
          <div class="space-y-1">
            <p class="text-[10px] font-black uppercase tracking-widest opacity-60">{{ $t('intake.duplicate.audit.title') }}</p>
            <h4 class="text-4xl font-black tracking-tighter">{{ $t('intake.duplicate.audit.matches', { count: 23 }) }}</h4>
          </div>
          <p class="text-[10px] font-bold uppercase tracking-widest opacity-80 leading-relaxed">
            {{ $t('intake.duplicate.audit.desc') }}
          </p>
        </div>
      </div>
    </div>

    <!-- Bottom Comparison Preview -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-10">
      <!-- Current Upload Preview -->
      <div class="space-y-6">
        <div class="flex items-center justify-between px-6">
          <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.duplicate.preview.curr') }}</h4>
          <LucideEye class="w-4 h-4 text-slate-300" />
        </div>
        <div class="glass rounded-lg h-[500px] relative overflow-hidden border border-slate-50 dark:border-slate-800 flex items-center justify-center">
          <!-- Skeleton Content with Overlay -->
          <div class="w-full h-full p-16 space-y-8 opacity-20">
            <div class="h-10 bg-slate-200 rounded-lg w-3/4"></div>
            <div class="h-6 bg-slate-200 rounded-lg w-full"></div>
            <div class="h-6 bg-slate-200 rounded-lg w-full"></div>
            <div class="h-32 bg-slate-200 rounded-2xl w-full"></div>
            <div class="h-6 bg-slate-200 rounded-lg w-2/3"></div>
          </div>
          <!-- Diff Overlay -->
          <div class="absolute inset-0 flex items-center justify-center pointer-events-none">
            <div class="px-8 py-4 bg-red-500/90 text-white rounded-xl shadow-2xl backdrop-blur-md transform -rotate-2 border-2 border-white/20 animate-pulse flex items-center gap-4">
              <LucideAlertCircle class="w-6 h-6" />
              <span class="text-sm font-black uppercase tracking-widest">{{ $t('intake.duplicate.preview.diff') }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Matched Archive Preview -->
      <div class="space-y-6">
        <div class="flex items-center justify-between px-6">
          <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.duplicate.preview.matched', { id: 'AKR-ADM-2023' }) }}</h4>
          <LucideHistory class="w-4 h-4 text-slate-300" />
        </div>
        <div class="glass rounded-lg h-[500px] border border-slate-50 dark:border-slate-800 flex items-center justify-center">
          <div class="p-10 bg-white dark:bg-slate-900 rounded-lg shadow-2xl border border-slate-100 dark:border-slate-800 flex flex-col items-center gap-6 group hover:scale-105 transition-all">
            <div class="w-20 h-20 rounded-lg bg-slate-50 dark:bg-slate-800 flex items-center justify-center text-slate-200 group-hover:text-blue-500 transition-colors">
              <LucideLock class="w-10 h-10" />
            </div>
            <div class="text-center space-y-1">
              <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('intake.duplicate.preview.archived') }}</p>
              <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('intake.duplicate.preview.modified', { date: '12 Jan 2024' }) }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideAlertTriangle, LucideSave, LucideLink, LucideShieldAlert, 
  LucideSearch, LucideEye, LucideAlertCircle, LucideHistory, LucideLock 
} from 'lucide-vue-next'

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
@reference "tailwindcss";
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}
</style>

