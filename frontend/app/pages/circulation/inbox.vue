<template>
  <div class="h-[calc(100vh-120px)] flex flex-col lg:flex-row gap-0 overflow-hidden -m-10" v-motion-fade>
    <!-- Left Side: Inbox List -->
    <div class="flex-grow bg-slate-50 dark:bg-slate-950 flex flex-col border-r border-slate-200 dark:border-slate-800">
      <div class="p-10 space-y-8 bg-white dark:bg-slate-900/50">
        <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
          <div class="space-y-1">
            <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('circulation.inbox.title') }}</h1>
            <p class="text-xs font-bold text-slate-400">{{ $t('circulation.inbox.subtitle') }}</p>
          </div>
          <div class="flex bg-slate-100 dark:bg-slate-800 p-1 rounded-xl">
            <button v-for="t in ['all', 'pending', 'completed', 'archived']" :key="t" :class="`px-4 py-2 rounded-lg text-[9px] font-black uppercase tracking-widest transition-all ${activeTab === t ? 'bg-white dark:bg-slate-900 text-[#1E3A5F] dark:text-white shadow-sm' : 'text-slate-400 hover:text-slate-600'}`" @click="activeTab = t">
              {{ $t(`circulation.inbox.tabs.${t}`) }}
            </button>
          </div>
        </div>
      </div>

      <div class="flex-grow overflow-y-auto p-10 space-y-4 custom-scrollbar">
        <div v-for="task in tasks" :key="task.id" :class="`p-8 rounded-[2rem] border transition-all cursor-pointer group hover:scale-[1.01] active:scale-[0.99] ${selectedTask?.id === task.id ? 'bg-white dark:bg-slate-900 border-[#1E3A5F] shadow-2xl shadow-blue-900/10' : 'bg-white/50 dark:bg-slate-900/30 border-slate-100 dark:border-slate-800 hover:border-slate-200'}`" @click="selectedTask = task">
          <div class="flex items-center gap-6">
            <div :class="`w-14 h-14 rounded-2xl flex items-center justify-center text-xl transition-transform group-hover:rotate-6 ${selectedTask?.id === task.id ? 'bg-[#1E3A5F] text-white' : 'bg-slate-100 dark:bg-slate-800 text-slate-400'}`">
              <LucideFileText v-if="task.type === 'doc'" class="w-6 h-6" />
              <LucideZap v-else class="w-6 h-6" />
            </div>
            <div class="flex-grow space-y-1.5 overflow-hidden">
              <div class="flex items-center justify-between">
                <span class="text-[9px] font-black text-slate-300 uppercase tracking-widest">{{ task.ref }}</span>
                <span class="text-[9px] font-black text-slate-300 uppercase">{{ task.time }}</span>
              </div>
              <h3 :class="`text-sm font-black truncate transition-colors ${selectedTask?.id === task.id ? 'text-[#1E3A5F] dark:text-white' : 'text-slate-600 dark:text-slate-300'}`">{{ task.title }}</h3>
              <div class="flex items-center gap-6">
                <div class="flex items-center gap-2 text-[9px] font-bold text-slate-400 uppercase">
                  <LucideUser class="w-3 h-3" />
                  {{ task.sender }}
                </div>
                <div class="flex items-center gap-2 text-[9px] font-bold text-slate-400 uppercase">
                  <LucideCalendar class="w-3 h-3" />
                  {{ task.date }}
                </div>
                <span v-if="task.badge" :class="`px-2 py-0.5 rounded text-[7px] font-black uppercase tracking-widest border ${task.badge === 'NEW' ? 'bg-blue-50 text-blue-500 border-blue-100' : task.badge === 'SIGNED' ? 'bg-green-50 text-green-500 border-green-100' : 'bg-amber-50 text-amber-500 border-amber-100'}`">{{ task.badge }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Right Side: Task Details -->
    <div class="w-full lg:w-[450px] bg-white dark:bg-slate-900 flex flex-col shadow-2xl z-10">
      <template v-if="selectedTask">
        <div class="flex-grow overflow-y-auto p-10 space-y-12 custom-scrollbar" v-motion-slide-right>
          <div class="space-y-6">
            <div class="flex items-center justify-between">
              <span class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ selectedTask.ref }}</span>
              <span class="px-3 py-1 bg-green-50 text-green-500 rounded-full text-[8px] font-black uppercase tracking-widest border border-green-100">Review Complete</span>
            </div>
            <h2 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase leading-tight tracking-tight">{{ selectedTask.title }}</h2>
          </div>

          <!-- Status Card -->
          <div class="p-8 bg-slate-50 dark:bg-slate-800/50 rounded-lg border border-slate-100 dark:border-slate-800 flex items-center gap-6">
            <div class="w-12 h-12 rounded-xl bg-[#1E3A5F] text-white flex items-center justify-center shadow-lg shadow-blue-900/20"><LucideShieldCheck class="w-6 h-6" /></div>
            <div class="space-y-1">
              <p class="text-[9px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('circulation.inbox.details.status_update') }}</p>
              <p class="text-[11px] font-bold text-slate-500 leading-tight">{{ $t('circulation.inbox.details.status_msg') }}</p>
            </div>
          </div>

          <!-- Info Grid -->
          <div class="grid grid-cols-2 gap-6">
            <div v-for="info in [
              { label: $t('circulation.inbox.details.info.sender'), val: selectedTask.sender, sub: 'Legal Department', icon: LucideUser },
              { label: $t('circulation.inbox.details.info.classification'), val: 'Highly Confidential', sub: 'External Partner', icon: LucideLock },
              { label: $t('circulation.inbox.details.info.length'), val: '6 Pages', sub: 'PDF Format (4.2 MB)', icon: LucideFileText },
              { label: $t('circulation.inbox.details.info.due_date'), val: 'Oct 30, 2026', sub: $t('circulation.inbox.details.info.days_left', { count: 8 }), icon: LucideCalendar, color: 'text-red-500' }
            ]" :key="info.label" class="p-6 bg-slate-50/50 dark:bg-slate-800/30 rounded-3xl border border-slate-50 dark:border-slate-800 space-y-4">
              <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ info.label }}</p>
              <div class="space-y-1">
                <p :class="`text-xs font-black uppercase tracking-tight ${info.color || 'text-[#1E3A5F] dark:text-white'}`">{{ info.val }}</p>
                <p :class="`text-[8px] font-bold uppercase ${info.color || 'text-slate-400'}`">{{ info.sub }}</p>
              </div>
            </div>
          </div>

          <!-- Actions -->
          <div class="space-y-4">
            <button class="w-full py-5 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40 hover:bg-[#152943] transition-all flex items-center justify-center gap-3 active:scale-95">
              <LucideDownload class="w-5 h-5" />
              {{ $t('circulation.inbox.details.actions.download') }}
            </button>
            <div class="grid grid-cols-2 gap-4">
              <button class="py-4 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-slate-600 dark:text-slate-300 rounded-xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center justify-center gap-2">
                <LucideEye class="w-4 h-4" />
                {{ $t('circulation.inbox.details.actions.preview') }}
              </button>
              <button class="py-4 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-slate-600 dark:text-slate-300 rounded-xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center justify-center gap-2">
                <LucideStamp class="w-4 h-4" />
                {{ $t('circulation.inbox.details.actions.disposisi') }}
              </button>
            </div>
          </div>

          <!-- Routing History -->
          <div class="space-y-8">
            <div class="flex items-center justify-between border-b border-slate-100 dark:border-slate-800 pb-4">
              <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.inbox.routing.title') }}</h4>
              <LucideChevronDown class="w-4 h-4 text-slate-300" />
            </div>
            <div class="space-y-10 relative before:absolute before:left-3.5 before:top-2 before:bottom-2 before:w-px before:bg-slate-100 dark:before:bg-slate-800">
              <div v-for="(step, i) in [
                { name: $t('circulation.inbox.routing.steps.head_legal'), status: 'Signature Recorded', time: 'Oct 24, 09:12', done: true },
                { name: $t('circulation.inbox.routing.steps.gm_legal') + ' ' + $t('circulation.inbox.routing.you'), status: 'Approved & Verified', time: 'Oct 24, 14:20', active: true },
                { name: $t('circulation.inbox.routing.steps.section_head'), status: 'Pending Signature', pending: true },
                { name: $t('circulation.inbox.routing.steps.archive'), status: 'Final Documentation', future: true }
              ]" :key="step.name" class="relative pl-12 space-y-1">
                <div :class="`absolute left-2 top-2 w-3 h-3 rounded-full border-4 border-white dark:border-slate-900 z-10 transition-colors ${step.done ? 'bg-green-500' : step.active ? 'bg-[#1E3A5F] ring-4 ring-blue-50 dark:ring-blue-900/20' : 'bg-slate-200 dark:bg-slate-800'}`"></div>
                <div class="flex items-center justify-between">
                  <p :class="`text-[11px] font-black uppercase ${step.active ? 'text-[#1E3A5F] dark:text-white' : 'text-slate-400'}`">{{ step.name }}</p>
                  <span v-if="step.time" class="text-[9px] font-bold text-slate-300 uppercase">{{ step.time }}</span>
                </div>
                <p :class="`text-[9px] font-bold uppercase tracking-tight ${step.done || step.active ? 'text-slate-500' : 'text-slate-300'}`">{{ step.status }}</p>
              </div>
            </div>
          </div>
        </div>
      </template>
      <div v-else class="flex-grow flex flex-col items-center justify-center p-20 text-center space-y-6 opacity-30">
        <div class="w-32 h-32 bg-slate-50 dark:bg-slate-800 rounded-full flex items-center justify-center text-slate-300"><LucideInbox class="w-16 h-16" /></div>
        <div class="space-y-2">
          <p class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tighter">No Task Selected</p>
          <p class="text-xs font-bold text-slate-400">Select a document from the inbox to view details.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { 
  LucideFileText, LucideZap, LucideUser, LucideCalendar, 
  LucideLock, LucideShieldCheck, LucideDownload, LucideEye, 
  LucideStamp, LucideChevronDown, LucideInbox 
} from 'lucide-vue-next'

const activeTab = ref('all')
const selectedTask = ref(null)

const tasks = [
  { id: 1, type: 'doc', ref: 'LGL.EXT.2026.0045', title: 'Lease Agreement - PT Global Tech Solutions', sender: 'Anita Wijaya (Legal)', date: 'Oct 24, 2026', time: '14:20 PM', badge: 'NEW' },
  { id: 2, type: 'doc', ref: 'HR.INT.2026.0882', title: 'Employee Handbook Revision V4', sender: 'Budi Santoso (HR)', date: 'Oct 23, 2026', time: 'Oct 23, 2026', badge: 'SIGNED' },
  { id: 3, type: 'doc', ref: 'FIN.PRO.2026.1104', title: 'Vendor Procurement: Hardware Phase 2', sender: 'Siska Pratama (Finance)', date: 'Oct 22, 2026', time: 'Oct 22, 2026', badge: 'DOWNLOADED' }
]

onMounted(() => {
  selectedTask.value = tasks[0]
})

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #E2E8F0; border-radius: 10px; }
.dark .custom-scrollbar::-webkit-scrollbar-thumb { background: #1E293B; }
</style>
