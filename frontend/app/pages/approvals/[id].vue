<template>
  <div class="max-w-7xl mx-auto space-y-8 pb-20">
    <!-- Header Navigation -->
    <div class="flex items-center justify-between" v-motion-fade>
      <button 
        @click="navigateTo('/dashboard')" 
        class="flex items-center gap-2 text-xs font-black text-slate-400 hover:text-primary-600 transition-colors uppercase tracking-widest group"
      >
        <LucideArrowLeft class="w-4 h-4 group-hover:-translate-x-1 transition-transform" />
        {{ $t('approvals.detail.back_to_queue') || 'Back to Queue' }}
      </button>
    </div>

    <!-- Title & Status Badges -->
    <div class="flex flex-wrap items-end justify-between gap-6" v-motion-fade>
      <div class="space-y-1">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
          {{ $t('approvals.detail.title') || 'Pre-Registration Review' }}
        </h1>
        <p class="text-sm font-bold text-slate-400 uppercase tracking-widest">REG-{{ doc?.id?.substring(0, 8).toUpperCase() || '2026-0234' }}</p>
      </div>
      <div class="flex items-center gap-3">
        <span class="px-4 py-2 bg-orange-50 text-orange-600 border border-orange-100 rounded-full text-[10px] font-black uppercase tracking-widest shadow-sm">
          {{ $t('approvals.detail.status.pending') || 'Pending Your Approval' }}
        </span>
        <span v-if="doc?.urgency === 'urgent' || true" class="px-4 py-2 bg-red-50 text-red-600 border border-red-100 rounded-full text-[10px] font-black uppercase tracking-widest shadow-sm">
          {{ $t('approvals.detail.status.urgent') || 'Urgent' }}
        </span>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
      <!-- Left Column: Details -->
      <div class="lg:col-span-2 space-y-8">
        <!-- Submission Details Card -->
        <div class="glass rounded-3xl p-10 space-y-10" v-motion-fade>
          <h3 class="text-xs font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('approvals.detail.sections.submission') || 'Submission Details' }}</h3>
          
          <div class="grid grid-cols-2 gap-12">
            <!-- Row 1 -->
            <div class="space-y-4">
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('approvals.detail.fields.submitted_by') || 'Submitted By' }}</p>
              <div class="flex items-center gap-4">
                <div class="w-10 h-10 rounded-full bg-slate-100 overflow-hidden ring-2 ring-white shadow-md">
                  <img :src="`https://i.pravatar.cc/100?u=${doc?.owner_name}`" class="w-full h-full object-cover" />
                </div>
                <p class="text-sm font-black text-[#1E3A5F]">{{ doc?.owner_name || 'Budi Qartono' }}</p>
              </div>
            </div>
            <div class="space-y-4">
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('approvals.detail.fields.doc_info') || 'Document Information' }}</p>
              <p class="text-sm font-black text-[#1E3A5F]">{{ doc?.title || 'Contract, Vendor Agreement' }}</p>
            </div>

            <!-- Row 2 -->
            <div class="space-y-4">
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('approvals.detail.fields.category') || 'Category' }}</p>
              <span class="px-3 py-1 bg-blue-50 text-blue-600 rounded-lg text-[10px] font-black uppercase tracking-widest border border-blue-100">
                {{ doc?.type_name || 'Legal Category' }}
              </span>
            </div>
            <div class="space-y-4">
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('approvals.detail.fields.department') || 'Department' }}</p>
              <p class="text-sm font-black text-[#1E3A5F]">{{ doc?.department_name || 'Procurement Dept' }}</p>
            </div>

            <!-- Row 3 -->
            <div class="space-y-4">
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('approvals.detail.fields.date') || 'Submission Date' }}</p>
              <p class="text-sm font-black text-[#1E3A5F]">{{ new Date(doc?.created_at).toLocaleDateString() || '15 Mar 2026' }}</p>
            </div>
            <div class="space-y-4">
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('approvals.detail.fields.count') || 'Document Count' }}</p>
              <p class="text-sm font-black text-[#1E3A5F]">{{ doc?.metadata?.page_count || 3 }} documents</p>
            </div>
          </div>

          <!-- Notes Box -->
          <div class="bg-amber-50/50 border-l-4 border-amber-400 p-6 rounded-r-2xl space-y-2">
            <p class="text-[10px] font-black text-amber-600 uppercase tracking-widest">{{ $t('approvals.detail.fields.notes') || 'Notes' }}</p>
            <p class="text-xs font-bold text-amber-800/80 leading-relaxed italic">
              {{ doc?.description || 'Please review the renewal terms on section 4.2 specifically. The vendor requested a 5% price escalation clause which has been pre-approved by finance.' }}
            </p>
          </div>

          <!-- Bulk Upload Info -->
          <div class="bg-slate-50 border border-slate-100 p-6 rounded-3xl flex items-center justify-between group cursor-pointer hover:bg-white transition-all shadow-sm">
            <div class="flex items-center gap-4">
              <div class="w-12 h-12 bg-white rounded-2xl flex items-center justify-center shadow-sm text-slate-400 group-hover:text-primary-600 transition-colors">
                <LucideLayers class="w-6 h-6" />
              </div>
              <div>
                <p class="text-xs font-black text-slate-700 uppercase tracking-tight">Part of Bulk Upload</p>
                <p class="text-[10px] text-slate-400 font-bold mt-0.5 uppercase tracking-widest">This submission belongs to Batch #BT-992</p>
              </div>
            </div>
            <button class="text-[10px] font-black text-primary-600 uppercase tracking-widest flex items-center gap-2 group-hover:gap-3 transition-all">
              View All 42 Documents in This Batch
              <LucideArrowRight class="w-3.5 h-3.5" />
            </button>
          </div>
        </div>

        <!-- Attachments List -->
        <div class="glass rounded-3xl p-10 space-y-8" v-motion-fade>
          <div class="flex items-center justify-between">
            <h3 class="text-xs font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('approvals.detail.sections.attachments') || 'Attachments' }} (3)</h3>
            <button class="text-[10px] font-black text-primary-600 uppercase tracking-widest hover:underline decoration-2 underline-offset-4">Download All (.zip)</button>
          </div>

          <div class="space-y-4">
            <div v-for="i in 2" :key="i" class="flex items-center justify-between p-6 bg-white border border-slate-100 rounded-2xl hover:border-primary-200 hover:shadow-lg transition-all group">
              <div class="flex items-center gap-5">
                <div class="w-12 h-12 bg-red-50 rounded-2xl flex items-center justify-center text-red-500 shadow-sm">
                  <LucideFileText class="w-6 h-6" />
                </div>
                <div>
                  <p class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">{{ i === 1 ? 'Vendor_Agreement_Draft.pdf' : 'Pricing_Annex_Q1.xlsx' }}</p>
                  <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest mt-1">1.2 MB • Updated 2 hours ago</p>
                </div>
              </div>
              <div class="flex items-center gap-3">
                <button class="px-4 py-2 bg-slate-50 text-slate-600 rounded-lg text-[10px] font-black uppercase tracking-widest hover:bg-slate-100 transition-all border border-slate-100">Preview</button>
                <button class="p-2 text-slate-300 hover:text-slate-600 transition-colors"><LucideDownload class="w-5 h-5" /></button>
              </div>
            </div>
          </div>
        </div>

        <!-- Manager Notes Input -->
        <div class="glass rounded-3xl p-10 space-y-6" v-motion-fade>
          <h3 class="text-xs font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('approvals.detail.sections.manager_notes') || 'Internal Manager Notes (Private)' }}</h3>
          <textarea 
            v-model="managerNotes"
            rows="4"
            class="w-full bg-slate-50 border border-slate-100 rounded-2xl p-6 text-sm font-medium focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500 transition-all outline-none"
            placeholder="Type notes for your records or other managers..."
          ></textarea>
        </div>
      </div>

      <!-- Right Column: Workflow -->
      <div class="space-y-8">
        <div class="glass rounded-3xl p-10 space-y-10" v-motion-fade>
          <h3 class="text-xs font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('approvals.detail.sections.workflow') || 'Workflow Status' }}</h3>
          
          <div class="relative pl-10 space-y-12">
            <!-- Vertical Line -->
            <div class="absolute left-[1.15rem] top-2 bottom-2 w-0.5 bg-slate-100"></div>

            <!-- Steps -->
            <div class="relative">
              <div class="absolute -left-[1.65rem] top-0 w-6 h-6 rounded-full bg-green-500 flex items-center justify-center text-white ring-4 ring-white shadow-sm z-10">
                <LucideCheck class="w-3.5 h-3.5" />
              </div>
              <div class="space-y-1">
                <p class="text-sm font-black text-[#1E3A5F]">{{ $t('approvals.detail.workflow.submitted') || 'Submitted' }}</p>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest leading-relaxed">
                  By {{ doc?.owner_name || 'Budi Qartono' }} • 15 Mar 2026, 09:12 AM
                </p>
              </div>
            </div>

            <div class="relative">
              <div class="absolute -left-[1.65rem] top-0 w-6 h-6 rounded-full bg-orange-500 flex items-center justify-center text-white ring-4 ring-white shadow-sm z-10 animate-pulse">
                <LucideClock class="w-3.5 h-3.5" />
              </div>
              <div class="space-y-1">
                <p class="text-sm font-black text-orange-600">{{ $t('approvals.detail.workflow.pending_manager') || 'Pending Manager Approval' }}</p>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest leading-relaxed">
                  Assigned to John Doe (You)
                </p>
              </div>
            </div>

            <div class="relative opacity-30">
              <div class="absolute -left-[1.65rem] top-0 w-6 h-6 rounded-full bg-slate-200 flex items-center justify-center text-white ring-4 ring-white shadow-sm z-10">
                <span class="text-[10px] font-black">3</span>
              </div>
              <div class="space-y-1">
                <p class="text-sm font-black text-slate-800">{{ $t('approvals.detail.workflow.compliance') || 'Legal Compliance Review' }}</p>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">Not started</p>
              </div>
            </div>

            <div class="relative opacity-30">
              <div class="absolute -left-[1.65rem] top-0 w-6 h-6 rounded-full bg-slate-200 flex items-center justify-center text-white ring-4 ring-white shadow-sm z-10">
                <span class="text-[10px] font-black">4</span>
              </div>
              <div class="space-y-1">
                <p class="text-sm font-black text-slate-800">{{ $t('approvals.detail.workflow.archiving') || 'Final Archiving' }}</p>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">Pending approval chain</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Bottom Actions Bar -->
    <div class="fixed bottom-0 left-0 right-0 bg-white/80 backdrop-blur-xl border-t border-slate-200 p-6 z-[100]" v-motion-slide-visible-bottom>
      <div class="max-w-7xl mx-auto flex items-center justify-between">
        <button 
          @click="navigateTo('/dashboard')"
          class="px-8 py-3 bg-white border border-slate-200 text-slate-600 rounded-xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all shadow-sm"
        >
          Back to Queue
        </button>
        <div class="flex items-center gap-4">
          <button 
            @click="rejectApproval"
            class="px-8 py-3 bg-white border border-red-200 text-red-600 rounded-xl text-xs font-black uppercase tracking-widest hover:bg-red-50 transition-all shadow-sm"
          >
            Reject
          </button>
          <button 
            @click="approveDocument"
            class="px-12 py-3 bg-green-500 hover:bg-green-600 text-white rounded-xl text-xs font-black uppercase tracking-widest flex items-center gap-3 shadow-xl shadow-green-500/30 transition-all group"
          >
            <LucideThumbsUp class="w-4 h-4 group-hover:scale-125 transition-transform" />
            Approve
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideArrowLeft, LucideLayers, LucideArrowRight, LucideFileText, 
  LucideDownload, LucideCheck, LucideClock, LucideThumbsUp 
} from 'lucide-vue-next'
import { useApi } from '~/composables/useApi'

const route = useRoute()
const { $api } = useApi()
const managerNotes = ref('')

// Fetch Data
const { data: docRes } = await useAsyncData(`approval-${route.params.id}`, () => 
  $api(`/documents/${route.params.id}`)
)
const doc = computed(() => docRes.value?.data)

const approveDocument = async () => {
  try {
    const res = await $api.post(`/documents/${route.params.id}/approve`, {
      notes: managerNotes.value
    })
    if (res.status === 'success') {
      alert('Document Approved Successfully!')
      navigateTo('/dashboard')
    }
  } catch (e) {
    console.error(e)
    alert('Failed to approve document')
  }
}

const rejectApproval = async () => {
  if (!managerNotes.value) {
    alert('Please provide rejection notes before rejecting.')
    return
  }
  
  if (confirm('Are you sure you want to reject this submission?')) {
    try {
      const res = await $api.post(`/documents/${route.params.id}/reject`, {
        notes: managerNotes.value
      })
      if (res.status === 'success') {
        alert('Document Rejected')
        navigateTo('/dashboard')
      }
    } catch (e) {
      console.error(e)
      alert('Failed to reject document')
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
</style>
