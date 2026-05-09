<template>
  <div class="max-w-7xl mx-auto space-y-10 pb-20">
    <!-- Header -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6" v-motion-fade>
      <div class="space-y-2">
        <h1 class="text-3xl font-black text-[#1E3A5F] tracking-tight uppercase">
          {{ $t('approvals.loans.title', { level }) }}
        </h1>
        <p class="text-slate-500 font-medium">
          {{ level === 'L1' ? $t('approvals.loans.subtitle_l1') : $t('approvals.loans.subtitle_l2') }}
        </p>
      </div>
      <div class="flex items-center gap-3">
        <div v-if="level === 'L2'" class="flex items-center gap-2 px-4 py-2 bg-emerald-50 text-emerald-600 rounded-full border border-emerald-100 text-[10px] font-black uppercase tracking-widest">
          <LucideShieldCheck class="w-3.5 h-3.5" />
          {{ $t('approvals.loans.l1_validated') }}
        </div>
        <button class="px-6 py-3 bg-white border border-slate-200 rounded-xl text-xs font-black uppercase tracking-widest text-slate-600 hover:bg-slate-50 transition-all flex items-center gap-2">
          <LucideFilter class="w-4 h-4" /> {{ $t('approvals.loans.filter') }}
        </button>
        <button class="px-6 py-3 bg-white border border-slate-200 rounded-xl text-xs font-black uppercase tracking-widest text-slate-600 hover:bg-slate-50 transition-all flex items-center gap-2">
          <LucideDownload class="w-4 h-4" /> {{ $t('approvals.loans.export') }}
        </button>
      </div>
    </div>

    <!-- Stat Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6" v-motion-slide-visible-bottom>
      <!-- Stat 1 -->
      <div class="glass p-8 rounded-[2rem] border border-slate-100 bg-white shadow-xl shadow-slate-200/50 relative overflow-hidden group hover:shadow-2xl hover:shadow-primary-900/10 transition-all" :class="level === 'L2' ? 'border-l-4 border-primary-500' : ''">
        <div class="space-y-4">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">
            {{ level === 'L1' ? $t('approvals.loans.stats.l1_waiting') : $t('approvals.loans.stats.l2_waiting') }}
          </p>
          <div class="flex items-end justify-between">
            <span class="text-4xl font-black text-[#1E3A5F]">{{ level === 'L1' ? '24' : '18' }}</span>
            <div class="p-3 bg-blue-50 text-blue-500 rounded-2xl group-hover:scale-110 transition-transform">
              <LucideClock v-if="level === 'L1'" class="w-6 h-6" />
              <LucideShieldCheck v-else class="w-6 h-6" />
            </div>
          </div>
        </div>
      </div>

      <!-- Stat 2 -->
      <div class="glass p-8 rounded-[2rem] border border-slate-100 bg-white shadow-xl shadow-slate-200/50 relative overflow-hidden group hover:shadow-2xl hover:shadow-primary-900/10 transition-all">
        <div class="space-y-4">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">
            {{ level === 'L1' ? $t('approvals.loans.stats.l1_approved') : $t('approvals.loans.stats.l2_release') }}
          </p>
          <div class="flex items-end justify-between">
            <span class="text-4xl font-black text-emerald-500">{{ level === 'L1' ? '12' : '8' }}</span>
            <div class="p-3 bg-emerald-50 text-emerald-500 rounded-2xl group-hover:scale-110 transition-transform">
              <LucideCheckCircle2 v-if="level === 'L1'" class="w-6 h-6" />
              <LucideBox v-else class="w-6 h-6" />
            </div>
          </div>
        </div>
      </div>

      <!-- Stat 3 -->
      <div class="glass p-8 rounded-[2rem] border border-slate-100 bg-white shadow-xl shadow-slate-200/50 relative overflow-hidden group hover:shadow-2xl hover:shadow-primary-900/10 transition-all">
        <div class="space-y-4">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">
            {{ level === 'L1' ? $t('approvals.loans.stats.l1_rejected') : $t('approvals.loans.stats.l2_pickup') }}
          </p>
          <div class="flex items-end justify-between">
            <span class="text-4xl font-black" :class="level === 'L1' ? 'text-red-500' : 'text-orange-500'">{{ level === 'L1' ? '3' : '12' }}</span>
            <div :class="`p-3 rounded-2xl group-hover:scale-110 transition-transform ${level === 'L1' ? 'bg-red-50 text-red-500' : 'bg-orange-50 text-orange-500'}`">
              <LucideXCircle v-if="level === 'L1'" class="w-6 h-6" />
              <LucideHandshake v-else class="w-6 h-6" />
            </div>
          </div>
        </div>
      </div>

      <!-- Stat 4 -->
      <div class="glass p-8 rounded-[2rem] border border-slate-100 bg-white shadow-xl shadow-slate-200/50 relative overflow-hidden group hover:shadow-2xl hover:shadow-primary-900/10 transition-all border-l-4" :class="level === 'L1' ? 'border-orange-500' : 'border-slate-200'">
        <div class="space-y-4">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">
            {{ level === 'L1' ? $t('approvals.loans.stats.l1_sla') : $t('approvals.loans.stats.l2_overdue') }}
          </p>
          <div class="flex items-end justify-between">
            <span class="text-4xl font-black" :class="level === 'L1' ? 'text-orange-500' : 'text-red-500'">{{ level === 'L1' ? '5' : '2' }}</span>
            <div :class="`p-3 rounded-2xl group-hover:scale-110 transition-transform ${level === 'L1' ? 'bg-orange-50 text-orange-500' : 'bg-red-50 text-red-500'}`">
              <LucideAlertTriangle v-if="level === 'L1'" class="w-6 h-6" />
              <LucideAlertCircle v-else class="w-6 h-6" />
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10 items-start">
      <!-- Request Queue Table -->
      <div class="lg:col-span-8 space-y-6" v-motion-slide-visible-bottom>
        <div class="flex items-center justify-between px-4">
          <h2 class="text-lg font-black text-[#1E3A5F] uppercase tracking-tight">
            {{ level === 'L1' ? $t('approvals.loans.queue.title_l1') : $t('approvals.loans.queue.title_l2') }}
          </h2>
          <div v-if="level === 'L2'" class="flex items-center gap-2 text-[10px] font-bold text-slate-400">
            <span class="w-2 h-2 rounded-full bg-emerald-500"></span>
            {{ $t('approvals.loans.queue.all_l1_approved') }}
          </div>
        </div>
        
        <div class="glass rounded-lg overflow-hidden border border-slate-100 bg-white shadow-xl shadow-slate-200/50">
          <table class="w-full text-left border-collapse">
            <thead>
              <tr class="bg-slate-50/50 text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100">
                <th class="px-8 py-5">{{ $t('approvals.loans.queue.table.req_no') }}</th>
                <th class="px-8 py-5">{{ $t('approvals.loans.queue.table.requestor') }}</th>
                <th class="px-8 py-5">{{ $t('approvals.loans.queue.table.dept') }}</th>
                <th v-if="level === 'L1'" class="px-8 py-5">{{ $t('approvals.loans.queue.table.method') }}</th>
                <th v-if="level === 'L2'" class="px-8 py-5">{{ $t('approvals.loans.queue.table.approver_l1') }}</th>
                <th class="px-8 py-5">{{ level === 'L1' ? $t('approvals.loans.queue.table.purpose') : $t('approvals.loans.queue.table.security') }}</th>
                <th class="px-8 py-5 text-center">{{ $t('approvals.loans.queue.table.docs') }}</th>
                <th class="px-8 py-5">{{ level === 'L1' ? $t('approvals.loans.queue.table.duration') : $t('approvals.loans.queue.table.priority') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50">
              <tr v-for="req in queue" :key="req.no" 
                  @click="selectedRequest = req"
                  :class="`group hover:bg-slate-50/50 transition-all cursor-pointer ${selectedRequest?.no === req.no ? 'bg-primary-50/30' : ''}`">
                <td class="px-8 py-7">
                  <p :class="`text-sm font-black uppercase tracking-tight ${selectedRequest?.no === req.no ? 'text-primary-600' : 'text-[#1E3A5F]'}`">{{ req.no }}</p>
                </td>
                <td class="px-8 py-7"><p class="text-sm font-bold text-slate-700">{{ req.requestor }}</p></td>
                <td class="px-8 py-7"><p class="text-xs font-bold text-slate-500 uppercase">{{ req.dept }}</p></td>
                
                <td v-if="level === 'L1'" class="px-8 py-7">
                  <span :class="`px-3 py-1 rounded-lg text-[9px] font-black uppercase tracking-widest ${req.method === 'PHYSICAL' ? 'bg-blue-50 text-blue-600 border border-blue-100' : 'bg-emerald-50 text-emerald-600 border border-emerald-100'}`">
                    {{ req.method }}
                  </span>
                </td>
                <td v-if="level === 'L2'" class="px-8 py-7">
                  <div class="flex items-center gap-2">
                    <img :src="req.l1Avatar" class="w-5 h-5 rounded-full" />
                    <p class="text-xs font-bold text-slate-600">{{ req.l1Approver }}</p>
                  </div>
                </td>

                <td class="px-8 py-7">
                  <template v-if="level === 'L1'">
                    <p class="text-xs font-bold text-slate-600 line-clamp-1">{{ req.purpose }}</p>
                  </template>
                  <template v-else>
                    <span :class="`px-3 py-1 rounded-lg text-[9px] font-black uppercase tracking-widest ${req.security === 'HIGH' ? 'bg-red-50 text-red-600 border border-red-100' : 'bg-blue-50 text-blue-600 border border-blue-100'}`">
                      {{ req.security }}
                    </span>
                  </template>
                </td>
                
                <td class="px-8 py-7 text-center"><span class="text-sm font-black text-slate-400">{{ req.docsCount }}</span></td>
                <td class="px-8 py-7">
                  <p v-if="level === 'L1'" class="text-xs font-bold text-slate-600">{{ req.duration }}</p>
                  <div v-else class="flex items-center gap-2">
                    <div :class="`w-2 h-2 rounded-full ${req.priority === 'High' ? 'bg-orange-500' : 'bg-slate-300'}`"></div>
                    <p class="text-xs font-bold text-slate-600">{{ req.priority }}</p>
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
          <div v-if="selectedRequest" :key="selectedRequest.no" class="glass rounded-lg bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 flex flex-col h-full overflow-hidden">
            <!-- Summary Header -->
            <div class="px-10 py-8 border-b border-slate-50 bg-slate-50/30 flex items-center justify-between">
              <div class="flex items-center gap-3">
                <LucideEye v-if="level === 'L1'" class="w-5 h-5 text-slate-400" />
                <LucideShieldCheck v-else class="w-5 h-5 text-emerald-500" />
                <h3 class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">
                  {{ level === 'L1' ? $t('approvals.loans.summary.title_l1') : $t('approvals.loans.summary.title_l2') }}
                </h3>
              </div>
              <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ selectedRequest.no }}</span>
            </div>

            <div class="flex-1 overflow-y-auto p-10 space-y-10 custom-scrollbar">
              <!-- Validation Status (L2 Only) -->
              <div v-if="level === 'L2'" class="p-6 bg-emerald-50 border border-emerald-100 rounded-3xl space-y-4" v-motion-slide-bottom>
                <div class="flex items-center justify-between">
                  <span class="text-[10px] font-black text-emerald-600 uppercase tracking-widest">{{ $t('approvals.loans.summary.l1_approval_status') }}</span>
                  <LucideCheckCircle2 class="w-4 h-4 text-emerald-500" />
                </div>
                <div class="flex items-center gap-4">
                  <img :src="selectedRequest.l1Avatar" class="w-10 h-10 rounded-xl shadow-sm" />
                  <div>
                    <p class="text-xs font-black text-emerald-700">{{ selectedRequest.l1Approver }}</p>
                    <p class="text-[9px] font-bold text-emerald-600/60 uppercase tracking-widest">{{ selectedRequest.l1Date }}</p>
                  </div>
                </div>
              </div>

              <!-- User Profile (L1 Only) -->
              <div v-if="level === 'L1'" class="flex items-center gap-5 p-6 bg-slate-50/50 rounded-3xl border border-slate-100">
                <div class="w-16 h-16 rounded-2xl overflow-hidden shadow-lg shadow-slate-200/50 border-2 border-white flex-shrink-0">
                  <img :src="selectedRequest.avatar" class="w-full h-full object-cover" />
                </div>
                <div>
                  <h4 class="text-lg font-black text-[#1E3A5F] leading-tight">{{ selectedRequest.requestor }}</h4>
                  <p class="text-xs font-bold text-slate-400 uppercase tracking-tight mt-1">{{ selectedRequest.role }}</p>
                </div>
              </div>

              <!-- Purpose -->
              <div class="space-y-4">
                <p class="text-[10px] font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('approvals.loans.summary.purpose_title') }}</p>
                <div class="relative p-6 bg-slate-50 border border-slate-100 rounded-3xl">
                  <span class="absolute -top-3 left-6 text-4xl text-slate-200 font-serif">“</span>
                  <p class="text-sm font-medium text-slate-600 leading-relaxed italic">{{ selectedRequest.purposeDesc }}</p>
                </div>
              </div>

              <!-- Physical Storage Validation (L2 Only) -->
              <div v-if="level === 'L2'" class="space-y-4" v-motion-slide-bottom>
                <p class="text-[10px] font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('approvals.loans.summary.storage_validation') }}</p>
                <div class="p-6 bg-slate-50 border border-slate-100 rounded-3xl space-y-6">
                  <div v-for="doc in selectedRequest.documents" :key="doc.name" class="flex items-start gap-4">
                    <div class="w-10 h-10 bg-white rounded-xl flex items-center justify-center shadow-sm border border-slate-100 flex-shrink-0">
                      <LucideFileText class="w-5 h-5 text-slate-400" />
                    </div>
                    <div class="flex-1">
                      <p class="text-[11px] font-black text-slate-700 uppercase tracking-tight">{{ doc.name }}</p>
                      <div class="flex items-center justify-between mt-1">
                        <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ doc.location || 'WH-TEMP-001' }}</p>
                        <span class="text-[9px] font-black text-emerald-600 uppercase">{{ $t('approvals.loans.summary.in_stock') }}</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <!-- Documents List (L1 Only) -->
              <div v-if="level === 'L1'" class="space-y-6">
                <p class="text-[10px] font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('approvals.loans.summary.docs_requested', { count: selectedRequest.docsCount }) }}</p>
                <div class="space-y-3">
                  <div v-for="doc in selectedRequest.documents" :key="doc.name" class="p-5 bg-white border border-slate-100 rounded-2xl hover:border-primary-200 transition-colors flex items-center justify-between group">
                    <div class="flex items-center gap-4">
                      <LucideFileText class="w-5 h-5 text-slate-400 group-hover:text-primary-500 transition-colors" />
                      <div>
                        <p class="text-xs font-black text-[#1E3A5F] group-hover:text-primary-600 transition-colors">{{ doc.name }}</p>
                        <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest mt-0.5">{{ doc.category }}</p>
                      </div>
                    </div>
                    <span :class="`px-2.5 py-1 rounded-lg text-[8px] font-black uppercase tracking-widest ${doc.sensitivityColor}`">{{ doc.sensitivity }}</span>
                  </div>
                </div>
              </div>

              <!-- Security Checklist (L2 Only) -->
              <div v-if="level === 'L2'" class="space-y-6" v-motion-slide-bottom>
                <p class="text-[10px] font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('approvals.loans.summary.security_checklist') }}</p>
                <div class="space-y-4">
                  <label v-for="check in checklist" :key="check" class="flex items-center gap-4 p-4 bg-white border border-slate-100 rounded-2xl cursor-pointer hover:bg-slate-50 transition-colors group">
                    <input type="checkbox" class="w-5 h-5 rounded border-2 border-slate-200 text-primary-600 focus:ring-primary-500/10" />
                    <span class="text-xs font-bold text-slate-600 group-hover:text-slate-900 transition-colors">{{ check }}</span>
                  </label>
                </div>
              </div>

              <!-- Policy Hint -->
              <div class="p-6 bg-blue-50 border border-blue-100 rounded-3xl flex gap-4 items-start">
                <LucideInfo class="w-5 h-5 text-blue-500 shrink-0 mt-0.5" />
                <p class="text-[10px] font-bold text-slate-500 leading-relaxed">
                  {{ level === 'L1' ? $t('approvals.loans.summary.policy_l1') : $t('approvals.loans.summary.policy_l2') }}
                </p>
              </div>
            </div>

            <!-- Actions Footer -->
            <div class="p-10 border-t border-slate-100 space-y-6">
              <Transition name="fade" mode="out-in">
                <!-- Rejection Mode -->
                <div v-if="isRejectMode" class="space-y-6" v-motion-slide-bottom>
                  <div class="space-y-3">
                    <div class="flex items-center justify-between">
                      <label class="text-[10px] font-black text-red-500 uppercase tracking-widest">{{ $t('approvals.loans.actions.rejection_reason') }} <span class="text-red-500">*</span></label>
                      <span class="text-[9px] font-bold text-slate-300 uppercase">{{ $t('approvals.loans.actions.rejection_hint') }}</span>
                    </div>
                    <textarea 
                      v-model="rejectionReason"
                      :placeholder="$t('approvals.loans.actions.rejection_placeholder')" 
                      rows="4" 
                      class="w-full p-6 bg-slate-50 border border-red-100 rounded-3xl text-sm font-bold text-slate-700 outline-none focus:ring-4 focus:ring-red-500/5 focus:border-red-500 transition-all resize-none"
                    ></textarea>
                  </div>
                  <div class="space-y-4">
                    <button @click="isRejectMode = false" class="w-full py-5 bg-white border border-slate-200 text-slate-400 hover:bg-slate-50 rounded-2xl text-xs font-black uppercase tracking-[0.2em] transition-all flex items-center justify-center gap-3">
                      <LucideShieldCheck class="w-5 h-5" />
                      {{ $t('approvals.loans.actions.btn_back_approve') }}
                    </button>
                    <button @click="confirmReject" :disabled="rejectionReason.length < 10" class="w-full py-5 bg-red-500 hover:bg-red-600 text-white rounded-2xl text-xs font-black uppercase tracking-[0.2em] shadow-xl shadow-red-900/20 transition-all flex items-center justify-center gap-3 disabled:opacity-50 disabled:cursor-not-allowed">
                      <LucideXCircle class="w-5 h-5" /> {{ $t('approvals.loans.actions.btn_confirm_reject') }}
                    </button>
                  </div>
                </div>

                <!-- Normal Mode -->
                <div v-else class="space-y-4">
                  <button @click="triggerApprove" class="w-full py-5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-xs font-black uppercase tracking-[0.2em] shadow-xl shadow-blue-900/20 transition-all flex items-center justify-center gap-3 group">
                    <LucideShieldCheck class="w-5 h-5 group-hover:scale-110 transition-transform" />
                    {{ level === 'L1' ? $t('approvals.loans.actions.btn_approve_l1') : $t('approvals.loans.actions.btn_approve_l2') }}
                  </button>
                  <button @click="isRejectMode = true" class="w-full py-5 bg-white border-2 border-red-100 text-red-500 hover:bg-red-50 rounded-2xl text-xs font-black uppercase tracking-[0.2em] transition-all flex items-center justify-center gap-3">
                    <LucideXCircle class="w-5 h-5" /> {{ $t('approvals.loans.actions.btn_reject') }}
                  </button>
                </div>
              </Transition>
              <p class="text-[9px] font-bold text-slate-400 text-center uppercase tracking-widest pt-2">
                {{ level === 'L1' ? $t('approvals.loans.actions.footer_hint_l1') : $t('approvals.loans.actions.footer_hint_l2') }}
              </p>
            </div>
          </div>
          
          <!-- Empty State -->
          <div v-else class="h-full glass rounded-lg bg-white border border-slate-100 flex flex-col items-center justify-center text-center p-10 space-y-6">
            <div class="w-20 h-20 bg-slate-50 rounded-full flex items-center justify-center text-slate-200">
              <LucideFileStack class="w-10 h-10" />
            </div>
            <div class="space-y-2">
              <h3 class="text-sm font-black text-slate-400 uppercase tracking-widest">{{ $t('approvals.loans.summary.empty.title') }}</h3>
              <p class="text-xs font-medium text-slate-300">{{ $t('approvals.loans.summary.empty.desc', { level }) }}</p>
            </div>
          </div>
        </Transition>
      </aside>
    </div>

    <!-- PIN Confirmation Modal -->
    <Transition name="scale">
      <div v-if="showPinModal" class="fixed inset-0 z-[100] flex items-center justify-center p-6 bg-[#1E3A5F]/40 backdrop-blur-sm">
        <div class="glass max-w-md w-full bg-white rounded-[3rem] p-12 shadow-2xl border border-white flex flex-col items-center text-center space-y-8" v-motion-pop>
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
import { ref, watch, computed } from 'vue'
import { 
  LucideFilter, LucideDownload, LucideClock, LucideCheckCircle2, LucideXCircle, LucideAlertTriangle, 
  LucideEye, LucideFileText, LucideInfo, LucideShieldCheck, LucideFileStack, LucideBox, LucideHandshake, LucideAlertCircle
} from 'lucide-vue-next'

