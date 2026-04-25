<template>
  <div class="max-w-[1600px] mx-auto space-y-10 pb-20">
    <!-- Header -->
    <div class="space-y-4" v-motion-fade>
      <div class="flex items-center gap-2 text-[10px] font-black text-slate-300 uppercase tracking-widest">
        <span>{{ $t('layout.menu.approvals') }}</span>
        <LucideChevronRight class="w-3 h-3" />
        <span class="text-primary-500">{{ $t('layout.menu.extensions') }}</span>
      </div>
      <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
        <div class="space-y-2">
          <h1 class="text-3xl font-black text-[#1E3A5F] tracking-tight uppercase">
            {{ $t('approvals.extensions.title') }}
          </h1>
          <p class="text-slate-500 font-medium">
            {{ $t('approvals.extensions.subtitle') }}
          </p>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-12 gap-10 items-start">
      <!-- Main Table -->
      <div class="col-span-8 space-y-6" v-motion-slide-visible-bottom>
        <div class="glass rounded-[2.5rem] overflow-hidden border border-slate-100 bg-white shadow-xl shadow-slate-200/50">
          <table class="w-full text-left">
            <thead>
              <tr class="bg-slate-50/50 text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100">
                <th class="px-8 py-6">{{ $t('approvals.extensions.table.req_no') }}</th>
                <th class="px-8 py-6">{{ $t('approvals.extensions.table.requestor') }}</th>
                <th class="px-8 py-6">{{ $t('approvals.extensions.table.current_due') }}</th>
                <th class="px-8 py-6">{{ $t('approvals.extensions.table.requested_days') }}</th>
                <th class="px-8 py-6">{{ $t('approvals.extensions.table.new_due') }}</th>
                <th class="px-8 py-6">{{ $t('approvals.extensions.table.reason') }}</th>
                <th class="px-8 py-6">{{ $t('approvals.extensions.table.risk_flag') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50">
              <tr v-for="req in queue" :key="req.no" 
                  @click="selectedRequest = req"
                  :class="`group cursor-pointer transition-all ${selectedRequest?.no === req.no ? 'bg-primary-50/30' : 'hover:bg-slate-50/50'}`">
                <td class="px-8 py-8">
                  <p :class="`text-sm font-black uppercase tracking-tight ${selectedRequest?.no === req.no ? 'text-primary-600' : 'text-[#1E3A5F]'}`">{{ req.no }}</p>
                </td>
                <td class="px-8 py-8">
                  <div class="flex items-center gap-3">
                    <img :src="req.avatar" class="w-8 h-8 rounded-full border-2 border-white shadow-sm" />
                    <p class="text-sm font-bold text-slate-700">{{ req.requestor }}</p>
                  </div>
                </td>
                <td class="px-8 py-8 text-xs font-bold text-slate-400 uppercase">{{ req.currentDue }}</td>
                <td class="px-8 py-8">
                  <span :class="`px-3 py-1 rounded-lg text-[9px] font-black uppercase tracking-widest ${req.days === 7 ? 'bg-blue-50 text-blue-600 border border-blue-100' : 'bg-slate-50 text-slate-400 border border-slate-100'}`">
                    {{ $t('approvals.extensions.table.days', { count: req.days }) }}
                  </span>
                </td>
                <td class="px-8 py-8 text-xs font-black text-[#1E3A5F] uppercase">{{ req.newDue }}</td>
                <td class="px-8 py-8">
                  <p class="text-xs font-bold text-slate-500">{{ req.reason }}</p>
                </td>
                <td class="px-8 py-8">
                  <span :class="`px-3 py-1 rounded-lg text-[9px] font-black uppercase tracking-widest ${req.risk === 'Low' ? 'bg-emerald-50 text-emerald-600 border border-emerald-100' : 'bg-orange-50 text-orange-600 border border-orange-100'}`">
                    {{ $t('approvals.extensions.table.risk_level', { level: req.risk }) }}
                  </span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Detail Panel -->
      <aside class="col-span-4" v-motion-slide-right>
        <div v-if="selectedRequest" class="glass bg-white border border-slate-100 rounded-[3rem] shadow-2xl shadow-slate-200/50 overflow-hidden flex flex-col">
          <!-- Detail Header -->
          <div class="px-10 py-8 border-b border-slate-50 flex items-center justify-between">
            <h3 class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('approvals.extensions.detail.title') }}</h3>
            <button @click="selectedRequest = null" class="p-2 hover:bg-slate-50 rounded-xl text-slate-300">
              <LucideX class="w-5 h-5" />
            </button>
          </div>

          <div class="p-10 space-y-10 overflow-y-auto custom-scrollbar max-h-[calc(100vh-25rem)]">
            <!-- Original Context -->
            <div class="space-y-6">
              <div class="flex items-center gap-3 text-[10px] font-black text-slate-400 uppercase tracking-widest">
                <LucideFolderOpen class="w-4 h-4 text-primary-500" />
                {{ $t('approvals.extensions.detail.context_title') }}
              </div>
              
              <div class="p-8 bg-slate-50/50 border border-slate-100 rounded-[2.5rem] space-y-6">
                <p class="text-[10px] font-black text-primary-600 uppercase tracking-widest">{{ $t('approvals.extensions.detail.docs_linked', { count: 5 }) }}</p>
                <div class="space-y-3">
                  <div v-for="doc in selectedRequest.docs" :key="doc" class="flex items-center gap-4 p-4 bg-white border border-slate-100 rounded-2xl">
                    <LucideFileText class="w-4 h-4 text-slate-300" />
                    <p class="text-xs font-black text-[#1E3A5F] uppercase tracking-tight">{{ doc }}</p>
                  </div>
                </div>
                <button class="w-full text-[9px] font-black text-slate-400 uppercase tracking-[0.2em] hover:text-primary-500 transition-colors">
                  {{ $t('approvals.extensions.detail.view_more', { count: 3 }) }}
                </button>
              </div>
            </div>

            <!-- User History -->
            <div class="grid grid-cols-2 gap-4">
              <div class="p-6 bg-slate-50 border border-slate-100 rounded-3xl space-y-2">
                <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest leading-relaxed">{{ $t('approvals.extensions.detail.history.return_title') }}</p>
                <p class="text-xl font-black text-emerald-500">12 <span class="text-[10px] uppercase font-bold text-slate-300">{{ $t('approvals.extensions.detail.history.return_success') }}</span></p>
              </div>
              <div class="p-6 bg-slate-50 border border-slate-100 rounded-3xl space-y-2">
                <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest leading-relaxed">{{ $t('approvals.extensions.detail.history.overdue_title') }}</p>
                <p class="text-xl font-black text-slate-300">0 <span class="text-[10px] uppercase font-bold">{{ $t('approvals.extensions.detail.history.overdue_active') }}</span></p>
              </div>
            </div>

            <!-- Prior Extensions -->
            <div class="p-6 bg-blue-50/30 border border-blue-100 rounded-3xl flex items-center justify-between">
              <div class="space-y-1">
                <p class="text-[10px] font-black text-[#1E3A5F] uppercase tracking-widest">{{ $t('approvals.extensions.detail.prior_ext.title') }}</p>
                <p class="text-[9px] font-bold text-slate-400">{{ $t('approvals.extensions.detail.prior_ext.first_request') }}</p>
              </div>
              <span class="text-2xl font-black text-[#1E3A5F]">0</span>
            </div>

            <!-- Risk Summary -->
            <div class="space-y-4">
              <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ $t('approvals.extensions.detail.risk_summary.title') }}</p>
              <div class="space-y-3">
                <div class="flex gap-3 items-start text-[10px] font-bold text-emerald-600">
                  <LucideCheckCircle2 class="w-4 h-4 shrink-0" />
                  {{ $t('approvals.extensions.detail.risk_summary.no_sensitive') }}
                </div>
                <div class="flex gap-3 items-start text-[10px] font-bold text-emerald-600">
                  <LucideCheckCircle2 class="w-4 h-4 shrink-0" />
                  {{ $t('approvals.extensions.detail.risk_summary.good_history') }}
                </div>
              </div>
            </div>
          </div>

          <!-- Actions -->
          <div class="p-10 border-t border-slate-50 space-y-4 bg-slate-50/20">
            <!-- Rejection Notes (Optional) -->
            <Transition name="fade">
              <div v-if="showRejectionNotes" class="space-y-4 pt-6" v-motion-slide-bottom>
                <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ $t('approvals.extensions.detail.rejection.title') }}</p>
                <textarea :placeholder="$t('approvals.extensions.detail.rejection.placeholder')" 
                          rows="3" 
                          class="w-full p-6 bg-slate-50 border border-slate-100 rounded-3xl text-sm font-medium text-slate-600 outline-none focus:ring-4 focus:ring-primary-500/5 focus:border-primary-500 transition-all resize-none"></textarea>
              </div>
            </Transition>

            <button @click="showPinModal = true" class="w-full py-5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-xs font-black uppercase tracking-[0.2em] shadow-xl shadow-blue-900/20 transition-all flex items-center justify-center gap-3">
              <LucideCheck class="w-4 h-4" /> {{ $t('approvals.extensions.detail.btn_approve') }}
            </button>
            <button @click="handleReject" class="w-full py-4 bg-white border border-red-100 text-red-500 hover:bg-red-50 rounded-2xl text-xs font-black uppercase tracking-[0.2em] transition-all flex items-center justify-center gap-3">
              <LucideXCircle class="w-4 h-4" /> {{ showRejectionNotes ? $t('approvals.extensions.detail.btn_confirm_reject') : $t('approvals.extensions.detail.btn_reject') }}
            </button>
            <p class="text-[9px] font-bold text-slate-300 text-center uppercase tracking-widest pt-2">
              {{ $t('approvals.extensions.detail.audit_hint') }}
            </p>
          </div>
        </div>

        <!-- Empty State -->
        <div v-else class="h-full glass rounded-[3rem] border border-slate-100 bg-white flex flex-col items-center justify-center text-center p-12 space-y-6">
          <div class="w-20 h-20 bg-slate-50 rounded-full flex items-center justify-center text-slate-200">
            <LucideHistory class="w-10 h-10" />
          </div>
          <div class="space-y-2">
            <h3 class="text-sm font-black text-slate-400 uppercase tracking-widest">{{ $t('approvals.extensions.empty.title') }}</h3>
            <p class="text-xs font-medium text-slate-300 px-10">{{ $t('approvals.extensions.empty.desc') }}</p>
          </div>
        </div>
      </aside>
    </div>

    <!-- PIN Confirmation Modal -->
    <Transition name="scale">
      <div v-if="showPinModal" class="fixed inset-0 z-[100] flex items-center justify-center p-6 bg-[#1E3A5F]/40 backdrop-blur-sm">
        <div class="glass max-w-md w-full bg-white rounded-[3rem] p-12 shadow-2xl border border-white flex flex-col items-center text-center space-y-8" v-motion-pop>
          <!-- Modal Icon -->
          <div class="w-20 h-20 bg-slate-50 rounded-full flex items-center justify-center text-[#1E3A5F] border border-slate-100">
            <LucideLock class="w-10 h-10" />
          </div>

          <!-- Title & Subtitle -->
          <div class="space-y-4">
            <h2 class="text-2xl font-black text-[#1E3A5F]">{{ $t('approvals.extensions.modal.title') }}</h2>
            <p class="text-xs font-bold text-slate-400 leading-relaxed px-4">
              {{ $t('approvals.extensions.modal.subtitle') }}
            </p>
          </div>

          <!-- PIN Inputs -->
          <div class="flex gap-3 justify-center">
            <input v-for="i in 6" :key="i" type="password" maxlength="1" 
                   class="w-12 h-16 bg-slate-50 border border-slate-100 rounded-2xl text-center text-xl font-black text-[#1E3A5F] outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all shadow-sm" />
          </div>

          <!-- Actions -->
          <div class="w-full space-y-4 pt-4">
            <button @click="confirmPin" class="w-full py-5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-sm font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 transition-all">
              {{ $t('approvals.extensions.modal.btn_confirm') }}
            </button>
            <button @click="showPinModal = false" class="text-xs font-black text-slate-400 hover:text-slate-600 uppercase tracking-widest transition-colors">
              {{ $t('approvals.extensions.modal.btn_cancel') }}
            </button>
          </div>

          <!-- Footer -->
          <div class="w-full pt-8 border-t border-slate-50 flex items-center justify-center gap-2">
            <LucideShieldCheck class="w-3 h-3 text-slate-300" />
            <p class="text-[9px] font-bold text-slate-300 uppercase tracking-widest">
              {{ $t('approvals.extensions.modal.footer') }}
            </p>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideChevronRight, LucideChevronDown, LucideX, LucideFolderOpen, 
  LucideFileText, LucideCheckCircle2, LucideCheck, LucideXCircle, LucideHistory,
  LucideLock, LucideShieldCheck
} from 'lucide-vue-next'

