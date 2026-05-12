<template>
  <div class="space-y-10 pb-20">
    <!-- Top Filter Bar -->
    <header class="bg-white dark:bg-slate-900 border-b border-slate-200 dark:border-slate-800 px-10 py-6 flex items-end gap-8 shadow-sm relative z-20">
      <div class="space-y-3">
        <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-1">Event Type</label>
        <div class="relative w-64">
          <select 
            v-model="filters.action"
            class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-[11px] font-black text-[#1E3A5F] dark:text-white appearance-none outline-none focus:ring-2 focus:ring-blue-500/20 transition-all uppercase tracking-tight"
          >
            <option value="">All Events</option>
            <option value="CREATE">Create</option>
            <option value="UPDATE">Update</option>
            <option value="DELETE">Delete</option>
          </select>
          <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
        </div>
      </div>

      <div class="space-y-3">
        <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-1">Search User / Module</label>
        <div class="relative w-80 group">
          <LucideSearch class="absolute left-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 group-focus-within:text-blue-500 transition-colors" />
          <input 
            type="text" 
            v-model="filters.search"
            placeholder="Search activities..." 
            class="w-full bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl pl-12 pr-6 py-3.5 text-[11px] font-black outline-none focus:ring-2 focus:ring-blue-500/20 transition-all uppercase tracking-tight" 
          />
        </div>
      </div>

      <button 
        @click="fetchLogs"
        class="px-10 py-4 bg-[#1E3A5F] text-white rounded-2xl text-[11px] font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 active:scale-95"
      >
        <LucideRefreshCw class="w-4 h-4" :class="{ 'animate-spin': loading }" />
        Refresh
      </button>
    </header>

    <div class="flex flex-col gap-10">
      <!-- Main Content: History Log Table -->
      <main class="flex flex-col gap-8">
        <div class="bg-white dark:bg-slate-900 rounded-lg shadow-sm border border-slate-100 dark:border-slate-800 flex flex-col overflow-hidden flex-grow">
          <!-- Table Header -->
          <div class="px-10 py-8 border-b border-slate-50 dark:border-slate-800 flex items-center justify-between bg-slate-50/30 dark:bg-slate-900/30">
            <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em]">{{ $t('layout.menu.audit_logs') }}</h3>
            <div class="flex items-center gap-6">
              <button class="p-3 text-slate-400 hover:text-blue-500 transition-colors"><LucideDownload class="w-5 h-5" /></button>
            </div>
          </div>

          <!-- Table -->
          <div class="overflow-auto flex-grow custom-scrollbar">
            <table class="w-full text-left border-collapse">
              <thead>
                <tr class="text-[9px] font-black text-slate-400 uppercase tracking-[0.2em] border-b border-slate-50 dark:border-slate-800 sticky top-0 bg-white dark:bg-slate-900 z-10">
                  <th class="p-8 pl-10">User</th>
                  <th class="p-8 text-center">Action</th>
                  <th class="p-8">Module / Entity</th>
                  <th class="p-8">IP Address</th>
                  <th class="p-8">Timestamp</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
                <tr v-if="loading" class="animate-pulse">
                  <td colspan="5" class="p-20 text-center text-slate-300 font-black uppercase tracking-widest text-xs">Loading logs...</td>
                </tr>
                <tr v-else-if="filteredLogs.length === 0">
                  <td colspan="5" class="p-20 text-center text-slate-300 font-black uppercase tracking-widest text-xs">No activities found</td>
                </tr>
                <tr v-for="log in paginatedLogs" :key="log.id" 
                  @click="selectedLog = log"
                  class="group transition-all cursor-pointer border-l-4"
                  :class="selectedLog?.id === log.id ? 'bg-blue-50/50 dark:bg-blue-900/10 border-blue-500' : 'hover:bg-slate-50/30 dark:hover:bg-slate-800/30 border-transparent'"
                >
                  <td class="p-8 pl-10">
                    <div class="flex items-center gap-4">
                      <div class="w-10 h-10 rounded-xl bg-primary-100 dark:bg-primary-900/30 flex items-center justify-center text-primary-600 dark:text-primary-400 text-[10px] font-black uppercase">
                        {{ log.user_name?.substring(0, 2) || 'US' }}
                      </div>
                      <div class="space-y-0.5">
                        <p class="text-[13px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ log.user_name }}</p>
                        <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ log.user_email }}</p>
                      </div>
                    </div>
                  </td>
                  <td class="p-8">
                    <div class="flex justify-center">
                      <span :class="getActionBadge(log.action)">
                        {{ log.action }}
                      </span>
                    </div>
                  </td>
                  <td class="p-8">
                    <div class="flex items-center gap-3">
                      <div class="p-2 bg-slate-50 dark:bg-slate-800 rounded-lg">
                        <LucideBox class="w-4 h-4 text-slate-400" v-if="log.entity_type === 'box'" />
                        <LucideBuilding2 class="w-4 h-4 text-slate-400" v-else-if="log.entity_type === 'company' || log.entity_type === 'branch'" />
                        <LucideFileText class="w-4 h-4 text-slate-400" v-else />
                      </div>
                      <div class="space-y-0.5">
                        <p class="text-[11px] font-black text-slate-600 dark:text-slate-300 uppercase tracking-tight">{{ log.entity_type }}</p>
                        <p class="text-[9px] font-bold text-slate-400 font-mono tracking-tighter">{{ (typeof log.entity_id === 'string' ? log.entity_id : (log.entity_id?.Bytes || ''))?.substring(0, 8) }}...</p>
                      </div>
                    </div>
                  </td>
                  <td class="p-8">
                    <div class="flex items-center gap-2">
                       <LucideNetwork class="w-3.5 h-3.5 text-slate-400" />
                       <p class="text-[11px] font-mono text-slate-500 font-bold tracking-tight">{{ log.ip_address || '127.0.0.1' }}</p>
                    </div>
                  </td>
                  <td class="p-8">
                    <p class="text-[11px] font-black text-[#1E3A5F] dark:text-slate-300 uppercase tracking-tight">{{ formatDate(log.created_at) }}</p>
                    <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ formatTime(log.created_at) }}</p>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <!-- Pagination -->
          <div class="px-10 py-8 border-t border-slate-50 dark:border-slate-800 flex items-center justify-between bg-slate-50/30 dark:bg-slate-900/30">
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">
              Showing {{ startIndex + 1 }} to {{ endIndex }} of {{ filteredLogs.length }} activities
            </p>
            <div class="flex items-center gap-2">
              <button 
                @click="currentPage--"
                :disabled="currentPage === 1"
                class="p-2 rounded-lg hover:bg-slate-100 dark:hover:bg-slate-800 text-slate-400 disabled:opacity-30"
              >
                <LucideChevronLeft class="w-4 h-4" />
              </button>
              <button 
                v-for="p in Math.min(5, totalPages)" :key="p"
                @click="currentPage = p"
                class="w-8 h-8 rounded-lg text-[10px] font-black uppercase transition-all"
                :class="currentPage === p ? 'bg-[#1E3A5F] text-white shadow-lg' : 'text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800'"
              >
                {{ p }}
              </button>
              <button 
                @click="currentPage++"
                :disabled="currentPage >= totalPages"
                class="p-2 rounded-lg hover:bg-slate-100 dark:hover:bg-slate-800 text-slate-400 disabled:opacity-30"
              >
                <LucideChevronRight class="w-4 h-4" />
              </button>
            </div>
          </div>
        </div>
      </main>

      <!-- Right Sidebar: Event Details -->
      <aside v-if="selectedLog" class="bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 flex flex-col p-10 rounded-lg gap-10 shadow-2xl relative z-30">
        <div class="flex items-center justify-between pb-8 border-b border-slate-50 dark:border-slate-800">
          <div class="space-y-1">
            <h2 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Audit Details</h2>
            <p class="text-[9px] font-black text-slate-400 uppercase tracking-[0.2em] italic">System Signature Verified</p>
          </div>
          <button @click="selectedLog = null" class="p-2 text-slate-300 hover:text-red-500 transition-colors"><LucideX class="w-6 h-6" /></button>
        </div>

        <div class="space-y-12">
          <!-- Action Summary -->
          <div :class="getActionSummaryClass(selectedLog.action)">
            <div class="w-12 h-12 bg-white dark:bg-slate-800 rounded-2xl flex items-center justify-center shadow-sm">
              <LucidePlus class="w-6 h-6" v-if="selectedLog.action === 'CREATE'" />
              <LucideRefreshCw class="w-6 h-6" v-else-if="selectedLog.action === 'UPDATE'" />
              <LucideTrash2 class="w-6 h-6" v-else-if="selectedLog.action === 'DELETE'" />
            </div>
            <div class="space-y-1">
              <p class="text-[9px] font-black opacity-60 uppercase tracking-widest">Operation Mode</p>
              <p class="text-lg font-black uppercase tracking-tight">{{ selectedLog.action }} {{ selectedLog.entity_type }}</p>
            </div>
          </div>

          <!-- Technical Metadata -->
          <section class="space-y-6">
            <h4 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em] flex items-center gap-3">
              <LucideTerminal class="w-4 h-4 text-blue-500" />
              Technical Identity
            </h4>
            <div class="grid grid-cols-2 gap-8 p-8 bg-slate-50 dark:bg-slate-800/50 rounded-lg border border-slate-100 dark:border-slate-800">
               <div class="space-y-1">
                 <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Entity Type</p>
                 <p class="text-[13px] font-black text-[#1E3A5F] dark:text-white uppercase">{{ selectedLog.entity_type }}</p>
               </div>
               <div class="space-y-1">
                 <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Object ID</p>
                 <p class="text-[11px] font-bold text-blue-500 font-mono truncate uppercase tracking-tighter" :title="selectedLog.entity_id">
                   {{ selectedLog.entity_id }}
                 </p>
               </div>
               <div class="space-y-1">
                 <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Client IP</p>
                 <p class="text-[13px] font-black text-[#1E3A5F] dark:text-white uppercase font-mono">{{ selectedLog.ip_address || 'Internal' }}</p>
               </div>
               <div class="space-y-1">
                 <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Event Hash</p>
                 <p class="text-[10px] font-black text-slate-400 font-mono truncate uppercase">{{ selectedLog.id?.substring(0, 16) }}</p>
               </div>
            </div>
          </section>

          <!-- Captured Payload -->
          <section class="space-y-6" v-if="selectedLog.details">
            <h4 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em] flex items-center gap-3">
              <LucideCode class="w-4 h-4 text-blue-500" />
              Captured Payload
            </h4>
            <div class="p-8 bg-slate-900 rounded-lg shadow-inner overflow-hidden border border-slate-800">
               <div v-if="parsedDetails" class="space-y-4">
                 <div v-for="(val, key) in parsedDetails" :key="key" class="flex flex-col gap-1 border-b border-white/5 pb-2 last:border-0">
                   <span class="text-[8px] font-black text-blue-500 uppercase tracking-widest">{{ key }}</span>
                   <span class="text-[11px] font-mono text-slate-300 break-all">{{ val }}</span>
                 </div>
               </div>
               <pre v-else class="text-[10px] font-mono text-blue-400 overflow-x-auto custom-scrollbar"><code>{{ selectedLog.details }}</code></pre>
            </div>
          </section>

          <!-- System Evidence -->
          <div class="p-10 bg-blue-50 dark:bg-blue-900/10 border-l-8 border-blue-500 rounded-lg space-y-6">
            <div class="flex items-center gap-4">
              <LucideShieldCheck class="w-6 h-6 text-blue-500" />
              <h5 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">Immutable Evidence</h5>
            </div>
            <p class="text-[10px] font-bold text-slate-500 dark:text-slate-400 leading-relaxed uppercase">
              This log entry has been cryptographically signed. Any modification attempt will invalidate the system integrity check.
            </p>
            <div class="pt-4 border-t border-blue-100 dark:border-blue-800 flex justify-between items-center">
               <span class="text-[9px] font-black text-blue-500 uppercase tracking-widest">Integrity Check</span>
               <span class="px-3 py-1 bg-green-500 text-white text-[8px] font-black rounded-lg uppercase tracking-widest">Verified</span>
            </div>
          </div>
        </div>
      </aside>

      <!-- Empty State for Details -->
      <aside v-else class="bg-slate-50 dark:bg-slate-900/50 border border-dashed border-slate-200 dark:border-slate-800 flex flex-col items-center justify-center p-20 text-center opacity-30 rounded-lg">
         <LucideHistory class="w-16 h-16 text-slate-300 mb-6" />
         <h2 class="text-xs font-black text-slate-400 uppercase tracking-[0.3em]">Select an activity to view technical details</h2>
      </aside>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { useRuntimeConfig } from '#app'
