<template>
  <div class="space-y-8 pb-10">
    <!-- Urgent Alert Banner -->
    <div v-motion-fade class="bg-orange-50 dark:bg-orange-900/20 border-l-4 border-orange-500 p-6 rounded-r-2xl flex items-center justify-between shadow-sm">
      <div class="flex items-center gap-4">
        <div class="w-10 h-10 rounded-full bg-orange-100 dark:bg-orange-800 flex items-center justify-center text-orange-600 dark:text-orange-300">
          <LucideAlertTriangle class="w-6 h-6" />
        </div>
        <div>
          <h3 class="font-black text-orange-900 dark:text-orange-200">{{ $t('dashboard.manager.alert_title') }}</h3>
          <p class="text-sm text-orange-700/80 dark:text-orange-400/80">{{ $t('dashboard.manager.alert_desc', { count: 12 }) }}</p>
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
              <tr v-for="item in queue" :key="item.id" :class="item.bgClass">
                <td class="px-8 py-6 border-l-4" :class="item.borderColor">
                  <p class="font-black text-sm text-slate-800 dark:text-slate-200">{{ item.id }}</p>
                  <p class="text-[10px] text-slate-400 font-bold mt-1 uppercase">{{ item.desc }}</p>
                </td>
                <td class="px-8 py-6">
                  <div class="flex items-center gap-3">
                    <img :src="item.avatar" class="w-8 h-8 rounded-full border-2 border-white dark:border-slate-800 shadow-sm" />
                    <span class="text-xs font-bold text-slate-700 dark:text-slate-300">{{ item.staff }}</span>
                  </div>
                </td>
                <td class="px-8 py-6">
                  <span :class="`text-xs font-black ${item.statusColor}`">{{ item.elapsed }}</span>
                </td>
                <td class="px-8 py-6">
                  <span :class="`px-3 py-1 rounded-full text-[9px] font-black uppercase ${item.statusBg}`">
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
              <circle cx="80" cy="80" r="70" fill="transparent" stroke="currentColor" stroke-width="12" stroke-dasharray="440" stroke-dashoffset="35" class="text-primary-600 dark:text-primary-400 transition-all duration-1000" />
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
              <img :src="person.avatar" class="w-8 h-8 rounded-full" />
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
          <div v-for="(log, index) in decisions" :key="index" class="flex gap-4">
            <div :class="`w-8 h-8 rounded-full flex-shrink-0 flex items-center justify-center ${log.bg}`">
              <component :is="log.icon" :class="`w-4 h-4 ${log.color}`" />
            </div>
            <div>
              <p class="text-sm text-slate-600 dark:text-slate-300">
                <span class="font-black text-slate-900 dark:text-white">{{ log.action }}</span> {{ log.target }}
              </p>
              <div class="flex items-center gap-3 mt-1.5">
                <span class="text-[10px] font-bold text-slate-400">{{ log.time }}</span>
                <span class="w-1 h-1 rounded-full bg-slate-200"></span>
                <span class="text-[10px] font-bold text-slate-400">{{ $t('dashboard.manager.decisions.submitter', { name: log.submitter }) }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const { t } = useI18n()
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
  LucideMessageSquare
} from 'lucide-vue-next'

const stats = [
  { label: t('dashboard.manager.stats.priority_pending'), value: '12', trend: '+15%', trendUp: true, progress: 65, barColor: 'bg-orange-500' },
  { label: t('dashboard.manager.stats.approved_today'), value: '07', trend: '+2%', trendUp: true, progress: 40, barColor: 'bg-green-500' },
  { label: t('dashboard.manager.stats.team_submissions'), value: '45', trend: '-5%', trendUp: false, progress: 80, barColor: 'bg-blue-600' },
  { label: t('dashboard.manager.stats.avg_approval_time'), value: '2.3h', trend: '-0.4h', trendUp: true, progress: 30, barColor: 'bg-primary-500' },
]