// Mock role detection (In real app, use auth store)
const userRole = ref('manager_doc_controller') // Default to Manager (L1)

// Level computed based on role
const level = computed(() => {
  return userRole.value === 'manager_doc_controller' ? 'L1' : 'L2'
})

const checklist = [
  'Document physical integrity verified',
  'Security clearance confirmed for requestor',
  'QR handover labels printed',
  'Pickup slot availability confirmed'
]

const queue = ref([
  {
    no: 'REQ-001',
    requestor: 'John Doe',
    dept: 'Finance',
    role: 'Finance • Senior Auditor',
    avatar: 'https://i.pravatar.cc/150?u=john',
    method: 'PHYSICAL',
    purpose: 'Audit Kepatuhan...',
    purposeDesc: 'Peminjaman dokumen keuangan periode Q3 2023 untuk audit internal kepatuhan operasional kantor cabang. Dokumen diperlukan dalam bentuk fisik.',
    docsCount: 5,
    duration: '3 Days',
    priority: 'High',
    security: 'HIGH',
    l1Approver: 'Sarah Connor',
    l1Avatar: 'https://i.pravatar.cc/150?u=sarah',
    l1Date: '24 Apr 2026, 14:20',
    documents: [
      { name: 'Q3_2023_Statement.pdf', category: 'INTERNAL LEDGER', sensitivity: 'HIGHLY SENSITIVE', sensitivityColor: 'bg-red-50 text-red-500 border border-red-100', location: 'WH-A1-S4-B2' },
      { name: 'Voucher_Expense_List.xlsx', category: 'OPERATIONS', sensitivity: 'RESTRICTED', sensitivityColor: 'bg-orange-50 text-orange-500 border border-orange-100', location: 'WH-B2-S1-B1' }
    ]
  },
  {
    no: 'REQ-002',
    requestor: 'Jane Smith',
    dept: 'Legal',
    role: 'Legal • Counsel',
    avatar: 'https://i.pravatar.cc/150?u=jane',
    method: 'DIGITAL',
    purpose: 'Contract Review',
    purposeDesc: 'Review kontrak kerjasama vendor IT untuk perpanjangan lisensi tahun 2024.',
    docsCount: 2,
    duration: '1 Day',
    priority: 'Normal',
    security: 'NORMAL',
    l1Approver: 'Robert Paulson',
    l1Avatar: 'https://i.pravatar.cc/150?u=robert',
    l1Date: '24 Apr 2026, 15:05',
    documents: [
      { name: 'Vendor_Agreement_V2.pdf', category: 'LEGAL', sensitivity: 'CONFIDENTIAL', sensitivityColor: 'bg-purple-50 text-purple-600 border border-purple-100', location: 'WH-C3-S2-B5' }
    ]
  }
])

