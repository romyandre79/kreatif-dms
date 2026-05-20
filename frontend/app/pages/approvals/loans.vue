<template>
  <div class="max-w-7xl mx-auto space-y-6 pb-12">
    <!-- Header -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-4" v-motion-fade>
      <div class="space-y-1">
        <h1 class="text-2xl font-black text-[#1E3A5F] tracking-tight uppercase">
          {{ $t('approvals.loans.title', { level }) }}
        </h1>
        <p class="text-sm text-slate-500 font-medium">
          {{ level === 'L1' ? $t('approvals.loans.subtitle_l1') : $t('approvals.loans.subtitle_l2') }}
        </p>
      </div>
      <div class="flex items-center gap-2">
        <div v-if="level === 'L2'" class="flex items-center gap-1.5 px-3 py-1.5 bg-emerald-50 text-emerald-600 rounded-full border border-emerald-100 text-[10px] font-black uppercase tracking-widest">
          <LucideShieldCheck class="w-3.5 h-3.5" />
          {{ $t('approvals.loans.l1_validated') }}
        </div>
        <button class="px-4 py-2 bg-white border border-slate-200 rounded-lg text-xs font-black uppercase tracking-widest text-slate-600 hover:bg-slate-50 transition-all flex items-center gap-1.5">
          <LucideFilter class="w-3.5 h-3.5" /> {{ $t('approvals.loans.filter') }}
        </button>
        <button class="px-4 py-2 bg-white border border-slate-200 rounded-lg text-xs font-black uppercase tracking-widest text-slate-600 hover:bg-slate-50 transition-all flex items-center gap-1.5">
          <LucideDownload class="w-3.5 h-3.5" /> {{ $t('approvals.loans.export') }}
        </button>
      </div>
    </div>

    <!-- Stat Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4" v-motion-slide-visible-bottom>
      <!-- Stat 1 -->
      <div class="glass p-5 rounded-lg border border-slate-100 bg-white shadow-lg shadow-slate-200/50 relative overflow-hidden group hover:shadow-xl hover:shadow-primary-900/10 transition-all" :class="level === 'L2' ? 'border-l-4 border-primary-500' : ''">
        <div class="space-y-2">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">
            {{ level === 'L1' ? $t('approvals.loans.stats.l1_waiting') : $t('approvals.loans.stats.l2_waiting') }}
          </p>
          <div class="flex items-end justify-between">
            <span class="text-3xl font-black text-[#1E3A5F]">{{ statsWaiting }}</span>
            <div class="p-2 bg-blue-50 text-blue-500 rounded-xl group-hover:scale-110 transition-transform">
              <LucideClock v-if="level === 'L1'" class="w-5 h-5" />
              <LucideShieldCheck v-else class="w-5 h-5" />
            </div>
          </div>
        </div>
      </div>

      <!-- Stat 2 -->
      <div class="glass p-5 rounded-lg border border-slate-100 bg-white shadow-lg shadow-slate-200/50 relative overflow-hidden group hover:shadow-xl hover:shadow-primary-900/10 transition-all">
        <div class="space-y-2">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">
            {{ level === 'L1' ? $t('approvals.loans.stats.l1_approved') : $t('approvals.loans.stats.l2_release') }}
          </p>
          <div class="flex items-end justify-between">
            <span class="text-3xl font-black text-emerald-500">{{ statsApproved }}</span>
            <div class="p-2 bg-emerald-50 text-emerald-500 rounded-xl group-hover:scale-110 transition-transform">
              <LucideCheckCircle2 v-if="level === 'L1'" class="w-5 h-5" />
              <LucideBox v-else class="w-5 h-5" />
            </div>
          </div>
        </div>
      </div>

      <!-- Stat 3 -->
      <div class="glass p-5 rounded-lg border border-slate-100 bg-white shadow-lg shadow-slate-200/50 relative overflow-hidden group hover:shadow-xl hover:shadow-primary-900/10 transition-all">
        <div class="space-y-2">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">
            {{ level === 'L1' ? $t('approvals.loans.stats.l1_rejected') : $t('approvals.loans.stats.l2_pickup') }}
          </p>
          <div class="flex items-end justify-between">
            <span class="text-3xl font-black" :class="level === 'L1' ? 'text-red-500' : 'text-orange-500'">{{ statsRejected }}</span>
            <div :class="`p-2 rounded-xl group-hover:scale-110 transition-transform ${level === 'L1' ? 'bg-red-50 text-red-500' : 'bg-orange-50 text-orange-500'}`">
              <LucideXCircle v-if="level === 'L1'" class="w-5 h-5" />
              <LucideHandshake v-else class="w-5 h-5" />
            </div>
          </div>
        </div>
      </div>

      <!-- Stat 4 -->
      <div class="glass p-5 rounded-lg border border-slate-100 bg-white shadow-lg shadow-slate-200/50 relative overflow-hidden group hover:shadow-xl hover:shadow-primary-900/10 transition-all border-l-4" :class="level === 'L1' ? 'border-orange-500' : 'border-slate-200'">
        <div class="space-y-2">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">
            {{ level === 'L1' ? $t('approvals.loans.stats.l1_sla') : $t('approvals.loans.stats.l2_overdue') }}
          </p>
          <div class="flex items-end justify-between">
            <span class="text-3xl font-black" :class="level === 'L1' ? 'text-orange-500' : 'text-red-500'">{{ statsSla }}</span>
            <div :class="`p-2 rounded-xl group-hover:scale-110 transition-transform ${level === 'L1' ? 'bg-orange-50 text-orange-500' : 'bg-red-50 text-red-500'}`">
              <LucideAlertTriangle v-if="level === 'L1'" class="w-5 h-5" />
              <LucideAlertCircle v-else class="w-5 h-5" />
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-6 items-start">
      <!-- Request Queue Table -->
      <div class="lg:col-span-8 space-y-4" v-motion-slide-visible-bottom>
        <div class="flex items-center justify-between px-4">
          <h2 class="text-lg font-black text-[#1E3A5F] uppercase tracking-tight">
            {{ level === 'L1' ? $t('approvals.loans.queue.title_l1') : $t('approvals.loans.queue.title_l2') }}
          </h2>
          <div v-if="level === 'L2'" class="flex items-center gap-2 text-[10px] font-bold text-slate-400">
            <span class="w-2 h-2 rounded-full bg-emerald-500"></span>
            {{ $t('approvals.loans.queue.all_l1_approved') }}
          </div>
        </div>
        
        <div class="glass rounded-lg overflow-hidden border border-slate-100 bg-white shadow-lg shadow-slate-200/50">
          <table class="w-full text-left border-collapse">
            <thead>
              <tr class="bg-slate-50/50 text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100">
                <th class="px-5 py-3.5">{{ $t('approvals.loans.queue.table.req_no') }}</th>
                <th class="px-4 py-3.5">{{ $t('approvals.loans.queue.table.requestor') }}</th>
                <th class="px-4 py-3.5">{{ $t('approvals.loans.queue.table.dept') }}</th>
                <th v-if="level === 'L1'" class="px-4 py-3.5">{{ $t('approvals.loans.queue.table.method') }}</th>
                <th v-if="level === 'L2'" class="px-4 py-3.5">{{ $t('approvals.loans.queue.table.approver_l1') }}</th>
                <th class="px-4 py-3.5">{{ level === 'L1' ? $t('approvals.loans.queue.table.purpose') : $t('approvals.loans.queue.table.security') }}</th>
                <th class="px-4 py-3.5 text-center">{{ $t('approvals.loans.queue.table.docs') }}</th>
                <th class="px-5 py-3.5">{{ level === 'L1' ? $t('approvals.loans.queue.table.duration') : $t('approvals.loans.queue.table.priority') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50">
              <tr v-for="req in queue" :key="req.no" 
                  @click="selectRequest(req)"
                  class="group hover:bg-slate-50/50 transition-all cursor-pointer"
                  :class="selectedRequest?.no === req.no ? 'bg-slate-50/70 border-l-4 border-[#1E3A5F]' : ''">
                <td class="px-5 py-4">
                  <p class="text-xs font-black uppercase tracking-tight text-[#1E3A5F] group-hover:text-primary-600 transition-colors">{{ req.no }}</p>
                </td>
                 <td class="px-4 py-4"><p class="text-xs font-bold text-slate-700">{{ req.userName }}</p></td>
                 <td class="px-4 py-4"><p class="text-[10px] font-bold text-slate-500 uppercase">{{ req.departmentName }}</p></td>
                 
                 <td v-if="level === 'L1'" class="px-4 py-4">
                   <div class="flex items-center gap-2">
                     <span :class="`px-2.5 py-1 rounded-md text-[9px] font-black uppercase tracking-widest ${req.method === 'DIGITAL' ? 'bg-emerald-50 text-emerald-600 border border-emerald-100' : 'bg-blue-50 text-blue-600 border border-blue-100'}`">
                       {{ req.method || 'PHYSICAL' }}
                     </span>
                     <span v-if="req.rawStatus === 'returned'" class="px-2 py-0.5 bg-red-100 text-red-600 border border-red-200 rounded-md text-[8px] font-black uppercase tracking-widest">
                       Returned
                     </span>
                   </div>
                 </td>
                 <td v-if="level === 'L2'" class="px-4 py-4">
                   <div class="flex items-center gap-1.5">
                     <div class="w-4 h-4 rounded-full bg-slate-200 flex items-center justify-center text-[9px] font-bold text-slate-600">L1</div>
                     <p class="text-[10px] font-bold text-slate-600">{{ req.l1Approver || 'System L1' }}</p>
                   </div>
                 </td>
 
                 <td class="px-4 py-4">
                   <template v-if="level === 'L1'">
                     <p class="text-[10px] font-bold text-slate-600 line-clamp-1">{{ req.purpose }}</p>
                   </template>
                   <template v-else>
                     <span :class="`px-2.5 py-1 rounded-md text-[9px] font-black uppercase tracking-widest ${req.security === 'HIGH' ? 'bg-red-50 text-red-600 border border-red-100' : 'bg-blue-50 text-blue-600 border border-blue-100'}`">
                       {{ req.security || 'NORMAL' }}
                     </span>
                   </template>
                 </td>
                 
                 <td class="px-4 py-4 text-center"><span class="text-xs font-black text-slate-400">{{ req.docsCount }}</span></td>
                 <td class="px-5 py-4">
                   <div class="flex items-center justify-between gap-3">
                     <p v-if="level === 'L1'" class="text-[10px] font-bold text-slate-600">{{ req.durationDays }} Hari</p>
                     <div v-else class="flex items-center gap-1.5">
                       <div :class="`w-1.5 h-1.5 rounded-full ${req.priority === 'High' ? 'bg-orange-500' : 'bg-slate-300'}`"></div>
                       <p class="text-[10px] font-bold text-slate-600">{{ req.priority || 'Medium' }}</p>
                     </div>
                     <button class="px-3 py-1 bg-[#1E3A5F] text-white rounded-md text-[9px] font-black uppercase tracking-widest opacity-0 group-hover:opacity-100 transition-all">
                       Review
                     </button>
                   </div>
                 </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Request Summary Sidebar -->
      <aside class="lg:col-span-4" v-motion-slide-visible-right>
        <Transition name="fade" mode="out-in">
          <div v-if="selectedRequest" :key="selectedRequest.no" class="glass rounded-lg bg-white border border-slate-100 shadow-xl shadow-slate-200/50 flex flex-col h-full overflow-hidden">
            <!-- Summary Header -->
            <div class="px-6 py-5 border-b border-slate-50 bg-slate-50/30 flex items-center justify-between">
              <div class="flex items-center gap-2.5">
                <LucideEye v-if="level === 'L1'" class="w-4 h-4 text-slate-400" />
                <LucideShieldCheck v-else class="w-4 h-4 text-emerald-500" />
                <h3 class="text-xs font-black text-[#1E3A5F] uppercase tracking-tight">
                  {{ level === 'L1' ? $t('approvals.loans.summary.title_l1') : $t('approvals.loans.summary.title_l2') }}
                </h3>
              </div>
              <span class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ selectedRequest.no }}</span>
            </div>

            <div class="flex-1 overflow-y-auto p-6 space-y-6 custom-scrollbar">
              <!-- Validation Status (L2 Only) -->
              <div v-if="level === 'L2'" class="p-4 bg-emerald-50 border border-emerald-100 rounded-2xl space-y-3" v-motion-slide-bottom>
                <div class="flex items-center justify-between">
                  <span class="text-[9px] font-black text-emerald-600 uppercase tracking-widest">{{ $t('approvals.loans.summary.l1_approval_status') }}</span>
                  <LucideCheckCircle2 class="w-3.5 h-3.5 text-emerald-500" />
                </div>
                <div class="flex items-center gap-3">
                  <img :src="selectedRequest.l1Avatar" class="w-8 h-8 rounded-lg shadow-sm" />
                  <div>
                    <p class="text-[11px] font-black text-emerald-700">{{ selectedRequest.l1Approver }}</p>
                    <p class="text-[8px] font-bold text-emerald-600/60 uppercase tracking-widest">{{ selectedRequest.l1Date }}</p>
                  </div>
                </div>
              </div>

              <!-- User Profile (L1 Only) -->
              <div v-if="level === 'L1'" class="flex items-center gap-4 p-4 bg-slate-50/50 rounded-2xl border border-slate-100">
                <div class="w-12 h-12 rounded-xl overflow-hidden shadow-md shadow-slate-200/50 border-2 border-white flex-shrink-0">
                  <img :src="selectedRequest.avatar" class="w-full h-full object-cover" />
                </div>
                <div>
                  <h4 class="text-sm font-black text-[#1E3A5F] leading-tight">{{ selectedRequest.requestor }}</h4>
                  <p class="text-[10px] font-bold text-slate-400 uppercase tracking-tight mt-0.5">{{ selectedRequest.role }}</p>
                </div>
              </div>

              <!-- L2 Rejection Reason (L1 Only) -->
              <div v-if="level === 'L1' && selectedRequest.l2RejectionReason" class="p-4 bg-red-50 border border-red-100 rounded-2xl space-y-2 mb-4" v-motion-slide-bottom>
                <div class="flex items-center gap-2">
                  <LucideAlertCircle class="w-4 h-4 text-red-500" />
                  <span class="text-[9px] font-black text-red-600 uppercase tracking-widest">{{ $t('approvals.loans.summary.returned_by_l2') || 'Dikembalikan oleh Kepala DC' }}</span>
                </div>
                <p class="text-[11px] font-medium text-red-700 leading-relaxed italic">"{{ selectedRequest.l2RejectionReason }}"</p>
              </div>

              <!-- Purpose -->
              <div class="space-y-2.5">
                <p class="text-[9px] font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('approvals.loans.summary.purpose_title') }}</p>
                <div class="relative p-4 bg-slate-50 border border-slate-100 rounded-2xl">
                  <span class="absolute -top-2 left-4 text-3xl text-slate-200 font-serif">“</span>
                  <p class="text-[11px] font-medium text-slate-600 leading-relaxed italic">{{ selectedRequest.purposeDesc }}</p>
                </div>
              </div>

              <!-- Physical Storage Validation (L2 Only) -->
              <div v-if="level === 'L2'" class="space-y-3" v-motion-slide-bottom>
                <p class="text-[9px] font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('approvals.loans.summary.storage_validation') }}</p>
                <div class="p-4 bg-slate-50 border border-slate-100 rounded-2xl space-y-4">
                  <div v-for="doc in selectedRequest.documents" :key="doc.name" class="flex items-start gap-3">
                    <div class="w-8 h-8 bg-white rounded-lg flex items-center justify-center shadow-sm border border-slate-100 flex-shrink-0">
                      <LucideFileText class="w-4 h-4 text-slate-400" />
                    </div>
                    <div class="flex-1">
                      <p class="text-[10px] font-black text-slate-700 uppercase tracking-tight">{{ doc.name }}</p>
                      <div class="flex items-center justify-between mt-0.5">
                        <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ doc.location || 'WH-TEMP-001' }}</p>
                        <span class="text-[8px] font-black text-emerald-600 uppercase">{{ $t('approvals.loans.summary.in_stock') }}</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <!-- Documents List (L1 Only) -->
              <div v-if="level === 'L1'" class="space-y-4">
                <p class="text-[9px] font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('approvals.loans.summary.docs_requested', { count: selectedRequest.docsCount }) }}</p>
                <div class="space-y-2">
                  <div v-for="doc in selectedRequest.documents" :key="doc.name" class="p-4 bg-white border border-slate-100 rounded-xl hover:border-primary-200 transition-colors flex items-center justify-between group">
                    <div class="flex items-center gap-3">
                      <LucideFileText class="w-4 h-4 text-slate-400 group-hover:text-primary-500 transition-colors" />
                      <div>
                        <p class="text-[11px] font-black text-[#1E3A5F] group-hover:text-primary-600 transition-colors">{{ doc.name }}</p>
                        <p class="text-[8px] font-bold text-slate-400 uppercase tracking-widest mt-0.5">{{ doc.category }}</p>
                      </div>
                    </div>
                    <span :class="`px-2 py-0.5 rounded-md text-[8px] font-black uppercase tracking-widest ${doc.sensitivityColor}`">{{ doc.sensitivity }}</span>
                  </div>
                </div>
              </div>

              <!-- Security Checklist (L2 Only) -->
              <div v-if="level === 'L2'" class="space-y-4" v-motion-slide-bottom>
                <p class="text-[9px] font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('approvals.loans.summary.security_checklist') }}</p>
                <div class="space-y-2">
                  <label v-for="check in checklist" :key="check" class="flex items-center gap-3 p-3 bg-white border border-slate-100 rounded-xl cursor-pointer hover:bg-slate-50 transition-colors group">
                    <input type="checkbox" class="w-4 h-4 rounded border-2 border-slate-200 text-primary-600 focus:ring-primary-500/10" />
                    <span class="text-[11px] font-bold text-slate-600 group-hover:text-slate-900 transition-colors">{{ check }}</span>
                  </label>
                </div>
              </div>

              <!-- Policy Hint -->
              <div class="p-4 bg-blue-50 border border-blue-100 rounded-2xl flex gap-3 items-start">
                <LucideInfo class="w-4 h-4 text-blue-500 shrink-0 mt-0.5" />
                <p class="text-[9px] font-bold text-slate-500 leading-relaxed">
                  {{ level === 'L1' ? $t('approvals.loans.summary.policy_l1') : $t('approvals.loans.summary.policy_l2') }}
                </p>
              </div>
            </div>

            <!-- Actions Footer -->
            <div class="p-6 border-t border-slate-100 space-y-4">
              <Transition name="fade" mode="out-in">
                <!-- Rejection Mode -->
                <div v-if="isRejectMode" class="space-y-4" v-motion-slide-bottom>
                  <div class="space-y-2">
                    <div class="flex items-center justify-between">
                      <label class="text-[9px] font-black text-red-500 uppercase tracking-widest">{{ $t('approvals.loans.actions.rejection_reason') }} <span class="text-red-500">*</span></label>
                      <span class="text-[8px] font-bold text-slate-300 uppercase">{{ $t('approvals.loans.actions.rejection_hint') }}</span>
                    </div>
                    <textarea 
                      v-model="rejectionReason"
                      :placeholder="$t('approvals.loans.actions.rejection_placeholder')" 
                      rows="3" 
                      class="w-full p-4 bg-slate-50 border border-red-100 rounded-2xl text-xs font-bold text-slate-700 outline-none focus:ring-4 focus:ring-red-500/5 focus:border-red-500 transition-all resize-none"
                    ></textarea>
                  </div>
                  <div class="space-y-2">
                    <button @click="isRejectMode = false" class="w-full py-3.5 bg-white border border-slate-200 text-slate-400 hover:bg-slate-50 rounded-xl text-[10px] font-black uppercase tracking-[0.2em] transition-all flex items-center justify-center gap-2">
                      <LucideShieldCheck class="w-4 h-4" />
                      {{ $t('approvals.loans.actions.btn_back_approve') }}
                    </button>
                    <button @click="confirmReject" :disabled="rejectionReason.length < 10" class="w-full py-3.5 bg-red-500 hover:bg-red-600 text-white rounded-xl text-[10px] font-black uppercase tracking-[0.2em] shadow-lg shadow-red-900/20 transition-all flex items-center justify-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed">
                      <LucideXCircle class="w-4 h-4" /> {{ $t('approvals.loans.actions.btn_confirm_reject') }}
                    </button>
                  </div>
                </div>

                <!-- Normal Mode -->
                <div v-else class="space-y-3">
                  <button @click="triggerApprove" class="w-full py-4 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-xl text-[10px] font-black uppercase tracking-[0.2em] shadow-lg shadow-blue-900/20 transition-all flex items-center justify-center gap-2 group">
                    <LucideShieldCheck class="w-4 h-4 group-hover:scale-110 transition-transform" />
                    {{ level === 'L1' ? $t('approvals.loans.actions.btn_approve_l1') : $t('approvals.loans.actions.btn_approve_l2') }}
                  </button>
                  <button @click="isRejectMode = true" class="w-full py-4 bg-white border-2 border-red-100 text-red-500 hover:bg-red-50 rounded-xl text-[10px] font-black uppercase tracking-[0.2em] transition-all flex items-center justify-center gap-2">
                    <LucideXCircle class="w-4 h-4" /> {{ $t('approvals.loans.actions.btn_reject') }}
                  </button>
                </div>
              </Transition>
              <p class="text-[8px] font-bold text-slate-400 text-center uppercase tracking-widest pt-1">
                {{ level === 'L1' ? $t('approvals.loans.actions.footer_hint_l1') : $t('approvals.loans.actions.footer_hint_l2') }}
              </p>
            </div>
          </div>
          
          <!-- Empty State -->
          <div v-else class="h-full glass rounded-lg bg-white border border-slate-100 flex flex-col items-center justify-center text-center p-8 space-y-4">
            <div class="w-16 h-16 bg-slate-50 rounded-full flex items-center justify-center text-slate-200">
              <LucideFileStack class="w-8 h-8" />
            </div>
            <div class="space-y-1">
              <h3 class="text-xs font-black text-slate-400 uppercase tracking-widest">{{ $t('approvals.loans.summary.empty.title') }}</h3>
              <p class="text-[10px] font-medium text-slate-300">{{ $t('approvals.loans.summary.empty.desc', { level }) }}</p>
            </div>
          </div>
        </Transition>
      </aside>
    </div>

    <!-- PIN Confirmation Modal -->
    <Transition name="scale">
      <div v-if="showPinModal" class="fixed inset-0 z-[100] flex items-center justify-center p-6 bg-[#1E3A5F]/40 backdrop-blur-sm">
        <div class="glass max-w-md w-full bg-white rounded-lg p-12 shadow-2xl border border-white flex flex-col items-center text-center space-y-8" v-motion-pop>
          <!-- Modal Icon -->
          <div class="w-20 h-20 bg-slate-50 rounded-2xl flex items-center justify-center text-[#1E3A5F] border border-slate-100">
            <LucideShieldCheck class="w-10 h-10" />
          </div>

          <!-- Title & Subtitle -->
          <div class="space-y-4">
            <h2 class="text-2xl font-black text-[#1E3A5F]">{{ $t('approvals.loans.modal_pin.title') }}</h2>
            <p class="text-xs font-bold text-slate-400 leading-relaxed px-4">
              {{ $t('approvals.loans.modal_pin.desc') }}
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
              {{ $t('approvals.loans.modal_pin.btn_confirm') }}
            </button>
            <button @click="showPinModal = false" class="text-xs font-black text-slate-400 hover:text-slate-600 uppercase tracking-widest transition-colors">
              {{ $t('approvals.loans.modal_pin.btn_cancel') }}
            </button>
          </div>

          <!-- Footer -->
          <div class="w-full pt-8 border-t border-slate-50">
            <p class="text-[9px] font-bold text-slate-300 uppercase tracking-widest">
              {{ $t('approvals.loans.modal_pin.footer') }}
            </p>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, watch, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '~/stores/auth'