const queue = [
  { 
    id: 'INV-2023-8821', desc: t('dashboard.manager.queue.items.inv_8821'), staff: t('dashboard.manager.top_submitters.items.indra'), avatar: 'https://i.pravatar.cc/150?u=budi',
    elapsed: t('dashboard.user.activity.items.three_hours'), status: 'OVERDUE', statusBg: 'bg-red-100 text-red-600', statusColor: 'text-red-500',
    bgClass: 'bg-red-50/50 dark:bg-red-900/10', borderColor: 'border-red-500'
  },
  { 
    id: 'DOC-XP-012', desc: t('dashboard.manager.queue.items.doc_xp_012'), staff: t('dashboard.manager.top_submitters.items.maya'), avatar: 'https://i.pravatar.cc/150?u=ani',
    elapsed: t('dashboard.user.activity.items.three_hours'), status: 'CRITICAL', statusBg: 'bg-red-100 text-red-600', statusColor: 'text-red-500',
    bgClass: 'bg-red-50/50 dark:bg-red-900/10', borderColor: 'border-red-400'
  },
  { 
    id: 'PR-990-2', desc: t('dashboard.manager.queue.items.pr_990_2'), staff: t('dashboard.manager.top_submitters.items.dedi'), avatar: 'https://i.pravatar.cc/150?u=joko',
    elapsed: t('dashboard.user.activity.items.one_hour'), status: 'PENDING', statusBg: 'bg-blue-100 text-blue-600', statusColor: 'text-slate-500',
    bgClass: '', borderColor: 'border-transparent'
  },
]

const topSubmitters = [
  { rank: '01', name: t('dashboard.manager.top_submitters.items.indra'), count: '12', avatar: 'https://i.pravatar.cc/150?u=1' },
  { rank: '02', name: t('dashboard.manager.top_submitters.items.maya'), count: '10', avatar: 'https://i.pravatar.cc/150?u=2' },
  { rank: '03', name: t('dashboard.manager.top_submitters.items.dedi'), count: '8', avatar: 'https://i.pravatar.cc/150?u=3' },
]

const bulkActions = [
  { label: t('dashboard.manager.bulk_actions.items.bulk_approve.label'), desc: t('dashboard.manager.bulk_actions.items.bulk_approve.desc'), icon: LucideLayoutGrid },
  { label: t('dashboard.manager.bulk_actions.items.deep_analytics.label'), desc: t('dashboard.manager.bulk_actions.items.deep_analytics.desc'), icon: LucideBarChart3 },
  { label: t('dashboard.manager.bulk_actions.items.archive_log.label'), desc: t('dashboard.manager.bulk_actions.items.archive_log.desc'), icon: LucideArchive },
  { label: t('dashboard.manager.bulk_actions.items.help_desk.label'), desc: t('dashboard.manager.bulk_actions.items.help_desk.desc'), icon: LucideLifeBuoy },
]

const decisions = [
  { action: t('dashboard.manager.decisions.action_approved'), target: t('dashboard.manager.decisions.items.tr_2201'), time: t('dashboard.user.activity.items.ten_mins'), submitter: 'Kevin S.', icon: LucideCheckCircle2, color: 'text-green-500', bg: 'bg-green-50 dark:bg-green-900/20' },
  { action: t('dashboard.manager.decisions.action_rejected'), target: t('dashboard.manager.decisions.items.rq_882'), time: t('dashboard.user.activity.items.one_hour'), submitter: 'Linda R.', icon: LucideXCircle, color: 'text-red-500', bg: 'bg-red-50 dark:bg-red-900/20' },
  { action: t('dashboard.manager.decisions.action_commented'), target: t('dashboard.manager.decisions.items.po_991'), time: t('dashboard.user.activity.items.three_hours'), submitter: 'Finance', icon: LucideMessageSquare, color: 'text-blue-500', bg: 'bg-blue-50 dark:bg-blue-900/20' },
]
</script>
