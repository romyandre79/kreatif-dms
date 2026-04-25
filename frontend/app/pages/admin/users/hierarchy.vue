<template>
  <div class="h-[calc(100vh-120px)] flex flex-col -m-10 overflow-hidden" v-motion-fade>
    <!-- Header: Filters & Alerts -->
    <div class="bg-white dark:bg-slate-900 border-b border-slate-100 dark:border-slate-800 p-4 px-10 flex flex-col md:flex-row md:items-center justify-between gap-6 z-20">
      <div class="flex flex-wrap items-center gap-10">
        <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.users.hierarchy.header.filter') }}</span>
        <div class="flex items-center gap-10">
          <div class="flex items-center gap-3 cursor-pointer group">
            <span class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight group-hover:text-primary-500 transition-colors">{{ $t('admin.users.hierarchy.header.dept') }}</span>
            <LucideChevronDown class="w-4 h-4 text-slate-300" />
          </div>
          <div class="flex items-center gap-3 cursor-pointer group">
            <span class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight group-hover:text-primary-500 transition-colors">{{ $t('admin.users.hierarchy.header.status') }}</span>
            <LucideChevronDown class="w-4 h-4 text-slate-300" />
          </div>
        </div>
      </div>
      <div class="flex items-center gap-8">
        <span class="text-[10px] font-black text-red-500 uppercase tracking-widest animate-pulse">{{ $t('admin.users.hierarchy.header.conflict') }}</span>
        <button class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest hover:underline">{{ $t('admin.users.hierarchy.header.save') }}</button>
      </div>
    </div>

    <div class="flex flex-grow overflow-hidden">
      <!-- Left Side: Organization Tree -->
      <div class="w-80 bg-slate-50 dark:bg-slate-950 border-r border-slate-200 dark:border-slate-800 flex flex-col">
        <div class="p-8">
          <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.users.hierarchy.tree.title') }}</h4>
        </div>
        <div class="flex-grow overflow-y-auto px-6 space-y-4 custom-scrollbar">
          <!-- IT Operations Tree -->
          <div class="space-y-4">
            <div class="flex items-center gap-3 cursor-pointer group">
              <LucideChevronDown class="w-4 h-4 text-slate-400 group-hover:text-[#1E3A5F]" />
              <LucideBuilding2 class="w-5 h-5 text-amber-500" />
              <span class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase">IT Operations Group</span>
            </div>
            <div class="ml-8 space-y-4 border-l-2 border-slate-100 dark:border-slate-800 pl-6 py-2">
              <div class="flex items-center gap-3 opacity-60">
                <LucideChevronDown class="w-3.5 h-3.5 text-slate-300" />
                <div class="w-7 h-7 rounded-lg bg-blue-50 text-blue-500 flex items-center justify-center"><LucideUser class="w-4 h-4" /></div>
                <div class="space-y-0.5">
                  <p class="text-[11px] font-black uppercase">Andi Wijaya</p>
                  <p class="text-[8px] font-bold text-slate-400 uppercase">Group Head</p>
                </div>
              </div>
              <div class="ml-6 space-y-4 border-l-2 border-slate-100 dark:border-slate-800 pl-6 py-2">
                <div class="flex items-center gap-3 p-3 bg-white dark:bg-slate-900 rounded-xl shadow-sm border border-slate-100 dark:border-slate-800">
                  <LucideChevronDown class="w-3.5 h-3.5 text-[#1E3A5F]" />
                  <div class="w-7 h-7 rounded-lg bg-[#1E3A5F] text-white flex items-center justify-center"><LucideUser class="w-4 h-4" /></div>
                  <div class="space-y-0.5">
                    <p class="text-[11px] font-black uppercase">Budi Santoso</p>
                    <p class="text-[8px] font-bold text-slate-400 uppercase">Section Head Infrastructure</p>
                  </div>
                </div>
                <div class="ml-6 space-y-4 border-l-2 border-slate-100 dark:border-slate-800 pl-6 py-2">
                  <div class="flex items-center gap-3 opacity-50"><p class="text-[11px] font-bold uppercase text-slate-400">Citra Lestari (Senior Staff)</p></div>
                  <div class="flex items-center gap-3">
                    <div class="w-2 h-2 rounded-full bg-red-500"></div>
                    <div class="space-y-0.5">
                      <p class="text-[11px] font-black uppercase text-red-500">Dedi Kurniawan</p>
                      <p class="text-[8px] font-bold text-red-400 uppercase tracking-widest">(Staff)</p>
                    </div>
                    <LucideAlertCircle class="w-3.5 h-3.5 text-red-400" />
                  </div>
                </div>
              </div>
            </div>
          </div>
          <!-- Finance & Accounting -->
          <div class="flex items-center gap-3 opacity-40 cursor-not-allowed">
            <LucideChevronRight class="w-4 h-4 text-slate-400" />
            <LucideBuilding2 class="w-5 h-5 text-slate-400" />
            <span class="text-sm font-black text-slate-500 uppercase">Finance & Accounting</span>
          </div>
        </div>
      </div>

      <!-- Center: Org Chart Workspace -->
      <div class="flex-grow bg-slate-50/30 dark:bg-slate-950/30 relative overflow-hidden flex items-center justify-center p-20">
        <div class="absolute top-10 left-10 flex gap-4 z-10">
          <button class="p-3 bg-white dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-xl shadow-sm text-slate-400 hover:text-[#1E3A5F] transition-all"><LucideSearch class="w-4 h-4" /></button>
          <button class="p-3 bg-white dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-xl shadow-sm text-slate-400 hover:text-[#1E3A5F] transition-all"><LucideZoomIn class="w-4 h-4" /></button>
          <button class="p-3 bg-white dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-xl shadow-sm text-slate-400 hover:text-[#1E3A5F] transition-all"><LucideMaximize2 class="w-4 h-4" /></button>
        </div>

        <!-- Simulated Org Chart Nodes -->
        <div class="relative w-full h-full flex items-center justify-center scale-90">
          <!-- Budi Santoso Node -->
          <div class="absolute top-1/4 -translate-y-1/2 left-1/4 w-64 p-6 bg-white dark:bg-slate-900 rounded-2xl shadow-2xl border-t-4 border-[#1E3A5F] space-y-4 z-20" v-motion-pop>
            <div class="flex items-center gap-4">
              <div class="w-12 h-12 rounded-xl bg-slate-100 dark:bg-slate-800 overflow-hidden"><LucideUser class="w-full h-full p-2 text-slate-300" /></div>
              <div class="space-y-0.5">
                <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Budi Santoso</p>
                <p class="text-[10px] font-bold text-slate-400 uppercase">Section Head</p>
              </div>
            </div>
          </div>

          <!-- Dedi Kurniawan Node -->
          <div class="absolute top-1/2 left-1/2 -translate-x-1/2 w-64 p-6 bg-white dark:bg-slate-900 rounded-2xl shadow-2xl border-2 border-red-500 space-y-4" v-motion-pop>
            <div class="flex items-center gap-4">
              <div class="w-12 h-12 rounded-xl bg-red-50 overflow-hidden"><LucideUser class="w-full h-full p-2 text-red-200" /></div>
              <div class="space-y-0.5">
                <p class="text-sm font-black text-red-500 uppercase tracking-tight">Dedi Kurniawan</p>
                <p class="text-[10px] font-bold text-red-400 uppercase">Staff</p>
              </div>
            </div>
            <div class="p-3 bg-red-50/50 rounded-xl border border-red-100">
              <p class="text-[8px] font-black text-red-400 uppercase tracking-widest text-center">Circular: Reports back to self</p>
            </div>
          </div>
          
          <!-- Connecting Lines (Visual Placeholder) -->
          <svg class="absolute inset-0 w-full h-full pointer-events-none opacity-10">
            <line x1="25%" y1="25%" x2="50%" y2="50%" stroke="currentColor" stroke-width="2" stroke-dasharray="8 8" />
          </svg>
        </div>
      </div>

      <!-- Right Side: Impact Preview -->
      <div class="w-[450px] bg-white dark:bg-slate-900 border-l border-slate-200 dark:border-slate-800 flex flex-col shadow-2xl z-10 overflow-hidden">
        <div class="flex-grow overflow-y-auto p-10 space-y-12 custom-scrollbar">
          <div class="flex items-center gap-4">
            <div class="w-10 h-10 rounded-xl bg-blue-50 text-blue-500 flex items-center justify-center"><LucideBarChart3 class="w-5 h-5" /></div>
            <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('admin.users.hierarchy.impact.title') }}</h3>
          </div>

          <!-- Impact Header Card -->
          <div class="p-10 bg-slate-50 dark:bg-slate-800/50 rounded-[3rem] border border-slate-100 dark:border-slate-800 space-y-10 relative overflow-hidden shadow-sm">
            <div class="absolute top-0 right-0 w-32 h-32 bg-slate-100/50 dark:bg-slate-800/30 -rotate-12 translate-x-10 -translate-y-10 rounded-[3rem] flex items-end justify-start p-6 text-slate-200"><LucideNetwork class="w-12 h-12" /></div>
            
            <div class="flex items-center gap-6">
              <div class="w-16 h-16 rounded-[1.5rem] bg-red-50 text-red-400 flex items-center justify-center border-2 border-red-100 shadow-lg shadow-red-500/5"><LucideAlertTriangle class="w-8 h-8" /></div>
              <div class="space-y-1">
                <h4 class="text-xl font-black text-slate-800 dark:text-white uppercase tracking-tight leading-none">Dedi Kurniawan</h4>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">Staff IT Support</p>
              </div>
            </div>

            <div class="grid grid-cols-2 gap-6 relative">
              <div class="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 w-8 h-8 bg-white dark:bg-slate-900 rounded-full flex items-center justify-center shadow-md z-10"><LucideArrowRight class="w-4 h-4 text-slate-300" /></div>
              <div class="p-6 bg-white dark:bg-slate-900 rounded-3xl border border-slate-100 dark:border-slate-800 space-y-2">
                <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.users.hierarchy.impact.current_sup') }}</p>
                <p class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase">Budi Santoso</p>
              </div>
              <div class="p-6 bg-blue-50 text-blue-500 rounded-3xl border border-blue-100 space-y-2">
                <p class="text-[8px] font-black uppercase tracking-widest opacity-60">{{ $t('admin.users.hierarchy.impact.new_sup') }}</p>
                <p class="text-xs font-black uppercase">Andi Wijaya</p>
              </div>
            </div>
          </div>

          <!-- Impacted Workflows -->
          <div class="space-y-8">
            <div class="flex items-center justify-between border-b border-slate-50 dark:border-slate-800 pb-4">
              <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.users.hierarchy.impact.workflows.title', { count: 3 }) }}</h4>
              <button class="text-[9px] font-black text-blue-500 uppercase tracking-widest hover:underline">{{ $t('admin.users.hierarchy.impact.workflows.sim') }}</button>
            </div>
            <div class="space-y-4">
              <div v-for="i in [1, 2, 3]" :key="i" :class="`p-6 rounded-2xl border flex items-center gap-5 transition-all cursor-pointer hover:scale-[1.02] ${i === 3 ? 'bg-red-50/50 border-red-100' : 'bg-slate-50/30 border-slate-50 dark:border-slate-800'}`">
                <div :class="`w-10 h-10 rounded-xl flex items-center justify-center ${i === 3 ? 'bg-white text-red-500 shadow-sm' : 'bg-white dark:bg-slate-800 text-slate-400 shadow-sm'}`">
                  <component :is="i === 1 ? LucideFileText : i === 2 ? LucideCreditCard : LucideShieldCheck" class="w-5 h-5" />
                </div>
                <div class="space-y-0.5">
                  <p :class="`text-[11px] font-black uppercase tracking-tight ${i === 3 ? 'text-red-600' : 'text-[#1E3A5F] dark:text-white'}`">{{ $t(`admin.users.hierarchy.impact.workflows.item_${i}.title`) }}</p>
                  <p :class="`text-[9px] font-bold uppercase ${i === 3 ? 'text-red-400' : 'text-slate-400'}`">{{ $t(`admin.users.hierarchy.impact.workflows.item_${i}.desc`) }}</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Reason Field -->
          <div class="space-y-4">
            <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.users.hierarchy.impact.reason.title') }}</h4>
            <textarea :placeholder="$t('admin.users.hierarchy.impact.reason.placeholder')" class="w-full h-32 px-6 py-5 bg-slate-50/50 dark:bg-slate-800/30 border border-slate-100 dark:border-slate-800 rounded-3xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all resize-none custom-scrollbar"></textarea>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideChevronDown, LucideBuilding2, LucideUser, LucideAlertCircle, 
  LucideSearch, LucideZoomIn, LucideMaximize2, LucideBarChart3, 
  LucideAlertTriangle, LucideArrowRight, LucideNetwork, LucideFileText, 
  LucideCreditCard, LucideShieldCheck, LucideChevronRight 
} from 'lucide-vue-next'

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #E2E8F0; border-radius: 10px; }
.dark .custom-scrollbar::-webkit-scrollbar-thumb { background: #1E293B; }
</style>
