<template>
  <div class="max-w-[1200px] mx-auto space-y-8 pb-20 pt-10">
    <!-- Success Header Card -->
    <div class="glass p-16 rounded-[4rem] bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 flex flex-col items-center text-center space-y-12" v-motion-pop>
      <!-- Success Icon -->
      <div class="w-28 h-28 bg-emerald-50 rounded-full flex items-center justify-center text-emerald-500 shadow-xl shadow-emerald-500/20">
        <LucideCheckCircle2 class="w-14 h-14" />
      </div>

      <!-- Message -->
      <div class="space-y-4">
        <p class="text-[11px] font-black text-emerald-500 uppercase tracking-[0.4em]">{{ $t('retention.purge_success.status') }}</p>
        <h2 class="text-3xl font-black text-[#1E3A5F] tracking-tight uppercase">{{ $t('retention.purge_success.title') }}</h2>
        <p class="text-xs font-bold text-slate-400 leading-relaxed max-w-[450px] mx-auto">
          {{ $t('retention.purge_success.desc') }}
        </p>
      </div>

      <!-- Actions -->
      <div class="flex items-center gap-4">
        <button @click="navigateTo('/retention/log-detail')" class="px-8 py-4 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-[10px] font-black uppercase tracking-[0.2em] shadow-xl shadow-blue-900/20 transition-all flex items-center gap-3">
          <LucideScrollText class="w-4 h-4" /> {{ $t('retention.purge_success.btn_log') }}
        </button>
        <button @click="navigateTo('/retention/history')" class="px-8 py-4 bg-slate-50 border border-slate-100 text-slate-400 hover:text-slate-600 rounded-2xl text-[10px] font-black uppercase tracking-widest transition-all">
          {{ $t('retention.purge_success.btn_history') }}
        </button>
      </div>
    </div>

    <!-- Transaction Summary -->
    <div class="glass p-12 rounded-lg bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-10" v-motion-slide-visible-bottom>
      <div class="flex items-center gap-4">
        <div class="p-3 bg-slate-50 text-[#1E3A5F] rounded-2xl">
          <LucideClipboardList class="w-6 h-6" />
        </div>
        <h2 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('retention.purge_success.summary_title') }}</h2>
      </div>

      <div class="grid grid-cols-3 gap-8">
        <div class="p-8 bg-slate-50/50 border border-slate-50 rounded-lg space-y-3">
          <p class="text-[9px] font-black text-slate-300 uppercase tracking-widest">{{ $t('retention.purge_success.stats.log_id') }}</p>
          <p class="text-lg font-black text-[#1E3A5F] tracking-tight uppercase">TR-PURGE-999</p>
        </div>
        <div class="p-8 bg-slate-50/50 border border-slate-50 rounded-lg space-y-3">
          <p class="text-[9px] font-black text-slate-300 uppercase tracking-widest">{{ $t('retention.purge_success.stats.actor') }}</p>
          <div class="flex items-center gap-3">
            <div class="w-8 h-8 rounded-full bg-primary-100 flex items-center justify-center text-primary-600">
              <LucideUser class="w-4 h-4" />
            </div>
            <p class="text-lg font-black text-[#1E3A5F] tracking-tight uppercase">Admin A</p>
          </div>
        </div>
        <div class="p-8 bg-slate-50/50 border border-slate-50 rounded-lg space-y-3">
          <p class="text-[9px] font-black text-slate-300 uppercase tracking-widest">{{ $t('retention.purge_success.stats.timestamp') }}</p>
          <div class="flex items-center gap-3">
            <LucideClock class="w-5 h-5 text-slate-300" />
            <p class="text-lg font-black text-[#1E3A5F] tracking-tight uppercase">2024-05-20 14:30</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Important Note -->
    <div class="p-8 bg-blue-50/50 border border-blue-100 rounded-lg flex items-start gap-6" v-motion-fade>
      <div class="p-3 bg-white/50 rounded-2xl text-blue-500 shadow-sm">
        <LucideInfo class="w-5 h-5" />
      </div>
      <div class="space-y-1">
        <h4 class="text-xs font-black text-blue-900 uppercase tracking-tight">{{ $t('retention.purge_success.footer_note_title') }}</h4>
        <p class="text-[11px] font-bold text-blue-700/70 leading-relaxed uppercase tracking-tight">
          {{ $t('retention.purge_success.footer_note_text') }}
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideCheckCircle2, LucideScrollText, LucideClipboardList, 
  LucideUser, LucideClock, LucideInfo 
} from 'lucide-vue-next'
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
</style>
