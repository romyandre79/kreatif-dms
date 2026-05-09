<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="space-y-1">
      <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('admin.import.excel.title') }}</h1>
      <p class="text-xs font-bold text-slate-500">{{ $t('admin.import.excel.subtitle') }}</p>
    </div>

    <!-- Stepper Section -->
    <div class="grid grid-cols-5 gap-4">
      <div v-for="s in [1, 2, 3, 4, 5]" :key="s" :class="`relative group cursor-pointer transition-all ${currentStep >= s ? 'opacity-100' : 'opacity-40 hover:opacity-60'}`">
        <div class="flex items-center gap-4">
          <div :class="`w-10 h-10 rounded-full flex items-center justify-center text-xs font-black transition-all ${currentStep === s ? 'bg-[#1E3A5F] text-white shadow-lg shadow-blue-900/20 scale-110' : currentStep > s ? 'bg-green-500 text-white' : 'bg-slate-200 dark:bg-slate-800 text-slate-400'}`">
            <LucideCheck v-if="currentStep > s" class="w-5 h-5" />
            <span v-else>{{ s.toString().padStart(2, '0') }}</span>
          </div>
          <div class="space-y-0.5">
            <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">PHASE {{ s.toString().padStart(2, '0') }}</p>
            <p :class="`text-[10px] font-black uppercase tracking-tight ${currentStep === s ? 'text-[#1E3A5F] dark:text-white' : 'text-slate-400'}`">{{ $t(`admin.import.excel.steps.s${s}`) }}</p>
          </div>
        </div>
        <div v-if="s < 5" :class="`absolute left-5 top-10 w-px h-6 bg-slate-100 dark:bg-slate-800 hidden md:block`" style="left: 20px; top: 40px; height: 1px; width: calc(100% - 40px)"></div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Main Area: Upload & Mapping -->
      <div class="lg:col-span-8 space-y-10">
        <!-- Upload Card -->
        <div class="glass p-12 rounded-[3rem] border-2 border-dashed border-slate-200 dark:border-slate-800 flex flex-col items-center justify-center text-center space-y-8 group hover:border-[#1E3A5F] transition-all">
          <div class="w-24 h-24 bg-slate-50 dark:bg-slate-800 rounded-[2rem] flex items-center justify-center text-slate-300 group-hover:scale-110 transition-transform"><LucideUploadCloud class="w-10 h-10" /></div>
          <div class="space-y-2">
            <h3 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('admin.import.excel.upload.title') }}</h3>
            <p class="text-xs font-bold text-slate-400 max-w-sm mx-auto leading-relaxed">{{ $t('admin.import.excel.upload.desc') }}</p>
          </div>
          <button class="px-10 py-4 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-2xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3">
            {{ $t('admin.import.excel.upload.btn') }}
            <LucideExternalLink class="w-4 h-4" />
          </button>
        </div>

        <!-- Mapping Interface -->
        <div class="glass rounded-[3rem] overflow-hidden">
          <div class="p-8 flex items-center justify-between border-b border-slate-50 dark:border-slate-800">
            <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('admin.import.excel.mapping.title') }}</h3>
            <button class="text-[9px] font-black text-blue-500 uppercase tracking-widest flex items-center gap-2 hover:underline transition-all">
              <LucideZap class="w-3.5 h-3.5" />
              {{ $t('admin.import.excel.mapping.btn_auto') }}
            </button>
          </div>
          <table class="w-full text-left">
            <thead>
              <tr class="bg-slate-50/50 dark:bg-slate-800/50 text-[8px] font-black text-slate-400 uppercase tracking-widest">
                <th class="p-6 pl-10">{{ $t('admin.import.excel.mapping.cols.source') }}</th>
                <th class="p-6">{{ $t('admin.import.excel.mapping.cols.target') }}</th>
                <th class="p-6">{{ $t('admin.import.excel.mapping.cols.type') }}</th>
                <th class="p-6 text-right pr-10">{{ $t('admin.import.excel.mapping.cols.conf') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
              <tr v-for="m in [
                { source: 'A: DOC_ID_REF', target: 'Document Reference ID', status: 'matched', type: 'STRING (UUID)', conf: '98%' },
                { source: 'B: CREATE_DT', target: 'Registration Date', status: 'suggested', type: 'DATETIME', conf: '72%' },
                { source: 'C: UNKNOWN_01', target: '-- Select Field --', status: 'unmapped', type: 'BLOB / TEXT', conf: '0%', color: 'text-red-500' }
              ]" :key="m.source" class="group hover:bg-slate-50/30 transition-all cursor-pointer">
                <td class="p-6 pl-10 font-bold text-xs text-slate-500 uppercase">{{ m.source }}</td>
                <td class="p-6">
                  <div class="flex items-center gap-4">
                    <div :class="`px-4 py-2 bg-slate-50 dark:bg-slate-800 rounded-lg text-[10px] font-bold border border-slate-100 dark:border-slate-800 ${m.status === 'unmapped' ? 'text-red-400 italic' : 'text-[#1E3A5F] dark:text-white'}`">{{ m.target }}</div>
                    <span :class="`text-[7px] font-black uppercase tracking-widest px-2 py-0.5 rounded border ${m.status === 'matched' ? 'bg-green-50 text-green-500 border-green-100' : m.status === 'suggested' ? 'bg-amber-50 text-amber-500 border-amber-100' : 'bg-red-50 text-red-500 border-red-100'}`">{{ $t(`admin.import.excel.mapping.status.${m.status}`) }}</span>
                  </div>
                </td>
                <td class="p-6 text-[10px] font-black text-slate-400 uppercase tracking-tight">{{ m.type }}</td>
                <td :class="`p-6 text-right pr-10 font-black text-xs ${m.color || 'text-blue-500'}`">{{ m.conf }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Right Column: Analysis & Security -->
      <div class="lg:col-span-4 space-y-10">
        <!-- Validation Summary -->
        <div class="glass p-10 rounded-[3rem] space-y-10 shadow-sm">
          <div class="flex items-center justify-between">
            <div class="space-y-1">
              <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.import.excel.summary.title') }}</h3>
              <p class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ $t('admin.import.excel.summary.rows', { count: '2,482' }) }}</p>
            </div>
            <div class="w-12 h-12 rounded-xl bg-blue-50 text-blue-500 flex items-center justify-center shadow-lg shadow-blue-500/10"><LucideBarChart3 class="w-6 h-6" /></div>
          </div>
          <div class="space-y-6">
            <div v-for="s in [
              { label: 'valid', val: '94.2%', color: 'bg-[#1E3A5F]' },
              { label: 'dup', val: '4.1%', color: 'bg-amber-500' },
              { label: 'errors', val: '1.7%', color: 'bg-red-500' }
            ]" :key="s.label" class="space-y-2">
              <div class="flex items-center justify-between text-[10px] font-black uppercase tracking-widest">
                <span class="text-slate-400">{{ $t(`admin.import.excel.summary.${s.label}`) }}</span>
                <span :class="s.label === 'valid' ? 'text-[#1E3A5F]' : s.label === 'dup' ? 'text-amber-500' : 'text-red-500'">{{ s.val }}</span>
              </div>
              <div class="h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
                <div :class="`h-full ${s.color}`" :style="{ width: s.val }"></div>
              </div>
            </div>
          </div>
        </div>

        <!-- Confidence Score -->
        <div class="bg-[#1E3A5F] p-10 rounded-[3rem] space-y-10 text-white shadow-2xl shadow-blue-900/40 relative overflow-hidden group">
          <div class="absolute top-0 right-0 w-32 h-32 bg-white/5 -rotate-12 translate-x-10 -translate-y-10 rounded-[3rem] flex items-end justify-start p-6 text-white/10 group-hover:scale-110 transition-transform"><LucideShieldCheck class="w-12 h-12" /></div>
          <div class="space-y-2">
            <h3 class="text-[10px] font-black uppercase tracking-widest opacity-60">{{ $t('admin.import.excel.confidence.title') }}</h3>
            <p class="text-5xl font-black tracking-tighter">88.4 <span class="text-xl opacity-40">/ 100</span></p>
          </div>
          <p class="text-xs font-bold leading-relaxed opacity-80">{{ $t('admin.import.excel.confidence.desc') }}</p>
          <button class="w-full py-4 bg-white/10 backdrop-blur-md rounded-2xl text-[9px] font-black uppercase tracking-widest border border-white/20 hover:bg-white/20 transition-all">
            {{ $t('admin.import.excel.confidence.btn') }}
          </button>
        </div>

        <!-- Execution Security -->
        <div class="glass p-10 rounded-[3rem] space-y-8">
          <div class="flex items-center gap-4">
            <LucideLock class="w-5 h-5 text-amber-500" />
            <h3 class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('admin.import.excel.security.title') }}</h3>
          </div>
          <p class="text-[10px] font-bold text-slate-500 leading-relaxed uppercase">{{ $t('admin.import.excel.security.desc') }}</p>
          <div class="flex gap-3 justify-center">
            <div v-for="i in 4" :key="i" class="w-12 h-14 bg-slate-50 dark:bg-slate-800 rounded-xl border border-slate-100 dark:border-slate-700 flex items-center justify-center text-xl font-black text-[#1E3A5F] dark:text-white">
              <span v-if="i <= pin.length" class="w-2 h-2 bg-[#1E3A5F] dark:bg-white rounded-full"></span>
            </div>
          </div>
          <button @click="pin = '12'" class="w-full py-5 bg-[#1E3A5F]/40 text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-xl shadow-blue-900/10 cursor-not-allowed">
            {{ $t('admin.import.excel.security.btn') }}
          </button>
        </div>
      </div>
    </div>

    <!-- Job Queue Monitor -->
    <div class="bg-slate-50/50 dark:bg-slate-900/50 p-10 rounded-[3rem] border border-slate-100 dark:border-slate-800 space-y-10 shadow-inner" v-motion-slide-visible-bottom>
      <div class="flex items-center justify-between">
        <div class="space-y-1">
          <h3 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('admin.import.excel.queue.title') }}</h3>
          <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('admin.import.excel.queue.subtitle') }}</p>
        </div>
        <span class="px-4 py-1.5 bg-white dark:bg-slate-800 text-[#1E3A5F] dark:text-white rounded-full text-[9px] font-black uppercase tracking-widest border border-slate-100 dark:border-slate-700 flex items-center gap-2 shadow-sm">
          <span class="w-1.5 h-1.5 bg-blue-500 rounded-full animate-pulse"></span>
          {{ $t('admin.import.excel.queue.active', { id: 'IMPR-8821' }) }}
        </span>
      </div>

      <div class="p-10 bg-white dark:bg-slate-900 rounded-lg border border-slate-100 dark:border-slate-800 shadow-xl space-y-10 relative overflow-hidden">
        <div class="flex flex-col md:flex-row md:items-center justify-between gap-10">
          <div class="flex-grow space-y-6">
            <div class="flex items-center gap-6">
              <span class="px-4 py-1.5 bg-blue-50 text-blue-500 rounded-lg text-[9px] font-black uppercase tracking-widest border border-blue-100">PROCESSING</span>
              <h4 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">DMS_BULK_UPLOAD_Q3_REPORTS.XLSX</h4>
            </div>
            <div class="space-y-4">
              <div class="h-2.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
                <div class="h-full bg-[#1E3A5F] transition-all duration-1000" style="width: 64%"></div>
              </div>
              <div class="grid grid-cols-3 gap-10">
                <div class="flex items-center gap-4">
                  <div class="w-10 h-10 rounded-xl bg-green-50 text-green-500 flex items-center justify-center shadow-sm"><LucideCheckCircle2 class="w-5 h-5" /></div>
                  <div class="space-y-0.5">
                    <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.import.excel.queue.stats.processed') }}</p>
                    <p class="text-sm font-black text-[#1E3A5F] dark:text-white">1,588 <span class="text-[10px] text-slate-300">/ 2,482</span></p>
                  </div>
                </div>
                <div class="flex items-center gap-4">
                  <div class="w-10 h-10 rounded-xl bg-red-50 text-red-500 flex items-center justify-center shadow-sm"><LucideAlertCircle class="w-5 h-5" /></div>
                  <div class="space-y-0.5">
                    <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.import.excel.queue.stats.failed') }}</p>
                    <p class="text-sm font-black text-red-500">12 <span class="text-[10px] text-red-300">Rows</span></p>
                  </div>
                </div>
                <div class="flex items-center gap-4">
                  <div class="w-10 h-10 rounded-xl bg-slate-50 dark:bg-slate-800 text-slate-400 flex items-center justify-center shadow-sm"><LucideClock class="w-5 h-5" /></div>
                  <div class="space-y-0.5">
                    <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.import.excel.queue.stats.est') }}</p>
                    <p class="text-sm font-black text-slate-600 dark:text-slate-300">2m 14s</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="flex flex-col gap-4 min-w-[200px]">
            <button class="w-full py-4 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-slate-600 dark:text-slate-300 rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-50 transition-all">{{ $t('admin.import.excel.queue.actions.pause') }}</button>
            <button class="w-full py-4 bg-red-50 text-red-500 rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-red-100 transition-all">{{ $t('admin.import.excel.queue.actions.abort') }}</button>
          </div>
        </div>
        <div class="absolute right-10 top-1/2 -translate-y-1/2 text-[80px] font-black text-[#1E3A5F]/5 pointer-events-none select-none">64%</div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideCheck, LucideUploadCloud, LucideExternalLink, LucideZap, 
  LucideBarChart3, LucideShieldCheck, LucideLock, LucideCheckCircle2, 
  LucideAlertCircle, LucideClock 
} from 'lucide-vue-next'

const currentStep = ref(1)
const pin = ref('12')

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
@reference "tailwindcss";
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #E2E8F0; border-radius: 10px; }
.dark .custom-scrollbar::-webkit-scrollbar-thumb { background: #1E293B; }
</style>

