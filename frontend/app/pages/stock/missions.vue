<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Page Header -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-1">
        <h1 class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
          {{ $t('stock.missions.title') }} 
          <span class="text-slate-400 font-bold normal-case text-lg">{{ $t('stock.missions.head_view') }}</span>
        </h1>
        <p class="text-xs font-bold text-slate-500">{{ $t('stock.missions.subtitle') }}</p>
      </div>
      
      <div class="flex flex-wrap items-center gap-4">
        <button @click="openCreateModal" class="px-6 py-3 bg-primary-500 text-white rounded-xl text-xs font-black uppercase tracking-widest hover:bg-primary-600 shadow-lg shadow-primary-500/20 transition-all flex items-center gap-2">
          <LucidePlus class="w-4 h-4" /> Create Audit Mission
        </button>
        
        <div class="flex items-center bg-white dark:bg-slate-900 p-1 rounded-xl shadow-sm border border-slate-100 dark:border-slate-800">
          <button v-for="t in ['all', 'scheduled', 'completed']" :key="t" :class="`px-6 py-2.5 rounded-lg text-xs font-black uppercase tracking-widest transition-all ${activeTab === t ? 'bg-[#1E3A5F] text-white shadow-lg' : 'text-slate-400 hover:text-slate-600'}`" @click="activeTab = t">
            {{ $t(`stock.missions.tabs.${t}`) }}
          </button>
        </div>
      </div>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
      <div class="glass p-8 rounded-lg flex items-center justify-between group hover:border-primary-500/30 transition-all">
        <div class="space-y-1">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('stock.missions.stats.pending') }}</p>
          <div class="flex items-baseline gap-3">
            <span class="text-3xl font-black text-[#1E3A5F] dark:text-white">{{ stats.pending_missions }}</span>
            <span class="text-[9px] font-black text-orange-500 uppercase tracking-widest">{{ $t('stock.missions.stats.approval_req') }}</span>
          </div>
        </div>
        <div class="w-12 h-12 rounded-2xl bg-orange-50 dark:bg-orange-900/20 text-orange-500 flex items-center justify-center"><LucideClipboardList class="w-6 h-6" /></div>
      </div>
      
      <div class="glass p-8 rounded-lg flex items-center justify-between group hover:border-primary-500/30 transition-all">
        <div class="space-y-1">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('stock.missions.stats.active') }}</p>
          <div class="flex items-baseline gap-3">
            <span class="text-3xl font-black text-[#1E3A5F] dark:text-white">{{ stats.active_missions }}</span>
            <span class="text-[9px] font-black text-blue-500 uppercase tracking-widest">{{ $t('stock.missions.stats.on_mission') }}</span>
          </div>
        </div>
        <div class="w-12 h-12 rounded-2xl bg-blue-50 dark:bg-blue-900/20 text-blue-500 flex items-center justify-center"><LucideUsers class="w-6 h-6" /></div>
      </div>

      <div class="glass p-8 rounded-lg flex items-center justify-between group hover:border-primary-500/30 transition-all">
        <div class="space-y-1">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('stock.missions.stats.rate') }}</p>
          <div class="flex items-baseline gap-3">
            <span class="text-3xl font-black text-[#1E3A5F] dark:text-white">{{ stats.accuracy_rate.toFixed(0) }}%</span>
            <span class="text-[9px] font-black text-green-500 uppercase tracking-widest">{{ $t('stock.missions.stats.vs_last_week', { value: '+2.4%' }) }}</span>
          </div>
        </div>
        <div class="w-12 h-12 rounded-2xl bg-green-50 dark:bg-green-900/20 text-green-500 flex items-center justify-center"><LucideCheckCircle2 class="w-6 h-6" /></div>
      </div>
    </div>

    <!-- Main Table Section -->
    <div class="glass rounded-lg overflow-hidden" v-motion-slide-visible-bottom>
      <div class="p-8 border-b border-slate-100 dark:border-slate-800 flex flex-col md:flex-row md:items-center justify-between gap-6 bg-white/30">
        <div class="relative flex-grow max-w-md">
          <LucideSearch class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
          <input v-model="searchQuery" type="text" :placeholder="$t('stock.missions.search_placeholder')" class="w-full pl-11 pr-5 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all" />
        </div>
        <div class="flex items-center gap-3">
          <button class="px-6 py-3 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-[#1E3A5F] dark:text-white rounded-xl text-xs font-black uppercase tracking-widest flex items-center gap-2 hover:bg-slate-50 transition-all shadow-sm">
            <LucideFilter class="w-4 h-4" /> Filter
          </button>
          <button class="px-6 py-3 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-[#1E3A5F] dark:text-white rounded-xl text-xs font-black uppercase tracking-widest flex items-center gap-2 hover:bg-slate-50 transition-all shadow-sm">
            <LucideDownload class="w-4 h-4" /> Export
          </button>
        </div>
      </div>

      <div class="overflow-x-auto">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800 bg-slate-50/30 dark:bg-slate-900/30">
              <th class="p-8 px-10">{{ $t('stock.missions.table.cols.target') }}</th>
              <th class="p-8">{{ $t('stock.missions.table.cols.schedule') }}</th>
              <th class="p-8">{{ $t('stock.missions.table.cols.operator') }}</th>
              <th class="p-8 text-center">{{ $t('stock.missions.table.cols.type') }}</th>
              <th class="p-8">{{ $t('stock.missions.table.cols.status') }}</th>
              <th class="p-8 text-center px-10">{{ $t('stock.missions.table.cols.actions') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50 dark:divide-slate-800/50">
            <tr v-for="mission in filteredMissions" :key="mission.id" class="hover:bg-slate-50/50 dark:hover:bg-slate-800/20 transition-all group">
              <td class="p-8 px-10">
                <div class="space-y-1">
                  <p class="text-sm font-black text-[#1E3A5F] dark:text-white tracking-tight">{{ mission.title }}</p>
                  <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ mission.target_area?.String || 'All Areas' }}</p>
                </div>
              </td>
              <td class="p-8">
                <div class="flex items-center gap-3">
                  <LucideCalendar class="w-4 h-4 text-slate-300" />
                  <p class="text-xs font-black text-slate-600 dark:text-slate-300">Created: {{ new Date(mission.created_at?.Time || mission.created_at).toLocaleDateString() }}</p>
                </div>
              </td>
              <td class="p-8">
                <div class="flex items-center gap-3">
                  <div class="w-8 h-8 rounded-full bg-primary-100 dark:bg-primary-900/30 flex items-center justify-center text-[10px] font-black text-primary-600">
                    {{ (mission.assigned_operator_name || 'OP').substring(0, 2).toUpperCase() }}
                  </div>
                  <p class="text-xs font-black text-slate-700 dark:text-slate-200">{{ mission.assigned_operator_name || 'Unassigned' }}</p>
                </div>
              </td>
              <td class="p-8 text-center">
                <span :class="`px-3 py-1 rounded-md text-[8px] font-black tracking-widest ${mission.target_area?.Valid ? 'bg-purple-50 text-purple-500' : 'bg-blue-50 text-blue-500'}`">
                  {{ mission.target_area?.Valid ? 'SPOT CHECK' : 'SCHEDULED' }}
                </span>
              </td>
              <td class="p-8">
                <div class="flex items-center gap-2">
                  <span :class="`w-2 h-2 rounded-full ${getStatusDotClass(mission.status?.String)}`"></span>
                  <span :class="`text-xs font-black ${getStatusTextClass(mission.status?.String)}`">{{ getStatusLabel(mission.status?.String) }}</span>
                </div>
              </td>
              <td class="p-8 text-center px-10">
                <button @click="handleAction(mission)" class="px-5 py-2 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-lg text-[9px] font-black uppercase tracking-widest transition-all">
                  {{ getActionLabel(mission.status?.String) }}
                </button>
              </td>
            </tr>
            <tr v-if="filteredMissions.length === 0">
              <td colspan="6" class="p-16 text-center text-xs font-bold text-slate-400 italic">
                No audit missions found.
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="p-8 border-t border-slate-100 dark:border-slate-800 flex items-center justify-between bg-slate-50/10">
        <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">
          Showing {{ totalMissions > 0 ? (page - 1) * limit + 1 : 0 }}-{{ Math.min(page * limit, totalMissions) }} of {{ totalMissions }} missions
        </p>
        <div class="flex items-center gap-2">
          <button :disabled="page === 1" @click="page--" class="w-8 h-8 flex items-center justify-center rounded-lg border border-slate-100 dark:border-slate-800 text-slate-400 hover:bg-slate-50 transition-all disabled:opacity-50"><LucideChevronLeft class="w-4 h-4" /></button>
          <button v-for="p in Math.ceil(totalMissions / limit)" :key="p" @click="page = p" :class="`w-8 h-8 flex items-center justify-center rounded-lg text-[10px] font-black transition-all ${page === p ? 'bg-[#1E3A5F] text-white shadow-lg' : 'border border-slate-100 dark:border-slate-800 text-slate-400 hover:bg-slate-50'}`">{{ p }}</button>
          <button :disabled="page * limit >= totalMissions" @click="page++" class="w-8 h-8 flex items-center justify-center rounded-lg border border-slate-100 dark:border-slate-800 text-slate-400 hover:bg-slate-50 transition-all disabled:opacity-50"><LucideChevronRight class="w-4 h-4" /></button>
        </div>
      </div>
    </div>

    <!-- Tip Box -->
    <div class="p-10 bg-primary-50/30 dark:bg-primary-900/10 border border-primary-100/50 dark:border-primary-800/30 rounded-lg flex gap-8" v-motion-slide-visible-bottom>
      <div class="w-14 h-14 rounded-2xl bg-white dark:bg-slate-800 text-primary-500 flex items-center justify-center shadow-sm shrink-0">
        <LucideInfo class="w-7 h-7" />
      </div>
      <div class="space-y-2">
        <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('stock.missions.tip.title') }}</h4>
        <p class="text-xs font-bold text-slate-500 dark:text-slate-400 leading-relaxed max-w-4xl">
          {{ $t('stock.missions.tip.desc') }}
        </p>
      </div>
    </div>

    <!-- Create Mission Modal -->
    <div v-if="showCreateModal" class="fixed inset-0 bg-[#1E3A5F]/40 backdrop-blur-sm z-50 flex items-center justify-center p-10" @click.self="showCreateModal = false">
      <div class="bg-white dark:bg-slate-900 rounded-lg w-full max-w-2xl overflow-hidden shadow-2xl border border-slate-100 dark:border-slate-800" v-motion-pop>
        <div class="bg-[#1E3A5F] px-10 py-6 flex items-center justify-between text-white">
          <h3 class="text-sm font-black uppercase tracking-widest">Create Audit Mission</h3>
          <button @click="showCreateModal = false" class="text-white/60 hover:text-white transition-colors">
            <LucideX class="w-6 h-6" />
          </button>
        </div>
        
        <form @submit.prevent="handleCreateMission" class="p-10 space-y-8">
          <div class="space-y-2">
            <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Mission Title <span class="text-red-500">*</span></label>
            <input type="text" v-model="newMission.title" required placeholder="e.g. Audit Legal Docs Q2" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all" />
          </div>
          
          <div class="space-y-2">
            <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Description</label>
            <textarea v-model="newMission.description" placeholder="Describe the scope or details of this audit..." rows="3" class="w-full px-5 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-3xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all resize-none"></textarea>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="space-y-2">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Assigned Operator <span class="text-red-500">*</span></label>
              <div class="relative">
                <select v-model="newMission.assigned_to" required class="w-full pl-5 pr-12 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold appearance-none outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all cursor-pointer">
                  <option value="" disabled selected>Select Operator</option>
                  <option v-for="op in operators" :key="op.id" :value="op.id">{{ op.full_name }}</option>
                </select>
                <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 pointer-events-none" />
              </div>
            </div>
            
            <div class="space-y-2">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Target Area / Zone <span class="text-red-500">*</span></label>
              <input type="text" v-model="newMission.target_area" required placeholder="e.g. Legal, Racks A1-A5" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all" />
            </div>
          </div>
          
          <div class="flex items-center justify-end gap-4 pt-6 border-t border-slate-50 dark:border-slate-800">
            <button type="button" @click="showCreateModal = false" class="px-6 py-3.5 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-slate-400 rounded-xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all">Cancel</button>
            <button type="submit" :disabled="isSubmitting" class="px-8 py-3.5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 transition-all flex items-center gap-2">
              <span v-if="isSubmitting">Creating...</span>
              <span v-else>Create Mission</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, computed } from 'vue'
