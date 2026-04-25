<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-1">
        <h1 class="text-4xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
          {{ $t('admin.integration.title') || 'Integration Status Monitor' }}
        </h1>
        <p class="text-slate-500 font-bold">
          {{ $t('admin.integration.subtitle') || 'Real-time health telemetry for key document management services.' }}
        </p>
      </div>
      <div class="flex items-center gap-4">
        <button class="px-6 py-4 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-slate-600 dark:text-slate-300 rounded-2xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center gap-2 shadow-sm">
          <LucideDownload class="w-4 h-4" />
          {{ $t('admin.integration.btn_download') || 'Download Integration Report' }}
        </button>
        <button class="px-10 py-4 bg-[#1E3A5F] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 active:scale-95">
          <LucideRefreshCcw class="w-4 h-4" />
          {{ $t('admin.integration.btn_refresh') || 'Refresh All Services' }}
        </button>
      </div>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
      <div v-for="(stat, i) in stats" :key="i" 
        class="glass p-10 rounded-[2.5rem] relative overflow-hidden group border-l-8"
        :class="stat.borderClass"
        v-motion-slide-visible-bottom
        :delay="i * 100"
      >
        <div class="flex items-start justify-between relative z-10">
          <div class="space-y-1">
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ stat.title }}</p>
            <p class="text-4xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ stat.value }}</p>
          </div>
          <div :class="`w-12 h-12 rounded-2xl flex items-center justify-center ${stat.bg} ${stat.color} shadow-sm border ${stat.borderColor}`">
            <component :is="stat.icon" class="w-6 h-6" />
          </div>
        </div>
        <div class="mt-4 flex items-center gap-2 relative z-10">
          <span :class="`text-[9px] font-black uppercase tracking-widest ${stat.subColor}`">{{ stat.subtitle }}</span>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Main Table -->
      <div class="lg:col-span-8 space-y-8">
        <div class="glass rounded-[3rem] overflow-hidden shadow-sm border border-slate-50 dark:border-slate-800" v-motion-slide-visible-bottom>
          <div class="px-10 py-8 border-b border-slate-50 dark:border-slate-800 flex items-center justify-between bg-slate-50/50 dark:bg-slate-900/50">
            <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('admin.integration.table.title') || 'Integration Service Table' }}</h3>
            <div class="flex items-center gap-2">
              <span class="w-2.5 h-2.5 rounded-full bg-green-500 animate-pulse"></span>
              <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.integration.table.live') || 'Live Telemetry Active' }}</span>
            </div>
          </div>
          <div class="overflow-x-auto">
            <table class="w-full text-left">
              <thead>
                <tr class="text-[9px] font-black text-slate-400 uppercase tracking-[0.2em] border-b border-slate-50 dark:border-slate-800">
                  <th class="p-8 pl-10">{{ $t('admin.integration.table.cols.name') || 'Service Name' }}</th>
                  <th class="p-8">{{ $t('admin.integration.table.cols.host') || 'Host Address' }}</th>
                  <th class="p-8 text-center">{{ $t('admin.integration.table.cols.status') || 'Status' }}</th>
                  <th class="p-8 pr-10">{{ $t('admin.integration.table.cols.latency') || 'Latency' }}</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
                <tr v-for="s in services" :key="s.name" class="group hover:bg-slate-50/50 dark:hover:bg-slate-800/30 transition-all cursor-pointer">
                  <td class="p-8 pl-10">
                    <div class="flex items-center gap-5">
                      <div class="w-12 h-12 rounded-2xl bg-slate-50 dark:bg-slate-800 flex items-center justify-center text-[#1E3A5F] dark:text-slate-300 border border-slate-100 dark:border-slate-700 shadow-sm group-hover:scale-110 transition-transform">
                        <component :is="s.icon" class="w-5 h-5" />
                      </div>
                      <div>
                        <p class="text-sm font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">{{ s.name }}</p>
                        <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ s.type }}</p>
                      </div>
                    </div>
                  </td>
                  <td class="p-8 font-mono text-[11px] text-slate-500">{{ s.host }}</td>
                  <td class="p-8">
                    <div class="flex justify-center">
                      <span :class="`px-4 py-1.5 rounded-full text-[9px] font-black uppercase tracking-widest border ${s.statusColor}`">
                        {{ s.status }}
                      </span>
                    </div>
                  </td>
                  <td class="p-8 pr-10">
                    <div class="flex items-center gap-4">
                      <div class="flex gap-0.5 items-end h-6">
                        <div v-for="(v, idx) in s.trend" :key="idx" 
                          :style="{ height: v + '%' }" 
                          :class="`w-1 rounded-t-sm transition-all ${idx === s.trend.length -1 ? s.trendColor : 'bg-slate-200 dark:bg-slate-700'}`">
                        </div>
                      </div>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- Right Sidebar -->
      <div class="lg:col-span-4 space-y-10">
        <!-- Live Alert Feed -->
        <div class="glass p-10 rounded-[3rem] space-y-8" v-motion-slide-visible-bottom>
          <div class="flex items-center justify-between border-b border-slate-50 dark:border-slate-800 pb-6">
            <div class="flex items-center gap-3">
              <LucideRadio class="w-4 h-4 text-red-500 animate-pulse" />
              <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('admin.integration.alerts.title') || 'Live Alert Feed' }}</h3>
            </div>
            <span class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">T+ 0.5s</span>
          </div>

          <div class="space-y-6">
            <div v-for="alert in alerts" :key="alert.id" 
              class="p-6 rounded-2xl border-l-4 space-y-2 relative group hover:scale-[1.02] transition-all cursor-pointer"
              :class="alert.bgClass"
            >
              <div class="flex items-center gap-3">
                <component :is="alert.icon" :class="`w-4 h-4 ${alert.iconColor}`" />
                <h4 :class="`text-[11px] font-black uppercase tracking-widest ${alert.titleColor}`">{{ alert.title }}</h4>
              </div>
              <p class="text-[11px] font-bold text-slate-500 leading-relaxed">{{ alert.desc }}</p>
              <div class="flex items-center justify-between pt-1">
                <span class="text-[9px] font-bold text-slate-400">{{ alert.time }}</span>
                <span class="text-[8px] font-black text-slate-300 uppercase tracking-tighter font-mono">{{ alert.code }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Latency Trend -->
        <div class="glass p-10 rounded-[3rem] space-y-8" v-motion-slide-visible-bottom>
          <div class="space-y-1">
            <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.integration.trend.title') || 'Latency Trend (60M)' }}</h3>
          </div>
          <div class="h-48 flex items-end justify-between gap-1.5">
            <div v-for="(v, i) in latencyTrend" :key="i" 
              :style="{ height: v + '%' }" 
              :class="`flex-grow rounded-sm transition-all group relative ${i === 3 ? 'bg-blue-400' : 'bg-[#1E3A5F] dark:bg-blue-900/40'}`"
            >
              <!-- Tooltip -->
              <div class="absolute bottom-full left-1/2 -translate-x-1/2 mb-2 px-2 py-1 bg-slate-900 text-[8px] text-white rounded opacity-0 group-hover:opacity-100 transition-opacity pointer-events-none whitespace-nowrap z-20">
                {{ v }}ms
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideDownload, LucideRefreshCcw, LucideCheckCircle2, LucideActivity, 
  LucideAlertTriangle, LucideRepeat, LucideDatabase, LucideHardDrive, 
  LucideScan, LucidePrinter, LucideCpu, LucideMail, LucideNetwork, 
  LucideRadio, LucideXCircle, LucideClock, LucideInfo
} from 'lucide-vue-next'

