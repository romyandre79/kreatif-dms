<template>
  <div class="space-y-12 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-10">
      <div class="space-y-2 max-w-2xl">
        <h1 class="text-4xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('intake.path_review.title') }}</h1>
        <p class="text-sm font-bold text-slate-500 uppercase tracking-tight">{{ $t('intake.path_review.subtitle') }}</p>
      </div>
      <div class="glass p-6 px-10 rounded-[2.5rem] flex items-center gap-6 shadow-sm border border-slate-50 dark:border-slate-800">
        <div class="space-y-1">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.path_review.health') }}</p>
          <p class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tighter">{{ $t('intake.path_review.health_val', { curr: '84.2', total: '100' }) }}</p>
        </div>
        <div class="w-12 h-12 bg-blue-50 text-blue-500 rounded-2xl flex items-center justify-center shadow-inner">
          <LucideDatabase class="w-6 h-6" />
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Column: Folder Tree Preview -->
      <div class="lg:col-span-8 space-y-10">
        <div class="glass rounded-[4rem] overflow-hidden shadow-sm border border-slate-50 dark:border-slate-800 h-full flex flex-col">
          <div class="p-10 flex items-center justify-between border-b border-slate-50 dark:border-slate-800 bg-white/50 dark:bg-slate-900/50 backdrop-blur-md">
            <div class="flex items-center gap-6">
              <div class="w-12 h-12 bg-[#1E3A5F] text-white rounded-2xl flex items-center justify-center shadow-xl shadow-blue-900/20">
                <LucideNetwork class="w-6 h-6" />
              </div>
              <div class="space-y-1">
                <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('intake.path_review.target.title') }}</h3>
                <p class="text-[11px] font-bold text-blue-500 uppercase font-mono tracking-tighter">{{ $t('intake.path_review.target.formula') }}</p>
              </div>
            </div>
            <LucideRotateCcw class="w-5 h-5 text-slate-300 hover:text-blue-500 cursor-pointer transition-all" />
          </div>

          <!-- Tree View -->
          <div class="p-16 flex-grow space-y-8 bg-slate-50/30 dark:bg-slate-950/30">
            <div class="space-y-6">
              <!-- Root -->
              <div class="flex items-center gap-4">
                <LucideFolder class="w-10 h-10 text-amber-500 fill-current" />
                <div class="flex items-center gap-3">
                  <span class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">ARSIP</span>
                  <span class="px-2 py-0.5 bg-slate-100 text-slate-400 rounded text-[7px] font-black uppercase tracking-widest">{{ $t('intake.path_review.tree.root') }}</span>
                </div>
              </div>

              <!-- PT -->
              <div class="pl-16 space-y-6 border-l-2 border-slate-100 dark:border-slate-800">
                <div class="flex items-center gap-4 relative">
                  <div class="absolute -left-[66px] top-1/2 w-16 h-0.5 bg-slate-100 dark:border-slate-800"></div>
                  <LucideFolder class="w-8 h-8 text-amber-500 fill-current" />
                  <div class="flex items-center gap-3">
                    <span class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">AKIRADATA_PRATAMA</span>
                    <span class="text-[9px] font-bold text-blue-500 uppercase tracking-widest opacity-60">({PT})</span>
                  </div>
                </div>

                <!-- Dept -->
                <div class="pl-16 space-y-6 border-l-2 border-slate-100 dark:border-slate-800">
                  <div class="flex items-center gap-4 relative">
                    <div class="absolute -left-[66px] top-1/2 w-16 h-0.5 bg-slate-100 dark:border-slate-800"></div>
                    <LucideFolder class="w-8 h-8 text-amber-500 fill-current" />
                    <div class="flex items-center gap-3">
                      <span class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">FINANCE_TAX</span>
                      <span class="text-[9px] font-bold text-blue-500 uppercase tracking-widest opacity-60">({Dept})</span>
                    </div>
                  </div>

                  <!-- Year -->
                  <div class="pl-16 space-y-6 border-l-2 border-slate-100 dark:border-slate-800">
                    <div class="flex items-center gap-4 relative">
                      <div class="absolute -left-[66px] top-1/2 w-16 h-0.5 bg-slate-100 dark:border-slate-800"></div>
                      <LucideFolder class="w-8 h-8 text-amber-500 fill-current" />
                      <div class="flex items-center gap-3">
                        <span class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">2024</span>
                        <span class="text-[9px] font-bold text-blue-500 uppercase tracking-widest opacity-60">({tahun})</span>
                      </div>
                    </div>

                    <!-- New Folder (Type) -->
                    <div class="pl-16 space-y-6 border-l-2 border-slate-100 dark:border-slate-800">
                      <div class="flex items-center gap-4 relative group">
                        <div class="absolute -left-[66px] top-1/2 w-16 h-0.5 bg-slate-100 dark:border-slate-800"></div>
                        <LucideFolderPlus class="w-8 h-8 text-amber-500" />
                        <div class="flex items-center gap-3">
                          <span class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">INV3ICE_3N_</span>
                          <span class="text-[9px] font-bold text-blue-500 uppercase tracking-widest opacity-60">({tipe})</span>
                          <span class="px-3 py-1 bg-blue-50 text-blue-500 rounded-lg text-[7px] font-black uppercase tracking-widest border border-blue-100">{{ $t('intake.path_review.tree.new_folder') }}</span>
                        </div>
                      </div>

                      <!-- Document -->
                      <div class="pl-16 space-y-6">
                        <div class="flex items-center gap-6 p-8 bg-white dark:bg-slate-900 rounded-[2.5rem] border border-slate-100 dark:border-slate-800 shadow-xl shadow-slate-200/50 dark:shadow-black/50 relative">
                          <div class="absolute -left-[66px] top-1/2 w-16 h-0.5 bg-slate-100 dark:border-slate-800"></div>
                          <div class="w-12 h-12 bg-blue-50 text-blue-500 rounded-2xl flex items-center justify-center shadow-inner">
                            <LucideFileText class="w-6 h-6" />
                          </div>
                          <div class="space-y-1">
                            <p class="text-[13px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">INV-2024-00192_PT_AD.pdf</p>
                            <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">2.4 MB • {{ $t('intake.path_review.tree.verified') }}</p>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Column: Rules & Options -->
      <div class="lg:col-span-4 space-y-10">
        <div class="glass p-10 rounded-[3.5rem] space-y-10 shadow-sm border border-slate-50 dark:border-slate-800">
          <div class="flex items-center gap-4">
            <LucideFolders class="w-6 h-6 text-[#1E3A5F] dark:text-white" />
            <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('intake.path_review.rules.title') }}</h3>
          </div>

          <div class="space-y-8">
            <div class="space-y-4">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.path_review.rules.policy') }}</label>
              <div class="p-8 bg-white dark:bg-slate-800 border-2 border-slate-50 dark:border-slate-700 rounded-3xl space-y-4 shadow-sm">
                <div class="flex items-center gap-4">
                  <div class="w-8 h-8 bg-green-50 text-green-500 rounded-xl flex items-center justify-center">
                    <LucideShieldCheck class="w-4 h-4" />
                  </div>
                  <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('intake.path_review.rules.dept_policy', { name: 'FINANCE' }) }}</h4>
                </div>
                <p class="text-[10px] font-bold text-slate-400 uppercase leading-relaxed tracking-tight">
                  {{ $t('intake.path_review.rules.policy_desc', { standard: 'Financial Compliance (ISO-27001)' }) }}
                </p>
              </div>
            </div>

            <div class="space-y-4">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('intake.path_review.rules.formula_preview') }}</label>
              <div class="p-8 bg-[#1E3A5F] rounded-3xl space-y-4 shadow-2xl shadow-blue-900/40">
                <div class="flex flex-wrap items-center gap-3">
                  <span class="px-3 py-1.5 bg-blue-900/50 text-blue-300 rounded-lg text-[9px] font-black font-mono">/ARSIP/</span>
                  <span class="px-3 py-1.5 bg-blue-500 text-white rounded-lg text-[9px] font-black font-mono">SCORP</span>
                  <span class="text-blue-300 font-mono">/</span>
                  <span class="px-3 py-1.5 bg-blue-500 text-white rounded-lg text-[9px] font-black font-mono">SDEPT</span>
                  <span class="text-blue-300 font-mono">/</span>
                  <span class="px-3 py-1.5 bg-blue-500 text-white rounded-lg text-[9px] font-black font-mono">SYEAR</span>
                  <span class="text-blue-300 font-mono">/</span>
                  <span class="px-3 py-1.5 bg-blue-500 text-white rounded-lg text-[9px] font-black font-mono">SDOC_TYPE</span>
                </div>
              </div>
            </div>

            <div class="pt-8 border-t border-slate-50 dark:border-slate-800 space-y-8">
              <div class="flex items-center justify-between group cursor-pointer">
                <div class="space-y-1">
                  <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('intake.path_review.options.create') }}</p>
                  <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('intake.path_review.options.create_desc') }}</p>
                </div>
                <div class="w-14 h-7 bg-blue-500 rounded-full relative shadow-inner">
                  <div class="absolute right-1 top-1 w-5 h-5 bg-white rounded-full shadow-md"></div>
                </div>
              </div>
              <div class="flex items-center justify-between group cursor-pointer opacity-40">
                <div class="space-y-1">
                  <p class="text-[11px] font-black text-slate-400 uppercase tracking-tight">{{ $t('intake.path_review.options.retention') }}</p>
                  <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('intake.path_review.options.retention_desc') }}</p>
                </div>
                <div class="w-14 h-7 bg-slate-200 rounded-full relative">
                  <div class="absolute left-1 top-1 w-5 h-5 bg-white rounded-full"></div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Footer Actions -->
    <div class="flex flex-col md:flex-row items-center justify-between gap-10">
      <button class="flex items-center gap-4 text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest hover:-translate-x-2 transition-transform group">
        <LucideArrowLeft class="w-5 h-5 group-hover:scale-125 transition-transform" />
        {{ $t('intake.path_review.actions.back') }}
      </button>
      <button class="px-16 py-6 bg-[#1E3A5F] text-white rounded-[2rem] text-sm font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40 hover:bg-[#152943] transition-all flex items-center gap-6 active:scale-95 group">
        {{ $t('intake.path_review.actions.commit') }}
        <div class="w-10 h-10 bg-white/20 rounded-xl flex items-center justify-center group-hover:rotate-12 transition-transform">
          <LucideCheckCircle2 class="w-6 h-6" />
        </div>
      </button>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideDatabase, LucideNetwork, LucideRotateCcw, LucideFolder, 
  LucideFolderPlus, LucideFileText, LucideFolders, LucideShieldCheck, 
  LucideArrowLeft, LucideCheckCircle2 
} from 'lucide-vue-next'

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}
</style>
