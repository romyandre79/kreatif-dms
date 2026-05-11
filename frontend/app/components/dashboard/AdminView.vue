<template>
  <div class="space-y-8 pb-10">
    <!-- Greeting & Search Header -->
    <PageHeader 
      :title="$t('dashboard.greetings.morning', { name: user?.full_name?.split(' ')[0] || 'User' })"
      :subtitle="$t('dashboard.header.operational_overview', { date: currentDate })"
    >
      <template #actions>
        <div class="flex flex-wrap items-center gap-4">
          <button 
            @click="navigateTo('/documents/upload')"
            class="flex items-center gap-2 px-6 py-3 bg-[#1E3A5F] hover:bg-[#152943] text-white font-bold rounded-xl shadow-lg shadow-blue-900/20 transition-all"
          >
            <LucidePlus class="w-5 h-5" />
            {{ $t('dashboard.header.btn_submit_new') }}
          </button>
          <div class="relative flex-1 min-w-[200px]">
            <LucideSearch class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
            <input
              type="text"
              :placeholder="$t('dashboard.header.search_placeholder')"
              class="w-full pl-12 pr-4 py-3 text-sm rounded-2xl border border-slate-200 dark:border-slate-800 bg-white dark:bg-slate-900 shadow-sm focus:outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all font-medium"
            />
          </div>
        </div>
      </template>
    </PageHeader>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-5 gap-6">
      <!-- Daily Received -->
      <div v-motion-slide-visible-bottom class="glass p-6 rounded-2xl flex flex-col justify-between">
        <div class="flex items-center justify-between mb-6">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest leading-tight" v-html="$t('dashboard.stats.daily_received').replace('\n', '<br>')"></p>
          <LucideInbox class="w-5 h-5 text-green-500" />
        </div>
        <div>
          <p class="text-3xl font-black text-slate-900 dark:text-white">{{ stats?.daily_received || 0 }}</p>
          <div class="flex items-center gap-1 text-[10px] font-black text-green-500 mt-1 uppercase tracking-tighter">
            <LucideTrendingUp class="w-3 h-3" />
            {{ $t('dashboard.stats.vs_yesterday', { percent: 0 }) }}
          </div>
        </div>
      </div>

      <!-- Daily Scanned -->
      <div v-motion-slide-visible-bottom class="glass p-6 rounded-2xl flex flex-col justify-between">
        <div class="flex items-center justify-between mb-6">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest leading-tight" v-html="$t('dashboard.stats.daily_scanned').replace('\n', '<br>')"></p>
          <LucideScan class="w-5 h-5 text-blue-500" />
        </div>
        <div>
          <p class="text-3xl font-black text-slate-900 dark:text-white">{{ stats?.daily_scanned || 0 }}</p>
          <div class="flex items-center gap-1 text-[10px] font-black text-green-500 mt-1 uppercase tracking-tighter">
            <LucideTrendingUp class="w-3 h-3" />
            {{ $t('dashboard.stats.vs_yesterday', { percent: 0 }) }}
          </div>
        </div>
      </div>

      <!-- Daily Processed -->
      <div v-motion-slide-visible-bottom class="glass p-6 rounded-2xl flex flex-col justify-between">
        <div class="flex items-center justify-between mb-6">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest leading-tight" v-html="$t('dashboard.stats.daily_processed').replace('\n', '<br>')"></p>
          <LucideFileCheck class="w-5 h-5 text-primary-500" />
        </div>
        <div>
          <p class="text-3xl font-black text-slate-900 dark:text-white">{{ stats?.daily_processed || 0 }}</p>
          <div class="flex items-center gap-1 text-[10px] font-black text-green-500 mt-1 uppercase tracking-tighter">
            <LucideTrendingUp class="w-3 h-3" />
            {{ $t('dashboard.stats.vs_yesterday', { percent: 0 }) }}
          </div>
        </div>
      </div>

      <!-- Warehouse Capacity -->
      <div v-motion-slide-visible-bottom class="glass p-6 rounded-2xl flex flex-col justify-between">
        <div class="flex items-center justify-between mb-6">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest leading-tight" v-html="$t('dashboard.stats.warehouse_capacity').replace('\n', '<br>')"></p>
          <LucideLayers class="w-5 h-5 text-teal-500" />
        </div>
        <div>
          <p class="text-xl font-black text-slate-900 dark:text-white">{{ formatBytes(stats?.total_storage_size || 0) }} / 10TB</p>
          <div class="h-1.5 w-full bg-slate-100 dark:bg-slate-800 rounded-full mt-3 overflow-hidden">
            <div class="h-full bg-teal-500 rounded-full" :style="`width: ${Math.min(((stats?.total_storage_size || 0) / (10 * 1024 * 1024 * 1024 * 1024)) * 100, 100)}%`"></div>
          </div>
        </div>
      </div>

      <!-- Active Tasks -->
      <div v-motion-slide-visible-bottom class="bg-[#2D9B7B] p-6 rounded-2xl flex flex-col justify-between text-white shadow-xl shadow-teal-900/20">
        <div class="flex items-center justify-between mb-6">
          <p class="text-[10px] font-black uppercase tracking-widest leading-tight text-teal-50" v-html="$t('dashboard.stats.active_tasks').replace('\n', '<br>')"></p>
          <LucideClipboardList class="w-5 h-5 text-teal-50" />
        </div>
        <div>
          <p class="text-4xl font-black">{{ stats?.active_tasks || 0 }}</p>
          <p class="text-[10px] font-bold text-teal-50 mt-1">{{ $t('dashboard.stats.immediate_attention') }}</p>
        </div>
      </div>

      <!-- SLA Compliance -->
      <div v-motion-slide-visible-bottom class="glass p-6 rounded-2xl flex flex-col justify-between border-l-4 border-primary-500">
        <div class="flex items-center justify-between mb-6">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest leading-tight">SLA<br>Compliance</p>
          <LucideActivity class="w-5 h-5 text-primary-500" />
        </div>
        <div>
          <p class="text-3xl font-black text-slate-900 dark:text-white">{{ stats?.compliance_rate?.toFixed(1) || 100 }}%</p>
          <div class="flex items-center gap-1 text-[10px] font-black text-slate-400 mt-1 uppercase tracking-tighter">
            Avg: {{ formatTime(stats?.avg_approval_time) }}
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="grid grid-cols-1 xl:grid-cols-3 gap-8">
      <!-- Priority Task Queue -->
      <div class="xl:col-span-2 glass rounded-2xl overflow-hidden" v-motion-slide-visible-bottom>
        <div class="px-8 py-6 flex items-center justify-between">
          <h3 class="font-black text-lg text-[#1E3A5F] dark:text-white">{{ $t('dashboard.queue.priority_tasks') }}</h3>
          <button class="text-primary-600 font-black text-xs hover:underline uppercase tracking-widest">{{ $t('dashboard.queue.view_all_tasks') }}</button>
        </div>
        <div class="overflow-x-auto">
          <table class="w-full text-left">
            <thead>
              <tr class="text-slate-400 text-[10px] uppercase font-black tracking-widest border-b border-slate-100 dark:border-slate-800">
                <th class="px-8 py-4">{{ $t('dashboard.queue.table.manifest') }}</th>
                <th class="px-8 py-4">{{ $t('dashboard.queue.table.type') }}</th>
                <th class="px-8 py-4">{{ $t('dashboard.queue.table.status') }}</th>
                <th class="px-8 py-4 text-right">{{ $t('dashboard.queue.table.urgency') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50 dark:divide-slate-900">
              <tr v-for="task in dashboardData?.tasks" :key="task.id" class="hover:bg-slate-50/50 dark:hover:bg-slate-800/20 transition-colors">
                <td class="px-8 py-6">
                  <p class="font-black text-sm text-[#1E3A5F] dark:text-slate-200">{{ task.id.substring(0, 8) }}...</p>
                  <p class="text-[10px] text-slate-400 font-bold mt-0.5">{{ task.entity_type }}</p>
                </td>
                <td class="px-8 py-6">
                  <span class="text-sm font-bold text-slate-600 dark:text-slate-400">{{ task.entity_type }}</span>
                </td>
                <td class="px-8 py-6">
                  <span class="text-xs font-black text-primary-600 uppercase tracking-tighter">{{ task.status }}</span>
                </td>
                <td class="px-8 py-6 text-right">
                  <span :class="`px-3 py-1 rounded-lg text-[9px] font-black uppercase ${task.level > 1 ? 'bg-red-50 text-red-500' : 'bg-orange-50 text-orange-500'}`">
                    {{ $t('dashboard.queue.level') }} {{ task.level }}
                  </span>
                </td>
              </tr>
              <tr v-if="!dashboardData?.tasks?.length">
                <td colspan="4" class="px-8 py-10 text-center text-slate-400 font-bold text-xs uppercase tracking-widest">
                  {{ $t('dashboard.queue.no_tasks') }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Recent Activity -->
      <div class="glass rounded-2xl p-8 flex flex-col" v-motion-slide-visible-bottom>
        <h3 class="font-black text-lg text-[#1E3A5F] dark:text-white mb-10 text-center">{{ $t('dashboard.activity.title_recent') }}</h3>
        <div class="space-y-10 flex-1">
          <div v-for="(act, index) in dashboardData?.activities" :key="index" class="flex gap-6 relative">
            <div v-if="index !== dashboardData.activities.length - 1" class="absolute left-3.5 top-8 w-0.5 h-10 bg-slate-100 dark:bg-slate-800"></div>
            <div :class="`w-7 h-7 rounded-full flex-shrink-0 flex items-center justify-center z-10 border-4 border-white dark:border-slate-900 ${getActivityStyles(act.action).bg}`">
              <component :is="getActivityStyles(act.action).icon" :class="`w-3 h-3 ${getActivityStyles(act.action).color}`" />
            </div>
            <div>
              <p class="text-sm font-bold text-[#1E3A5F] dark:text-slate-200">{{ act.action }} {{ act.entity_type }}</p>
              <p class="text-xs text-slate-400 mt-1 leading-relaxed">{{ $t('dashboard.activity.by', { name: act.user_name }) }}</p>
              <p class="text-[10px] text-slate-300 dark:text-slate-500 font-black uppercase mt-1.5">{{ timeAgo(act.created_at) }}</p>
            </div>
          </div>
          <div v-if="!dashboardData?.activities?.length" class="text-center py-10 text-slate-400 font-bold text-xs uppercase tracking-widest">
             {{ $t('dashboard.activity.no_activity') }}
          </div>
        </div>
        <button class="w-full mt-10 py-3.5 bg-slate-50 dark:bg-slate-800/50 rounded-xl text-[10px] font-black uppercase tracking-widest text-slate-400 hover:bg-slate-100 transition-all">
          {{ $t('dashboard.activity.btn_logs') }}
        </button>
      </div>
    </div>

    <!-- System Update Banner -->
    <div v-if="dashboardData?.announcement?.active" class="bg-gradient-to-br from-[#1E3A5F] to-[#152943] rounded-lg p-6 text-white relative overflow-hidden group shadow-2xl shadow-blue-900/40 flex items-center justify-between" v-motion-slide-visible-bottom>
      <div class="relative z-10 flex items-center gap-8">
        <div class="w-16 h-16 bg-white/10 rounded-2xl flex items-center justify-center shrink-0 ring-1 ring-white/20">
           <LucideInfo class="w-8 h-8 text-white" />
        </div>
        <div>
          <h3 class="font-black text-xl mb-1 leading-tight">{{ dashboardData.announcement.title }}</h3>
          <p class="text-blue-100/70 text-sm font-medium leading-relaxed max-w-2xl">
            {{ dashboardData.announcement.message }}
          </p>
        </div>
      </div>
      <div class="relative z-10">
        <button 
          @click="isNotesOpen = true"
          class="px-8 py-3 bg-white text-[#1E3A5F] rounded-xl font-black text-xs uppercase tracking-widest hover:bg-blue-50 transition-all shadow-xl shadow-black/20"
        >
          {{ $t('dashboard.user.system_update.btn_notes') }}
        </button>
      </div>
      <!-- Decorative bubbles -->
      <div class="absolute -right-10 -bottom-10 w-64 h-64 bg-white/5 rounded-full blur-3xl group-hover:scale-125 transition-transform duration-1000"></div>
      <div class="absolute -left-10 -top-10 w-32 h-32 bg-primary-500/10 rounded-full blur-2xl"></div>
    </div>

    <!-- Infrastructure Status -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6" v-motion-slide-visible-bottom>
      <!-- Scanner -->
      <div class="glass p-6 rounded-2xl flex items-center gap-5 group hover:border-green-500/30 transition-all">
        <div :class="`w-14 h-14 rounded-2xl flex items-center justify-center transition-all ${isScannerActive ? 'bg-green-50 dark:bg-green-900/20 text-green-500' : 'bg-slate-50 dark:bg-slate-900/20 text-slate-400'}`">
          <LucidePrinter class="w-7 h-7" />
        </div>
        <div>
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('dashboard.infrastructure.scanner') }}</p>
          <p class="text-sm font-black text-slate-800 dark:text-white">{{ scannerNode?.name || $t('dashboard.infrastructure.no_scanner') }}</p>
          <p :class="`text-[10px] font-bold uppercase ${isScannerActive ? 'text-green-500' : 'text-slate-400'}`">
            {{ isScannerActive ? $t('dashboard.infrastructure.connected') : $t('dashboard.infrastructure.offline') }}
          </p>
        </div>
      </div>

      <!-- OCR Engine -->
      <div class="glass p-6 rounded-2xl flex items-center gap-5 group hover:border-primary-500/30 transition-all">
        <div :class="`w-14 h-14 rounded-2xl flex items-center justify-center transition-all ${isOCRActive ? 'bg-primary-50 dark:bg-primary-900/20 text-primary-500' : 'bg-slate-50 dark:bg-slate-900/20 text-slate-400'}`">
          <LucideCpu class="w-7 h-7" />
        </div>
        <div>
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('dashboard.infrastructure.ocr_engine') }}</p>
          <p class="text-sm font-black text-slate-800 dark:text-white">{{ ocrNode?.name || 'PaddleOCR' }}</p>
          <p :class="`text-[10px] font-bold uppercase ${isOCRActive ? 'text-green-500' : 'text-slate-400'}`">
             {{ isOCRActive ? $t('dashboard.infrastructure.active') : $t('dashboard.infrastructure.offline') }}
          </p>
        </div>
      </div>

      <!-- Database -->
      <div class="glass p-6 rounded-2xl flex items-center gap-5 group hover:border-teal-500/30 transition-all">
        <div class="w-14 h-14 rounded-2xl bg-teal-50 dark:bg-teal-900/20 flex items-center justify-center text-teal-500">
          <LucideDatabase class="w-7 h-7" />
        </div>
        <div>
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('dashboard.infrastructure.database') }}</p>
          <p class="text-sm font-black text-slate-800 dark:text-white">PostgreSQL Cluster</p>
          <p class="text-[10px] font-bold text-green-500 uppercase">{{ $t('dashboard.infrastructure.synchronized') }}</p>
        </div>
      </div>
    </div>

    <!-- Release Notes Modal -->
    <Modal 
      v-model="isNotesOpen" 
      :title="dashboardData?.announcement?.title || 'Release Notes'"
    >
      <div class="prose prose-slate dark:prose-invert max-w-none">
        <div v-html="parseMarkdown(dashboardData?.announcement?.notes || '')"></div>
      </div>
      <template #footer>
        <button 
          @click="isNotesOpen = false"
          class="px-6 py-2 bg-[#1E3A5F] text-white rounded-xl font-black text-xs uppercase tracking-widest hover:bg-[#152943] transition-all"
        >
          {{ $t('common.close') || 'Tutup' }}
        </button>
      </template>
    </Modal>
  </div>
