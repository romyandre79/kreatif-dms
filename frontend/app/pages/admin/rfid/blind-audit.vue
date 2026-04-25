<template>
  <div class="min-h-screen bg-[#F8FAFC] flex justify-center p-4">
    <div class="w-full max-w-md bg-white shadow-2xl rounded-[3rem] overflow-hidden flex flex-col relative border border-slate-100" v-motion-fade>
      
      <!-- Header -->
      <header class="p-8 border-b border-slate-50 flex items-center justify-between">
        <div class="space-y-1">
          <h1 class="text-xl font-black text-[#1E3A5F] tracking-tight uppercase">PT Akiradata DMS</h1>
          <p class="text-[10px] font-black text-blue-500 uppercase tracking-widest">Mission: F-22 Blind Audit</p>
        </div>
        <div class="flex items-center gap-4">
          <div class="text-right">
            <div class="flex items-center gap-1.5 text-[9px] font-black text-green-500 uppercase tracking-widest">
              <LucideWifi class="w-3 h-3" />
              Connected
            </div>
            <p class="text-[10px] font-mono font-black text-slate-400">00:14:22</p>
          </div>
          <button class="p-3 bg-slate-50 rounded-2xl text-slate-400 hover:text-[#1E3A5F] transition-colors">
            <LucideSettings class="w-5 h-5" />
          </button>
        </div>
      </header>

      <!-- Auditor Info Card -->
      <div class="px-8 pt-8">
        <div class="p-6 bg-slate-50 rounded-[2rem] border-l-4 border-[#1E3A5F] flex items-center justify-between">
          <div class="flex items-center gap-4">
            <img src="https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?auto=format&fit=crop&q=80&w=100" class="w-12 h-12 rounded-2xl object-cover shadow-sm" />
            <div>
              <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">Auditor Active</p>
              <p class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">Supriyadi Pratama</p>
            </div>
          </div>
          <div class="text-right">
            <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">Device ID</p>
            <p class="text-[10px] font-black text-[#1E3A5F] font-mono">RFID-X880</p>
          </div>
        </div>
      </div>

      <!-- Main Progress Circle -->
      <div class="p-10 flex flex-col items-center space-y-10">
        <div class="relative w-64 h-64 flex items-center justify-center">
          <!-- Circular Progress SVG -->
          <svg class="w-full h-full -rotate-90">
            <circle cx="128" cy="128" r="110" stroke="currentColor" stroke-width="12" fill="transparent" class="text-slate-100" />
            <circle cx="128" cy="128" r="110" stroke="currentColor" stroke-width="12" fill="transparent" stroke-dasharray="690" stroke-dashoffset="545" class="text-[#1E3A5F]" />
          </svg>
          <div class="absolute inset-0 flex flex-col items-center justify-center space-y-1">
            <span class="text-6xl font-black text-[#1E3A5F] tracking-tighter">42</span>
            <span class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">Scanned Items</span>
          </div>
        </div>

        <div class="w-full grid grid-cols-2 gap-8 border-t border-slate-50 pt-8">
          <div class="text-center space-y-1">
            <p class="text-2xl font-black text-[#1E3A5F]">158</p>
            <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Remaining Target</p>
          </div>
          <div class="text-center space-y-1">
            <p class="text-2xl font-black text-[#1E3A5F]">21%</p>
            <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Completion</p>
          </div>
        </div>
      </div>

      <!-- Active Location Card -->
      <div class="px-8 mb-8">
        <div class="p-8 bg-blue-50 rounded-[2.5rem] flex items-center justify-between group cursor-pointer hover:bg-blue-100 transition-all">
          <div class="flex items-center gap-6">
            <div class="w-14 h-14 rounded-2xl bg-[#1E3A5F] flex items-center justify-center text-white shadow-lg">
              <LucideScanQrCode class="w-7 h-7" />
            </div>
            <div class="space-y-1">
              <p class="text-[10px] font-black text-blue-600 uppercase tracking-widest">Active Location (Rack ID)</p>
              <p class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">Scan Rack QR first</p>
            </div>
          </div>
          <LucideRefreshCw class="w-5 h-5 text-blue-400 group-hover:rotate-180 transition-transform duration-700" />
        </div>
      </div>

      <!-- Recent Scans -->
      <div class="px-8 pb-40 space-y-6">
        <div class="flex items-center justify-between px-2">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">Recent Scans (Blind Mode)</h3>
          <button class="text-[9px] font-black text-blue-500 uppercase tracking-widest">Live Feed</button>
        </div>

        <div class="space-y-3">
          <div v-for="scan in recentScans" :key="scan.tag" class="p-5 bg-white border border-slate-100 rounded-2xl flex items-center justify-between shadow-sm border-l-4" :class="scan.border">
            <div class="space-y-1">
              <p class="text-xs font-black text-[#1E3A5F] font-mono tracking-tighter">{{ scan.tag }}</p>
              <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ scan.time }} • {{ scan.shelf }}</p>
            </div>
            <span :class="`px-3 py-1 rounded-lg text-[8px] font-black uppercase tracking-widest ${scan.badge}`">
              {{ scan.status }}
            </span>
          </div>
        </div>
      </div>

      <!-- Floating Bottom Panel -->
      <div class="absolute bottom-0 left-0 w-full p-6 bg-white/80 backdrop-blur-xl border-t border-slate-50 flex flex-col gap-6">
        <div class="flex gap-4">
          <button class="flex-grow flex items-center justify-center gap-3 py-6 bg-slate-100 rounded-[1.5rem] text-[10px] font-black text-[#1E3A5F] uppercase tracking-widest">
            <LucideMapPin class="w-5 h-5" />
            Scan Rack
          </button>
          <button class="flex-grow flex items-center justify-center gap-3 py-6 bg-blue-50 rounded-[1.5rem] text-[10px] font-black text-blue-200 uppercase tracking-widest cursor-not-allowed">
            <LucideRss class="w-5 h-5" />
            Scan Item
          </button>
        </div>

        <!-- Tab Bar -->
        <nav class="flex items-center justify-between px-4 pb-2">
          <button class="flex flex-col items-center gap-1 text-blue-500">
            <div class="p-2.5 bg-blue-50 rounded-xl"><LucideRss class="w-5 h-5" /></div>
            <span class="text-[8px] font-black uppercase tracking-widest">Scan</span>
          </button>
          <button class="flex flex-col items-center gap-1 text-slate-400 hover:text-[#1E3A5F] transition-colors">
            <LucidePause class="w-5 h-5" />
            <span class="text-[8px] font-black uppercase tracking-widest">Pause</span>
          </button>
          <button class="flex flex-col items-center gap-1 text-slate-400 hover:text-[#1E3A5F] transition-colors">
            <LucideAlertTriangle class="w-5 h-5" />
            <span class="text-[8px] font-black uppercase tracking-widest">Alerts</span>
          </button>
          <button class="flex flex-col items-center gap-1 text-slate-400 hover:text-[#1E3A5F] transition-colors">
            <LucideCheckCircle2 class="w-5 h-5" />
            <span class="text-[8px] font-black uppercase tracking-widest">Finish</span>
          </button>
        </nav>
      </div>

    </div>
  </div>
</template>

<script setup>
import { 
  LucideSettings, LucideWifi, LucideScanQrCode, LucideRefreshCw, 
  LucideMapPin, LucideRss, LucidePause, LucideAlertTriangle, LucideCheckCircle2 
} from 'lucide-vue-next'

const recentScans = [
  { tag: 'TAG-8829-XJ', time: '14:21:05', shelf: 'Shelf A-2', status: 'Match', badge: 'bg-green-50 text-green-500', border: 'border-green-500' },
  { tag: 'TAG-1044-BY', time: '14:19:30', shelf: 'Shelf B-1', status: 'Missing', badge: 'bg-amber-50 text-amber-500', border: 'border-amber-500' },
  { tag: 'TAG-0012-ZZ', time: '14:10:12', shelf: 'Bulk Area', status: 'Extra', badge: 'bg-blue-50 text-blue-600', border: 'border-blue-500' }
]

definePageMeta({
  layout: 'none'
})
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  @apply bg-transparent;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  @apply bg-slate-200 dark:bg-slate-800 rounded-full;
}
</style>
