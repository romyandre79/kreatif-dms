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
          <button class="p-3 bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 text-slate-400 relative hover:text-primary-500 transition-colors shadow-sm">
            <LucideBell class="w-5 h-5" />
            <span class="absolute top-3 right-3 w-2 h-2 bg-red-500 rounded-full border-2 border-white dark:border-slate-900"></span>
          </button>
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
          <p class="text-3xl font-black text-slate-900 dark:text-white">128</p>
          <div class="flex items-center gap-1 text-[10px] font-black text-green-500 mt-1 uppercase tracking-tighter">
            <LucideTrendingUp class="w-3 h-3" />
            {{ $t('dashboard.stats.vs_yesterday', { percent: 12 }) }}
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
          <p class="text-3xl font-black text-slate-900 dark:text-white">94</p>
          <div class="flex items-center gap-1 text-[10px] font-black text-green-500 mt-1 uppercase tracking-tighter">
            <LucideTrendingUp class="w-3 h-3" />
            {{ $t('dashboard.stats.vs_yesterday', { percent: 5 }) }}
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
          <p class="text-3xl font-black text-slate-900 dark:text-white">82</p>
          <div class="flex items-center gap-1 text-[10px] font-black text-green-500 mt-1 uppercase tracking-tighter">
            <LucideTrendingUp class="w-3 h-3" />
            {{ $t('dashboard.stats.vs_yesterday', { percent: 8 }) }}
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
          <p class="text-2xl font-black text-slate-900 dark:text-white">7.2TB / 10TB</p>
          <div class="h-1.5 w-full bg-slate-100 dark:bg-slate-800 rounded-full mt-3 overflow-hidden">
            <div class="h-full bg-teal-500 rounded-full" style="width: 72%"></div>
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
          <p class="text-4xl font-black">31</p>
          <p class="text-[10px] font-bold text-teal-50 mt-1">{{ $t('dashboard.stats.immediate_attention') }}</p>
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
              <tr v-for="task in taskQueue" :key="task.id" class="hover:bg-slate-50/50 dark:hover:bg-slate-800/20 transition-colors">
                <td class="px-8 py-6">
                  <p class="font-black text-sm text-[#1E3A5F] dark:text-slate-200">{{ task.id }}</p>
                  <p class="text-[10px] text-slate-400 font-bold mt-0.5">{{ task.desc }}</p>
                </td>
                <td class="px-8 py-6">
                  <span class="text-sm font-bold text-slate-600 dark:text-slate-400">{{ task.type }}</span>
                </td>
                <td class="px-8 py-6">
                  <span class="text-xs font-black text-primary-600 uppercase tracking-tighter">{{ task.status }}</span>
                </td>
                <td class="px-8 py-6 text-right">
                  <span :class="`px-3 py-1 rounded-lg text-[9px] font-black uppercase ${task.urgencyBg}`">
                    {{ task.urgency }}
                  </span>
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
          <div v-for="(act, index) in activities" :key="index" class="flex gap-6 relative">
            <div v-if="index !== activities.length - 1" class="absolute left-3.5 top-8 w-0.5 h-10 bg-slate-100 dark:bg-slate-800"></div>
            <div :class="`w-7 h-7 rounded-full flex-shrink-0 flex items-center justify-center z-10 border-4 border-white dark:border-slate-900 ${act.bg}`">
              <component :is="act.icon" :class="`w-3 h-3 ${act.color}`" />
            </div>
            <div>
              <p class="text-sm font-bold text-[#1E3A5F] dark:text-slate-200">{{ act.title }}</p>
              <p class="text-xs text-slate-400 mt-1 leading-relaxed">{{ act.desc }}</p>
              <p class="text-[10px] text-slate-300 dark:text-slate-500 font-black uppercase mt-1.5">{{ act.time }}</p>
            </div>
          </div>
        </div>
        <button class="w-full mt-10 py-3.5 bg-slate-50 dark:bg-slate-800/50 rounded-xl text-[10px] font-black uppercase tracking-widest text-slate-400 hover:bg-slate-100 transition-all">
          {{ $t('dashboard.activity.btn_logs') }}
        </button>
      </div>
    </div>

    <!-- Infrastructure Status -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6" v-motion-slide-visible-bottom>
      <!-- Scanner -->
      <div class="glass p-6 rounded-2xl flex items-center gap-5 group hover:border-green-500/30 transition-all">
        <div class="w-14 h-14 rounded-2xl bg-green-50 dark:bg-green-900/20 flex items-center justify-center text-green-500">
          <LucidePrinter class="w-7 h-7" />
        </div>
        <div>
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('dashboard.infrastructure.scanner') }}</p>
          <p class="text-sm font-black text-slate-800 dark:text-white">Canon DR-M260</p>
          <p class="text-[10px] font-bold text-green-500 uppercase">{{ $t('dashboard.infrastructure.connected') }}</p>
        </div>
      </div>

      <!-- OCR Engine -->
      <div class="glass p-6 rounded-2xl flex items-center gap-5 group hover:border-primary-500/30 transition-all">
        <div class="w-14 h-14 rounded-2xl bg-primary-50 dark:bg-primary-900/20 flex items-center justify-center text-primary-500">
          <LucideCpu class="w-7 h-7" />
        </div>
        <div>
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('dashboard.infrastructure.ocr_engine') }}</p>
          <p class="text-sm font-black text-slate-800 dark:text-white">Abbyy Finereader</p>
          <p class="text-[10px] font-bold text-green-500 uppercase">{{ $t('dashboard.infrastructure.active_accuracy', { accuracy: 99.2 }) }}</p>
        </div>
      </div>

      <!-- Database -->
      <div class="glass p-6 rounded-2xl flex items-center gap-5 group hover:border-teal-500/30 transition-all">
        <div class="w-14 h-14 rounded-2xl bg-teal-50 dark:bg-teal-900/20 flex items-center justify-center text-teal-500">
          <LucideDatabase class="w-7 h-7" />
        </div>
        <div>
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('dashboard.infrastructure.database') }}</p>
          <p class="text-sm font-black text-slate-800 dark:text-white">Oracle Cloud</p>
          <p class="text-[10px] font-bold text-green-500 uppercase">{{ $t('dashboard.infrastructure.synchronized') }}</p>
        </div>
      </div>
    </div>
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
  LucidePlus
} from 'lucide-vue-next'
import { useAuthStore } from '~/stores/auth'
import PageHeader from '~/components/PageHeader.vue'