const selectedRequest = ref(queue.value[0])
const isRejectMode = ref(false)
const rejectionReason = ref('')
const showPinModal = ref(false)

watch(selectedRequest, () => {
  isRejectMode.value = false
  rejectionReason.value = ''
  showPinModal.value = false
})

const triggerApprove = () => {
  if (level.value === 'L2') {
    showPinModal.value = true
  } else {
    processApprove()
  }
}

const processApprove = () => {
  alert(`Request ${selectedRequest.value.no} approved by ${level.value}`)
  queue.value = queue.value.filter(r => r.no !== selectedRequest.value.no)
  selectedRequest.value = queue.value[0] || null
}

const confirmPin = () => {
  showPinModal.value = false
  processApprove()
}

const confirmReject = () => {
  if (rejectionReason.value.length >= 10) {
    alert(`Request ${selectedRequest.value.no} rejected by ${level.value}. Reason: ${rejectionReason.value}`)
    queue.value = queue.value.filter(r => r.no !== selectedRequest.value.no)
    selectedRequest.value = queue.value[0] || null
    isRejectMode.value = false
    rejectionReason.value = ''
  }
}
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #E2E8F0; border-radius: 10px; }
.fade-enter-active, .fade-leave-active { transition: opacity 0.2s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; }
</style>
