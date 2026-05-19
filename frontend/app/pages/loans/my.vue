<template>
  <div class="max-w-[1600px] mx-auto grid grid-cols-12 gap-10 items-start pb-20">
    <!-- Success Banner -->
    <Transition name="fade">
      <div v-if="submissionSuccess" class="col-span-12 p-4 bg-emerald-50 border border-emerald-100 rounded-2xl flex items-center justify-between gap-4 text-emerald-700 font-bold text-sm shadow-sm mb-4" v-motion-slide-top>
        <div class="flex items-center gap-4">
          <div class="w-8 h-8 bg-emerald-500 rounded-full flex items-center justify-center text-white shrink-0 shadow-lg shadow-emerald-500/20">
            <LucideCheck class="w-5 h-5" />
          </div>
          <p v-if="submissionSuccessDays">
            Perpanjangan {{ submissionSuccessDays }} hari untuk {{ selectedLoan?.no }} berhasil diajukan. Menunggu persetujuan tim Legal.
          </p>
          <p v-else>{{ $t('loans.my.success_banner') }}</p>
        </div>
        <button @click="submissionSuccess = false" class="text-emerald-400 hover:text-emerald-600">
          <LucideX class="w-4 h-4" />
        </button>
      </div>
    </Transition>

    <!-- Main Content: Loan Table -->
    <div class="col-span-8 space-y-10" v-motion-fade>
    <!-- Header Section -->
    <div class="col-span-12">
      <PageHeader 
        :title="$t('loans.my.title')"
        :subtitle="$t('loans.my.subtitle')"
      >
        <template #actions>
          <button @click="showExtensionModal = true" :disabled="selectedLoan?.status === 'EXTENSION PENDING'" class="px-8 py-4 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-xs font-black uppercase tracking-[0.2em] shadow-xl shadow-blue-900/20 transition-all flex items-center gap-3 group disabled:opacity-50 disabled:cursor-not-allowed">
            <LucideHistory class="w-4 h-4 group-hover:rotate-12 transition-transform" />
            {{ $t('loans.my.btn_request_ext') }}
          </button>
        </template>
      </PageHeader>
    </div>

      <!-- Filters/Tabs -->
      <div class="flex items-center gap-12 border-b border-slate-100 pb-1">
        <button v-for="tab in tabs" :key="tab.value"
                @click="activeTab = tab.value"
                :class="`pb-5 text-xs font-black uppercase tracking-widest transition-all relative ${activeTab === tab.value ? 'text-[#1E3A5F]' : 'text-slate-300 hover:text-slate-500'}`">
          {{ $t(tab.labelKey) }}
          <div v-if="activeTab === tab.value" class="absolute bottom-0 left-0 right-0 h-1 bg-[#1E3A5F] rounded-full" v-motion-pop></div>
        </button>
      </div>

      <!-- Table Card -->
      <div class="glass rounded-lg bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 overflow-hidden">
        <div class="overflow-x-auto">
          <table class="w-full text-left min-w-[700px]">
            <thead>
              <tr class="bg-slate-50/50 text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100">
                <th class="px-10 py-6">{{ $t('loans.my.table.req_no') }}</th>
                <th class="px-10 py-6">{{ $t('loans.my.table.total_docs') }}</th>
                <th class="px-10 py-6">{{ $t('loans.my.table.checkout_date') }}</th>
                <th class="px-10 py-6">{{ $t('loans.my.table.due_date') }}</th>
                <th class="px-10 py-6">{{ $t('loans.my.table.remaining_days') }}</th>
                <th class="px-10 py-6">{{ $t('loans.my.table.status') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50">
              <tr v-if="filteredLoans.length === 0">
                <td colspan="6" class="px-10 py-16 text-center text-sm font-bold text-slate-400 italic">
                  {{ isLoading ? 'Loading...' : 'Tidak ada data peminjaman.' }}
                </td>
              </tr>
              <tr v-for="loan in filteredLoans" :key="loan.no" 
                  @click="selectedLoan = loan"
                  :class="`group cursor-pointer transition-all ${selectedLoan?.no === loan.no ? 'bg-primary-50/30' : 'hover:bg-slate-50/50'}`">
                <td class="px-10 py-8">
                  <div class="flex items-center gap-4">
                    <div :class="`w-4 h-4 rounded-full border-2 transition-all flex items-center justify-center ${selectedLoan?.no === loan.no ? 'border-[#1E3A5F]' : 'border-slate-200'}`">
                      <div v-if="selectedLoan?.no === loan.no" class="w-1.5 h-1.5 bg-[#1E3A5F] rounded-full"></div>
                    </div>
                    <span :class="`text-sm font-black uppercase tracking-tight ${selectedLoan?.no === loan.no ? 'text-[#1E3A5F]' : 'text-slate-700'}`">{{ loan.no }}</span>
                  </div>
                </td>
                <td class="px-10 py-8 text-xs font-bold text-slate-500 uppercase">{{ $t('loans.my.table.docs_count', { count: loan.docsCount }) }}</td>
                <td class="px-10 py-8 text-xs font-bold text-slate-500 uppercase">{{ loan.checkoutDate }}</td>
                <td class="px-10 py-8 text-xs font-bold text-slate-500 uppercase">{{ loan.dueDate }}</td>
                <td class="px-10 py-8">
                  <span :class="`px-4 py-2 rounded-full text-[9px] font-black uppercase tracking-widest ${loan.remainingColor}`">
                    {{ loan.remainingText }}
                  </span>
                </td>
                <td class="px-10 py-8">
                  <span :class="`px-3 py-1 rounded-lg text-[9px] font-black uppercase tracking-widest ${loan.statusColor}`">
                    {{ loan.status }}
                  </span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <!-- Right Sidebar: Details & Timeline -->
    <div class="col-span-4 space-y-8" v-motion-slide-right>
      <div v-if="selectedLoan" class="glass bg-white border border-slate-100 rounded-lg shadow-2xl shadow-slate-200/50 overflow-hidden flex flex-col">
        <div class="p-10 space-y-10">
          <div class="flex items-center justify-between">
            <h3 class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('loans.my.detail.title') }}</h3>
            <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ selectedLoan.no }}</span>
          </div>

          <!-- Time Remaining Big Card -->
          <div class="p-8 bg-orange-50/50 border border-orange-100 rounded-lg text-center space-y-4">
            <p class="text-[10px] font-black text-orange-400 uppercase tracking-widest">{{ $t('loans.my.detail.time_remaining') }}</p>
            <div class="space-y-1">
              <p class="text-4xl font-black text-orange-500">{{ selectedLoan.remainingText }}</p>
              <p class="text-[10px] font-bold text-orange-400 uppercase">{{ $t('loans.my.detail.expires_hint', { date: selectedLoan.dueDate }) }}</p>
            </div>
          </div>

          <!-- Tracking Progress Timeline -->
          <div class="space-y-8 pt-4">
            <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ $t('loans.my.detail.tracking_progress') }}</p>
            <div class="space-y-0 relative">
              <div class="absolute left-[1.125rem] top-2 bottom-8 w-0.5 bg-slate-100"></div>
              
              <!-- Approved -->
              <div class="relative flex items-start gap-6 pb-10">
                <div class="relative z-10 w-9 h-9 rounded-full bg-[#1E3A5F] text-white flex items-center justify-center shadow-lg shadow-blue-900/10">
                  <LucideCheck class="w-4 h-4" />
                </div>
                <div>
                  <h4 class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('loans.my.detail.steps.approved') }}</h4>
                  <p class="text-[10px] font-bold text-slate-400 uppercase mt-1">{{ selectedLoan.approvedDate }}</p>
                </div>
              </div>

              <!-- Ready -->
              <div class="relative flex items-start gap-6 pb-10">
                <div class="relative z-10 w-9 h-9 rounded-full bg-[#1E3A5F] text-white flex items-center justify-center shadow-lg shadow-blue-900/10">
                  <LucideCheck class="w-4 h-4" />
                </div>
                <div>
                  <h4 class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('loans.my.detail.steps.ready') }}</h4>
                  <p class="text-[10px] font-bold text-slate-400 uppercase mt-1">{{ selectedLoan.readyDate }}</p>
                </div>
              </div>

              <!-- On Loan -->
              <div :class="`relative flex items-start gap-6 ${selectedLoan.status === 'EXTENSION PENDING' ? 'pb-10' : ''}`">
                <div :class="`relative z-10 w-9 h-9 rounded-full flex items-center justify-center shadow-lg ${selectedLoan.status === 'EXTENSION PENDING' ? 'bg-[#1E3A5F] text-white' : 'bg-white border-2 border-[#1E3A5F]'}`">
                  <LucideCheck v-if="selectedLoan.status === 'EXTENSION PENDING'" class="w-4 h-4" />
                  <div v-else class="w-2 h-2 bg-[#1E3A5F] rounded-full animate-pulse"></div>
                </div>
                <div>
                  <h4 class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('loans.my.detail.steps.on_loan') }}</h4>
                  <p class="text-[10px] font-bold text-slate-400 uppercase mt-1">
                    {{ selectedLoan.status === 'EXTENSION PENDING' ? selectedLoan.onLoanDate : $t('loans.my.detail.steps.current_status') }}
                  </p>
                </div>
              </div>

              <!-- Extension Requested (Pending State) -->
              <div v-if="selectedLoan.status === 'EXTENSION PENDING'" class="relative flex items-start gap-6">
                <div class="relative z-10 w-9 h-9 rounded-full bg-white border-2 border-orange-400 flex items-center justify-center shadow-lg">
                  <div class="w-2 h-2 bg-orange-500 rounded-full animate-pulse"></div>
                </div>
                <div>
                  <h4 class="text-sm font-black text-orange-500 uppercase tracking-tight">{{ $t('loans.my.detail.steps.ext_requested') }}</h4>
                  <p class="text-[10px] font-bold text-slate-400 uppercase mt-1">{{ $t('loans.my.detail.steps.waiting_legal') }}</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Policy Alerts -->
          <div class="space-y-4 pt-10 border-t border-slate-50">
            <div class="flex gap-4 items-start p-4 bg-slate-50/50 rounded-2xl">
              <LucideInfo class="w-4 h-4 text-slate-400 shrink-0 mt-0.5" />
              <p class="text-[10px] font-bold text-slate-400 leading-relaxed uppercase tracking-tight">
                {{ $t('loans.my.detail.policy.notif_hint') }}
              </p>
            </div>
            <div class="flex gap-4 items-start p-6 bg-red-50 border border-red-100 rounded-3xl">
              <LucideAlertTriangle class="w-5 h-5 text-red-500 shrink-0 mt-0.5" />
              <div class="space-y-1">
                <p class="text-[11px] font-black text-red-600 uppercase tracking-widest">{{ $t('loans.my.detail.policy.penalty_title') }}</p>
                <p class="text-[10px] font-bold text-red-500/70 leading-relaxed uppercase">{{ $t('loans.my.detail.policy.penalty_desc') }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Footer CTA (Changes when pending) -->
      <div v-if="selectedLoan?.status === 'EXTENSION PENDING'" class="glass p-10 bg-slate-50 rounded-lg border border-slate-100 shadow-xl space-y-6" v-motion-slide-bottom>
        <div class="space-y-2">
          <h3 class="text-xl font-black text-slate-700 uppercase tracking-tight">{{ $t('loans.my.detail.ext_pending.title') }}</h3>
          <p class="text-xs font-medium text-slate-500 leading-relaxed">
            {{ $t('loans.my.detail.ext_pending.desc') }}
          </p>
        </div>
      </div>
      <div v-else class="glass p-10 bg-[#1E3A5F] rounded-lg text-white shadow-2xl shadow-blue-900/30 space-y-6" v-motion-slide-bottom>
        <div class="space-y-2">
          <h3 class="text-xl font-black uppercase tracking-tight">{{ $t('loans.my.detail.ext_promo.title') }}</h3>
          <p class="text-xs font-medium text-blue-200/80 leading-relaxed">
            {{ $t('loans.my.detail.ext_promo.desc') }}
          </p>
        </div>
        <button @click="showExtensionModal = true" class="w-full py-5 bg-white text-[#1E3A5F] rounded-2xl text-xs font-black uppercase tracking-[0.2em] shadow-xl hover:bg-blue-50 transition-all">
          {{ $t('loans.my.detail.ext_promo.btn') }}
        </button>
      </div>
    </div>
    
    <!-- Request Extension Modal -->
    <Transition name="scale">
      <div v-if="showExtensionModal" class="fixed inset-0 z-[100] flex items-center justify-center p-6 bg-[#1E3A5F]/40 backdrop-blur-sm">
        <div class="glass max-w-lg w-full bg-white rounded-lg p-12 shadow-2xl border border-white space-y-10" v-motion-pop>
          <!-- Modal Header -->
          <div class="flex items-start justify-between">
            <div class="space-y-1">
              <h2 class="text-2xl font-black text-[#1E3A5F]">{{ $t('loans.my.modal.title') }}</h2>
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('loans.my.modal.doc_label', { no: selectedLoan?.no }) }}</p>
            </div>
            <button @click="showExtensionModal = false" class="p-2 hover:bg-slate-50 rounded-xl text-slate-300 transition-colors">
              <LucideX class="w-6 h-6" />
            </button>
          </div>

          <div class="space-y-8">
            <!-- Duration Input -->
            <div class="space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest flex items-center gap-2">
                {{ $t('loans.my.modal.days_label') }} <span class="text-red-500">*</span>
              </label>
              <div class="relative group">
                 <input type="number" v-model.number="extensionDays" min="1" max="14"
                       class="w-full pl-6 pr-16 py-5 bg-slate-50 border border-slate-100 rounded-2xl text-sm font-black text-[#1E3A5F] outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all" />
                <span class="absolute right-6 top-1/2 -translate-y-1/2 text-[10px] font-black text-slate-300 uppercase">{{ $t('loans.my.modal.days_suffix') }}</span>
              </div>
              <p class="text-[10px] font-bold text-slate-300">{{ $t('loans.my.modal.days_hint') }}</p>
            </div>

            <!-- Reason Input -->
            <div class="space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest flex items-center gap-2">
                {{ $t('loans.my.modal.reason_label') }} <span class="text-red-500">*</span>
              </label>
              <textarea v-model="extensionReason"
                        :placeholder="$t('loans.my.modal.reason_placeholder')" 
                        rows="4" 
                        class="w-full p-6 bg-slate-50 border border-slate-100 rounded-2xl text-sm font-medium text-slate-600 outline-none focus:ring-4 focus:ring-primary-500/5 focus:border-primary-500 transition-all resize-none"></textarea>
            </div>

            <!-- Info Box -->
            <div class="p-6 bg-blue-50/50 border border-blue-100 rounded-lg flex gap-4 items-start">
              <LucideInfo class="w-5 h-5 text-blue-500 shrink-0 mt-0.5" />
              <p class="text-[11px] font-bold text-slate-500 leading-relaxed">
                {{ $t('loans.my.modal.info_box') }}
              </p>
            </div>
          </div>

          <!-- Footer Actions -->
          <div class="flex items-center justify-end gap-6 pt-4">
            <button @click="showExtensionModal = false" class="text-xs font-black text-slate-400 hover:text-slate-600 uppercase tracking-widest transition-colors">
              {{ $t('loans.my.modal.btn_cancel') }}
            </button>
            <button @click="submitExtension" class="px-8 py-5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-xs font-black uppercase tracking-[0.2em] shadow-xl shadow-blue-900/20 transition-all flex items-center gap-3">
              <LucideSend class="w-4 h-4" /> {{ $t('loans.my.modal.btn_submit') }}
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { 
  LucideHistory, LucideCheck, LucideInfo, LucideAlertTriangle, LucideX, LucideSend 
} from 'lucide-vue-next'
import PageHeader from '~/components/PageHeader.vue'
import { useApi } from '~/composables/useApi'

