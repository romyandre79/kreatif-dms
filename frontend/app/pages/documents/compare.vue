<template>
  <div class="max-w-7xl mx-auto space-y-10 pb-32" v-motion-fade>
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-1">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
          Version Detail & Compare <span class="text-slate-400">(F-31)</span>
        </h1>
        <p class="text-sm font-bold text-slate-500">
          Audit report for Technical Spec-v02 vs v03
        </p>
      </div>
      <div class="flex items-center gap-4">
        <button class="px-6 py-4 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-slate-600 dark:text-slate-300 rounded-2xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center gap-2 shadow-sm">
          <LucideDownload class="w-4 h-4" />
          Export PDF
        </button>
        <button class="px-8 py-4 bg-[#1E3A5F] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 active:scale-95">
          <LucideChevronLeft class="w-4 h-4" />
          Kembali ke Daftar Versi
        </button>
      </div>
    </div>

    <!-- Top Snapshot Cards -->
    <div class="grid grid-cols-1 md:grid-cols-12 gap-8">
      <!-- Version A -->
      <div class="md:col-span-4 lg:col-span-4">
        <div class="glass p-10 rounded-lg border-l-8 border-blue-500/30 space-y-8" v-motion-slide-visible-bottom>
          <div class="flex items-center justify-between">
            <span class="px-4 py-1.5 bg-blue-50 dark:bg-blue-900/30 text-blue-500 text-[10px] font-black uppercase tracking-widest rounded-lg">Version A (V2)</span>
            <span class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">24 May 2023 10:20 AM</span>
          </div>
          <div class="space-y-6">
            <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest flex items-center gap-3">
              <LucideRotateCcw class="w-4 h-4 text-blue-500" />
              Metadata Snapshot
            </h3>
            <div class="space-y-4">
              <div v-for="item in versionA" :key="item.label" class="flex justify-between items-center text-[11px]">
                <span class="font-bold text-slate-400 uppercase tracking-widest">{{ item.label }}</span>
                <span :class="`font-black uppercase tracking-tight ${item.color || 'text-[#1E3A5F] dark:text-white'}`">{{ item.value }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Version B -->
      <div class="md:col-span-4 lg:col-span-4">
        <div class="glass p-10 rounded-lg border-l-8 border-indigo-500/30 space-y-8" v-motion-slide-visible-bottom :delay="100">
          <div class="flex items-center justify-between">
            <span class="px-4 py-1.5 bg-indigo-50 dark:bg-indigo-900/30 text-indigo-500 text-[10px] font-black uppercase tracking-widest rounded-lg">Version B (V3)</span>
            <span class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">12 June 2023 14:45 PM</span>
          </div>
          <div class="space-y-6">
            <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest flex items-center gap-3">
              <LucideCheckCircle2 class="w-4 h-4 text-indigo-500" />
              Metadata Snapshot
            </h3>
            <div class="space-y-4">
              <div v-for="item in versionB" :key="item.label" class="flex justify-between items-center text-[11px]">
                <span class="font-bold text-slate-400 uppercase tracking-widest">{{ item.label }}</span>
                <span :class="`font-black uppercase tracking-tight ${item.color || 'text-[#1E3A5F] dark:text-white'}`">{{ item.value }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Change Approval Status -->
      <div class="md:col-span-4 lg:col-span-4">
        <div class="glass p-10 rounded-lg space-y-8 h-full" v-motion-slide-visible-bottom :delay="200">
          <h3 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em]">Change Approval Status</h3>
          <div class="flex items-center gap-5 p-4 bg-green-50/50 dark:bg-green-900/20 rounded-3xl border border-green-100 dark:border-green-800">
            <LucideCheckCircle2 class="w-8 h-8 text-green-500" />
            <div>
              <p class="text-[11px] font-black text-green-600 uppercase tracking-tight">Approved by Auditor</p>
              <p class="text-[9px] font-bold text-slate-500 uppercase tracking-widest">Setyo Wahono — 13 June 2023</p>
            </div>
          </div>
          <div class="space-y-3">
            <div class="flex justify-between items-center">
              <span class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Change Integrity</span>
              <span class="text-[9px] font-black text-green-500 uppercase tracking-widest">98% Match</span>
            </div>
            <div class="h-2 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
              <div class="h-full bg-green-500 w-[98%]"></div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content Grid -->
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Textual Diff Preview -->
      <div class="lg:col-span-8 space-y-8">
        <div class="glass rounded-[3rem] overflow-hidden" v-motion-slide-visible-bottom>
          <div class="px-10 py-8 border-b border-slate-50 dark:border-slate-800 flex items-center justify-between bg-slate-50/30 dark:bg-slate-900/30">
            <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest flex items-center gap-3">
              <LucideFileText class="w-5 h-5 text-blue-500" />
              Textual Diff Preview
            </h3>
            <div class="flex items-center gap-6">
              <div class="flex items-center gap-2">
                <div class="w-3 h-3 bg-green-200 dark:bg-green-900/50 rounded-sm"></div>
                <span class="text-[9px] font-black text-green-600 uppercase tracking-widest">Added</span>
              </div>
              <div class="flex items-center gap-2">
                <div class="w-3 h-3 bg-red-200 dark:bg-red-900/50 rounded-sm"></div>
                <span class="text-[9px] font-black text-red-600 uppercase tracking-widest">Removed</span>
              </div>
            </div>
          </div>
          <div class="p-10 space-y-10">
            <div class="space-y-6 font-mono text-sm leading-relaxed text-slate-600 dark:text-slate-400">
              <p>Section 1.2: Architecture Overview</p>
              <p>
                The current system architecture utilizes 
                <span class="text-red-500 line-through decoration-2">legacy REST endpoints</span> 
                <span class="bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-400 px-1 rounded-sm underline decoration-2">modern GraphQL federated schemas</span> 
                to handle data requests from the 
                <span class="bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-400 px-1 rounded-sm block mt-2">increased complexity of document relationships in Enterprise Tier.</span>
              </p>
              <p>Section 2.0: Security Protocol</p>
              <p>
                All document intake processes 
                <span class="text-red-500 line-through decoration-2">should be monitored by local admin</span> 
                <span class="bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-400 px-1 rounded-sm underline decoration-2">must be audited through the LDAP Sync module using multi-tenant keys</span>.
                Encryption at rest is 
                <span class="bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-400 px-1 rounded-sm">now enforced using AES-256 standards</span>.
              </p>
            </div>

            <!-- Editor Notes -->
            <div class="p-8 border-2 border-dashed border-slate-200 dark:border-slate-800 rounded-3xl bg-slate-50/50 dark:bg-slate-900/50 space-y-4">
              <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Editor Notes</p>
              <p class="text-xs font-bold text-slate-500 italic leading-relaxed">
                "Updated the security protocols to match the new LDAP synchronization policy established in Q2. Also expanded the architecture section to include the new federated schema details." — Ani Wijaya
              </p>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Sidebar -->
      <div class="lg:col-span-4 space-y-10">
        <!-- Action History Log -->
        <div class="glass p-10 rounded-[3rem] space-y-10" v-motion-slide-visible-bottom>
          <div class="flex items-center justify-between border-b border-slate-50 dark:border-slate-800 pb-6">
            <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">Action History Log</h3>
            <LucideHistory class="w-4 h-4 text-slate-400" />
          </div>

          <div class="space-y-8 relative">
            <!-- Timeline Line -->
            <div class="absolute left-[15px] top-0 bottom-0 w-0.5 bg-slate-100 dark:bg-slate-800"></div>

            <div v-for="log in history" :key="log.time" class="flex items-start gap-6 relative z-10">
              <div :class="`w-8 h-8 rounded-full flex items-center justify-center border-2 border-white dark:border-slate-900 shadow-sm ${log.bg}`">
                <component :is="log.icon" :class="`w-3.5 h-3.5 ${log.color}`" />
              </div>
              <div class="space-y-0.5">
                <p class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ log.action }}</p>
                <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ log.user }} — {{ log.time }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Scan CTA -->
        <div class="bg-gradient-to-br from-[#1E3A5F] to-[#152943] p-10 rounded-[3rem] space-y-6 shadow-2xl shadow-blue-900/30 relative overflow-hidden group" v-motion-slide-visible-bottom :delay="300">
          <div class="space-y-2 relative z-10">
            <h3 class="text-lg font-black text-white uppercase tracking-tight">Butuh Dokumen Baru?</h3>
            <p class="text-[11px] font-bold text-blue-200 leading-relaxed uppercase">Pindai dokumen fisik Anda langsung ke sistem repository untuk perbandingan versi instan.</p>
          </div>
          <button class="w-full py-4 bg-white/10 hover:bg-white/20 border border-white/20 text-white rounded-2xl text-[10px] font-black uppercase tracking-widest flex items-center justify-center gap-3 transition-all relative z-10">
            <LucideScan class="w-4 h-4" />
            Scan Dokumen Baru
          </button>
          <LucideFileText class="absolute -right-6 -bottom-6 w-32 h-32 text-white/5 -rotate-12 group-hover:scale-110 transition-transform" />
        </div>
      </div>
    </div>

    <!-- Bottom Difference Summary -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6" v-motion-slide-visible-bottom>
      <div class="glass p-8 rounded-lg flex flex-col justify-center gap-2">
        <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Text Growth</p>
        <p class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tighter uppercase">+1,240 <span class="text-xs text-slate-400">words</span></p>
      </div>
      <div class="glass p-8 rounded-lg flex flex-col justify-center gap-2">
        <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Deletions</p>
        <p class="text-2xl font-black text-red-500 tracking-tighter uppercase">45 <span class="text-xs text-slate-400">lines</span></p>
      </div>
      <div class="glass p-8 rounded-lg flex flex-col justify-center gap-2">
        <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Additions</p>
        <p class="text-2xl font-black text-green-500 tracking-tighter uppercase">112 <span class="text-xs text-slate-400">lines</span></p>
      </div>
      <div class="glass p-8 rounded-lg flex flex-col justify-center gap-2">
        <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Complexity</p>
        <p class="text-2xl font-black text-blue-500 tracking-tighter uppercase">High <span class="text-xs text-slate-400">Risk</span></p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideDownload, LucideRotateCcw, LucideCheckCircle2, LucideChevronLeft, 
  LucideFileText, LucideHistory, LucideScan, LucideEye, LucideRepeat, 
  LucideFileUp, LucidePlusCircle
} from 'lucide-vue-next'

