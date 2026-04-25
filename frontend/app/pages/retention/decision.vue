<template>
  <div class="max-w-[1600px] mx-auto space-y-10 pb-20">
    <!-- Header -->
    <div class="space-y-4" v-motion-fade>
      <div class="flex items-center gap-2">
        <span class="px-3 py-1 bg-orange-50 text-orange-600 border border-orange-100 rounded-full text-[9px] font-black uppercase tracking-widest flex items-center gap-1.5">
          <div class="w-1.5 h-1.5 bg-orange-500 rounded-full animate-pulse"></div>
          {{ $t('retention.decision.status_pending') }}
        </span>
      </div>
      <div class="space-y-2">
        <h1 class="text-3xl font-black text-[#1E3A5F] tracking-tight uppercase">{{ $t('retention.decision.title') }}</h1>
        <p class="text-slate-500 font-medium">{{ $t('retention.decision.subtitle', { id: '#RET-2023-089' }) }}</p>
      </div>
    </div>

    <div class="grid grid-cols-12 gap-10 items-start">
      <!-- Left Sidebar: Process Tracker -->
      <aside class="col-span-3 glass p-10 rounded-[3rem] bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-10" v-motion-slide-left>
        <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">{{ $t('retention.decision.tracker.title') }}</h3>
        
        <div class="space-y-0 relative">
          <div class="absolute left-[1.125rem] top-2 bottom-8 w-0.5 bg-slate-100"></div>
          
          <!-- Step 1 -->
          <div class="relative flex items-start gap-6 pb-12">
            <div class="relative z-10 w-9 h-9 rounded-full bg-[#1E3A5F] text-white flex items-center justify-center shadow-lg shadow-blue-900/10">
              <LucideCheck class="w-4 h-4" />
            </div>
            <div>
              <h4 class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('retention.decision.tracker.initiation') }}</h4>
              <p class="text-[10px] font-bold text-slate-400 uppercase mt-1">{{ $t('retention.decision.tracker.init_desc', { date: 'Oct 12, 2023' }) }}</p>
            </div>
          </div>

          <!-- Step 2 -->
          <div class="relative flex items-start gap-6 pb-12">
            <div class="relative z-10 w-9 h-9 rounded-full bg-[#1E3A5F] text-white flex items-center justify-center shadow-lg shadow-blue-900/10">
              <LucideCheck class="w-4 h-4" />
            </div>
            <div>
              <h4 class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('retention.decision.tracker.verification') }}</h4>
              <p class="text-[10px] font-bold text-slate-400 uppercase mt-1">{{ $t('retention.decision.tracker.verif_desc') }}</p>
            </div>
          </div>

          <!-- Step 3 (Active) -->
          <div class="relative flex items-start gap-6 pb-12">
            <div class="relative z-10 w-9 h-9 rounded-full bg-white border-2 border-[#1E3A5F] flex items-center justify-center shadow-lg">
              <div class="w-2 h-2 bg-[#1E3A5F] rounded-full animate-pulse"></div>
            </div>
            <div>
              <h4 class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('retention.decision.tracker.decision') }}</h4>
              <p class="text-[10px] font-black text-orange-500 uppercase mt-1">{{ $t('retention.decision.tracker.dec_desc') }}</p>
            </div>
          </div>

          <!-- Step 4 -->
          <div class="relative flex items-start gap-6">
            <div class="relative z-10 w-9 h-9 rounded-full bg-slate-50 border border-slate-100 flex items-center justify-center text-slate-200">
              <LucideTrash2 class="w-4 h-4" />
            </div>
            <div>
              <h4 class="text-sm font-black text-slate-300 uppercase tracking-tight">{{ $t('retention.decision.tracker.final') }}</h4>
              <p class="text-[10px] font-bold text-slate-200 uppercase mt-1">{{ $t('retention.decision.tracker.final_desc') }}</p>
            </div>
          </div>
        </div>
      </aside>

      <!-- Main Content: Input Form -->
      <div class="col-span-9 space-y-10" v-motion-fade>
        <div class="glass p-12 rounded-[3.5rem] bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-12 relative overflow-hidden">
          <div class="absolute right-12 top-12 p-3 bg-slate-50 text-slate-200 rounded-2xl">
            <LucideFileSignature class="w-8 h-8" />
          </div>

          <div class="space-y-1">
            <p class="text-[10px] font-black text-primary-500 uppercase tracking-[0.2em]">{{ $t('retention.decision.form.actor') }}</p>
            <h2 class="text-2xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('retention.decision.form.title') }}</h2>
          </div>

          <div class="space-y-10">
            <!-- Decision Status -->
            <div class="space-y-6">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest flex items-center gap-2">
                {{ $t('retention.decision.form.status_label') }} <span class="text-red-500">*</span>
              </label>
              <div class="grid grid-cols-2 gap-6">
                <label @click="decision = 'extended'" 
                       :class="`p-8 border-2 rounded-[2.5rem] cursor-pointer transition-all flex items-center gap-6 ${decision === 'extended' ? 'border-[#1E3A5F] bg-blue-50/20' : 'border-slate-50 bg-slate-50/50 hover:border-slate-100'}`">
                  <div :class="`w-6 h-6 rounded-full border-2 flex items-center justify-center transition-all ${decision === 'extended' ? 'border-[#1E3A5F]' : 'border-slate-200'}`">
                    <div v-if="decision === 'extended'" class="w-2 h-2 bg-[#1E3A5F] rounded-full"></div>
                  </div>
                  <div class="space-y-1">
                    <p class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('retention.decision.form.extended') }}</p>
                    <p class="text-[10px] font-bold text-slate-400 uppercase">{{ $t('retention.decision.form.extended_desc') }}</p>
                  </div>
                </label>

                <label @click="decision = 'destroy'" 
                       :class="`p-8 border-2 rounded-[2.5rem] cursor-pointer transition-all flex items-center gap-6 ${decision === 'destroy' ? 'border-[#1E3A5F] bg-blue-50/20' : 'border-slate-50 bg-slate-50/50 hover:border-slate-100'}`">
                  <div :class="`w-6 h-6 rounded-full border-2 flex items-center justify-center transition-all ${decision === 'destroy' ? 'border-[#1E3A5F]' : 'border-slate-200'}`">
                    <div v-if="decision === 'destroy'" class="w-2 h-2 bg-[#1E3A5F] rounded-full"></div>
                  </div>
                  <div class="space-y-1">
                    <p class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('retention.decision.form.destroy') }}</p>
                    <p class="text-[10px] font-bold text-slate-400 uppercase">{{ $t('retention.decision.form.destroy_desc') }}</p>
                  </div>
                </label>
              </div>
            </div>

            <!-- Info Box -->
            <div class="p-8 bg-slate-50/80 border border-slate-100 rounded-[2.5rem] flex items-start gap-6">
              <LucideInfo class="w-5 h-5 text-primary-500 shrink-0 mt-0.5" />
              <p class="text-xs font-bold text-slate-500 leading-loose uppercase tracking-tight">
                {{ $t('retention.decision.form.info_text') }}
              </p>
            </div>

            <!-- Reference Input -->
            <div class="space-y-4">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest flex items-center gap-2">
                {{ $t('retention.decision.form.ref_label') }} <span class="text-red-500">*</span>
              </label>
              <div class="relative group">
                <div class="absolute left-6 top-1/2 -translate-y-1/2 text-slate-300 font-black text-xl">#</div>
                <input type="text" placeholder="e.g., BAST/OFF/2023/X/0042" 
                       class="w-full pl-12 pr-10 py-5 bg-slate-50 border border-slate-100 rounded-[2rem] text-sm font-black text-[#1E3A5F] outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all tracking-widest placeholder:tracking-normal placeholder:font-medium" />
              </div>
            </div>

            <!-- Upload Area -->
            <div class="space-y-4">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('retention.decision.form.upload_label') }}</label>
              <div class="aspect-[4/1] border-2 border-dashed border-slate-100 rounded-[2.5rem] flex flex-col items-center justify-center text-center p-8 space-y-3 bg-slate-50/30 group hover:border-primary-200 transition-all cursor-pointer">
                <div class="w-12 h-12 bg-white rounded-2xl flex items-center justify-center shadow-sm text-slate-300 group-hover:text-primary-500 transition-colors">
                  <LucideUploadCloud class="w-6 h-6" />
                </div>
                <div class="space-y-1">
                  <p class="text-[11px] font-black text-slate-500 uppercase tracking-widest">{{ $t('retention.decision.form.upload_desc') }}</p>
                  <p class="text-[9px] font-bold text-slate-300 uppercase tracking-widest">PDF, JPG or PNG (MAX. 5MB)</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Footer Actions -->
          <div class="flex items-center justify-end gap-6 pt-6 border-t border-slate-50">
            <button @click="navigateTo('/retention/export')" class="px-8 py-4 bg-white border border-slate-200 text-slate-400 rounded-2xl text-[11px] font-black uppercase tracking-widest hover:bg-slate-50 transition-all">
              {{ $t('retention.decision.form.btn_cancel') }}
            </button>
            <button @click="navigateTo('/retention/shredding')" class="px-10 py-5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-[11px] font-black uppercase tracking-[0.2em] shadow-xl shadow-blue-900/20 transition-all flex items-center gap-3 group">
              {{ $t('retention.decision.form.btn_submit') }} <LucideArrowRight class="w-4 h-4 group-hover:translate-x-1 transition-transform" />
            </button>
          </div>
        </div>

        <!-- Batch Summary Cards -->
        <div class="space-y-6 pt-4">
          <h3 class="text-xs font-black text-[#1E3A5F] uppercase tracking-[0.2em] px-2">{{ $t('retention.decision.form.summary_title') }}</h3>
          <div class="grid grid-cols-4 gap-6">
            <div class="p-8 bg-white border border-slate-100 rounded-[2.5rem] shadow-xl shadow-slate-200/30 space-y-1">
              <p class="text-[9px] font-black text-slate-300 uppercase tracking-widest">{{ $t('retention.decision.form.stats.items') }}</p>
              <p class="text-xl font-black text-[#1E3A5F]">1,240</p>
            </div>
            <div class="p-8 bg-white border border-slate-100 rounded-[2.5rem] shadow-xl shadow-slate-200/30 space-y-1">
              <p class="text-[9px] font-black text-slate-300 uppercase tracking-widest">{{ $t('retention.decision.form.stats.volume') }}</p>
              <p class="text-xl font-black text-[#1E3A5F]">45.2 m³</p>
            </div>
            <div class="p-8 bg-white border border-slate-100 rounded-[2.5rem] shadow-xl shadow-slate-200/30 space-y-1">
              <p class="text-[9px] font-black text-slate-300 uppercase tracking-widest">{{ $t('retention.decision.form.stats.source') }}</p>
              <p class="text-xl font-black text-[#1E3A5F] uppercase">Finance</p>
            </div>
            <div class="p-8 bg-white border border-slate-100 rounded-[2.5rem] shadow-xl shadow-slate-200/30 space-y-1">
              <p class="text-[9px] font-black text-slate-300 uppercase tracking-widest">{{ $t('retention.decision.form.stats.risk') }}</p>
              <p class="text-xl font-black text-orange-500 uppercase">Medium</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideCheck, LucideTrash2, LucideFileSignature, LucideInfo, 
  LucideUploadCloud, LucideArrowRight 
} from 'lucide-vue-next'

const decision = ref('destroy')
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
</style>
