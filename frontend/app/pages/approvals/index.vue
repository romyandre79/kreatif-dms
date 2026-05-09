<template>
  <div class="max-w-7xl mx-auto space-y-8 pb-32">
    <!-- Header -->
    <div class="flex flex-col gap-6" v-motion-fade>
      <div class="flex items-center justify-between">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
          {{ $t('approvals.list.title') || 'All Pending Submissions' }}
        </h1>
        <div class="flex items-center gap-2 text-xs font-bold text-slate-400">
          <span class="w-2 h-2 bg-primary-500 rounded-full animate-pulse"></span>
          {{ totalCount }} Pending Documents
        </div>
      </div>

      <!-- Search & Filters -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
        <div class="md:col-span-2 relative group">
          <LucideSearch class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 group-focus-within:text-primary-500 transition-colors" />
          <input 
            v-model="searchQuery"
            type="text" 
            placeholder="Search by ID, name, or document title..."
            class="w-full bg-white border border-slate-200 rounded-xl py-3 pl-12 pr-4 text-sm font-medium outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all shadow-sm"
          />
        </div>
        <div class="relative">
          <select 
            v-model="filterDept"
            class="w-full bg-white border border-slate-200 rounded-xl py-3 px-4 text-sm font-bold appearance-none outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all shadow-sm cursor-pointer"
          >
            <option value="">All Departments</option>
            <option v-for="dept in departments" :key="dept.id" :value="dept.id">{{ dept.name }}</option>
          </select>
          <LucideChevronDown class="absolute right-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 pointer-events-none" />
        </div>
        <div class="relative">
          <input 
            type="text" 
            placeholder="Oct 01, 2023 - Oct 31, 2023"
            class="w-full bg-white border border-slate-200 rounded-xl py-3 px-4 text-sm font-bold outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all shadow-sm"
            readonly
          />
          <LucideCalendar class="absolute right-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
        </div>
      </div>
    </div>

    <!-- Table Container -->
    <div class="glass rounded-3xl overflow-hidden" v-motion-fade>
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-slate-50/50 border-b border-slate-100">
            <th class="pl-8 py-5 w-12">
              <input 
                type="checkbox" 
                :checked="isAllSelected"
                @change="toggleSelectAll"
                class="w-4 h-4 rounded border-slate-300 text-primary-600 focus:ring-primary-500 cursor-pointer"
              />
            </th>
            <th class="px-6 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('approvals.list.table.reg_id') || 'Registration ID' }}</th>
            <th class="px-6 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('approvals.list.table.submitted_by') || 'Submitted By' }}</th>
            <th class="px-6 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('approvals.list.table.doc_info') || 'Document Info' }}</th>
            <th class="px-6 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('approvals.list.table.date') || 'Submission Date' }}</th>
            <th class="px-6 py-5 text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('approvals.list.table.urgency') || 'Urgency' }}</th>
            <th class="pr-8 py-5 text-right text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('approvals.list.table.actions') || 'Actions' }}</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-50">
          <tr 
            v-for="item in items" 
            :key="item.id"
            class="group hover:bg-slate-50/50 transition-colors cursor-pointer"
            @click="navigateTo(`/approvals/${item.entity_id || item.id}`)"
          >
            <td class="pl-8 py-6" @click.stop>
              <input 
                type="checkbox" 
                v-model="selectedIds"
                :value="item.id"
                class="w-4 h-4 rounded border-slate-300 text-primary-600 focus:ring-primary-500 cursor-pointer"
              />
            </td>
            <td class="px-6 py-6">
              <span class="text-xs font-black text-[#1E3A5F] uppercase tracking-tight">REQ-{{ item.id?.substring(0, 8).toUpperCase() }}</span>
            </td>
            <td class="px-6 py-6">
              <div class="flex items-center gap-3">
                <div class="w-8 h-8 rounded-full bg-slate-100 overflow-hidden ring-2 ring-white shadow-sm">
                  <img :src="`https://i.pravatar.cc/100?u=${item.owner_name}`" class="w-full h-full object-cover" />
                </div>
                <div>
                  <p class="text-xs font-black text-[#1E3A5F] leading-tight">{{ item.owner_name || 'User' }}</p>
                  <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest mt-0.5">{{ item.department_name || 'Department' }}</p>
                </div>
              </div>
            </td>
            <td class="px-6 py-6">
              <div class="space-y-0.5">
                <p class="text-xs font-black text-[#1E3A5F] leading-tight">{{ item.title || 'Untitled Document' }}</p>
                <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ item.type_name || 'Submission' }} • {{ item.metadata?.page_count || 1 }} Files</p>
              </div>
            </td>
            <td class="px-6 py-6">
              <div class="space-y-0.5">
                <p class="text-xs font-black text-slate-700 leading-tight">{{ new Date(item.created_at).toLocaleDateString('en-GB', { day: '2-digit', month: 'short', year: 'numeric' }).toUpperCase() }}</p>
                <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ new Date(item.created_at).toLocaleTimeString('en-GB', { hour: '2-digit', minute: '2-digit' }) }}</p>
              </div>
            </td>
            <td class="px-6 py-6">
              <span 
                :class="`px-3 py-1 rounded-full text-[9px] font-black uppercase tracking-widest border ${
                  item.level > 1 
                    ? 'bg-red-50 text-red-500 border-red-100' 
                    : 'bg-slate-50 text-slate-500 border-slate-100'
                }`"
              >
                {{ item.level > 1 ? 'Urgent' : 'Normal' }}
              </span>
            </td>
            <td class="pr-8 py-6 text-right">
              <button 
                class="px-5 py-2 bg-[#1E3A5F] hover:bg-[#2A4B75] text-white rounded-lg text-[10px] font-black uppercase tracking-widest flex items-center gap-2 ml-auto transition-all group-hover:scale-105 active:scale-95 shadow-lg shadow-blue-900/10"
              >
                Review
                <LucideArrowRight class="w-3 h-3 group-hover:translate-x-1 transition-transform" />
              </button>
            </td>
          </tr>
        </tbody>
      </table>

      <!-- Pagination -->
      <div class="px-8 py-6 bg-slate-50/30 flex items-center justify-between border-t border-slate-100">
        <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">
          Showing 1-{{ items.length }} of {{ totalCount }} pending documents
        </p>
        <div class="flex items-center gap-2">
          <button class="w-8 h-8 rounded-lg bg-white border border-slate-200 flex items-center justify-center text-slate-400 hover:text-primary-600 hover:border-primary-200 transition-all">
            <LucideChevronLeft class="w-4 h-4" />
          </button>
          <button class="w-8 h-8 rounded-lg bg-white border border-slate-200 flex items-center justify-center text-slate-400 hover:text-primary-600 hover:border-primary-200 transition-all">
            <LucideChevronRight class="w-4 h-4" />
          </button>
        </div>
      </div>
    </div>

    <!-- Bulk Action Footer -->
    <div 
      v-if="selectedIds.length > 0"
      class="fixed bottom-8 left-1/2 -translate-x-1/2 w-[90%] max-w-4xl bg-[#1E3A5F]/95 backdrop-blur-xl border border-white/20 rounded-2xl p-4 shadow-2xl z-[100] flex items-center justify-between"
      v-motion-slide-bottom
    >
      <div class="flex items-center gap-4 pl-4">
        <div class="w-10 h-10 rounded-full bg-white/10 flex items-center justify-center text-white ring-4 ring-white/5">
          <LucideCheckCircle class="w-5 h-5" />
        </div>
        <div>
          <p class="text-sm font-black text-white">{{ selectedIds.length }} Documents Selected</p>
          <p class="text-[10px] text-white/50 font-bold uppercase tracking-widest">Bulk actions will be applied immediately.</p>
        </div>
      </div>
      
      <div class="flex items-center gap-3">
        <button 
          @click="bulkReject"
          class="px-6 py-3 bg-red-500 hover:bg-red-600 text-white rounded-xl text-[10px] font-black uppercase tracking-widest flex items-center gap-2 transition-all active:scale-95 shadow-lg shadow-red-500/20"
        >
          <LucideXCircle class="w-4 h-4" />
          Reject Selected
        </button>
        <button 
          @click="bulkApprove"
          class="px-6 py-3 bg-green-500 hover:bg-green-600 text-white rounded-xl text-[10px] font-black uppercase tracking-widest flex items-center gap-2 transition-all active:scale-95 shadow-lg shadow-green-500/20"
        >
          <LucideCheckCircle class="w-4 h-4" />
          Approve Selected
        </button>
        <div class="w-px h-8 bg-white/10 mx-2"></div>
        <button 
          @click="selectedIds = []"
          class="px-6 py-3 text-white/70 hover:text-white text-[10px] font-black uppercase tracking-widest transition-colors"
        >
          Cancel
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideSearch, LucideChevronDown, LucideCalendar, LucideArrowRight, 
  LucideChevronLeft, LucideChevronRight, LucideCheckCircle, LucideXCircle 
} from 'lucide-vue-next'
import { useApi } from '~/composables/useApi'

const { $api } = useApi()
const searchQuery = ref('')
const filterDept = ref('')
const selectedIds = ref([])

// Fetch Data
const { data: summaryRes } = await useAsyncData('approval-list', () => $api('/dashboard/summary'))
const items = computed(() => summaryRes.value?.data?.tasks || [])

// Search & Filter Logic
const filteredItems = computed(() => {
  let list = items.value
  
  if (searchQuery.value) {
    const q = searchQuery.value.toLowerCase()
    list = list.filter(i => 
      i.id.toLowerCase().includes(q) || 
      i.title?.toLowerCase().includes(q) || 
      i.owner_name?.toLowerCase().includes(q)
    )
  }
  
  if (filterDept.value) {
    list = list.filter(i => i.department_id === filterDept.value)
  }
  
  return list
})

const totalCount = computed(() => filteredItems.value.length)

// Mock departments - in real app fetch from API
const departments = [
  { id: 'f0a1b2c3-d4e5-4f6a-8b9c-0d1e2f3a4b5c', name: 'Finance Dept' },
  { id: 'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d', name: 'IT Support' },
  { id: 'b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e', name: 'Human Resources' }
]

const isAllSelected = computed(() => {
  return filteredItems.value.length > 0 && selectedIds.value.length === filteredItems.value.length
})

const toggleSelectAll = () => {
  if (isAllSelected.value) {
    selectedIds.value = []
  } else {
    selectedIds.value = filteredItems.value.map(i => i.id)
  }
}

const bulkApprove = async () => {
  if (confirm(`Approve ${selectedIds.value.length} documents?`)) {
    try {
      const res = await $api.post('/documents/bulk-approve', { ids: selectedIds.value })
      if (res.status === 'success') {
        alert('Bulk Approval Success')
        selectedIds.value = []
        refreshNuxtData('approval-list')
      }
    } catch (e) {
      console.error(e)
      alert('Bulk approval failed')
    }
  }
}

const bulkReject = async () => {
  if (confirm(`Reject ${selectedIds.value.length} documents?`)) {
    try {
      const res = await $api.post('/documents/bulk-reject', { ids: selectedIds.value })
      if (res.status === 'success') {
        alert('Bulk Rejection Success')
        selectedIds.value = []
        refreshNuxtData('approval-list')
      }
    } catch (e) {
      console.error(e)
      alert('Bulk rejection failed')
    }
  }
}
</script>

<style scoped>
.glass {
  background-color: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(40px);
  -webkit-backdrop-filter: blur(40px);
  border: 1px solid rgba(255, 255, 255, 1);
  box-shadow: 0 25px 50px -12px rgba(226, 232, 240, 0.5);
}

input[type="checkbox"] {
  border-radius: 4px;
  border-color: #cbd5e1;
  color: #2563eb;
  transition: all 0.2s;
}

input[type="checkbox"]:focus {
  --tw-ring-color: rgba(37, 99, 235, 0.1);
  ring: 4px;
}
</style>
