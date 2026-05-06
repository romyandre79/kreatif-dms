<template>
  <div class="space-y-8 pb-10">
    <PageHeader 
      :title="$t('dashboard.welcome', { name: user?.full_name?.split(' ')[0] || 'Manager' })"
      :subtitle="$t('dashboard.manager.alert_desc', { count: statsSummary?.active_tasks || 0 })"
    />
    <!-- Urgent Alert Banner -->
    <div v-if="statsSummary?.active_tasks > 0" v-motion-fade class="bg-orange-50 dark:bg-orange-900/20 border-l-4 border-orange-500 p-6 rounded-r-2xl flex items-center justify-between shadow-sm">
      <div class="flex items-center gap-4">
        <div class="w-10 h-10 rounded-full bg-orange-100 dark:bg-orange-800 flex items-center justify-center text-orange-600 dark:text-orange-300">
          <LucideAlertTriangle class="w-6 h-6" />
        </div>
        <div>
          <h3 class="font-black text-orange-900 dark:text-orange-200">{{ $t('dashboard.manager.alert_title') }}</h3>
          <p class="text-sm text-orange-700/80 dark:text-orange-400/80">{{ $t('dashboard.manager.alert_desc', { count: statsSummary?.active_tasks }) }}</p>
        </div>
      </div>
      <button class="px-6 py-2 bg-orange-600 hover:bg-orange-700 text-white font-black text-xs uppercase tracking-widest rounded-xl shadow-lg shadow-orange-600/30 transition-all">
        {{ $t('dashboard.manager.btn_review_all') }}
      </button>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div v-for="(stat, index) in stats" :key="index" v-motion-slide-visible-bottom class="glass p-6 rounded-2xl relative overflow-hidden group">
        <p class="text-xs text-slate-400 font-bold uppercase tracking-wider mb-4">{{ stat.label }}</p>
        <div class="flex items-end justify-between mb-4">
          <p class="text-4xl font-black text-slate-900 dark:text-white">{{ stat.value }}</p>
          <div :class="`flex items-center gap-1 text-xs font-bold ${stat.trendUp ? 'text-green-500' : 'text-red-500'}`">
            <component :is="stat.trendUp ? LucideTrendingUp : LucideTrendingDown" class="w-3 h-3" />
            {{ stat.trend }}
          </div>
        </div>
        <div :class="`h-1.5 w-full rounded-full bg-slate-100 dark:bg-slate-800 overflow-hidden`">
          <div :class="`h-full rounded-full ${stat.barColor}`" :style="{ width: stat.progress + '%' }"></div>
        </div>
      </div>
    </div>

    <!-- Middle Content -->
    <div class="grid grid-cols-1 xl:grid-cols-3 gap-8">
      <!-- Urgent Approval Queue -->
      <div class="xl:col-span-2 glass rounded-2xl overflow-hidden" v-motion-slide-visible-bottom>
        <div class="px-8 py-6 border-b border-slate-200 dark:border-slate-800 flex items-center justify-between">
          <h3 class="font-black text-lg text-slate-800 dark:text-white">{{ $t('dashboard.manager.queue.title') }}</h3>
          <div class="flex bg-slate-100 dark:bg-slate-800 p-1 rounded-xl">
            <button class="px-4 py-1.5 text-xs font-bold rounded-lg bg-white dark:bg-slate-700 shadow-sm">{{ $t('dashboard.manager.queue.tabs.urgent') }}</button>
            <button class="px-4 py-1.5 text-xs font-bold text-slate-500 hover:text-slate-700 transition-colors">{{ $t('dashboard.manager.queue.tabs.normal') }}</button>
            <button class="px-4 py-1.5 text-xs font-bold text-slate-500 hover:text-slate-700 transition-colors">{{ $t('dashboard.manager.queue.tabs.all') }}</button>
          </div>
        </div>
        <div class="overflow-x-auto">
          <table class="w-full text-left">
            <thead>
              <tr class="text-slate-400 text-[10px] uppercase font-black tracking-widest bg-slate-50/50 dark:bg-slate-900/50">
                <th class="px-8 py-4">{{ $t('dashboard.manager.queue.table.submission') }}</th>
                <th class="px-8 py-4">{{ $t('dashboard.manager.queue.table.staff_name') }}</th>
                <th class="px-8 py-4">{{ $t('dashboard.manager.queue.table.days_elapsed') }}</th>
                <th class="px-8 py-4">{{ $t('dashboard.manager.queue.table.status') }}</th>
                <th class="px-8 py-4 text-right">{{ $t('dashboard.manager.queue.table.actions') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-100 dark:divide-slate-800">
              <tr v-for="item in dashboardData?.tasks" :key="item.id" :class="item.level > 1 ? 'bg-red-50/50 dark:bg-red-900/10' : ''">
                <td class="px-8 py-6 border-l-4" :class="item.level > 1 ? 'border-red-500' : 'border-transparent'">
                  <p class="font-black text-sm text-slate-800 dark:text-slate-200">{{ item.id.substring(0, 8) }}...</p>
                  <p class="text-[10px] text-slate-400 font-bold mt-1 uppercase">{{ item.entity_type }}</p>
                </td>
                <td class="px-8 py-6">
                  <div class="flex items-center gap-3">
                    <div class="w-8 h-8 rounded-full bg-slate-100 dark:bg-slate-800 flex items-center justify-center text-slate-400">
                       <LucideUser class="w-4 h-4" />
                    </div>
                    <span class="text-xs font-bold text-slate-700 dark:text-slate-300">{{ $t('common.system') || 'System' }}</span>
                  </div>
                </td>
                <td class="px-8 py-6">
                  <span :class="`text-xs font-black ${item.level > 1 ? 'text-red-500' : 'text-slate-500'}`">{{ timeAgo(item.created_at) }}</span>
                </td>
                <td class="px-8 py-6">
                  <span :class="`px-3 py-1 rounded-full text-[9px] font-black uppercase ${item.level > 1 ? 'bg-red-100 text-red-600' : 'bg-blue-100 text-blue-600'}`">
                    {{ item.status }}
                  </span>
                </td>
                <td class="px-8 py-6 text-right">
                  <div class="flex justify-end gap-2">
                    <button class="w-8 h-8 rounded-full bg-green-500/10 text-green-500 flex items-center justify-center hover:bg-green-500 hover:text-white transition-all shadow-sm">
                      <LucideCheck class="w-4 h-4" />
                    </button>
                    <button class="w-8 h-8 rounded-full bg-red-500/10 text-red-500 flex items-center justify-center hover:bg-red-500 hover:text-white transition-all shadow-sm">
                      <LucideX class="w-4 h-4" />
                    </button>
                  </div>
                </td>
              </tr>
              <tr v-if="!dashboardData?.tasks?.length">
                <td colspan="5" class="px-8 py-10 text-center text-slate-400 font-bold text-xs uppercase tracking-widest">
                  {{ $t('dashboard.manager.queue.no_pending') }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Right Column -->
      <div class="space-y-8">
        <!-- SLA Compliance -->
        <div class="glass rounded-3xl p-8 flex flex-col items-center justify-center text-center" v-motion-slide-visible-bottom>
          <h3 class="text-xs font-black uppercase tracking-widest text-slate-400 mb-8 self-start">{{ $t('dashboard.manager.compliance.title') }}</h3>
          <div class="relative w-40 h-40 flex items-center justify-center">
            <svg class="w-full h-full transform -rotate-90">
              <circle cx="80" cy="80" r="70" fill="transparent" stroke="currentColor" stroke-width="12" class="text-slate-100 dark:text-slate-800" />
              <circle cx="80" cy="80" r="70" fill="transparent" stroke="currentColor" stroke-width="12" stroke-dasharray="440" :stroke-dashoffset="440 - (440 * 0.92)" class="text-primary-600 dark:text-primary-400 transition-all duration-1000" />
            </svg>
            <div class="absolute inset-0 flex flex-col items-center justify-center">
              <span class="text-3xl font-black text-slate-800 dark:text-white">92%</span>
              <span class="text-[10px] font-bold text-slate-400 uppercase">{{ $t('dashboard.manager.compliance.target', { target: 95 }) }}</span>
            </div>
          </div>
        </div>

        <!-- Top Submitters -->
        <div class="glass rounded-3xl p-8" v-motion-slide-visible-bottom>
          <h3 class="text-xs font-black uppercase tracking-widest text-slate-400 mb-8">{{ $t('dashboard.manager.top_submitters.title') }}</h3>
          <div class="space-y-6">
            <div v-for="(person, index) in topSubmitters" :key="index" class="flex items-center gap-4">
              <span class="text-xs font-black text-slate-300 w-4">{{ person.rank }}</span>
              <div class="w-8 h-8 rounded-full bg-slate-100 dark:bg-slate-800 flex items-center justify-center text-slate-400">
                <LucideUser class="w-4 h-4" />
              </div>
              <div class="flex-1 min-w-0">
                <p class="text-sm font-bold text-slate-700 dark:text-slate-200 truncate">{{ person.name }}</p>
              </div>
              <span class="text-xs font-black text-slate-800 dark:text-white">{{ person.count }} <span class="text-[10px] text-slate-400 ml-0.5">{{ $t('dashboard.manager.top_submitters.docs') }}</span></span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Bottom Content -->
    <div class="grid grid-cols-1 xl:grid-cols-2 gap-8">
      <!-- Quick Bulk Actions -->
      <div class="glass rounded-2xl p-8" v-motion-slide-visible-bottom>
        <h3 class="text-xs font-black uppercase tracking-widest text-slate-400 mb-8">{{ $t('dashboard.manager.bulk_actions.title') }}</h3>
        <div class="grid grid-cols-2 gap-4">
          <button v-for="action in bulkActions" :key="action.label" class="flex items-center gap-4 p-5 bg-slate-50 dark:bg-slate-900/50 rounded-2xl hover:bg-primary-50 dark:hover:bg-primary-900/20 transition-all border border-transparent hover:border-primary-200 group">
            <div class="w-12 h-12 rounded-xl bg-white dark:bg-slate-800 flex items-center justify-center text-slate-400 group-hover:text-primary-500 transition-colors shadow-sm">
              <component :is="action.icon" class="w-6 h-6" />
            </div>
            <div class="text-left">
              <p class="text-xs font-black text-slate-700 dark:text-slate-200">{{ action.label }}</p>
              <p class="text-[9px] text-slate-400 font-bold mt-0.5 uppercase tracking-tighter">{{ action.desc }}</p>
            </div>
          </button>
        </div>
      </div>

      <!-- Recent Decisions Log -->
      <div class="glass rounded-2xl p-8" v-motion-slide-visible-bottom>
        <h3 class="text-xs font-black uppercase tracking-widest text-slate-400 mb-8">{{ $t('dashboard.manager.decisions.title') }}</h3>
        <div class="space-y-6">
          <div v-for="(log, index) in dashboardData?.activities?.slice(0, 5)" :key="index" class="flex gap-4">
            <div :class="`w-8 h-8 rounded-full flex-shrink-0 flex items-center justify-center ${getActivityStyles(log.action).bg}`">
              <component :is="getActivityStyles(log.action).icon" :class="`w-4 h-4 ${getActivityStyles(log.action).color}`" />
            </div>
            <div>
              <p class="text-sm text-slate-600 dark:text-slate-300">
                <span class="font-black text-slate-900 dark:text-white">{{ log.action }}</span> {{ log.entity_type }}
              </p>
              <div class="flex items-center gap-3 mt-1.5">
                <span class="text-[10px] font-bold text-slate-400">{{ timeAgo(log.created_at) }}</span>
                <span class="w-1 h-1 rounded-full bg-slate-200"></span>
                <span class="text-[10px] font-bold text-slate-400">{{ $t('dashboard.manager.decisions.submitter', { name: log.user_name }) }}</span>
              </div>
            </div>
          </div>
          <div v-if="!dashboardData?.activities?.length" class="text-center py-10 text-slate-400 font-bold text-xs uppercase tracking-widest">
            {{ $t('dashboard.manager.decisions.no_decisions') }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideAlertTriangle, 
  LucideTrendingUp, 
  LucideTrendingDown, 
  LucideCheck, 
  LucideX,
  LucideLayoutGrid,
  LucideBarChart3,
  LucideArchive,
  LucideLifeBuoy,
  LucideCheckCircle2,
  LucideXCircle,
  LucideMessageSquare,
  LucideUser,
  LucidePlus,
  LucideEdit,
  LucideTrash,
  LucideUpload,
  LucideActivity
} from 'lucide-vue-next'
import { useAuthStore } from '~/stores/auth'
import { useApi } from '~/composables/useApi'
import PageHeader from '~/components/PageHeader.vue'
import { formatBytes, timeAgo } from '~/utils/format'

const { t } = useI18n()
const auth = useAuthStore()
const { $api } = useApi()
const user = computed(() => auth.user)
const config = useRuntimeConfig()

// Data Fetching
const { data: dashboardData, refresh: refreshDashboard } = await useAsyncData('dashboard-summary', async () => {
  const res = await $api(`${config.public.apiBase}/dashboard/summary`)
  return res.data
})

const statsSummary = computed(() => dashboardData.value?.stats)

const stats = computed(() => [
  { label: t('dashboard.manager.stats.priority_pending'), value: statsSummary.value?.active_tasks || 0, trend: '+0%', trendUp: true, progress: Math.min((statsSummary.value?.active_tasks || 0) * 5, 100), barColor: 'bg-orange-500' },
  { label: t('dashboard.manager.stats.approved_today'), value: statsSummary.value?.daily_processed || 0, trend: '+0%', trendUp: true, progress: Math.min((statsSummary.value?.daily_processed || 0) * 10, 100), barColor: 'bg-green-500' },
  { label: t('dashboard.manager.stats.team_submissions'), value: statsSummary.value?.daily_received || 0, trend: '+0%', trendUp: true, progress: Math.min((statsSummary.value?.daily_received || 0) * 5, 100), barColor: 'bg-blue-600' },
  { label: t('dashboard.manager.stats.avg_approval_time'), value: '2.3h', trend: '-0h', trendUp: true, progress: 30, barColor: 'bg-primary-500' },
])

const topSubmitters = [
  { rank: '01', name: 'Indra M.', count: '0', avatar: '' },
  { rank: '02', name: 'Maya A.', count: '0', avatar: '' },
  { rank: '03', name: 'Dedi K.', count: '0', avatar: '' },
]

const bulkActions = [
  { label: t('dashboard.manager.bulk_actions.items.bulk_approve.label'), desc: t('dashboard.manager.bulk_actions.items.bulk_approve.desc'), icon: LucideLayoutGrid },
  { label: t('dashboard.manager.bulk_actions.items.deep_analytics.label'), desc: t('dashboard.manager.bulk_actions.items.deep_analytics.desc'), icon: LucideBarChart3 },
  { label: t('dashboard.manager.bulk_actions.items.archive_log.label'), desc: t('dashboard.manager.bulk_actions.items.archive_log.desc'), icon: LucideArchive },
  { label: t('dashboard.manager.bulk_actions.items.help_desk.label'), desc: t('dashboard.manager.bulk_actions.items.help_desk.desc'), icon: LucideLifeBuoy },
]

const getActivityStyles = (action) => {
  switch (action) {
    case 'CREATE': return { icon: LucidePlus, color: 'text-green-500', bg: 'bg-green-50 dark:bg-green-900/20' }
    case 'UPDATE': return { icon: LucideEdit, color: 'text-primary-500', bg: 'bg-primary-50 dark:bg-primary-900/20' }
    case 'DELETE': return { icon: LucideTrash, color: 'text-red-500', bg: 'bg-red-50 dark:bg-red-900/20' }
    case 'UPLOAD': return { icon: LucideUpload, color: 'text-teal-500', bg: 'bg-teal-50 dark:bg-teal-900/20' }
    default: return { icon: LucideActivity, color: 'text-slate-500', bg: 'bg-slate-50 dark:bg-slate-900/20' }
  }
}

// Polling for updates
let timer
onMounted(() => {
  timer = setInterval(() => {
    refreshDashboard()
  }, 30000)
})

onUnmounted(() => {
  if (timer) clearInterval(timer)
})
</script>