const auth = useAuthStore()
const { t } = useI18n()
const user = computed(() => auth.user)

const currentDate = computed(() => {
  const d = new Date();
  const day = String(d.getDate()).padStart(2, '0');
  const month = String(d.getMonth() + 1).padStart(2, '0');
  const year = d.getFullYear();
  return `${day}.${month}.${year}`;
})

const taskQueue = [
  { id: 'MFST-2023-0941', desc: 'Finance Dept (Legal Vouchers)', type: t('dashboard.queue.items.bulk_scan'), status: t('dashboard.queue.items.on_schedule'), urgency: 'HIGH', urgencyBg: 'bg-red-50 text-red-500' },
  { id: 'MFST-2023-0942', desc: 'HR Services (Employee Files)', type: t('dashboard.queue.items.indexing'), status: t('dashboard.queue.items.on_schedule'), urgency: 'MEDIUM', urgencyBg: 'bg-orange-50 text-orange-500' },
  { id: 'MFST-2023-0899', desc: 'Procurement (Vendor Contracts)', type: t('dashboard.queue.items.receive'), status: t('dashboard.queue.items.on_schedule'), urgency: 'LOW', urgencyBg: 'bg-slate-50 text-slate-400' },
  { id: 'MFST-2023-0945', desc: 'Board Records (Annual 2022)', type: t('dashboard.queue.items.bulk_scan'), status: t('dashboard.queue.items.on_schedule'), urgency: 'HIGH', urgencyBg: 'bg-red-50 text-red-500' },
]

const activities = [
  { title: t('dashboard.activity.items.ocr_completed', { id: '2023-0948' }), desc: t('dashboard.activity.items.ocr_sub', { count: 42 }), time: t('dashboard.activity.items.mins_ago', { count: 12 }), icon: LucideFileStack, color: 'text-primary-500', bg: 'bg-primary-50 dark:bg-primary-900/20' },
  { title: t('retention.approaching.title'), desc: '340 documents from Finance Dept are reaching 10-year retention limit.', time: t('dashboard.activity.items.mins_ago', { count: 45 }), icon: LucideAlertCircle, color: 'text-red-500', bg: 'bg-red-50 dark:bg-red-900/20' },
  { title: t('dashboard.infrastructure.database'), desc: 'Zone A-04 marked as 95% full after bulk deposit MFST-0881.', time: t('dashboard.activity.items.hour_ago'), icon: LucideLayers, color: 'text-teal-500', bg: 'bg-teal-50 dark:bg-teal-900/20' },
  { title: t('dashboard.activity.items.security_login'), desc: t('dashboard.activity.items.security_sub', { ip: '10.20.44.12' }), time: '4 hours ago', icon: LucideUser, color: 'text-green-500', bg: 'bg-green-50 dark:bg-green-900/20' },
]
</script>