import { useAsyncData } from '#app'
import { useApi } from '~/composables/useApi'
import { 
  LucideFilter, LucideDownload, LucideClock, LucideCheckCircle2, LucideXCircle, LucideAlertTriangle, 
  LucideEye, LucideFileText, LucideInfo, LucideShieldCheck, LucideFileStack, LucideBox, LucideHandshake, LucideAlertCircle
} from 'lucide-vue-next'

const route = useRoute()
const router = useRouter()
const auth = useAuthStore()
const { $api } = useApi()

// Role & Level detection
const userRole = computed(() => (auth.user?.role || '').toLowerCase())
const level = computed(() => {
  if (userRole.value === 'manajer') {
    return 'L1'
  }
  if (userRole.value === 'kepala doc controller' || userRole.value === 'kepala dc') {
    return 'L2'
  }
  return 'NONE'
})

const checklist = [
  'Document physical integrity verified',
  'Security clearance confirmed for requestor',
  'QR handover labels printed',
  'Pickup slot availability confirmed'
]

// Real queue & selectedRequest
const queue = ref([])
const selectedRequest = ref(null)
const isRejectMode = ref(false)
const rejectionReason = ref('')
const showPinModal = ref(false)
const isLoading = ref(false)

// Stat Counters
const statsWaiting = computed(() => {
  if (level.value === 'L2') {
    return queue.value.filter(r => r.rawStatus === 'l1_approved').length
  }
  return queue.value.filter(r => r.rawStatus === 'pending' || r.rawStatus === 'returned').length
})
const statsApproved = computed(() => {
  if (level.value === 'L2') {
    // Siap Diambil / Ready for Release
    return queue.value.filter(r => r.rawStatus === 'l2_approved').length
  }
  return queue.value.filter(r => r.rawStatus === 'l1_approved' || r.rawStatus === 'l2_approved' || r.rawStatus === 'active').length
})
const statsRejected = computed(() => {
  if (level.value === 'L2') {
    // Telah Diambil / Picked Up
    return queue.value.filter(r => r.rawStatus === 'active' || r.rawStatus === 'returned').length
  }
  return queue.value.filter(r => r.rawStatus === 'rejected').length
})
const statsSla = computed(() => {
  return queue.value.filter(r => r.rawStatus === 'overdue').length
})