import { 
  LucideClipboardList, LucideUsers, LucideCheckCircle2, 
  LucideSearch, LucideFilter, LucideDownload, LucideCalendar, 
  LucideMoreVertical, LucideChevronLeft, LucideChevronRight, LucideInfo,
  LucidePlus, LucideX, LucideChevronDown
} from 'lucide-vue-next'

const { $api } = useApi()
const toast = useToast()

const activeTab = ref('all')
const page = ref(1)
const limit = ref(10)
const totalMissions = ref(0)
const searchQuery = ref('')

const stats = ref({
  pending_missions: 0,
  active_missions: 0,
  accuracy_rate: 0
})

const missions = ref([])
const operators = ref([])
const showCreateModal = ref(false)

const newMission = ref({
  title: '',
  description: '',
  assigned_to: '',
  target_area: ''
})
const isSubmitting = ref(false)

const fetchStats = async () => {
  try {
    const res = await $api('/stock/missions/stats')
    if (res.success && res.data) {
      stats.value = res.data
    }
  } catch (err) {
    console.error('Failed to fetch stock stats:', err)
  }
}

const fetchMissions = async () => {
  try {
    const res = await $api(`/stock/missions?tab=${activeTab.value}&page=${page.value}&limit=${limit.value}`)
    if (res.success && res.data) {
      missions.value = res.data.items || []
      totalMissions.value = res.data.total || 0
    }
  } catch (err) {
    console.error('Failed to fetch stock missions:', err)
    toast.error('Failed to load audit missions.')
  }
}