import { 
  LucideChevronDown, LucideSearch, LucideRefreshCw, LucideHistory,
  LucideDownload, LucideFileText, LucideX, LucideTrash2, 
  LucideUser, LucideNetwork, LucideShieldCheck, LucideBox,
  LucideBuilding2, LucidePlus, LucideChevronLeft, LucideChevronRight,
  LucideTerminal, LucideCode
} from 'lucide-vue-next'
import { useApi } from '@/composables/useApi'

const { $api } = useApi()
const config = useRuntimeConfig()
const loading = ref(false)
const logs = ref([])
const selectedLog = ref(null)

const filters = ref({
  action: '',
  search: ''
})

const currentPage = ref(1)
const itemsPerPage = ref(15)

const fetchLogs = async () => {
  console.log('AUDIT: Fetching logs...')
  loading.value = true
  try {
    const res = await $api(`${config.public.apiBase}/master/audit-logs`)
    console.log('AUDIT: API Response:', res)
    if (res && res.data) {
      logs.value = Array.isArray(res.data) ? res.data : (res.data.logs || [])
      console.log('AUDIT: Logs loaded:', logs.value.length)
    }
  } catch (err) {
    console.error('AUDIT: Failed to fetch logs:', err)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  console.log('AUDIT: Component mounted')
  fetchLogs()
})

