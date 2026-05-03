<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="space-y-1">
      <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('intake.external.title') }}</h1>
      <p class="text-xs font-bold text-slate-500 uppercase tracking-widest">{{ $t('intake.external.subtitle') }}</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Column: Timeline & SOP -->
      <div class="lg:col-span-3 space-y-10">
        <div class="glass p-10 rounded-[3rem] space-y-10 shadow-sm border border-slate-50 dark:border-slate-800">
          <h3 class="text-[10px] font-black text-blue-500 uppercase tracking-widest">{{ $t('intake.external.timeline.title') }}</h3>
          <div class="space-y-8 relative">
            <div class="absolute left-4 top-2 bottom-2 w-0.5 bg-slate-100 dark:bg-slate-800"></div>
            <div v-for="(step, idx) in [
              { id: 's1', status: 'done', icon: LucideCheckCircle2 },
              { id: 's2', status: 'active', icon: LucideCircleDot },
              { id: 's3', status: 'pending', icon: LucideCircle },
              { id: 's4', status: 'pending', icon: LucideCircle }
            ]" :key="step.id" class="flex items-start gap-6 relative group">
              <div :class="`w-8 h-8 rounded-full flex items-center justify-center border-4 border-white dark:border-slate-900 shadow-sm z-10 transition-all ${step.status === 'done' ? 'bg-blue-500 text-white' : step.status === 'active' ? 'bg-white dark:bg-slate-800 text-blue-500 border-blue-500' : 'bg-white dark:bg-slate-800 text-slate-200'}`">
                <component :is="step.icon" class="w-4 h-4" />
              </div>
              <div class="space-y-1">
                <p :class="`text-[11px] font-black uppercase tracking-tight ${step.status === 'pending' ? 'text-slate-300' : 'text-[#1E3A5F] dark:text-white'}`">{{ $t(`intake.external.timeline.${step.id}`) }}</p>
                <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ $t(`intake.external.timeline.${step.id}_sub`) }}</p>
              </div>
            </div>
          </div>
        </div>

        <div class="p-10 bg-blue-50/50 dark:bg-blue-900/10 rounded-[3rem] border border-blue-100 dark:border-blue-800 space-y-6 relative overflow-hidden group">
          <LucideInfo class="absolute top-0 right-0 w-24 h-24 text-blue-500/10 -rotate-12 translate-x-6 -translate-y-6 group-hover:scale-110 transition-transform" />
          <div class="flex items-center gap-4">
            <div class="w-10 h-10 rounded-xl bg-blue-500 text-white flex items-center justify-center shadow-lg shadow-blue-500/20">
              <LucideInfo class="w-5 h-5" />
            </div>
            <h4 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('intake.external.sop.title') }}</h4>
          </div>
          <p class="text-[10px] font-bold text-slate-500 leading-relaxed uppercase tracking-tight">
            {{ $t('intake.external.sop.desc') }}
          </p>
        </div>
      </div>

      <!-- Center Column: Form & Scan -->
      <div class="lg:col-span-6 space-y-10">
        <!-- Registration Form -->
        <div class="glass p-10 rounded-[3rem] space-y-10 shadow-sm border border-slate-50 dark:border-slate-800">
          <div class="flex items-center justify-between">
            <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('intake.external.form.title') }}</h3>
            <span class="px-4 py-1.5 bg-blue-50 text-blue-500 rounded-lg text-[9px] font-black uppercase tracking-widest">{{ $t('intake.external.form.mandatory') }}</span>
          </div>

          <div class="grid grid-cols-2 gap-8">
            <div class="space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.external.form.sender') }}</label>
              <input type="text" value="PT. Global Digital Niaga" class="w-full h-14 px-6 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold outline-none focus:border-blue-500 transition-all uppercase" />
            </div>
            <div class="space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.external.form.ext_no') }}</label>
              <input type="text" :placeholder="$t('intake.external.form.ext_placeholder')" class="w-full h-14 px-6 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold outline-none focus:border-blue-500 transition-all" />
            </div>
            <div class="space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.external.form.date') }}</label>
              <div class="relative">
                <input type="text" value="10/24/2023" class="w-full h-14 px-6 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold outline-none" />
                <LucideCalendar class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300" />
              </div>
            </div>
            <div class="space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.external.form.recipient') }}</label>
              <input type="text" value="Direktur Keuangan" class="w-full h-14 px-6 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold outline-none uppercase" />
            </div>
            <div class="col-span-2 space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.external.form.subject') }}</label>
              <textarea class="w-full p-6 bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800 rounded-[2rem] text-xs font-bold leading-relaxed outline-none h-24 resize-none uppercase">Permohonan Kerjasama Infrastruktur Jaringan Tahap II - Wilayah Jabodetabek</textarea>
            </div>
          </div>
        </div>

        <!-- Scan Panel -->
        <div class="glass p-10 rounded-[3rem] space-y-10 shadow-sm border border-slate-50 dark:border-slate-800">
          <div class="flex items-center justify-between">
            <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('intake.external.scan.title') }}</h3>
            <div class="flex items-center gap-4">
              <LucideSearch class="w-4 h-4 text-slate-400 cursor-pointer" />
              <LucideSearch class="w-4 h-4 text-slate-400 cursor-pointer" />
            </div>
          </div>

          <div class="grid grid-cols-12 gap-8">
            <!-- Sidebar Preview -->
            <div class="col-span-4 space-y-6">
              <div class="relative rounded-2xl overflow-hidden border-4 border-blue-500 shadow-xl group cursor-pointer">
                <div class="aspect-[3/4] bg-slate-200 opacity-50"></div>
                <div class="absolute inset-0 flex flex-col items-center justify-center text-center p-4">
                  <LucideFileText class="w-8 h-8 text-[#1E3A5F] mb-2" />
                  <p class="text-[10px] font-black text-[#1E3A5F] uppercase">{{ $t('intake.external.scan.page', { curr: 1, total: 3 }) }}</p>
                  <p class="text-[8px] font-bold text-slate-400">FRONT_FACE.JPG</p>
                </div>
              </div>
              <div class="aspect-[3/4] bg-slate-50 dark:bg-slate-900 rounded-2xl border-2 border-dashed border-slate-100 dark:border-slate-800 flex flex-col items-center justify-center gap-4 text-center group cursor-pointer hover:bg-slate-100 transition-all">
                <LucidePlus class="w-6 h-6 text-slate-300 group-hover:text-[#1E3A5F]" />
                <div class="space-y-0.5">
                  <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest group-hover:text-[#1E3A5F]">Page 2</p>
                  <p class="text-[7px] font-bold text-slate-300 uppercase">{{ $t('intake.external.scan.add') }}</p>
                </div>
              </div>
            </div>

            <!-- Main Scan View -->
            <div class="col-span-8 bg-slate-50 dark:bg-slate-950 rounded-[2.5rem] p-10 h-[600px] overflow-y-auto custom-scrollbar relative border border-slate-100 dark:border-slate-800">
              <div class="bg-white shadow-2xl p-12 space-y-10 min-h-[1000px] relative">
                <!-- Watermark/Skeleton doc -->
                <div class="h-20 bg-slate-800/80 w-1/3 mb-10"></div>
                <div class="space-y-4 opacity-30">
                  <div class="h-4 bg-slate-400 w-full rounded"></div>
                  <div class="h-4 bg-slate-400 w-full rounded"></div>
                  <div class="h-4 bg-slate-400 w-5/6 rounded"></div>
                </div>
                <div class="space-y-4 pt-10">
                  <div class="h-4 bg-slate-800 w-full rounded"></div>
                  <div class="h-4 bg-slate-800 w-full rounded"></div>
                  <div class="h-4 bg-slate-800 w-full rounded"></div>
                  <div class="h-4 bg-slate-800 w-4/6 rounded"></div>
                </div>
                <!-- Overlay Rescan Button -->
                <div class="absolute inset-0 flex items-center justify-center pointer-events-none group">
                  <button class="px-8 py-4 bg-[#1E3A5F]/90 text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-2xl backdrop-blur-md pointer-events-auto hover:bg-[#152943] transition-all flex items-center gap-4 shadow-blue-900/40">
                    <LucideRotateCcw class="w-4 h-4" />
                    {{ $t('intake.external.scan.rescan') }}
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Column: Readiness & Actions -->
      <div class="lg:col-span-3 space-y-10">
        <div class="glass p-10 rounded-[3rem] space-y-10 shadow-sm border border-slate-50 dark:border-slate-800">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.external.readiness.title') }}</h3>
          
          <div class="space-y-8">
            <div class="p-6 bg-slate-50/50 dark:bg-slate-800/50 rounded-2xl border border-slate-100 dark:border-slate-800 space-y-4">
              <div class="flex items-center justify-between">
                <p class="text-[10px] font-bold text-slate-400 uppercase">{{ $t('intake.external.readiness.appending') }}</p>
                <span class="text-[8px] font-black text-green-500 uppercase tracking-widest">{{ $t('intake.external.readiness.ready') }}</span>
              </div>
              <div class="flex items-center gap-4">
                <div class="w-10 h-10 rounded-xl bg-white dark:bg-slate-900 flex items-center justify-center text-slate-400 shadow-sm">
                  <LucideDatabase class="w-5 h-5" />
                </div>
                <div class="space-y-0.5">
                  <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase">Metadata Mapping</p>
                  <p class="text-[8px] font-bold text-slate-400 uppercase">Target: SQL_DMS_PROD</p>
                </div>
              </div>
            </div>

            <div class="p-6 bg-slate-50/50 dark:bg-slate-800/50 rounded-2xl border border-slate-100 dark:border-slate-800 space-y-4">
              <div class="flex items-center justify-between">
                <p class="text-[10px] font-bold text-slate-400 uppercase">{{ $t('intake.external.readiness.task_inbox') }}</p>
                <span class="text-[8px] font-black text-amber-500 uppercase tracking-widest animate-pulse">{{ $t('intake.external.readiness.syncing') }}</span>
              </div>
              <div class="flex items-center gap-4">
                <div class="w-10 h-10 rounded-xl bg-white dark:bg-slate-900 flex items-center justify-center text-slate-400 shadow-sm">
                  <LucideInbox class="w-5 h-5" />
                </div>
                <div class="space-y-0.5">
                  <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase">Recipient Availability</p>
                  <p class="text-[8px] font-bold text-slate-400 uppercase">Queue: Dept_General_Affairs</p>
                </div>
              </div>
            </div>
          </div>

          <div class="space-y-6 pt-10 border-t border-slate-50 dark:border-slate-800">
            <div class="flex items-center justify-between px-2">
              <span class="text-[10px] font-bold text-slate-400 uppercase">{{ $t('intake.external.readiness.id') }}</span>
              <span class="text-[11px] font-black text-[#1E3A5F] dark:text-white font-mono">REG-2023-X9921</span>
            </div>
            <div class="flex items-center justify-between px-2">
              <span class="text-[10px] font-bold text-slate-400 uppercase">{{ $t('intake.external.readiness.priority') }}</span>
              <div class="flex items-center gap-2">
                <div class="w-2 h-2 bg-red-500 rounded-full animate-pulse shadow-sm shadow-red-500/20"></div>
                <span class="text-[9px] font-black text-red-500 uppercase tracking-widest">{{ $t('intake.external.readiness.high') }}</span>
              </div>
            </div>
          </div>
        </div>

        <div class="space-y-4">
          <button class="w-full py-5 bg-[#1E3A5F] text-white rounded-2xl text-[11px] font-black uppercase tracking-widest shadow-2xl shadow-blue-900/30 hover:bg-[#152943] transition-all flex items-center justify-center gap-4 group active:scale-95">
            <LucideSend class="w-5 h-5 group-hover:translate-x-1 group-hover:-translate-y-1 transition-transform" />
            {{ $t('intake.external.actions.publish') }}
          </button>
          <button class="w-full py-5 bg-slate-100 dark:bg-slate-800 text-[#1E3A5F] dark:text-white rounded-2xl text-[11px] font-black uppercase tracking-widest hover:bg-slate-200 transition-all flex items-center justify-center gap-4 active:scale-95">
            <LucideSave class="w-5 h-5" />
            {{ $t('intake.external.actions.draft') }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideCheckCircle2, LucideCircleDot, LucideCircle, LucideInfo, 
  LucideCalendar, LucideSearch, LucideFileText, LucidePlus, 
  LucideRotateCcw, LucideDatabase, LucideInbox, LucideSend, LucideSave 
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

