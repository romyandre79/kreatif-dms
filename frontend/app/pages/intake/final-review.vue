<template>
  <div class="space-y-12 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="space-y-4">
      <div class="flex items-center gap-4">
        <span class="px-4 py-1.5 bg-blue-50 text-blue-500 rounded-lg text-[9px] font-black uppercase tracking-widest border border-blue-100 shadow-sm shadow-blue-500/10">{{ $t('intake.final_review.badge') }}</span>
      </div>
      <div class="space-y-1">
        <h1 class="text-4xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('intake.final_review.title') }}</h1>
        <p class="text-sm font-bold text-slate-500 uppercase tracking-tight">{{ $t('intake.final_review.subtitle') }}</p>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Column: Validation & Output -->
      <div class="lg:col-span-7 space-y-10">
        <!-- Top Stats -->
        <div class="grid grid-cols-2 gap-8">
          <div class="glass p-10 rounded-[3.5rem] space-y-4 shadow-sm border border-slate-50 dark:border-slate-800 border-b-8 border-blue-500">
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.final_review.stats.completeness') }}</p>
            <div class="flex items-center gap-6">
              <p class="text-5xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">95%</p>
              <LucideCheckCircle2 class="w-8 h-8 text-green-500 shadow-xl shadow-green-500/20" />
            </div>
          </div>
          <div class="glass p-10 rounded-[3.5rem] space-y-4 shadow-sm border border-slate-50 dark:border-slate-800 border-b-8 border-[#1E3A5F]">
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.final_review.stats.ocr') }}</p>
            <div class="flex items-center gap-6">
              <p class="text-5xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">92%</p>
              <LucideCheckCircle2 class="w-8 h-8 text-green-500 shadow-xl shadow-green-500/20" />
            </div>
          </div>
        </div>

        <!-- Validation Table -->
        <div class="glass rounded-[4rem] overflow-hidden shadow-sm border border-slate-50 dark:border-slate-800 bg-white">
          <div class="p-10 border-b border-slate-50 dark:border-slate-800 flex items-center gap-4">
            <LucideTerminal class="w-6 h-6 text-[#1E3A5F]" />
            <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('intake.final_review.validation.title') }}</h3>
          </div>
          <div class="p-4">
            <table class="w-full text-left">
              <tbody class="divide-y divide-slate-50">
                <tr v-for="row in [
                  { key: 'duplicate', val: 'passed', icon: LucideCheckCircle2, color: 'text-green-500' },
                  { key: 'id', val: 'AKD-2024-QX-09882', mono: true },
                  { key: 'name', val: 'Legal_Procurement_PT_Akiradata_Q1.pdf' },
                  { key: 'path', val: 'ROOT / ARCHIVE / 2024 / PROCUREMENT / LEGAL', icon: LucideFolder, color: 'text-blue-500' },
                  { key: 'channel', val: 'SINGLE INPUT', badge: true }
                ]" :key="row.key" class="group hover:bg-slate-50/50 transition-all">
                  <td class="p-8 pl-10 w-1/3 text-[11px] font-black text-slate-400 uppercase tracking-widest">{{ $t(`intake.final_review.validation.${row.key}`) }}</td>
                  <td class="p-8 pr-10">
                    <div class="flex items-center gap-4">
                      <component :is="row.icon" v-if="row.icon" :class="`w-4 h-4 ${row.color}`" />
                      <span v-if="row.badge" class="px-3 py-1 bg-slate-100 text-slate-400 rounded-lg text-[8px] font-black uppercase tracking-widest">{{ row.val }}</span>
                      <p v-else :class="`text-[13px] font-black uppercase tracking-tight ${row.mono ? 'font-mono text-blue-500' : 'text-[#1E3A5F]'} ${row.key === 'passed' ? 'text-green-500' : ''}`">
                        {{ row.key === 'passed' ? $t('intake.final_review.validation.passed') : row.val }}
                      </p>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- Verification Preview -->
        <div class="glass p-10 rounded-[4rem] shadow-sm border border-slate-50 dark:border-slate-800 bg-slate-900 overflow-hidden relative group">
          <div class="aspect-video bg-slate-800 rounded-lg flex items-center justify-center relative overflow-hidden">
            <!-- Digital Twin Lines -->
            <div class="absolute inset-0 opacity-10 pointer-events-none">
              <div class="absolute top-0 left-0 w-full h-full grid grid-cols-12 gap-1">
                <div v-for="i in 12" :key="i" class="h-full border-l border-white/50"></div>
              </div>
            </div>
            <p class="text-[10px] font-black text-white/40 uppercase tracking-[0.5em] animate-pulse relative z-10">{{ $t('intake.final_review.verification') }}</p>
          </div>
        </div>
      </div>

      <!-- Right Column: Risk & Controls -->
      <div class="lg:col-span-5 space-y-10">
        <!-- Risk Checklist -->
        <div class="glass p-12 rounded-[4rem] space-y-10 shadow-sm border border-slate-50 dark:border-slate-800">
          <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('intake.final_review.risk.title') }}</h3>
          <div class="space-y-8">
            <div v-for="risk in ['pii', 'ldap', 'hash', 'retention']" :key="risk" class="flex items-start gap-6 group">
              <div class="w-10 h-10 rounded-xl bg-green-500 text-white flex items-center justify-center shrink-0 shadow-lg shadow-green-500/20 group-hover:scale-110 transition-transform">
                <LucideCheck class="w-6 h-6" />
              </div>
              <div class="space-y-1">
                <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t(`intake.final_review.risk.${risk}`) }}</h4>
                <p class="text-[10px] font-bold text-slate-400 uppercase leading-relaxed tracking-tight">{{ $t(`intake.final_review.risk.${risk}_desc`) }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Publish Controls -->
        <div class="glass p-12 rounded-[4rem] space-y-12 shadow-sm border border-slate-50 dark:border-slate-800 relative bg-white">
          <div class="text-center space-y-1">
            <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('intake.final_review.controls.title') }}</h3>
          </div>

          <div class="space-y-6">
            <button class="w-full py-6 bg-[#1E3A5F] text-white rounded-[2rem] text-sm font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40 hover:bg-[#152943] transition-all flex items-center justify-center gap-6 active:scale-95 group">
              <LucideSend class="w-6 h-6 group-hover:translate-x-2 group-hover:-translate-y-2 transition-transform" />
              {{ $t('intake.final_review.controls.commit') }}
            </button>
            <button class="w-full py-6 bg-slate-100 dark:bg-slate-800 text-[#1E3A5F] dark:text-white rounded-[2rem] text-sm font-black uppercase tracking-widest hover:bg-slate-200 transition-all flex items-center justify-center gap-6 active:scale-95 group">
              <LucideUsers class="w-6 h-6" />
              {{ $t('intake.final_review.controls.supervisor') }}
            </button>
            <div class="grid grid-cols-2 gap-6 pt-6 border-t border-slate-50">
              <button class="py-5 bg-white dark:bg-slate-900 border-2 border-slate-50 dark:border-slate-700 text-[10px] font-black text-[#1E3A5F] dark:text-white rounded-2xl hover:bg-slate-50 transition-all flex items-center justify-center gap-3">
                <LucideSave class="w-4 h-4" />
                {{ $t('intake.final_review.controls.draft') }}
              </button>
              <button class="py-5 bg-red-50 text-red-500 border-2 border-red-100 rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-red-100 transition-all flex items-center justify-center gap-3">
                <LucideTrash2 class="w-4 h-4" />
                {{ $t('intake.final_review.controls.purge') }}
              </button>
            </div>
          </div>

          <div class="text-center space-y-2 pt-10 border-t border-slate-50">
            <p class="text-[9px] font-bold text-slate-300 uppercase tracking-widest">{{ $t('intake.final_review.controls.auth', { id: 'AUTH_98X_DMS' }) }}</p>
            <p class="text-[9px] font-bold text-slate-300 uppercase tracking-widest">{{ $t('intake.final_review.controls.time', { val: '2024-10-27 14:32:01 WIB' }) }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideCheckCircle2, LucideTerminal, LucideFolder, 
  LucideCheck, LucideSend, LucideUsers, LucideSave, LucideTrash2 
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