const versionA = [
  { label: 'Author', value: 'Budi Santoso' },
  { label: 'Classification', value: 'INTERNAL', color: 'text-blue-500' },
  { label: 'Pages', value: '12' },
  { label: 'File Size', value: '4.2 MB' }
]

const versionB = [
  { label: 'Author', value: 'Ani Wijaya' },
  { label: 'Classification', value: 'CONFIDENTIAL', color: 'text-red-500' },
  { label: 'Pages', value: '15' },
  { label: 'File Size', value: '5.1 MB' }
]

const history = [
  { action: 'Viewed by Admin', user: 'John Doe', time: '2 mins ago', icon: LucideEye, bg: 'bg-blue-50', color: 'text-blue-500' },
  { action: 'Compared Versions', user: 'John Doe', time: '5 mins ago', icon: LucideRepeat, bg: 'bg-indigo-50', color: 'text-indigo-500' },
  { action: 'Exported Comparison Report', user: 'Budi Santoso', time: '2 hours ago', icon: LucideFileUp, bg: 'bg-slate-100', color: 'text-slate-500' },
  { action: 'Draft Created', user: 'Ani Wijaya', time: 'Yesterday', icon: LucidePlusCircle, bg: 'bg-slate-50', color: 'text-slate-400' }
]

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
@reference "tailwindcss";
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/5 shadow-sm;
}
</style>

