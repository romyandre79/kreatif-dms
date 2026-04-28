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
      <div v-motion-slide-visible-bottom class="glass p-6 rounded-2xl relative overflow-hidden group">
        <p class="text-xs text-slate-400 font-bold uppercase tracking-wider mb-4">{{ $t('dashboard.user.stats.my_submissions') }}</p>
        <div class="flex items-end justify-between">
          <p class="text-4xl font-black text-slate-900 dark:text-white">24</p>
          <div class="flex items-center gap-1 text-green-500 text-xs font-bold bg-green-500/10 px-2 py-1 rounded-lg">
            <LucideTrendingUp class="w-3 h-3" />
            +12%
          </div>
        </div>
      </div>

      <!-- Pending Approvals -->
      <div v-motion-slide-visible-bottom class="glass p-6 rounded-2xl relative overflow-hidden group">
        <p class="text-xs text-slate-400 font-bold uppercase tracking-wider mb-4">{{ $t('dashboard.user.stats.pending_approvals') }}</p>
        <div class="flex items-center justify-between">
          <p class="text-4xl font-black text-slate-900 dark:text-white">3</p>
          <div class="bg-orange-100 dark:bg-orange-900/30 text-orange-600 dark:text-orange-400 px-3 py-1.5 rounded-lg text-[10px] font-black uppercase leading-tight" v-html="$t('dashboard.user.stats.action_required')">
          </div>
        </div>
      </div>

      <!-- Active Loans -->
      <div v-motion-slide-visible-bottom class="glass p-6 rounded-2xl relative overflow-hidden group">
        <p class="text-xs text-slate-400 font-bold uppercase tracking-wider mb-4">{{ $t('dashboard.user.stats.active_loans') }}</p>
        <div class="flex items-center justify-between">
          <p class="text-4xl font-black text-slate-900 dark:text-white">2</p>
          <div class="bg-green-100 dark:bg-green-900/30 text-green-600 dark:text-green-400 px-3 py-1 rounded-full text-[10px] font-black uppercase">
            {{ $t('dashboard.infrastructure.active_accuracy', { accuracy: '' }).split('(')[0].trim() }}
          </div>
        </div>
      </div>

      <!-- Search History -->
      <div v-motion-slide-visible-bottom class="glass p-6 rounded-2xl relative overflow-hidden group">
        <p class="text-xs text-slate-400 font-bold uppercase tracking-wider mb-4">{{ $t('dashboard.user.stats.search_history') }}</p>
        <div class="flex items-end justify-between">
          <p class="text-4xl font-black text-slate-900 dark:text-white">47</p>
          <p class="text-slate-400 text-xs font-medium mb-1">{{ $t('dashboard.user.stats.today') }}</p>
        </div>
      </div>
    </div>

    <!-- Middle Grid -->
    <div class="grid grid-cols-1 xl:grid-cols-3 gap-8">
      <!-- Items Requiring My Action -->
      <div class="xl:col-span-2 glass rounded-2xl overflow-hidden" v-motion-slide-visible-bottom>
        <div class="px-6 py-5 border-b border-slate-200 dark:border-slate-800 flex items-center justify-between">
          <h3 class="font-black text-lg text-slate-800 dark:text-white">{{ $t('dashboard.user.tasks.title') }}</h3>
          <button class="text-primary-600 dark:text-primary-400 text-sm font-bold hover:underline">{{ $t('dashboard.queue.view_all_tasks') }}</button>
        </div>
        <div class="overflow-x-auto">
          <table class="w-full text-left">
            <thead>
              <tr class="text-slate-400 text-[10px] uppercase font-black tracking-widest bg-slate-50/50 dark:bg-slate-900/50">
                <th class="px-6 py-4">{{ $t('dashboard.user.tasks.table.task_type') }}</th>
                <th class="px-6 py-4">{{ $t('dashboard.user.tasks.table.doc_manifest') }}</th>
                <th class="px-6 py-4">{{ $t('dashboard.user.tasks.table.due_date') }}</th>
                <th class="px-6 py-4 text-right">{{ $t('dashboard.user.tasks.table.action') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-100 dark:divide-slate-800">
              <tr v-for="task in tasks" :key="task.id" class="group hover:bg-slate-50/50 dark:hover:bg-slate-900/50 transition-colors">
                <td class="px-6 py-5">
                  <div class="flex items-center gap-3">
                    <div :class="`w-2 h-2 rounded-full ${task.dotColor}`"></div>
                    <span class="font-bold text-sm text-slate-700 dark:text-slate-200">{{ task.type }}</span>
                  </div>
                </td>
                <td class="px-6 py-5">
                  <span class="text-xs font-bold text-slate-400 tracking-tighter">{{ task.id }}</span>
                </td>
                <td class="px-6 py-5">
                  <div class="text-xs font-bold text-slate-500">
                    <p>{{ task.dueDate }}</p>
                    <p class="text-[10px] text-slate-400 mt-0.5">{{ task.dueTime }}</p>
                  </div>
                </td>
                <td class="px-6 py-5 text-right">
                  <button :class="`px-5 py-2 rounded-lg text-xs font-black transition-all ${task.btnClass}`">
                    {{ task.btnLabel }}
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Quick Actions -->
      <div class="glass rounded-3xl p-8" v-motion-slide-visible-bottom>
        <h3 class="font-black text-lg text-slate-800 dark:text-white mb-8">{{ $t('dashboard.user.quick_actions.title') }}</h3>
        <div class="grid grid-cols-2 gap-4">
          <button v-for="action in quickActions" :key="action.label" class="flex flex-col items-center justify-center p-6 bg-slate-50 dark:bg-slate-900/50 rounded-2xl hover:bg-primary-50 dark:hover:bg-primary-900/20 transition-all border border-transparent hover:border-primary-200 group">
            <div class="w-12 h-12 rounded-xl flex items-center justify-center mb-4 text-slate-400 group-hover:text-primary-500 transition-colors">
              <component :is="action.icon" class="w-6 h-6" />
            </div>
            <span class="text-[10px] font-black uppercase tracking-widest text-slate-600 dark:text-slate-400 group-hover:text-primary-600">{{ action.label }}</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Bottom Grid -->
    <div class="grid grid-cols-1 xl:grid-cols-3 gap-8">
      <!-- Recent Activities -->
      <div class="xl:col-span-2 glass rounded-2xl p-8" v-motion-slide-visible-bottom>
        <h3 class="font-black text-lg text-slate-800 dark:text-white mb-8">{{ $t('dashboard.user.activity.title') }}</h3>
        <div class="space-y-8">
          <div v-for="(activity, index) in activities" :key="index" class="flex gap-6 relative">
            <div v-if="index !== activities.length - 1" class="absolute left-3 top-8 w-0.5 h-8 bg-slate-100 dark:bg-slate-800"></div>
            <div :class="`w-6 h-6 rounded-full flex-shrink-0 flex items-center justify-center z-10 ${activity.dot}`">
              <div class="w-2.5 h-2.5 rounded-full bg-white dark:bg-slate-900"></div>
            </div>
            <div class="flex-1 flex justify-between items-start gap-4">
              <div>
                <p class="text-sm font-bold text-slate-600 dark:text-slate-300 leading-relaxed" v-html="activity.text"></p>
                <p class="text-xs text-slate-400 font-bold mt-1.5">{{ activity.time }}</p>
              </div>
              <button class="text-slate-400 hover:text-primary-500 text-xs font-black uppercase tracking-tighter">{{ activity.action }}</button>
            </div>
          </div>
        </div>
      </div>

      <!-- System Update Banner -->
      <div class="bg-[#1E3A5F] rounded-2xl p-8 text-white relative overflow-hidden group shadow-2xl shadow-blue-900/40" v-motion-slide-visible-bottom>
        <div class="relative z-10">
          <h3 class="font-black text-xl mb-4">{{ $t('dashboard.user.system_update.title') }}</h3>
          <p class="text-blue-100/80 text-sm leading-relaxed mb-8">
            {{ $t('dashboard.user.system_update.desc') }}
          </p>
          <button class="w-full py-3.5 bg-white text-[#1E3A5F] rounded-xl font-black text-xs uppercase tracking-widest hover:bg-blue-50 transition-all">
            {{ $t('dashboard.user.system_update.btn_notes') }}
          </button>
        </div>
        <LucideInfo class="absolute -right-8 -bottom-8 w-48 h-48 text-white/5 opacity-10 rotate-12" />
      </div>
    </div>

    <!-- Loan History Table -->
    <div class="glass rounded-2xl overflow-hidden" v-motion-slide-visible-bottom>
      <div class="px-6 py-6 border-b border-slate-200 dark:border-slate-800 flex items-center justify-between">
        <h3 class="font-black text-lg text-slate-800 dark:text-white">{{ $t('dashboard.user.loan_history.title') }}</h3>
        <div class="flex gap-2">
          <button class="px-4 py-2 bg-slate-50 dark:bg-slate-900 rounded-lg text-xs font-black text-slate-500 border border-slate-200 dark:border-slate-800 hover:bg-slate-100">{{ $t('dashboard.user.loan_history.btn_export') }}</button>
          <button class="px-4 py-2 bg-slate-50 dark:bg-slate-900 rounded-lg text-xs font-black text-slate-500 border border-slate-200 dark:border-slate-800 hover:bg-slate-100">{{ $t('dashboard.user.loan_history.btn_filter') }}</button>
        </div>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full text-left">
          <thead>
            <tr class="text-slate-400 text-[10px] uppercase font-black tracking-widest bg-slate-50/50 dark:bg-slate-900/50">
              <th class="px-6 py-4">{{ $t('dashboard.user.loan_history.table.loan_id') }}</th>
              <th class="px-6 py-4">{{ $t('dashboard.user.loan_history.table.doc_title') }}</th>
              <th class="px-6 py-4">{{ $t('dashboard.user.loan_history.table.loan_date') }}</th>
              <th class="px-6 py-4">{{ $t('dashboard.user.loan_history.table.expected_return') }}</th>
              <th class="px-6 py-4">{{ $t('dashboard.user.loan_history.table.status') }}</th>
              <th class="px-6 py-4 text-right">{{ $t('dashboard.user.loan_history.table.action') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-100 dark:divide-slate-800">
            <tr v-for="loan in loans" :key="loan.id" class="hover:bg-slate-50/50 dark:hover:bg-slate-900/50 transition-colors">
              <td class="px-6 py-5 text-xs font-bold text-slate-400">{{ loan.id }}</td>
              <td class="px-6 py-5">
                <div class="flex items-center gap-3">
                  <LucideFileText class="w-4 h-4 text-slate-400" />
                  <span class="font-bold text-sm text-slate-700 dark:text-slate-200">{{ loan.title }}</span>
                </div>
              </td>
              <td class="px-6 py-5 text-xs font-bold text-slate-500">{{ loan.date }}</td>
              <td class="px-6 py-5 text-xs font-bold text-slate-500">{{ loan.returnDate }}</td>
              <td class="px-6 py-5">
                <span 
                  class="px-2.5 py-1 rounded-lg text-[10px] font-black uppercase"
                  :class="loan.status === 'ACTIVE' ? 'bg-green-100 text-green-700 dark:bg-green-900/30 dark:text-green-400' : 'bg-slate-100 text-slate-500 dark:bg-slate-800 dark:text-slate-400'"
                >
                  {{ loan.status }}
                </span>
              </td>
              <td class="px-6 py-5 text-right">
                <button class="p-2 text-slate-400 hover:text-slate-600">
                  <LucideMoreVertical class="w-5 h-5" />
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
const { t } = useI18n()
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
  LucideBookOpen
} from 'lucide-vue-next'
import { useAuthStore } from '~/stores/auth'
import PageHeader from '~/components/PageHeader.vue'

const auth = useAuthStore()
const user = computed(() => auth.user)

const tasks = [
  { id: 'MF-2023-8982', type: t('dashboard.user.tasks.items.print_manifest'), dotColor: 'bg-orange-500', dueDate: t('dashboard.user.tasks.items.today'), dueTime: '14:00', btnLabel: t('dashboard.user.tasks.items.execute'), btnClass: 'bg-slate-100 text-slate-600 hover:bg-slate-200' },
  { id: 'DOC-A4-7721', type: t('dashboard.user.tasks.items.doc_return'), dotColor: 'bg-blue-500', dueDate: t('dashboard.user.tasks.items.tomorrow'), dueTime: '', btnLabel: t('dashboard.user.tasks.items.process'), btnClass: 'bg-slate-100 text-slate-600 hover:bg-slate-200' },
  { id: 'VLD-2023-1102', type: t('dashboard.user.tasks.items.phys_val'), dotColor: 'bg-orange-500', dueDate: 'Sep 28,', dueTime: '2023', btnLabel: t('dashboard.user.tasks.items.verify'), btnClass: 'bg-slate-100 text-slate-600 hover:bg-slate-200' },
]

const quickActions = [
  { label: t('dashboard.user.quick_actions.items.submit'), icon: LucideUpload },
  { label: t('dashboard.user.quick_actions.items.loan'), icon: LucideBookOpen },
  { label: t('dashboard.user.quick_actions.items.search'), icon: LucideSearch },
  { label: t('dashboard.user.quick_actions.items.track'), icon: LucideLocateFixed },
]

const activities = [
  { text: t('dashboard.user.activity.items.doc_approved', { id: 'AKR-V-002' }), time: t('dashboard.user.activity.items.ten_mins'), action: t('dashboard.user.activity.items.details'), dot: 'bg-green-500' },
  { text: t('dashboard.user.activity.items.new_submission', { id: 'INV-9982-Jakarta' }), time: t('dashboard.user.activity.items.one_hour'), action: t('dashboard.user.activity.items.view'), dot: 'bg-blue-500' },
  { text: t('dashboard.user.activity.items.loan_req_sig', { id: 'LR-772' }), time: t('dashboard.user.activity.items.three_hours'), action: t('dashboard.user.activity.items.sign_now'), dot: 'bg-orange-500' },
]

const loans = [
  { id: 'LN-2023-45', title: 'Financial Audit Q3 2023', date: 'Sep 20, 2023', returnDate: 'Sep 30, 2023', status: 'ACTIVE' },
  { id: 'LN-2023-44', title: 'Legal Agreement - PT ABC', date: 'Sep 15, 2023', returnDate: 'Sep 18, 2023', status: 'RETURNED' },
]
</script>
