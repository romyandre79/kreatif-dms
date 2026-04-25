<template>
  <div class="space-y-12 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="space-y-2">
      <div class="flex items-center gap-3 text-[10px] font-black text-primary-500 uppercase tracking-[0.3em]">
        <LucideSettings2 class="w-4 h-4" />
        {{ $t('admin.archival.numbering.header_tag') }}
      </div>
      <h1 class="text-4xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('admin.archival.numbering.title') }}</h1>
      <p class="text-xs font-bold text-slate-500 max-w-2xl leading-relaxed">{{ $t('admin.archival.numbering.subtitle') }}</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Column: Numbering & Collision -->
      <div class="lg:col-span-7 space-y-10">
        <div class="glass p-12 rounded-[3rem] space-y-12 relative overflow-hidden">
          <div class="absolute top-0 right-0 w-64 h-64 bg-[#1E3A5F]/5 -translate-y-1/2 translate-x-1/2 rounded-full flex items-end justify-start p-12 opacity-10">
            <LucideFingerprint class="w-32 h-32" />
          </div>
          
          <div class="flex items-center justify-between">
            <div class="space-y-1">
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.archival.numbering.numbering.label') }}</p>
              <h2 class="text-5xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">#AD-2024-08912</h2>
            </div>
            <div class="text-right space-y-2">
              <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.archival.numbering.numbering.status') }}</p>
              <span class="px-4 py-1 bg-green-50 text-green-500 rounded-full text-[9px] font-black uppercase tracking-widest border border-green-100">{{ $t('admin.archival.numbering.numbering.ready') }}</span>
            </div>
          </div>

          <div class="space-y-4">
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.archival.numbering.numbering.preview') }}</p>
            <div class="p-10 bg-slate-50 dark:bg-slate-900 border-l-4 border-[#1E3A5F] rounded-2xl">
              <p class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-widest leading-relaxed">
                [IT] . [AKD] . [INV] . [HARDWARE] . [2024] . [00452]
              </p>
            </div>
          </div>

          <div class="grid grid-cols-3 gap-6">
            <div v-for="d in [
              { id: 'schema', val: 'Standard F-09' },
              { id: 'increment', val: '+1 Sequential' },
              { id: 'type', val: 'PDF/A-3b' }
            ]" :key="d.id" class="p-6 bg-white dark:bg-slate-800 rounded-2xl border border-slate-100 dark:border-slate-700 space-y-1.5 shadow-sm">
              <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t(`admin.archival.numbering.numbering.details.${d.id}`) }}</p>
              <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ d.val }}</p>
            </div>
          </div>
        </div>

        <!-- Collision Checker -->
        <div class="p-10 bg-slate-50/50 dark:bg-slate-900/50 rounded-[3rem] border border-slate-100 dark:border-slate-800 space-y-8 shadow-inner">
          <div class="flex items-center gap-4">
            <LucideShieldCheck class="w-5 h-5 text-[#1E3A5F] dark:text-white" />
            <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('admin.archival.numbering.collision.title') }}</h3>
          </div>
          <div class="space-y-4">
            <div class="p-6 bg-green-50/50 border border-green-100 rounded-2xl flex items-center justify-between">
              <div class="flex items-center gap-4">
                <LucideCheckCircle2 class="w-5 h-5 text-green-500" />
                <span class="text-xs font-bold text-green-600">{{ $t('admin.archival.numbering.collision.primary') }}</span>
              </div>
              <span class="text-[8px] font-black text-green-400 uppercase">Checked 0.2ms ago</span>
            </div>
            <div class="p-6 bg-slate-100/50 border border-slate-200 rounded-2xl flex items-center gap-4 text-slate-500">
              <LucideDatabase class="w-5 h-5" />
              <span class="text-xs font-bold uppercase tracking-tight">{{ $t('admin.archival.numbering.collision.secondary') }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Column: Folder Path Preview -->
      <div class="lg:col-span-5 bg-slate-100/50 dark:bg-slate-900/50 rounded-[3rem] p-10 border border-slate-100 dark:border-slate-800 flex flex-col gap-10">
        <div class="space-y-8 flex-grow">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-4">
              <LucideFolderTree class="w-6 h-6 text-[#1E3A5F]" />
              <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('admin.archival.numbering.path.title') }}</h3>
            </div>
            <span class="text-[9px] font-black text-primary-500 uppercase tracking-widest">{{ $t('admin.archival.numbering.path.schema') }}</span>
          </div>

          <div class="space-y-6 overflow-y-auto px-4 py-2 custom-scrollbar">
            <!-- Simulated Folder Tree -->
            <div class="flex items-center gap-3 opacity-30">
              <LucideFolder class="w-4 h-4" />
              <span class="text-[10px] font-bold uppercase">/ ARSIP</span>
            </div>
            <div class="ml-6 flex items-center gap-3 opacity-40">
              <LucideChevronRight class="w-3.5 h-3.5" />
              <LucideFolderOpen class="w-4 h-4" />
              <span class="text-[10px] font-bold uppercase">AKIRADATA</span>
            </div>
            <div class="ml-12 flex items-center gap-3 opacity-60">
              <LucideChevronRight class="w-3.5 h-3.5" />
              <LucideUsers class="w-4 h-4" />
              <span class="text-[10px] font-bold uppercase tracking-tight">INFORMATION TECHNOLOGY</span>
            </div>
            <div class="ml-20 flex items-center gap-3 opacity-80">
              <LucideChevronRight class="w-3.5 h-3.5" />
              <LucideCalendar class="w-4 h-4" />
              <span class="text-[10px] font-black uppercase">2024</span>
            </div>
            <div class="ml-28 flex items-center gap-3">
              <LucideChevronRight class="w-3.5 h-3.5 text-slate-300" />
              <LucideDatabase class="w-4 h-4 text-slate-400" />
              <span class="text-[10px] font-black uppercase text-slate-500">INVOICE</span>
            </div>
            <div class="ml-36 p-4 bg-[#1E3A5F]/5 border-l-4 border-[#1E3A5F] rounded-xl flex items-center gap-3" v-motion-slide-right>
              <LucideTag class="w-4 h-4 text-[#1E3A5F]" />
              <span class="text-[10px] font-black uppercase text-[#1E3A5F] tracking-widest">[TAG] VENDOR_PAYMENT</span>
            </div>
          </div>
        </div>

        <div class="space-y-4">
          <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('admin.archival.numbering.path.rules.title') }}</p>
          <div class="flex flex-wrap gap-2">
            <span v-for="rule in ['year', 'dept', 'tag']" :key="rule" class="px-4 py-1.5 bg-white dark:bg-slate-800 text-[#1E3A5F] dark:text-white rounded-lg text-[8px] font-black uppercase tracking-widest border border-slate-100 dark:border-slate-700 shadow-sm">
              {{ $t(`admin.archival.numbering.path.rules.${rule}`) }}
            </span>
          </div>
        </div>
      </div>
    </div>

    <!-- Final Confirmation Bar -->
    <div class="bg-white dark:bg-slate-900 p-10 rounded-[3rem] border border-slate-100 dark:border-slate-800 shadow-2xl flex flex-col md:flex-row items-center justify-between gap-10">
      <div class="flex items-center gap-8">
        <div class="w-20 h-20 rounded-2xl bg-slate-100 dark:bg-slate-800 overflow-hidden shadow-inner flex items-center justify-center">
          <LucidePrinter class="w-10 h-10 text-slate-300" />
        </div>
        <div class="space-y-1">
          <h3 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('admin.archival.numbering.confirm.title') }}</h3>
          <p class="text-xs font-bold text-slate-400 uppercase">{{ $t('admin.archival.numbering.confirm.desc') }}</p>
        </div>
      </div>
      <div class="flex items-center gap-6">
        <button class="px-10 py-5 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-slate-600 dark:text-slate-300 rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-slate-50 transition-all">
          {{ $t('admin.archival.numbering.confirm.btn_review') }}
        </button>
        <button class="px-16 py-5 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40 hover:bg-[#152943] transition-all flex items-center gap-4 active:scale-95 group">
          {{ $t('admin.archival.numbering.confirm.btn_commit') }}
          <LucideArrowRight class="w-5 h-5 group-hover:translate-x-1 transition-transform" />
        </button>
      </div>
    </div>

    <!-- Metadata Verification Summary -->
    <div class="space-y-6">
      <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-4">{{ $t('admin.archival.numbering.summary.title') }}</h4>
      <div class="bg-slate-50/50 dark:bg-slate-950/50 p-8 rounded-[2.5rem] border border-slate-100 dark:border-slate-800 grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-8">
        <div v-for="(v, k) in { entity: 'AKD Corp', division: 'IT Services', type: 'Hardware Inv', date: '24 Oct 2024', user: 'Admin_Cent', revision: 'v1.0 (Final)' }" :key="k" class="space-y-1">
          <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">{{ $t(`admin.archival.numbering.summary.cols.${k}`) }}</p>
          <p class="text-[11px] font-black text-slate-600 dark:text-slate-300 uppercase tracking-tight">{{ v }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideSettings2, LucideFingerprint, LucideShieldCheck, LucideCheckCircle2, 
  LucideDatabase, LucideFolderTree, LucideFolder, LucideFolderOpen, 
  LucideChevronRight, LucideUsers, LucideCalendar, LucideTag, 
  LucidePrinter, LucideArrowRight 
} from 'lucide-vue-next'

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #E2E8F0; border-radius: 10px; }
.dark .custom-scrollbar::-webkit-scrollbar-thumb { background: #1E293B; }
</style>