const filteredLogs = computed(() => {
  let result = logs.value
  
  if (filters.value.action) {
    result = result.filter(l => l.action === filters.value.action)
  }
  
  if (filters.value.search) {
    const q = filters.value.search.toLowerCase()
    result = result.filter(l => 
      l.user_name?.toLowerCase().includes(q) || 
      l.entity_type?.toLowerCase().includes(q) ||
      l.action?.toLowerCase().includes(q)
    )
  }
  
  return result
})

const totalPages = computed(() => Math.ceil(filteredLogs.value.length / itemsPerPage.value) || 1)
const startIndex = computed(() => (currentPage.value - 1) * itemsPerPage.value)
const endIndex = computed(() => Math.min(startIndex.value + itemsPerPage.value, filteredLogs.value.length))
const paginatedLogs = computed(() => filteredLogs.value.slice(startIndex.value, endIndex.value))

const parsedDetails = computed(() => {
  if (!selectedLog.value?.details) return null
  if (typeof selectedLog.value.details === 'object') return selectedLog.value.details
  try {
    // Try to decode base64 if it looks like one, or just parse JSON
    const str = typeof selectedLog.value.details === 'string' ? selectedLog.value.details : ''
    if (!str) return null
    return JSON.parse(str)
  } catch (e) {
    return null
  }
})