const activeTab = ref('All')
const tabs = [
  { labelKey: 'loans.my.tabs.all', value: 'All' },
  { labelKey: 'loans.my.tabs.due_soon', value: 'Due Soon' },
  { labelKey: 'loans.my.tabs.overdue', value: 'Overdue' }
]
const showExtensionModal = ref(false)
const submissionSuccess = ref(false)
const extensionDays = ref(7)
const extensionReason = ref('')
const submissionSuccessDays = ref(null)

const loans = ref([])
const selectedLoan = ref(null)
const isLoading = ref(false)

const { $api } = useApi()
const config = useRuntimeConfig()

const fetchLoans = async () => {
  isLoading.value = true
  try {
    const res = await $api(`${config.public.apiBase}/loans/my`)
    if (res && res.data) {
      loans.value = res.data
      if (loans.value.length > 0 && !selectedLoan.value) {
        selectedLoan.value = loans.value[0]
      }
    }
  } catch (err) {
    console.error('Failed to fetch user loans:', err)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  fetchLoans()
})

const filteredLoans = computed(() => {
  if (activeTab.value === 'All') return loans.value
  if (activeTab.value === 'Due Soon') {
    return loans.value.filter(l => {
      if (!l.remainingText) return false
      const match = l.remainingText.match(/(\d+)\s+Days?\s+Left/i)
      if (match) {
        const days = parseInt(match[1])
        return days <= 5
      }
      return false
    })
  }
  if (activeTab.value === 'Overdue') {
    return loans.value.filter(l => l.status.toLowerCase() === 'overdue' || l.rawStatus === 'overdue')
  }
  return loans.value
})

const submitExtension = () => {
  // Update state for demo/placeholder for now
  if (selectedLoan.value) {
    selectedLoan.value.status = 'EXTENSION PENDING'
    selectedLoan.value.statusColor = 'bg-orange-50 text-orange-500 border border-orange-100'
    submissionSuccessDays.value = extensionDays.value
    submissionSuccess.value = true
    showExtensionModal.value = false
    extensionReason.value = ''
    
    // Auto hide success banner after 5s
    setTimeout(() => {
      submissionSuccess.value = false
    }, 5000)
  }
}
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
</style>
