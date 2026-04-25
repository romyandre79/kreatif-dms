<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Stepper & Header -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-2">
        <div class="flex items-center gap-2 text-[9px] font-black text-slate-400 uppercase tracking-widest">
          <span>{{ $t('circulation.receive.breadcrumbs.form') }}</span>
          <LucideChevronRight class="w-3 h-3" />
          <span class="text-primary-500">{{ $t('circulation.receive.breadcrumbs.scanning') }}</span>
          <LucideChevronRight class="w-3 h-3" />
          <span>{{ $t('circulation.receive.breadcrumbs.validation') }}</span>
        </div>
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
          {{ $t('circulation.receive.title') }}
        </h1>
        <p class="text-xs font-bold text-slate-500">{{ $t('circulation.receive.subtitle') }}</p>
      </div>

      <div class="flex items-center gap-6 bg-white dark:bg-slate-900 p-4 rounded-2xl shadow-sm border border-slate-100 dark:border-slate-800">
        <div class="w-10 h-10 rounded-full bg-[#1E3A5F] text-white flex items-center justify-center text-[10px] font-black">2</div>
        <div class="space-y-1.5 w-48">
          <div class="flex items-center justify-between text-[10px] font-black uppercase tracking-widest">
            <span class="text-[#1E3A5F] dark:text-white">SCANNED (40%)</span>
          </div>
          <div class="h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
            <div class="h-full bg-[#1E3A5F]" style="width: 40%"></div>
          </div>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Column: Scanner Controls -->
      <div class="lg:col-span-4 space-y-8">
        <!-- Scanner Connection -->
        <div class="glass p-8 rounded-[2.5rem] space-y-6">
          <div class="flex items-center justify-between">
            <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.receive.scanner.title') }}</h4>
            <span class="px-3 py-1 bg-green-50 text-green-500 rounded-full text-[8px] font-black uppercase tracking-widest border border-green-100 flex items-center gap-1.5">
              <span class="w-1.5 h-1.5 bg-green-500 rounded-full animate-pulse"></span>
              {{ $t('circulation.receive.scanner.ready') }}
            </span>
          </div>
          <div class="flex items-center gap-5 p-6 bg-slate-50 dark:bg-slate-800/50 rounded-2xl border border-slate-100 dark:border-slate-800">
            <div class="w-12 h-12 rounded-xl bg-white dark:bg-slate-900 text-[#1E3A5F] dark:text-white flex items-center justify-center shadow-sm">
              <LucidePrinter class="w-6 h-6" />
            </div>
            <div class="space-y-0.5">
              <p class="text-sm font-black text-[#1E3A5F] dark:text-white">{{ $t('circulation.receive.scanner.name') }}</p>
              <p class="text-[10px] font-bold text-slate-400">{{ $t('circulation.receive.scanner.details') }}</p>
            </div>
          </div>
        </div>

        <!-- Scanning Actions -->
        <div class="space-y-4">
          <button class="w-full py-5 bg-[#1E3A5F] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-2xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center justify-center gap-3 active:scale-95 group">
            <LucideScanLine class="w-5 h-5 group-hover:scale-110 transition-transform" />
            {{ $t('circulation.receive.actions.btn_scan') }}
          </button>
          <button class="w-full py-5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-slate-500 dark:text-slate-400 rounded-2xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center justify-center gap-3">
            <LucideUpload class="w-5 h-5" />
            {{ $t('circulation.receive.actions.btn_upload') }}
          </button>
        </div>

        <!-- Scanning Checklist -->
        <div class="p-8 bg-slate-50/50 dark:bg-slate-900/50 border border-slate-100 dark:border-slate-800 rounded-[2rem] space-y-6">
          <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.receive.checklist.title') }}</h4>
          <ul class="space-y-4">
            <li v-for="i in [1, 2, 3]" :key="i" class="flex gap-4 text-xs font-bold text-slate-600 dark:text-slate-300 leading-relaxed">
              <span class="text-[#1E3A5F] font-black">{{ i }}.</span>
              {{ $t(`circulation.receive.checklist.item_${i}`) }}
            </li>
          </ul>
        </div>

        <!-- Scanned Pages List -->
        <div class="space-y-6">
          <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-2">{{ $t('circulation.receive.pages.title', { count: 3 }) }}</h4>
          <div class="space-y-4">
            <div v-for="p in [1, 2]" :key="p" class="flex items-center justify-between p-4 bg-white dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl group hover:border-primary-500/30 transition-all shadow-sm">
              <div class="flex items-center gap-4">
                <LucideGripVertical class="w-4 h-4 text-slate-200 cursor-move" />
                <div class="w-12 h-16 bg-slate-50 dark:bg-slate-800 rounded-lg overflow-hidden border border-slate-100 dark:border-slate-700 flex items-center justify-center">
                  <img v-if="p === 1" src="https://images.unsplash.com/photo-1544005313-94ddf0286df2?q=80&w=100&auto=format&fit=crop" class="w-full h-full object-cover" />
                  <LucideFileImage v-else class="w-6 h-6 text-slate-300" />
                </div>
                <div class="space-y-0.5">
                  <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase">Page {{ p }}.jpg</p>
                  <p class="text-[9px] font-bold text-slate-400 uppercase tracking-tight">245 KB • 2480 × 3508</p>
                </div>
              </div>
              <button class="p-2 text-slate-300 hover:text-red-500 transition-colors"><LucideTrash2 class="w-4 h-4" /></button>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Column: Document Viewer & QC -->
      <div class="lg:col-span-8 space-y-8">
        <!-- Viewer Toolbar -->
        <div class="bg-white dark:bg-slate-900 p-4 rounded-2xl border border-slate-100 dark:border-slate-800 flex items-center justify-between shadow-sm">
          <div class="flex items-center gap-2">
            <button class="p-2 text-slate-400 hover:text-[#1E3A5F] transition-colors"><LucideSearch class="w-4 h-4" /></button>
            <span class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest mx-2">100%</span>
            <button class="p-2 text-slate-400 hover:text-[#1E3A5F] transition-colors"><LucideZoomIn class="w-4 h-4" /></button>
            <div class="w-px h-4 bg-slate-100 dark:bg-slate-800 mx-2"></div>
            <button class="p-2 text-slate-400 hover:text-[#1E3A5F] transition-colors"><LucideMaximize2 class="w-4 h-4" /></button>
          </div>
          <div class="flex items-center gap-4">
            <button class="p-2 text-slate-400 hover:text-[#1E3A5F] transition-colors"><LucideChevronLeft class="w-5 h-5" /></button>
            <span class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('circulation.receive.viewer.page_info', { current: 1, total: 5 }) }}</span>
            <button class="p-2 text-slate-400 hover:text-[#1E3A5F] transition-colors"><LucideChevronRight class="w-5 h-5" /></button>
          </div>
          <span class="px-4 py-1.5 bg-green-50 text-green-500 rounded-full text-[9px] font-black uppercase tracking-widest border border-green-100 flex items-center gap-2">
            <LucideShieldCheck class="w-3.5 h-3.5" />
            {{ $t('circulation.receive.viewer.quality') }}
          </span>
        </div>

        <!-- Viewer Canvas -->
        <div class="bg-slate-100 dark:bg-slate-950 rounded-[3rem] aspect-[1/1.4] relative shadow-inner overflow-hidden border border-slate-200 dark:border-slate-800 flex items-center justify-center group">
          <div class="absolute inset-0 bg-primary-500/5 opacity-0 group-hover:opacity-100 transition-opacity"></div>
          
          <!-- Mock Document -->
          <div class="w-[80%] h-[90%] bg-white dark:bg-slate-900 rounded-xl shadow-2xl p-16 space-y-12 relative" v-motion-fade>
            <div class="space-y-4">
              <h3 class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tighter uppercase">DOCUMENTARY EVIDENCE</h3>
              <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">REFERENCE: BP-07/AKIRA/2023/QX</p>
            </div>
            
            <div class="space-y-6">
              <div v-for="l in 3" :key="l" class="h-4 bg-slate-50 dark:bg-slate-800 rounded-md" :style="{ width: [100, 80, 95][l-1] + '%' }"></div>
            </div>

            <div class="aspect-video bg-slate-50 dark:bg-slate-800 rounded-2xl flex items-center justify-center border-2 border-dashed border-slate-200 dark:border-slate-700">
              <span class="text-[10px] font-black text-slate-300 uppercase tracking-[0.5em]">[ OFFICIAL SCAN DATA AREA ]</span>
            </div>

            <div class="space-y-6 pt-10">
              <div v-for="l in 4" :key="l" class="h-3 bg-slate-50 dark:bg-slate-800 rounded-md" :style="{ width: [90, 100, 70, 85][l-1] + '%' }"></div>
            </div>

            <div class="absolute bottom-16 left-16 space-y-2 opacity-30">
              <div class="w-24 h-12 bg-slate-100 dark:bg-slate-800 rounded"></div>
              <p class="text-[7px] font-black text-slate-400 uppercase tracking-widest">AUTHORIZED SIGNATORY</p>
            </div>
            
            <!-- Watermark -->
            <div class="absolute inset-0 flex items-center justify-center pointer-events-none opacity-[0.03] rotate-[-35deg] scale-150">
              <span class="text-[120px] font-black text-[#1E3A5F] dark:text-white whitespace-nowrap">PREVIEW</span>
            </div>
          </div>
        </div>

        <!-- Quality Control Footer -->
        <div class="glass p-10 rounded-[2.5rem] space-y-8" v-motion-slide-visible-bottom>
          <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('circulation.receive.quality.title') }}</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div v-for="qc in [1, 2, 3, 4]" :key="qc" class="flex items-center gap-4 group cursor-pointer" @click="qcChecked[qc-1] = !qcChecked[qc-1]">
              <div :class="`w-6 h-6 rounded-lg border-2 flex items-center justify-center transition-all ${qcChecked[qc-1] ? 'bg-[#1E3A5F] border-[#1E3A5F] shadow-lg shadow-blue-900/20' : 'bg-white border-slate-200'}`">
                <LucideCheck v-if="qcChecked[qc-1]" class="w-4 h-4 text-white" />
              </div>
              <span class="text-xs font-bold text-slate-600 dark:text-slate-300 group-hover:text-[#1E3A5F] transition-colors">{{ $t(`circulation.receive.quality.item_${qc}`) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideChevronRight, LucidePrinter, LucideScanLine, LucideUpload, 
  LucideGripVertical, LucideTrash2, LucideSearch, LucideZoomIn,
  LucideMaximize2, LucideChevronLeft, LucideChevronRight as LucideChevronRightIcon,
  LucideShieldCheck, LucideCheck, LucideFileImage
} from 'lucide-vue-next'

const qcChecked = ref([true, true, true, true])
</script>

<style scoped>
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}
</style>
