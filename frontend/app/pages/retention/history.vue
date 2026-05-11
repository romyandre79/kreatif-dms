<template>
  <div class="max-w-[1600px] mx-auto space-y-8 pb-20">
    <!-- Header -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6" v-motion-fade>
      <div class="space-y-2">
        <h1 class="text-3xl font-black text-[#1E3A5F] tracking-tight uppercase">{{ $t('retention.history.title') }}</h1>
        <p class="text-slate-500 font-medium">{{ $t('retention.history.subtitle') }}</p>
      </div>
      <div class="flex items-center gap-4">
        <button class="px-6 py-4 bg-white border border-slate-200 text-slate-400 rounded-2xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center gap-2">
          <LucideSearch class="w-4 h-4" /> {{ $t('retention.history.btn_search') }}
        </button>
      </div>
    </div>

    <!-- Active Missions Alert -->
    <div v-if="pendingCount > 0" class="p-8 bg-orange-50 border border-orange-100 rounded-3xl flex items-center justify-between shadow-xl shadow-orange-900/5" v-motion-slide-top>
      <div class="flex items-center gap-6">
        <div class="p-3 bg-orange-500 text-white rounded-2xl shadow-lg shadow-orange-500/20">
          <LucideAlertTriangle class="w-6 h-6" />
        </div>
        <div class="space-y-1">
          <h3 class="text-sm font-black text-orange-900 uppercase tracking-tight">{{ $t('retention.history.alert.title') }}</h3>
          <p class="text-[11px] font-bold text-orange-700/70 uppercase tracking-widest">{{ $t('retention.history.alert.desc', { count: pendingCount }) }}</p>
        </div>
      </div>
      <button @click="navigateTo('/retention/shredding')" class="px-6 py-3 bg-orange-500 hover:bg-orange-600 text-white rounded-xl text-[10px] font-black uppercase tracking-widest transition-all shadow-lg shadow-orange-500/20 flex items-center gap-2">
        {{ $t('retention.history.alert.btn_execute') }} <LucideArrowRight class="w-4 h-4" />
      </button>
    </div>

    <!-- Stats -->
    <div class="grid grid-cols-4 gap-8" v-motion-fade>
      <div v-for="stat in stats" :key="stat.label" class="glass p-8 rounded-lg bg-white border border-slate-100 shadow-xl shadow-slate-200/50 space-y-4">
        <p class="text-[9px] font-black text-slate-300 uppercase tracking-widest">{{ stat.label }}</p>
        <div class="flex items-end justify-between">
          <p class="text-3xl font-black text-[#1E3A5F]">{{ stat.value }}</p>
          <component :is="stat.icon" class="w-6 h-6 text-slate-100" />
        </div>
      </div>
    </div>

    <!-- History Table -->
    <div class="glass rounded-lg bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 overflow-hidden" v-motion-slide-visible-bottom>
      <div class="p-10 border-b border-slate-50 flex items-center justify-between">
        <h2 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('retention.history.table.title') }}</h2>
        <div class="flex items-center gap-4">
          <button class="p-3 text-slate-300 hover:text-primary-500 hover:bg-primary-50 rounded-2xl transition-all">
            <LucideFilter class="w-5 h-5" />
          </button>
        </div>
      </div>

      <div class="overflow-x-auto">
        <table class="w-full text-left">
          <thead>
            <tr class="text-[10px] font-black text-slate-300 uppercase tracking-widest border-b border-slate-50 bg-slate-50/30">
              <th class="px-10 py-6">{{ $t('retention.history.table.col_id') }}</th>
              <th class="px-6 py-6">{{ $t('retention.history.table.col_dept') }}</th>
              <th class="px-6 py-6">{{ $t('retention.history.table.col_items') }}</th>
              <th class="px-6 py-6">{{ $t('retention.history.table.col_date') }}</th>
              <th class="px-6 py-6">{{ $t('retention.history.table.col_status') }}</th>
              <th class="px-10 py-6 text-right">{{ $t('retention.history.table.col_actions') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50">
            <tr v-for="batch in history" :key="batch.id" class="group hover:bg-slate-50/50 transition-colors">
              <td class="px-10 py-8">
                <p class="text-sm font-black text-[#1E3A5F] tracking-tight group-hover:text-primary-600 transition-colors">{{ batch.id }}</p>
              </td>
              <td class="px-6 py-8">
                <p class="text-xs font-bold text-slate-500 uppercase">{{ batch.dept }}</p>
              </td>
              <td class="px-6 py-8">
                <p class="text-xs font-black text-slate-400 uppercase tracking-tight">{{ batch.items }} records</p>
              </td>
              <td class="px-6 py-8">
                <p class="text-xs font-bold text-slate-400 uppercase tracking-tight">{{ batch.date || '—' }}</p>
              </td>
              <td class="px-6 py-8">
                <span :class="`px-3 py-1 rounded-lg text-[9px] font-black uppercase tracking-widest border ${batch.statusColor}`">
                  {{ batch.status }}
                </span>
              </td>
              <td class="px-10 py-8 text-right">
                <div class="flex items-center justify-end gap-2">
                  <button v-if="batch.status === 'APPROVED'" @click="navigateTo('/retention/shredding')" 
                          class="px-4 py-2 bg-[#1E3A5F] text-white rounded-xl text-[9px] font-black uppercase tracking-widest shadow-lg shadow-blue-900/10 hover:bg-[#152943] transition-all">
                    {{ $t('retention.history.table.btn_execute') }}
                  </button>
                  <button class="p-3 text-slate-200 hover:text-primary-500 hover:bg-primary-50 rounded-xl transition-all">
                    <LucideEye class="w-4 h-4" />
                  </button>
                  <button v-if="batch.status === 'COMPLETED'" class="p-3 text-slate-200 hover:text-emerald-500 hover:bg-emerald-50 rounded-xl transition-all">
                    <LucideFileDown class="w-4 h-4" />
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideSearch, LucideAlertTriangle, LucideArrowRight, 
  LucideFilter, LucideEye, LucideFileDown, LucideTrash2,
  LucideFileCheck, LucideShieldAlert, LucideBox
} from 'lucide-vue-next'

const pendingCount = ref(1)

const stats = [
  { label: 'TOTAL PURGED', value: '14,204', icon: LucideTrash2 },
  { label: 'LAST MONTH', value: '1,240', icon: LucideFileCheck },
  { label: 'SECURITY HOLDS', value: '8', icon: LucideShieldAlert },
  { label: 'TOTAL BATCHES', value: '42', icon: LucideBox }
]

const history = ref([
  {
    id: 'BATCH-2024-001',
    dept: 'Finance',
    items: 1240,
    date: '',
    status: 'APPROVED',
    statusColor: 'bg-emerald-50 text-emerald-600 border-emerald-100'
  },
  {
    id: 'BATCH-2023-089',
    dept: 'Legal',
    items: 850,
    date: 'Dec 12, 2023',
    status: 'COMPLETED',
    statusColor: 'bg-slate-50 text-slate-400 border-slate-100'
  },
  {
    id: 'BATCH-2023-045',
    dept: 'HR & Talent',
    items: 320,
    date: 'Oct 05, 2023',
    status: 'COMPLETED',
    statusColor: 'bg-slate-50 text-slate-400 border-slate-100'
  }
])
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
</style>