</template>

<script setup>
import { 
  LucideSearch, 
  LucideBell, 
  LucideInbox, 
  LucideTrendingUp, 
  LucideScan, 
  LucideFileCheck,
  LucideLayers,
  LucideClipboardList,
  LucideAlertCircle,
  LucideRefreshCcw,
  LucideUser,
  LucidePrinter,
  LucideCpu,
  LucideDatabase,
  LucideFileStack,
  LucidePlus,
  LucideActivity,
  LucideTrash,
  LucideEdit,
  LucideUpload,
  LucideInfo
} from 'lucide-vue-next'
import { useAuthStore } from '~/stores/auth'
import { useApi } from '~/composables/useApi'
import PageHeader from '~/components/PageHeader.vue'
import Modal from '~/components/Modal.vue'
import { formatBytes, timeAgo, parseMarkdown } from '~/utils/format'

const auth = useAuthStore()
const { $api } = useApi()
const { t } = useI18n()
const user = computed(() => auth.user)
const config = useRuntimeConfig()
const isNotesOpen = ref(false)

const currentDate = computed(() => {
  const d = new Date();
  const day = String(d.getDate()).padStart(2, '0');
  const month = String(d.getMonth() + 1).padStart(2, '0');
  const year = d.getFullYear();
  return `${day}.${month}.${year}`;
})

