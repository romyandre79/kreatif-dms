<template>
  <div class="flex h-[calc(100vh-theme(spacing.32))] gap-10" v-motion-fade>
    <!-- Left Section: Department List -->
    <div class="w-80 flex flex-col gap-10 shrink-0">
      <div class="flex items-center justify-between px-2">
        <h2 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('warehouse.departments.title') }}</h2>
        <LucideFilter class="w-5 h-5 text-blue-500 cursor-pointer" />
      </div>

      <div class="flex-grow space-y-6 overflow-y-auto pr-4 custom-scrollbar">
        <div v-for="d in [
          { name: 'Legal & Compliance', status: 'active', count: 14, active: true },
          { name: 'Finance & Tax', status: 'reserved', count: 22, resCount: 4 },
          { name: 'Human Resources', status: 'locked', count: 8 },
          { name: 'Procurement', status: 'standard', count: 45 },
          { name: 'IT Infrastructure', status: 'standard', count: 5 }
        ]" :key="d.name" :class="`glass p-8 rounded-[2.5rem] space-y-4 shadow-sm cursor-pointer transition-all border-l-8 ${d.active ? 'border-blue-500 bg-slate-50/50 scale-105' : 'border-transparent hover:bg-slate-50/30'}`">
          <div class="flex items-center justify-between">
            <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight leading-tight max-w-[150px]">{{ d.name }}</h4>
            <span :class="`px-3 py-1 rounded text-[8px] font-black uppercase tracking-widest ${d.status === 'active' ? 'bg-blue-50 text-blue-500' : d.status === 'reserved' ? 'bg-blue-50 text-blue-500' : d.status === 'locked' ? 'bg-red-50 text-red-500' : 'bg-slate-50 text-slate-400'}`">
              <LucideLock v-if="d.status === 'locked'" class="w-3 h-3 inline mr-1" />
              {{ $t(`warehouse.departments.statuses.${d.status}`, { count: d.resCount }) }}
            </span>
          </div>
          <div class="space-y-2">
            <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('warehouse.departments.zoned', { count: d.count }) }}</p>
            <div class="h-1 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
              <div :class="`h-full ${d.status === 'locked' ? 'bg-red-500' : 'bg-blue-500'}`" :style="{ width: (d.count / 50 * 100) + '%' }"></div>
            </div>
          </div>
        </div>
      </div>

      <button class="w-full py-5 bg-slate-100 dark:bg-slate-800 text-[#1E3A5F] dark:text-white rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-200 transition-all flex items-center justify-center gap-4 active:scale-95">
        <LucidePlus class="w-4 h-4" />
        {{ $t('warehouse.departments.add') }}
      </button>
    </div>

    <!-- Center Section: Warehouse Map -->
    <div class="flex-grow flex flex-col gap-8">
      <div class="glass p-10 rounded-[3rem] flex items-center justify-between shadow-sm border border-slate-50 dark:border-slate-800">
        <div class="flex items-center gap-8">
          <div v-for="l in ['legal', 'finance', 'generic', 'locked']" :key="l" class="flex items-center gap-3">
            <div :class="`w-6 h-6 rounded-lg ${l === 'legal' ? 'bg-blue-500 shadow-lg shadow-blue-500/20' : l === 'finance' ? 'bg-[#1E3A5F]' : l === 'generic' ? 'bg-amber-500' : 'bg-slate-200'} flex items-center justify-center` ">
              <LucideLock v-if="l === 'locked'" class="w-3 h-3 text-slate-400" />
            </div>
            <span class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t(`warehouse.departments.map.legend.${l}`) }}</span>
          </div>
        </div>
        <div class="flex items-center gap-6">
          <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.departments.map.layout', { id: '02-B' }) }}</span>
          <div class="flex gap-2">
            <LucideSearch class="w-4 h-4 text-slate-300" />
            <LucideSearch class="w-4 h-4 text-slate-300" />
          </div>
        </div>
      </div>

      <div class="flex-grow glass p-12 rounded-[4rem] flex flex-col justify-between shadow-sm border border-slate-50 dark:border-slate-800">
        <!-- Grid Map -->
        <div class="grid grid-cols-8 gap-4 flex-grow">
          <div v-for="r in Array.from({length: 64}, (_, i) => ({
            id: String.fromCharCode(65 + Math.floor(i / 8)) + (i % 8 + 1),
            type: i < 8 ? 'grey' : i < 16 ? 'gold' : i < 48 ? 'blue' : 'darkblue',
            locked: [4, 20, 28, 30, 48, 60].includes(i)
          }))" :key="r.id" :class="`aspect-square rounded-xl flex items-center justify-center text-[10px] font-black tracking-tighter transition-all hover:scale-110 cursor-pointer ${
            r.locked ? 'bg-slate-50 text-slate-200 border border-slate-100' :
            r.type === 'blue' ? 'bg-blue-500 text-white shadow-lg shadow-blue-500/20' :
            r.type === 'darkblue' ? 'bg-[#1E3A5F] text-white' :
            r.type === 'gold' ? 'bg-amber-500 text-white' :
            'bg-slate-100 text-slate-400'
          }`">
            <LucideLock v-if="r.locked" class="w-4 h-4" />
            <span v-else>{{ r.id }}</span>
          </div>
        </div>

        <div class="pt-10 border-t border-slate-50 dark:border-slate-800 flex items-center justify-between text-[10px] font-black text-slate-400 uppercase tracking-widest">
          <span>{{ $t('warehouse.departments.map.capacity', { count: '1,240' }) }}</span>
          <span>{{ $t('warehouse.departments.map.allocated', { val: 82 }) }}</span>
        </div>
      </div>
    </div>

    <!-- Right Section: Rule Editor -->
    <div class="w-96 space-y-10 shrink-0">
      <div class="glass p-10 rounded-[3rem] space-y-10 shadow-sm border border-slate-50 dark:border-slate-800">
        <div class="flex items-center gap-4">
          <LucideEdit class="w-6 h-6 text-[#1E3A5F]" />
          <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('warehouse.departments.editor.title') }}</h3>
        </div>

        <div class="space-y-8">
          <div class="space-y-4">
            <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.departments.editor.action') }}</label>
            <div class="flex p-1.5 bg-slate-50 dark:bg-slate-800 rounded-2xl h-14 border border-slate-100 dark:border-slate-800">
              <button class="flex-1 bg-[#1E3A5F] text-white rounded-xl text-[9px] font-black uppercase tracking-widest shadow-lg">{{ $t('warehouse.departments.editor.reserve') }}</button>
              <button class="flex-1 text-slate-400 rounded-xl text-[9px] font-black uppercase tracking-widest">{{ $t('warehouse.departments.editor.release') }}</button>
            </div>
          </div>

          <div class="space-y-4">
            <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.departments.editor.target') }}</label>
            <div class="relative group">
              <select class="w-full h-14 pl-6 pr-12 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold outline-none appearance-none cursor-pointer">
                <option>Legal & Compliance (ACTIVE)</option>
              </select>
              <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300 pointer-events-none" />
            </div>
          </div>

          <div class="space-y-4">
            <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.departments.editor.date') }}</label>
            <div class="relative">
              <input type="text" value="11/24/2023" class="w-full h-14 px-6 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold outline-none" />
              <LucideCalendar class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300" />
            </div>
          </div>

          <!-- Conflict Warning -->
          <div class="p-8 bg-red-50 border-l-4 border-red-500 rounded-2xl space-y-2">
            <div class="flex items-center gap-3">
              <LucideAlertTriangle class="w-4 h-4 text-red-500" />
              <h5 class="text-[10px] font-black text-red-600 uppercase tracking-widest">{{ $t('warehouse.departments.editor.conflict') }}</h5>
            </div>
            <p class="text-[10px] font-bold text-red-400 leading-relaxed italic uppercase">{{ $t('warehouse.departments.editor.conflict_desc') }}</p>
          </div>
        </div>

        <button class="w-full py-5 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40 hover:bg-[#152943] transition-all">
          {{ $t('warehouse.departments.editor.btn_apply') }}
        </button>
      </div>

      <!-- Active Schedule -->
      <div class="space-y-8">
        <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-2">{{ $t('warehouse.departments.schedule.title') }}</h4>
        <div class="space-y-6">
          <div v-for="s in [
            { id: 'FINANCE-R0D', date: '12 Dec', desc: 'Temporary Audit Space Allocation [Zone D]', color: 'border-blue-800 bg-[#1E3A5F]' },
            { id: 'LEGAL-M01', permanent: true, desc: 'Master Deed Storage (High-Sec Vault)', color: 'border-blue-500 bg-blue-500 shadow-xl shadow-blue-500/20' },
            { id: 'GEN-A02', date: '02 Jan', desc: 'Peak Season Buffer Allocation', color: 'border-amber-500 bg-amber-500' }
          ]" :key="s.id" :class="`p-8 rounded-[2.5rem] border-l-8 text-white space-y-3 ${s.color}`">
            <div class="flex items-center justify-between">
              <h5 class="text-[11px] font-black tracking-widest uppercase">{{ s.id }}</h5>
              <span class="text-[8px] font-black uppercase tracking-widest opacity-60">{{ s.permanent ? $t('warehouse.departments.schedule.permanent') : $t('warehouse.departments.schedule.expires', { date: s.date }) }}</span>
            </div>
            <p class="text-[10px] font-bold uppercase tracking-tight opacity-90">{{ s.desc }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideFilter, LucidePlus, LucideLock, LucideSearch, 
  LucideEdit, LucideChevronDown, LucideCalendar, LucideAlertTriangle 
} from 'lucide-vue-next'

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
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
