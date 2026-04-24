<template>
  <div class="max-w-7xl mx-auto space-y-8 pb-32">
    <!-- Header -->
    <div class="flex items-center justify-between" v-motion-fade>
      <div>
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight">{{ $t('manager_review.title') }}</h1>
        <p class="text-slate-500 font-medium mt-1">{{ $t('manager_review.subtitle') }}</p>
      </div>
      <div class="flex items-center gap-4">
        <span class="px-4 py-2 bg-blue-50 dark:bg-blue-900/20 text-[#1E3A5F] dark:text-blue-400 rounded-xl text-xs font-black uppercase tracking-widest border border-blue-100 dark:border-blue-800">
          12 Pending
        </span>
      </div>
    </div>

    <!-- Filters Bar -->
    <div class="grid grid-cols-1 md:grid-cols-12 gap-4" v-motion-fade>
      <!-- Search -->
      <div class="md:col-span-5 relative group">
        <LucideSearch class="absolute left-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 group-focus-within:text-[#1E3A5F] transition-colors" />
        <input 
          type="text" 
          :placeholder="$t('manager_review.filters.search')"
          class="w-full pl-12 pr-6 py-4 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-2xl text-sm font-bold outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all shadow-sm"
        />
      </div>

      <!-- Department Filter -->
      <div class="md:col-span-3 relative">
        <select class="w-full pl-6 pr-12 py-4 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-2xl text-sm font-bold appearance-none outline-none focus:border-primary-500 transition-all shadow-sm cursor-pointer">
          <option selected>{{ $t('manager_review.filters.dept_placeholder') }}</option>
          <option>Finance</option>
          <option>IT Support</option>
          <option>Human Resources</option>
        </select>
        <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 pointer-events-none" />
      </div>

      <!-- Date Range Filter -->
      <div class="md:col-span-4 relative group">
        <LucideCalendar class="absolute left-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 group-focus-within:text-[#1E3A5F] transition-colors" />
        <input 
          type="text" 
          placeholder="Oct 01, 2023 - Oct 31, 2023"
          class="w-full pl-12 pr-6 py-4 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-2xl text-sm font-bold outline-none focus:border-primary-500 transition-all shadow-sm"
        />
      </div>
    </div>

    <!-- Approvals Table -->
    <div class="glass rounded-[2rem] overflow-hidden shadow-xl shadow-slate-200/50 dark:shadow-none" v-motion-fade>
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-slate-50/50 dark:bg-slate-900/50 border-b border-slate-100 dark:border-slate-800">
            <th class="p-6 w-16">
              <div @click="toggleAll" class="w-6 h-6 rounded-lg border-2 border-slate-200 dark:border-slate-700 flex items-center justify-center cursor-pointer hover:border-primary-500 transition-all">
                <LucideCheck v-if="isAllSelected" class="w-3.5 h-3.5 text-primary-500" />
              </div>
            </th>
            <th class="p-6 text-[10px] font-black text-slate-400 uppercase tracking-[0.1em]">{{ $t('manager_review.table.reg_id') }}</th>
            <th class="p-6 text-[10px] font-black text-slate-400 uppercase tracking-[0.1em]">{{ $t('manager_review.table.submitted_by') }}</th>
            <th class="p-6 text-[10px] font-black text-slate-400 uppercase tracking-[0.1em]">{{ $t('manager_review.table.doc_info') }}</th>
            <th class="p-6 text-[10px] font-black text-slate-400 uppercase tracking-[0.1em]">{{ $t('manager_review.table.submission_date') }}</th>
            <th class="p-6 text-[10px] font-black text-slate-400 uppercase tracking-[0.1em] text-center">{{ $t('manager_review.table.urgency') }}</th>
            <th class="p-6 text-[10px] font-black text-slate-400 uppercase tracking-[0.1em] text-right">{{ $t('manager_review.table.actions') }}</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-50 dark:divide-slate-800/50">
          <tr v-for="req in pendingRequests" :key="req.id" :class="`group hover:bg-slate-50/50 dark:hover:bg-slate-800/20 transition-all ${selectedIds.includes(req.id) ? 'bg-primary-50/20 dark:bg-primary-900/10' : ''}`">
            <td class="p-6">
              <div @click="toggleSelect(req.id)" class="w-6 h-6 rounded-lg border-2 border-slate-200 dark:border-slate-700 flex items-center justify-center cursor-pointer group-hover:border-primary-500 transition-all" :class="selectedIds.includes(req.id) ? 'border-primary-500 bg-primary-500/10' : ''">
                <LucideCheck v-if="selectedIds.includes(req.id)" class="w-3.5 h-3.5 text-primary-500" />
              </div>
            </td>
            <td class="p-6">
              <span class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ req.id }}</span>
            </td>
            <td class="p-6">
              <div class="flex items-center gap-4">
                <div class="w-10 h-10 rounded-full bg-slate-100 dark:bg-slate-800 overflow-hidden border-2 border-white dark:border-slate-900 shadow-sm">
                  <img :src="req.avatar" class="w-full h-full object-cover" />
                </div>
                <div>
                  <p class="text-sm font-black text-[#1E3A5F] dark:text-white leading-tight">{{ req.user }}</p>
                  <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest mt-0.5">{{ req.dept }}</p>
                </div>
              </div>
            </td>
            <td class="p-6">
              <div>
                <p class="text-sm font-bold text-slate-700 dark:text-slate-200 leading-tight">{{ req.title }}</p>
                <div class="flex items-center gap-2 mt-1">
                  <span class="text-[10px] font-bold text-slate-400">{{ req.type }}</span>
                  <span class="w-1 h-1 rounded-full bg-slate-300"></span>
                  <span class="text-[10px] font-black text-[#1E3A5F] dark:text-blue-400 uppercase tracking-tighter">{{ $t('manager_review.table.files', { count: req.files }) }}</span>
                </div>
              </div>
            </td>
            <td class="p-6">
              <div>
                <p class="text-xs font-black text-slate-500 uppercase tracking-tight">{{ req.date }}</p>
                <p class="text-[10px] font-bold text-slate-400 mt-0.5">{{ req.time }}</p>
              </div>
            </td>
            <td class="p-6">
              <div class="flex justify-center">
                <span :class="`px-3 py-1 rounded-full text-[9px] font-black uppercase tracking-[0.1em] ${req.urgency === 'URGENT' ? 'bg-red-50 text-red-500 border border-red-100' : 'bg-slate-50 text-slate-400 border border-slate-100'}`">
                  {{ req.urgency }}
                </span>
              </div>
            </td>
            <td class="p-6 text-right">
              <button class="inline-flex items-center gap-2 px-6 py-2.5 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest hover:bg-[#152943] transition-all group shadow-lg shadow-blue-900/10">
                {{ $t('manager_review.table.review_btn') }}
                <LucideArrowRight class="w-3.5 h-3.5 group-hover:translate-x-1 transition-transform" />
              </button>
            </td>
          </tr>
        </tbody>
      </table>

      <!-- Table Footer / Pagination -->
      <div class="p-6 bg-slate-50/30 dark:bg-slate-900/30 flex items-center justify-between border-t border-slate-100 dark:border-slate-800">
        <p class="text-[11px] font-bold text-slate-400 uppercase tracking-widest">
          {{ $t('manager_review.table.showing', { start: 1, end: 3, total: 12 }) }}
        </p>
        <div class="flex gap-2">
          <button class="p-2 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-lg text-slate-400 hover:text-[#1E3A5F] transition-colors disabled:opacity-50" disabled>
            <LucideChevronLeft class="w-4 h-4" />
          </button>
          <button class="p-2 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-lg text-slate-400 hover:text-[#1E3A5F] transition-colors">
            <LucideChevronRight class="w-4 h-4" />
          </button>
        </div>
      </div>
    </div>

    <!-- Sticky Selection Footer -->
    <Transition 
      enter-active-class="transition duration-500 ease-out" 
      enter-from-class="translate-y-full opacity-0" 
      enter-to-class="translate-y-0 opacity-100"
      leave-active-class="transition duration-300 ease-in"
      leave-from-class="translate-y-0 opacity-100"
      leave-to-class="translate-y-full opacity-0"
    >
      <div v-if="selectedIds.length > 0" class="fixed bottom-0 left-0 lg:left-64 right-0 p-6 glass border-t border-slate-200/50 dark:border-white/5 z-[60] shadow-[0_-10px_40px_rgba(0,0,0,0.05)]">
        <div class="max-w-7xl mx-auto flex items-center justify-between">
          <div class="flex items-center gap-5">
            <div class="w-12 h-12 rounded-2xl bg-blue-50 dark:bg-blue-900/20 flex items-center justify-center text-[#1E3A5F] dark:text-blue-400 border border-blue-100 dark:border-blue-800">
              <LucideCheckCircle2 class="w-6 h-6" />
            </div>
            <div>
              <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">
                {{ $t('manager_review.footer.selected', { count: selectedIds.length }) }}
              </p>
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mt-0.5">
                {{ $t('manager_review.footer.bulk_hint') }}
              </p>
            </div>
          </div>
          
          <div class="flex items-center gap-6">
            <button class="text-xs font-black uppercase tracking-widest text-slate-400 hover:text-slate-600 transition-colors" @click="selectedIds = []">
              {{ $t('manager_review.footer.btn_cancel') }}
            </button>
            <div class="flex items-center gap-3">
              <button class="px-8 py-3.5 bg-red-500 hover:bg-red-600 text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-lg shadow-red-500/20 transition-all flex items-center gap-3">
                <LucideXCircle class="w-4 h-4" />
                {{ $t('manager_review.footer.btn_reject') }}
              </button>
              <button class="px-8 py-3.5 bg-green-500 hover:bg-green-600 text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-lg shadow-green-500/20 transition-all flex items-center gap-3">
                <LucideCheckCircle2 class="w-4 h-4" />
                {{ $t('manager_review.footer.btn_approve') }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { 
  LucideSearch, LucideChevronDown, LucideCalendar, LucideCheck, 
  LucideArrowRight, LucideChevronLeft, LucideChevronRight,
  LucideCheckCircle2, LucideXCircle
} from 'lucide-vue-next'

const { t } = useI18n()

const selectedIds = ref([])

const pendingRequests = [
  {
    id: 'REG-2023-9941',
    user: 'Sarah Jenkins',
    avatar: 'https://i.pravatar.cc/150?u=sarah',
    dept: 'FINANCE DEPT',
    title: 'Q3 Tax Compliance Form',
    type: 'Legal Document',
    files: 4,
    date: '24 OCT 2023',
    time: '09:45 AM',
    urgency: 'URGENT'
  },
  {
    id: 'REG-2023-9938',
    user: 'Michael Chen',
    avatar: 'https://i.pravatar.cc/150?u=michael',
    dept: 'IT SUPPORT',
    title: 'Server Upgrade Procurement',
    type: 'Invoice',
    files: 1,
    date: '24 OCT 2023',
    time: '08:12 AM',
    urgency: 'NORMAL'
  },
  {
    id: 'REG-2023-9912',
    user: 'Amanda S. Rio',
    avatar: 'https://i.pravatar.cc/150?u=amanda',
    dept: 'HUMAN RESOURCES',
    title: 'Employment Contract Update',
    type: 'Contract',
    files: 2,
    date: '20 OCT 2023',
    time: '02:55 PM',
    urgency: 'NORMAL'
  }
]

const toggleSelect = (id) => {
  if (selectedIds.value.includes(id)) {
    selectedIds.value = selectedIds.value.filter(i => i !== id)
  } else {
    selectedIds.value.push(id)
  }
}

const isAllSelected = computed(() => {
  return pendingRequests.length > 0 && selectedIds.value.length === pendingRequests.length
})

const toggleAll = () => {
  if (isAllSelected.value) {
    selectedIds.value = []
  } else {
    selectedIds.value = pendingRequests.map(r => r.id)
  }
}
</script>

<style scoped>
.glass {
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 255, 255, 0.3);
}

.dark .glass {
  background: rgba(15, 23, 42, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.05);
}
</style>
