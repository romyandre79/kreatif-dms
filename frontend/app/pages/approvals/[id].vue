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
        <p class="text-sm font-bold text-slate-400 uppercase tracking-widest">{{ doc?.reg_no || doc?.id || route.params.id }}</p>
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
          <h3 class="text-xs font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('approvals.detail.sections.submission') }}</h3>
          
          <div class="grid grid-cols-2 gap-12">
            <!-- Row 1 -->
            <div class="space-y-4">
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('approvals.list.table.reg_id') || 'ID REGISTRASI' }}</p>
              <p class="text-sm font-black text-[#1E3A5F] font-mono">{{ doc?.id || doc?.entity_id || '---' }}</p>
            </div>
            <div class="space-y-4">
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('approvals.detail.fields.submitted_by') }}</p>
              <div class="flex items-center gap-4">
                <div class="w-10 h-10 rounded-full bg-slate-100 overflow-hidden ring-2 ring-white shadow-md">
                  <img :src="`https://ui-avatars.com/api/?name=${doc?.owner_name || 'U'}&background=random`" class="w-full h-full object-cover" />
                </div>
                <p class="text-sm font-black text-[#1E3A5F]">{{ doc?.owner_name || '---' }}</p>
              </div>
            </div>

            <!-- Row 2 -->
            <div class="space-y-4 col-span-2">
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('approvals.detail.fields.doc_info') }}</p>
              <p class="text-sm font-black text-[#1E3A5F]">{{ doc?.description || doc?.title || '---' }}</p>
            </div>

            <!-- Row 3 -->
            <div class="space-y-4">
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('approvals.detail.fields.category') }}</p>
              <span class="px-3 py-1 bg-blue-50 text-blue-600 rounded-lg text-[10px] font-black uppercase tracking-widest border border-blue-100">
                {{ doc?.type_name || '---' }}
              </span>
            </div>
            <div class="space-y-4">
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('approvals.detail.fields.department') }}</p>
              <p class="text-sm font-black text-[#1E3A5F]">{{ doc?.department_name || '---' }}</p>
            </div>

            <!-- Row 4 -->
            <div class="space-y-4">
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('approvals.detail.fields.date') }}</p>
              <p class="text-sm font-black text-[#1E3A5F]">{{ doc?.created_at ? new Date(doc.created_at).toLocaleDateString() : '---' }}</p>
            </div>
            <div class="space-y-4">
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('approvals.detail.fields.count') }}</p>
              <p class="text-sm font-black text-[#1E3A5F]">{{ doc?.metadata?.page_count || 1 }} {{ $t('approvals.detail.fields.files') }}</p>
            </div>
          </div>

          <!-- Notes Box -->
          <div v-if="doc?.description" class="bg-amber-50/50 border-l-4 border-amber-400 p-6 rounded-r-2xl space-y-2">
            <p class="text-[10px] font-black text-amber-600 uppercase tracking-widest">{{ $t('approvals.detail.fields.notes') }}</p>
            <p class="text-xs font-bold text-amber-800/80 leading-relaxed italic">
              {{ doc?.description }}
            </p>
          </div>

          <!-- Bulk Upload Info (Conditional) -->
          <div v-if="doc?.batch_id" class="bg-slate-50 border border-slate-100 p-6 rounded-3xl flex items-center justify-between group cursor-pointer hover:bg-white transition-all shadow-sm">
            <div class="flex items-center gap-4">
              <div class="w-12 h-12 bg-white rounded-2xl flex items-center justify-center shadow-sm text-slate-400 group-hover:text-primary-600 transition-colors">
                <LucideLayers class="w-6 h-6" />
              </div>
              <div>
                <p class="text-xs font-black text-slate-700 uppercase tracking-tight">Part of Bulk Upload</p>
                <p class="text-[10px] text-slate-400 font-bold mt-0.5 uppercase tracking-widest">This submission belongs to Batch #{{ doc.batch_id }}</p>
              </div>
            </div>
            <button @click="navigateTo(`/approvals?batch=${doc.batch_id}`)" class="text-[10px] font-black text-primary-600 uppercase tracking-widest flex items-center gap-2 group-hover:gap-3 transition-all">
              View Batch
              <LucideArrowRight class="w-3.5 h-3.5" />
            </button>
          </div>
        </div>

        <!-- Attachments List -->
        <div class="glass rounded-3xl p-10 space-y-8" v-motion-fade>
          <div class="flex items-center justify-between">
            <h3 class="text-xs font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('approvals.detail.sections.attachments') }} ({{ attachedFiles.length }})</h3>
          </div>

          <div class="space-y-4">
            <div v-for="file in attachedFiles" :key="file.id" class="flex items-center justify-between p-6 bg-white border border-slate-100 rounded-2xl hover:border-primary-200 hover:shadow-lg transition-all group">
              <div class="flex items-center gap-5">
                <div class="w-12 h-12 bg-red-50 rounded-2xl flex items-center justify-center text-red-500 shadow-sm">
                  <LucideFileText class="w-6 h-6" />
                </div>
                <div>
                  <p class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">{{ file.filename || file.file_name || 'Document File' }}</p>
                  <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest mt-1">{{ (file.size / 1024).toFixed(1) }} KB</p>
                </div>
              </div>
              <div class="flex items-center gap-3">
                <button @click="previewFile(file)" class="px-4 py-2 bg-slate-50 text-slate-600 rounded-lg text-[10px] font-black uppercase tracking-widest hover:bg-slate-100 transition-all border border-slate-100">Preview</button>
                <a :href="file.url || file.file_path" download class="p-2 text-slate-300 hover:text-slate-600 transition-colors"><LucideDownload class="w-5 h-5" /></a>
              </div>
            </div>
            <div v-if="!attachedFiles.length" class="text-center py-10">
              <LucideFileStack class="w-10 h-10 text-slate-200 mx-auto mb-4" />
              <p class="text-xs font-bold text-slate-300 uppercase tracking-widest">No attachments found</p>
            </div>
          </div>
        </div>

        <!-- Manager Notes Input -->
        <div class="glass rounded-3xl p-10 space-y-6" v-motion-fade>
          <h3 class="text-xs font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('approvals.detail.sections.manager_notes') }}</h3>
          <textarea 
            v-model="managerNotes"
            rows="4"
            class="w-full bg-slate-50 border border-slate-100 rounded-2xl p-6 text-sm font-medium focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500 transition-all outline-none"
            :placeholder="$t('approvals.loans.actions.rejection_placeholder')"
          ></textarea>
        </div>
      </div>

      <!-- Right Column: Workflow -->
      <div class="space-y-8">
        <div class="glass rounded-3xl p-10 space-y-10" v-motion-fade>
          <h3 class="text-xs font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('approvals.detail.sections.workflow') }}</h3>
          
          <div class="relative pl-10 space-y-12">
            <!-- Vertical Line -->
            <div class="absolute left-[1.15rem] top-2 bottom-2 w-0.5 bg-slate-100"></div>

            <!-- Steps -->
            <div class="relative">
              <div class="absolute -left-[1.65rem] top-0 w-6 h-6 rounded-full bg-green-500 flex items-center justify-center text-white ring-4 ring-white shadow-sm z-10">
                <LucideCheck class="w-3.5 h-3.5" />
              </div>
              <div class="space-y-1">
                <p class="text-sm font-black text-[#1E3A5F]">{{ $t('approvals.detail.workflow.submitted') }}</p>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest leading-relaxed">
                  By {{ doc?.owner_name }} • {{ doc?.created_at ? new Date(doc.created_at).toLocaleString() : '---' }}
                </p>
              </div>
            </div>

            <div class="relative">
              <div class="absolute -left-[1.65rem] top-0 w-6 h-6 rounded-full bg-orange-500 flex items-center justify-center text-white ring-4 ring-white shadow-sm z-10 animate-pulse">
                <LucideClock class="w-3.5 h-3.5" />
              </div>
              <div class="space-y-1">
                <p class="text-sm font-black text-orange-600">{{ $t('approvals.detail.workflow.pending_manager') }}</p>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest leading-relaxed">
                  Assigned to You
                </p>
              </div>
            </div>

            <div class="relative opacity-30">
              <div class="absolute -left-[1.65rem] top-0 w-6 h-6 rounded-full bg-slate-200 flex items-center justify-center text-white ring-4 ring-white shadow-sm z-10">
                <span class="text-[10px] font-black">3</span>
              </div>
              <div class="space-y-1">
                <p class="text-sm font-black text-slate-800">{{ $t('approvals.detail.workflow.compliance') }}</p>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">Not started</p>
              </div>
            </div>

            <div class="relative opacity-30">
              <div class="absolute -left-[1.65rem] top-0 w-6 h-6 rounded-full bg-slate-200 flex items-center justify-center text-white ring-4 ring-white shadow-sm z-10">
                <span class="text-[10px] font-black">4</span>
              </div>
              <div class="space-y-1">
                <p class="text-sm font-black text-slate-800">{{ $t('approvals.detail.workflow.archiving') }}</p>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">Pending approval chain</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Bottom Actions Bar -->
    <div class="fixed bottom-0 left-0 lg:left-[280px] right-0 bg-white border-t border-slate-100 p-8 z-[100] shadow-[0_-20px_50px_-12px_rgba(0,0,0,0.05)]" v-motion-slide-visible-bottom>
      <div class="max-w-7xl mx-auto flex items-center justify-between">
        <button 
          @click="navigateTo('/dashboard')"
          class="px-8 py-4 bg-white border border-slate-200 text-slate-600 rounded-2xl text-[10px] font-black uppercase tracking-[0.2em] hover:bg-slate-50 transition-all shadow-sm"
        >
          {{ $t('approvals.detail.back_to_queue') || 'Kembali ke Antrean' }}
        </button>
        <div class="flex items-center gap-6">
          <button 
            @click="showRejectModal = true"
            class="px-10 py-4 bg-white border border-red-200 text-red-600 rounded-2xl text-[10px] font-black uppercase tracking-[0.2em] hover:bg-red-50 transition-all shadow-sm"
          >
            {{ $t('approvals.loans.actions.btn_reject') || 'Tolak Permohonan' }}
          </button>
            <button 
              @click="showApproveModal = true"
              class="px-14 py-4 bg-emerald-500 hover:bg-emerald-600 text-white rounded-2xl text-[10px] font-black uppercase tracking-[0.2em] flex items-center gap-3 shadow-xl shadow-emerald-500/30 transition-all group"
            >
              <LucideThumbsUp class="w-4 h-4 group-hover:scale-125 transition-transform" />
              {{ $t('approvals.loans.actions.btn_approve_l1') || 'Setujui Permohonan' }}
            </button>
          </div>
        </div>
      </div>

      <!-- Reject Submission Modal -->
      <Transition
        enter-active-class="transition duration-300 ease-out"
        enter-from-class="opacity-0"
        enter-to-class="opacity-100"
        leave-active-class="transition duration-200 ease-in"
        leave-from-class="opacity-100"
        leave-to-class="opacity-0"
      >
        <div v-if="showRejectModal" class="fixed inset-0 z-[200] flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm">
          <div 
            class="bg-white dark:bg-slate-900 w-full max-w-2xl rounded-lg shadow-2xl overflow-hidden max-h-[90vh] flex flex-col"
            v-motion-slide-visible-bottom
          >
            <!-- Modal Header -->
            <div class="px-10 py-8 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between">
              <h2 class="text-2xl font-black text-slate-800 dark:text-white uppercase tracking-tight">{{ $t('approvals.modals.reject.title') }}</h2>
              <button @click="showRejectModal = false" class="p-2 text-slate-400 hover:text-slate-600 transition-colors">
                <LucideX class="w-6 h-6" />
              </button>
            </div>

            <!-- Modal Body -->
            <div class="p-10 space-y-10 overflow-y-auto flex-grow custom-scrollbar">
              <!-- Summary Box (Reddish) -->
              <div class="p-8 bg-red-50/50 border border-red-100 rounded-3xl grid grid-cols-2 gap-x-12 gap-y-6">
                <div class="space-y-1">
                  <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">REG ID</p>
                  <p class="text-xs font-black text-slate-800">{{ doc?.reg_no || doc?.id || route.params.id }}</p>
                </div>
                <div class="space-y-1">
                  <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">SUBMITTED BY</p>
                  <p class="text-xs font-bold text-slate-600">{{ doc?.owner_name || '---' }}</p>
                </div>
                <div class="space-y-1">
                  <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">DOCUMENT TYPE</p>
                  <p class="text-xs font-bold text-slate-600 truncate">{{ doc?.type_name || '---' }}</p>
                </div>
                <div class="space-y-1">
                  <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">QUANTITY</p>
                  <p class="text-xs font-bold text-slate-600">{{ attachedFiles.length }} physical documents</p>
                </div>
              </div>

              <!-- Reason Selection -->
              <div class="space-y-6">
                <div class="flex items-center gap-2">
                  <LucideAlertTriangle class="w-4 h-4 text-orange-500" />
                  <h3 class="text-xs font-black text-orange-600 uppercase tracking-widest">{{ $t('approvals.modals.reject.reason_title') }}</h3>
                </div>
                
                <div class="grid grid-cols-2 gap-6">
                  <label v-for="reason in rejectReasons" :key="reason.id" class="flex items-center gap-3 cursor-pointer group">
                    <div class="relative flex items-center justify-center">
                      <input 
                        type="radio" 
                        v-model="rejectionReason" 
                        :value="reason.label"
                        name="rejection_reason"
                        class="peer appearance-none w-5 h-5 border-2 border-slate-200 rounded-full checked:border-red-500 transition-all"
                      />
                      <div class="absolute w-2.5 h-2.5 rounded-full bg-red-500 transform scale-0 peer-checked:scale-100 transition-transform"></div>
                    </div>
                    <span :class="`text-sm font-bold transition-colors ${rejectionReason === reason ? 'text-slate-900' : 'text-slate-500 group-hover:text-slate-700'}`">
                      {{ reason }}
                    </span>
                  </label>
                </div>
              </div>

              <!-- Detailed Explanation -->
              <div class="space-y-4">
                <textarea 
                  v-model="rejectionNotes"
                  rows="4"
                  maxlength="500"
                  class="w-full bg-white border border-slate-200 rounded-2xl p-6 text-sm font-medium focus:ring-4 focus:ring-red-500/10 focus:border-red-500 transition-all outline-none resize-none"
                  :placeholder="$t('approvals.modals.reject.placeholder')"
                ></textarea>
                <div class="flex items-center justify-between">
                  <p class="text-[10px] font-bold text-slate-400">{{ $t('approvals.modals.reject.hint') }}</p>
                  <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ rejectionNotes.length }} / 500</p>
                </div>
              </div>

              <!-- Next Steps Info (Blue) -->
              <div class="p-6 bg-blue-50/50 border border-blue-100 rounded-2xl space-y-4">
                <div class="flex items-center gap-2">
                  <LucideInfo class="w-4 h-4 text-blue-500" />
                  <h3 class="text-[10px] font-black text-blue-800 uppercase tracking-widest">{{ $t('approvals.modals.reject.next_steps') }}</h3>
                </div>
                <ul class="space-y-2">
                  <li class="flex items-center gap-3 text-[10px] font-bold text-slate-500">
                    <div class="w-1 h-1 rounded-full bg-blue-400"></div>
                    {{ $t('approvals.modals.reject.step1') }}
                  </li>
                  <li class="flex items-center gap-3 text-[10px] font-bold text-slate-500">
                    <div class="w-1 h-1 rounded-full bg-blue-400"></div>
                    {{ $t('approvals.modals.reject.step2') }}
                  </li>
                  <li class="flex items-center gap-3 text-[10px] font-bold text-slate-500">
                    <div class="w-1 h-1 rounded-full bg-blue-400"></div>
                    {{ $t('approvals.modals.reject.step3') }}
                  </li>
                </ul>
              </div>
            </div>

            <!-- Modal Footer -->
            <div class="px-10 py-8 bg-slate-50 dark:bg-slate-900/50 border-t border-slate-100 dark:border-slate-800 flex items-center justify-end gap-6">
              <button 
                @click="showRejectModal = false"
                class="text-xs font-black uppercase tracking-widest text-slate-400 hover:text-slate-600 transition-colors"
              >
                {{ $t('approvals.modals.reject.btn_cancel') }}
              </button>
              <button 
                @click="confirmReject"
                :disabled="submitting || !rejectionReason"
                class="px-10 py-4 bg-[#EF4444] hover:bg-[#DC2626] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-red-500/20 transition-all flex items-center gap-3 disabled:opacity-50"
              >
                <LucideX class="w-4 h-4" />
                {{ $t('approvals.modals.reject.btn_confirm') }}
              </button>
            </div>
          </div>
        </div>
      </Transition>

      <!-- Approval Confirmation Modal -->
      <Transition
        enter-active-class="transition duration-300 ease-out"
        enter-from-class="opacity-0"
        enter-to-class="opacity-100"
        leave-active-class="transition duration-200 ease-in"
        leave-from-class="opacity-100"
        leave-to-class="opacity-0"
      >
        <div v-if="showApproveModal" class="fixed inset-0 z-[200] flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm">
          <div 
            class="bg-white dark:bg-slate-900 w-full max-w-lg rounded-lg shadow-2xl overflow-hidden"
            v-motion-slide-visible-bottom
          >
            <!-- Modal Header -->
            <div class="px-8 py-6 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between">
              <h2 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('approvals.modals.approve.title') }}</h2>
              <button @click="showApproveModal = false" class="p-2 text-slate-400 hover:text-slate-600 transition-colors">
                <LucideX class="w-5 h-5" />
              </button>
            </div>

            <!-- Modal Body -->
            <div class="p-8 space-y-8">
              <!-- Summary Box -->
              <div class="p-6 bg-slate-50 dark:bg-slate-950/30 rounded-2xl space-y-4">
                <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('approvals.modals.approve.summary') }}</p>
                <div class="space-y-3">
                  <div class="flex justify-between items-center text-xs">
                    <span class="font-bold text-slate-500">{{ $t('approvals.modals.approve.ref_id') }}</span>
                    <span class="font-black text-[#1E3A5F] dark:text-white font-mono">{{ doc?.reg_no || doc?.id || route.params.id }}</span>
                  </div>
                  <div class="flex justify-between items-center text-xs">
                    <span class="font-bold text-slate-500">{{ $t('approvals.modals.approve.type') }}</span>
                    <span class="font-black text-[#1E3A5F] dark:text-white">{{ doc?.type_name || '---' }}</span>
                  </div>
                </div>
              </div>

              <!-- Comments Field -->
              <div class="space-y-3">
                <label class="text-[10px] font-black text-[#1E3A5F] dark:text-slate-300 uppercase tracking-widest">{{ $t('approvals.modals.approve.comments') }}</label>
                <textarea 
                  v-model="managerNotes"
                  rows="4"
                  class="w-full bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-2xl p-5 text-sm font-medium focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all outline-none resize-none"
                  :placeholder="$t('approvals.modals.approve.placeholder')"
                ></textarea>
              </div>

              <!-- Info Box -->
              <div class="p-4 bg-blue-50/50 dark:bg-blue-900/10 border border-blue-100 dark:border-blue-800/30 rounded-xl flex items-start gap-3">
                <LucideInfo class="w-4 h-4 text-blue-500 mt-0.5 flex-shrink-0" />
                <p class="text-[10px] font-bold text-slate-500 dark:text-slate-400 leading-relaxed uppercase tracking-widest">
                  {{ $t('approvals.modals.approve.info') }}
                </p>
              </div>
            </div>

            <!-- Modal Footer -->
            <div class="px-8 py-6 bg-slate-50 dark:bg-slate-900/50 border-t border-slate-100 dark:border-slate-800 flex items-center gap-4">
              <button 
                @click="showApproveModal = false"
                class="flex-1 py-4 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-[10px] font-black uppercase tracking-widest text-slate-500 hover:text-[#1E3A5F] transition-all"
              >
                {{ $t('approvals.modals.approve.btn_cancel') }}
              </button>
              <button 
                @click="confirmApprove"
                :disabled="submitting"
                class="flex-1 py-4 bg-[#10B981] hover:bg-[#059669] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-lg shadow-emerald-500/20 transition-all disabled:opacity-50"
              >
                {{ submitting ? $t('approvals.modals.approve.processing') : $t('approvals.modals.approve.btn_confirm') }}
              </button>
            </div>
          </div>
        </div>
      </Transition>
    </div>
  </template>