const fetchOperators = async () => {
  try {
    const res = await $api('/users')
    if (res.success && res.data) {
      operators.value = res.data || []
    }
  } catch (err) {
    console.error('Failed to fetch operators:', err)
  }
}

const openCreateModal = async () => {
  showCreateModal.value = true
  await fetchOperators()
}

const handleCreateMission = async () => {
  if (!newMission.value.title || !newMission.value.assigned_to || !newMission.value.target_area) {
    toast.warning('Please fill in all required fields.')
    return
  }
  isSubmitting.value = true
  try {
    const res = await $api('/stock/missions', {
      method: 'POST',
      body: newMission.value
    })
    if (res.success) {
      toast.success('Audit mission created successfully!')
      showCreateModal.value = false
      newMission.value = { title: '', description: '', assigned_to: '', target_area: '' }
      await fetchMissions()
      await fetchStats()
    } else {
      toast.error(res.message || 'Failed to create mission.')
    }
  } catch (err) {
    console.error(err)
    toast.error(err.data?.message || 'Error creating audit mission.')
  } finally {
    isSubmitting.value = false
  }
}

const handleAction = async (mission) => {
  const status = (mission.status?.String || '').toLowerCase()
  if (status === 'pending' || status === 'not started' || status === '') {
    try {
      const res = await $api(`/stock/missions/${mission.id}/start`, { method: 'POST' })
      if (res.success) {
        toast.success('Mission started successfully.')
        navigateTo(`/stock/scan?session_id=${mission.session_id?.Bytes || ''}&mission_id=${mission.id}`)
      } else {
        toast.error(res.message || 'Failed to start mission.')
      }
    } catch (err) {
      console.error(err)
      toast.error(err.data?.message || 'Error starting mission.')
    }
  } else if (status === 'in_progress') {
    navigateTo(`/stock/scan?session_id=${mission.session_id?.Bytes || ''}&mission_id=${mission.id}`)
  } else if (status === 'completed' || status === 'approved') {
    navigateTo(`/stock/reconciliation?session_id=${mission.session_id?.Bytes || ''}&mission_id=${mission.id}`)
  }
}

