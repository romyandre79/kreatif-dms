<template>
  <div class="space-y-8">
    <PageHeader 
      title="Integration Status Monitor" 
      description="Real-time health telemetry for key document management services."
    >
      <template #actions>
        <div class="flex items-center gap-3">
          <button @click="downloadReport" class="flex items-center gap-2 px-4 py-2.5 bg-white/5 hover:bg-white/10 text-slate-300 rounded-xl border border-white/10 transition-all font-bold text-sm">
            <LucideDownload class="w-4 h-4" />
            Download Integration Report
          </button>
          <button @click="refreshAll" class="flex items-center gap-2 px-4 py-2.5 bg-primary-600 hover:bg-primary-500 text-white rounded-xl shadow-lg shadow-primary-500/20 transition-all font-bold text-sm">
            <LucideRefreshCw class="w-4 h-4" :class="{ 'animate-spin': refreshing }" />
            Refresh All Services
          </button>
        </div>
      </template>
    </PageHeader>

    <!-- Top Stats Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-4 gap-6">
      <div v-for="stat in topStats" :key="stat.label" class="relative group">
        <div :class="`absolute inset-x-0 -bottom-px h-1 transition-all duration-300 ${stat.color}`"></div>
        <div class="bg-white/80 dark:bg-slate-900/50 backdrop-blur-xl border border-slate-200 dark:border-slate-800 p-6 rounded-2xl">
          <div class="flex items-center justify-between mb-4">
            <span class="text-xs font-black text-slate-400 uppercase tracking-widest">{{ stat.label }}</span>
            <component :is="stat.icon" :class="`w-5 h-5 ${stat.iconColor}`" />
          </div>
          <div class="flex items-end gap-2">
            <span class="text-3xl font-black text-slate-900 dark:text-white tracking-tight">{{ stat.value }}</span>
            <span v-if="stat.suffix" class="text-sm font-bold text-slate-400 mb-1.5">{{ stat.suffix }}</span>
          </div>
          <div class="mt-4 flex items-center gap-2">
            <div :class="`w-2 h-2 rounded-full ${stat.statusColor} animate-pulse`"></div>
            <span :class="`text-xs font-bold uppercase tracking-tighter ${stat.statusTextColor}`">{{ stat.statusText }}</span>
          </div>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
      <!-- Main Integration Table -->
      <div class="lg:col-span-8 bg-white/80 dark:bg-slate-900/50 backdrop-blur-xl border border-slate-200 dark:border-slate-800 rounded-3xl overflow-hidden flex flex-col">
        <div class="p-6 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between bg-slate-50/50 dark:bg-slate-950/20">
          <div class="flex items-center gap-3">
            <h3 class="text-lg font-black text-slate-900 dark:text-white uppercase tracking-tight">Integration Service Table</h3>
            <div class="flex items-center gap-2 px-2.5 py-1 bg-green-500/10 rounded-full border border-green-500/20">
              <div class="w-1.5 h-1.5 rounded-full bg-green-500 animate-pulse"></div>
              <span class="text-[10px] font-black text-green-500 uppercase tracking-widest">Live Telemetry Active</span>
            </div>
          </div>
        </div>
        
        <div class="overflow-x-auto">
          <table class="w-full text-left border-collapse">
            <thead>
              <tr class="text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">
                <th class="px-8 py-4">Service Name</th>
                <th class="px-6 py-4">Host Address</th>
                <th class="px-6 py-4">Status</th>
                <th class="px-6 py-4">Latency</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50 dark:divide-slate-800/50">
              <tr v-for="service in services" :key="service.name" class="group hover:bg-slate-50 dark:hover:bg-white/5 transition-colors">
                <td class="px-8 py-5">
                  <div class="flex items-center gap-4">
                    <div :class="`w-10 h-10 rounded-xl flex items-center justify-center border transition-all duration-300 ${service.status === 'online' ? 'bg-white dark:bg-slate-800 border-slate-100 dark:border-slate-700 shadow-sm group-hover:scale-110' : 'bg-red-500/5 border-red-500/20'}`">
                      <component :is="service.icon" :class="`w-5 h-5 ${service.status === 'online' ? 'text-primary-500' : 'text-red-500'}`" />
                    </div>
                    <div>
                      <p class="text-sm font-black text-slate-900 dark:text-white leading-tight">{{ service.name }}</p>
                      <p class="text-[10px] font-bold text-slate-400 uppercase tracking-tighter mt-0.5">{{ service.category }}</p>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-5">
                  <span class="text-xs font-mono font-medium text-slate-500 dark:text-slate-400">{{ service.host }}</span>
                </td>
                <td class="px-6 py-5">
                  <div v-if="service.status === 'online'" class="inline-flex items-center px-3 py-1 bg-green-500/10 text-green-600 dark:text-green-400 rounded-full border border-green-500/20 text-[10px] font-black uppercase tracking-widest">
                    Selesai
                  </div>
                  <div v-else-if="service.status === 'warning'" class="inline-flex flex-col items-center px-3 py-1 bg-amber-500/10 text-amber-600 dark:text-amber-400 rounded-lg border border-amber-500/20 text-[10px] font-black uppercase tracking-widest leading-tight text-center">
                    Latency <span>Spike</span>
                  </div>
                  <div v-else class="inline-flex items-center px-3 py-1 bg-red-500/10 text-red-600 dark:text-red-400 rounded-full border border-red-500/20 text-[10px] font-black uppercase tracking-widest">
                    Gagal
                  </div>
                </td>
                <td class="px-6 py-5">
                  <div class="flex items-center gap-3">
                    <div class="flex flex-col">
                      <span class="text-[10px] font-black text-slate-900 dark:text-white">{{ service.lastLatency }}ms</span>
                      <div class="w-16 h-4 flex items-end gap-0.5 mt-1">
                        <div v-for="(v, i) in service.latencyHistory" :key="i" 
                          class="w-1.5 rounded-full transition-all duration-500" 
                          :class="[v > 200 ? 'bg-red-500' : v > 100 ? 'bg-amber-500' : 'bg-green-500']"
                          :style="`height: ${Math.min(v/5, 100)}%`"
                        ></div>
                      </div>
                    </div>
                    <div :class="`w-1.5 h-1.5 rounded-full ${service.status === 'online' ? 'bg-green-500' : 'bg-red-500'}`"></div>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Right Column: Alert Feed & Trends -->
      <div class="lg:col-span-4 space-y-8">
        <!-- Live Alert Feed -->
        <div class="bg-white/80 dark:bg-slate-900/50 backdrop-blur-xl border border-slate-200 dark:border-slate-800 rounded-3xl p-6">
          <div class="flex items-center justify-between mb-6">
            <div class="flex items-center gap-2">
              <LucideActivity class="w-4 h-4 text-red-500 animate-pulse" />
              <h3 class="text-sm font-black text-slate-900 dark:text-white uppercase tracking-widest">Live Alert Feed</h3>
            </div>
            <span class="text-[10px] font-mono font-bold text-slate-400">T+ 0.5s</span>
          </div>
          
          <div class="space-y-4">
            <div v-for="alert in alerts" :key="alert.id" 
              class="relative pl-6 py-4 pr-4 rounded-2xl overflow-hidden transition-all hover:scale-[1.02] cursor-default"
              :class="alert.bg"
            >
              <div :class="`absolute left-0 top-0 bottom-0 w-1.5 ${alert.accent}`"></div>
              <div class="flex gap-4">
                <LucideAlertTriangle v-if="alert.type === 'error'" class="w-5 h-5 text-red-500 flex-shrink-0 mt-0.5" />
                <LucideTimer v-else-if="alert.type === 'warning'" class="w-5 h-5 text-amber-500 flex-shrink-0 mt-0.5" />
                <LucideCheckCircle2 v-else class="w-5 h-5 text-blue-500 flex-shrink-0 mt-0.5" />
                
                <div class="flex-1 min-w-0">
                  <h4 class="text-xs font-black uppercase tracking-widest mb-1" :class="alert.titleColor">{{ alert.title }}</h4>
                  <p class="text-xs font-medium text-slate-500 dark:text-slate-400 leading-relaxed">{{ alert.message }}</p>
                  <div class="mt-3 flex items-center justify-between">
                    <span class="text-[10px] font-bold text-slate-400 tracking-widest">{{ alert.time }}</span>
                    <span class="text-[10px] font-mono font-bold text-slate-300 dark:text-slate-600 tracking-tighter">{{ alert.code }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="bg-white/80 dark:bg-slate-900/50 backdrop-blur-xl border border-slate-200 dark:border-slate-800 rounded-3xl p-6">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-6">Latency Trend (Recent Checks)</h3>
          <div class="h-40 flex items-end justify-between gap-1 px-2">
            <div v-for="(v, i) in globalTrend" :key="i" 
              class="flex-1 rounded-t-sm transition-all duration-700 hover:opacity-80" 
              :class="[v > 70 ? 'bg-red-500/40 border-t-2 border-red-500' : i === globalTrend.length - 1 ? 'bg-primary-500 shadow-[0_0_15px_rgba(59,130,246,0.5)]' : 'bg-slate-700/40']"
              :style="`height: ${Math.min(v, 100)}%`"
            ></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>


<script setup>
import { ref, onMounted, onUnmounted, computed } from 'vue'
import { useRuntimeConfig } from '#app'
import { 
  LucideDownload, 
  LucideRefreshCw, 
  LucideActivity, 
  LucideServer, 
  LucideDatabase, 
  LucideHardDrive, 
  LucideScan, 
  LucidePrinter, 
  LucideHash, 
  LucideMail, 
  LucideAlertTriangle, 
  LucideCheckCircle2, 
  LucideTimer 
} from 'lucide-vue-next'

const refreshing = ref(false)
const nodes = ref([])
const { $api } = useApi()
const config = useRuntimeConfig()

const fetchNodes = async () => {
  try {
    refreshing.value = true
    const res = await $api(`${config.public.apiBase}/master/integration/status`)
    
    console.log('API Response:', res)
    if (res && res.data) {
      nodes.value = res.data
    }
  } catch (err) {
    console.error('Failed to fetch integration status:', err)
  } finally {
    refreshing.value = false
  }
}

const getIcon = (type) => {
  switch (type) {
    case 'LDAP': return LucideServer
    case 'DATABASE': return LucideDatabase
    case 'S3': return LucideHardDrive
    case 'SCANNER': return LucideScan
    case 'PRINTER': return LucidePrinter
    case 'SMTP': return LucideMail
    default: return LucideHash
  }
}

const services = computed(() => {
  return nodes.value.map(node => ({
    name: node.name,
    category: node.service_type,
    host: node.endpoint,
    status: node.status || 'offline',
    icon: getIcon(node.service_type),
    latencyHistory: node.latency_history || [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    lastLatency: node.last_latency || 0,
    lastError: node.last_error
  }))
})

const topStats = computed(() => {
  const onlineCount = nodes.value.filter(n => n.status === 'online').length
  const totalCount = nodes.value.length
  const avgLatency = nodes.value.reduce((acc, curr) => acc + (curr.last_latency || 0), 0) / (totalCount || 1)
  const criticalIssues = nodes.value.filter(n => n.status === 'offline' && n.is_critical).length

  return [
    { label: 'Overall Health', value: totalCount ? ((onlineCount/totalCount)*100).toFixed(1) : '0', suffix: '%', icon: LucideCheckCircle2, iconColor: 'text-green-500', color: 'bg-green-500', statusColor: 'bg-green-500', statusTextColor: 'text-green-500', statusText: `${onlineCount}/${totalCount} Services Online` },
    { label: 'Avg Latency', value: avgLatency.toFixed(0), suffix: 'ms', icon: LucideTimer, iconColor: 'text-blue-500', color: 'bg-blue-500', statusColor: 'bg-blue-500', statusTextColor: 'text-blue-500', statusText: 'Global Average' },
    { label: 'Critical Alerts', value: String(criticalIssues).padStart(2, '0'), suffix: '', icon: LucideAlertTriangle, iconColor: 'text-red-500', color: 'bg-red-500', statusColor: 'bg-red-500', statusTextColor: 'text-red-500', statusText: criticalIssues > 0 ? 'Action Required' : 'All Critical Services OK' },
    { label: 'Sync Operations', value: '1.2k', suffix: '', icon: LucideRefreshCw, iconColor: 'text-indigo-500', color: 'bg-indigo-500', statusColor: 'bg-indigo-500', statusTextColor: 'text-indigo-500', statusText: 'Last 60 Minutes' },
  ]
})

const alerts = computed(() => {
  const result = []
  nodes.value.forEach(node => {
    if (node.status === 'offline') {
      result.push({
        id: node.id,
        type: 'error',
        title: 'Service Offline',
        message: `${node.name} is currently unreachable at ${node.endpoint}.`,
        time: node.last_check_at ? new Date(node.last_check_at).toLocaleTimeString() : 'N/A',
        code: 'ERR_CONN_TIMEOUT',
        bg: 'bg-red-500/5 dark:bg-red-500/10',
        accent: 'bg-red-500',
        titleColor: 'text-red-600 dark:text-red-400'
      })
    } else if (node.last_latency > 100) {
      result.push({
        id: node.id,
        type: 'warning',
        title: 'High Latency',
        message: `${node.name} response time is elevated (${node.last_latency}ms).`,
        time: node.last_check_at ? new Date(node.last_check_at).toLocaleTimeString() : 'N/A',
        code: 'WRN_LATENCY_SPIKE',
        bg: 'bg-amber-500/5 dark:bg-amber-500/10',
        accent: 'bg-amber-500',
        titleColor: 'text-amber-600 dark:text-amber-400'
      })
    }
  })

  // Add a default info alert if everything is OK
  if (result.length === 0 && nodes.value.length > 0) {
    result.push({
      id: 'all-ok',
      type: 'info',
      title: 'System Healthy',
      message: 'All integration nodes are responding within normal parameters.',
      time: new Date().toLocaleTimeString(),
      code: 'INF_HEALTH_OK',
      bg: 'bg-blue-500/5 dark:bg-blue-500/10',
      accent: 'bg-blue-500',
      titleColor: 'text-blue-600 dark:text-blue-400'
    })
  }

  return result
})

const globalTrend = computed(() => {
  if (!nodes.value.length) return Array(15).fill(0)
  
  // Find the max length of history (up to 20)
  const maxLength = Math.max(...nodes.value.map(n => (n.latency_history || []).length), 0)
  if (maxLength === 0) return Array(15).fill(0)

  const trend = []
  for (let i = 0; i < maxLength; i++) {
    let sum = 0
    let count = 0
    nodes.value.forEach(node => {
      const history = node.latency_history || []
      // Align to the end of the array
      const offset = history.length - maxLength + i
      if (offset >= 0 && offset < history.length) {
        sum += history[offset]
        count++
      }
    })
    // Scale for percentage display (assuming 500ms is 100%)
    trend.push(count > 0 ? (sum / count) / 5 : 0)
  }
  return trend
})

const downloadReport = async () => {
  try {
    const res = await $api(`${config.public.apiBase}/master/integration/report`, {
      responseType: 'blob'
    })
    
    // Create a link to download the blob
    const url = window.URL.createObjectURL(new Blob([res]))
    const link = document.createElement('a')
    link.href = url
    link.setAttribute('download', `integration_report_${new Date().getTime()}.csv`)
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
  } catch (err) {
    console.error('Failed to download report:', err)
  }
}

const refreshAll = () => {
  fetchNodes()
}

onMounted(() => {
  fetchNodes()
  // Auto refresh every 30s to match backend
  const timer = setInterval(fetchNodes, 30000)
  onUnmounted(() => clearInterval(timer))
})
</script>

<style scoped>
/* Glassmorphism subtle polish */
.bg-white\/80 {
  backdrop-filter: blur(24px) saturate(180%);
}
</style>