<script setup>
const { t } = useI18n()
const route = useRoute()
const { $api } = useApi()

const rejectReasons = computed(() => [
  { id: 'incomplete', label: t('approvals.modals.reject.reasons.incomplete') },
  { id: 'unclear', label: t('approvals.modals.reject.reasons.unclear') },
  { id: 'wrong_dept', label: t('approvals.modals.reject.reasons.wrong_dept') },
  { id: 'invalid', label: t('approvals.modals.reject.reasons.invalid') },
  { id: 'other', label: t('approvals.modals.reject.reasons.other') }
])

import { ref, computed } from 'vue'
import { 
  LucideArrowLeft, LucideLayers, LucideArrowRight, LucideFileText, 
  LucideDownload, LucideCheck, LucideClock, LucideThumbsUp, LucideX,
  LucideInfo, LucideFileStack
} from 'lucide-vue-next'

const managerNotes = ref('')
const showApproveModal = ref(false)
const submitting = ref(false)

// Fetch Data
const { data: docRes } = await useAsyncData(`approval-${route.params.id}`, () => 
  $api(`/documents/${route.params.id}`)
)
const doc = computed(() => docRes.value?.data)

// Robust file detection
const attachedFiles = computed(() => {
  if (!doc.value) return []
  
  const files = []
  const nested = doc.value.files || doc.value.attachments || doc.value.metadata?.files || doc.value.metadata?.attachments || []
  files.push(...nested)
  
  if (doc.value.file_name || doc.value.filename) {
    const exists = files.some(f => (f.filename || f.file_name) === (doc.value.filename || doc.value.file_name))
    if (!exists) {
      files.push({
        id: doc.value.id,
        filename: doc.value.filename || doc.value.file_name,
        file_name: doc.value.filename || doc.value.file_name,
        size: doc.value.file_size || doc.value.size || 0,
        url: doc.value.url || doc.value.file_path || `${useRuntimeConfig().public.apiBase}/documents/${doc.value.id}/preview?token=${useAuthStore().accessToken}`
      })
    }
  }
  return files
})

