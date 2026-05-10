<template>
  <div class="space-y-8 pb-10">
    <!-- Action Header -->
    <PageHeader 
      :title="$t('dashboard.greetings.morning', { name: user?.full_name?.split(' ')[0] || 'User' })"
      :subtitle="$t('dashboard.user.overview_desc')"
    >
      <template #actions>
        <button 
          @click="navigateTo('/documents/upload')"
          class="flex items-center gap-2 px-6 py-3 bg-[#1E3A5F] hover:bg-[#152943] text-white font-bold rounded-xl shadow-lg shadow-blue-900/20 transition-all"
        >
          <LucidePlus class="w-5 h-5" />
          {{ $t('dashboard.header.btn_submit_new') }}
        </button>
      </template>
    </PageHeader>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- My Submissions -->
      <div v-motion-slide-visible-bottom class="glass p-4 rounded-lg relative overflow-hidden group shadow-sm hover:shadow-xl transition-all duration-500">
        <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-6">{{ $t('dashboard.user.stats.my_submissions') }}</p>
        <div class="flex items-end justify-between">
          <p class="text-5xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ statsSummary?.daily_received || 0 }}</p>
          <div class="flex items-center gap-1.5 text-[#2D9B7B] text-[10px] font-black bg-[#2D9B7B]/10 px-3 py-1.5 rounded-xl uppercase tracking-tighter mb-1">
            <LucideTrendingUp class="w-3.5 h-3.5" />
            +0%
          </div>
        </div>
      </div>

      <!-- Pending Approvals -->
      <div v-motion-slide-visible-bottom class="glass p-4 rounded-lg relative overflow-hidden group shadow-sm hover:shadow-xl transition-all duration-500">
        <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-6">{{ $t('dashboard.user.stats.pending_approvals') }}</p>
        <div class="flex items-center justify-between">
          <p class="text-5xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ statsSummary?.active_tasks || 0 }}</p>
          <div class="bg-orange-50 dark:bg-orange-900/30 text-orange-600 dark:text-orange-400 px-4 py-2 rounded-xl text-[9px] font-black uppercase leading-tight tracking-widest text-center">
            ACTION<br>REQUIRED
          </div>
        </div>
      </div>

      <!-- Active Loans -->
      <div v-motion-slide-visible-bottom class="glass p-4 rounded-lg relative overflow-hidden group shadow-sm hover:shadow-xl transition-all duration-500">
        <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-6">{{ $t('dashboard.user.stats.active_loans') }}</p>
        <div class="flex items-center justify-between">
          <p class="text-5xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ statsSummary?.active_loans || 0 }}</p>
          <div class="bg-green-50 dark:bg-green-900/30 text-green-600 dark:text-green-400 px-4 py-2 rounded-xl text-[10px] font-black uppercase tracking-widest">
            ACTIVE
          </div>
        </div>
      </div>

      <!-- Search History -->
      <div v-motion-slide-visible-bottom class="glass p-4 rounded-lg relative overflow-hidden group shadow-sm hover:shadow-xl transition-all duration-500">
        <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-6">{{ $t('dashboard.user.stats.search_history') }}</p>
        <div class="flex items-end justify-between">
          <p class="text-5xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ statsSummary?.search_count || 0 }}</p>
          <p class="text-slate-400 text-[11px] font-bold mb-2 uppercase tracking-widest">Today</p>
        </div>
      </div>
    </div>

    <!-- Middle Grid -->
    <div class="grid grid-cols-1 xl:grid-cols-3 gap-8">
      <!-- Items Requiring My Action -->
      <div class="xl:col-span-2 glass rounded-lg overflow-hidden shadow-sm" v-motion-slide-visible-bottom>
        <div class="px-8 py-7 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between">
          <h3 class="font-black text-lg text-[#1E3A5F] dark:text-white">{{ $t('dashboard.user.tasks.title') }}</h3>
          <button class="text-primary-600 dark:text-primary-400 text-xs font-black uppercase tracking-widest hover:underline">{{ $t('dashboard.queue.view_all_tasks') }}</button>
        </div>
        <div class="overflow-x-auto">
          <table class="w-full text-left">
            <thead>
              <tr class="text-slate-400 text-[10px] uppercase font-black tracking-widest bg-slate-50/30 dark:bg-slate-900/30">
                <th class="px-8 py-5">{{ $t('dashboard.user.tasks.table.task_type') }}</th>
                <th class="px-8 py-5">{{ $t('dashboard.user.tasks.table.doc_manifest') }}</th>
                <th class="px-8 py-5">{{ $t('dashboard.user.tasks.table.due_date') }}</th>
                <th class="px-8 py-5 text-right">{{ $t('dashboard.user.tasks.table.action') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-100/50 dark:divide-slate-800/50">
              <tr v-for="task in dashboardData?.tasks" :key="task.id" class="group hover:bg-slate-50/50 dark:hover:bg-slate-900/50 transition-colors">
                <td class="px-8 py-6">
                  <div class="flex items-center gap-4">
                    <div :class="`w-2.5 h-2.5 rounded-full ring-4 ring-offset-2 ring-transparent group-hover:ring-offset-0 transition-all ${task.level > 1 ? 'bg-orange-500 shadow-[0_0_10px_rgba(249,115,22,0.4)]' : 'bg-blue-500 shadow-[0_0_10px_rgba(59,130,246,0.4)]'}`"></div>
                    <span class="font-bold text-sm text-slate-700 dark:text-slate-200">{{ formatTaskType(task.entity_type) }}</span>
                  </div>
                </td>
                <td class="px-8 py-6">
                  <span class="text-xs font-bold text-slate-400 tracking-tighter">{{ task.id.substring(0, 8) }}...</span>
                </td>
                <td class="px-8 py-6">
                  <div class="text-xs font-bold text-slate-500">
                    <p>{{ timeAgo(task.created_at) }}</p>
                  </div>
                </td>
                <td class="px-8 py-6 text-right">
                  <button 
                    @click="handleTaskAction(task)"
                    class="px-6 py-2.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-[10px] font-black uppercase tracking-widest text-[#1E3A5F] dark:text-white hover:bg-slate-50 transition-all flex items-center gap-2 group/btn shadow-sm"
                  >
                    {{ 
                      task.entity_type === 'document_rejection' ? 'Edit' : 
                      (task.level > 1 ? $t('dashboard.user.tasks.items.execute') : 
                      (task.entity_type.includes('Scan') ? $t('dashboard.user.tasks.items.process') : $t('dashboard.user.tasks.items.verify'))) 
                    }}
                    <LucideArrowRight class="w-3.5 h-3.5 group-hover/btn:translate-x-1 transition-transform" />
                  </button>
                </td>
              </tr>
              <tr v-if="!dashboardData?.tasks?.length">
                <td colspan="4" class="px-8 py-10 text-center text-slate-400 font-bold text-xs uppercase tracking-widest">
                  {{ $t('dashboard.user.tasks.no_tasks') }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Quick Actions -->
      <div class="glass rounded-lg p-4 flex flex-col" v-motion-slide-visible-bottom>
        <h3 class="font-black text-lg text-[#1E3A5F] dark:text-white mb-10">{{ $t('dashboard.user.quick_actions.title') }}</h3>
        <div class="grid grid-cols-2 gap-6 flex-1">
          <button 
            v-for="action in quickActions" 
            :key="action.label" 
            @click="navigateTo(action.path)"
            class="flex flex-col items-center justify-center p-4 bg-slate-50 dark:bg-slate-900/50 rounded-3xl hover:bg-white dark:hover:bg-slate-800 transition-all border-2 border-transparent hover:border-primary-100 dark:hover:border-primary-900/30 group shadow-sm hover:shadow-xl"
          >
            <div class="w-16 h-16 rounded-2xl bg-white dark:bg-slate-800 flex items-center justify-center mb-5 text-primary-500 shadow-lg shadow-primary-500/5 group-hover:scale-110 transition-transform">
              <component :is="action.icon" class="w-8 h-8" />
            </div>
            <span class="text-[11px] font-black uppercase tracking-[0.1em] text-[#1E3A5F] dark:text-slate-300 text-center">{{ action.label }}</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Bottom Grid -->
    <div class="grid grid-cols-1 xl:grid-cols-3 gap-8">
      <!-- Recent Activities -->
      <div class="xl:col-span-2 glass rounded-lg p-4 shadow-sm" v-motion-slide-visible-bottom>
        <h3 class="font-black text-lg text-[#1E3A5F] dark:text-white mb-10">{{ $t('dashboard.user.activity.title') }}</h3>
        <div class="space-y-10">
          <div v-for="(activity, index) in dashboardData?.activities" :key="index" class="flex gap-8 relative">
            <div v-if="index !== (dashboardData?.activities?.length || 0) - 1" class="absolute left-[11px] top-4 w-0.5 h-10 bg-slate-100 dark:bg-slate-800"></div>
            <div :class="`w-6 h-6 rounded-full flex-shrink-0 flex items-center justify-center z-10 ring-4 ring-white dark:ring-[#0A0F1C] ${getActivityStyles(activity.action).dot}`">
              <div class="w-2.5 h-2.5 rounded-full bg-white dark:bg-[#0A0F1C]"></div>
            </div>
            <div class="flex-1 flex justify-between items-start gap-6">
              <div class="max-w-[80%]">
                <p class="text-sm font-bold text-slate-700 dark:text-slate-300 leading-relaxed">
                   <span class="font-black">{{ activity.action }}</span> {{ activity.entity_type }} ({{ activity.entity_id?.substring(0, 8) }}...)
                </p>
                <p class="text-xs text-slate-400 font-black uppercase mt-2 tracking-tighter">{{ timeAgo(activity.created_at) }}</p>
              </div>
              <button 
                @click="navigateTo(`/documents/${activity.entity_id}`)"
                class="text-slate-300 hover:text-primary-500 text-[10px] font-black uppercase tracking-widest transition-colors"
              >
                {{ $t('dashboard.user.activity.items.view') }}
              </button>
            </div>
          </div>
          <div v-if="!dashboardData?.activities?.length" class="text-center py-10 text-slate-400 font-bold text-xs uppercase tracking-widest">
            {{ $t('dashboard.user.activity.no_activity') }}
          </div>
        </div>
      </div>

      <!-- System Update Banner -->
      <div v-if="dashboardData?.announcement?.active" class="bg-gradient-to-br from-[#1E3A5F] to-[#152943] rounded-lg p-4 text-white relative overflow-hidden group shadow-2xl shadow-blue-900/40 flex flex-col justify-between" v-motion-slide-visible-bottom>
        <div class="relative z-10">
          <div class="w-14 h-14 bg-white/10 rounded-2xl flex items-center justify-center mb-8 ring-1 ring-white/20">
             <LucideInfo class="w-7 h-7 text-white" />
          </div>
          <h3 class="font-black text-2xl mb-4 leading-tight">{{ dashboardData.announcement.title }}</h3>
          <p class="text-blue-100/70 text-sm font-medium leading-relaxed mb-10">
            {{ dashboardData.announcement.message }}
          </p>
        </div>
        <div class="relative z-10">
          <button 
            @click="isNotesOpen = true"
            class="w-full py-4 bg-white text-[#1E3A5F] rounded-2xl font-black text-xs uppercase tracking-widest hover:bg-blue-50 transition-all shadow-xl shadow-black/20"
          >
            {{ $t('dashboard.user.system_update.btn_notes') }}
          </button>
        </div>
        <!-- Decorative bubbles -->
        <div class="absolute -right-10 -bottom-10 w-64 h-64 bg-white/5 rounded-full blur-3xl group-hover:scale-125 transition-transform duration-1000"></div>
        <div class="absolute -left-10 -top-4 w-32 h-32 bg-primary-500/10 rounded-full blur-2xl"></div>
      </div>
    </div>

    <!-- Loan History Table -->
    <div class="glass rounded-lg overflow-hidden shadow-sm" v-motion-slide-visible-bottom>
      <div class="px-10 py-8 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between bg-slate-50/20 dark:bg-slate-900/20">
        <h3 class="font-black text-lg text-[#1E3A5F] dark:text-white">{{ $t('dashboard.user.loan_history.title') }}</h3>
        <div class="flex gap-3">
          <button class="px-5 py-2.5 bg-white dark:bg-slate-900 rounded-xl text-[10px] font-black text-slate-500 border border-slate-200 dark:border-slate-800 hover:bg-slate-50 hover:text-primary-500 uppercase tracking-widest transition-all">{{ $t('dashboard.user.loan_history.btn_export') }}</button>
          <button class="px-5 py-2.5 bg-white dark:bg-slate-900 rounded-xl text-[10px] font-black text-slate-500 border border-slate-200 dark:border-slate-800 hover:bg-slate-50 hover:text-primary-500 uppercase tracking-widest transition-all">{{ $t('dashboard.user.loan_history.btn_filter') }}</button>
        </div>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full text-left">
          <thead>
            <tr class="text-slate-400 text-[10px] uppercase font-black tracking-widest border-b border-slate-50 dark:border-slate-900">
              <th class="px-10 py-5">{{ $t('dashboard.user.loan_history.table.loan_id') }}</th>
              <th class="px-10 py-5">{{ $t('dashboard.user.loan_history.table.doc_title') }}</th>
              <th class="px-10 py-5">{{ $t('dashboard.user.loan_history.table.loan_date') }}</th>
              <th class="px-10 py-5">{{ $t('dashboard.user.loan_history.table.expected_return') }}</th>
              <th class="px-10 py-5">{{ $t('dashboard.user.loan_history.table.status') }}</th>
              <th class="px-10 py-5 text-right">{{ $t('dashboard.user.loan_history.table.action') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-100/50 dark:divide-slate-800/50">
            <tr v-for="loan in dashboardData?.loans" :key="loan.id" class="group hover:bg-slate-50/50 dark:hover:bg-slate-900/50 transition-colors">
              <td class="px-10 py-6 font-bold text-sm text-slate-400 tracking-tighter">{{ loan.id.substring(0, 8) }}...</td>
              <td class="px-10 py-6">
                <div class="flex items-center gap-3">
                  <LucideFileText class="w-4 h-4 text-slate-400" />
                  <span class="text-sm font-bold text-[#1E3A5F] dark:text-slate-200">{{ loan.document_title }}</span>
                </div>
              </td>
              <td class="px-10 py-6 text-sm font-bold text-slate-500">{{ new Date(loan.borrow_date).toLocaleDateString() }}</td>
              <td class="px-10 py-6 text-sm font-bold" :class="loan.status === 'OVERDUE' ? 'text-red-500' : 'text-slate-500'">{{ new Date(loan.due_date).toLocaleDateString() }}</td>
              <td class="px-10 py-6">
                <span :class="`px-3 py-1.5 rounded-xl text-[9px] font-black uppercase tracking-widest ${loan.status === 'active' ? 'bg-green-50 text-green-600' : loan.status === 'returned' ? 'bg-blue-50 text-blue-600' : 'bg-red-50 text-red-600'}`">
                  {{ loan.status }}
                </span>
              </td>
              <td class="px-10 py-6 text-right">
                <button class="p-2 text-slate-300 hover:text-[#1E3A5F] transition-colors"><LucideMoreVertical class="w-5 h-5" /></button>
              </td>
            </tr>
            <tr v-if="!dashboardData?.loans?.length">
              <td colspan="6" class="px-10 py-10 text-center text-slate-400 font-bold text-xs uppercase tracking-widest">
                {{ $t('dashboard.user.loan_history.no_records') }}
              </td>
            </tr>
          </tbody>
        </table>
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
  LucidePlus, 
  LucideTrendingUp, 
  LucideUpload, 
  LucideFileSearch, 
  LucideSearch, 
  LucideLocateFixed,
  LucideFileText,
  LucideInfo,
  LucideMoreVertical,
  LucideBookOpen,
  LucideActivity,
  LucideEdit,
  LucideTrash
} from 'lucide-vue-next'
import { useAuthStore } from '~/stores/auth'
import { useApi } from '~/composables/useApi'
import PageHeader from '~/components/PageHeader.vue'
import Modal from '~/components/Modal.vue'
import { timeAgo, parseMarkdown } from '~/utils/format'

const { t } = useI18n()
const auth = useAuthStore()
const { $api } = useApi()
const user = computed(() => auth.user)
const config = useRuntimeConfig()
const isNotesOpen = ref(false)

// Data Fetching
const { data: dashboardData, refresh: refreshDashboard } = await useAsyncData('dashboard-summary', async () => {
  const res = await $api('/dashboard/summary')
  return res.data
})

const statsSummary = computed(() => dashboardData.value?.stats)

const getActivityStyles = (action) => {
  switch (action) {
    case 'CREATE': return { dot: 'bg-green-500', icon: LucidePlus, color: 'text-green-500', bg: 'bg-green-50 dark:bg-green-900/20' }
    case 'UPDATE': return { dot: 'bg-primary-500', icon: LucideEdit, color: 'text-primary-500', bg: 'bg-primary-50 dark:bg-primary-900/20' }
    case 'DELETE': return { dot: 'bg-red-500', icon: LucideTrash, color: 'text-red-500', bg: 'bg-red-50 dark:bg-red-900/20' }
    case 'UPLOAD': return { dot: 'bg-teal-500', icon: LucideUpload, color: 'text-teal-500', bg: 'bg-teal-50 dark:bg-teal-900/20' }
    default: return { dot: 'bg-slate-300', icon: LucideActivity, color: 'text-slate-500', bg: 'bg-slate-50 dark:bg-slate-900/20' }
  }
}

const quickActions = [
  { label: t('dashboard.user.quick_actions.items.submit'), icon: LucideUpload, path: '/documents/upload' },
  { label: t('dashboard.user.quick_actions.items.loan'), icon: LucideBookOpen, path: '/loans/request' },
  { label: t('dashboard.user.quick_actions.items.search'), icon: LucideSearch, path: '/documents' },
  { label: t('dashboard.user.quick_actions.items.track'), icon: LucideLocateFixed, path: '/documents' },
]


const handleTaskAction = (task) => {
  if (task.entity_type === 'document_rejection') {
    navigateTo(`/documents/upload?id=${task.entity_id}`)
  } else if (task.entity_type.includes('Scan')) {
    navigateTo('/stock/scan')
  } else if (task.entity_type.includes('Loan')) {
    navigateTo('/approvals/loans')
  } else {
    navigateTo('/approvals/submissions')
  }
}

const formatTaskType = (type) => {
  switch (type) {
    case 'document_rejection': return 'Revisi Dokumen'
    case 'document_upload': return 'Persetujuan Dokumen'
    case 'bulk_scan': return 'Pemindaian Massal'
    case 'loan_request': return 'Permohonan Pinjaman'
    default: return type
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