const stats = [
  { 
    title: 'OVERALL HEALTH', 
    value: '94.2%', 
    subtitle: 'SELESAI • 7/7 SERVICES ONLINE', 
    icon: LucideCheckCircle2, 
    bg: 'bg-green-50 dark:bg-green-900/20', 
    color: 'text-green-500', 
    borderColor: 'border-green-100 dark:border-green-800',
    borderClass: 'border-green-500',
    subColor: 'text-green-500'
  },
  { 
    title: 'AVG LATENCY', 
    value: '12ms', 
    subtitle: 'GLOBAL AVERAGE', 
    icon: LucideActivity, 
    bg: 'bg-blue-50 dark:bg-blue-900/20', 
    color: 'text-blue-500', 
    borderColor: 'border-blue-100 dark:border-blue-800',
    borderClass: 'border-blue-500',
    subColor: 'text-slate-400'
  },
  { 
    title: 'ACTIVE ALERTS', 
    value: '02', 
    subtitle: 'GAGAL • DATABASE WARNING', 
    icon: LucideAlertTriangle, 
    bg: 'bg-red-50 dark:bg-red-900/20', 
    color: 'text-red-500', 
    borderColor: 'border-red-100 dark:border-red-800',
    borderClass: 'border-red-500',
    subColor: 'text-red-500'
  },
  { 
    title: 'SYNC OPERATIONS', 
    value: '1.2k', 
    subtitle: 'LAST 60 MINUTES', 
    icon: LucideRepeat, 
    bg: 'bg-purple-50 dark:bg-purple-900/20', 
    color: 'text-purple-500', 
    borderColor: 'border-purple-100 dark:border-purple-800',
    borderClass: 'border-purple-500',
    subColor: 'text-slate-400'
  }
]

