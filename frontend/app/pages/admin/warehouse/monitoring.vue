<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-1">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('warehouse.monitoring.title') }}</h1>
        <p class="text-xs font-bold text-slate-500">{{ $t('warehouse.monitoring.subtitle') }}</p>
      </div>
      <div class="flex items-center gap-4">
        <button class="px-6 py-3 text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest flex items-center gap-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl transition-all border border-slate-200 dark:border-slate-700">
          <LucideFilter class="w-4 h-4" />
          {{ $t('warehouse.monitoring.header.filter') }}
        </button>
        <button class="px-10 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 group">
          <LucideZap class="w-4 h-4 fill-current group-hover:scale-110 transition-transform" />
          {{ $t('warehouse.monitoring.header.generate') }}
        </button>
      </div>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="glass p-8 rounded-[2.5rem] space-y-4 shadow-sm border-t-8 border-blue-500">
        <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.monitoring.stats.total') }}</p>
        <div class="flex items-center gap-4">
          <p class="text-4xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">1,248</p>
          <span class="text-[8px] font-black text-green-500 uppercase tracking-widest">{{ $t('warehouse.monitoring.stats.mom', { val: 2 }) }}</span>
        </div>
      </div>
      <div class="glass p-8 rounded-[2.5rem] space-y-4 shadow-sm">
        <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.monitoring.stats.capacity') }}</p>
        <div class="flex items-end justify-between">
          <div class="space-y-1">
            <p class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">82.4%</p>
            <p class="text-[10px] font-bold text-slate-400 uppercase">/ 1,028 units</p>
          </div>
          <span class="px-3 py-1 bg-slate-50 dark:bg-slate-800 rounded text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.monitoring.stats.optimal') }}</span>
        </div>
      </div>
      <div class="glass p-8 rounded-[2.5rem] space-y-4 shadow-sm border-t-8 border-red-500">
        <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.monitoring.stats.reserved') }}</p>
        <div class="flex items-center gap-4">
          <div class="w-10 h-10 rounded-xl bg-red-50 text-red-500 flex items-center justify-center shadow-sm">
            <LucideAlertTriangle class="w-5 h-5" />
          </div>
          <div class="space-y-0.5">
            <p class="text-3xl font-black text-red-600 tracking-tighter">12</p>
            <p class="text-[8px] font-black text-red-400 uppercase tracking-widest">{{ $t('warehouse.monitoring.stats.critical') }}</p>
          </div>
        </div>
      </div>
      <div class="glass p-8 rounded-[2.5rem] space-y-4 shadow-sm relative overflow-hidden group">
        <LucideCompass class="absolute top-0 right-0 w-24 h-24 text-slate-100 dark:text-slate-800 -rotate-12 translate-x-6 -translate-y-6 group-hover:scale-110 transition-transform" />
        <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest z-10 relative">{{ $t('warehouse.monitoring.stats.accuracy') }}</p>
        <p class="text-4xl font-black text-[#1E3A5F] dark:text-white tracking-tighter z-10 relative">96.8%</p>
      </div>
    </div>

    <!-- Heatmap Grid -->
    <div class="glass p-10 rounded-[4rem] space-y-10 shadow-sm border border-slate-50 dark:border-slate-800">
      <div class="flex flex-col md:flex-row md:items-center justify-between gap-6 border-b border-slate-50 dark:border-slate-800 pb-10">
        <div class="space-y-1">
          <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('warehouse.monitoring.heatmap.title') }}</h3>
          <p class="text-[10px] font-bold text-slate-400 uppercase">{{ $t('warehouse.monitoring.heatmap.subtitle') }}</p>
        </div>
        <div class="flex flex-wrap items-center gap-6">
          <div v-for="l in ['empty', 'low', 'medium', 'high', 'reserved']" :key="l" class="flex items-center gap-3">
            <div :class="`w-4 h-4 rounded-md ${
              l === 'empty' ? 'bg-slate-100' : 
              l === 'low' ? 'bg-blue-100' : 
              l === 'medium' ? 'bg-blue-500' : 
              l === 'high' ? 'bg-[#1E3A5F]' : 
              'bg-red-500'
            }`"></div>
            <span class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t(`warehouse.monitoring.heatmap.legend.${l}`) }}</span>
          </div>
        </div>
      </div>

      <div class="grid grid-cols-8 md:grid-cols-12 lg:grid-cols-16 gap-3">
        <div v-for="i in 128" :key="i" :class="`aspect-square rounded-lg transition-all hover:scale-110 cursor-pointer shadow-sm ${
          [5, 12, 18, 45, 67, 88, 112].includes(i) ? 'bg-red-500 shadow-lg shadow-red-500/20' :
          [2, 22, 34, 56, 78, 90, 110].includes(i) ? 'bg-[#1E3A5F]' :
          [1, 10, 30, 50, 70, 90, 110].includes(i) ? 'bg-slate-100' :
          i % 3 === 0 ? 'bg-blue-500' : 'bg-blue-100'
        }`"></div>
      </div>
    </div>

    <!-- Recommendations Table -->
    <div class="glass rounded-[3rem] overflow-hidden shadow-sm border border-slate-50 dark:border-slate-800">
      <div class="p-10 border-b border-slate-50 dark:border-slate-800 space-y-1">
        <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('warehouse.monitoring.reco.title') }}</h3>
        <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest flex items-center gap-2">
          <LucideInfo class="w-3.5 h-3.5" />
          {{ $t('warehouse.monitoring.reco.logic') }}
        </p>
      </div>
      <table class="w-full text-left">
        <thead>
          <tr class="bg-slate-50/50 dark:bg-slate-800/50 text-[8px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-50 dark:border-slate-800">
            <th class="p-8 pl-10">{{ $t('warehouse.monitoring.reco.table.rack') }}</th>
            <th class="p-8">{{ $t('warehouse.monitoring.reco.table.dept') }}</th>
            <th class="p-8">{{ $t('warehouse.monitoring.reco.table.reason') }}</th>
            <th class="p-8 text-center">{{ $t('warehouse.monitoring.reco.table.score') }}</th>
            <th class="p-8 pr-10 text-right">{{ $t('warehouse.monitoring.reco.table.action') }}</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
          <tr v-for="r in [
            { id: 'A12-L4', zone: 'Premium Archive', dept: 'FINANCE & TAX', reason: 'Kedekatan dengan entry point utama; Optimal untuk dokumen audit.', score: 98, color: 'text-blue-500' },
            { id: 'C08-L2', zone: 'Reserved HR', dept: 'HUMAN RESOURCES', reason: 'Zonasi Departemen Sesuai; Mengurangi fragmen data.', score: 92, color: 'text-[#1E3A5F]' },
            { id: 'F04-L1', zone: 'Deep Storage', dept: 'LEGAL & CORP', reason: 'Dokumen inaktif; Penempatan di area sirkulasi rendah.', score: 85, color: 'text-slate-400' }
          ]" :key="r.id" class="group hover:bg-slate-50/30 transition-all cursor-pointer">
            <td class="p-8 pl-10">
              <div class="flex items-center gap-4">
                <div class="w-10 h-10 rounded-xl bg-slate-50 dark:bg-slate-800 flex items-center justify-center text-blue-500 font-black text-[10px] border border-slate-100 dark:border-slate-700">
                  {{ r.id.split('-')[0] }}
                </div>
                <div class="space-y-0.5">
                  <p class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Rack {{ r.id }}</p>
                  <p class="text-[8px] font-bold text-slate-400 uppercase tracking-widest">Zone: {{ r.zone }}</p>
                </div>
              </div>
            </td>
            <td class="p-8">
              <span class="px-4 py-1.5 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-700 rounded-full text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ r.dept }}</span>
            </td>
            <td class="p-8 text-[11px] font-bold text-slate-500 uppercase leading-relaxed max-w-md">
              {{ r.reason }}
            </td>
            <td class="p-8 text-center">
              <div class="space-y-2">
                <p class="text-lg font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ r.score }}%</p>
                <div class="w-20 mx-auto h-1 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
                  <div class="h-full bg-blue-500" :style="{ width: r.score + '%' }"></div>
                </div>
              </div>
            </td>
            <td class="p-8 pr-10 text-right">
              <button class="px-6 py-3 bg-[#1E3A5F] text-white rounded-xl text-[10px] font-black uppercase tracking-widest hover:bg-[#152943] transition-all shadow-lg shadow-blue-900/20 active:scale-95">
                {{ $t('warehouse.monitoring.reco.table.apply') }}
              </button>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="p-8 bg-slate-50/30 dark:bg-slate-800/30 text-center">
        <button class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest flex items-center gap-2 mx-auto hover:underline group">
          {{ $t('warehouse.monitoring.reco.view_all', { count: 24 }) }}
          <LucideArrowRight class="w-4 h-4 group-hover:translate-x-1 transition-transform" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideFilter, LucideZap, LucideAlertTriangle, LucideCompass, 
  LucideInfo, LucideArrowRight 
} from 'lucide-vue-next'

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}
</style>