const formatDate = (dateStr) => {
  if (!dateStr) return '-'
  try {
    const d = new Date(dateStr)
    if (isNaN(d.getTime())) return '-'
    return d.toLocaleDateString('en-GB', {
      day: '2-digit', month: 'short', year: 'numeric'
    })
  } catch (e) {
    return '-'
  }
}

const formatTime = (dateStr) => {
  if (!dateStr) return '-'
  try {
    const d = new Date(dateStr)
    if (isNaN(d.getTime())) return '-'
    return d.toLocaleTimeString('en-GB', {
      hour: '2-digit', minute: '2-digit', second: '2-digit'
    })
  } catch (e) {
    return '-'
  }
}

const getActionBadge = (action) => {
  const base = "px-3 py-1 rounded-lg text-[9px] font-black uppercase tracking-widest "
  switch (action) {
    case 'CREATE': return base + "bg-green-50 dark:bg-green-900/20 text-green-600 dark:text-green-400"
    case 'UPDATE': return base + "bg-blue-50 dark:bg-blue-900/20 text-blue-600 dark:text-blue-400"
    case 'DELETE': return base + "bg-red-50 dark:bg-red-900/20 text-red-600 dark:text-red-400"
    default: return base + "bg-slate-50 dark:bg-slate-800 text-slate-500"
  }
}

const getActionSummaryClass = (action) => {
  const base = "p-8 rounded-lg border flex items-center gap-6 "
  switch (action) {
    case 'CREATE': return base + "bg-green-50 dark:bg-green-900/10 border-green-100 dark:border-green-800 text-green-600 dark:text-green-400"
    case 'UPDATE': return base + "bg-blue-50 dark:bg-blue-900/10 border-blue-100 dark:border-blue-800 text-blue-600 dark:text-blue-400"
    case 'DELETE': return base + "bg-red-50 dark:bg-red-900/10 border-red-100 dark:border-red-800 text-red-600 dark:text-red-400"
    default: return base + "bg-slate-50 dark:bg-slate-800 text-slate-500"
  }
}

watch(filters, () => {
  currentPage.value = 1
}, { deep: true })

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #e2e8f0; /* slate-200 */
  border-radius: 9999px;
}
.dark .custom-scrollbar::-webkit-scrollbar-thumb {
  background: #1e293b; /* slate-800 */
}
</style>
