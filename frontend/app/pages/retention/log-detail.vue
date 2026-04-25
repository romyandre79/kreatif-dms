<template>
  <div class="max-w-[1400px] mx-auto space-y-8 pb-20">
    <!-- Header -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6" v-motion-fade>
      <div class="space-y-2">
        <h1 class="text-3xl font-black text-[#1E3A5F] tracking-tight uppercase">{{ $t('retention.log_detail.title') }}</h1>
        <p class="text-slate-500 font-medium text-xs font-bold uppercase tracking-widest">{{ $t('retention.log_detail.subtitle') }}</p>
      </div>
      <span class="px-4 py-2 bg-emerald-50 text-emerald-600 border border-emerald-100 rounded-xl text-[10px] font-black uppercase tracking-widest flex items-center gap-2">
        <div class="w-2 h-2 bg-emerald-500 rounded-full"></div> {{ $t('retention.log_detail.status_success') }}
      </span>
    </div>

    <!-- Immutable Banner -->
    <div class="p-4 bg-slate-100/80 border border-slate-200 rounded-2xl flex items-center gap-3" v-motion-slide-top>
      <LucideLock class="w-4 h-4 text-slate-400" />
      <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('retention.log_detail.immutable_note') }}</p>
    </div>

    <div class="grid grid-cols-12 gap-10 items-start">
      <!-- Left Sidebar: Document Snapshot & Evidence -->
      <div class="col-span-4 space-y-8" v-motion-slide-left>
        <!-- Document Snapshot -->
        <div class="glass p-10 rounded-[3rem] bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-10">
          <div class="flex items-center gap-3 text-[10px] font-black text-slate-400 uppercase tracking-widest">
            <LucideFileText class="w-4 h-4 text-primary-500" />
            {{ $t('retention.log_detail.snapshot.title') }}
          </div>

          <div class="space-y-8">
            <div class="space-y-1">
              <p class="text-[9px] font-black text-slate-300 uppercase tracking-widest">{{ $t('retention.log_detail.snapshot.doc_title') }}</p>
              <p class="text-lg font-black text-[#1E3A5F] leading-tight">Laporan Keuangan Q4 2023 - Final</p>
            </div>
            <div class="space-y-1">
              <p class="text-[9px] font-black text-slate-300 uppercase tracking-widest">{{ $t('retention.log_detail.snapshot.doc_id') }}</p>
              <p class="text-sm font-black text-slate-400 tracking-tight uppercase">DOC-2023-DISP-88291-K</p>
            </div>
            <div class="space-y-1">
              <p class="text-[9px] font-black text-slate-300 uppercase tracking-widest">{{ $t('retention.log_detail.snapshot.owner') }}</p>
              <p class="text-xs font-bold text-slate-500">Divisi Keuangan - Finance Ops</p>
            </div>
            <div class="space-y-2">
              <p class="text-[9px] font-black text-slate-300 uppercase tracking-widest">{{ $t('retention.log_detail.snapshot.class') }}</p>
              <span class="px-3 py-1 bg-red-50 text-red-500 border border-red-100 rounded-lg text-[9px] font-black uppercase tracking-widest">Highly Confidential</span>
            </div>
          </div>
        </div>

        <!-- Disposal Evidence -->
        <div class="glass p-10 rounded-[3rem] bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-8">
          <div class="flex items-center gap-3 text-[10px] font-black text-slate-400 uppercase tracking-widest">
            <LucidePaperclip class="w-4 h-4 text-primary-500" />
            {{ $t('retention.log_detail.evidence.title') }}
          </div>

          <div class="p-6 bg-slate-50 border border-slate-100 rounded-3xl flex items-center justify-between group cursor-pointer hover:bg-primary-50 transition-all">
            <div class="flex items-center gap-4">
              <div class="w-10 h-10 bg-white rounded-xl flex items-center justify-center text-red-500 shadow-sm">
                <LucideFileText class="w-5 h-5" />
              </div>
              <div class="space-y-0.5">
                <p class="text-[10px] font-black text-[#1E3A5F] uppercase tracking-tight truncate max-w-[120px]">BAST_Dispos...</p>
                <p class="text-[8px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('retention.log_detail.evidence.verified') }}</p>
              </div>
            </div>
            <LucideEye class="w-4 h-4 text-slate-200 group-hover:text-primary-500 transition-colors" />
          </div>
        </div>
      </div>

      <!-- Main Section: Disposal Audit Proof Timeline -->
      <div class="col-span-8 glass p-12 rounded-[3.5rem] bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-12" v-motion-fade>
        <div class="flex items-center gap-3 text-[10px] font-black text-slate-400 uppercase tracking-widest">
          <LucideListChecks class="w-4 h-4 text-primary-500" />
          {{ $t('retention.log_detail.timeline.audit_proof') }}
        </div>

        <div class="relative space-y-0 px-20">
          <div class="absolute left-1/2 top-0 bottom-12 w-0.5 bg-slate-50 -translate-x-1/2"></div>

          <!-- Timeline Items -->
          <div class="space-y-24">
            <!-- 1. Alert (Right) -->
            <div class="relative flex justify-end">
              <div class="absolute left-1/2 -translate-x-1/2 w-10 h-10 rounded-full bg-slate-50 border border-slate-100 flex items-center justify-center text-slate-300 z-10">
                <LucideBell class="w-4 h-4" />
              </div>
              <div class="w-[45%] glass p-8 rounded-[2rem] border border-slate-50 shadow-xl shadow-slate-100/50 space-y-3" v-motion-slide-right>
                <div class="flex items-center justify-between">
                  <h4 class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight leading-tight">{{ $t('retention.log_detail.timeline.trigger.title') }}</h4>
                  <p class="text-[10px] font-black text-slate-300">12/01/2024</p>
                </div>
                <p class="text-[11px] font-bold text-slate-400 leading-relaxed uppercase">{{ $t('retention.log_detail.timeline.trigger.desc') }}</p>
              </div>
            </div>

            <!-- 2. Metadata (Left) -->
            <div class="relative flex justify-start">
              <div class="absolute left-1/2 -translate-x-1/2 w-10 h-10 rounded-full bg-slate-50 border border-slate-100 flex items-center justify-center text-slate-300 z-10">
                <LucideFileSpreadsheet class="w-4 h-4" />
              </div>
              <div class="w-[45%] glass p-8 rounded-[2rem] border border-slate-50 shadow-xl shadow-slate-100/50 space-y-3" v-motion-slide-left>
                <div class="flex items-center justify-between">
                  <h4 class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight leading-tight">{{ $t('retention.log_detail.timeline.export.title') }}</h4>
                  <p class="text-[10px] font-black text-slate-300">14/01/2024</p>
                </div>
                <p class="text-[11px] font-bold text-slate-400 leading-relaxed uppercase">{{ $t('retention.log_detail.timeline.export.desc') }}</p>
              </div>
            </div>

            <!-- 3. Decision (Right) -->
            <div class="relative flex justify-end">
              <div class="absolute left-1/2 -translate-x-1/2 w-10 h-10 rounded-full bg-slate-50 border border-slate-100 flex items-center justify-center text-slate-300 z-10">
                <LucideGavel class="w-4 h-4" />
              </div>
              <div class="w-[45%] glass p-8 rounded-[2rem] border border-slate-50 shadow-xl shadow-slate-100/50 space-y-3" v-motion-slide-right>
                <div class="flex items-center justify-between">
                  <h4 class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight leading-tight">{{ $t('retention.log_detail.timeline.decision.title') }}</h4>
                  <p class="text-[10px] font-black text-slate-300">20/01/2024</p>
                </div>
                <p class="text-[11px] font-bold text-slate-400 leading-relaxed uppercase">{{ $t('retention.log_detail.timeline.decision.desc') }}</p>
              </div>
            </div>

            <!-- 4. Shredding (Left) -->
            <div class="relative flex justify-start">
              <div class="absolute left-1/2 -translate-x-1/2 w-10 h-10 rounded-full bg-slate-50 border border-slate-100 flex items-center justify-center text-slate-300 z-10">
                <LucideScissors class="w-4 h-4" />
              </div>
              <div class="w-[45%] glass p-8 rounded-[2rem] border border-slate-50 shadow-xl shadow-slate-100/50 space-y-3" v-motion-slide-left>
                <div class="flex items-center justify-between">
                  <h4 class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight leading-tight">{{ $t('retention.log_detail.timeline.shredding.title') }}</h4>
                  <p class="text-[10px] font-black text-slate-300">21/01/2024</p>
                </div>
                <p class="text-[11px] font-bold text-slate-400 leading-relaxed uppercase">{{ $t('retention.log_detail.timeline.shredding.desc') }}</p>
              </div>
            </div>

            <!-- 5. BAST (Right) -->
            <div class="relative flex justify-end">
              <div class="absolute left-1/2 -translate-x-1/2 w-10 h-10 rounded-full bg-slate-50 border border-slate-100 flex items-center justify-center text-slate-300 z-10">
                <LucideCloudUpload class="w-4 h-4" />
              </div>
              <div class="w-[45%] glass p-8 rounded-[2rem] border border-slate-50 shadow-xl shadow-slate-100/50 space-y-3" v-motion-slide-right>
                <div class="flex items-center justify-between">
                  <h4 class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight leading-tight">{{ $t('retention.log_detail.timeline.bast.title') }}</h4>
                  <p class="text-[10px] font-black text-slate-300">21/01/2024</p>
                </div>
                <p class="text-[11px] font-bold text-slate-400 leading-relaxed uppercase">{{ $t('retention.log_detail.timeline.bast.desc') }}</p>
              </div>
            </div>

            <!-- 6. Purged (Center) -->
            <div class="relative flex flex-col items-center">
              <div class="w-14 h-14 rounded-full bg-emerald-500 text-white flex items-center justify-center shadow-xl shadow-emerald-500/30 z-10" v-motion-pop>
                <LucideCheck class="w-8 h-8" />
              </div>
              <div class="mt-8 w-full max-w-sm glass p-8 rounded-[2rem] border-2 border-emerald-500 bg-emerald-50/10 shadow-2xl shadow-emerald-500/10 text-center space-y-3" v-motion-fade>
                <div class="flex items-center justify-between">
                  <h4 class="text-sm font-black text-emerald-600 uppercase tracking-tight leading-tight">{{ $t('retention.log_detail.timeline.purged.title') }}</h4>
                  <p class="text-[10px] font-black text-emerald-500">22/01/2024</p>
                </div>
                <p class="text-[11px] font-bold text-emerald-700/60 leading-relaxed uppercase">{{ $t('retention.log_detail.timeline.purged.desc') }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Bottom Actions -->
    <div class="flex justify-center pt-10" v-motion-fade>
      <button @click="navigateTo('/retention/history')" class="px-10 py-5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-[11px] font-black uppercase tracking-[0.2em] shadow-2xl shadow-blue-900/20 transition-all flex items-center gap-4 group">
        <LucideArrowLeft class="w-4 h-4 group-hover:-translate-x-1 transition-transform" /> {{ $t('retention.log_detail.btn_back') }}
      </button>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideLock, LucideFileText, LucidePaperclip, LucideEye, 
  LucideListChecks, LucideBell, LucideFileSpreadsheet, 
  LucideGavel, LucideScissors, LucideCloudUpload, LucideCheck,
  LucideArrowLeft
} from 'lucide-vue-next'
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
</style>
