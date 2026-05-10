<template>
  <div class="max-w-7xl mx-auto space-y-8 pb-32">
    <!-- Top Navigation -->
    <div class="flex items-center justify-between" v-motion-fade>
      <div class="space-y-4">
        <button @click="navigateTo('/dashboard/manager/review')" class="flex items-center gap-2 text-xs font-black text-slate-400 hover:text-[#1E3A5F] transition-colors uppercase tracking-widest">
          <LucideArrowLeft class="w-4 h-4" />
          {{ $t('review_detail.back_to_queue') }}
        </button>
        <div class="space-y-1">
          <h1 class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">{{ $t('review_detail.title') }}</h1>
          <p class="text-sm font-black text-slate-400 uppercase tracking-widest">REG-2026-0234</p>
        </div>
      </div>
      <div class="flex items-center gap-3">
        <span class="px-4 py-2 bg-orange-50 dark:bg-orange-900/20 text-orange-600 dark:text-orange-400 rounded-full text-[10px] font-black uppercase tracking-widest border border-orange-100 dark:border-orange-900/50">
          {{ $t('review_detail.status_pending') }}
        </span>
        <span class="px-4 py-2 bg-red-50 dark:bg-red-900/20 text-red-500 rounded-full text-[10px] font-black uppercase tracking-widest border border-red-100 dark:border-red-900/50">
          {{ $t('review_detail.status_urgent') }}
        </span>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8 items-start">
      <!-- Main Column -->
      <div class="lg:col-span-8 space-y-8" v-motion-fade>
        <!-- Submission Details Card -->
        <div class="glass rounded-lg p-10 space-y-10 shadow-xl shadow-slate-200/50 dark:shadow-none">
          <div class="flex items-center gap-3">
            <div class="w-1.5 h-6 bg-[#1E3A5F] rounded-full"></div>
            <h2 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em]">{{ $t('review_detail.submission_details.title') }}</h2>
          </div>

          <div class="grid grid-cols-2 gap-x-12 gap-y-10">
            <div class="space-y-2"><p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('review_detail.submission_details.submitted_by') }}</p><div class="flex items-center gap-3"><div class="w-10 h-10 rounded-full border-2 border-white shadow-sm overflow-hidden bg-slate-100"><img src="https://i.pravatar.cc/150?u=budi" class="w-full h-full object-cover" /></div><p class="text-sm font-black text-[#1E3A5F] dark:text-white">Budi Qartono</p></div></div>
            <div class="space-y-2"><p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('review_detail.submission_details.doc_info') }}</p><p class="text-sm font-black text-slate-700 dark:text-slate-200">Contract, Vendor Agreement</p></div>
            <div class="space-y-2"><p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('review_detail.submission_details.category') }}</p><span class="inline-block px-3 py-1 bg-blue-50 dark:bg-blue-900/20 text-[#1E3A5F] dark:text-blue-400 rounded-lg text-[10px] font-black uppercase tracking-tight">Legal Category</span></div>
            <div class="space-y-2"><p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('review_detail.submission_details.dept') }}</p><p class="text-sm font-bold text-slate-700 dark:text-slate-300">Procurement Dept</p></div>
            <div class="space-y-2"><p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('review_detail.submission_details.date') }}</p><p class="text-sm font-bold text-slate-700 dark:text-slate-300">15 Mar 2026</p></div>
            <div class="space-y-2"><p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('review_detail.submission_details.count') }}</p><p class="text-sm font-bold text-slate-700 dark:text-slate-300">3 documents</p></div>
          </div>

          <div class="p-6 bg-orange-50/50 dark:bg-orange-900/10 border-l-4 border-orange-400 rounded-r-2xl">
            <p class="text-[10px] font-black text-orange-600 uppercase tracking-widest mb-2">{{ $t('review_detail.submission_details.notes_title') }}</p>
            <p class="text-xs font-bold text-slate-600 dark:text-slate-300 leading-relaxed italic">"Please review the renewal terms on section 4.2 specifically. The vendor requested a 5% price escalation clause which has been pre-approved by finance."</p>
          </div>

          <div class="p-6 bg-slate-50 dark:bg-slate-900/50 border border-slate-100 dark:border-slate-800 rounded-3xl flex items-center justify-between">
            <div class="flex items-center gap-4"><div class="w-12 h-12 rounded-2xl bg-white dark:bg-slate-800 flex items-center justify-center shadow-sm"><LucideLayers class="w-6 h-6 text-slate-400" /></div><div><p class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('review_detail.submission_details.bulk_title') }}</p><p class="text-[10px] font-bold text-slate-400 mt-1 uppercase tracking-widest">{{ $t('review_detail.submission_details.bulk_desc', { id: '#BT-992' }) }}</p></div></div>
            <button class="text-[10px] font-black text-[#1E3A5F] dark:text-blue-400 uppercase tracking-widest hover:translate-x-1 transition-transform flex items-center gap-2">{{ $t('review_detail.submission_details.view_batch', { count: 42 }) }}<LucideArrowRight class="w-3 h-3" /></button>
          </div>
        </div>

        <div class="glass rounded-lg p-10 space-y-8 shadow-xl shadow-slate-200/50 dark:shadow-none">
          <div class="flex items-center justify-between"><h2 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em]">{{ $t('review_detail.attachments.title', { count: 3 }) }}</h2><button class="text-[10px] font-black text-[#1E3A5F] dark:text-blue-400 uppercase tracking-widest hover:text-blue-600 transition-colors">{{ $t('review_detail.attachments.download_all') }}</button></div>
          <div class="space-y-4">
            <div v-for="file in files" :key="file.name" class="flex items-center justify-between p-4 bg-slate-50/50 dark:bg-slate-900/30 border border-slate-100 dark:border-slate-800 rounded-2xl group hover:border-primary-500/50 transition-all">
              <div class="flex items-center gap-4"><div :class="`w-12 h-12 rounded-xl flex items-center justify-center ${file.iconBg}`"><component :is="file.icon" class="w-6 h-6" :class="file.iconColor" /></div><div><p class="text-sm font-black text-[#1E3A5F] dark:text-white tracking-tight">{{ file.name }}</p><p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest mt-0.5">{{ file.size }} • {{ file.updated }}</p></div></div>
              <div class="flex items-center gap-2"><button class="px-6 py-2.5 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-[10px] font-black uppercase tracking-widest text-[#1E3A5F] dark:text-white hover:bg-slate-50 transition-all">{{ $t('review_detail.attachments.preview') }}</button><button class="p-2.5 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-slate-400 hover:text-[#1E3A5F] transition-all"><LucideDownload class="w-4 h-4" /></button></div>
            </div>
          </div>
        </div>

        <div class="glass rounded-lg p-10 space-y-6 shadow-xl shadow-slate-200/50 dark:shadow-none"><h2 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em]">{{ $t('review_detail.internal_notes.title') }}</h2><textarea :placeholder="$t('review_detail.internal_notes.placeholder')" rows="4" class="w-full p-6 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-3xl text-sm font-bold outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all resize-none"></textarea></div>
      </div>

      <div class="lg:col-span-4 space-y-8" v-motion-fade>
        <div class="glass rounded-lg p-10 space-y-8 shadow-xl shadow-slate-200/50 dark:shadow-none"><h2 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-[0.2em]">{{ $t('review_detail.workflow.title') }}</h2><div class="space-y-12 relative"><div class="absolute left-[11px] top-4 bottom-4 w-0.5 bg-slate-100 dark:bg-slate-800"></div>
          <div class="flex gap-6 relative"><div class="w-6 h-6 rounded-full bg-green-500 text-white flex items-center justify-center z-10 shadow-lg shadow-green-500/20"><LucideCheck class="w-3.5 h-3.5" /></div><div class="space-y-1"><p class="text-sm font-black text-slate-800 dark:text-white">{{ $t('review_detail.workflow.step1') }}</p><p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest leading-relaxed">{{ $t('review_detail.workflow.step1_desc', { name: 'Budi Qartono', date: '15 Mar 2026, 09:12 AM' }) }}</p></div></div>
          <div class="flex gap-6 relative"><div class="w-6 h-6 rounded-full bg-orange-500 text-white flex items-center justify-center z-10 shadow-lg animate-pulse"><LucideClock class="w-3.5 h-3.5" /></div><div class="space-y-1"><p class="text-sm font-black text-orange-600">{{ $t('review_detail.workflow.step2') }}</p><p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest leading-relaxed">{{ $t('review_detail.workflow.step2_desc', { name: 'John Doe' }) }}</p></div></div>
          <div class="flex gap-6 relative opacity-50"><div class="w-6 h-6 rounded-full bg-white dark:bg-slate-800 border-2 border-slate-100 dark:border-slate-800 text-slate-400 flex items-center justify-center z-10 font-black text-[10px]">3</div><div class="space-y-1"><p class="text-sm font-black text-slate-400">{{ $t('review_detail.workflow.step3') }}</p><p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('review_detail.workflow.step3_not_started') }}</p></div></div>
          <div class="flex gap-6 relative opacity-50"><div class="w-6 h-6 rounded-full bg-white dark:bg-slate-800 border-2 border-slate-100 dark:border-slate-800 text-slate-400 flex items-center justify-center z-10 font-black text-[10px]">4</div><div class="space-y-1"><p class="text-sm font-black text-slate-400">{{ $t('review_detail.workflow.step4') }}</p><p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('review_detail.workflow.step4_pending') }}</p></div></div>
        </div></div>
      </div>
    </div>

    <!-- Sticky Footer -->
    <div class="fixed bottom-0 left-0 lg:left-64 right-0 p-6 glass border-t border-slate-200/50 dark:border-white/5 z-20" v-motion-slide-bottom>
      <div class="max-w-7xl mx-auto flex items-center justify-between">
        <button @click="navigateTo('/dashboard/manager/review')" class="px-8 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-xs font-black uppercase tracking-widest text-slate-600 dark:text-slate-400 hover:bg-slate-50 transition-all">{{ $t('review_detail.back_to_queue') }}</button>
        <div class="flex items-center gap-4">
          <button @click="showRejectModal = true" class="px-10 py-3.5 bg-white dark:bg-slate-900 border border-red-200 dark:border-red-900 text-red-500 rounded-xl text-xs font-black uppercase tracking-widest hover:bg-red-50 transition-all">{{ $t('review_detail.footer.btn_reject') }}</button>
          <button @click="showApprovalModal = true" class="flex items-center gap-3 px-12 py-3.5 bg-green-500 hover:bg-green-600 text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-green-900/20 transition-all group"><LucideThumbsUp class="w-4 h-4 group-hover:scale-110 transition-transform" />{{ $t('review_detail.footer.btn_approve') }}</button>
        </div>
      </div>
    </div>

    <!-- Confirm Approval Modal -->
    <Transition enter-active-class="transition duration-300 ease-out" enter-from-class="opacity-0" enter-to-class="opacity-100"><div v-if="showApprovalModal" class="fixed inset-0 z-[70] flex items-center justify-center p-4 bg-slate-950/60 backdrop-blur-md"><div class="bg-white dark:bg-slate-900 w-full max-w-md rounded-lg shadow-2xl relative overflow-hidden" v-motion-pop><button @click="showApprovalModal = false" class="absolute top-6 right-6 p-2 text-slate-400 hover:text-slate-600 transition-colors"><LucideX class="w-4 h-4" /></button><div class="p-10 space-y-8"><h3 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('review_detail.approval_modal.title') }}</h3><div class="p-6 bg-blue-50/50 dark:bg-blue-900/10 rounded-3xl border border-blue-100/50 dark:border-blue-800/30 space-y-4"><p class="text-[10px] font-black text-blue-500 uppercase tracking-widest">{{ $t('review_detail.approval_modal.summary_title') }}</p><div class="flex justify-between items-center text-sm"><span class="text-slate-400 font-bold">{{ $t('review_detail.approval_modal.ref_id') }}</span><span class="text-[#1E3A5F] dark:text-white font-black">REG-2026-0234</span></div><div class="flex justify-between items-center text-sm"><span class="text-slate-400 font-bold">{{ $t('review_detail.approval_modal.type') }}</span><span class="text-[#1E3A5F] dark:text-white font-black">Contract Renewal</span></div></div><div class="space-y-3"><label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('review_detail.approval_modal.comments_title') }}</label><textarea :placeholder="$t('review_detail.approval_modal.comments_placeholder')" rows="3" class="w-full p-4 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-sm font-bold outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all resize-none"></textarea></div><div class="flex items-start gap-3 p-4 bg-blue-50/30 dark:bg-blue-900/5 rounded-2xl border border-blue-100/30"><LucideInfo class="w-4 h-4 text-blue-400 mt-0.5 flex-shrink-0" /><p class="text-[10px] font-medium text-slate-500 leading-relaxed">{{ $t('review_detail.approval_modal.info_text') }}</p></div><div class="grid grid-cols-2 gap-4"><button @click="showApprovalModal = false" class="py-4 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-2xl text-xs font-black uppercase tracking-widest text-slate-400 hover:bg-slate-50 transition-all">{{ $t('review_detail.approval_modal.btn_cancel') }}</button><button class="py-4 bg-green-500 hover:bg-green-600 text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-lg shadow-green-500/20 transition-all">{{ $t('review_detail.approval_modal.btn_confirm') }}</button></div></div></div></div></Transition>

    <!-- Reject Submission Modal -->
    <Transition enter-active-class="transition duration-300 ease-out" enter-from-class="opacity-0" enter-to-class="opacity-100">
      <div v-if="showRejectModal" class="fixed inset-0 z-[70] flex items-center justify-center p-4 bg-slate-950/60 backdrop-blur-md">
        <div class="bg-white dark:bg-slate-900 w-full max-w-lg rounded-lg shadow-2xl relative overflow-hidden" v-motion-pop>
          <button @click="showRejectModal = false" class="absolute top-6 right-6 p-2 text-slate-400 hover:text-slate-600 transition-colors"><LucideX class="w-4 h-4" /></button>
          <div class="p-10 space-y-8">
            <h3 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('review_detail.reject_modal.title') }}</h3>
            <div class="p-8 bg-red-50/50 dark:bg-red-900/10 rounded-3xl border border-red-100/30 grid grid-cols-2 gap-8">
              <div class="space-y-1"><p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('review_detail.reject_modal.summary.reg_id') }}</p><p class="text-sm font-black text-[#1E3A5F] dark:text-white">REG-2026-0234</p></div>
              <div class="space-y-1"><p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('review_detail.reject_modal.summary.submitted_by') }}</p><p class="text-sm font-bold text-slate-700 dark:text-slate-300">Budi Qartono (QA/E Staff)</p></div>
              <div class="space-y-1"><p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('review_detail.reject_modal.summary.type') }}</p><p class="text-sm font-bold text-slate-700 dark:text-slate-300">Contract - Vendor Agreement with PT. XYZ</p></div>
              <div class="space-y-1"><p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('review_detail.reject_modal.summary.qty') }}</p><p class="text-sm font-bold text-slate-700 dark:text-slate-300">3 physical documents</p></div>
            </div>
            <div class="space-y-4">
              <div class="flex items-center gap-2 text-orange-600"><LucideAlertTriangle class="w-4 h-4" /><p class="text-[10px] font-black uppercase tracking-widest">{{ $t('review_detail.reject_modal.reason_title') }}</p></div>
              <div class="grid grid-cols-2 gap-4">
                <label v-for="(label, key) in $tm('review_detail.reject_modal.rereasons')" :key="key" class="flex items-center gap-3 cursor-pointer group">
                  <div class="relative flex items-center justify-center"><input type="radio" v-model="rejectReason" name="reason" :value="key" class="peer appearance-none w-5 h-5 border-2 border-slate-200 rounded-full checked:border-red-500 transition-all" /><div class="absolute w-2.5 h-2.5 rounded-full bg-red-500 transform scale-0 peer-checked:scale-100 transition-transform"></div></div>
                  <span class="text-xs font-bold text-slate-500 group-hover:text-slate-700 transition-colors">{{ $rt(label) }}</span>
                </label>
                <!-- Manual loop since tm/rt is acting up with nested objects sometimes -->
                <label v-for="reason in ['incomplete', 'duplicate', 'policy', 'category', 'approvals', 'other']" :key="reason" class="flex items-center gap-3 cursor-pointer group">
                  <div class="relative flex items-center justify-center"><input type="radio" v-model="rejectReason" name="reason" :value="reason" class="peer appearance-none w-5 h-5 border-2 border-slate-200 rounded-full checked:border-red-500 transition-all" /><div class="absolute w-2.5 h-2.5 rounded-full bg-red-500 transform scale-0 peer-checked:scale-100 transition-transform"></div></div>
                  <span class="text-xs font-bold text-slate-500 group-hover:text-slate-700 transition-colors">{{ $t(`review_detail.reject_modal.reasons.${reason}`) }}</span>
                </label>
              </div>
            </div>
            <div class="space-y-2"><textarea :placeholder="$t('review_detail.reject_modal.placeholder')" rows="3" class="w-full p-4 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-sm font-bold outline-none focus:border-red-500 transition-all resize-none"></textarea><p class="text-[10px] font-medium text-slate-400 leading-relaxed">{{ $t('review_detail.reject_modal.hint') }}</p></div>
            <div class="p-6 bg-blue-50/30 dark:bg-blue-900/5 rounded-3xl border border-blue-100/30 space-y-3"><div class="flex items-center gap-2 text-[#1E3A5F] dark:text-blue-400"><LucideInfo class="w-4 h-4" /><p class="text-[10px] font-black uppercase tracking-widest">{{ $t('review_detail.reject_modal.next_title') }}</p></div><ul class="space-y-1.5"><li v-for="step in $tm('review_detail.reject_modal.next_steps')" :key="step" class="text-[10px] font-medium text-slate-500 leading-relaxed flex gap-2"><span class="text-blue-400">•</span> {{ $rt(step) }}</li></ul></div>
            <div class="flex items-center justify-between pt-4"><button @click="showRejectModal = false" class="text-xs font-black uppercase tracking-widest text-slate-400 hover:text-slate-600 transition-colors">{{ $t('review_detail.reject_modal.btn_cancel') }}</button><button class="px-10 py-4 bg-red-500 hover:bg-red-600 text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-lg shadow-red-500/20 transition-all flex items-center gap-3"><LucideXCircle class="w-4 h-4" />{{ $t('review_detail.reject_modal.btn_confirm') }}</button></div>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideArrowLeft, LucideLayers, LucideArrowRight, LucideFileText, LucideDownload, LucideFileSpreadsheet,
  LucideCheck, LucideClock, LucideThumbsUp, LucideX, LucideInfo, LucideAlertTriangle, LucideXCircle
} from 'lucide-vue-next'

const showApprovalModal = ref(false)
const showRejectModal = ref(false)
const rejectReason = ref('other')

const files = [
  { name: 'Vendor_Agreement_Draft.pdf', size: '1.2 MB', updated: 'Updated 2 hours ago', icon: LucideFileText, iconColor: 'text-red-500', iconBg: 'bg-red-50 dark:bg-red-900/20' },
  { name: 'Pricing_Annex_Q1.xlsx', size: '450 KB', updated: 'Updated 2 hours ago', icon: LucideFileSpreadsheet, iconColor: 'text-green-500', iconBg: 'bg-green-50 dark:bg-green-900/20' }
]
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
.dark .glass { background: rgba(15, 23, 42, 0.6); border: 1px solid rgba(255, 255, 255, 0.05); }
</style>
