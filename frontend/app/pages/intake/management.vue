<template>
  <div class="space-y-12 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="space-y-2 max-w-3xl">
      <h1 class="text-4xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('intake.management.title') }}</h1>
      <p class="text-sm font-bold text-slate-500 leading-relaxed uppercase tracking-tight">{{ $t('intake.management.subtitle') }}</p>
    </div>

    <!-- Mode Selector Cards -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-10">
      <div v-for="mode in [
        { id: 'single', icon: LucideScan, color: 'bg-blue-500' },
        { id: 'bulk', icon: LucideLayoutGrid, color: 'bg-amber-500' },
        { id: 'mailroom', icon: LucideMail, color: 'bg-blue-900' }
      ]" :key="mode.id" class="glass group p-10 rounded-[4rem] space-y-10 shadow-sm hover:shadow-2xl hover:-translate-y-2 transition-all border border-slate-50 dark:border-slate-800 relative overflow-hidden">
        <div :class="`w-20 h-20 rounded-[2rem] ${mode.color} text-white flex items-center justify-center shadow-2xl transition-transform group-hover:scale-110` ">
          <component :is="mode.icon" class="w-10 h-10" />
        </div>
        <div class="space-y-4 relative z-10">
          <h3 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight leading-tight">{{ $t(`intake.management.modes.${mode.id}.title`) }}</h3>
          <p class="text-xs font-bold text-slate-400 uppercase tracking-widest leading-relaxed">{{ $t(`intake.management.modes.${mode.id}.desc`) }}</p>
        </div>
        <button class="flex items-center gap-2 text-blue-500 font-black text-xs uppercase tracking-widest hover:translate-x-2 transition-transform relative z-10">
          {{ $t(`intake.management.modes.${mode.id}.btn`) }}
          <LucideArrowRight class="w-4 h-4" />
        </button>
        <!-- Background Accent -->
        <div :class="`absolute top-0 right-0 w-32 h-32 ${mode.color} opacity-5 -translate-x-6 -translate-y-6 rounded-full group-hover:scale-150 transition-transform` "></div>
      </div>
    </div>

    <!-- Recent Jobs Table -->
    <div class="glass rounded-[4rem] overflow-hidden shadow-sm border border-slate-50 dark:border-slate-800">
      <div class="p-12 flex items-center justify-between border-b border-slate-50 dark:border-slate-800">
        <h3 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('intake.management.jobs.title') }}</h3>
        <button class="flex items-center gap-2 text-[10px] font-black text-blue-500 uppercase tracking-widest hover:underline group">
          {{ $t('intake.management.jobs.view_all') }}
          <LucideExternalLink class="w-4 h-4 group-hover:translate-x-1 group-hover:-translate-y-1 transition-transform" />
        </button>
      </div>

      <div class="overflow-x-auto">
        <table class="w-full text-left">
          <thead>
            <tr class="bg-slate-50/50 dark:bg-slate-800/50 text-[9px] font-black text-slate-400 uppercase tracking-widest">
              <th class="p-10 pl-12">{{ $t('intake.management.jobs.cols.id') }}</th>
              <th class="p-10">{{ $t('intake.management.jobs.cols.type') }}</th>
              <th class="p-10">{{ $t('intake.management.jobs.cols.time') }}</th>
              <th class="p-10">{{ $t('intake.management.jobs.cols.status') }}</th>
              <th class="p-10 pr-12 text-right">ACTIONS</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
            <tr v-for="job in [
              { id: 'INV-2023-OCT-0012', name: 'Digitalization of AP Invoices', type: 'Bulk Import', time: 'Today, 14:22 WIB', status: 'running' },
              { id: 'MAIL-EXT-44910', name: 'OJK Regulation Letter', type: 'Mailroom Intake', time: 'Today, 10:05 WIB', status: 'success' },
              { id: 'USER-ENTRY-881', name: 'Customer KYC Form', type: 'Single Input', time: 'Yesterday, 16:45 WIB', status: 'draft' },
              { id: 'EXCEL-UPLOAD-992', name: 'Historical Archive 1990-2000', type: 'Bulk Import', time: 'Yesterday, 11:20 WIB', status: 'failed' }
            ]" :key="job.id" class="group hover:bg-slate-50/30 transition-all cursor-pointer">
              <td class="p-10 pl-12">
                <div class="space-y-1">
                  <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ job.id }}</p>
                  <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ job.name }}</p>
                </div>
              </td>
              <td class="p-10">
                <span class="px-4 py-1.5 bg-blue-50 dark:bg-slate-800 text-blue-500 rounded-lg text-[9px] font-black uppercase tracking-widest border border-blue-100 dark:border-slate-700">{{ job.type }}</span>
              </td>
              <td class="p-10">
                <p class="text-xs font-bold text-slate-500 uppercase tracking-tight leading-relaxed">{{ job.time }}</p>
              </td>
              <td class="p-10">
                <div class="flex items-center gap-3">
                  <LucideRotateCcw v-if="job.status === 'running'" class="w-4 h-4 text-blue-500 animate-spin" />
                  <LucideCheckCircle2 v-if="job.status === 'success'" class="w-4 h-4 text-green-500" />
                  <LucideFileText v-if="job.status === 'draft'" class="w-4 h-4 text-slate-400" />
                  <LucideAlertCircle v-if="job.status === 'failed'" class="w-4 h-4 text-red-500" />
                  <span :class="`text-[9px] font-black uppercase tracking-widest ${
                    job.status === 'running' ? 'text-blue-500' :
                    job.status === 'success' ? 'text-green-500' :
                    job.status === 'draft' ? 'text-slate-400' :
                    'text-red-500'
                  }`">
                    {{ $t(`intake.management.jobs.statuses.${job.status}`) }}
                  </span>
                </div>
              </td>
              <td class="p-10 pr-12 text-right">
                <button class="w-10 h-10 rounded-xl hover:bg-slate-100 dark:hover:bg-slate-800 flex items-center justify-center transition-all text-slate-300 hover:text-[#1E3A5F]">
                  <LucideMoreVertical class="w-5 h-5" />
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Bottom Section: Stats & Guide -->
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <div class="lg:col-span-3">
        <div class="glass p-10 rounded-[3rem] space-y-10 shadow-sm border border-slate-50 dark:border-slate-800 h-full flex flex-col justify-between bg-[#1E3A5F] text-white">
          <div class="space-y-4">
            <p class="text-[10px] font-black text-blue-300 uppercase tracking-widest">{{ $t('intake.management.stats.today') }}</p>
            <div class="space-y-1">
              <p class="text-4xl font-black tracking-tighter">1,284</p>
              <p class="text-[10px] font-black text-green-400 uppercase tracking-widest flex items-center gap-2">
                <LucideTrendingUp class="w-3.5 h-3.5" />
                {{ $t('intake.management.stats.vs_yesterday', { val: 12 }) }}
              </p>
            </div>
          </div>
          <div class="space-y-4">
            <div class="flex items-center justify-between text-[10px] font-black uppercase tracking-widest text-blue-300">
              <span>{{ $t('intake.management.stats.ocr_cap') }}</span>
              <span>68%</span>
            </div>
            <div class="h-1.5 bg-blue-900 rounded-full overflow-hidden shadow-inner">
              <div class="h-full bg-blue-400" style="width: 68%"></div>
            </div>
          </div>
        </div>
      </div>

      <div class="lg:col-span-9">
        <div class="glass p-10 rounded-[3.5rem] shadow-sm border border-slate-50 dark:border-slate-800 flex flex-col md:flex-row items-center gap-10">
          <div class="flex-grow space-y-8">
            <div class="space-y-4">
              <h3 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight leading-tight">{{ $t('intake.management.guide.title') }}</h3>
              <p class="text-xs font-bold text-slate-400 uppercase tracking-widest leading-relaxed max-w-2xl">
                {{ $t('intake.management.guide.desc') }}
              </p>
            </div>
            <div class="flex flex-wrap items-center gap-6">
              <button class="px-8 py-4 bg-white dark:bg-slate-800 border-2 border-slate-50 dark:border-slate-700 text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest rounded-2xl hover:bg-slate-50 transition-all shadow-sm">
                {{ $t('intake.management.guide.btn_template') }}
              </button>
              <button class="px-8 py-4 bg-white dark:bg-slate-800 border-2 border-slate-50 dark:border-slate-700 text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest rounded-2xl hover:bg-slate-50 transition-all shadow-sm">
                {{ $t('intake.management.guide.btn_api') }}
              </button>
            </div>
          </div>
          <div class="w-full md:w-64 aspect-square rounded-[3rem] bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 flex items-center justify-center p-10 group overflow-hidden">
            <div class="relative w-full h-full">
              <div class="absolute inset-0 bg-blue-500/10 blur-3xl rounded-full scale-150 animate-pulse"></div>
              <LucidePrinter class="w-full h-full text-[#1E3A5F] relative z-10 group-hover:scale-110 transition-transform duration-700" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideScan, LucideLayoutGrid, LucideMail, LucideArrowRight, 
  LucideExternalLink, LucideRotateCcw, LucideCheckCircle2, 
  LucideFileText, LucideAlertCircle, LucideMoreVertical, 
  LucideTrendingUp, LucidePrinter 
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

