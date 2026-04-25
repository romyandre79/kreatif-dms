<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-1">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('warehouse.generation.title') }}</h1>
        <p class="text-xs font-bold text-slate-500 uppercase tracking-widest">{{ $t('warehouse.generation.subtitle') }}</p>
      </div>
      <div class="flex items-center gap-4">
        <button class="px-6 py-3 text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest flex items-center gap-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl transition-all border border-slate-200 dark:border-slate-700">
          <LucideDownload class="w-4 h-4" />
          {{ $t('warehouse.generation.header.export') }}
        </button>
        <button class="px-10 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 group">
          <LucidePrinter class="w-4 h-4 group-hover:scale-110 transition-transform" />
          {{ $t('warehouse.generation.header.print') }}
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Column: Template Selector & Batch Controls -->
      <div class="lg:col-span-3 space-y-10">
        <div class="glass p-10 rounded-[3rem] space-y-8 shadow-sm border border-slate-50 dark:border-slate-800">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.generation.templates.title') }}</h3>
          <div class="space-y-4">
            <div v-for="t in [
              { id: 'child', icon: LucideFileText, active: true },
              { id: 'parent', icon: LucideBox },
              { id: 'rack', icon: LucideLayoutGrid }
            ]" :key="t.id" :class="`p-6 rounded-2xl border-2 transition-all cursor-pointer group ${t.active ? 'border-blue-500 bg-blue-50/30 shadow-lg shadow-blue-500/5' : 'border-slate-50 dark:border-slate-800 hover:border-slate-200'}`">
              <div class="flex items-center gap-6">
                <div :class="`w-12 h-12 rounded-xl flex items-center justify-center transition-colors ${t.active ? 'bg-blue-500 text-white' : 'bg-slate-50 dark:bg-slate-800 text-slate-300 group-hover:text-[#1E3A5F]'}`">
                  <component :is="t.icon" class="w-6 h-6" />
                </div>
                <div class="space-y-0.5">
                  <h4 :class="`text-[11px] font-black uppercase tracking-tight ${t.active ? 'text-[#1E3A5F] dark:text-white' : 'text-slate-400'}`">{{ $t(`warehouse.generation.templates.${t.id}`) }}</h4>
                  <p class="text-[8px] font-bold text-slate-400 uppercase tracking-widest">{{ $t(`warehouse.generation.templates.${t.id}_desc`) }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="glass p-10 rounded-[3rem] space-y-8 shadow-sm border border-slate-50 dark:border-slate-800">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.generation.batch.title') }}</h3>
          <div class="space-y-6">
            <div class="space-y-3">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.generation.batch.sequence') }}</label>
              <input type="text" value="AK-2024-0001" class="w-full h-14 px-6 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold outline-none focus:border-blue-500 transition-all font-mono" />
            </div>
            <div class="space-y-3">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.generation.batch.qty') }}</label>
              <input type="number" value="25" class="w-full h-14 px-6 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold outline-none focus:border-blue-500 transition-all" />
            </div>
            <button class="w-full py-4 bg-slate-950 text-white rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-black transition-all active:scale-95 shadow-xl shadow-black/20">
              {{ $t('warehouse.generation.batch.btn') }}
            </button>
          </div>
        </div>
      </div>

      <!-- Center Column: Live Preview -->
      <div class="lg:col-span-6 space-y-10">
        <div class="glass rounded-[4rem] flex flex-col h-full shadow-sm border border-slate-50 dark:border-slate-800 relative overflow-hidden">
          <div class="p-10 flex items-center justify-between border-b border-slate-50 dark:border-slate-800 bg-white/50 dark:bg-slate-900/50 backdrop-blur-md z-10">
            <div class="flex items-center gap-4">
              <div class="w-2 h-2 bg-green-500 rounded-full animate-pulse shadow-sm shadow-green-500/20"></div>
              <h3 class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('warehouse.generation.preview.title') }}</h3>
            </div>
            <div class="flex items-center gap-4 text-slate-300">
              <LucideSearch class="w-4 h-4 cursor-pointer hover:text-blue-500 transition-colors" />
              <LucideMaximize2 class="w-4 h-4 cursor-pointer hover:text-blue-500 transition-colors" />
            </div>
          </div>

          <!-- Label Viewport -->
          <div class="flex-grow flex flex-col items-center justify-center p-20 bg-slate-50/30 dark:bg-slate-950/30 relative">
            <div class="absolute top-10 left-10">
              <p class="text-[8px] font-black text-slate-300 uppercase tracking-widest">SCALE: 1:1 PREVIEW</p>
            </div>
            
            <!-- RFID/QR Label Mockup -->
            <div class="w-full max-w-lg bg-white dark:bg-slate-900 shadow-2xl rounded-sm p-12 border border-slate-100 dark:border-slate-800 relative group transition-transform hover:scale-105 duration-500">
              <div class="flex gap-10">
                <!-- Left: QR/Brand -->
                <div class="w-32 space-y-4">
                  <div class="aspect-square bg-slate-800 dark:bg-white rounded-lg flex items-center justify-center p-2">
                    <LucideQrCode class="w-full h-full text-white dark:text-slate-800" />
                  </div>
                </div>
                <!-- Right: Content -->
                <div class="flex-grow space-y-6">
                  <div class="flex items-center justify-between">
                    <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">AKIRADATA PROPERTY</p>
                    <span class="px-3 py-1 bg-blue-50 text-blue-500 rounded text-[7px] font-black uppercase tracking-widest flex items-center gap-1.5 animate-pulse">
                      <LucideRss class="w-2.5 h-2.5" />
                      RFID ENABLED
                    </span>
                  </div>
                  <div class="space-y-1">
                    <h2 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight">DOC-2024-8892-F16</h2>
                    <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">CLASSIFICATION: RESTRICTED / PERMANENT ARCHIVE</p>
                  </div>
                  <div class="flex items-end justify-between pt-4 border-t border-slate-50 dark:border-slate-800">
                    <div class="grid grid-cols-2 gap-4">
                      <div>
                        <p class="text-[7px] font-black text-slate-300 uppercase">ENCODED DATE</p>
                        <p class="text-[9px] font-bold text-slate-500 uppercase tracking-tighter">2024-05-24</p>
                      </div>
                      <div>
                        <p class="text-[7px] font-black text-slate-300 uppercase">BATCH ID</p>
                        <p class="text-[9px] font-bold text-slate-500 uppercase tracking-tighter font-mono">B-7728-X</p>
                      </div>
                    </div>
                    <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest opacity-40">PT AKIRADATA</p>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Preview Info Footer -->
          <div class="p-10 grid grid-cols-2 gap-8 border-t border-slate-50 dark:border-slate-800 bg-white/50 dark:bg-slate-900/50 backdrop-blur-md">
            <div class="flex items-center gap-6 p-6 bg-slate-50/50 dark:bg-slate-800/50 rounded-[2rem] border border-slate-100 dark:border-slate-700">
              <div class="w-12 h-12 rounded-2xl bg-white dark:bg-slate-900 flex items-center justify-center text-blue-500 shadow-sm">
                <LucideRss class="w-6 h-6" />
              </div>
              <div class="space-y-0.5">
                <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.generation.preview.freq') }}</p>
                <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase">{{ $t('warehouse.generation.preview.freq_val') }}</p>
              </div>
            </div>
            <div class="flex items-center gap-6 p-6 bg-slate-50/50 dark:bg-slate-800/50 rounded-[2rem] border border-slate-100 dark:border-slate-700">
              <div class="w-12 h-12 rounded-2xl bg-white dark:bg-slate-900 flex items-center justify-center text-[#1E3A5F] dark:text-white shadow-sm">
                <LucideQrCode class="w-6 h-6" />
              </div>
              <div class="space-y-0.5">
                <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.generation.preview.qr') }}</p>
                <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase">{{ $t('warehouse.generation.preview.qr_val') }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Column: Encoding Options & Queue -->
      <div class="lg:col-span-3 space-y-10">
        <div class="glass p-10 rounded-[3rem] space-y-10 shadow-sm border border-slate-50 dark:border-slate-800">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.generation.encoding.title') }}</h3>
          <div class="space-y-4">
            <div v-for="o in ['ref', 'full']" :key="o" :class="`p-6 rounded-2xl border-2 transition-all cursor-pointer ${o === 'ref' ? 'border-blue-500 bg-blue-50/30 shadow-lg shadow-blue-500/5' : 'border-slate-50 dark:border-slate-800 hover:border-slate-100'}`">
              <div class="flex items-start gap-4">
                <div :class="`w-6 h-6 rounded-full border-4 flex items-center justify-center ${o === 'ref' ? 'border-blue-500 bg-blue-500 text-white' : 'border-slate-200'}`">
                  <div v-if="o === 'ref'" class="w-1.5 h-1.5 bg-white rounded-full"></div>
                </div>
                <div class="space-y-1">
                  <h4 :class="`text-[11px] font-black uppercase tracking-tight ${o === 'ref' ? 'text-[#1E3A5F] dark:text-white' : 'text-slate-400'}`">{{ $t(`warehouse.generation.encoding.${o}`) }}</h4>
                  <p class="text-[9px] font-bold text-slate-400 uppercase leading-tight">{{ $t(`warehouse.generation.encoding.${o}_desc`) }}</p>
                </div>
              </div>
            </div>
          </div>

          <div class="pt-6 border-t border-slate-50 dark:border-slate-800 space-y-6">
            <div class="flex items-center justify-between">
              <label class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('warehouse.generation.encoding.lock') }}</label>
              <div class="w-12 h-6 bg-blue-500 rounded-full relative shadow-inner">
                <div class="absolute right-1 top-1 w-4 h-4 bg-white rounded-full shadow-sm"></div>
              </div>
            </div>
            <div class="flex items-center justify-between opacity-40">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.generation.encoding.verify') }}</label>
              <div class="w-12 h-6 bg-slate-200 rounded-full relative">
                <div class="absolute left-1 top-1 w-4 h-4 bg-white rounded-full"></div>
              </div>
            </div>
          </div>
        </div>

        <!-- Print Queue -->
        <div class="glass rounded-[3rem] overflow-hidden shadow-sm border border-slate-50 dark:border-slate-800 flex flex-col h-[400px]">
          <div class="p-8 border-b border-slate-50 dark:border-slate-800 flex items-center justify-between bg-white/50 dark:bg-slate-900/50">
            <h3 class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('warehouse.generation.queue.title') }}</h3>
            <span class="px-3 py-1 bg-blue-50 text-blue-500 rounded text-[8px] font-black uppercase tracking-widest">{{ $t('warehouse.generation.queue.count', { count: 12 }) }}</span>
          </div>
          
          <div class="flex-grow p-6 overflow-y-auto custom-scrollbar space-y-3">
            <div v-for="i in 12" :key="i" class="p-4 bg-slate-50/50 dark:bg-slate-800/50 rounded-xl border border-slate-100 dark:border-slate-700 flex items-center justify-between group hover:bg-slate-50 transition-all cursor-default">
              <div class="space-y-0.5">
                <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white tracking-tighter uppercase font-mono">AK-2024-000{{ i }}</p>
                <p class="text-[7px] font-bold text-slate-400 uppercase">Child Doc F-16</p>
              </div>
              <LucideX class="w-3.5 h-3.5 text-slate-300 hover:text-red-500 cursor-pointer transition-colors" />
            </div>
          </div>

          <div class="p-8 bg-slate-50/30 dark:bg-slate-800/30 border-t border-slate-50 dark:border-slate-800 space-y-6">
            <div class="flex items-center justify-between px-2">
              <span class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('warehouse.generation.queue.printer') }}</span>
              <div class="flex items-center gap-2">
                <div class="w-1.5 h-1.5 bg-green-500 rounded-full"></div>
                <span class="text-[9px] font-black text-green-600 uppercase tracking-widest">{{ $t('warehouse.generation.queue.online') }}</span>
              </div>
            </div>
            <button class="w-full py-4 bg-white dark:bg-slate-800 border-2 border-slate-100 dark:border-slate-700 text-[#1E3A5F] dark:text-white rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-50 transition-all">
              {{ $t('warehouse.generation.queue.btn_clear') }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideDownload, LucidePrinter, LucideFileText, LucideBox, 
  LucideLayoutGrid, LucideSearch, LucideMaximize2, LucideQrCode, 
  LucideRss, LucideX 
} from 'lucide-vue-next'

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
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
