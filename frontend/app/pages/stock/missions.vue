<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Page Header -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-1">
        <h1 class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
          {{ $t('stock.missions.title') }} 
          <span class="text-slate-400 font-bold normal-case text-lg">{{ $t('stock.missions.head_view') }}</span>
        </h1>
        <p class="text-xs font-bold text-slate-500">{{ $t('stock.missions.subtitle') }}</p>
      </div>
      
      <div class="flex items-center bg-white dark:bg-slate-900 p-1 rounded-xl shadow-sm border border-slate-100 dark:border-slate-800">
        <button v-for="t in ['all', 'scheduled', 'completed']" :key="t" :class="`px-6 py-2.5 rounded-lg text-xs font-black uppercase tracking-widest transition-all ${activeTab === t ? 'bg-[#1E3A5F] text-white shadow-lg' : 'text-slate-400 hover:text-slate-600'}`" @click="activeTab = t">
          {{ $t(`stock.missions.tabs.${t}`) }}
        </button>
      </div>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
      <div class="glass p-8 rounded-lg flex items-center justify-between group hover:border-primary-500/30 transition-all">
        <div class="space-y-1">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('stock.missions.stats.pending') }}</p>
          <div class="flex items-baseline gap-3">
            <span class="text-3xl font-black text-[#1E3A5F] dark:text-white">12</span>
            <span class="text-[9px] font-black text-orange-500 uppercase tracking-widest">{{ $t('stock.missions.stats.approval_req') }}</span>
          </div>
        </div>
        <div class="w-12 h-12 rounded-2xl bg-orange-50 dark:bg-orange-900/20 text-orange-500 flex items-center justify-center"><LucideClipboardList class="w-6 h-6" /></div>
      </div>
      
      <div class="glass p-8 rounded-lg flex items-center justify-between group hover:border-primary-500/30 transition-all">
        <div class="space-y-1">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('stock.missions.stats.active') }}</p>
          <div class="flex items-baseline gap-3">
            <span class="text-3xl font-black text-[#1E3A5F] dark:text-white">8</span>
            <span class="text-[9px] font-black text-blue-500 uppercase tracking-widest">{{ $t('stock.missions.stats.on_mission') }}</span>
          </div>
        </div>
        <div class="w-12 h-12 rounded-2xl bg-blue-50 dark:bg-blue-900/20 text-blue-500 flex items-center justify-center"><LucideUsers class="w-6 h-6" /></div>
      </div>

      <div class="glass p-8 rounded-lg flex items-center justify-between group hover:border-primary-500/30 transition-all">
        <div class="space-y-1">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('stock.missions.stats.rate') }}</p>
          <div class="flex items-baseline gap-3">
            <span class="text-3xl font-black text-[#1E3A5F] dark:text-white">94%</span>
            <span class="text-[9px] font-black text-green-500 uppercase tracking-widest">{{ $t('stock.missions.stats.vs_last_week', { value: '+2.4%' }) }}</span>
          </div>
        </div>
        <div class="w-12 h-12 rounded-2xl bg-green-50 dark:bg-green-900/20 text-green-500 flex items-center justify-center"><LucideCheckCircle2 class="w-6 h-6" /></div>
      </div>
    </div>

    <!-- Main Table Section -->
    <div class="glass rounded-lg overflow-hidden" v-motion-slide-visible-bottom>
      <div class="p-8 border-b border-slate-100 dark:border-slate-800 flex flex-col md:flex-row md:items-center justify-between gap-6 bg-white/30">
        <div class="relative flex-grow max-w-md">
          <LucideSearch class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
          <input type="text" :placeholder="$t('stock.missions.search_placeholder')" class="w-full pl-11 pr-5 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all" />
        </div>
        <div class="flex items-center gap-3">
          <button class="px-6 py-3 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-[#1E3A5F] dark:text-white rounded-xl text-xs font-black uppercase tracking-widest flex items-center gap-2 hover:bg-slate-50 transition-all shadow-sm">
            <LucideFilter class="w-4 h-4" /> Filter
          </button>
          <button class="px-6 py-3 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-[#1E3A5F] dark:text-white rounded-xl text-xs font-black uppercase tracking-widest flex items-center gap-2 hover:bg-slate-50 transition-all shadow-sm">
            <LucideDownload class="w-4 h-4" /> Export
          </button>
        </div>
      </div>

      <div class="overflow-x-auto">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800 bg-slate-50/30 dark:bg-slate-900/30">
              <th class="p-8 px-10">{{ $t('stock.missions.table.cols.target') }}</th>
              <th class="p-8">{{ $t('stock.missions.table.cols.schedule') }}</th>
              <th class="p-8">{{ $t('stock.missions.table.cols.operator') }}</th>
              <th class="p-8 text-center">{{ $t('stock.missions.table.cols.type') }}</th>
              <th class="p-8">{{ $t('stock.missions.table.cols.status') }}</th>
              <th class="p-8 text-center px-10">{{ $t('stock.missions.table.cols.actions') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50 dark:divide-slate-800/50">
            <tr v-for="mission in missions" :key="mission.id" class="hover:bg-slate-50/50 dark:hover:bg-slate-800/20 transition-all group">
              <td class="p-8 px-10">
                <div class="space-y-1">
                  <p class="text-sm font-black text-[#1E3A5F] dark:text-white tracking-tight">{{ mission.target }}</p>
                  <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ mission.subtarget }}</p>
                </div>
              </td>
              <td class="p-8">
                <div class="flex items-center gap-3">
                  <LucideCalendar :class="`w-4 h-4 ${mission.isLate ? 'text-orange-500' : 'text-slate-300'}`" />
                  <p :class="`text-xs font-black ${mission.isLate ? 'text-orange-600' : 'text-slate-600 dark:text-slate-300'}`">Deadline {{ mission.deadline }}</p>
                </div>
              </td>
              <td class="p-8">
                <div class="flex items-center gap-3">
                  <div class="w-8 h-8 rounded-full bg-primary-100 dark:bg-primary-900/30 flex items-center justify-center text-[10px] font-black text-primary-600">
                    {{ mission.operatorCode }}
                  </div>
                  <p class="text-xs font-black text-slate-700 dark:text-slate-200">{{ mission.operator }}</p>
                </div>
              </td>
              <td class="p-8 text-center">
                <span :class="`px-3 py-1 rounded-md text-[8px] font-black tracking-widest ${mission.type === 'SCHEDULED' ? 'bg-blue-50 text-blue-500' : 'bg-purple-50 text-purple-500'}`">
                  {{ mission.type }}
                </span>
              </td>
              <td class="p-8">
                <div class="flex items-center gap-2">
                  <span :class="`w-2 h-2 rounded-full ${mission.status === 'In Progress' ? 'bg-blue-500 shadow-[0_0_8px_rgba(59,130,246,0.5)]' : 'bg-slate-300'}`"></span>
                  <span :class="`text-xs font-black ${mission.status === 'In Progress' ? 'text-blue-600' : 'text-slate-400'}`">{{ mission.status }}</span>
                </div>
              </td>
              <td class="p-8 text-center px-10">
                <button class="p-2 text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 transition-colors"><LucideMoreVertical class="w-5 h-5" /></button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="p-8 border-t border-slate-100 dark:border-slate-800 flex items-center justify-between bg-slate-50/10">
        <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('stock.missions.pagination', { start: 1, end: 3, total: 12 }) }}</p>
        <div class="flex items-center gap-2">
          <button class="w-8 h-8 flex items-center justify-center rounded-lg border border-slate-100 dark:border-slate-800 text-slate-400 hover:bg-slate-50 transition-all"><LucideChevronLeft class="w-4 h-4" /></button>
          <button class="w-8 h-8 flex items-center justify-center rounded-lg bg-[#1E3A5F] text-white text-[10px] font-black shadow-lg">1</button>
          <button class="w-8 h-8 flex items-center justify-center rounded-lg border border-slate-100 dark:border-slate-800 text-slate-400 text-[10px] font-black hover:bg-slate-50 transition-all">2</button>
          <button class="w-8 h-8 flex items-center justify-center rounded-lg border border-slate-100 dark:border-slate-800 text-slate-400 text-[10px] font-black hover:bg-slate-50 transition-all">3</button>
          <button class="w-8 h-8 flex items-center justify-center rounded-lg border border-slate-100 dark:border-slate-800 text-slate-400 hover:bg-slate-50 transition-all"><LucideChevronRight class="w-4 h-4" /></button>
        </div>
      </div>
    </div>

    <!-- Tip Box -->
    <div class="p-10 bg-primary-50/30 dark:bg-primary-900/10 border border-primary-100/50 dark:border-primary-800/30 rounded-lg flex gap-8" v-motion-slide-visible-bottom>
      <div class="w-14 h-14 rounded-2xl bg-white dark:bg-slate-800 text-primary-500 flex items-center justify-center shadow-sm shrink-0">
        <LucideInfo class="w-7 h-7" />
      </div>
      <div class="space-y-2">
        <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('stock.missions.tip.title') }}</h4>
        <p class="text-xs font-bold text-slate-500 dark:text-slate-400 leading-relaxed max-w-4xl">
          {{ $t('stock.missions.tip.desc') }}
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideClipboardList, LucideUsers, LucideCheckCircle2, 
  LucideSearch, LucideFilter, LucideDownload, LucideCalendar, 
  LucideMoreVertical, LucideChevronLeft, LucideChevronRight, LucideInfo
} from 'lucide-vue-next'

const activeTab = ref('all')

const missions = [
  { id: 1, target: 'Zonation Dept. Legal', subtarget: 'Racks A1 to A5', deadline: '2026-04-15', operator: 'Admin Budi', operatorCode: 'AB', type: 'SCHEDULED', status: 'Not Started', isLate: false },
  { id: 2, target: 'Warehouse North', subtarget: 'Pallet Zone P9-P12', deadline: '2026-04-12', operator: 'Siti Dahlan', operatorCode: 'SD', type: 'SPOT CHECK', status: 'In Progress', isLate: true },
  { id: 3, target: 'IT Asset Storage', subtarget: 'Rack C3 (High Value)', deadline: '2026-04-20', operator: 'Reza Kurnia', operatorCode: 'RK', type: 'SCHEDULED', status: 'Not Started', isLate: false }
]
</script>

<style scoped>
@reference "tailwindcss";
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}
</style>

