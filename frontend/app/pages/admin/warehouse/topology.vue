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
      <div class="lg:col-span-4 space-y-10 h-[800px] flex flex-col">
        <div class="glass p-10 rounded-[4rem] space-y-10 shadow-sm border border-slate-50 dark:border-slate-800 flex-grow flex flex-col overflow-hidden">
          <div class="flex items-center justify-between">
            <h3 class="text-[11px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.topology.explorer.title') }}</h3>
            <LucideSearch class="w-4 h-4 text-slate-300 cursor-pointer" />
          </div>

          <div class="flex-grow overflow-y-auto custom-scrollbar pr-4 space-y-6">
            <div v-for="hub in [
              { id: 'JKT', name: 'Jakarta Hub (Main)', children: [
                { id: 'WH-A', name: 'Warehouse A - Archives', children: [
                  { id: 'SEC-1', name: 'Section 1: Legal', status: 'full' },
                  { id: 'SEC-2', name: 'Section 2: Finance', status: 'warning' }
                ]},
                { id: 'WH-B', name: 'Warehouse B - Bulk' }
              ]},
              { id: 'BDG', name: 'Bandung Hub' }
            ]" :key="hub.id" class="space-y-4">
              <div class="flex items-center gap-4 group cursor-pointer">
                <LucideChevronRight class="w-4 h-4 text-slate-300 group-hover:text-blue-500 transition-colors" />
                <LucideBuilding2 class="w-6 h-6 text-[#1E3A5F] dark:text-white" />
                <span class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ hub.name }}</span>
              </div>
              
              <div v-if="hub.children" class="pl-12 space-y-4 border-l-2 border-slate-50 dark:border-slate-800">
                <div v-for="wh in hub.children" :key="wh.id" class="space-y-4">
                  <div class="flex items-center gap-4 group cursor-pointer">
                    <LucideWarehouse class="w-5 h-5 text-blue-500" />
                    <span class="text-xs font-bold text-slate-500 uppercase tracking-tight">{{ wh.name }}</span>
                  </div>
                  
                  <div v-if="wh.children" class="pl-10 space-y-3 border-l-2 border-blue-50 dark:border-blue-900/30">
                    <div v-for="sec in wh.children" :key="sec.id" :class="`flex items-center justify-between p-4 rounded-xl transition-all cursor-pointer ${sec.status === 'full' ? 'bg-red-50/50 hover:bg-red-50' : 'hover:bg-slate-50'}`">
                      <div class="flex items-center gap-4">
                        <div :class="`w-2 h-2 rounded-full ${sec.status === 'full' ? 'bg-red-500' : 'bg-amber-500'}`"></div>
                        <span :class="`text-[11px] font-black uppercase tracking-tight ${sec.status === 'full' ? 'text-red-700' : 'text-amber-700'}`">{{ sec.id }}</span>
                        <span class="text-[9px] font-bold text-slate-400 uppercase">{{ sec.name }}</span>
                      </div>
                      <LucideArrowRight class="w-3.5 h-3.5 text-slate-300" />
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Column: Visual Map / Graph -->
      <div class="lg:col-span-8 space-y-10 h-[800px] flex flex-col">
        <div class="glass rounded-[5rem] overflow-hidden shadow-sm border border-slate-50 dark:border-slate-800 flex-grow relative bg-slate-50/30 dark:bg-slate-950/30">
          <!-- Graph Toolbar -->
          <div class="absolute top-10 left-10 z-10 flex items-center gap-4">
            <div class="glass px-6 py-4 rounded-2xl flex items-center gap-6 shadow-xl shadow-slate-900/5 border border-white">
              <LucideZoomIn class="w-5 h-5 text-slate-300 cursor-pointer hover:text-blue-500 transition-colors" />
              <LucideZoomOut class="w-5 h-5 text-slate-300 cursor-pointer hover:text-blue-500 transition-colors" />
              <div class="w-px h-6 bg-slate-100 mx-2"></div>
              <LucideFocus class="w-5 h-5 text-[#1E3A5F] cursor-pointer" />
            </div>
            <div class="glass px-6 py-4 rounded-2xl flex items-center gap-4 shadow-xl shadow-slate-900/5 border border-white">
              <span class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.topology.map.layers') }}</span>
              <div class="flex items-center gap-2">
                <span class="px-3 py-1 bg-blue-500 text-white rounded-lg text-[8px] font-black uppercase tracking-widest">SENSORS</span>
                <span class="px-3 py-1 bg-white text-[#1E3A5F] rounded-lg text-[8px] font-black uppercase tracking-widest border border-slate-100">THERMAL</span>
              </div>
            </div>
          </div>

          <!-- Topology Visualization (Mockup) -->
          <div class="w-full h-full flex flex-col items-center justify-center p-20 relative overflow-hidden">
            <!-- Grid Background -->
            <div class="absolute inset-0 opacity-10 pointer-events-none" style="background-image: radial-gradient(circle, #cbd5e1 1px, transparent 1px); background-size: 40px 40px;"></div>
            
            <!-- Nodes & Edges Visual -->
            <div class="relative w-full h-full flex items-center justify-center">
              <!-- Central Hub -->
              <div class="w-32 h-32 bg-[#1E3A5F] rounded-full flex flex-col items-center justify-center text-white shadow-[0_0_100px_rgba(30,58,95,0.3)] z-20 border-8 border-white group cursor-pointer hover:scale-110 transition-transform duration-700">
                <LucideBuilding2 class="w-10 h-10 mb-2" />
                <span class="text-[9px] font-black uppercase tracking-widest">JKT-HUB</span>
              </div>

              <!-- Orbiting Warehouses (Connecting lines handled via SVG or CSS absolute positioning) -->
              <div v-for="(pos, i) in [
                { t: '10%', l: '20%', id: 'WH-A', val: '84%', color: 'text-red-500' },
                { t: '15%', l: '70%', id: 'WH-B', val: '12%', color: 'text-blue-500' },
                { t: '60%', l: '10%', id: 'WH-C', val: '42%', color: 'text-amber-500' },
                { t: '75%', l: '80%', id: 'WH-D', val: '68%', color: 'text-blue-500' }
              ]" :key="i" :style="{ top: pos.t, left: pos.l }" class="absolute w-48 h-48 flex flex-col items-center justify-center group cursor-pointer">
                <!-- Connector Line (Simplified mockup) -->
                <div class="absolute top-1/2 left-1/2 w-[300px] h-0.5 bg-slate-200 -z-10 origin-left" :style="{ transform: `rotate(${i * 90 + 45}deg)` }"></div>
                
                <div class="glass p-8 rounded-[2.5rem] border-2 border-white shadow-2xl flex flex-col items-center text-center space-y-4 group-hover:-translate-y-4 transition-all duration-500">
                  <div class="w-14 h-14 rounded-2xl bg-slate-50 flex items-center justify-center text-[#1E3A5F] shadow-inner mb-2">
                    <LucideWarehouse class="w-8 h-8" />
                  </div>
                  <div class="space-y-1">
                    <h4 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ pos.id }}</h4>
                    <p :class="`text-xl font-black ${pos.color} tracking-tighter`">{{ pos.val }}</p>
                    <p class="text-[8px] font-bold text-slate-300 uppercase tracking-widest">UTILIZATION</p>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Map Info Overlay -->
          <div class="absolute bottom-10 right-10 z-10 glass p-8 rounded-[2.5rem] w-80 space-y-6 shadow-2xl border border-white">
            <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.topology.map.selected') }}</h4>
            <div class="flex items-center gap-6">
              <div class="w-16 h-16 rounded-2xl bg-red-50 flex items-center justify-center text-red-500 shadow-inner">
                <LucideAlertCircle class="w-8 h-8" />
              </div>
              <div class="space-y-1">
                <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">WH-A SECTION 1</p>
                <p class="text-[9px] font-bold text-red-500 uppercase tracking-widest">{{ $t('warehouse.topology.map.alert') }}</p>
              </div>
            </div>
            <div class="pt-4 border-t border-slate-100 flex items-center justify-between">
              <span class="text-[9px] font-bold text-slate-400 uppercase">{{ $t('warehouse.topology.map.last_sync') }}</span>
              <span class="text-[9px] font-black text-[#1E3A5F] uppercase font-mono">14:22:15</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideActivity, LucideSettings, LucideSearch, LucideBuilding2, 
  LucideChevronRight, LucideWarehouse, LucideArrowRight, 
  LucideZoomIn, LucideZoomOut, LucideFocus, LucideAlertCircle 
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

