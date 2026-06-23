<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-10">
      <div class="space-y-1">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('warehouse.topology.title') }}</h1>
        <p class="text-xs font-bold text-slate-500 uppercase tracking-widest">{{ $t('warehouse.topology.subtitle') }}</p>
      </div>
      <div class="flex items-center gap-6">
        <div class="flex items-center gap-3 glass px-6 py-3 rounded-2xl border border-slate-50">
          <LucideActivity class="w-4 h-4 text-green-500" />
          <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('warehouse.topology.sensors', { val: '42' }) }}</p>
        </div>
        <button class="px-8 py-3.5 bg-[#1E3A5F] text-white rounded-2xl text-[11px] font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-4 group">
          <LucideSettings class="w-4 h-4 group-hover:rotate-90 transition-transform duration-500" />
          {{ $t('warehouse.topology.btn_config') }}
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Column: Tree Explorer -->
      <div class="lg:col-span-4 space-y-6 min-h-[500px] flex flex-col">
        <div class="glass p-6 rounded-3xl space-y-6 shadow-sm border border-slate-50 dark:border-slate-800 flex-grow flex flex-col overflow-hidden">
          <div class="flex items-center justify-between">
            <h3 class="text-[11px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.topology.explorer.title') }}</h3>
            <LucideSearch class="w-4 h-4 text-slate-300 cursor-pointer" />
          </div>

          <div class="flex-grow overflow-y-auto custom-scrollbar pr-4 space-y-4">
            <div v-for="hub in [
              { id: 'JKT', name: 'Jakarta Hub (Main)', children: [
                { id: 'WH-A', name: 'Warehouse A - Archives', children: [
                  { id: 'SEC-1', name: 'Section 1: Legal', status: 'full' },
                  { id: 'SEC-2', name: 'Section 2: Finance', status: 'warning' }
                ]},
                { id: 'WH-B', name: 'Warehouse B - Bulk' }
              ]},
              { id: 'BDG', name: 'Bandung Hub' }
            ]" :key="hub.id" class="space-y-3">
              <div class="flex items-center gap-3 group cursor-pointer">
                <LucideChevronRight class="w-4 h-4 text-slate-300 group-hover:text-blue-500 transition-colors" />
                <LucideBuilding2 class="w-5 h-5 text-[#1E3A5F] dark:text-white" />
                <span class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ hub.name }}</span>
              </div>
              
              <div v-if="hub.children" class="pl-9 space-y-3 border-l-2 border-slate-50 dark:border-slate-800">
                <div v-for="wh in hub.children" :key="wh.id" class="space-y-3">
                  <div class="flex items-center gap-3 group cursor-pointer">
                    <LucideWarehouse class="w-4 h-4 text-blue-500" />
                    <span class="text-[11px] font-bold text-slate-500 uppercase tracking-tight">{{ wh.name }}</span>
                  </div>
                  
                  <div v-if="wh.children" class="pl-8 space-y-2 border-l-2 border-blue-50 dark:border-blue-900/30">
                    <div v-for="sec in wh.children" :key="sec.id" :class="`flex items-center justify-between p-3 rounded-xl transition-all cursor-pointer ${sec.status === 'full' ? 'bg-red-50/50 hover:bg-red-50' : 'hover:bg-slate-50 dark:hover:bg-slate-800'}`">
                      <div class="flex items-center gap-3">
                        <div :class="`w-1.5 h-1.5 rounded-full ${sec.status === 'full' ? 'bg-red-500' : 'bg-amber-500'}`"></div>
                        <span :class="`text-[10px] font-black uppercase tracking-tight ${sec.status === 'full' ? 'text-red-700' : 'text-amber-700'}`">{{ sec.id }}</span>
                        <span class="text-[9px] font-bold text-slate-400 uppercase truncate max-w-[100px]">{{ sec.name }}</span>
                      </div>
                      <LucideArrowRight class="w-3 h-3 text-slate-300" />
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Column: Visual Map Placeholder -->
      <div class="lg:col-span-8 min-h-[500px] flex flex-col">
        <div class="glass rounded-3xl overflow-hidden shadow-sm border border-slate-100 dark:border-slate-800 flex-grow relative bg-slate-50/50 dark:bg-slate-900/50 flex items-center justify-center p-10 text-center">
          
          <div class="absolute inset-0 opacity-10 pointer-events-none" style="background-image: radial-gradient(circle, #cbd5e1 1px, transparent 1px); background-size: 30px 30px;"></div>
          
          <div class="relative z-10 max-w-md mx-auto space-y-6">
            <div class="w-24 h-24 bg-white dark:bg-slate-800 rounded-[2rem] shadow-xl shadow-slate-200/50 dark:shadow-slate-900/50 flex items-center justify-center mx-auto border border-slate-100 dark:border-slate-700 transform -rotate-6 hover:rotate-0 transition-transform duration-500">
              <LucideMap class="w-10 h-10 text-slate-300 dark:text-slate-500" />
            </div>
            <div class="space-y-2">
              <h3 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Live Map Visualization</h3>
              <p class="text-xs font-bold text-slate-400 leading-relaxed">
                Fitur Peta Visual Interaktif belum diaktifkan. Modul ini akan tersedia pada pembaruan sistem tahap berikutnya untuk memantau kapasitas penyimpanan secara real-time.
              </p>
            </div>
            
            <button disabled class="px-6 py-2.5 bg-slate-100 dark:bg-slate-800 text-slate-400 dark:text-slate-500 rounded-xl text-[10px] font-black uppercase tracking-widest cursor-not-allowed border border-slate-200 dark:border-slate-700">
              Segera Hadir
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideActivity, LucideSettings, LucideSearch, LucideBuilding2, 
  LucideChevronRight, LucideWarehouse, LucideArrowRight, LucideMap
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

