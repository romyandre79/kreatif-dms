<template>
  <div class="max-w-full mx-auto space-y-10 pb-32" v-motion-fade>
    <!-- Header -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-1">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
          Executive KPI Dashboard <span class="text-slate-400">(F-45)</span>
        </h1>
        <p class="text-sm font-bold text-slate-500">
          Real-time enterprise data governance monitoring
        </p>
      </div>
      <div class="flex items-center gap-4">
        <div class="flex bg-slate-100 dark:bg-slate-900 rounded-2xl p-1 gap-1">
          <button class="px-6 py-3 bg-[#1E3A5F] text-white rounded-xl text-[10px] font-black uppercase tracking-widest shadow-lg">Live View</button>
          <div class="relative group">
            <button class="px-6 py-3 text-slate-400 text-[10px] font-black uppercase tracking-widest hover:text-[#1E3A5F] transition-colors flex items-center gap-2">
              Q3 2023
              <LucideChevronDown class="w-3 h-3" />
            </button>
          </div>
        </div>
        <button class="px-6 py-3 text-slate-400 hover:text-[#1E3A5F] text-[10px] font-black uppercase tracking-widest transition-colors">Export</button>
      </div>
    </div>

    <!-- Summary Cards -->
    <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-5 gap-6">
      <div v-for="(stat, i) in topStats" :key="i" 
        class="glass p-8 rounded-lg space-y-6 relative overflow-hidden group"
        v-motion-slide-visible-bottom
        :delay="i * 100"
      >
        <div class="flex items-center justify-between">
          <div :class="`w-10 h-10 rounded-xl flex items-center justify-center ${stat.bg} ${stat.color} shadow-sm`">
            <component :is="stat.icon" class="w-5 h-5" />
          </div>
          <span v-if="stat.badge" :class="`px-2 py-0.5 rounded-lg text-[8px] font-black uppercase tracking-widest ${stat.badgeClass}`">
            {{ stat.badge }}
          </span>
        </div>
        <div class="space-y-1">
          <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ stat.title }}</p>
          <p class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ stat.value }}</p>
        </div>
      </div>
    </div>

    <!-- Main Grid -->
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Monthly Growth Chart -->
      <div class="lg:col-span-8">
        <div class="glass p-10 rounded-lg space-y-10" v-motion-slide-visible-bottom>
          <div class="flex items-center justify-between px-2">
            <div class="space-y-1">
              <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">Monthly Growth by Department</h3>
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-tighter">Comparative analysis across core business units</p>
            </div>
            <div class="flex items-center gap-6">
              <div class="flex items-center gap-2">
                <div class="w-3 h-3 rounded-full bg-[#1E3A5F]"></div>
                <span class="text-[9px] font-black text-[#1E3A5F] dark:text-slate-400 uppercase tracking-widest">Operations</span>
              </div>
              <div class="flex items-center gap-2">
                <div class="w-3 h-3 rounded-full bg-blue-200"></div>
                <span class="text-[9px] font-black text-blue-300 uppercase tracking-widest">Legal</span>
              </div>
            </div>
          </div>

          <!-- Mock Bar Chart -->
          <div class="h-80 flex items-end justify-between gap-6 px-4">
            <div v-for="m in months" :key="m.name" class="flex-grow flex flex-col items-center gap-4 group">
              <div class="w-full flex flex-col-reverse gap-0.5 rounded-t-lg overflow-hidden transition-all group-hover:opacity-80">
                <div :style="{ height: (m.ops * 2.5) + 'px' }" class="bg-[#1E3A5F] w-full"></div>
                <div :style="{ height: (m.legal * 2.5) + 'px' }" class="bg-blue-100 dark:bg-blue-900/40 w-full"></div>
              </div>
              <span class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ m.name }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Sidebar Info -->
      <div class="lg:col-span-4 space-y-10">
        <!-- Access Activity -->
        <div class="glass p-10 rounded-lg space-y-10" v-motion-slide-visible-bottom :delay="100">
          <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em] border-b border-slate-50 dark:border-slate-800 pb-6">Access Activity by Role</h3>
          <div class="space-y-8">
            <div v-for="role in roles" :key="role.name" class="space-y-3">
              <div class="flex justify-between items-center">
                <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ role.name }}</span>
                <span class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase">{{ role.value }}%</span>
              </div>
              <div class="h-2 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
                <div :style="{ width: role.value + '%' }" class="h-full bg-[#1E3A5F] rounded-full"></div>
              </div>
            </div>
          </div>
        </div>

        <!-- SLA Performance -->
        <div class="glass p-10 rounded-lg space-y-8" v-motion-slide-visible-bottom :delay="200">
          <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">Approval SLA Performance</h3>
          <div class="flex items-center justify-between">
            <div class="space-y-1">
              <p class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">4.2h</p>
              <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">Average Response</p>
            </div>
            <div class="w-16 h-16 rounded-2xl border-4 border-blue-500/20 flex items-center justify-center">
              <span class="text-xs font-black text-[#1E3A5F] dark:text-blue-400">92%</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Risk Alerts Section -->
    <div class="space-y-6" v-motion-slide-visible-bottom>
      <div class="flex items-center justify-between px-2">
        <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest flex items-center gap-3">
          <LucideAlertTriangle class="w-5 h-5 text-red-500 animate-pulse" />
          Real-time Risk Alerts
        </h3>
        <button class="text-[10px] font-black text-blue-500 hover:text-blue-600 uppercase tracking-widest transition-colors">View All Protocols</button>
      </div>

      <div class="glass rounded-lg overflow-hidden shadow-sm border border-slate-50 dark:border-slate-800">
        <div class="overflow-x-auto">
          <table class="w-full text-left">
            <thead>
              <tr class="bg-slate-50/50 dark:bg-slate-900/50 text-[9px] font-black text-slate-400 uppercase tracking-[0.2em] border-b border-slate-50 dark:border-slate-800">
                <th class="p-8 pl-10">Severity</th>
                <th class="p-8">Risk Description</th>
                <th class="p-8">Department</th>
                <th class="p-8 text-center">Status</th>
                <th class="p-8">Timestamp</th>
                <th class="p-8 pr-10 text-right">Action</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
              <tr v-for="alert in alerts" :key="alert.id" class="group hover:bg-slate-50/50 dark:hover:bg-slate-800/30 transition-all cursor-pointer">
                <td class="p-8 pl-10">
                  <span :class="`px-3 py-1 rounded-lg text-[8px] font-black uppercase tracking-widest border ${alert.severityClass}`">
                    {{ alert.severity }}
                  </span>
                </td>
                <td class="p-8">
                  <div class="space-y-1">
                    <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ alert.title }}</p>
                    <p class="text-[10px] font-bold text-slate-400 uppercase tracking-tighter max-w-xs">{{ alert.desc }}</p>
                  </div>
                </td>
                <td class="p-8 text-[11px] font-black text-slate-500 uppercase tracking-tight">{{ alert.dept }}</td>
                <td class="p-8">
                  <div class="flex items-center justify-center gap-2">
                    <div :class="`w-2 h-2 rounded-full ${alert.statusColor}`"></div>
                    <span class="text-[10px] font-black text-slate-700 dark:text-slate-300 uppercase tracking-widest">{{ alert.status }}</span>
                  </div>
                </td>
                <td class="p-8 text-[11px] font-black text-slate-500 uppercase tracking-tight">{{ alert.time }}</td>
                <td class="p-8 pr-10 text-right">
                  <button class="px-6 py-2 bg-slate-50 dark:bg-slate-800 hover:bg-[#1E3A5F] hover:text-white border border-slate-200 dark:border-slate-700 text-[#1E3A5F] dark:text-slate-300 rounded-xl text-[9px] font-black uppercase tracking-widest transition-all">
                    {{ alert.btnText }}
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideHardDrive, LucideTrendingUp, LucideFolder, LucideAlertCircle, 
  LucideShieldCheck, LucideChevronDown, LucideDownload, 
  LucideAlertTriangle, LucideHistory 
} from 'lucide-vue-next'

