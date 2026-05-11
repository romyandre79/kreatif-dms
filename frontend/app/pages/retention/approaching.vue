<template>
  <div class="max-w-[1600px] mx-auto space-y-8 pb-20">
    <!-- Header -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6" v-motion-fade>
      <div class="space-y-2">
        <h1 class="text-3xl font-black text-[#1E3A5F] tracking-tight uppercase">{{ $t('retention.approaching.title') }}</h1>
        <p class="text-slate-500 font-medium">{{ $t('retention.approaching.subtitle') }}</p>
      </div>
      <div class="flex items-center gap-4">
        <button class="px-6 py-4 bg-white border border-slate-200 text-slate-400 rounded-2xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center gap-2">
          <LucideFilter class="w-4 h-4" />
          {{ $t('common.filter') || 'Filter' }}
        </button>
      </div>
    </div>

      <div class="glass p-10 rounded-lg bg-white border border-slate-100 shadow-xl shadow-slate-200/50 space-y-6 relative overflow-hidden group">
        <div class="absolute right-0 top-0 p-8 text-slate-50 group-hover:text-orange-50 transition-colors">
          <LucideHand class="w-20 h-20" />
        </div>
        <div class="flex items-center justify-between relative z-10">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">{{ $t('retention.approaching.stats.hold_title') }}</p>
          <LucidePauseCircle class="w-6 h-6 text-orange-500" />
        </div>
        <div class="space-y-1 relative z-10">
          <div class="flex items-end gap-4">
            <p class="text-5xl font-black text-[#1E3A5F]">8</p>
            <div class="flex items-center gap-1 text-red-500 font-black text-xs mb-2">
              -1% <LucideTrendingDown class="w-3 h-3" />
            </div>
          </div>
          <p class="text-xs font-bold text-slate-400">{{ $t('retention.approaching.stats.hold_desc') }}</p>
        </div>
      </div>

    <!-- Info Banner -->
    <div class="p-8 bg-[#1E3A5F] rounded-3xl border border-blue-900/20 flex items-start gap-6 shadow-xl shadow-blue-900/10" v-motion-fade>
      <div class="p-3 bg-white/10 rounded-2xl text-blue-200">
        <LucideInfo class="w-6 h-6" />
      </div>
      <p class="text-sm font-medium text-blue-50/90 leading-relaxed pt-1">
        {{ $t('retention.approaching.banner') }}
      </p>
    </div>

    <!-- Main List -->
    <div class="glass rounded-lg bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 overflow-hidden" v-motion-slide-visible-bottom>
      <div class="p-10 border-b border-slate-50 flex items-center justify-between">
        <div class="flex items-center gap-4">
          <h2 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('retention.approaching.table.title') }}</h2>
          <span class="px-3 py-1 bg-slate-50 text-slate-400 rounded-lg text-[9px] font-black uppercase tracking-widest border border-slate-100">{{ $t('retention.approaching.table.live_data') }}</span>
        </div>
        <div class="flex items-center gap-6">
          <div class="flex items-center gap-3">
            <div class="w-12 h-6 bg-primary-500 rounded-full relative cursor-pointer shadow-inner">
              <div class="absolute right-1 top-1 w-4 h-4 bg-white rounded-full shadow-sm"></div>
            </div>
            <span class="text-xs font-bold text-slate-500">{{ $t('retention.approaching.table.showing', { count: 10, total: 45 }) }}</span>
          </div>
        </div>
      </div>

      <div class="overflow-x-auto">
        <table class="w-full text-left">
          <thead>
            <tr class="text-[10px] font-black text-slate-300 uppercase tracking-widest border-b border-slate-50 bg-slate-50/30">
              <th class="px-10 py-6"><input type="checkbox" class="w-5 h-5 rounded border-2 border-slate-200" /></th>
              <th class="px-6 py-6">{{ $t('retention.approaching.table.id') }}</th>
              <th class="px-6 py-6">{{ $t('retention.approaching.table.doc_title') }}</th>
              <th class="px-6 py-6">{{ $t('retention.approaching.table.dept') }}</th>
              <th class="px-6 py-6">{{ $t('retention.approaching.table.age') }}</th>
              <th class="px-6 py-6">{{ $t('retention.approaching.table.rule') }}</th>
              <th class="px-6 py-6 text-center">{{ $t('retention.approaching.table.status') }}</th>
              <th class="px-10 py-6 text-right">{{ $t('retention.approaching.table.actions') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50">
            <tr v-for="doc in items" :key="doc.id" class="group hover:bg-slate-50/50 transition-colors">
              <td class="px-10 py-8">
                <input type="checkbox" :checked="doc.id === 'DOC-2019-452'" class="w-5 h-5 rounded border-2 border-slate-200 text-primary-500" />
              </td>
              <td class="px-6 py-8">
                <p class="text-sm font-black text-[#1E3A5F] tracking-tight">{{ doc.id }}</p>
              </td>
              <td class="px-6 py-8">
                <p class="text-sm font-bold text-slate-600 truncate max-w-[250px]">{{ doc.title }}</p>
              </td>
              <td class="px-6 py-8">
                <p class="text-xs font-bold text-slate-500 uppercase">{{ doc.dept }}</p>
              </td>
              <td class="px-6 py-8">
                <p class="text-xs font-black text-slate-400 uppercase tracking-tight">{{ doc.age }} {{ $t('retention.approaching.table.years') || 'Years' }}</p>
              </td>
              <td class="px-6 py-8">
                <p class="text-xs font-bold text-slate-500">{{ doc.rule }}</p>
              </td>
              <td class="px-6 py-8">
                <div class="flex justify-center">
                  <span :class="`px-3 py-1 rounded-lg text-[9px] font-black uppercase tracking-widest flex items-center gap-2 ${doc.statusColor}`">
                    <div v-if="doc.status === 'DUE FOR REVIEW'" class="w-1.5 h-1.5 rounded-full bg-primary-500"></div>
                    <div v-if="doc.status === 'ON HOLD'" class="w-1.5 h-1.5 rounded-full bg-orange-500"></div>
                    {{ doc.status }}
                  </span>
                </div>
              </td>
              <td class="px-10 py-8 text-right">
                <button class="p-3 text-slate-300 hover:text-primary-500 hover:bg-primary-50 rounded-2xl transition-all">
                  <LucideEye class="w-5 h-5" />
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="p-10 bg-slate-50/30 flex items-center justify-between border-t border-slate-50">
        <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ $t('retention.approaching.table.page', { current: 1, total: 5 }) || 'Page 1 of 5' }}</p>
        <div class="flex items-center gap-2">
          <button class="px-4 py-2 border border-slate-200 text-slate-300 rounded-xl text-[10px] font-black uppercase tracking-widest hover:bg-white transition-all disabled:opacity-50" disabled>{{ $t('retention.approaching.table.prev') || 'Prev' }}</button>
          <button class="w-10 h-10 bg-[#1E3A5F] text-white rounded-xl text-xs font-black">1</button>
          <button class="w-10 h-10 bg-white border border-slate-200 text-slate-400 rounded-xl text-xs font-black hover:bg-slate-50">2</button>
          <button class="px-4 py-2 border border-slate-200 text-slate-400 rounded-xl text-[10px] font-black uppercase tracking-widest hover:bg-white transition-all">{{ $t('retention.approaching.table.next') || 'Next' }}</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const { t } = useI18n()