// Fetch all loans
const fetchLoans = async () => {
  console.log('fetchLoans: Started. Level:', level.value, 'User:', auth.user)
  isLoading.value = true
  try {
    console.log('fetchLoans: Making API request to /loans?limit=100')
    const res = await $api('/loans?limit=100')
    console.log('fetchLoans: Response received:', res)
    if ((res.success || res.status === 'success') && res.data) {
      // Filter list of loans based on approval role/level
      const allLoans = res.data || []
      console.log('fetchLoans: allLoans count:', allLoans.length)
      if (level.value === 'L1') {
        // Manager only sees loans that are pending or returned by L2
        queue.value = allLoans.filter(l => l.rawStatus === 'pending' || l.rawStatus === 'returned')
      } else if (level.value === 'L2') {
        // L2 (Kepala DC) sees loans that are L1 approved or ready for release
        queue.value = allLoans.filter(l => l.rawStatus === 'l1_approved' || l.rawStatus === 'l2_approved' || l.rawStatus === 'active' || l.rawStatus === 'returned' || l.rawStatus === 'overdue')
      } else {
        // Unauthorized roles see empty queue
        queue.value = []
      }
      console.log('fetchLoans: Filtered queue count:', queue.value.length)

      // Check query parameter to auto-select target
      const targetId = route.query.id
      console.log('fetchLoans: targetId from query:', targetId)
      if (targetId) {
        // Search in the UNFILTERED allLoans list so that we can find it even if it's not in the filtered queue!
        const found = allLoans.find(l => l.id === targetId || l.no === targetId)
        console.log('fetchLoans: found in allLoans:', found)
        if (found) {
          // If we found it, temporarily add it to our queue so it displays in the table!
          const inQueue = queue.value.some(l => l.id === found.id)
          if (!inQueue) {
            queue.value.unshift(found)
          }
          selectRequest(found)
        } else {
          // Fallback if not in allLoans, try to fetch directly if targetId looks like UUID
          const isUuid = /^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/i.test(targetId)
          console.log('fetchLoans: fallback targetId. isUuid:', isUuid)
          if (isUuid) {
            selectRequest({ id: targetId })
          } else if (queue.value.length > 0) {
            selectRequest(queue.value[0])
          }
        }
      } else if (queue.value.length > 0) {
        selectRequest(queue.value[0])
      }
    } else {
      console.warn('fetchLoans: res.status !== success or res.data is falsy:', res)
    }
  } catch (err) {
    console.error('fetchLoans: Failed to fetch loans:', err)
  } finally {
    isLoading.value = false
  }
}