const services = [
  { 
    name: 'AD/LDAP Directory', 
    type: 'Directory', 
    host: '10.0.1.45:389', 
    status: 'SELESAI', 
    statusColor: 'bg-green-50 text-green-500 border-green-100', 
    icon: LucideNetwork,
    trend: [40, 45, 42, 38, 41, 40],
    trendColor: 'bg-green-500'
  },
  { 
    name: 'SQL Database', 
    type: 'Database', 
    host: 'db-prod-cluster.internal', 
    status: 'LATENCY SPIKE', 
    statusColor: 'bg-amber-50 text-amber-500 border-amber-100', 
    icon: LucideDatabase,
    trend: [30, 40, 35, 90, 85, 95],
    trendColor: 'bg-red-500'
  },
  { 
    name: 'S3/File Storage', 
    type: 'Object Storage', 
    host: 's3.ap-southeast-1.aws', 
    status: 'SELESAI', 
    statusColor: 'bg-green-50 text-green-500 border-green-100', 
    icon: LucideHardDrive,
    trend: [20, 22, 21, 20, 23, 21],
    trendColor: 'bg-green-500'
  },
  { 
    name: 'Network Scanner', 
    type: 'Hardware', 
    host: '192.168.20.104', 
    status: 'SELESAI', 
    statusColor: 'bg-green-50 text-green-500 border-green-100', 
    icon: LucideScan,
    trend: [10, 15, 12, 11, 14, 13],
    trendColor: 'bg-green-500'
  },
  { 
    name: 'Zebra Label Printer', 
    type: 'Hardware', 
    host: '192.168.20.115', 
    status: 'GAGAL', 
    statusColor: 'bg-red-50 text-red-500 border-red-100', 
    icon: LucidePrinter,
    trend: [0, 0, 0, 0, 0, 0],
    trendColor: 'bg-slate-300'
  },
  { 
    name: 'RFID Encoder', 
    type: 'Hardware', 
    host: 'localhost:8088', 
    status: 'SELESAI', 
    statusColor: 'bg-green-50 text-green-500 border-green-100', 
    icon: LucideCpu,
    trend: [5, 8, 7, 6, 9, 8],
    trendColor: 'bg-green-500'
  },
  { 
    name: 'Email SMTP Relay', 
    type: 'Mail Service', 
    host: 'smtp.office365.com:587', 
    status: 'SELESAI', 
    statusColor: 'bg-green-50 text-green-500 border-green-100', 
    icon: LucideMail,
    trend: [100, 95, 98, 92, 96, 94],
    trendColor: 'bg-green-500'
  }
]

const alerts = [
  {
    id: 1,
    title: 'SERVICE OFFLINE',
    desc: 'Zebra Label Printer has lost connection to terminal 104.',
    time: '14:23:45',
    code: 'ERR_CONN_TIMEOUT',
    icon: LucideXCircle,
    iconColor: 'text-red-500',
    titleColor: 'text-red-600',
    bgClass: 'bg-red-50/50 border-red-200'
  },
  {
    id: 2,
    title: 'LATENCY THRESHOLD EXCEEDED',
    desc: 'SQL Database cluster reporting elevated response times (142ms).',
    time: '14:18:22',
    code: 'WRN_HIGH_LATENCY',
    icon: LucideClock,
    iconColor: 'text-amber-500',
    titleColor: 'text-amber-600',
    bgClass: 'bg-amber-50/50 border-amber-200'
  },
  {
    id: 3,
    title: 'HEARTBEAT RESTORED',
    desc: 'Identity AD/LDAP synchronization successfully validated.',
    time: '13:55:01',
    code: 'INF_HEARTBEAT_OK',
    icon: LucideInfo,
    iconColor: 'text-blue-500',
    titleColor: 'text-blue-600',
    bgClass: 'bg-blue-50/50 border-blue-200'
  }
]

const latencyTrend = [40, 35, 45, 80, 25, 30, 35, 42, 45, 38, 30, 20]

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}
</style>