const getStatusLabel = (status) => {
  const s = (status || '').toLowerCase()
  if (s === 'in_progress') return 'In Progress'
  if (s === 'completed') return 'Completed'
  if (s === 'approved') return 'Approved'
  return 'Pending'
}

const getStatusDotClass = (status) => {
  const s = (status || '').toLowerCase()
  if (s === 'in_progress') return 'bg-blue-500 shadow-[0_0_8px_rgba(59,130,246,0.5)]'
  if (s === 'completed') return 'bg-green-500 shadow-[0_0_8px_rgba(34,197,94,0.5)]'
  if (s === 'approved') return 'bg-emerald-500 shadow-[0_0_8px_rgba(16,185,129,0.5)]'
  return 'bg-slate-300'
}

const getStatusTextClass = (status) => {
  const s = (status || '').toLowerCase()
  if (s === 'in_progress') return 'text-blue-600'
  if (s === 'completed') return 'text-green-600'
  if (s === 'approved') return 'text-emerald-600'
  return 'text-slate-400'
}

const getActionLabel = (status) => {
  const s = (status || '').toLowerCase()
  if (s === 'in_progress') return 'Resume Audit'
  if (s === 'completed' || s === 'approved') return 'View Report'
  return 'Start Audit'
}

const filteredMissions = computed(() => {
  if (!searchQuery.value) return missions.value
  const q = searchQuery.value.toLowerCase()
  return missions.value.filter(m => 
    m.title.toLowerCase().includes(q) || 
    (m.assigned_operator_name && m.assigned_operator_name.toLowerCase().includes(q)) || 
    (m.target_area?.String && m.target_area.String.toLowerCase().includes(q))
  )
})

watch([activeTab, page], () => {
  fetchMissions()
})

onMounted(() => {
  fetchStats()
  fetchMissions()
})
</script>

<style scoped>
@reference "tailwindcss";
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}
</style>
