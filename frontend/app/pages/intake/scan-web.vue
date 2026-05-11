<template>
  <div class="space-y-8 pb-20" v-motion-fade>
    <!-- Breadcrumb & Header -->
    <div class="space-y-4">
      <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.scan_web.breadcrumb') }}</p>
      <div class="flex items-center gap-4">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('intake.scan_web.title') }}</h1>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
      <!-- Left Column: Scanner Connection -->
      <div class="lg:col-span-4 glass p-10 rounded-lg space-y-10 shadow-sm border border-slate-50 dark:border-slate-800">
        <div class="flex items-center justify-between">
          <h3 class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('intake.scan_web.connection.title') }}</h3>
          <span class="flex items-center gap-2 px-3 py-1 bg-green-50 text-green-500 rounded-full text-[8px] font-black uppercase tracking-widest border border-green-100">
            <span class="w-1.5 h-1.5 bg-green-500 rounded-full animate-pulse"></span>
            {{ $t('intake.scan_web.connection.connected') }}
          </span>
        </div>

        <div class="p-8 bg-slate-50 dark:bg-slate-900 rounded-lg border border-slate-100 dark:border-slate-800 space-y-6">
          <div class="flex items-start gap-6">
            <div class="w-16 h-16 rounded-2xl bg-white dark:bg-slate-800 flex items-center justify-center text-[#1E3A5F] dark:text-white shadow-sm border border-slate-100 dark:border-slate-800">
              <LucidePrinter class="w-8 h-8" />
            </div>
            <div class="space-y-1">
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('intake.scan_web.connection.device') }}</p>
              <h4 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">HP Digital Sender Flow</h4>
              <div class="flex items-center gap-2 text-[10px] font-bold text-slate-400">
                <LucideNetwork class="w-3.5 h-3.5" />
                192.168.1.104
              </div>
            </div>
          </div>
        </div>

        <div class="space-y-6 px-2">
          <div class="flex items-center justify-between text-[11px] font-bold uppercase">
            <span class="text-slate-400">{{ $t('intake.scan_web.connection.driver') }}</span>
            <span class="text-[#1E3A5F] dark:text-white font-black">v2.4.1 (Latest)</span>
          </div>
          <div class="flex items-center justify-between text-[11px] font-bold uppercase">
            <span class="text-slate-400">{{ $t('intake.scan_web.connection.paper') }}</span>
            <span class="text-green-500 font-black">{{ $t('intake.scan_web.connection.ready') }}</span>
          </div>
        </div>

        <div class="pt-10 border-t border-slate-50 dark:border-slate-800">
          <p class="text-[10px] font-bold text-slate-400 leading-relaxed italic">
            {{ $t('intake.scan_web.connection.note') }}
          </p>
        </div>
      </div>

      <!-- Right Column: Scan Configuration -->
      <div class="lg:col-span-8 glass p-10 rounded-lg space-y-10 shadow-sm border border-slate-50 dark:border-slate-800">
        <h3 class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('intake.scan_web.config.title') }}</h3>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-10">
          <!-- Color Mode -->
          <div class="space-y-3">
            <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.scan_web.config.color') }}</label>
            <div class="relative group">
              <select class="w-full h-14 pl-6 pr-12 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold outline-none appearance-none cursor-pointer focus:border-[#1E3A5F] transition-all">
                <option>Auto Detect (Recommended)</option>
                <option>Full Color</option>
                <option>Grayscale</option>
                <option>Black & White</option>
              </select>
              <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300 pointer-events-none group-focus-within:text-[#1E3A5F]" />
            </div>
          </div>

          <!-- Resolution -->
          <div class="space-y-3">
            <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.scan_web.config.resolution') }}</label>
            <div class="relative group">
              <select class="w-full h-14 pl-6 pr-12 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold outline-none appearance-none cursor-pointer focus:border-[#1E3A5F] transition-all">
                <option>300 DPI - Standard</option>
                <option>600 DPI - High Quality</option>
                <option>1200 DPI - Photo</option>
              </select>
              <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300 pointer-events-none group-focus-within:text-[#1E3A5F]" />
            </div>
          </div>

          <!-- Duplex -->
          <div class="space-y-3">
            <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.scan_web.config.duplex') }}</label>
            <div class="flex p-1.5 bg-slate-50 dark:bg-slate-900 rounded-2xl border border-slate-100 dark:border-slate-800 h-14">
              <button 
                @click="isDuplex = true"
                :class="`flex-1 rounded-xl text-[10px] font-black uppercase tracking-widest transition-all ${isDuplex ? 'bg-[#1E3A5F] text-white shadow-lg' : 'text-slate-400 hover:text-slate-600'}`"
              >
                {{ $t('intake.scan_web.config.double') }}
              </button>
              <button 
                @click="isDuplex = false"
                :class="`flex-1 rounded-xl text-[10px] font-black uppercase tracking-widest transition-all ${!isDuplex ? 'bg-[#1E3A5F] text-white shadow-lg' : 'text-slate-400 hover:text-slate-600'}`"
              >
                {{ $t('intake.scan_web.config.single') }}
              </button>
            </div>
          </div>

          <!-- Page Size -->
          <div class="space-y-3">
            <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.scan_web.config.size') }}</label>
            <div class="relative group">
              <select class="w-full h-14 pl-6 pr-12 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold outline-none appearance-none cursor-pointer focus:border-[#1E3A5F] transition-all">
                <option>A4 (210 x 297 mm)</option>
                <option>Letter (8.5 x 11 in)</option>
                <option>Legal (8.5 x 14 in)</option>
              </select>
              <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300 pointer-events-none group-focus-within:text-[#1E3A5F]" />
            </div>
          </div>
        </div>

        <!-- Warning Alert -->
        <div class="p-8 bg-red-50 border-l-4 border-red-500 rounded-2xl flex gap-6 items-start">
          <LucideAlertTriangle class="w-6 h-6 text-red-500 shrink-0" />
          <div class="space-y-1">
            <h5 class="text-[11px] font-black text-red-600 uppercase tracking-widest">Warning</h5>
            <p class="text-[11px] font-bold text-red-400 leading-relaxed italic">
              {{ $t('intake.scan_web.config.warning') }}
            </p>
          </div>
        </div>

        <!-- Action Footer -->
        <div class="pt-10 border-t border-slate-50 dark:border-slate-800 flex flex-col md:flex-row md:items-center justify-between gap-6">
          <div class="flex items-center gap-4 text-slate-400">
            <LucideInfo class="w-5 h-5" />
            <p class="text-[11px] font-bold uppercase tracking-tight">{{ $t('intake.scan_web.config.footer') }}</p>
          </div>
          <button class="px-12 py-4 bg-[#1E3A5F] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-2xl shadow-blue-900/30 hover:bg-[#152943] transition-all flex items-center gap-4 active:scale-95 group">
            <LucideScan class="w-5 h-5 group-hover:rotate-12 transition-transform" />
            {{ $t('intake.scan_web.config.btn_start') }}
          </button>
        </div>
      </div>
    </div>

    <!-- Bottom Status / Idle Gallery -->
    <div class="glass p-20 rounded-[4rem] flex flex-col items-center justify-center text-center space-y-8 opacity-40 shadow-inner">
      <div class="w-24 h-24 rounded-lg bg-slate-50 dark:bg-slate-800 flex items-center justify-center text-slate-300 border-4 border-dashed border-slate-100 dark:border-slate-800">
        <LucideScanFace class="w-12 h-12" />
      </div>
      <div class="space-y-2">
        <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('intake.scan_web.status.idle') }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucidePrinter, LucideNetwork, LucideChevronDown, LucideAlertTriangle, 
  LucideInfo, LucideScan, LucideScanFace 
} from 'lucide-vue-next'

const isDuplex = ref(true)

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

