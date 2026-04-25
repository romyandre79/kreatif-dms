<template>
  <div class="flex h-screen bg-[#F8FAFC] dark:bg-slate-950 overflow-hidden" v-motion-fade>
    <!-- Left Sidebar: Selection Directory -->
    <aside class="w-80 bg-white dark:bg-slate-900 border-r border-slate-200 dark:border-slate-800 flex flex-col p-8 overflow-y-auto custom-scrollbar">
      <div class="space-y-8 h-full flex flex-col">
        <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em]">Selection Directory</h3>
        
        <div class="relative group">
          <LucideSearch class="absolute left-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 group-focus-within:text-blue-500 transition-colors" />
          <input type="text" placeholder="Filter by Unit/ID..." class="w-full bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-700 rounded-2xl pl-12 pr-6 py-4 text-xs font-bold outline-none focus:ring-2 focus:ring-blue-500/20 transition-all" />
        </div>

        <div class="space-y-3 flex-grow">
          <button v-for="user in directory" :key="user.id" 
            class="w-full p-6 rounded-3xl transition-all text-left border-2"
            :class="user.active ? 'bg-blue-50/50 border-[#1E3A5F] shadow-lg shadow-blue-900/10' : 'bg-white dark:bg-slate-900 border-transparent hover:bg-slate-50'"
          >
            <div class="flex items-center gap-4">
              <div :class="`w-12 h-12 rounded-2xl flex items-center justify-center text-white font-black text-xs ${user.bg}`">
                {{ user.initials }}
              </div>
              <div>
                <p :class="`text-sm font-black uppercase tracking-tight ${user.active ? 'text-[#1E3A5F]' : 'text-slate-600 dark:text-slate-300'}`">{{ user.name }}</p>
                <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ user.dept }} • {{ user.id }}</p>
              </div>
            </div>
          </button>
        </div>

        <button class="w-full py-5 border-2 border-dashed border-slate-200 dark:border-slate-800 text-[10px] font-black text-slate-400 uppercase tracking-widest rounded-3xl hover:bg-slate-50 transition-all flex items-center justify-center gap-3">
          <LucidePlus class="w-4 h-4" />
          Add Filter Layer
        </button>
      </div>
    </aside>

    <!-- Main Content Area: Activity Stream -->
    <main class="flex-grow flex flex-col overflow-hidden">
      <!-- Header -->
      <header class="bg-white dark:bg-slate-900 border-b border-slate-200 dark:border-slate-800 px-10 py-8 flex items-center justify-between shadow-sm relative z-10">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">User Scorecard & Activity Timeline</h1>
        <div class="flex items-center gap-4">
          <button class="px-6 py-4 border border-slate-200 dark:border-slate-800 text-slate-600 dark:text-slate-300 rounded-2xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center gap-3 shadow-sm">
            <LucideZap class="w-4 h-4 text-blue-500" />
            Compare 2 users
          </button>
          <button class="px-8 py-4 bg-[#1E3A5F] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 active:scale-95">
            <LucideDownload class="w-4 h-4" />
            Export user report
          </button>
        </div>
      </header>

      <!-- Timeline Section -->
      <div class="flex-grow p-10 overflow-auto custom-scrollbar flex flex-col gap-10">
        <div class="glass rounded-[3rem] p-10 space-y-12 min-h-full flex flex-col">
          <div class="flex items-center justify-between px-4 border-b border-slate-50 dark:border-slate-800 pb-8 bg-slate-50/30 dark:bg-slate-900/30 rounded-t-[2.5rem] -mx-10 -mt-10 p-10">
            <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em]">Activity Stream</h3>
            <div class="flex items-center gap-8">
              <div class="flex bg-white dark:bg-slate-800 rounded-xl p-1 shadow-sm border border-slate-100 dark:border-slate-700">
                <button class="px-4 py-2 text-[9px] font-black uppercase tracking-widest text-slate-400 hover:text-[#1E3A5F]">Day</button>
                <button class="px-4 py-2 bg-slate-50 dark:bg-slate-700 rounded-lg text-[9px] font-black uppercase tracking-widest text-[#1E3A5F] dark:text-white">Week</button>
                <button class="px-4 py-2 text-[9px] font-black uppercase tracking-widest text-slate-400 hover:text-[#1E3A5F]">Month</button>
              </div>
              <div class="flex items-center gap-4 text-[10px] font-black text-slate-500">
                <span>May 12 - May 18, 2024</span>
                <div class="flex gap-2">
                  <button class="p-1 hover:text-blue-500"><LucideChevronLeft class="w-4 h-4" /></button>
                  <button class="p-1 hover:text-blue-500"><LucideChevronRight class="w-4 h-4" /></button>
                </div>
              </div>
            </div>
          </div>

          <div class="space-y-16 relative pl-12 flex-grow">
            <!-- Timeline Line -->
            <div class="absolute left-[23px] top-0 bottom-0 w-0.5 bg-slate-100 dark:bg-slate-800"></div>

            <div v-for="event in timeline" :key="event.time" class="relative group">
              <!-- Timeline Dot -->
              <div :class="`absolute -left-[12.5px] top-0 w-6 h-6 rounded-full border-4 border-white dark:border-slate-900 shadow-sm z-10 ${event.dot}`"></div>
              
              <div class="space-y-6">
                <div class="flex items-center justify-between">
                  <h4 :class="`text-[10px] font-black uppercase tracking-[0.2em] ${event.labelColor}`">{{ event.label }}</h4>
                  <span class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ event.time }}</span>
                </div>

                <div class="p-8 bg-slate-50/50 dark:bg-slate-900/50 rounded-[2rem] border border-slate-100 dark:border-slate-800 space-y-4 group-hover:bg-white dark:group-hover:bg-slate-900 transition-all shadow-sm group-hover:shadow-xl group-hover:shadow-blue-900/5" :class="event.bg">
                  <h5 class="text-base font-black text-[#1E3A5F] dark:text-white tracking-tight">{{ event.title }}</h5>
                  <p class="text-xs font-bold text-slate-500 leading-relaxed max-w-2xl" v-html="event.desc"></p>
                  
                  <div v-if="event.actions" class="flex items-center gap-6 pt-2">
                    <button v-for="action in event.actions" :key="action.text" :class="`text-[9px] font-black uppercase tracking-widest hover:underline ${action.color}`">
                      {{ action.text }}
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Bottom Input -->
          <div class="flex items-center gap-4 bg-slate-50 dark:bg-slate-800 rounded-3xl p-2 pl-6">
            <input type="text" placeholder="Submit compliance note for this user..." class="flex-grow bg-transparent border-none outline-none text-xs font-bold text-slate-600 dark:text-slate-300" />
            <button class="px-8 py-4 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-[#152943] transition-all">Submit Note</button>
          </div>
        </div>
      </div>
    </main>

    <!-- Right Sidebar: Metrics & Risk -->
    <aside class="w-80 bg-white dark:bg-slate-900 border-l border-slate-200 dark:border-slate-800 flex flex-col p-8 overflow-y-auto custom-scrollbar space-y-8">
      <div v-for="metric in metrics" :key="metric.label" class="p-8 rounded-[2.5rem] border border-slate-50 dark:border-slate-800 bg-white dark:bg-slate-900 shadow-sm space-y-6">
        <div class="flex justify-between items-start">
          <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest max-w-[120px]">{{ metric.label }}</p>
          <span v-if="metric.badge" :class="`px-2 py-0.5 rounded-lg text-[8px] font-black uppercase tracking-widest ${metric.badgeClass}`">
            {{ metric.badge }}
          </span>
        </div>
        <div class="space-y-4">
          <div class="flex items-baseline gap-2">
            <p class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ metric.value }}</p>
            <span v-if="metric.subValue" class="text-xs font-black text-slate-400 uppercase tracking-widest">{{ metric.subValue }}</span>
          </div>
          <div v-if="metric.progress" class="h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
            <div :style="{ width: metric.progress + '%' }" :class="`h-full rounded-full ${metric.progressColor}`"></div>
          </div>
          <button v-if="metric.link" class="text-[9px] font-black text-blue-500 hover:underline uppercase tracking-widest">{{ metric.link }}</button>
        </div>
      </div>

      <!-- Risk Assessment -->
      <div class="bg-[#1E3A5F] p-8 rounded-[3rem] space-y-8 shadow-2xl shadow-blue-900/30 mt-auto">
        <h3 class="text-xs font-black text-white uppercase tracking-[0.2em]">Risk Assessment</h3>
        <div class="flex gap-4">
          <div class="flex-grow p-4 bg-white/5 rounded-2xl border border-white/10 space-y-1">
            <p class="text-[7px] font-black text-blue-300 uppercase tracking-widest">Current Score</p>
            <p class="text-2xl font-black text-white tracking-tighter">720</p>
          </div>
          <div class="flex-grow p-4 bg-white/5 rounded-2xl border border-white/10 space-y-1">
            <p class="text-[7px] font-black text-blue-300 uppercase tracking-widest">Risk Level</p>
            <p class="text-sm font-black text-amber-400 uppercase tracking-widest">Moderate</p>
          </div>
        </div>
        <p class="text-[10px] font-bold text-blue-200 leading-relaxed italic">
          "Higher than normal download activity detected outside business hours."
        </p>
        <button class="w-full py-4 bg-white/10 hover:bg-white text-[#1E3A5F] hover:text-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest transition-all shadow-lg border border-white/20">
          Initiate Deep Audit
        </button>
      </div>
    </aside>
  </div>
