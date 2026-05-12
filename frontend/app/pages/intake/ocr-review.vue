<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-1">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('intake.ocr_review.title') }}</h1>
        <p class="text-xs font-bold text-slate-500">{{ $t('intake.ocr_review.subtitle') }}</p>
      </div>
      <div class="flex items-center gap-4">
        <button class="px-6 py-3 text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl transition-all border border-slate-200 dark:border-slate-700">
          {{ $t('intake.ocr_review.header.rescan') }}
        </button>
        <button class="px-10 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3">
          <LucideCheckCircle2 class="w-4 h-4" />
          {{ $t('intake.ocr_review.header.accept') }}
        </button>
      </div>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div v-for="s in [
        { id: 'status', val: 'COMPLETED', color: 'text-green-500', icon: LucideActivity },
        { id: 'confidence', val: '92.4%', color: 'text-green-500', icon: LucideZap, progress: true },
        { id: 'lang', val: 'Indonesian (ID)', color: 'text-blue-500', icon: LucideGlobe },
        { id: 'duplicate', val: 'LOW RISK', sub: 'Matching content < 5%', color: 'text-green-500', icon: LucideCopyCheck, badge: true }
      ]" :key="s.id" class="glass p-8 rounded-lg space-y-4 shadow-sm border border-slate-50 dark:border-slate-800">
        <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t(`intake.ocr_review.stats.${s.id}`) }}</p>
        <div class="flex items-center gap-4">
          <div :class="`w-8 h-8 rounded-lg flex items-center justify-center bg-slate-50 dark:bg-slate-800 ${s.color}`">
            <component :is="s.icon" class="w-4 h-4" />
          </div>
          <div class="space-y-1">
            <p :class="`text-2xl font-black ${s.id === 'status' || s.id === 'duplicate' ? s.color : 'text-[#1E3A5F] dark:text-white'} tracking-tighter`">{{ s.val }}</p>
            <p v-if="s.sub" class="text-[8px] font-black text-slate-400 uppercase">{{ s.sub }}</p>
            <div v-if="s.progress" class="w-24 h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
              <div class="h-full bg-green-500" style="width: 92.4%"></div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Column: Document Viewer -->
      <div class="lg:col-span-7 flex flex-col gap-6">
        <div class="glass rounded-lg overflow-hidden flex flex-col shadow-xl border border-slate-50 dark:border-slate-800 h-[800px]">
          <!-- Viewer Toolbar -->
          <div class="p-6 bg-slate-50/50 dark:bg-slate-800/50 flex items-center justify-between border-b border-slate-100 dark:border-slate-700">
            <div class="flex items-center gap-6">
              <span class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('intake.ocr_review.viewer.page', { curr: 1, total: 3 }) }}</span>
              <span class="px-3 py-1 bg-white dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-lg text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.ocr_review.viewer.layer') }}</span>
            </div>
            <div class="flex items-center gap-4">
              <LucideSearch class="w-4 h-4 text-slate-400 cursor-pointer hover:text-[#1E3A5F]" />
              <LucideSearch class="w-4 h-4 text-slate-400 cursor-pointer hover:text-[#1E3A5F]" />
            </div>
          </div>

          <!-- Document Canvas -->
          <div class="flex-grow relative bg-slate-200/30 dark:bg-slate-950 p-10 overflow-auto custom-scrollbar flex justify-center items-start">
            <div class="w-[600px] bg-white shadow-2xl p-16 space-y-10 relative">
              <p class="text-[8px] font-bold text-slate-300 text-right uppercase tracking-widest">INTERNAL - CONFIDENTIAL</p>
              <h4 class="text-xs font-black text-[#1E3A5F] uppercase italic">PT AKIRADATA LOGISTICS SOLUTIONS</h4>
              
              <!-- Highlighted Reference -->
              <div class="relative inline-block px-2 py-1 bg-blue-100/40 border border-blue-400 rounded-sm">
                <p class="text-xs font-black text-blue-600 uppercase font-mono">REF NO: AR-2023-X990/LOG</p>
              </div>

              <!-- Highlighted Title -->
              <div class="relative space-y-2">
                <div class="absolute -left-8 top-0 w-5 h-5 bg-[#1E3A5F] text-white rounded-full flex items-center justify-center text-[10px] font-black">1</div>
                <div class="absolute -left-8 top-10 w-5 h-5 bg-[#1E3A5F] text-white rounded-full flex items-center justify-center text-[10px] font-black">2</div>
                <h2 class="text-xl font-black text-[#1E3A5F] uppercase leading-tight tracking-tight">
                  PERSETUJUAN PENGADAAN PERANGKAT SERVER
                </h2>
              </div>

              <p class="text-[11px] font-bold text-slate-600 leading-relaxed">
                Berdasarkan hasil verifikasi tim teknis pada tanggal <span class="text-blue-600 border-b-2 border-blue-200">15 OKTOBER 2023</span>, maka dengan ini diputuskan bahwa PT SOLUSI TEKNOLOGI ABADI terpilih sebagai vendor utama dalam proyek modernisasi infrastruktur data center.
              </p>

              <!-- Extracted Regions -->
              <div class="grid grid-cols-2 gap-8 pt-6">
                <div class="border-2 border-dashed border-slate-200 p-6 rounded-xl space-y-3 relative">
                  <span class="absolute -top-3 left-4 bg-white px-2 text-[8px] font-black text-slate-400 uppercase">{{ $t('intake.ocr_review.viewer.region', { id: '01' }) }}</span>
                  <p class="text-[10px] font-black text-[#1E3A5F] uppercase">NILAI KONTRAK: IDR 450.000.000</p>
                </div>
                <div class="border-2 border-dashed border-slate-200 p-6 rounded-xl space-y-3 relative">
                  <span class="absolute -top-3 left-4 bg-white px-2 text-[8px] font-black text-slate-400 uppercase">{{ $t('intake.ocr_review.viewer.region', { id: '02' }) }}</span>
                  <p class="text-[10px] font-black text-[#1E3A5F] uppercase">NPWP: 01.234.567.8-901.000</p>
                </div>
              </div>

              <div class="pt-10 flex flex-col items-center gap-4 opacity-20">
                <LucidePenTool class="w-10 h-10 text-slate-300" />
                <p class="text-[10px] font-black text-slate-300 uppercase italic">Signature Detected</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Viewer Footer -->
        <div class="bg-[#1E3A5F] p-4 rounded-2xl flex items-center justify-between text-white shadow-xl shadow-blue-900/20">
          <div class="flex items-center gap-4">
            <span class="px-3 py-1 bg-white/10 rounded-lg text-[9px] font-black uppercase tracking-widest">{{ $t('intake.ocr_review.viewer.selected') }}</span>
            <p class="text-xs font-black uppercase tracking-tight">DOC_INVOICE_091.PDF</p>
          </div>
          <div class="flex items-center gap-6">
            <button class="flex items-center gap-2 text-[10px] font-black uppercase tracking-widest hover:text-blue-300 transition-colors">
              <LucideChevronLeft class="w-4 h-4" />
              {{ $t('intake.ocr_review.viewer.prev') }}
            </button>
            <span class="text-[11px] font-black font-mono">01 / 15</span>
            <button class="flex items-center gap-2 text-[10px] font-black uppercase tracking-widest hover:text-blue-300 transition-colors">
              {{ $t('intake.ocr_review.viewer.next') }}
              <LucideChevronRight class="w-4 h-4" />
            </button>
          </div>
        </div>
      </div>

      <!-- Right Column: Suggestions Panel -->
      <div class="lg:col-span-5 space-y-8">
        <div class="glass p-10 rounded-lg space-y-10 shadow-sm border border-slate-50 dark:border-slate-800 h-full flex flex-col">
          <div class="flex items-center justify-between">
            <div class="space-y-1">
              <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('intake.ocr_review.suggestions.title') }}</h3>
              <p class="text-[9px] font-bold text-blue-500 uppercase">{{ $t('intake.ocr_review.suggestions.pending', { count: 3 }) }}</p>
            </div>
            <button class="text-[9px] font-black text-slate-300 uppercase tracking-widest hover:text-[#1E3A5F]">{{ $t('intake.ocr_review.suggestions.logs') }}</button>
          </div>

          <div class="flex-grow space-y-10 overflow-y-auto pr-4 custom-scrollbar">
            <!-- Ref Field -->
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.ocr_review.suggestions.ref') }}</label>
                <span class="flex items-center gap-2 text-[8px] font-black text-green-500 uppercase tracking-widest">
                  <span class="w-1.5 h-1.5 bg-green-500 rounded-full"></span>
                  {{ $t('intake.ocr_review.suggestions.high', { val: 98 }) }}
                </span>
              </div>
              <div class="flex gap-4">
                <input type="text" value="AR-2023-X990/LOG" class="flex-grow h-14 px-6 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold font-mono outline-none focus:border-[#1E3A5F]" />
                <button class="w-14 h-14 bg-green-50 text-green-500 rounded-2xl flex items-center justify-center hover:bg-green-100 transition-all"><LucideCheck class="w-5 h-5" /></button>
                <button class="w-14 h-14 bg-slate-50 text-slate-400 rounded-2xl flex items-center justify-center hover:bg-slate-100 transition-all"><LucideEdit2 class="w-5 h-5" /></button>
              </div>
            </div>

            <!-- Title Field -->
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.ocr_review.suggestions.title_field') }}</label>
                <span class="flex items-center gap-2 text-[8px] font-black text-amber-500 uppercase tracking-widest">
                  <span class="w-1.5 h-1.5 bg-amber-500 rounded-full"></span>
                  {{ $t('intake.ocr_review.suggestions.required', { val: 72 }) }}
                </span>
              </div>
              <div class="relative">
                <textarea class="w-full p-6 bg-slate-50 dark:bg-slate-900 border-2 border-amber-100 rounded-2xl text-xs font-bold leading-relaxed outline-none focus:border-amber-500 h-32 resize-none">PERSETUJUAN PENGADAAN PERANGKAT SERVER UTAMA MODERNISASI</textarea>
                <div class="absolute right-4 top-4 flex flex-col gap-3">
                  <button class="w-10 h-10 bg-amber-50 text-amber-500 rounded-xl flex items-center justify-center shadow-sm"><LucideZap class="w-4 h-4" /></button>
                  <button class="w-10 h-10 bg-white dark:bg-slate-800 text-slate-300 rounded-xl border border-slate-100 dark:border-slate-700 flex items-center justify-center shadow-sm"><LucideX class="w-4 h-4" /></button>
                </div>
              </div>
              <p class="text-[9px] font-bold text-amber-600 italic">Note: Multiple title candidates detected on page 1.</p>
            </div>

            <!-- Entities -->
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.ocr_review.suggestions.entities') }}</label>
                <span class="flex items-center gap-2 text-[8px] font-black text-green-500 uppercase tracking-widest">
                  <span class="w-1.5 h-1.5 bg-green-500 rounded-full"></span>
                  {{ $t('intake.ocr_review.suggestions.auto', { val: 88 }) }}
                </span>
              </div>
              <div class="flex flex-wrap gap-3">
                <div v-for="e in [
                  { id: 'org', val: 'PT SOLUSI TEKNOLOGI ABADI' },
                  { id: 'person', val: 'AHMAD SUBARJO' }
                ]" :key="e.val" class="px-4 py-2 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-full flex items-center gap-3">
                  <span class="text-[8px] font-black text-blue-500 uppercase tracking-widest">{{ e.id.toUpperCase() }}</span>
                  <span class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase">{{ e.val }}</span>
                  <LucideX class="w-3.5 h-3.5 text-slate-300 cursor-pointer hover:text-red-500" />
                </div>
              </div>
              <button class="w-full py-4 border-2 border-dashed border-slate-100 dark:border-slate-800 text-[9px] font-black text-slate-400 uppercase tracking-widest rounded-2xl hover:bg-slate-50 transition-all">
                {{ $t('intake.ocr_review.suggestions.add') }}
              </button>
            </div>

            <!-- Expired -->
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <label class="text-[10px] font-black text-red-500 uppercase tracking-widest">{{ $t('intake.ocr_review.suggestions.expired') }}</label>
                <span class="flex items-center gap-2 text-[8px] font-black text-red-500 uppercase tracking-widest">
                  <span class="w-1.5 h-1.5 bg-red-500 rounded-full animate-pulse"></span>
                  {{ $t('intake.ocr_review.suggestions.low', { val: 34 }) }}
                </span>
              </div>
              <div class="flex gap-4">
                <div class="flex-grow h-14 px-6 bg-red-50/30 border-2 border-red-100 rounded-2xl flex items-center text-xs font-bold text-slate-400 italic">
                  {{ $t('intake.ocr_review.suggestions.not_found') }}
                </div>
                <button class="px-8 h-14 bg-red-500 text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-lg shadow-red-500/20 hover:bg-red-600 transition-all">
                  {{ $t('intake.ocr_review.suggestions.manual') }}
                </button>
              </div>
            </div>
          </div>

          <button class="w-full py-5 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40 hover:bg-[#152943] transition-all flex items-center justify-center gap-4">
            <LucideSave class="w-4 h-4" />
            {{ $t('intake.ocr_review.suggestions.btn_save') }}
          </button>
        </div>

        <!-- Historical Data Card -->
        <div class="p-8 bg-blue-50/50 dark:bg-blue-900/10 rounded-lg border border-blue-100 dark:border-blue-800 flex items-center gap-6 group">
          <div class="w-14 h-14 rounded-2xl bg-white dark:bg-slate-800 flex items-center justify-center text-blue-500 shadow-sm border border-slate-100 dark:border-slate-800">
            <LucideHistory class="w-6 h-6" />
          </div>
          <div class="flex-grow space-y-1">
            <h5 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('intake.ocr_review.historical.title') }}</h5>
            <p class="text-[9px] font-bold text-slate-400 leading-relaxed">{{ $t('intake.ocr_review.historical.desc') }}</p>
          </div>
          <button class="text-[10px] font-black text-blue-500 uppercase tracking-widest hover:underline">{{ $t('intake.ocr_review.historical.btn') }}</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideCheckCircle2, LucideActivity, LucideZap, LucideGlobe, 
  LucideCopyCheck, LucideSearch, LucidePenTool, LucideChevronLeft, 
  LucideChevronRight, LucideCheck, LucideEdit2, LucideX, LucideSave, 
  LucideHistory 
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