const topStats = [
  { title: 'Total Active Archive', value: '1.24 PB', icon: LucideHardDrive, bg: 'bg-blue-50', color: 'text-blue-500', badge: '+4.2%', badgeClass: 'bg-green-50 text-green-500' },
  { title: 'Monthly Growth', value: '+14.8%', icon: LucideTrendingUp, bg: 'bg-amber-50', color: 'text-amber-500', badge: '↑ 12k', badgeClass: 'bg-green-50 text-green-500' },
  { title: 'Active Loans', value: '3,492', icon: LucideFolder, bg: 'bg-indigo-50', color: 'text-indigo-500' },
  { title: 'Overdue Count', value: '128', icon: LucideAlertCircle, bg: 'bg-red-50', color: 'text-red-500', badge: 'High', badgeClass: 'bg-red-100 text-red-600' },
  { title: 'Compliance Score', value: '98.4%', icon: LucideShieldCheck, bg: 'bg-blue-50', color: 'text-blue-600' }
]

const months = [
  { name: 'Jan', ops: 60, legal: 20 },
  { name: 'Feb', ops: 70, legal: 25 },
  { name: 'Mar', ops: 85, legal: 30 },
  { name: 'Apr', ops: 95, legal: 35 },
  { name: 'May', ops: 80, legal: 25 },
  { name: 'Jun', ops: 110, legal: 40 }
]

const roles = [
  { name: 'Administrator', value: 42 },
  { name: 'Auditor', value: 28 },
  { name: 'Department Head', value: 18 },
  { name: 'Staff', value: 12 }
]

const alerts = [
  {
    id: 1,
    severity: 'Critical',
    severityClass: 'bg-red-50 text-red-500 border-red-100',
    title: 'Unusual Bulk Export Activity',
    desc: 'User ID: AD-9902 detected 450+ downloads in 5min',
    dept: 'Finance & Treasury',
    status: 'Investigating',
    statusColor: 'bg-red-500',
    time: '14:22:10 WIB',
    btnText: 'Action'
  },
  {
    id: 2,
    severity: 'Medium',
    severityClass: 'bg-amber-50 text-amber-500 border-amber-100',
    title: 'Expired Retention Policy',
    desc: 'Batch #44910 reached mandatory destruction date',
    dept: 'Legal / Litigation',
    status: 'Queued',
    statusColor: 'bg-slate-400',
    time: '12:05:44 WIB',
    btnText: 'Review'
  },
  {
    id: 3,
    severity: 'Low',
    severityClass: 'bg-blue-50 text-blue-500 border-blue-100',
    title: 'LDAP Sync Delayed',
    desc: 'System synchronization took >300ms latency',
    dept: 'Infrastructure',
    status: 'Resolved',
    statusColor: 'bg-slate-900',
    time: '10:15:22 WIB',
    btnText: 'Log'
  }
]

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
@reference "tailwindcss";
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/5 shadow-sm;
}
</style>

