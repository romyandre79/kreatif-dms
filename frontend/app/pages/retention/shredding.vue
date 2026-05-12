<template>
  <div class="max-w-[1600px] mx-auto space-y-10 pb-20">
    <!-- Header -->
    <div class="space-y-4" v-motion-fade>
      <div class="flex items-center gap-2 text-[10px] font-black text-slate-300 uppercase tracking-widest">
        <span>Home</span>
        <LucideChevronRight class="w-3 h-3" />
        <span>Retention & Disposal</span>
        <LucideChevronRight class="w-3 h-3" />
        <span class="text-primary-500">Disposal History</span>
      </div>
      <div class="space-y-2">
        <h1 class="text-3xl font-black text-[#1E3A5F] tracking-tight uppercase">{{ $t('retention.shredding.title') }}</h1>
        <p class="text-slate-500 font-medium">{{ $t('retention.shredding.subtitle') }}</p>
      </div>
    </div>

    <div class="grid grid-cols-12 gap-10 items-start">
      <!-- Left Sidebar: Batch Identity -->
      <aside class="col-span-3 glass p-10 rounded-lg bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-10" v-motion-slide-left>
        <div class="space-y-1">
          <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ $t('retention.shredding.identity.title') }}</p>
          <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('retention.shredding.identity.label') }}</p>
          <p class="text-xl font-black text-[#1E3A5F]">BATCH-2024-001</p>
        </div>

        <div class="space-y-4 pt-6 border-t border-slate-50">
          <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('retention.shredding.identity.status_label') }}</p>
          <span class="px-3 py-1 bg-emerald-50 text-emerald-600 border border-emerald-100 rounded-full text-[9px] font-black uppercase tracking-widest flex items-center gap-2 w-fit">
            <div class="w-1.5 h-1.5 bg-emerald-500 rounded-full"></div>
            {{ $t('retention.shredding.identity.status_approved') }}
          </span>
        </div>

        <div class="space-y-4 pt-6 border-t border-slate-50">
          <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('retention.shredding.identity.docs_label') }}</p>
          <div class="p-4 bg-slate-50 border border-slate-100 rounded-2xl flex items-center gap-3 group cursor-pointer hover:bg-primary-50 transition-colors">
            <LucideFileText class="w-4 h-4 text-primary-500" />
            <p class="text-[10px] font-black text-[#1E3A5F] uppercase truncate">BAST_2024_001_SIGNED.pdf</p>
          </div>
        </div>
      </aside>

      <!-- Main Content -->
      <div class="col-span-9 space-y-8" v-motion-fade>
        <!-- Pre-Execution Checklist -->
        <div class="glass p-10 rounded-lg bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-8">
          <h3 class="text-sm font-black text-[#1E3A5F] uppercase tracking-[0.2em]">{{ $t('retention.shredding.checklist.title') }}</h3>
          
          <div class="space-y-4">
            <label class="flex items-start gap-6 p-6 bg-slate-50/50 border border-slate-50 rounded-2xl cursor-pointer group hover:bg-white hover:border-primary-100 transition-all">
              <input type="checkbox" class="w-6 h-6 mt-1 rounded border-2 border-slate-200 text-primary-500 focus:ring-primary-500/10" />
              <div class="space-y-1">
                <p class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight group-hover:text-primary-600 transition-colors">{{ $t('retention.shredding.checklist.verify_qty') }}</p>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest leading-relaxed">{{ $t('retention.shredding.checklist.verify_qty_desc') }}</p>
              </div>
            </label>

            <label class="flex items-start gap-6 p-6 bg-slate-50/50 border border-slate-50 rounded-2xl cursor-pointer group hover:bg-white hover:border-primary-100 transition-all">
              <input type="checkbox" class="w-6 h-6 mt-1 rounded border-2 border-slate-200 text-primary-500 focus:ring-primary-500/10" />
              <div class="space-y-1">
                <p class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight group-hover:text-primary-600 transition-colors">{{ $t('retention.shredding.checklist.witness') }}</p>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest leading-relaxed">{{ $t('retention.shredding.checklist.witness_desc') }}</p>
              </div>
            </label>

            <label class="flex items-start gap-6 p-6 bg-slate-50/50 border border-slate-50 rounded-2xl cursor-pointer group hover:bg-white hover:border-primary-100 transition-all">
              <input type="checkbox" class="w-6 h-6 mt-1 rounded border-2 border-slate-200 text-primary-500 focus:ring-primary-500/10" />
              <div class="space-y-1">
                <p class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight group-hover:text-primary-600 transition-colors">{{ $t('retention.shredding.checklist.vendor') }}</p>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest leading-relaxed">{{ $t('retention.shredding.checklist.vendor_desc') }}</p>
              </div>
            </label>
          </div>
        </div>

        <!-- Execution Log Details -->
        <div class="glass p-10 rounded-lg bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-10">
          <h3 class="text-sm font-black text-[#1E3A5F] uppercase tracking-[0.2em]">{{ $t('retention.shredding.log.title') }}</h3>
          
          <div class="grid grid-cols-2 gap-10">
            <div class="space-y-4">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest flex items-center gap-2">
                {{ $t('retention.shredding.log.date_label') }} <span class="text-red-500">*</span>
              </label>
              <div class="relative group">
                <input type="text" value="10/24/2024, 09:00 AM" 
                       class="w-full px-6 py-5 bg-slate-50 border border-slate-100 rounded-2xl text-sm font-black text-[#1E3A5F] outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all tracking-tight" />
              </div>
            </div>

            <div class="space-y-4">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('retention.shredding.log.operator_label') }}</label>
              <div class="w-full px-6 py-5 bg-slate-50 border border-slate-100 rounded-2xl text-sm font-bold text-slate-400 tracking-tight">
                Admin - Ahmad Riyadi
              </div>
            </div>
          </div>

          <div class="space-y-4 pt-6">
            <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('retention.shredding.log.notes_label') }}</label>
            <textarea :placeholder="$t('retention.shredding.log.notes_placeholder')" 
                      rows="4" 
                      class="w-full p-8 bg-slate-50 border border-slate-100 rounded-2xl text-sm font-medium text-slate-600 outline-none focus:ring-4 focus:ring-primary-500/5 focus:border-primary-500 transition-all resize-none"></textarea>
          </div>
        </div>

        <!-- Footer Actions -->
        <div class="flex items-center justify-between pt-6">
          <button @click="navigateTo('/retention/decision')" class="text-[11px] font-black text-slate-400 hover:text-slate-600 uppercase tracking-[0.2em] transition-colors px-10">
            {{ $t('retention.shredding.footer.btn_cancel') }}
          </button>
          <button @click="navigateTo('/retention/upload-bast')" class="px-12 py-6 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-3xl text-xs font-black uppercase tracking-[0.2em] shadow-2xl shadow-blue-900/20 transition-all flex items-center gap-4 group">
            <LucideCheckCircle2 class="w-5 h-5 group-hover:scale-110 transition-transform" /> {{ $t('retention.shredding.footer.btn_confirm') }}
          </button>
        </div>
      </div>

      <!-- Upload Full Width Area -->
      <div class="col-span-12 glass p-12 rounded-[3.5rem] bg-slate-50/50 border border-slate-100 shadow-xl space-y-10" v-motion-slide-visible-bottom>
        <div class="flex flex-col items-center justify-center text-center space-y-6 py-10">
          <div class="w-20 h-20 bg-white rounded-lg flex items-center justify-center shadow-sm text-slate-200">
            <LucideCamera class="w-10 h-10" />
          </div>
          <div class="space-y-2">
            <h3 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('retention.shredding.evidence.title') }}</h3>
            <p class="text-xs font-bold text-slate-400 uppercase tracking-widest leading-loose">{{ $t('retention.shredding.evidence.desc') }}</p>
          </div>
          <button class="px-10 py-4 bg-white border border-slate-200 text-[#1E3A5F] rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-50 transition-all shadow-sm">
            {{ $t('retention.shredding.evidence.btn_choose') }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideChevronRight, LucideFileText, LucideCheckCircle2, 
  LucideCamera, LucideTrash2 
} from 'lucide-vue-next'
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
</style>