const showPinModal = ref(false)
const showRejectionNotes = ref(false)

const queue = ref([
  {
    no: 'REQ-2023-001',
    requestor: 'Budi Santoso',
    avatar: 'https://i.pravatar.cc/150?u=budi',
    currentDue: 'Oct 15, 2023',
    days: 7,
    newDue: 'Oct 22, 2023',
    reason: 'Audit preparation',
    risk: 'Low',
    docs: ['Sertifikat Tanah A-12', 'Akta Jual Beli #901-22']
  },
  {
    no: 'REQ-2023-002',
    requestor: 'Siti Aminah',
    avatar: 'https://i.pravatar.cc/150?u=siti',
    currentDue: 'Oct 16, 2023',
    days: 14,
    newDue: 'Oct 30, 2023',
    reason: 'Project closure delays',
    risk: 'Med',
    docs: ['Contract_Vendor_IT.pdf', 'Legal_Opinion_Q4.docx']
  }
])

const selectedRequest = ref(queue.value[0])

const confirmPin = () => {
  alert(`Perpanjangan untuk ${selectedRequest.value.no} telah disetujui.`)
  showPinModal.value = false
}

const handleReject = () => {
  if (!showRejectionNotes.value) {
    showRejectionNotes.value = true
    return
  }
  
  alert(`Permohonan ${selectedRequest.value.no} telah ditolak.`)
  showRejectionNotes.value = false
}
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #E2E8F0; border-radius: 10px; }
</style>
