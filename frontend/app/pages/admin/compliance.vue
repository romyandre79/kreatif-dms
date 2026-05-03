<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-1">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('compliance.title') }}</h1>
        <p class="text-xs font-bold text-slate-500">{{ $t('compliance.subtitle') }}</p>
      </div>
      <div class="flex items-center gap-4">
        <button class="px-6 py-3 text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest flex items-center gap-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl transition-all border border-slate-200 dark:border-slate-700">
          <LucideFileSpreadsheet class="w-4 h-4" />
          {{ $t('compliance.header.export') }}
        </button>
        <button class="px-10 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 group">
          <LucideZap class="w-4 h-4 group-hover:scale-110 transition-transform" />
          {{ $t('compliance.header.run') }}
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Column: Status Verifikasi -->
      <div class="lg:col-span-3 space-y-8">
        <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-4">{{ $t('compliance.verification.title') }}</h4>
        <div class="space-y-4">
          <div v-for="v in [
            { id: 'tls', status: 'active', color: 'text-green-500', bg: 'bg-green-50' },
            { id: 'ssl', status: 'critical', color: 'text-red-500', bg: 'bg-red-50' },
            { id: 'aes', status: 'verified', color: 'text-green-500', bg: 'bg-green-50' },
            { id: 'patch', status: 'warning', color: 'text-amber-500', bg: 'bg-amber-50' },
            { id: 'audit', status: 'synced', color: 'text-green-500', bg: 'bg-green-50' }
          ]" :key="v.id" class="p-6 bg-white dark:bg-slate-900 rounded-3xl border border-slate-100 dark:border-slate-800 space-y-4 shadow-sm hover:scale-[1.02] transition-all cursor-pointer">
            <div class="flex items-start justify-between">
              <div class="space-y-1">
                <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t(`compliance.verification.${v.id}`) }}</p>
                <p class="text-[9px] font-bold text-slate-400 uppercase">{{ $t(`compliance.verification.${v.id}_desc`) }}</p>
              </div>
              <LucideCheckCircle2 v-if="v.status !== 'critical' && v.status !== 'warning'" class="w-4 h-4 text-green-500" />
              <LucideAlertTriangle v-else :class="`w-4 h-4 ${v.color}`" />
            </div>
            <span :class="`px-3 py-1 rounded text-[8px] font-black uppercase tracking-widest border ${v.bg} ${v.color} border-${v.color}/20`">{{ v.status.toUpperCase() }}</span>
          </div>
        </div>
      </div>

      <!-- Center Column: Score & History -->
      <div class="lg:col-span-6 space-y-10">
        <!-- Score Card -->
        <div class="glass p-10 rounded-[3rem] space-y-10 shadow-xl border-t-8 border-[#1E3A5F]">
          <div class="flex items-center justify-between">
            <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('compliance.score.title') }}</h3>
            <div class="flex items-center gap-2 text-[8px] font-bold text-slate-400 uppercase">
              <LucideClock class="w-3.5 h-3.5" />
              {{ $t('compliance.score.last_scan', { time: '09:42' }) }}
            </div>
          </div>
          <div class="grid grid-cols-4 gap-6">
            <div class="col-span-1 flex flex-col items-center justify-center p-6 bg-blue-50/50 dark:bg-blue-900/10 rounded-[2rem]">
              <p class="text-4xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">84%</p>
              <p class="text-[8px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest text-center mt-2">{{ $t('compliance.score.health') }}</p>
            </div>
            <div v-for="s in [
              { id: 'critical', val: '01', color: 'text-red-500' },
              { id: 'warnings', val: '03', color: 'text-amber-500' },
              { id: 'passed', val: '12', color: 'text-green-500' }
            ]" :key="s.id" class="flex flex-col items-center justify-center">
              <p :class="`text-3xl font-black ${s.color} tracking-tighter`">{{ s.val }}</p>
              <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest mt-1">{{ $t(`compliance.score.${s.id}`) }}</p>
            </div>
          </div>
          <div class="space-y-4">
            <div class="h-2.5 w-full bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden flex shadow-inner">
              <div class="h-full bg-green-500" style="width: 75%"></div>
              <div class="h-full bg-amber-500" style="width: 15%"></div>
              <div class="h-full bg-red-500" style="width: 10%"></div>
            </div>
            <div class="flex items-center justify-between text-[8px] font-black text-slate-400 uppercase tracking-widest">
              <span>Deployment Integrity (F-57 Protocol)</span>
              <span>{{ $t('compliance.score.threshold') }}</span>
            </div>
          </div>
        </div>

        <!-- History Table -->
        <div class="glass rounded-[3rem] overflow-hidden shadow-sm">
          <div class="p-8 flex items-center justify-between border-b border-slate-50 dark:border-slate-800">
            <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('compliance.history.title') }}</h3>
            <button class="text-[9px] font-black text-blue-500 uppercase tracking-widest hover:underline transition-all">{{ $t('compliance.history.view_all') }}</button>
          </div>
          <table class="w-full text-left">
            <thead>
              <tr class="bg-slate-50/50 dark:bg-slate-800/50 text-[8px] font-black text-slate-400 uppercase tracking-widest">
                <th class="p-6 pl-10">{{ $t('compliance.history.cols.source') }}</th>
                <th class="p-6">{{ $t('compliance.history.cols.id') }}</th>
                <th class="p-6">{{ $t('compliance.history.cols.result') }}</th>
                <th class="p-6 pr-10 text-right">{{ $t('compliance.history.cols.time') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
              <tr v-for="h in [
                { source: 'Node-04 Firewall', id: 'F57-002-C', result: 'PASSED', color: 'text-green-500', bg: 'bg-green-50', time: '09:42:11' },
                { source: 'LDAP Auth Sync', id: 'F57-009-S', result: 'WARNING', color: 'text-amber-500', bg: 'bg-amber-50', time: '09:41:05' },
                { source: 'Kernel Hardening', id: 'F57-012-H', result: 'PASSED', color: 'text-green-500', bg: 'bg-green-50', time: '09:40:22' }
              ]" :key="h.id" class="group hover:bg-slate-50/30 transition-all cursor-pointer">
                <td class="p-6 pl-10">
                  <div class="space-y-0.5">
                    <p class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ h.source }}</p>
                  </div>
                </td>
                <td class="p-6 text-[10px] font-bold text-slate-400 uppercase font-mono">{{ h.id }}</td>
                <td class="p-6">
                  <span :class="`px-3 py-1 rounded text-[8px] font-black uppercase tracking-widest border ${h.bg} ${h.color} border-${h.color}/20`">{{ h.result }}</span>
                </td>
                <td class="p-6 pr-10 text-right text-[10px] font-bold text-slate-400 uppercase">{{ h.time }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Right Column: Remediation -->
      <div class="lg:col-span-3 space-y-10">
        <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-4">{{ $t('compliance.remediation.title') }}</h4>
        <div class="space-y-6">
          <div v-for="t in [
            { id: 'cert', type: 'urgent', color: 'border-red-500', btn: 'btn_renew' },
            { id: 'kb', type: 'patch', color: 'border-amber-500', btn: 'btn_patch' },
            { id: 'ssh', type: 'hardening', color: 'border-blue-500', btn: 'btn_auto' }
          ]" :key="t.id" class="p-8 bg-white dark:bg-slate-900 rounded-[2.5rem] border-l-8 border-slate-100 dark:border-slate-800 transition-all hover:border-[#1E3A5F] shadow-sm relative group overflow-hidden">
            <div :class="`absolute top-0 left-0 w-1.5 h-full ${t.color}`"></div>
            <div class="flex items-center justify-between mb-4">
              <span :class="`text-[8px] font-black uppercase tracking-widest ${t.type === 'urgent' ? 'text-red-500' : t.type === 'patch' ? 'text-amber-500' : 'text-blue-500'}`">{{ $t(`compliance.remediation.${t.type}`) }}</span>
              <LucideExternalLink class="w-3.5 h-3.5 text-slate-300 opacity-0 group-hover:opacity-100 transition-opacity" />
            </div>
            <div class="space-y-2 mb-8">
              <h5 class="text-[13px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t(`compliance.remediation.${t.id}`) }}</h5>
              <p class="text-[10px] font-bold text-slate-400 leading-relaxed">{{ $t(`compliance.remediation.${t.id}_desc`) }}</p>
            </div>
            <button :class="`w-full py-4 rounded-xl text-[9px] font-black uppercase tracking-widest transition-all ${t.type === 'urgent' ? 'bg-red-500 text-white shadow-lg shadow-red-500/20 hover:bg-red-600' : 'bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-[#1E3A5F] dark:text-white hover:bg-slate-50'}`">
              {{ $t(`compliance.remediation.${t.btn}`) }}
            </button>
          </div>
        </div>

        <!-- Manual Card -->
        <div class="bg-[#1E3A5F] p-8 rounded-[2.5rem] text-white space-y-6 relative overflow-hidden group shadow-2xl shadow-blue-900/40">
          <LucideFileText class="absolute top-0 right-0 w-32 h-32 text-white/5 -rotate-12 translate-x-10 -translate-y-10 group-hover:scale-110 transition-transform" />
          <div class="space-y-2">
            <h4 class="text-[11px] font-black uppercase tracking-widest opacity-60">{{ $t('compliance.manual.title') }}</h4>
            <p class="text-xs font-bold leading-relaxed opacity-80">{{ $t('compliance.manual.desc') }}</p>
          </div>
          <button class="flex items-center gap-2 text-[10px] font-black uppercase tracking-widest group-hover:gap-4 transition-all">
            {{ $t('compliance.manual.btn') }}
            <LucideArrowRight class="w-4 h-4" />
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideFileSpreadsheet, LucideZap, LucideCheckCircle2, LucideAlertTriangle, 
  LucideClock, LucideExternalLink, LucideFileText, LucideArrowRight 
} from 'lucide-vue-next'

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
@reference "tailwindcss";

.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}
</style>
