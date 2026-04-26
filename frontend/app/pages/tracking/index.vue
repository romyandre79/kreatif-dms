<template>
  <div class="max-w-5xl mx-auto space-y-8 pb-20">
    <!-- Success/Reject Alert Banner -->
    <div v-if="status === 'approved'" class="p-4 bg-emerald-50 border border-emerald-100 rounded-2xl flex items-center gap-4 text-emerald-700 font-bold text-sm shadow-sm" v-motion-slide-top>
      <div class="w-8 h-8 bg-emerald-500 rounded-full flex items-center justify-center text-white shrink-0 shadow-lg shadow-emerald-500/20">
        <LucideCheck class="w-5 h-5" />
      </div>
      <p>{{ $t('loans.tracking.success_banner') }}</p>
    </div>

    <div v-if="status === 'rejected'" class="p-8 bg-red-50 border border-red-100 rounded-3xl flex items-start gap-6 shadow-sm" v-motion-slide-top>
      <div class="w-12 h-12 bg-red-500 rounded-2xl flex items-center justify-center text-white shrink-0 shadow-lg shadow-red-500/20">
        <LucideAlertCircle class="w-6 h-6" />
      </div>
      <div class="space-y-1">
        <h3 class="text-lg font-black text-red-600 tracking-tight">{{ $t('loans.tracking.rejection_banner.title') }}</h3>
        <p class="text-sm font-medium text-red-700/80 leading-relaxed italic">
          {{ $t('loans.tracking.rejection_banner.example') }}
        </p>
      </div>
    </div>

    <!-- Header Card -->
    <div class="glass rounded-[2.5rem] bg-white border border-slate-100 p-10 shadow-xl shadow-slate-200/50 space-y-10" v-motion-fade>
      <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
        <div class="space-y-3">
          <div class="flex items-center gap-4">
            <h1 class="text-3xl font-black text-[#1E3A5F]">{{ $t('loans.tracking.header.request_id', { id: 'BP-2023-001' }) }}</h1>
            <span :class="`px-4 py-1.5 rounded-full text-[10px] font-black uppercase tracking-widest border transition-colors ${
              status === 'approved' ? 'bg-emerald-50 text-emerald-500 border-emerald-100' : 
              status === 'rejected' ? 'bg-red-50 text-red-500 border-red-100' :
              'bg-orange-50 text-orange-500 border-orange-100'}`">
              {{ $t(`loans.tracking.status.${status}`) }}
            </span>
          </div>
          <p class="text-xs font-bold text-slate-400 flex items-center gap-2">
            <LucideCalendar class="w-3.5 h-3.5" />
            {{ $t('loans.tracking.header.submitted_at', { date: 'Oct 24, 2023', time: '10:30 AM' }) }}
          </p>
        </div>
        <div class="flex items-center gap-3">
          <button v-if="status === 'rejected'" class="px-6 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest hover:bg-[#152943] transition-all flex items-center gap-2 shadow-lg shadow-blue-900/10">
            <LucideSend class="w-4 h-4" /> {{ $t('loans.tracking.header.btn_resubmit') }}
          </button>
          <button class="px-6 py-3 bg-slate-50 border border-slate-200 rounded-xl text-xs font-black uppercase tracking-widest text-slate-600 hover:bg-slate-100 transition-all flex items-center gap-2">
            <LucideRotateCcw class="w-4 h-4" /> {{ $t('loans.tracking.header.btn_refresh') }}
          </button>
        </div>
      </div>

      <!-- Metadata Grid -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-10 pt-10 border-t border-slate-50">
        <div class="space-y-2">
          <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ $t('loans.tracking.meta.method') }}</p>
          <div class="flex items-center gap-2 text-sm font-black text-[#1E3A5F]">
            <LucidePenTool class="w-4 h-4 text-primary-500" />
            {{ $t('loans.tracking.meta.digital_sign') }}
          </div>
        </div>
        <div class="space-y-2">
          <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ $t('loans.tracking.meta.purpose') }}</p>
          <p class="text-sm font-black text-[#1E3A5F]">Contract Renewal - Vendor Q4</p>
        </div>
        <div class="space-y-2">
          <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">
            {{ status === 'approved' ? $t('loans.tracking.meta.next_step') : status === 'rejected' ? $t('loans.tracking.meta.last_status') : $t('loans.tracking.meta.eta') }}
          </p>
          <p :class="`text-sm font-black transition-colors ${
            status === 'approved' ? 'text-emerald-600' : 
            status === 'rejected' ? 'text-red-500' :
            'text-[#1E3A5F]'}`">
            {{ status === 'approved' ? $t('loans.tracking.meta.waiting_pickup') : status === 'rejected' ? $t('loans.tracking.meta.rejected_by', { name: 'L2 Head Legal' }) : $t('loans.tracking.meta.eta_value') }}
          </p>
        </div>
      </div>
    </div>

    <!-- Tracking Progress Section -->
    <div class="glass rounded-[2.5rem] bg-white border border-slate-100 shadow-xl shadow-slate-200/50 overflow-hidden" v-motion-slide-visible-bottom>
      <div class="px-10 py-8 border-b border-slate-50">
        <h2 class="text-lg font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('loans.tracking.progress.title') }}</h2>
      </div>

      <div class="p-10 space-y-0 relative">
        <!-- Vertical Line -->
        <div class="absolute left-[3.25rem] top-10 bottom-24 w-0.5 bg-slate-100"></div>

        <!-- Step 1: Completed -->
        <div class="relative flex items-start gap-8 pb-12">
          <div class="relative z-10 w-10 h-10 rounded-full bg-emerald-500 text-white flex items-center justify-center shadow-lg shadow-emerald-500/20">
            <LucideCheck class="w-5 h-5" />
          </div>
          <div class="flex-1 pt-1 space-y-1">
            <div class="flex items-center justify-between">
              <h3 class="text-sm font-black text-emerald-600 uppercase tracking-tight">{{ $t('loans.tracking.progress.steps.step1.title') }}</h3>
              <span class="text-[10px] font-bold text-slate-400">Oct 24, 10:30 AM</span>
            </div>
            <p class="text-xs font-medium text-slate-500 leading-relaxed">
              {{ $t('loans.tracking.progress.steps.step1.desc') }}
            </p>
          </div>
        </div>

        <!-- Step 2: Completed -->
        <div class="relative flex items-start gap-8 pb-12">
          <div class="relative z-10 w-10 h-10 rounded-full bg-emerald-500 text-white flex items-center justify-center shadow-lg shadow-emerald-500/20">
            <LucideCheck class="w-5 h-5" />
          </div>
          <div class="flex-1 pt-1 space-y-1">
            <div class="flex items-center justify-between">
              <h3 class="text-sm font-black text-emerald-600 uppercase tracking-tight">{{ $t('loans.tracking.progress.steps.step2.title') }}</h3>
              <span class="text-[10px] font-bold text-slate-400">Oct 24, 14:15 PM</span>
            </div>
            <p class="text-xs font-medium text-slate-500 leading-relaxed">
              {{ $t('loans.tracking.progress.steps.step2.desc', { name: 'John Doe' }) }}
            </p>
          </div>
        </div>

        <!-- Step 3: L2 Decision -->
        <div class="relative flex items-start gap-8 pb-12">
          <div :class="`relative z-10 w-10 h-10 rounded-full flex items-center justify-center shadow-lg transition-all ${
            status === 'approved' ? 'bg-emerald-500 text-white shadow-emerald-500/20' : 
            status === 'rejected' ? 'bg-red-500 text-white shadow-red-500/20' :
            'bg-[#1E3A5F] text-white shadow-blue-900/20'}`">
            <LucideCheck v-if="status === 'approved'" class="w-5 h-5" />
            <LucideX v-else-if="status === 'rejected'" class="w-5 h-5" />
            <LucideMoreHorizontal v-else class="w-5 h-5" />
          </div>
          <div class="flex-1 pt-1 space-y-4">
            <div class="flex items-center justify-between">
              <h3 :class="`text-sm font-black uppercase tracking-tight transition-colors ${
                status === 'approved' ? 'text-emerald-600' : 
                status === 'rejected' ? 'text-red-600' :
                'text-[#1E3A5F]'}`">
                {{ $t('loans.tracking.progress.steps.step3.title') }}
              </h3>
              <span v-if="status === 'approved'" class="text-[10px] font-bold text-slate-400">Oct 25, 09:45 AM</span>
              <span v-else-if="status === 'rejected'" class="px-2 py-0.5 bg-red-50 text-red-600 rounded-md text-[8px] font-black uppercase tracking-widest border border-red-100">
                {{ $t('loans.tracking.progress.steps.step3.status_rejected') }}
              </span>
              <span v-else class="px-2 py-0.5 bg-blue-50 text-blue-600 rounded-md text-[8px] font-black uppercase tracking-widest border border-blue-100">
                {{ $t('loans.tracking.progress.steps.step3.status_in_progress') }}
              </span>
            </div>
            
            <!-- Detail Card -->
            <div :class="`p-6 border rounded-[2rem] flex gap-4 items-start border-l-4 transition-all ${
              status === 'approved' ? 'bg-emerald-50/30 border-emerald-100 border-l-emerald-500' : 
              status === 'rejected' ? 'bg-red-50/30 border-red-100 border-l-red-500' :
              'bg-slate-50/50 border-slate-100 border-l-primary-500'}`">
              <div class="p-2 bg-white rounded-lg shadow-sm border border-slate-100">
                <LucideShieldCheck v-if="status === 'approved'" class="w-4 h-4 text-emerald-500" />
                <LucideShieldAlert v-else-if="status === 'rejected'" class="w-4 h-4 text-red-500" />
                <LucideKey v-else class="w-4 h-4 text-primary-500" />
              </div>
              <div class="space-y-1">
                <h4 :class="`text-xs font-black uppercase tracking-tight transition-colors ${
                  status === 'approved' ? 'text-emerald-700' : 
                  status === 'rejected' ? 'text-red-700' :
                  'text-primary-700'}`">
                  {{ status === 'approved' ? $t('loans.tracking.progress.steps.step3.card_approved') : status === 'rejected' ? $t('loans.tracking.progress.steps.step3.card_rejected') : $t('loans.tracking.progress.steps.step3.card_pending') }}
                </h4>
                <p class="text-[10px] font-medium text-slate-500 leading-loose">
                  {{ status === 'approved' ? $t('loans.tracking.progress.steps.step3.desc_approved') : 
                     status === 'rejected' ? $t('loans.tracking.progress.steps.step3.desc_rejected', { date: 'Oct 25, 09:45 AM' }) :
                     $t('loans.tracking.progress.steps.step3.desc_pending') }}
                </p>
              </div>
            </div>
          </div>
        </div>

        <!-- Step 4: Admin Preparation -->
        <div class="relative flex items-start gap-8">
          <div :class="`relative z-10 w-10 h-10 rounded-full flex items-center justify-center transition-all ${
            status === 'approved' ? 'bg-[#1E3A5F] text-white shadow-lg shadow-blue-900/20' : 
            'bg-white border-2 border-slate-100 text-slate-300'}`">
            <LucidePackage class="w-5 h-5" />
          </div>
          <div :class="`flex-1 pt-1 space-y-1 transition-all ${status === 'approved' ? '' : 'opacity-50'}`">
            <div class="flex items-center justify-between">
              <h3 class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('loans.tracking.progress.steps.step4.title') }}</h3>
              <span v-if="status === 'approved'" class="px-2 py-0.5 bg-blue-50 text-blue-600 rounded-md text-[8px] font-black uppercase tracking-widest border border-blue-100">
                {{ $t('loans.tracking.progress.steps.step4.status_in_progress') }}
              </span>
            </div>
            <p class="text-xs font-medium text-slate-500 leading-relaxed">
              {{ status === 'approved' ? $t('loans.tracking.progress.steps.step4.desc_approved') : 
                 status === 'rejected' ? $t('loans.tracking.progress.steps.step4.desc_rejected') :
                 $t('loans.tracking.progress.steps.step4.desc_pending') }}
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Rejection History (Empty State when no history, but shown when rejected) -->
    <div v-if="status !== 'rejected'" class="flex flex-col items-center justify-center py-10 opacity-30" v-motion-fade>
      <LucideInfo class="w-8 h-8 text-slate-400 mb-2" />
      <p class="text-xs font-bold text-slate-400 uppercase tracking-widest">{{ $t('loans.tracking.empty_history') }}</p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideCalendar, LucideRotateCcw, LucidePenTool, LucideCheck, 
  LucideMoreHorizontal, LucideKey, LucidePackage, LucideInfo, LucideShieldCheck, LucideAlertCircle, LucideSend, LucideX, LucideShieldAlert
} from 'lucide-vue-next'

const status = ref('waiting_l2') // Options: 'waiting_l1', 'waiting_l2', 'approved', 'rejected'
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
</style>
