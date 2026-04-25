<template>
  <div class="max-w-[1600px] mx-auto grid grid-cols-12 gap-8 items-start pb-20">
    <!-- Main Content -->
    <div class="col-span-8 space-y-8">
      <!-- Header -->
      <div class="space-y-4" v-motion-fade>
        <div class="flex items-center gap-3">
          <h1 class="text-3xl font-black text-[#1E3A5F] tracking-tight uppercase">{{ $t('retention.export.title') }}</h1>
        </div>
        <div class="flex items-center gap-4 text-xs font-bold text-slate-400 uppercase tracking-widest">
          <span class="flex items-center gap-2 text-emerald-500">
            <div class="w-2 h-2 bg-emerald-500 rounded-full animate-pulse"></div>
            {{ $t('retention.export.status_validated') }}
          </span>
          <span class="text-slate-300">—</span>
          <span>{{ $t('retention.export.subtitle') }}</span>
        </div>
      </div>

      <!-- Download Card -->
      <div class="glass p-10 rounded-[3rem] bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 relative overflow-hidden group" v-motion-slide-visible-bottom>
        <div class="absolute top-0 left-0 right-0 h-1.5 bg-emerald-500"></div>
        <div class="flex items-center gap-8">
          <div class="w-24 h-24 bg-emerald-50 rounded-[2rem] flex items-center justify-center text-emerald-500 shadow-lg shadow-emerald-500/10 group-hover:scale-110 transition-transform duration-500">
            <LucideFileSpreadsheet class="w-12 h-12" />
          </div>
          <div class="space-y-6 flex-1">
            <div class="space-y-1">
              <h3 class="text-xl font-black text-[#1E3A5F] tracking-tight uppercase">Retention_Evaluation_Q1.xlsx</h3>
              <div class="flex items-center gap-6 text-[10px] font-bold text-slate-400 uppercase tracking-widest">
                <span class="flex items-center gap-1.5"><LucideUser class="w-3.5 h-3.5" /> {{ $t('retention.export.download_card.generated_by', { name: 'Admin A' }) }}</span>
                <span class="flex items-center gap-1.5"><LucideClock class="w-3.5 h-3.5" /> 2024-05-20 10:00</span>
              </div>
            </div>
            <div class="flex items-center gap-4">
              <button class="px-8 py-4 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-xs font-black uppercase tracking-[0.2em] shadow-xl shadow-blue-900/20 transition-all flex items-center gap-3">
                <LucideDownload class="w-4 h-4" /> {{ $t('retention.export.download_card.btn_download') }}
              </button>
              <button class="px-6 py-4 bg-white border border-slate-200 text-slate-400 hover:text-slate-600 rounded-2xl text-xs font-black uppercase tracking-widest transition-all flex items-center gap-2">
                <LucideEye class="w-4 h-4" /> {{ $t('retention.export.download_card.btn_preview') }}
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Compliance Checklist Reminder -->
      <div class="glass p-10 rounded-[3rem] bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-10" v-motion-slide-visible-bottom>
        <div class="flex items-center gap-4">
          <div class="p-3 bg-orange-50 text-orange-500 rounded-2xl">
            <LucideCheckSquare class="w-6 h-6" />
          </div>
          <h2 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('retention.export.checklist.title') }}</h2>
        </div>

        <p class="text-sm font-medium text-slate-500 leading-relaxed px-2">
          {{ $t('retention.export.checklist.desc') }}
        </p>

        <div class="grid grid-cols-3 gap-6">
          <div class="p-8 bg-slate-50/50 border border-slate-100 rounded-[2.5rem] space-y-6 hover:bg-white hover:shadow-xl hover:shadow-slate-200/30 transition-all group">
            <div class="flex items-center justify-between text-slate-300 group-hover:text-primary-500 transition-colors">
              <p class="text-[9px] font-black uppercase tracking-widest">{{ $t('retention.export.checklist.litigation_title') }}</p>
              <LucideGavel class="w-4 h-4" />
            </div>
            <div class="space-y-2">
              <h4 class="text-xs font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('retention.export.checklist.litigation_title') }}</h4>
              <p class="text-[10px] font-bold text-slate-400 leading-relaxed uppercase">{{ $t('retention.export.checklist.litigation_desc') }}</p>
            </div>
          </div>

          <div class="p-8 bg-slate-50/50 border border-slate-100 rounded-[2.5rem] space-y-6 hover:bg-white hover:shadow-xl hover:shadow-slate-200/30 transition-all group">
            <div class="flex items-center justify-between text-slate-300 group-hover:text-emerald-500 transition-colors">
              <p class="text-[9px] font-black uppercase tracking-widest">{{ $t('retention.export.checklist.financial_title') }}</p>
              <LucideBanknote class="w-4 h-4" />
            </div>
            <div class="space-y-2">
              <h4 class="text-xs font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('retention.export.checklist.financial_title') }}</h4>
              <p class="text-[10px] font-bold text-slate-400 leading-relaxed uppercase">{{ $t('retention.export.checklist.financial_desc') }}</p>
            </div>
          </div>

          <div class="p-8 bg-slate-50/50 border border-slate-100 rounded-[2.5rem] space-y-6 hover:bg-white hover:shadow-xl hover:shadow-slate-200/30 transition-all group">
            <div class="flex items-center justify-between text-slate-300 group-hover:text-purple-500 transition-colors">
              <p class="text-[9px] font-black uppercase tracking-widest">{{ $t('retention.export.checklist.internal_title') }}</p>
              <LucideSearch class="w-4 h-4" />
            </div>
            <div class="space-y-2">
              <h4 class="text-xs font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('retention.export.checklist.internal_title') }}</h4>
              <p class="text-[10px] font-bold text-slate-400 leading-relaxed uppercase">{{ $t('retention.export.checklist.internal_desc') }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Right Sidebar -->
    <div class="col-span-4 space-y-8" v-motion-slide-right>
      <!-- Offline Process Note -->
      <div class="glass p-10 rounded-[3rem] bg-slate-50/50 border border-slate-100 shadow-xl space-y-8">
        <div class="flex items-center gap-3 text-slate-400 uppercase tracking-widest text-[10px] font-black">
          <LucideInfo class="w-4 h-4 text-primary-500" />
          {{ $t('retention.export.sidebar.note_title') }}
        </div>
        <p class="text-xs font-bold text-slate-500 italic leading-loose">
          {{ $t('retention.export.sidebar.note_text') }}
        </p>

        <div class="pt-8 border-t border-slate-200 space-y-6">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('retention.export.sidebar.next_steps') }}</p>
          <button @click="navigateTo('/retention/decision')" class="w-full p-8 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-[2rem] shadow-2xl shadow-blue-900/30 transition-all flex items-center justify-between group">
            <span class="text-sm font-black text-left uppercase tracking-tight">
              {{ $t('retention.export.sidebar.btn_wait') }}
            </span>
            <LucideArrowRight class="w-6 h-6 group-hover:translate-x-2 transition-transform" />
          </button>
          <p class="text-[9px] font-bold text-slate-300 text-center uppercase tracking-widest leading-loose px-4">
            {{ $t('retention.export.sidebar.footer_hint') }}
          </p>
        </div>
      </div>

      <!-- Export Visualization -->
      <div class="glass p-8 rounded-[3rem] bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-8">
        <div class="flex items-center gap-3 text-slate-400 uppercase tracking-widest text-[10px] font-black">
          <LucideLineChart class="w-4 h-4 text-primary-500" />
          {{ $t('retention.export.sidebar.viz_title') }}
        </div>
        <div class="aspect-video bg-slate-100 rounded-2xl overflow-hidden relative group">
          <img src="https://images.unsplash.com/photo-1460925895917-afdab827c52f?auto=format&fit=crop&q=80&w=400" 
               class="w-full h-full object-cover opacity-30 mix-blend-multiply transition-transform duration-700 group-hover:scale-110" />
          <div class="absolute inset-0 flex items-center justify-center">
            <LucideFileSpreadsheet class="w-12 h-12 text-slate-300" />
          </div>
        </div>
        <div class="space-y-3">
          <div class="flex items-center justify-between text-[10px] font-black uppercase tracking-widest">
            <span class="text-slate-400">{{ $t('retention.export.sidebar.total_records') }}</span>
            <span class="text-[#1E3A5F]">1,248 items</span>
          </div>
          <div class="h-1.5 bg-slate-100 rounded-full overflow-hidden">
            <div class="h-full bg-[#1E3A5F] w-[85%] rounded-full shadow-lg shadow-blue-500/20"></div>
          </div>
        </div>
      </div>
    </div>v>
  </div>
</template>

<script setup>
import { 
  LucideFileSpreadsheet, LucideUser, LucideClock, LucideDownload, 
  LucideEye, LucideCheckSquare, LucideGavel, LucideBanknote, 
  LucideSearch, LucideInfo, LucideArrowRight, LucideLineChart 
} from 'lucide-vue-next'
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
</style>
