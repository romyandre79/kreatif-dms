<template>
  <div class="max-w-[1600px] mx-auto space-y-10 pb-20">
    <!-- Header -->
    <div class="space-y-2" v-motion-fade>
      <h1 class="text-3xl font-black text-[#1E3A5F] tracking-tight uppercase">{{ $t('retention.upload_bast.title') }}</h1>
      <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">
        {{ $t('retention.upload_bast.subtitle') }}
      </p>
    </div>

    <div class="grid grid-cols-12 gap-10 items-start">
      <div class="col-span-6 space-y-8">
        <!-- Upload Dropzone -->
        <div class="glass p-10 rounded-[3rem] bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-8" v-motion-slide-left>
          <div class="flex items-center gap-3 text-[10px] font-black text-slate-400 uppercase tracking-widest">
            <LucideFilePlus class="w-4 h-4 text-primary-500" />
            {{ $t('retention.upload_bast.dropzone.title') }}
          </div>

          <div class="aspect-square border-2 border-dashed border-slate-100 rounded-[3rem] flex flex-col items-center justify-center text-center p-12 space-y-6 bg-slate-50/30 group hover:border-primary-200 transition-all cursor-pointer">
            <div class="w-20 h-20 bg-white rounded-lg flex items-center justify-center shadow-sm text-slate-200 group-hover:text-primary-500 transition-colors">
              <LucideUploadCloud class="w-10 h-10" />
            </div>
            <div class="space-y-2">
              <h3 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('retention.upload_bast.dropzone.label') }}</h3>
              <p class="text-xs font-bold text-slate-400 uppercase tracking-widest leading-loose max-w-[250px] mx-auto">
                {{ $t('retention.upload_bast.dropzone.desc') }}
              </p>
            </div>
            <button class="px-10 py-4 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest hover:bg-[#152943] transition-all shadow-lg shadow-blue-900/20">
              {{ $t('retention.upload_bast.dropzone.btn_select') }}
            </button>
          </div>
        </div>

        <!-- Validation Checklist -->
        <div class="glass p-10 rounded-[3rem] bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-8" v-motion-slide-visible-bottom>
          <div class="flex items-center gap-3 text-[10px] font-black text-slate-400 uppercase tracking-widest">
            <LucideClipboardCheck class="w-4 h-4 text-emerald-500" />
            {{ $t('retention.upload_bast.checklist.title') }}
          </div>

          <div class="space-y-4">
            <label class="flex items-start gap-6 p-6 bg-slate-50/50 border border-slate-50 rounded-2xl cursor-pointer group hover:bg-white hover:border-emerald-100 transition-all">
              <input type="checkbox" class="w-6 h-6 mt-1 rounded border-2 border-slate-200 text-emerald-500 focus:ring-emerald-500/10" />
              <div class="space-y-1">
                <p class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight group-hover:text-emerald-600 transition-colors">{{ $t('retention.upload_bast.checklist.sig_present') }}</p>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest leading-relaxed">{{ $t('retention.upload_bast.checklist.sig_desc') }}</p>
              </div>
            </label>

            <label class="flex items-start gap-6 p-6 bg-slate-50/50 border border-slate-50 rounded-2xl cursor-pointer group hover:bg-white hover:border-emerald-100 transition-all">
              <input type="checkbox" class="w-6 h-6 mt-1 rounded border-2 border-slate-200 text-emerald-500 focus:ring-emerald-500/10" />
              <div class="space-y-1">
                <p class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight group-hover:text-emerald-600 transition-colors">{{ $t('retention.upload_bast.checklist.id_match') }}</p>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest leading-relaxed">{{ $t('retention.upload_bast.checklist.id_desc') }}</p>
              </div>
            </label>

            <label class="flex items-start gap-6 p-6 bg-slate-50/50 border border-slate-50 rounded-2xl cursor-pointer group hover:bg-white hover:border-emerald-100 transition-all">
              <input type="checkbox" class="w-6 h-6 mt-1 rounded border-2 border-slate-200 text-emerald-500 focus:ring-emerald-500/10" />
              <div class="space-y-1">
                <p class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight group-hover:text-emerald-600 transition-colors">{{ $t('retention.upload_bast.checklist.pages_scanned') }}</p>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest leading-relaxed">{{ $t('retention.upload_bast.checklist.pages_desc') }}</p>
              </div>
            </label>
          </div>
        </div>
      </div>

      <!-- Right Side: File Preview -->
      <div class="col-span-6 space-y-8" v-motion-slide-right>
        <div class="glass p-10 rounded-[3rem] bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-8">
          <div class="flex items-center gap-3 text-[10px] font-black text-slate-400 uppercase tracking-widest">
            <LucideEye class="w-4 h-4 text-primary-500" />
            {{ $t('retention.upload_bast.preview.title') }}
          </div>

          <div class="aspect-[3/4] bg-slate-50 rounded-lg border border-slate-100 overflow-hidden relative group">
            <div class="absolute inset-0 flex flex-col items-center justify-center p-12 text-center space-y-6">
              <div class="w-32 h-40 bg-white rounded-xl shadow-xl border border-slate-50 relative p-6 flex flex-col items-center justify-center gap-4">
                <LucideFileText class="w-12 h-12 text-slate-100" />
                <div class="h-1.5 w-full bg-slate-50 rounded-full"></div>
                <div class="h-1.5 w-2/3 bg-slate-50 rounded-full"></div>
              </div>
              <div class="space-y-2">
                <p class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">BAST_2023_BATCH_44.pdf</p>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">2.4 MB • Uploaded 2 mins ago</p>
              </div>
            </div>
            <!-- Mock document background -->
            <div class="absolute inset-0 bg-gradient-to-t from-slate-50 to-transparent opacity-60"></div>
          </div>

          <div class="space-y-6">
            <button @click="navigateTo('/retention/purge-confirm')" class="w-full py-6 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-3xl text-xs font-black uppercase tracking-[0.2em] shadow-xl shadow-blue-900/20 transition-all flex items-center justify-center gap-4 group">
              {{ $t('retention.upload_bast.preview.btn_validate') }} <LucideArrowRight class="w-5 h-5 group-hover:translate-x-1 transition-transform" />
            </button>
            <p class="text-[9px] font-bold text-slate-300 text-center uppercase tracking-widest leading-loose">
              {{ $t('retention.upload_bast.preview.footer_hint') }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideFilePlus, LucideUploadCloud, LucideClipboardCheck, 
  LucideEye, LucideFileText, LucideArrowRight 
} from 'lucide-vue-next'
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
</style>
