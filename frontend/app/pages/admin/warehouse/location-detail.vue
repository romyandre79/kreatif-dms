<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Column: Location Info -->
      <div class="lg:col-span-3 space-y-10">
        <div class="glass p-10 rounded-lg space-y-10 shadow-sm border border-slate-50 dark:border-slate-800">
          <div class="flex items-center justify-between">
            <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('warehouse.detail.title') }}</h3>
            <span class="px-4 py-1.5 bg-blue-50 text-blue-500 rounded-lg text-[9px] font-black uppercase tracking-widest">{{ $t('warehouse.detail.active') }}</span>
          </div>

          <div class="space-y-8">
            <div class="space-y-1.5">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.detail.id') }}</label>
              <p class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tight">F-15 / Row 18</p>
            </div>
            <div class="space-y-1.5">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.detail.zone') }}</label>
              <p class="text-xs font-bold text-slate-500 uppercase tracking-tight">Archives A - High Security</p>
            </div>
            <div class="space-y-1.5">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.detail.media') }}</label>
              <p class="text-xs font-bold text-slate-500 uppercase tracking-tight">Standard Box (30×40×25)</p>
            </div>
          </div>

          <div class="pt-8 border-t border-slate-50 dark:border-slate-800 space-y-8">
            <h4 class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('warehouse.detail.limits.title') }}</h4>
            <div class="space-y-6">
              <div v-for="l in ['max', 'soft', 'hard']" :key="l" class="flex items-center justify-between">
                <span class="text-[10px] font-bold text-slate-400 uppercase tracking-tight">{{ $t(`warehouse.detail.limits.${l}`) }}</span>
                <input type="text" :value="l === 'max' ? '150' : l === 'soft' ? '85' : '95'" class="w-16 h-10 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-xl text-center text-xs font-black text-[#1E3A5F] dark:text-white outline-none focus:border-blue-500 transition-all" />
              </div>
            </div>
          </div>

          <div class="pt-8 border-t border-slate-50 dark:border-slate-800 space-y-6">
            <div class="flex items-center justify-between">
              <h4 class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('warehouse.detail.status.title') }}</h4>
              <span class="text-[8px] font-black text-red-500 uppercase tracking-widest animate-pulse">{{ $t('warehouse.detail.status.at_capacity') }}</span>
            </div>
            <div class="space-y-3">
              <div class="h-2.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden shadow-inner">
                <div class="h-full bg-red-500 shadow-lg shadow-red-500/20" style="width: 94%"></div>
              </div>
              <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest text-center">{{ $t('warehouse.detail.status.util', { curr: 141, total: 150, val: 94 }) }}</p>
            </div>
          </div>

          <button class="w-full py-5 bg-[#1E3A5F] text-white rounded-2xl text-[11px] font-black uppercase tracking-widest shadow-2xl shadow-blue-900/30 hover:bg-[#152943] transition-all flex items-center justify-center gap-4 active:scale-95 group">
            <LucideSave class="w-5 h-5 group-hover:scale-110 transition-transform" />
            {{ $t('warehouse.detail.btn_save') }}
          </button>
        </div>
      </div>

      <!-- Center Column: Slot Configuration -->
      <div class="lg:col-span-6 space-y-10">
        <div class="glass rounded-lg overflow-hidden shadow-sm border border-slate-50 dark:border-slate-800 h-full flex flex-col">
          <div class="p-10 flex items-center justify-between border-b border-slate-50 dark:border-slate-800 bg-white/50 dark:bg-slate-900/50 backdrop-blur-md">
            <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('warehouse.detail.container.title') }}</h3>
            <div class="flex items-center gap-4">
              <LucideFilter class="w-4 h-4 text-slate-300 cursor-pointer" />
              <LucideDownload class="w-4 h-4 text-slate-300 cursor-pointer" />
            </div>
          </div>

          <div class="flex-grow overflow-y-auto custom-scrollbar">
            <table class="w-full text-left border-collapse">
              <thead>
                <tr class="bg-slate-50/50 dark:bg-slate-800/50 text-[9px] font-black text-slate-400 uppercase tracking-widest">
                  <th class="p-8 pl-12">{{ $t('warehouse.detail.container.cols.code') }}</th>
                  <th class="p-8">{{ $t('warehouse.detail.container.cols.box') }}</th>
                  <th class="p-8">{{ $t('warehouse.detail.container.cols.date') }}</th>
                  <th class="p-8 pr-12">{{ $t('warehouse.detail.container.cols.status') }}</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
                <tr v-for="slot in [
                  { code: 'F-15-A1', box: 'BOX-2023-00452', date: '12 Oct 2023', status: 'filled' },
                  { code: 'F-15-A2', box: null, date: '-', status: 'available' },
                  { code: 'F-15-A3', box: 'BOX-2023-00981', date: '14 Oct 2023', status: 'overflow', warn: true },
                  { code: 'F-15-A4', box: 'BOX-2023-01023', date: '15 Oct 2023', status: 'filled' },
                  { code: 'F-15-A5', box: 'BOX-2023-01114', date: '15 Oct 2023', status: 'filled' },
                  { code: 'F-15-A6', box: 'Maintenance', date: '-', status: 'locked' }
                ]" :key="slot.code" :class="`group hover:bg-slate-50/30 transition-all ${slot.status === 'overflow' ? 'bg-red-50/30' : ''}`">
                  <td class="p-8 pl-12 text-[11px] font-black text-blue-500 uppercase tracking-widest font-mono">{{ slot.code }}</td>
                  <td class="p-8">
                    <p :class="`text-[11px] font-bold uppercase tracking-tight ${slot.box && slot.box !== 'Maintenance' ? 'text-[#1E3A5F] dark:text-white' : 'text-slate-300'}`">
                      {{ slot.box || $t('warehouse.detail.container.empty') }}
                    </p>
                  </td>
                  <td class="p-8 text-[10px] font-bold text-slate-400 uppercase tracking-widest font-mono">{{ slot.date }}</td>
                  <td class="p-8 pr-12">
                    <div class="flex items-center gap-4">
                      <span :class="`px-3 py-1 rounded text-[8px] font-black uppercase tracking-widest ${
                        slot.status === 'filled' ? 'bg-green-50 text-green-500' :
                        slot.status === 'available' ? 'bg-slate-100 text-slate-400' :
                        slot.status === 'overflow' ? 'bg-red-100 text-red-500' :
                        'bg-amber-50 text-amber-500'
                      }`">
                        {{ $t(`warehouse.detail.container.${slot.status}`) }}
                      </span>
                      <LucideAlertTriangle v-if="slot.warn" class="w-3.5 h-3.5 text-red-500 animate-pulse" />
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- Right Column: AI & Incoming -->
      <div class="lg:col-span-3 space-y-10">
        <div class="space-y-6">
          <div class="flex items-center gap-3 px-4">
            <LucideZap class="w-4 h-4 text-blue-500 fill-current" />
            <h3 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('warehouse.detail.ai.title') }}</h3>
          </div>

          <div class="space-y-6">
            <!-- Optimal Relocation Card -->
            <div class="glass p-8 rounded-lg border-l-8 border-blue-500 space-y-6 shadow-xl shadow-blue-500/5">
              <div class="flex items-start gap-4">
                <div class="w-10 h-10 rounded-xl bg-blue-50 text-blue-500 flex items-center justify-center shrink-0">
                  <LucideRepeat class="w-5 h-5" />
                </div>
                <div class="space-y-2">
                  <h4 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('warehouse.detail.ai.relocation') }}</h4>
                  <p class="text-[10px] font-bold text-slate-500 uppercase leading-relaxed italic">{{ $t('warehouse.detail.ai.relo_desc') }}</p>
                </div>
              </div>
              <button class="w-full py-4 text-[10px] font-black text-blue-500 uppercase tracking-widest border-2 border-blue-100 rounded-xl hover:bg-blue-50 transition-all">
                {{ $t('warehouse.detail.ai.btn_move') }}
              </button>
            </div>

            <!-- Retention Alert Card -->
            <div class="glass p-8 rounded-lg border-l-8 border-amber-500 space-y-4">
              <div class="flex items-start gap-4">
                <div class="w-10 h-10 rounded-xl bg-amber-50 text-amber-500 flex items-center justify-center shrink-0">
                  <LucideCalendarClock class="w-5 h-5" />
                </div>
                <div class="space-y-2">
                  <h4 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('warehouse.detail.ai.retention') }}</h4>
                  <p class="text-[10px] font-bold text-slate-500 uppercase leading-relaxed italic">{{ $t('warehouse.detail.ai.reten_desc') }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="space-y-6 pt-6">
          <h3 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest px-4">{{ $t('warehouse.detail.incoming.title') }}</h3>
          <div class="space-y-4">
            <div v-for="inc in [
              { id: 'BOX-2023-01255', desc: 'Legal Docs / 24kg', icon: LucidePackage },
              { id: 'BOX-2023-01256', desc: 'Finance FY22 / 18kg', icon: LucidePackage },
              { id: 'BOX-2023-01257', desc: 'Human Resources / 12kg', icon: LucidePackage }
            ]" :key="inc.id" class="p-6 bg-slate-50/50 dark:bg-slate-800/50 rounded-2xl border border-slate-100 dark:border-slate-700 flex items-center gap-6 group hover:bg-slate-50 transition-all cursor-pointer">
              <div class="w-12 h-12 bg-white dark:bg-slate-900 rounded-xl border border-slate-100 dark:border-slate-700 flex items-center justify-center text-blue-500 shadow-sm group-hover:scale-110 transition-transform">
                <component :is="inc.icon" class="w-6 h-6" />
              </div>
              <div class="space-y-0.5">
                <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tighter">{{ inc.id }}</p>
                <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ inc.desc }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideSave, LucideFilter, LucideDownload, LucideAlertTriangle, 
  LucideZap, LucideRepeat, LucideCalendarClock, LucidePackage 
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
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  @apply bg-transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  @apply bg-slate-200 dark:bg-slate-800 rounded-full;
}
</style>