</template>

<script setup>
import { 
  LucideUsers, LucideSearch, LucidePlus, LucideChevronLeft, 
  LucideChevronRight, LucideDownload, LucideZap 
} from 'lucide-vue-next'

const directory = [
  { name: 'Andi Maulana', dept: 'DEPT. FINANCE', id: 'ID-8821', initials: 'AM', bg: 'bg-[#1E3A5F]', active: true },
  { name: 'Siti Pertiwi', dept: 'DEPT. HRD', id: 'ID-9042', initials: 'SP', bg: 'bg-slate-400' },
  { name: 'Bambang Kus...', dept: 'DEPT. IT', id: 'ID-1123', initials: 'BK', bg: 'bg-slate-400' },
  { name: 'Dewi Wijaya', dept: 'DEPT. LEGAL', id: 'ID-7762', initials: 'DW', bg: 'bg-slate-400' }
]

const timeline = [
  {
    label: 'Policy Violation Flag',
    labelColor: 'text-red-500',
    time: 'Today, 14:22',
    dot: 'bg-red-500',
    bg: 'bg-red-50/30',
    title: 'Massive Document Download Triggered',
    desc: 'System detected 142 PDF exports within 3 minutes from <span class="text-blue-500 underline font-black">SEC_FIN_2024</span>. Exceeds individual velocity threshold.',
    actions: [
      { text: 'Lock Account', color: 'text-red-500' },
      { text: 'Ignore Once', color: 'text-slate-400' }
    ]
  },
  {
    label: 'Sensitive Access',
    labelColor: 'text-[#1E3A5F]',
    time: 'Yesterday, 09:45',
    dot: 'bg-[#1E3A5F]',
    title: 'Accessed "Q1_Tax_Audit_Internal_Draft"',
    desc: 'Location: Head Office (192.168.1.42). Verified via MFA-Token #921.'
  },
  {
    label: 'Approval Action',
    labelColor: 'text-green-500',
    time: 'May 14, 16:10',
    dot: 'bg-green-500',
    title: 'Selesai: Payroll Batch Approval',
    desc: 'Successfully authorized release for 250 records. Verification: Signature PKI-Valid.'
  },
  {
    label: 'Overdue Behavior Impact',
    labelColor: 'text-amber-700',
    time: 'May 13, 11:30',
    dot: 'bg-amber-700',
    title: 'Recurring Late Submission: Document #F-22',
    desc: 'Archiving task has been overdue for 72 hours. Impacting departmental compliance score.'
  }
]

const metrics = [
  { label: 'Total Accesses', value: '1,482', badge: '↑ 12%', badgeClass: 'bg-green-50 text-green-500', progress: 60, progressColor: 'bg-[#1E3A5F]' },
  { label: 'Sensitive Doc Attempts', value: '28', badge: '⚠ High', badgeClass: 'bg-red-50 text-red-500' },
  { label: 'Approval Actions', value: '156', badge: 'Consistent', badgeClass: 'bg-green-50 text-green-500' },
  { label: 'Overdue Behavior Impact', value: '8.2', subValue: '/ 10', progress: 82, progressColor: 'bg-amber-500' },
  { label: 'Policy Violation Flags', value: '3', link: 'View Audit Trail', labelColor: 'text-red-500' }
]

definePageMeta({
  layout: 'none'
})
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  @apply bg-transparent;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  @apply bg-slate-200 dark:bg-slate-800 rounded-full;
}

.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/5 shadow-sm;
}
</style>