// Select a request and load its detailed items
const selectRequest = async (req) => {
  if (!req) {
    selectedRequest.value = null
    return
  }

  isRejectMode.value = false
  rejectionReason.value = ''
  showPinModal.value = false

  // Create deep copy with initial fields
  selectedRequest.value = { ...req, documents: [] }

  try {
    const res = await $api(`/loans/${req.id}`)
    if ((res.success || res.status === 'success') && res.data) {
      const detail = res.data.loan
      const items = res.data.items || []

      selectedRequest.value = {
        id: detail.id,
        no: detail.request_no,
        userName: detail.user_name,
        departmentName: detail.department_name,
        requestor: detail.user_name || 'System User',
        role: detail.department_name || 'Department',
        avatar: `https://i.pravatar.cc/100?u=${detail.user_name || 'unknown'}`,
        l1Avatar: `https://i.pravatar.cc/100?u=manager_${detail.department_name || 'unknown'}`,
        purpose: detail.purpose,
        purposeDesc: detail.purpose || 'No purpose description provided.',
        durationDays: detail.duration_days,
        docsCount: detail.items_count,
        rawStatus: detail.status,
        l1Approver: 'Manager ' + (detail.department_name || ''),
        l1Date: detail.l1_approved_at ? new Date(detail.l1_approved_at).toLocaleString() : '-',
        l2RejectionReason: detail.l2_rejection_reason?.String || '',
        documents: items.map(item => ({
          name: item.document_title || item.document_filename || 'Unnamed Document',
          category: item.category || 'General',
          sensitivity: item.sensitivity || 'NORMAL',
          sensitivityColor: item.sensitivity === 'HIGH' ? 'bg-red-50 text-red-500 border border-red-100' : 'bg-slate-50 text-slate-400 border border-slate-100',
          location: 'WH-TEMP-001'
        }))
      }
    }
  } catch (err) {
    console.error('Failed to fetch loan details:', err)
  }
}