// Data Fetching
const { data: dashboardData, refresh: refreshDashboard } = await useAsyncData(`admin-summary-${user.value?.id}`, async () => {
  const res = await $api('/dashboard/summary')
  return res.data
})

const { data: integrationData, refresh: refreshIntegration } = await useAsyncData('health-status', async () => {
  const res = await $api('/master/integration/status')
  return res.data
})

const stats = computed(() => dashboardData.value?.stats)

// Integration Node Helpers
const scannerNode = computed(() => integrationData.value?.nodes?.find(n => n.service_type === 'SCANNER_LOCAL' || n.service_type === 'SCANNER_NETWORK'))
const isScannerActive = computed(() => scannerNode.value?.is_active)

const ocrNode = computed(() => integrationData.value?.nodes?.find(n => n.service_type === 'OCR'))
const isOCRActive = computed(() => ocrNode.value?.is_active)

const getActivityStyles = (action) => {
  switch (action) {
    case 'CREATE': return { icon: LucidePlus, color: 'text-green-500', bg: 'bg-green-50 dark:bg-green-900/20' }
    case 'UPDATE': return { icon: LucideEdit, color: 'text-primary-500', bg: 'bg-primary-50 dark:bg-primary-900/20' }
    case 'DELETE': return { icon: LucideTrash, color: 'text-red-500', bg: 'bg-red-50 dark:bg-red-900/20' }
    case 'UPLOAD': return { icon: LucideUpload, color: 'text-teal-500', bg: 'bg-teal-50 dark:bg-teal-900/20' }
    default: return { icon: LucideActivity, color: 'text-slate-500', bg: 'bg-slate-50 dark:bg-slate-900/20' }
  }
}

const formatTime = (seconds) => {
  if (!seconds || seconds <= 0) return '0h'
  if (seconds < 3600) return `${(seconds / 60).toFixed(1)}m`
  return `${(seconds / 3600).toFixed(1)}h`
}

// Polling for updates
let timer
onMounted(() => {
  refreshDashboard()
  refreshIntegration()
  
  timer = setInterval(() => {
    refreshDashboard()
    refreshIntegration()
  }, 10000)
})

onUnmounted(() => {
  if (timer) clearInterval(timer)
})
</script>

