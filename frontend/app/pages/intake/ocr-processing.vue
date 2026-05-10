<template>
  <div class="space-y-12 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('intake.processing.title') }}</h1>
      <div class="flex items-center gap-6">
        <div class="px-6 py-2 bg-blue-50 dark:bg-slate-800 rounded-full flex items-center gap-4 border border-blue-100 dark:border-slate-700">
          <div class="w-2 h-2 bg-blue-500 rounded-full animate-ping"></div>
          <span class="text-[9px] font-black text-blue-500 uppercase tracking-widest">{{ $t('intake.processing.running') }}</span>
        </div>
        <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('intake.processing.started', { time: '4m 12s' }) }}</p>
      </div>
    </div>

    <!-- Pipeline Stepper -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
      <div v-for="(step, idx) in [
        { id: 's1', status: 'done', icon: LucideScanLine },
        { id: 's2', status: 'active', icon: LucideTarget, progress: 84 },
        { id: 's3', status: 'waiting', icon: LucideLanguages },
        { id: 's4', status: 'waiting', icon: LucideBadgeCheck }
      ]" :key="step.id" :class="`glass p-8 rounded-[3rem] space-y-6 shadow-sm border-2 transition-all ${step.status === 'active' ? 'border-blue-500 shadow-xl shadow-blue-500/10 scale-105' : 'border-slate-50 dark:border-slate-800'}`">
        <div class="flex items-center justify-between">
          <div :class="`w-12 h-12 rounded-2xl flex items-center justify-center transition-colors ${step.status === 'done' ? 'bg-green-500 text-white shadow-lg shadow-green-500/20' : step.status === 'active' ? 'bg-[#1E3A5F] text-white shadow-lg shadow-blue-900/20' : 'bg-slate-50 dark:bg-slate-800 text-slate-300'}`">
            <component :is="step.icon" class="w-6 h-6" />
          </div>
          <LucideMoreHorizontal v-if="step.status === 'active'" class="w-5 h-5 text-blue-500" />
          <LucideCheckCircle2 v-if="step.status === 'done'" class="w-5 h-5 text-green-500" />
        </div>
        <div class="space-y-2">
          <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t(`intake.processing.pipeline.${step.id}`) }}</h4>
          <p class="text-[9px] font-bold text-slate-400 uppercase leading-relaxed tracking-tight">
            {{ step.id === 's2' ? $t(`intake.processing.pipeline.${step.id}_desc`, { curr: 42, total: 50 }) : $t(`intake.processing.pipeline.${step.id}_desc`) }}
          </p>
        </div>
        <div class="pt-4 border-t border-slate-50 dark:border-slate-800 flex items-center justify-between">
          <span :class="`text-[8px] font-black uppercase tracking-widest ${step.status === 'done' ? 'text-green-500' : step.status === 'active' ? 'text-blue-500' : 'text-slate-300'}`">
            {{ $t(`intake.processing.pipeline.${step.status}`) }}
          </span>
          <span v-if="step.progress" class="text-[10px] font-black text-blue-500">{{ step.progress }}%</span>
        </div>
        <div v-if="step.progress" class="h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
          <div class="h-full bg-blue-500 shadow-sm" :style="{ width: step.progress + '%' }"></div>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Column: Result Quality -->
      <div class="lg:col-span-4 space-y-10">
        <div class="glass p-10 rounded-[3.5rem] space-y-10 shadow-sm border border-slate-50 dark:border-slate-800">
          <div class="flex items-center gap-4">
            <LucideBarChart3 class="w-6 h-6 text-blue-500" />
            <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('intake.processing.quality.title') }}</h3>
          </div>

          <div class="space-y-10">
            <div class="glass p-8 rounded-lg bg-slate-50/50 dark:bg-slate-900/50 border border-slate-100 dark:border-slate-800 flex flex-col items-center text-center space-y-3">
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.processing.quality.avg') }}</p>
              <div class="space-y-1">
                <p class="text-5xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">92.4%</p>
                <p class="text-[9px] font-black text-green-500 uppercase tracking-widest flex items-center justify-center gap-2">
                  <LucideTrendingUp class="w-3 h-3" />
                  {{ $t('intake.processing.quality.vs_prev', { val: 2.1 }) }}
                </p>
              </div>
            </div>

            <div class="grid grid-cols-2 gap-6">
              <div class="p-8 bg-slate-50/30 dark:bg-slate-900/30 rounded-3xl border border-slate-100 dark:border-slate-800 space-y-2">
                <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.processing.quality.pages') }}</p>
                <p class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">50</p>
              </div>
              <div class="p-8 bg-slate-50/30 dark:bg-slate-900/30 rounded-3xl border border-slate-100 dark:border-slate-800 space-y-2">
                <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.processing.quality.errors') }}</p>
                <p class="text-3xl font-black text-red-500 tracking-tighter">0</p>
              </div>
            </div>

            <div class="space-y-4">
              <div class="flex items-center justify-between text-[10px] font-black uppercase tracking-widest text-slate-400">
                <span>{{ $t('intake.processing.quality.chars') }}</span>
                <span class="text-[#1E3A5F] dark:text-white">142,802</span>
              </div>
              <div class="h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden shadow-inner">
                <div class="h-full bg-blue-500" style="width: 100%"></div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Column: Warnings -->
      <div class="lg:col-span-8 space-y-10">
        <div class="glass p-10 rounded-[3.5rem] space-y-10 shadow-sm border border-slate-50 dark:border-slate-800 h-full flex flex-col">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-4">
              <LucideAlertTriangle class="w-6 h-6 text-red-500" />
              <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('intake.processing.warnings.title') }}</h3>
            </div>
            <span class="px-6 py-1.5 bg-red-50 text-red-500 rounded-full text-[9px] font-black uppercase tracking-widest">{{ $t('intake.processing.warnings.issues', { count: 12 }) }}</span>
          </div>

          <div class="flex-grow">
            <table class="w-full text-left border-collapse">
              <thead>
                <tr class="bg-slate-50/50 dark:bg-slate-800/50 text-[9px] font-black text-slate-400 uppercase tracking-widest">
                  <th class="p-8 pl-12">{{ $t('intake.processing.warnings.cols.page') }}</th>
                  <th class="p-8">{{ $t('intake.processing.warnings.cols.field') }}</th>
                  <th class="p-8">{{ $t('intake.processing.warnings.cols.conf') }}</th>
                  <th class="p-8 pr-12 text-right">{{ $t('intake.processing.warnings.cols.action') }}</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
                <tr v-for="w in [
                  { page: '#04', field: 'Nomor Akta Notaris', context: 'Handwritten signature overlap', conf: 45 },
                  { page: '#12', field: 'Tanggal Terbit', context: 'Blurry timestamp in footer', conf: 62 },
                  { page: '#28', field: 'Alamat Domisili', context: 'Unrecognized character set', conf: 78 }
                ]" :key="w.page" class="group hover:bg-slate-50/30 transition-all cursor-pointer">
                  <td class="p-8 pl-12 text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tighter">{{ w.page }}</td>
                  <td class="p-8">
                    <div class="space-y-1">
                      <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ w.field }}</p>
                      <p class="text-[8px] font-bold text-slate-400 uppercase italic tracking-widest">"{{ w.context }}"</p>
                    </div>
                  </td>
                  <td class="p-8">
                    <div class="flex items-center gap-4">
                      <div class="w-20 h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
                        <div :class="`h-full ${w.conf < 50 ? 'bg-red-500' : 'bg-amber-500'}`" :style="{ width: w.conf + '%' }"></div>
                      </div>
                      <span :class="`text-[10px] font-black ${w.conf < 50 ? 'text-red-500' : 'text-amber-500'}`">{{ w.conf }}%</span>
                    </div>
                  </td>
                  <td class="p-8 pr-12 text-right">
                    <button class="text-[10px] font-black text-blue-500 uppercase tracking-widest hover:underline">{{ $t('intake.processing.warnings.inspect') }}</button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <div class="p-8 bg-slate-50/30 dark:bg-slate-800/30 text-center">
            <button class="text-[10px] font-black text-slate-400 uppercase tracking-widest flex items-center gap-2 mx-auto hover:text-[#1E3A5F] group transition-all">
              {{ $t('intake.processing.warnings.view_all', { count: 12 }) }}
              <LucideChevronDown class="w-4 h-4 group-hover:translate-y-1 transition-transform" />
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Footer Actions -->
    <div class="glass p-10 rounded-[3rem] shadow-sm border border-slate-50 dark:border-slate-800 bg-slate-50/30 dark:bg-slate-800/30 flex flex-col md:flex-row items-center justify-between gap-10">
      <div class="flex items-center gap-6">
        <button class="px-8 py-4 bg-white dark:bg-slate-800 border-2 border-slate-50 dark:border-slate-700 text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest rounded-2xl hover:bg-slate-50 transition-all flex items-center gap-3">
          <LucideRotateCcw class="w-4 h-4" />
          {{ $t('intake.processing.actions.rerun') }}
        </button>
        <button class="px-8 py-4 bg-white dark:bg-slate-800 border-2 border-slate-50 dark:border-slate-700 text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest rounded-2xl hover:bg-slate-50 transition-all flex items-center gap-3">
          <LucideDownload class="w-4 h-4" />
          {{ $t('intake.processing.actions.download') }}
        </button>
      </div>
      <button class="px-12 py-4 bg-[#1E3A5F] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-2xl shadow-blue-900/30 hover:bg-[#152943] transition-all flex items-center gap-4 group">
        {{ $t('intake.processing.actions.continue') }}
        <LucideArrowRight class="w-5 h-5 group-hover:translate-x-2 transition-transform" />
      </button>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideScanLine, LucideTarget, LucideLanguages, LucideBadgeCheck, 
  LucideMoreHorizontal, LucideCheckCircle2, LucideBarChart3, 
  LucideTrendingUp, LucideAlertTriangle, LucideChevronDown, 
  LucideRotateCcw, LucideDownload, LucideArrowRight 
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

