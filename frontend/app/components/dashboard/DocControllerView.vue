<template>
  <div class="space-y-8 pb-10">
    <!-- Header Section -->
    <PageHeader 
      :title="$t('dashboard.welcome', { name: user?.full_name || 'User' })"
      :subtitle="$t('dashboard.header.enterprise_dashboard')"
    >
      <template #actions>
        <div class="flex items-center gap-3">
          <button 
            @click="navigateTo('/documents/upload')"
            class="flex items-center gap-2 px-6 py-2.5 bg-[#1E3A5F] hover:bg-[#152943] text-white font-bold rounded-xl shadow-lg shadow-blue-900/20 transition-all mr-2"
          >
            <LucidePlus class="w-4 h-4" />
            {{ $t('dashboard.header.btn_submit_new') }}
          </button>
          <div class="bg-slate-100 dark:bg-slate-800 px-4 py-2 rounded-xl flex items-center gap-2 border border-slate-200 dark:border-slate-700">
            <LucideDatabase class="w-4 h-4 text-slate-500" />
            <span class="text-xs font-black text-slate-700 dark:text-slate-300 uppercase tracking-tighter">{{ $t('dashboard.stats.storage_percent', { percent: 72 }) }}</span>
          </div>
          <div class="bg-orange-50 dark:bg-orange-900/20 px-4 py-2 rounded-xl flex items-center gap-2 border border-orange-100 dark:border-orange-800">
            <LucideZap class="w-4 h-4 text-orange-500" />
            <span class="text-xs font-black text-orange-700 dark:text-orange-400 uppercase tracking-tighter">{{ $t('dashboard.stats.active_tasks_count', { count: 31 }) }}</span>
          </div>
        </div>
      </template>
    </PageHeader>

    <!-- Top Stats -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- Inbound Pipeline -->
      <div v-motion-slide-visible-bottom class="glass p-6 rounded-2xl">
        <div class="flex items-start justify-between mb-4">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('dashboard.stats.inbound_pipeline') }}</p>
          <LucideArrowRightCircle class="w-5 h-5 text-slate-300" />
        </div>
        <p class="text-3xl font-black text-slate-900 dark:text-white">{{ $t('dashboard.stats.units', { count: '1,248' }) }}</p>
        <p class="text-xs font-bold text-green-500 mt-2">{{ $t('dashboard.stats.growth', { percent: 12 }) }}</p>
      </div>

      <!-- Loan Operations -->
      <div v-motion-slide-visible-bottom class="glass p-6 rounded-2xl">
        <div class="flex items-start justify-between mb-4">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('dashboard.stats.loans') }}</p>
          <LucidePackageCheck class="w-5 h-5 text-orange-400" />
        </div>
        <p class="text-3xl font-black text-slate-900 dark:text-white">{{ $t('dashboard.stats.active', { count: 42 }) }}</p>
        <p class="text-xs font-bold text-red-500 mt-2">{{ $t('dashboard.stats.overdue', { count: 8 }) }}</p>
      </div>

      <!-- WH Capacity -->
      <div v-motion-slide-visible-bottom class="glass p-6 rounded-2xl flex items-center gap-6">
        <div class="relative w-16 h-16 flex items-center justify-center">
          <svg class="w-full h-full transform -rotate-90">
            <circle cx="32" cy="32" r="28" fill="transparent" stroke="currentColor" stroke-width="6" class="text-slate-100 dark:text-slate-800" />
            <circle cx="32" cy="32" r="28" fill="transparent" stroke="currentColor" stroke-width="6" stroke-dasharray="176" stroke-dashoffset="49" class="text-primary-500" />
          </svg>
          <span class="absolute text-[10px] font-black">72%</span>
        </div>
        <div>
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-1">{{ $t('dashboard.stats.warehouse_capacity') }}</p>
          <p class="text-lg font-black text-slate-900 dark:text-white">18.4TB / 25TB</p>
        </div>
      </div>

      <!-- System Health -->
      <div v-motion-slide-visible-bottom class="glass p-6 rounded-2xl">
        <div class="flex items-start justify-between mb-4">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('dashboard.stats.system_health') }}</p>
          <LucideCheckCircle2 class="w-5 h-5 text-green-500" />
        </div>
        <p class="text-2xl font-black text-slate-900 dark:text-white">{{ $t('dashboard.stats.all_nominal') }}</p>
        <p class="text-[10px] font-bold text-slate-400 mt-1 uppercase">{{ $t('dashboard.stats.uptime', { percent: 99.9 }) }}</p>
      </div>
    </div>

    <!-- Priority Operation Queue -->
    <div class="glass rounded-2xl overflow-hidden" v-motion-slide-visible-bottom>
      <div class="px-8 py-6 flex items-center justify-between">
        <div class="flex items-center gap-2">
          <LucideAlertCircle class="w-5 h-5 text-red-500" />
          <h3 class="font-black text-lg text-[#1E3A5F] dark:text-white">{{ $t('dashboard.queue.priority_operations') }}</h3>
        </div>
        <button class="text-primary-600 font-black text-sm hover:underline">{{ $t('dashboard.queue.view_all_queue') }}</button>
      </div>
      
      <!-- Tabs -->
      <div class="px-8 flex border-b border-slate-100 dark:border-slate-800 gap-8">
        <button v-for="tab in queueTabs" :key="tab.label" :class="`pb-4 text-xs font-black uppercase tracking-widest transition-all relative ${tab.active ? 'text-primary-600' : 'text-slate-400 hover:text-slate-600'}`">
          {{ tab.label }} ({{ tab.count }})
          <div v-if="tab.active" class="absolute bottom-0 left-0 w-full h-1 bg-primary-600 rounded-t-full"></div>
        </button>
      </div>

      <div class="overflow-x-auto">
        <table class="w-full text-left">
          <thead>
            <tr class="text-slate-400 text-[10px] uppercase font-black tracking-widest bg-slate-50/50 dark:bg-slate-900/50">
              <th class="px-8 py-4">{{ $t('dashboard.queue.table.ref_id') }}</th>
              <th class="px-8 py-4">{{ $t('dashboard.queue.table.client_dept') }}</th>
              <th class="px-8 py-4">{{ $t('dashboard.queue.table.type') }}</th>
              <th class="px-8 py-4">{{ $t('dashboard.queue.table.sla_status') }}</th>
              <th class="px-8 py-4 text-right">{{ $t('dashboard.queue.table.action') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-100 dark:divide-slate-800">
            <tr v-for="item in queueItems" :key="item.id" class="hover:bg-slate-50/50 dark:hover:bg-slate-900/50 transition-colors">
              <td class="px-8 py-6 font-black text-sm text-slate-800 dark:text-slate-200">{{ item.id }}</td>
              <td class="px-8 py-6">
                <p class="text-sm font-black text-slate-700 dark:text-slate-300">{{ item.client }}</p>
                <p class="text-[10px] text-slate-400 font-bold uppercase">{{ item.dept }}</p>
              </td>
              <td class="px-8 py-6 text-sm font-bold text-slate-500">{{ item.type }}</td>
              <td class="px-8 py-6">
                <span :class="`px-3 py-1 rounded-lg text-[9px] font-black uppercase ${item.slaBg}`">
                  {{ item.sla }}
                </span>
              </td>
              <td class="px-8 py-6 text-right">
                <button :class="`px-5 py-2.5 rounded-xl text-xs font-black transition-all ${item.btnClass}`">
                  {{ item.btnLabel }}
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Activity and Alerts Row -->
    <div class="grid grid-cols-1 xl:grid-cols-3 gap-8">
      <!-- Activity Feed -->
      <div class="xl:col-span-2 glass rounded-2xl p-8" v-motion-slide-visible-bottom>
        <h3 class="font-black text-xl text-[#1E3A5F] dark:text-white mb-8">{{ $t('dashboard.activity.title_today') }}</h3>
        <div class="space-y-8">
          <div v-for="(act, index) in activities" :key="index" class="flex gap-6 relative">
            <div :class="`w-3 h-3 rounded-full mt-1 flex-shrink-0 ${act.dot}`"></div>
            <div>
              <p class="text-sm font-bold text-slate-700 dark:text-slate-200 leading-relaxed" v-html="act.text"></p>
              <div class="flex items-center gap-3 mt-1">
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-tighter" v-html="act.sub"></p>
                <span class="w-1 h-1 rounded-full bg-slate-200"></span>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-tighter">{{ act.time }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Alerts & Warnings -->
      <div class="glass rounded-2xl p-8" v-motion-slide-visible-bottom>
        <div class="flex items-center justify-between mb-8">
          <h3 class="font-black text-xl text-[#1E3A5F] dark:text-white">{{ $t('dashboard.alerts.title') }}</h3>
          <span class="bg-red-500 text-white text-[9px] font-black uppercase px-2 py-1 rounded-md">{{ $t('dashboard.alerts.action_required', { count: 6 }) }}</span>
        </div>
        <div class="space-y-4">
          <div v-for="alert in alerts" :key="alert.title" :class="`p-5 rounded-2xl border-l-4 ${alert.bg} ${alert.border}`">
            <div class="flex gap-4">
              <component :is="alert.icon" :class="`w-5 h-5 flex-shrink-0 ${alert.iconColor}`" />
              <div>
                <p :class="`text-sm font-black ${alert.titleColor}`">{{ alert.title }}</p>
                <p :class="`text-[11px] mt-1 font-bold leading-relaxed ${alert.descColor}`">{{ alert.desc }}</p>
              </div>
            </div>
          </div>
          
          <!-- Sync Status -->
          <div class="flex items-center gap-4 p-5">
            <LucideRefreshCcw class="w-5 h-5 text-slate-400 animate-spin-slow" />
            <div>
              <p class="text-sm font-black text-slate-700 dark:text-slate-300">{{ $t('dashboard.alerts.sync_status') }}</p>
              <p class="text-[11px] text-slate-400 font-bold leading-relaxed">{{ $t('dashboard.alerts.sync_desc', { count: 12 }) }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Quick Access Tools -->
    <div class="glass rounded-2xl p-10" v-motion-slide-visible-bottom>
      <h3 class="text-xs font-black uppercase tracking-widest text-slate-400 mb-10">{{ $t('dashboard.tools.title') }}</h3>
      <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-6">
        <button v-for="tool in tools" :key="tool.label" class="flex flex-col items-center group">
          <div class="w-16 h-16 rounded-2xl bg-white dark:bg-slate-900 border border-slate-100 dark:border-slate-800 shadow-sm flex items-center justify-center text-slate-400 group-hover:bg-primary-500 group-hover:text-white transition-all group-hover:-translate-y-1 mb-4">
            <component :is="tool.icon" class="w-7 h-7" />
          </div>
          <span class="text-[11px] font-black text-slate-600 dark:text-slate-400 uppercase tracking-tighter">{{ tool.label }}</span>
        </button>
      </div>
    </div>

    <!-- Footer Stats -->
    <div class="grid grid-cols-2 lg:grid-cols-4 gap-8 px-4" v-motion-fade>
      <div v-for="stat in footerStats" :key="stat.label" class="text-center">
        <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">{{ stat.label }}</p>
        <p class="text-2xl font-black text-slate-900 dark:text-white">{{ stat.value }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideDatabase, 
  LucideZap, 
  LucideArrowRightCircle, 
  LucidePackageCheck,
  LucideCheckCircle2,
  LucideAlertCircle,
  LucidePrinter,
  LucideScan,
  LucideClipboardList,
  LucideMap,
  LucideUsers,
  LucideHelpCircle,
  LucideRefreshCcw,
  LucidePlus
} from 'lucide-vue-next'
import { useAuthStore } from '~/stores/auth'
import PageHeader from '~/components/PageHeader.vue'

const auth = useAuthStore()
const { t } = useI18n()
const user = computed(() => auth.user)

const queueTabs = [
  { label: t('dashboard.queue.tabs.all'), count: 31, active: true },
  { label: t('dashboard.queue.tabs.receive'), count: 12, active: false },
  { label: t('dashboard.queue.tabs.scan'), count: 8, active: false },
  { label: t('dashboard.queue.tabs.metadata'), count: 6, active: false },
  { label: t('dashboard.queue.tabs.loan_prep'), count: 5, active: false },
]

const queueItems = [
  { id: '#REQ-2023-0891', client: 'Legal Department', dept: 'Confidential Contracts', type: t('dashboard.queue.items.bulk_scan'), sla: t('dashboard.queue.items.overdue_2h'), slaBg: 'bg-red-100 text-red-600', btnLabel: t('dashboard.queue.items.start_scanning'), btnClass: 'bg-[#1E3A5F] text-white hover:bg-[#152943]' },
  { id: '#PKG-4501-A', client: 'Corporate Finance', dept: 'Physical Invoices', type: t('dashboard.queue.items.receive'), sla: t('dashboard.queue.items.due_45m'), slaBg: 'bg-orange-100 text-orange-600', btnLabel: t('dashboard.queue.items.receive_pkg'), btnClass: 'bg-blue-100 text-blue-700 hover:bg-blue-200' },
  { id: '#MTD-8821-C', client: 'Human Resources', dept: 'Employee Records', type: t('dashboard.queue.items.indexing'), sla: t('dashboard.queue.items.on_schedule'), slaBg: 'bg-slate-100 text-slate-500', btnLabel: t('dashboard.queue.items.apply_metadata'), btnClass: 'bg-blue-100 text-blue-700 hover:bg-blue-200' },
]

const activities = [
  { text: t('dashboard.activity.items.ocr_completed', { id: '772' }), sub: t('dashboard.activity.items.ocr_sub', { count: '14,201' }), time: t('dashboard.activity.items.mins_ago', { count: 2 }), dot: 'bg-blue-500' },
  { text: t('dashboard.activity.items.box_received', { name: 'Budi S.', id: 'BX-902' }), sub: t('dashboard.activity.items.box_sub', { loc: 'D-12' }), time: t('dashboard.activity.items.mins_ago', { count: 15 }), dot: 'bg-green-500' },
  { text: t('dashboard.activity.items.loan_pending', { id: 'LR-202' }), sub: t('dashboard.activity.items.loan_sub', { name: 'Diana', dept: 'Marketing' }), time: t('dashboard.activity.items.mins_ago', { count: 45 }), dot: 'bg-orange-500' },
  { text: t('dashboard.activity.items.security_login'), sub: t('dashboard.activity.items.security_sub', { ip: '192.168.1.45' }), time: t('dashboard.activity.items.hour_ago'), dot: 'bg-slate-300' },
]

const alerts = [
  { title: 'Overdue Retention Review', desc: '128 records are past disposal date. Action required by Legal compliance.', icon: LucideAlertCircle, bg: 'bg-red-50 dark:bg-red-900/10', border: 'border-red-500', iconColor: 'text-red-500', titleColor: 'text-red-900 dark:text-red-200', descColor: 'text-red-700 dark:text-red-400' },
  { title: 'Storage Threshold Warning', desc: "Cloud Bucket 'DMS-Primary' is at 92% capacity. Consider scaling.", icon: LucideDatabase, bg: 'bg-orange-50 dark:bg-orange-900/10', border: 'border-orange-500', iconColor: 'text-orange-500', titleColor: 'text-orange-900 dark:text-orange-200', descColor: 'text-orange-700 dark:text-orange-400' },
]

const tools = [
  { label: t('dashboard.tools.label_printing'), icon: LucidePrinter },
  { label: t('dashboard.tools.quick_scan'), icon: LucideScan },
  { label: t('dashboard.tools.audit_logs'), icon: LucideClipboardList },
  { label: t('dashboard.tools.inventory_map'), icon: LucideMap },
  { label: t('dashboard.tools.user_access'), icon: LucideUsers },
  { label: t('dashboard.tools.it_support'), icon: LucideHelpCircle },
]

const footerStats = [
  { label: t('dashboard.footer_stats.throughput'), value: '4.2k pgs/hr' },
  { label: t('dashboard.footer_stats.turnaround'), value: '5.8 hrs' },
  { label: t('dashboard.footer_stats.ocr_accuracy'), value: '99.2%' },
  { label: t('dashboard.footer_stats.sla_satisfaction'), value: '94.8%' },
]
</script>

<style scoped>
.animate-spin-slow {
  animation: spin 8s linear infinite;
}
@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}
</style>