const previewFile = (file) => {
  if (file.id) navigateTo(`/documents/${file.id}`)
  else alert('Preview not available for this file')
}

const confirmApprove = async () => {
  submitting.value = true
  try {
    const res = await $api(`/documents/${route.params.id}/approve`, {
      method: 'POST',
      body: { notes: managerNotes.value }
    })
    if (res.status === 'success' || res.success) {
      alert('Document Approved Successfully!')
      showApproveModal.value = false
      navigateTo('/dashboard')
    }
  } catch (e) {
    console.error(e)
    alert('Failed to approve document')
  } finally {
    submitting.value = false
  }
}

const showRejectModal = ref(false)
const rejectionReason = ref('')
const rejectionNotes = ref('')

const confirmReject = async () => {
  if (!rejectionReason.value) {
    alert('Please select a rejection reason')
    return
  }
  
  submitting.value = true
  try {
    const res = await $api(`/documents/${route.params.id}/reject`, {
      method: 'POST',
      body: { 
        reason: rejectionReason.value,
        notes: rejectionNotes.value 
      }
    })
    if (res.status === 'success' || res.success) {
      alert('Document Rejected Successfully')
      showRejectModal.value = false
      navigateTo('/dashboard')
    }
  } catch (e) {
    console.error(e)
    alert('Failed to reject document')
  } finally {
    submitting.value = false
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