import { ref } from 'vue'
import { 
  LucideFilter, LucideCheckCircle2, LucideClipboardList, 
  LucideAlertTriangle, LucidePauseCircle, LucideTrendingUp, 
  LucideTrendingDown, LucideInfo, LucideEye, LucideFileClock,
  LucideAlertCircle, LucideHand
} from 'lucide-vue-next'

const items = ref([
  {
    id: 'DOC-2023-001',
    title: 'Quarterly Financial Statement Q1 2023',
    dept: 'Finance',
    age: 7.1,
    rule: '7Y Retention (FIN-01)',
    status: t('retention.approaching.stats.overdue_title'),
    statusColor: 'bg-red-50 text-red-500 border border-red-100'
  },
  {
    id: 'DOC-2019-452',
    title: 'Employee Contract - J. Doe',
    dept: 'HR & Talent',
    age: 4.9,
    rule: '5Y Retention (HR-04)',
    status: t('retention.approaching.stats.due_title'),
    statusColor: 'bg-blue-50 text-primary-600 border border-blue-100'
  },
  {
    id: 'DOC-2021-118',
    title: 'Vendor Agreement - TechCorp',
    dept: 'Legal',
    age: 3.0,
    rule: '3Y Retention (LEG-02)',
    status: t('retention.approaching.stats.hold_title'),
    statusColor: 'bg-orange-50 text-orange-600 border border-orange-100'
  },
  {
    id: 'DOC-2023-882',
    title: 'Compliance Audit Report 2023',
    dept: 'Compliance',
    age: 1.0,
    rule: '1Y Retention (COM-01)',
    status: t('retention.approaching.stats.due_title'),
    statusColor: 'bg-blue-50 text-primary-600 border border-blue-100'
  }
])
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
</style>
