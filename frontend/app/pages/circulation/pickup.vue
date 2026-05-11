<template>
  <div class="max-w-[1600px] mx-auto flex gap-8 items-start">
    <!-- Main Content -->
    <div class="flex-1 space-y-8">
      <!-- Stats Summary -->
      <div class="grid grid-cols-4 gap-6" v-motion-slide-visible-bottom>
        <div class="glass p-8 rounded-3xl bg-white border border-slate-100 shadow-xl shadow-slate-200/50">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] mb-4">{{ $t('circulation.pickup.stats.waiting') }}</p>
          <p class="text-4xl font-black text-[#1E3A5F]">12</p>
        </div>
        <div class="glass p-8 rounded-3xl bg-white border border-slate-100 shadow-xl shadow-slate-200/50">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] mb-4">{{ $t('circulation.pickup.stats.ready_today') }}</p>
          <p class="text-4xl font-black text-[#1E3A5F]">8</p>
        </div>
        <div class="glass p-8 rounded-3xl bg-white border border-slate-100 shadow-xl shadow-slate-200/50">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] mb-4">{{ $t('circulation.pickup.stats.sla_risk') }}</p>
          <p class="text-4xl font-black text-red-500">2</p>
        </div>
        <div class="glass p-8 rounded-3xl bg-white border border-slate-100 shadow-xl shadow-slate-200/50">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] mb-4 text-orange-500">{{ $t('circulation.pickup.stats.high_priority') }}</p>
          <p class="text-4xl font-black text-orange-500">5</p>
        </div>
      </div>

      <!-- Queue Header -->
      <div class="flex items-center justify-between">
        <h2 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('circulation.pickup.active_requests', { count: 12 }) }}</h2>
        <div class="flex items-center gap-4">
          <div class="relative group w-80">
            <LucideSearch class="absolute left-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 group-focus-within:text-primary-500 transition-colors" />
            <input type="text" :placeholder="$t('circulation.pickup.search_placeholder')" 
                   class="w-full pl-12 pr-6 py-4 bg-white border border-slate-200 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all shadow-sm" />
          </div>
          <button class="p-4 bg-white border border-slate-200 rounded-2xl text-slate-400 hover:text-slate-600 transition-all shadow-sm">
            <LucideFilter class="w-5 h-5" />
          </button>
        </div>
      </div>

      <!-- Requests List -->
      <div class="space-y-4">
        <div v-for="req in queue" :key="req.no" 
             @click="selectedRequest = req"
             :class="`group p-8 bg-white rounded-lg border transition-all cursor-pointer flex items-center justify-between shadow-xl shadow-slate-200/30 ${selectedRequest?.no === req.no ? 'border-primary-500 ring-4 ring-primary-500/10' : 'border-slate-100 hover:border-slate-300'}`">
          <div class="flex items-center gap-8">
            <div class="flex items-center gap-4 min-w-[200px]">
              <div :class="`w-2 h-2 rounded-full ${req.priority === 'High' ? 'bg-red-500 animate-pulse' : 'bg-slate-300'}`"></div>
              <div>
                <p class="text-lg font-black text-[#1E3A5F] group-hover:text-primary-600 transition-colors">{{ req.no }}</p>
                <p class="text-xs font-bold text-slate-600">{{ req.requestor }}</p>
                <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest mt-1">{{ $t('circulation.pickup.table.approved_at', { time: req.approvedAt }) }}</p>
              </div>
            </div>
            
            <div class="min-w-[150px]">
              <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest mb-1">{{ $t('circulation.pickup.table.dept') }}</p>
              <p class="text-xs font-bold text-slate-700 uppercase">{{ req.dept }}</p>
            </div>

            <div class="min-w-[120px]">
              <span class="px-3 py-1 bg-emerald-50 text-emerald-600 border border-emerald-100 rounded-lg text-[9px] font-black uppercase tracking-widest">
                {{ $t('circulation.pickup.table.status_complete') }}
              </span>
            </div>

            <div class="flex items-center gap-2 text-slate-400">
              <LucideFileStack class="w-4 h-4" />
              <span class="text-xs font-black">{{ $t('circulation.pickup.table.docs_count', { count: req.docsCount }) }}</span>
            </div>
          </div>

          <button class="px-8 py-4 bg-slate-50 border border-slate-200 text-[#1E3A5F] rounded-2xl text-xs font-black uppercase tracking-widest hover:bg-primary-500 hover:text-white hover:border-primary-500 transition-all">
            {{ $t('circulation.pickup.table.btn_review') }}
          </button>
        </div>
      </div>
    </div>

    <!-- Sidebar Detail Panel -->
    <aside class="w-[480px] h-[calc(100vh-10rem)] sticky top-32 flex flex-col glass bg-white border border-slate-100 rounded-lg shadow-2xl shadow-slate-200/50 overflow-hidden" v-motion-slide-right>
      <div v-if="selectedRequest" class="flex flex-col h-full">
        <!-- Sidebar Header -->
        <div class="px-10 py-8 border-b border-slate-50 flex items-center justify-between">
          <div>
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">{{ $t('circulation.pickup.detail.title') }}</p>
            <h3 class="text-xl font-black text-[#1E3A5F] tracking-tight uppercase">{{ selectedRequest.no }}</h3>
          </div>
          <button @click="selectedRequest = null" class="p-3 hover:bg-slate-50 rounded-2xl text-slate-400 transition-colors">
            <LucideX class="w-5 h-5" />
          </button>
        </div>

        <div class="flex-1 overflow-y-auto p-10 space-y-10 custom-scrollbar">
          <!-- Approval Summary Card -->
          <div class="p-8 bg-slate-50/50 border border-slate-100 rounded-lg space-y-8">
            <div class="flex items-center justify-between">
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.pickup.detail.approved_summary') }}</p>
            </div>
            <div class="grid grid-cols-2 gap-6">
              <div>
                <p class="text-[9px] font-bold text-slate-400 uppercase mb-1">{{ $t('circulation.pickup.detail.reason') }}</p>
                <p class="text-xs font-black text-[#1E3A5F] uppercase">{{ selectedRequest.reason }}</p>
              </div>
              <div>
                <p class="text-[9px] font-bold text-slate-400 uppercase mb-1">{{ $t('circulation.pickup.detail.priority') }}</p>
                <p :class="`text-xs font-black uppercase ${selectedRequest.priority === 'High' ? 'text-red-500' : 'text-emerald-500'}`">
                  {{ $t('circulation.pickup.detail.priority_level', { level: selectedRequest.priority }) }}
                </p>
              </div>
            </div>
            <div>
              <p class="text-[9px] font-bold text-slate-400 uppercase mb-2">{{ $t('circulation.pickup.detail.reviewer_note') }}</p>
              <p class="text-[11px] font-medium text-slate-600 leading-relaxed italic">
                "{{ selectedRequest.reviewerNote }}"
              </p>
            </div>
          </div>

          <!-- Document Checklist -->
          <div class="space-y-6">
            <div class="flex items-center justify-between">
              <h4 class="text-xs font-black text-[#1E3A5F] uppercase tracking-widest">{{ $t('circulation.pickup.detail.checklist_title', { count: selectedRequest.docsCount }) }}</h4>
              <span class="text-[9px] font-bold text-slate-400 uppercase">{{ $t('circulation.pickup.detail.checklist_step', { current: 3, total: selectedRequest.docsCount }) }}</span>
            </div>
            <div class="space-y-4">
              <div v-for="doc in selectedRequest.documents" :key="doc.name" class="p-6 bg-white border border-slate-100 rounded-3xl space-y-6 group hover:border-primary-100 transition-colors">
                <div class="flex items-start justify-between">
                  <div class="flex items-start gap-4">
                    <div class="p-3 bg-slate-50 rounded-2xl text-slate-400 group-hover:text-primary-500 transition-colors">
                      <LucideFileText class="w-5 h-5" />
                    </div>
                    <div class="space-y-1">
                      <p class="text-xs font-black text-[#1E3A5F] group-hover:text-primary-600 transition-colors">{{ doc.name }}</p>
                      <p class="text-[10px] font-bold text-slate-400 flex items-center gap-1.5">
                        <LucideMapPin class="w-3.5 h-3.5" />
                        {{ doc.location }}
                      </p>
                    </div>
                  </div>
                  <div v-if="doc.status === 'found'" class="text-emerald-500">
                    <LucideCheckCircle2 class="w-5 h-5" />
                  </div>
                </div>

                <div class="grid grid-cols-3 gap-3">
                  <button @click="doc.status = 'found'" :class="`py-3 rounded-xl text-[9px] font-black uppercase tracking-widest transition-all ${doc.status === 'found' ? 'bg-emerald-500 text-white shadow-lg shadow-emerald-500/20' : 'bg-slate-50 text-slate-400 hover:bg-slate-100'}`">
                    {{ $t('circulation.pickup.detail.btn_found') }}
                  </button>
                  <button @click="doc.status = 'not_found'" :class="`py-3 rounded-xl text-[9px] font-black uppercase tracking-widest transition-all ${doc.status === 'not_found' ? 'bg-red-500 text-white shadow-lg shadow-red-500/20' : 'bg-slate-50 text-slate-400 hover:bg-slate-100'}`">
                    {{ $t('circulation.pickup.detail.btn_not_found') }}
                  </button>
                  <button @click="doc.status = 'damaged'" :class="`py-3 rounded-xl text-[9px] font-black uppercase tracking-widest transition-all ${doc.status === 'damaged' ? 'bg-orange-500 text-white shadow-lg shadow-orange-500/20' : 'bg-slate-50 text-slate-400 hover:bg-slate-100'}`">
                    {{ $t('circulation.pickup.detail.btn_damaged') }}
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- Admin Note -->
          <div class="space-y-4 pt-6 border-t border-slate-50">
            <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.pickup.detail.admin_note') }}</label>
            <textarea :placeholder="$t('circulation.pickup.detail.admin_note_placeholder')" 
                      rows="3" 
                      class="w-full p-6 bg-slate-50 border border-slate-100 rounded-3xl text-sm font-medium text-slate-600 outline-none focus:ring-4 focus:ring-primary-500/5 focus:border-primary-500 transition-all resize-none"></textarea>
          </div>
        </div>

        <!-- Footer Actions -->
        <div class="p-10 border-t border-slate-50 bg-slate-50/20 space-y-4">
          <button class="w-full py-5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-xs font-black uppercase tracking-[0.2em] shadow-xl shadow-blue-900/20 transition-all">
            {{ $t('circulation.pickup.detail.btn_ready') }}
          </button>
          <div class="grid grid-cols-2 gap-4">
            <button class="py-4 border border-slate-200 text-slate-400 hover:bg-slate-50 rounded-2xl text-[9px] font-black uppercase tracking-widest transition-all flex items-center justify-center gap-2">
              <LucideAlertTriangle class="w-3.5 h-3.5" /> {{ $t('circulation.pickup.detail.btn_escalate') }}
            </button>
            <button class="py-4 border border-slate-200 text-slate-400 hover:bg-slate-50 rounded-2xl text-[9px] font-black uppercase tracking-widest transition-all flex items-center justify-center gap-2">
              <LucideRotateCcw class="w-3.5 h-3.5" /> {{ $t('circulation.pickup.detail.btn_return') }}
            </button>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-else class="h-full flex flex-col items-center justify-center text-center p-12 space-y-6">
        <div class="w-24 h-24 bg-slate-50 rounded-full flex items-center justify-center text-slate-200">
          <LucideBox class="w-12 h-12" />
        </div>
        <div class="space-y-2">
          <h3 class="text-sm font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.pickup.empty.title') }}</h3>
          <p class="text-xs font-medium text-slate-300 leading-relaxed px-10">
            {{ $t('circulation.pickup.empty.desc') }}
          </p>
        </div>
      </div>
    </aside>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideSearch, LucideFilter, LucideFileStack, LucideX, 
  LucideFileText, LucideMapPin, LucideCheckCircle2, 
  LucideAlertTriangle, LucideRotateCcw, LucideBox
} from 'lucide-vue-next'

