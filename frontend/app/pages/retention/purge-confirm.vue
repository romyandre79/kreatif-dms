<template>
  <div class="max-w-[1600px] mx-auto space-y-12 pb-20 relative">
    <!-- Header -->
    <div class="space-y-4" v-motion-fade>
      <div class="flex items-center gap-2">
        <span class="px-3 py-1 bg-slate-100 text-slate-500 rounded-full text-[9px] font-black uppercase tracking-widest">
          {{ $t('retention.purge_confirm.header_actor') }}
        </span>
      </div>
      <div class="space-y-2">
        <h1 class="text-3xl font-black text-[#1E3A5F] tracking-tight uppercase">{{ $t('retention.purge_confirm.title') }}</h1>
        <p class="text-slate-500 font-medium">{{ $t('retention.purge_confirm.subtitle') }}</p>
      </div>
    </div>

    <!-- Success Banner -->
    <div class="p-8 bg-emerald-50 border border-emerald-100 rounded-3xl flex items-center gap-6 shadow-xl shadow-emerald-900/5" v-motion-slide-top>
      <div class="w-12 h-12 bg-emerald-500 text-white rounded-2xl flex items-center justify-center shadow-lg shadow-emerald-500/20">
        <LucideCheckCircle2 class="w-6 h-6" />
      </div>
      <div class="space-y-1">
        <h3 class="text-sm font-black text-emerald-900 uppercase tracking-tight">{{ $t('retention.purge_confirm.banner.title') }}</h3>
        <p class="text-[11px] font-bold text-emerald-700/70 uppercase tracking-widest">{{ $t('retention.purge_confirm.banner.desc') }}</p>
      </div>
    </div>

    <!-- Main Execution Card -->
    <div class="glass p-12 rounded-[3.5rem] bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 flex gap-12 items-center" v-motion-slide-visible-bottom>
      <div class="w-1/3 aspect-[4/3] bg-slate-100 rounded-lg relative overflow-hidden group">
        <img src="https://images.unsplash.com/photo-1568667256549-094345857637?auto=format&fit=crop&q=80&w=400" 
             class="w-full h-full object-cover opacity-60 mix-blend-multiply group-hover:scale-110 transition-transform duration-1000" />
        <div class="absolute inset-0 flex items-center justify-center">
          <div class="w-20 h-20 bg-white/90 backdrop-blur-sm rounded-lg shadow-2xl flex items-center justify-center text-slate-400">
            <LucideFolderMinus class="w-10 h-10" />
          </div>
        </div>
      </div>
      
      <div class="flex-1 space-y-8">
        <div class="space-y-3">
          <p class="text-[11px] font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('retention.purge_confirm.content.status') }}</p>
          <h2 class="text-2xl font-black text-[#1E3A5F] uppercase tracking-tight leading-tight">
            {{ $t('retention.purge_confirm.content.title') }}
          </h2>
          <p class="text-sm font-medium text-slate-400 leading-relaxed max-w-[600px]">
            {{ $t('retention.purge_confirm.content.desc') }}
          </p>
        </div>

        <button @click="showModal = true" class="px-12 py-6 bg-red-500 hover:bg-red-600 text-white rounded-3xl text-xs font-black uppercase tracking-[0.2em] shadow-2xl shadow-red-900/20 transition-all flex items-center gap-4 group">
          <LucideTrash2 class="w-5 h-5 group-hover:scale-110 transition-transform" /> {{ $t('retention.purge_confirm.content.btn_initiate') }}
        </button>
      </div>
    </div>

    <!-- Modal Preview Area -->
    <div class="pt-20 space-y-12" v-motion-fade>
      <div class="flex flex-col items-center gap-4">
        <div class="w-1.5 h-1.5 bg-red-400 rounded-full animate-ping"></div>
        <p class="text-[10px] font-black text-red-400 uppercase tracking-[0.4em] text-center">{{ $t('retention.purge_confirm.preview_label') || 'PURGE CONFIRMATION MODAL PREVIEW' }}</p>
      </div>
      
      <div class="flex items-center justify-center p-12 bg-slate-50/30 rounded-[4rem] border border-dashed border-slate-200">
        <!-- The Preview Card -->
        <div class="glass max-w-lg w-full bg-white rounded-[3.5rem] p-16 shadow-2xl border border-red-50 flex flex-col items-center text-center space-y-12 relative overflow-hidden" v-motion-pop>
          <div class="absolute top-0 left-0 right-0 h-2 bg-gradient-to-r from-red-400 to-red-600"></div>
          
          <div class="w-24 h-24 bg-red-50 rounded-full flex items-center justify-center text-red-500 shadow-xl shadow-red-500/10">
            <LucideAlertTriangle class="w-12 h-12" />
          </div>

          <div class="space-y-6">
            <h2 class="text-3xl font-black text-[#1E3A5F] tracking-tight">{{ $t('retention.purge_confirm.modal.title') }}</h2>
            <p class="text-sm font-bold text-slate-500 leading-relaxed px-4" v-html="$t('retention.purge_confirm.modal.desc')"></p>
          </div>

          <div class="w-full flex gap-6">
            <button class="flex-1 py-6 bg-white border-2 border-slate-100 text-slate-400 rounded-2xl text-[11px] font-black uppercase tracking-widest hover:bg-slate-50 hover:border-slate-200 transition-all">
              {{ $t('retention.purge_confirm.modal.btn_cancel') }}
            </button>
            <button class="flex-1 py-6 bg-red-500 hover:bg-red-600 text-white rounded-2xl text-[11px] font-black uppercase tracking-widest shadow-xl shadow-red-900/20 transition-all">
              {{ $t('retention.purge_confirm.modal.btn_confirm') }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Real Functional Modal -->
    <Transition name="scale">
      <div v-if="showModal" class="fixed inset-0 z-[100] flex items-center justify-center p-6 bg-[#1E3A5F]/60 backdrop-blur-md">
        <div class="glass max-w-lg w-full bg-white rounded-[3rem] p-12 shadow-2xl border border-red-50 flex flex-col items-center text-center space-y-10" v-motion-pop>
          <div class="w-24 h-24 bg-red-50 rounded-full flex items-center justify-center text-red-500">
            <LucideAlertTriangle class="w-12 h-12" />
          </div>

          <div class="space-y-6">
            <h2 class="text-2xl font-black text-[#1E3A5F]">{{ $t('retention.purge_confirm.modal.title') }}</h2>
            <p class="text-[13px] font-bold text-slate-500 leading-relaxed px-4" v-html="$t('retention.purge_confirm.modal.desc')"></p>
          </div>

          <div class="w-full flex gap-4">
            <button @click="showModal = false" class="flex-1 py-5 bg-white border border-slate-200 text-slate-400 rounded-2xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all">
              {{ $t('retention.purge_confirm.modal.btn_cancel') }}
            </button>
            <button @click="handleFinalPurge" class="flex-1 py-5 bg-red-500 hover:bg-red-600 text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-red-900/20 transition-all">
              {{ $t('retention.purge_confirm.modal.btn_confirm') }}
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideCheckCircle2, LucideFolderMinus, LucideTrash2, 
  LucideAlertTriangle 
} from 'lucide-vue-next'

const showModal = ref(false)

const handleFinalPurge = () => {
  navigateTo('/retention/purge-success')
}
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }

.scale-enter-active, .scale-leave-active { transition: all 0.3s ease; }
.scale-enter-from, .scale-leave-to { opacity: 0; transform: scale(0.9); }
</style>
