<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-10">
      <div class="space-y-1">
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('warehouse.override.title') }}</h1>
        <p class="text-xs font-bold text-slate-500 uppercase tracking-widest">{{ $t('warehouse.override.subtitle') }}</p>
      </div>
      <div class="flex items-center gap-4">
        <div class="glass px-6 py-3 rounded-2xl flex items-center gap-3 border border-slate-50">
          <LucideDatabase class="w-4 h-4 text-blue-500" />
          <span class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('warehouse.override.list.count', { count: '124' }) }}</span>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Column: Rack List & Selection -->
      <div class="lg:col-span-5 space-y-8">
        <div class="glass rounded-[3.5rem] overflow-hidden shadow-sm border border-slate-50 dark:border-slate-800">
          <div class="p-8 border-b border-slate-50 dark:border-slate-800 bg-slate-50/30">
            <div class="relative">
              <LucideSearch class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
              <input 
                type="text" 
                placeholder="Search rack ID (e.g. A1-01)..." 
                class="w-full bg-white dark:bg-slate-900 border-none rounded-2xl py-3 pl-12 pr-4 text-xs font-bold focus:ring-2 focus:ring-blue-500 outline-none"
              />
            </div>
          </div>

          <div class="max-h-[600px] overflow-y-auto custom-scrollbar">
            <div 
              v-for="rack in [
                { id: 'RACK-A1-01', filled: 85, status: 'auto', type: 'Legal Docs' },
                { id: 'RACK-A1-02', filled: 100, status: 'manual', type: 'Finance Archives' },
                { id: 'RACK-B2-09', filled: 40, status: 'auto', type: 'General' },
                { id: 'RACK-C3-15', filled: 92, status: 'locked', type: 'Confidential' },
                { id: 'RACK-D4-04', filled: 0, status: 'auto', type: 'Unassigned' }
              ]" 
              :key="rack.id"
              class="p-6 border-b border-slate-50 dark:border-slate-800 hover:bg-slate-50/50 cursor-pointer transition-all group"
            >
              <div class="flex items-center justify-between">
                <div class="flex items-center gap-4">
                  <div class="w-12 h-12 rounded-2xl bg-white dark:bg-slate-800 flex items-center justify-center shadow-inner group-hover:scale-110 transition-transform">
                    <LucideBox class="w-6 h-6 text-[#1E3A5F] dark:text-white" />
                  </div>
                  <div class="space-y-1">
                    <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ rack.id }}</p>
                    <p class="text-[9px] font-bold text-slate-400 uppercase">{{ rack.type }}</p>
                  </div>
                </div>
                <div class="text-right space-y-1">
                  <span :class="`px-3 py-1 rounded-lg text-[8px] font-black uppercase tracking-widest ${
                    rack.status === 'auto' ? 'bg-green-50 text-green-600' : 
                    rack.status === 'manual' ? 'bg-amber-50 text-amber-600' : 
                    'bg-red-50 text-red-600'
                  }`">
                    {{ rack.status === 'auto' ? $t('warehouse.override.list.normal') : (rack.status === 'manual' ? $t('warehouse.override.list.full') : 'LOCKED') }}
                  </span>
                  <p class="text-[9px] font-bold text-slate-500 uppercase">{{ $t('warehouse.override.list.filled', { val: rack.filled }) }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Column: Detail & Configuration -->
      <div class="lg:col-span-7 space-y-10">
        <div class="glass p-12 rounded-[4rem] space-y-12 shadow-sm border border-slate-50 dark:border-slate-800">
          <!-- Selection Header -->
          <div class="flex items-center justify-between">
            <div class="space-y-2">
              <h3 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('warehouse.override.detail.title') }}</h3>
              <div class="flex items-center gap-3">
                <span class="text-[10px] font-black text-blue-500 uppercase tracking-widest">RACK-A1-02</span>
                <span class="w-1 h-1 bg-slate-300 rounded-full"></span>
                <span class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">Zone A, Section 1</span>
              </div>
            </div>
            <LucideLock class="w-8 h-8 text-amber-500 opacity-20" />
          </div>

          <!-- Quick Stats Grid -->
          <div class="grid grid-cols-3 gap-6">
            <div v-for="stat in [
              { label: $t('warehouse.override.detail.stats.sensor'), val: 'ACTIVE', color: 'text-green-500' },
              { label: $t('warehouse.override.detail.stats.capacity'), val: '100%', color: 'text-[#1E3A5F]' },
              { label: 'LAST UPDATED', val: '2m ago', color: 'text-slate-400' }
            ]" :key="stat.label" class="bg-slate-50/50 dark:bg-slate-900/50 p-6 rounded-3xl space-y-1 border border-slate-100 dark:border-slate-800">
              <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ stat.label }}</p>
              <p :class="`text-sm font-black uppercase ${stat.color}`">{{ stat.val }}</p>
            </div>
          </div>

          <!-- Configuration Form -->
          <div class="space-y-10 pt-4">
            <div class="space-y-4">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest flex items-center justify-between">
                {{ $t('warehouse.override.config.title') }}
                <span class="px-3 py-1 bg-amber-500 text-white rounded-lg text-[8px] tracking-tight">{{ $t('warehouse.override.config.lock_hint') }}</span>
              </label>
              <div class="grid grid-cols-2 gap-6">
                <button class="p-6 rounded-lg border-2 border-slate-50 dark:border-slate-800 flex flex-col items-center text-center space-y-3 hover:border-blue-500 transition-all group">
                  <LucideZap class="w-6 h-6 text-slate-300 group-hover:text-blue-500" />
                  <span class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase">{{ $t('warehouse.override.list.auto') }}</span>
                </button>
                <button class="p-6 rounded-lg border-2 border-amber-500 bg-amber-50/30 flex flex-col items-center text-center space-y-3 group">
                  <LucideShieldAlert class="w-6 h-6 text-amber-500" />
                  <span class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase">{{ $t('warehouse.override.list.full') }}</span>
                </button>
              </div>
            </div>

            <!-- Recommendation Engine Toggle -->
            <div class="flex items-center justify-between p-8 rounded-lg bg-slate-50/30 dark:bg-slate-900/30 border border-slate-100 dark:border-slate-800">
              <div class="space-y-1">
                <h4 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('warehouse.override.config.engine') }}</h4>
                <p class="text-[9px] font-bold text-slate-400 uppercase max-w-sm">{{ $t('warehouse.override.config.engine_desc') }}</p>
              </div>
              <div class="w-14 h-8 bg-blue-500 rounded-full p-1 cursor-pointer">
                <div class="w-6 h-6 bg-white rounded-full ml-auto shadow-sm"></div>
              </div>
            </div>

            <!-- Reason Audit Trail -->
            <div class="space-y-4">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('warehouse.override.config.reason') }}</label>
              <textarea 
                :placeholder="$t('warehouse.override.config.reason_placeholder')"
                class="w-full bg-slate-50/50 dark:bg-slate-900/50 border-2 border-slate-100 dark:border-slate-800 rounded-lg p-6 text-xs font-bold focus:ring-2 focus:ring-blue-500 outline-none min-h-[120px] custom-scrollbar"
              ></textarea>
            </div>

            <!-- Action Buttons -->
            <div class="flex items-center gap-6">
              <button class="flex-grow py-5 bg-[#1E3A5F] text-white rounded-3xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all">
                {{ $t('warehouse.override.config.btn_apply') }}
              </button>
              <button class="px-10 py-5 bg-white border-2 border-slate-100 text-slate-400 rounded-3xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all">
                {{ $t('warehouse.override.config.btn_cancel') }}
              </button>
            </div>
          </div>
        </div>

        <!-- Alert Notification Box -->
        <div class="glass p-8 rounded-[3rem] border border-amber-100 bg-amber-50/20 flex items-start gap-6">
          <div class="w-12 h-12 rounded-2xl bg-amber-500/10 flex items-center justify-center text-amber-600">
            <LucideInfo class="w-6 h-6" />
          </div>
          <div class="space-y-1">
            <h4 class="text-[10px] font-black text-amber-600 uppercase tracking-widest">{{ $t('warehouse.override.config.alert') }}</h4>
            <p class="text-[9px] font-bold text-amber-500/80 uppercase tracking-tight leading-relaxed">
              {{ $t('warehouse.override.config.desc') }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideDatabase, LucideSearch, LucideBox, LucideLock, 
  LucideZap, LucideShieldAlert, LucideInfo 
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