const queue = ref([
  {
    no: 'REQ-2024-001',
    requestor: 'John Doe',
    dept: 'Human Resources',
    approvedAt: '2h ago',
    priority: 'High',
    docsCount: 5,
    reason: 'Legal Review',
    reviewerNote: 'Ensure all employment contracts are original copies.',
    documents: [
      { name: 'Employment_Contract_JD01.pdf', location: 'Cabinet A1, Box 05', status: 'found' },
      { name: 'ID_Card_Copy_JD01.pdf', location: 'Cabinet A1, Box 05', status: 'none' },
      { name: 'NDA_Agreement_JD01.pdf', location: 'Cabinet B2, Box 12', status: 'none' }
    ]
  },
  {
    no: 'REQ-2024-004',
    requestor: 'Jane Smith',
    dept: 'Procurement',
    approvedAt: '4h ago',
    priority: 'Normal',
    docsCount: 3,
    reason: 'Vendor Audit',
    reviewerNote: 'Check for wet signatures on page 12.',
    documents: [
      { name: 'Vendor_Agreement_V2.pdf', location: 'Cabinet C3, Box 01', status: 'none' }
    ]
  },
  {
    no: 'REQ-2024-007',
    requestor: 'Robert King',
    dept: 'Finance',
    approvedAt: '1d ago',
    priority: 'Normal',
    docsCount: 12,
    reason: 'Internal Audit',
    reviewerNote: 'Full ledger documents for Q4.',
    documents: [
      { name: 'Q4_Tax_Report.pdf', location: 'Cabinet D1, Box 08', status: 'none' }
    ]
  }
])

const selectedRequest = ref(queue.value[0])
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #E2E8F0; border-radius: 10px; }
</style>
