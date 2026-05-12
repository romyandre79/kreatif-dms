<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-1">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('admin.backup.title') }}</h1>
        <p class="text-xs font-bold text-slate-500">{{ $t('admin.backup.subtitle') }}</p>
      </div>
      <div class="flex items-center gap-4">
        <button class="px-6 py-3 text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest flex items-center gap-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl transition-all border border-slate-200 dark:border-slate-700">
          {{ $t('admin.backup.header.verify') }}
        </button>
        <button class="px-10 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 group">
          <LucidePlay class="w-4 h-4 fill-current group-hover:scale-110 transition-transform" />
          {{ $t('admin.backup.header.run') }}
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Column: Schedule Settings -->
      <div class="lg:col-span-3 space-y-10">
        <div class="glass p-10 rounded-lg space-y-10 shadow-sm border-t-8 border-slate-100 dark:border-slate-800">
          <div class="flex items-center gap-4">
            <LucideClock class="w-5 h-5 text-blue-500" />
            <h3 class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('admin.backup.settings.title') }}</h3>
          </div>
          
          <div class="space-y-8">
            <!-- Frequency -->
            <div class="space-y-3">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.backup.settings.freq') }}</label>
              <div class="relative">
                <select class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold outline-none appearance-none cursor-pointer">
                  <option>Daily at 02:00 AM</option>
                  <option>Weekly on Sunday</option>
                </select>
                <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300 pointer-events-none" />
              </div>
            </div>

            <!-- Scope -->
            <div class="space-y-4">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.backup.settings.scope') }}</label>
              <div class="space-y-3">
                <div v-for="s in ['full', 'db', 'media']" :key="s" class="p-5 bg-white dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl flex items-center gap-4 cursor-pointer group hover:border-[#1E3A5F] transition-all" @click="selectedScope = s">
                  <div :class="`w-5 h-5 rounded-full border-2 flex items-center justify-center transition-all ${selectedScope === s ? 'border-[#1E3A5F]' : 'border-slate-200'}`">
                    <div v-if="selectedScope === s" class="w-2.5 h-2.5 bg-[#1E3A5F] rounded-full"></div>
                  </div>
                  <div class="space-y-0.5">
                    <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t(`admin.backup.settings.scopes.${s}`) }}</p>
                    <p v-if="s === 'full'" class="text-[8px] font-bold text-slate-400 uppercase tracking-widest">Entire database & filesystem</p>
                    <p v-if="s === 'db'" class="text-[8px] font-bold text-slate-400 uppercase tracking-widest">SQL & Key-Value stores</p>
                    <p v-if="s === 'media'" class="text-[8px] font-bold text-slate-400 uppercase tracking-widest">Attachments and archives</p>
                  </div>
                </div>
              </div>
            </div>

            <!-- Destination -->
            <div class="space-y-3">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.backup.settings.dest') }}</label>
              <div class="p-6 bg-slate-50/50 dark:bg-slate-900/50 rounded-2xl border border-slate-100 dark:border-slate-800 flex items-center justify-between">
                <div class="flex items-center gap-4">
                  <LucideCloud class="w-5 h-5 text-slate-400" />
                  <span class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase">AWS S3 - Tokyo-02</span>
                </div>
                <span class="text-[7px] font-black text-green-500 uppercase tracking-widest">CONNECTED</span>
              </div>
            </div>

            <!-- Retention -->
            <div class="space-y-3">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.backup.settings.retention') }}</label>
              <div class="flex items-center gap-4">
                <input type="text" value="30" class="w-16 px-4 py-4 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-xl text-xs font-bold text-center outline-none" />
                <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.backup.settings.retention_val', { days: '', iters: 50 }) }}</span>
              </div>
            </div>

            <button class="w-full py-4 bg-slate-100 dark:bg-slate-800 text-[#1E3A5F] dark:text-white rounded-2xl text-[9px] font-black uppercase tracking-widest hover:bg-slate-200 transition-all">
              {{ $t('admin.backup.settings.btn_update') }}
            </button>
          </div>
        </div>

        <!-- Health Check -->
        <div class="bg-blue-50/50 dark:bg-blue-900/10 p-8 rounded-lg border border-blue-100 dark:border-blue-800 space-y-6">
          <div class="flex items-center justify-between">
            <h4 class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('admin.backup.settings.health.title') }}</h4>
            <LucideShieldCheck class="w-4 h-4 text-[#1E3A5F]" />
          </div>
          <div class="space-y-6">
            <div class="flex items-center justify-between">
              <span class="text-[10px] font-bold text-slate-500 uppercase">{{ $t('admin.backup.settings.health.last') }}</span>
              <span class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase">3h 24m ago</span>
            </div>
            <div class="space-y-2">
              <div class="flex items-center justify-between text-[10px] font-bold text-slate-500 uppercase">
                <span>{{ $t('admin.backup.settings.health.storage') }}</span>
                <span class="font-black text-[#1E3A5F] dark:text-white">4.2 TB</span>
              </div>
              <div class="h-1.5 bg-white dark:bg-slate-800 rounded-full overflow-hidden">
                <div class="h-full bg-[#1E3A5F]" style="width: 65%"></div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Center Column: History -->
      <div class="lg:col-span-6 space-y-10">
        <div class="glass rounded-lg overflow-hidden shadow-xl border-t-8 border-slate-50 dark:border-slate-800">
          <div class="p-10 flex items-center justify-between border-b border-slate-50 dark:border-slate-800">
            <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('admin.backup.history.title') }}</h3>
            <div class="flex items-center gap-3 text-[9px] font-black text-slate-300 uppercase tracking-widest cursor-pointer hover:text-slate-500 transition-colors">
              <LucideFilter class="w-4 h-4" />
              {{ $t('admin.backup.history.filter') }}
            </div>
          </div>
          <table class="w-full text-left">
            <thead>
              <tr class="bg-slate-50/50 dark:bg-slate-800/50 text-[8px] font-black text-slate-400 uppercase tracking-widest">
                <th class="p-8 pl-10">{{ $t('admin.backup.history.cols.id') }}</th>
                <th class="p-8">{{ $t('admin.backup.history.cols.time') }}</th>
                <th class="p-8">{{ $t('admin.backup.history.cols.size') }}</th>
                <th class="p-8">{{ $t('admin.backup.history.cols.status') }}</th>
                <th class="p-8 pr-10 text-right">{{ $t('admin.backup.history.cols.checksum') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
              <tr v-for="h in [
                { id: 'BK-9928', date: '20 Oct 2023', time: '02:00', duration: '12m 45s', size: '12.4 GB', status: 'ok', checksum: 'e3b9c442' },
                { id: 'BK-9927', date: '19 Oct 2023', time: '02:00', duration: '11m 58s', size: '12.2 GB', status: 'ok', checksum: 'a94af8e5' },
                { id: 'BK-9926', date: '18 Oct 2023', time: '02:00', duration: '4m 12s', size: '4.1 GB', status: 'fail', checksum: 'TIMEOUT ERROR', color: 'text-red-500' },
                { id: 'BK-9925', date: '17 Oct 2023', time: '02:00', duration: '12m 45s', size: '12.5 GB', status: 'ok', checksum: 'f1e2d3c4' },
                { id: 'BK-9924', date: '16 Oct 2023', time: '02:00', duration: '12m 30s', size: '12.3 GB', status: 'ok', checksum: 'c5b1a3f2' }
              ]" :key="h.id" class="group hover:bg-slate-50/30 transition-all cursor-pointer">
                <td class="p-8 pl-10 text-xs font-black text-blue-500 uppercase font-mono">#{{ h.id }}</td>
                <td class="p-8">
                  <div class="space-y-0.5">
                    <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ h.date }}</p>
                    <p class="text-[8px] font-bold text-slate-400 uppercase">{{ h.time }} <span class="ml-2">Duration: {{ h.duration }}</span></p>
                  </div>
                </td>
                <td class="p-8 text-[11px] font-black text-slate-600 dark:text-slate-300 uppercase">{{ h.size }}</td>
                <td class="p-8">
                  <span :class="`text-[9px] font-black uppercase tracking-widest ${h.status === 'ok' ? 'text-slate-800' : 'text-red-500 bg-red-50 px-3 py-1 rounded'}`">{{ $t(`admin.backup.history.status.${h.status}`) }}</span>
                </td>
                <td :class="`p-8 pr-10 text-right text-[10px] font-black uppercase font-mono ${h.color || 'text-slate-300'}`">{{ h.checksum }}</td>
              </tr>
            </tbody>
          </table>
          <div class="p-8 flex items-center justify-between border-t border-slate-50 dark:border-slate-800 text-[10px] font-black text-slate-400 uppercase">
            <span>Showing 5 of 1,245 snapshots</span>
            <div class="flex gap-4">
              <LucideChevronLeft class="w-4 h-4 cursor-pointer hover:text-[#1E3A5F]" />
              <LucideChevronRight class="w-4 h-4 cursor-pointer hover:text-[#1E3A5F]" />
            </div>
          </div>
        </div>

        <!-- Redundancy & Encryption Info -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-10">
          <div class="p-8 bg-slate-50/50 dark:bg-slate-900/50 rounded-lg border border-slate-100 dark:border-slate-800 space-y-4 shadow-inner relative overflow-hidden group">
            <LucideGlobe class="absolute top-0 right-0 w-32 h-32 text-slate-200/30 -rotate-12 translate-x-10 -translate-y-10 group-hover:scale-110 transition-transform" />
            <div class="space-y-1">
              <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.backup.history.redundancy.title') }}</p>
              <h4 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('admin.backup.history.redundancy.val') }}</h4>
            </div>
            <div class="flex items-center gap-3 text-[9px] font-black text-green-500 uppercase tracking-widest">
              <LucideCheckCircle2 class="w-3.5 h-3.5" />
              {{ $t('admin.backup.history.redundancy.desc') }}
            </div>
          </div>
          <div class="p-8 bg-slate-50/50 dark:bg-slate-900/50 rounded-lg border border-slate-100 dark:border-slate-800 space-y-4 shadow-inner relative overflow-hidden group">
            <LucideLock class="absolute top-0 right-0 w-32 h-32 text-slate-200/30 -rotate-12 translate-x-10 -translate-y-10 group-hover:scale-110 transition-transform" />
            <div class="space-y-1">
              <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.backup.history.encryption.title') }}</p>
              <h4 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('admin.backup.history.encryption.val') }}</h4>
            </div>
            <div class="flex items-center gap-3 text-[9px] font-black text-blue-500 uppercase tracking-widest">
              <LucideShieldCheck class="w-3.5 h-3.5" />
              {{ $t('admin.backup.history.encryption.desc') }}
            </div>
          </div>
        </div>
      </div>

      <!-- Right Column: Restore Wizard -->
      <div class="lg:col-span-3 space-y-10">
        <div class="glass p-10 rounded-lg space-y-12 shadow-2xl border border-blue-50 dark:border-blue-900/30 relative">
          <div class="flex items-center gap-4 border-b border-slate-50 dark:border-slate-800 pb-8">
            <LucideRotateCcw class="w-6 h-6 text-blue-500" />
            <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('admin.backup.restore.title') }}</h3>
          </div>

          <!-- Step 1: Point Selector -->
          <div class="space-y-6">
            <div class="flex items-center justify-between">
              <h4 class="text-[11px] font-black text-slate-800 dark:text-white uppercase tracking-tight">{{ $t('admin.backup.restore.s1.title') }}</h4>
              <span class="text-[8px] font-black text-blue-500 uppercase tracking-widest">Recommended</span>
            </div>
            <div class="p-8 bg-blue-50 border-2 border-blue-100 rounded-3xl space-y-4 relative group cursor-pointer transition-all">
              <LucideCheckCircle2 class="absolute top-6 right-6 w-5 h-5 text-blue-500" />
              <div class="space-y-1">
                <p class="text-xs font-black text-[#1E3A5F] uppercase leading-tight">{{ $t('admin.backup.restore.s1.latest') }}</p>
                <p class="text-[9px] font-bold text-blue-400 uppercase">20 Oct 2023 — 02:00 AM</p>
              </div>
            </div>
            <button class="w-full py-4 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-700 rounded-2xl text-[9px] font-black text-slate-400 uppercase tracking-widest hover:border-[#1E3A5F] transition-all">
              {{ $t('admin.backup.restore.s1.manual') }}
            </button>
          </div>

          <!-- Step 2: Impact Preview -->
          <div class="space-y-6">
            <h4 class="text-[11px] font-black text-slate-800 dark:text-white uppercase tracking-tight">{{ $t('admin.backup.restore.s2.title') }}</h4>
            <div class="space-y-6 px-4">
              <div class="flex items-center justify-between border-b border-slate-50 dark:border-slate-800 pb-3">
                <span class="text-[10px] font-bold text-slate-500 uppercase">{{ $t('admin.backup.restore.s2.overwrite') }}</span>
                <span class="text-[11px] font-black text-red-500 uppercase tracking-tight">142,042 files</span>
              </div>
              <div class="flex items-center justify-between border-b border-slate-50 dark:border-slate-800 pb-3">
                <span class="text-[10px] font-bold text-slate-500 uppercase">{{ $t('admin.backup.restore.s2.est') }}</span>
                <span class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">18m 30s</span>
              </div>
              <div class="flex items-center justify-between">
                <span class="text-[10px] font-bold text-slate-500 uppercase">{{ $t('admin.backup.restore.s2.downtime') }}</span>
                <span class="text-[11px] font-black text-blue-500 uppercase tracking-tight">{{ $t('admin.backup.restore.s2.required') }}</span>
              </div>
            </div>
            <div class="p-6 bg-slate-50 dark:bg-slate-900 rounded-2xl border border-slate-100 dark:border-slate-800 italic">
              <p class="text-[9px] font-bold text-slate-400 leading-relaxed">{{ $t('admin.backup.restore.s2.note') }}</p>
            </div>
          </div>

          <div class="pt-8 border-t border-slate-50 dark:border-slate-800 space-y-10">
            <label class="flex gap-4 cursor-pointer group">
              <div :class="`mt-1 w-6 h-6 rounded border-2 flex items-center justify-center transition-all shrink-0 ${confirmed ? 'bg-[#1E3A5F] border-[#1E3A5F]' : 'bg-white border-slate-200 group-hover:border-slate-400'}`" @click="confirmed = !confirmed">
                <LucideCheck v-if="confirmed" class="w-4 h-4 text-white" />
              </div>
              <span class="text-[10px] font-bold text-slate-500 leading-relaxed uppercase">{{ $t('admin.backup.restore.auth') }}</span>
            </label>
            <div class="space-y-4">
              <button :class="`w-full py-5 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40 hover:bg-[#152943] transition-all flex items-center justify-center gap-4 active:scale-95 ${!confirmed ? 'opacity-50 grayscale cursor-not-allowed' : ''}`">
                <LucideRotateCcw class="w-4 h-4" />
                {{ $t('admin.backup.restore.btn_restore') }}
              </button>
              <p class="text-[8px] font-black text-slate-300 text-center uppercase tracking-widest">Verification token will be required next</p>
            </div>
          </div>
        </div>

        <!-- Maintenance Window -->
        <div class="p-8 bg-amber-50 dark:bg-amber-900/10 rounded-lg border border-amber-100 dark:border-amber-800 flex gap-5">
          <LucideInfo class="w-6 h-6 text-amber-500 shrink-0" />
          <div class="space-y-1">
            <h5 class="text-[10px] font-black text-amber-600 uppercase tracking-widest">{{ $t('admin.backup.restore.maint.title') }}</h5>
            <p class="text-[9px] font-bold text-amber-500 uppercase leading-relaxed">{{ $t('admin.backup.restore.maint.desc') }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucidePlay, LucideClock, LucideChevronDown, LucideCloud, 
  LucideShieldCheck, LucideFilter, LucideChevronLeft, LucideChevronRight, 
  LucideGlobe, LucideCheckCircle2, LucideLock, LucideRotateCcw, 
  LucideCheck, LucideInfo 
} from 'lucide-vue-next'

const selectedScope = ref('full')
const confirmed = ref(false)

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
