<template>
  <div class="max-w-6xl mx-auto space-y-10 pb-20" v-motion-fade>
    <!-- Breadcrumbs -->
    <nav class="flex items-center gap-2 text-[10px] font-black text-slate-400 uppercase tracking-widest">
      <span>Document Intake</span>
      <LucideChevronRight class="w-3 h-3" />
      <span class="text-[#1E3A5F] dark:text-white uppercase tracking-tighter">Scan from Web</span>
    </nav>

    <!-- Header -->
    <div class="space-y-1">
      <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
        Scan-to-Web Direct Launcher <span class="text-slate-400">(F-11)</span>
      </h1>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8 items-stretch">
      <!-- Left Sidebar: Scanner Connection -->
      <div class="lg:col-span-4 h-full">
        <div class="glass p-10 rounded-[3rem] space-y-10 h-full flex flex-col" v-motion-slide-visible-bottom>
          <div class="flex items-center justify-between">
            <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">Scanner Connection</h3>
            <div class="flex items-center gap-2">
              <div class="w-2 h-2 rounded-full bg-green-500 animate-pulse"></div>
              <span class="text-[9px] font-black text-green-500 uppercase tracking-widest">Connected</span>
            </div>
          </div>

          <!-- Device Info Card -->
          <div class="p-8 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-[2rem] space-y-6">
            <div class="w-14 h-14 rounded-2xl bg-white dark:bg-slate-800 flex items-center justify-center text-[#1E3A5F] dark:text-blue-400 shadow-sm border border-slate-50 dark:border-slate-700">
              <LucidePrinter class="w-7 h-7" />
            </div>
            <div class="space-y-1">
              <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">Device Name</p>
              <p class="text-xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">HP Digital Sender Flow</p>
              <div class="flex items-center gap-2 text-[10px] font-black text-slate-400">
                <LucideNetwork class="w-3 h-3" />
                192.168.1.104
              </div>
            </div>
          </div>

          <div class="space-y-6 flex-grow">
            <div class="flex items-center justify-between">
              <span class="text-[11px] font-bold text-slate-400 uppercase tracking-widest">Driver Status</span>
              <span class="text-[11px] font-black text-[#1E3A5F] dark:text-white">v2.4.1 (Latest)</span>
            </div>
            <div class="flex items-center justify-between border-b border-slate-50 dark:border-slate-800 pb-6">
              <span class="text-[11px] font-bold text-slate-400 uppercase tracking-widest">Paper Feed</span>
              <span class="text-[11px] font-black text-green-500 uppercase">Ready</span>
            </div>
          </div>

          <p class="text-[10px] font-bold text-slate-400 leading-relaxed italic">
            Note: Scanner hardware is locked for exclusive web session.
          </p>
        </div>
      </div>

      <!-- Right Panel: Configuration -->
      <div class="lg:col-span-8 h-full">
        <div class="glass p-10 rounded-[3rem] space-y-10 h-full flex flex-col" v-motion-slide-visible-bottom :delay="100">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">Scan Configuration</h3>

          <div class="grid grid-cols-2 gap-x-10 gap-y-8">
            <!-- Color Mode -->
            <div class="space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-1">Color Mode</label>
              <div class="relative">
                <select class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-[11px] font-black text-[#1E3A5F] dark:text-white appearance-none outline-none focus:ring-2 focus:ring-blue-500/20 transition-all uppercase tracking-tight">
                  <option>Auto Detect (Recommended)</option>
                  <option>Color</option>
                  <option>Grayscale</option>
                  <option>Black & White</option>
                </select>
                <LucideChevronDown class="absolute right-6 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
              </div>
            </div>

            <!-- Resolution -->
            <div class="space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-1">Resolution (DPI)</label>
              <div class="relative">
                <select class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-[11px] font-black text-[#1E3A5F] dark:text-white appearance-none outline-none focus:ring-2 focus:ring-blue-500/20 transition-all uppercase tracking-tight">
                  <option>300 DPI - Standard</option>
                  <option>150 DPI - Draft</option>
                  <option>600 DPI - High Quality</option>
                </select>
                <LucideChevronDown class="absolute right-6 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
              </div>
            </div>

            <!-- Duplex -->
            <div class="space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-1">Duplex Scanning</label>
              <div class="flex p-1 bg-slate-100 dark:bg-slate-900 rounded-2xl gap-1">
                <button 
                  @click="duplex = true"
                  class="flex-grow py-3 rounded-xl text-[10px] font-black uppercase tracking-widest transition-all"
                  :class="duplex ? 'bg-[#1E3A5F] text-white shadow-lg' : 'text-slate-400 hover:text-slate-600'"
                >
                  Double Sided
                </button>
                <button 
                  @click="duplex = false"
                  class="flex-grow py-3 rounded-xl text-[10px] font-black uppercase tracking-widest transition-all"
                  :class="!duplex ? 'bg-[#1E3A5F] text-white shadow-lg' : 'text-slate-400 hover:text-slate-600'"
                >
                  Single Sided
                </button>
              </div>
            </div>

            <!-- Page Size -->
            <div class="space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-1">Page Size</label>
              <div class="relative">
                <select class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-[11px] font-black text-[#1E3A5F] dark:text-white appearance-none outline-none focus:ring-2 focus:ring-blue-500/20 transition-all uppercase tracking-tight">
                  <option>A4 (210 x 297 mm)</option>
                  <option>F4 / Folio</option>
                  <option>Letter</option>
                  <option>Legal</option>
                </select>
                <LucideChevronDown class="absolute right-6 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
              </div>
            </div>
          </div>

          <!-- Warning Box -->
          <div class="p-8 bg-red-50/50 dark:bg-red-900/10 border-l-4 border-red-500 rounded-2xl flex gap-6">
            <LucideAlertTriangle class="w-6 h-6 text-red-500 flex-shrink-0" />
            <div class="space-y-1">
              <h4 class="text-[11px] font-black text-red-600 uppercase tracking-widest">Warning</h4>
              <p class="text-[11px] font-bold text-slate-500 leading-relaxed">
                Hasil scan masuk langsung ke browser, tanpa simpan manual ke PC. Pastikan koneksi internet stabil selama proses transfer data.
              </p>
            </div>
          </div>

          <!-- Action Footer -->
          <div class="pt-6 border-t border-slate-50 dark:border-slate-800 flex items-center justify-between mt-auto">
            <div class="flex items-center gap-3">
              <LucideInfo class="w-4 h-4 text-slate-400" />
              <p class="text-[10px] font-bold text-slate-500 uppercase tracking-tight">Klik Mulai Scan untuk trigger scanner fisik.</p>
            </div>
            <button class="px-10 py-4 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 transition-all flex items-center gap-3 active:scale-95">
              <LucideScan class="w-4 h-4" />
              Mulai Scan
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Idle State Footer Area -->
    <div class="p-20 bg-slate-50/30 dark:bg-slate-900/30 rounded-[4rem] border-2 border-dashed border-slate-100 dark:border-slate-800 flex flex-col items-center justify-center space-y-6 opacity-60 hover:opacity-100 transition-opacity" v-motion-slide-visible-bottom :delay="200">
      <div class="w-20 h-20 rounded-[2rem] bg-white dark:bg-slate-800 flex items-center justify-center text-slate-300 shadow-sm">
        <LucideCpu class="w-10 h-10" />
      </div>
      <div class="text-center space-y-1">
        <p class="text-sm font-black text-[#1E3A5F] dark:text-slate-400 uppercase tracking-[0.2em]">Scanner is idle. Ready for direct intake.</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideChevronRight, LucidePrinter, LucideNetwork, LucideChevronDown, 
  LucideAlertTriangle, LucideInfo, LucideScan, LucideCpu
} from 'lucide-vue-next'

const duplex = ref(true)

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