const triggerApprove = () => {
  if (level.value === 'L2') {
    showPinModal.value = true
  } else {
    processApprove()
  }
}

const processApprove = async () => {
  if (!selectedRequest.value) return
  
  try {
    const res = await $api(`/loans/${selectedRequest.value.id}/approve`, {
      method: 'POST'
    })
    if (res.status === 'success' || res.success) {
      alert('Loan request approved successfully!')
      await fetchLoans()
    } else {
      alert('Failed to approve loan request: ' + (res.message || 'Unknown error'))
    }
  } catch (err) {
    console.error('Error approving loan:', err)
    alert('Error approving loan request')
  }
}

const confirmPin = () => {
  showPinModal.value = false
  processApprove()
}

const confirmReject = async () => {
  if (!selectedRequest.value) return
  if (rejectionReason.value.length < 10) {
    alert('Rejection reason must be at least 10 characters long.')
    return
  }

  try {
    const res = await $api(`/loans/${selectedRequest.value.id}/reject`, {
      method: 'POST',
      body: { reason: rejectionReason.value }
    })
    if (res.status === 'success' || res.success) {
      alert('Loan request rejected successfully!')
      isRejectMode.value = false
      rejectionReason.value = ''
      
      // Clear URL to remove the persistent targetId that forces it to stay in the queue
      if (route.query.id) {
        router.push('/approvals/loans')
      } else {
        await fetchLoans()
      }
    } else {
      alert('Failed to reject loan request: ' + (res.message || 'Unknown error'))
    }
  } catch (err) {
    console.error('Error rejecting loan:', err)
    alert('Error rejecting loan request')
  }
}

watch(() => auth.user, (newVal) => {
  console.log('watch(auth.user) triggered:', newVal)
  if (newVal) {
    fetchLoans()
  }
}, { immediate: true })

onMounted(() => {
  console.log('onMounted triggered. auth.user:', auth.user)
  if (auth.user) {
    fetchLoans()
  }
})
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #E2E8F0; border-radius: 10px; }
.fade-enter-active, .fade-leave-active { transition: opacity 0.2s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; }
</style>
