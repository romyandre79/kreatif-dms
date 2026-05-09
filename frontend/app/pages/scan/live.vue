<template>
  <div class="max-w-7xl mx-auto space-y-8 pb-20" v-motion-fade>
    <!-- Breadcrumbs -->
    <nav class="flex items-center gap-2 text-[10px] font-black text-slate-400 uppercase tracking-widest">
      <span>Document Intake</span>
      <LucideChevronRight class="w-3 h-3" />
      <span class="text-[#1E3A5F] dark:text-white uppercase tracking-tighter">Live Scan Session</span>
    </nav>

    <!-- Header -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-1">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
          Screen S2 - Live Scan Preview
        </h1>
        <p class="text-sm font-bold text-slate-500">
          Browser Intake F-11 Engine v4.0
        </p>
      </div>
      <div class="glass px-6 py-4 rounded-2xl flex items-center gap-4">
        <div class="text-right">
          <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Captured to current browser session</p>
          <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">Secure Tunnel: Active</p>
        </div>
        <div class="w-2.5 h-2.5 rounded-full bg-green-500 animate-pulse"></div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8 items-start">
      <!-- Left Sidebar: Thumbnails -->
      <div class="lg:col-span-3 space-y-6">
        <div class="flex items-center justify-between px-2">
          <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em]">Thumbnails (12)</h3>
          <LucideFilter class="w-4 h-4 text-slate-400 cursor-pointer hover:text-blue-500 transition-colors" />
        </div>

        <div class="space-y-6 overflow-y-auto max-h-[800px] pr-2 custom-scrollbar">
          <!-- Scanning Progress Card -->
          <div class="p-6 bg-slate-50 dark:bg-slate-900 border-2 border-blue-500/20 rounded-[2rem] space-y-3">
            <p class="text-[10px] font-black text-slate-500 uppercase tracking-widest text-center">Scanning page 4 of 12</p>
            <div class="h-2 bg-slate-200 dark:bg-slate-800 rounded-full overflow-hidden">
              <div class="h-full bg-[#1E3A5F] w-[33%] animate-pulse"></div>
            </div>
          </div>

          <!-- Thumbnail List -->
          <div v-for="t in thumbnails" :key="t.page" 
            class="relative group rounded-lg p-4 transition-all duration-300 border-4 overflow-hidden"
            :class="t.status === 'active' ? 'border-[#1E3A5F] bg-slate-50 dark:bg-slate-900 shadow-xl shadow-blue-900/10' : 'border-transparent hover:border-slate-100 dark:hover:border-slate-800'"
          >
            <div class="aspect-[3/4] bg-white dark:bg-slate-800 rounded-[1.5rem] overflow-hidden shadow-sm flex items-center justify-center p-2">
              <img :src="t.src" class="w-full h-full object-cover rounded-xl" :class="t.status === 'error' ? 'opacity-50' : ''" />
            </div>
            <div class="mt-4 flex items-center justify-between px-2">
              <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Page {{ t.page }}</span>
              <div v-if="t.status === 'success'" class="w-4 h-4 rounded-full bg-green-500 flex items-center justify-center text-white">
                <LucideCheck class="w-2.5 h-2.5" />
              </div>
              <div v-else-if="t.status === 'active'" class="px-2 py-0.5 bg-blue-50 dark:bg-blue-900 text-[8px] font-black text-[#1E3A5F] dark:text-blue-400 uppercase tracking-tighter rounded">
                Active
              </div>
              <div v-else-if="t.status === 'error'" class="w-4 h-4 rounded-full bg-red-500 flex items-center justify-center text-white">
                <LucideAlertCircle class="w-3 h-3" />
              </div>
            </div>
            <!-- Error Border for Page 03 -->
            <div v-if="t.status === 'error'" class="absolute inset-0 border-2 border-red-500/30 rounded-lg pointer-events-none"></div>
          </div>
        </div>
      </div>

      <!-- Main Preview Area -->
      <div class="lg:col-span-6 space-y-6">
        <div class="glass rounded-[3rem] overflow-hidden flex flex-col h-[850px]" v-motion-slide-visible-bottom>
          <!-- Toolbar -->
          <div class="px-8 py-6 border-b border-slate-50 dark:border-slate-800 flex items-center justify-between bg-slate-50/30 dark:bg-slate-900/30">
            <div class="flex items-center gap-6">
              <div class="flex items-center gap-4 bg-white dark:bg-slate-800 rounded-xl p-1 shadow-sm border border-slate-100 dark:border-slate-700">
                <button class="p-2 text-slate-400 hover:text-blue-500 transition-colors"><LucideZoomIn class="w-4 h-4" /></button>
                <button class="p-2 text-slate-400 hover:text-blue-500 transition-colors"><LucideZoomOut class="w-4 h-4" /></button>
              </div>
              <button class="p-3 bg-white dark:bg-slate-800 rounded-xl shadow-sm border border-slate-100 dark:border-slate-700 text-slate-400 hover:text-blue-500"><LucideRotateCw class="w-4 h-4" /></button>
              <div class="flex items-center gap-3">
                <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Page 04</span>
                <span class="px-2 py-1 bg-slate-100 dark:bg-slate-800 rounded text-[9px] font-black text-slate-500 uppercase tracking-widest">100%</span>
              </div>
            </div>
            <div class="flex items-center gap-4">
              <button class="flex items-center gap-2 px-5 py-2.5 bg-red-50 dark:bg-red-900/20 text-red-500 text-[10px] font-black uppercase tracking-widest rounded-xl hover:bg-red-500 hover:text-white transition-all">
                <LucideTrash2 class="w-3.5 h-3.5" />
                Remove
              </button>
              <button class="flex items-center gap-2 px-5 py-2.5 bg-blue-50 dark:bg-blue-900/20 text-blue-500 text-[10px] font-black uppercase tracking-widest rounded-xl hover:bg-blue-500 hover:text-white transition-all">
                <LucideRefreshCcw class="w-3.5 h-3.5" />
                Retry
              </button>
            </div>
          </div>

          <!-- Image Area -->
          <div class="flex-grow bg-slate-100/50 dark:bg-slate-900/50 p-10 flex items-center justify-center relative overflow-hidden">
            <div class="w-full h-full bg-white dark:bg-slate-800 shadow-2xl rounded-sm flex items-center justify-center p-8 relative">
              <img src="https://images.unsplash.com/photo-1586769852836-bc069f19e1b6?auto=format&fit=crop&q=80&w=800" class="max-w-full max-h-full object-contain" />
              <!-- Scanning Line Overlay -->
              <div class="absolute top-1/2 left-0 w-full h-[1px] bg-blue-500 shadow-[0_0_15px_blue] animate-scan"></div>
            </div>
          </div>

          <!-- Bottom Status -->
          <div class="p-6 bg-slate-50/50 dark:bg-slate-900/50 text-center">
            <p class="text-[9px] font-black text-slate-400 uppercase tracking-[0.2em]">
              Resolusi Deteksi: 300 DPI • Mode: Warna • Format: PDF/A-1
            </p>
          </div>
        </div>
      </div>

      <!-- Right Sidebar: Metadata & Actions -->
      <div class="lg:col-span-3 space-y-8">
        <!-- Metadata Sesi -->
        <div class="glass p-8 rounded-lg space-y-8" v-motion-slide-visible-bottom>
          <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em] border-b border-slate-50 dark:border-slate-800 pb-6">Metadata Sesi</h3>
          <div class="space-y-6">
            <div class="space-y-1">
              <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">Session ID</p>
              <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight font-mono">DMS-2023-9981</p>
            </div>
            <div class="space-y-1">
              <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">Waktu Mulai</p>
              <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">14:20:05 WIB</p>
            </div>
            <div class="space-y-1">
              <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">Total Ukuran</p>
              <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">4.2 MB</p>
            </div>
          </div>
        </div>

        <!-- Action Panel -->
        <div class="space-y-4" v-motion-slide-visible-bottom :delay="100">
          <button class="w-full py-6 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-[2rem] shadow-2xl shadow-blue-900/20 transition-all active:scale-95 group">
            <p class="text-xs font-black uppercase tracking-[0.2em] mb-1">Gunakan Hasil Scan</p>
            <p class="text-[9px] font-bold text-blue-300 uppercase tracking-widest">Lanjut ke Distribusi Digital</p>
          </button>
          
          <button class="w-full py-4 border-2 border-slate-200 dark:border-slate-800 text-[#1E3A5F] dark:text-slate-300 rounded-[1.5rem] text-[10px] font-black uppercase tracking-widest hover:bg-slate-50 transition-all">
            Simpan Draft
          </button>

          <button class="w-full py-4 text-slate-400 hover:text-red-500 text-[10px] font-black uppercase tracking-widest transition-colors">
            Batalkan Sesi
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideChevronRight, LucideFilter, LucideSearch, LucideRotateCw, 
  LucideTrash2, LucideRefreshCcw, LucideZoomIn, LucideZoomOut, 
  LucideCheck, LucideAlertCircle, LucideClock
} from 'lucide-vue-next'

const thumbnails = [
  { page: '01', src: 'https://images.unsplash.com/photo-1586769852836-bc069f19e1b6?auto=format&fit=crop&q=80&w=200', status: 'success' },
  { page: '02', src: 'https://images.unsplash.com/photo-1568667256549-094345857637?auto=format&fit=crop&q=80&w=200', status: 'active' },
  { page: '03', src: 'https://images.unsplash.com/photo-1517842645767-c639042777db?auto=format&fit=crop&q=80&w=200', status: 'error' }
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

.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  @apply bg-transparent;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  @apply bg-slate-200 dark:bg-slate-800 rounded-full;
}

@keyframes scan {
  0% { top: 0; }
  100% { top: 100%; }
}

.animate-scan {
  animation: scan 3s linear infinite;
}
</style>

