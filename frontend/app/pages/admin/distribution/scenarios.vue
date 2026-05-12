<template>
  <div class="max-w-6xl mx-auto space-y-10 pb-32" v-motion-fade>
    <!-- Header -->
    <nav class="flex items-center gap-2 text-[10px] font-black text-slate-400 uppercase tracking-widest">
      <span>Distribution</span>
      <LucideChevronRight class="w-3 h-3" />
      <span class="text-[#1E3A5F] dark:text-white uppercase tracking-tighter">Delivery Method</span>
    </nav>

    <div class="space-y-1">
      <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
        Dual Scenario Digital Distribution Config
      </h1>
      <p class="text-sm font-bold text-slate-500">
        Konfigurasi metode pengiriman dokumen elektronik <span class="text-slate-400">(F-39)</span>.
      </p>
    </div>

    <!-- Scenarios Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
      <!-- Scenario A -->
      <div 
        @click="selectedScenario = 'A'"
        class="glass p-10 rounded-lg relative cursor-pointer transition-all duration-500 border-2"
        :class="selectedScenario === 'A' ? 'border-[#1E3A5F] shadow-2xl shadow-blue-900/10' : 'border-transparent hover:border-slate-200 dark:hover:border-slate-800'"
        v-motion-slide-visible-bottom
      >
        <div class="absolute -top-4 left-10 px-6 py-2 rounded-full text-[9px] font-black uppercase tracking-widest transition-colors shadow-sm"
          :class="selectedScenario === 'A' ? 'bg-[#1E3A5F] text-white' : 'bg-slate-100 text-slate-400 dark:bg-slate-800'"
        >
          Skenario A
        </div>

        <div class="flex items-start justify-between mb-8 pt-4">
          <div class="space-y-1">
            <h2 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">User Download Mandiri</h2>
            <p class="text-[11px] font-bold text-slate-500 leading-relaxed max-w-[280px]">
              Sistem membuat tautan publik yang aman untuk diunduh langsung oleh penerima.
            </p>
          </div>
          <div class="w-12 h-12 rounded-full bg-blue-50 dark:bg-blue-900/20 flex items-center justify-center text-blue-500 shadow-inner">
            <LucideLink class="w-5 h-5" />
          </div>
        </div>

        <!-- Inner Config Box -->
        <div class="p-8 bg-slate-50 dark:bg-slate-800/50 rounded-3xl space-y-6 border border-slate-100 dark:border-slate-800">
          <div class="space-y-3">
            <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">One-Time Link Generator</p>
            <div class="flex items-center gap-3">
              <div class="flex-grow bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-700 rounded-xl px-4 py-3 font-mono text-[10px] text-blue-500 truncate">
                https://dms.akiradata.id/dl/v4/sh-9218...
              </div>
              <button class="p-3 bg-slate-100 dark:bg-slate-800 rounded-xl hover:bg-slate-200 transition-colors">
                <LucideCopy class="w-4 h-4 text-slate-400" />
              </button>
            </div>
          </div>
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-2 text-red-500">
              <LucideClock class="w-3.5 h-3.5" />
              <span class="text-[10px] font-black uppercase">Expired 24 jam</span>
            </div>
            <button class="text-[9px] font-black text-red-500 uppercase tracking-widest hover:underline">Revoke Access</button>
          </div>
        </div>

        <div class="grid grid-cols-2 gap-4 mt-8">
          <div class="p-6 bg-slate-50 dark:bg-slate-800/30 rounded-2xl space-y-1">
            <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">Security Level</p>
            <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Multi-Factor</p>
          </div>
          <div class="p-6 bg-slate-50 dark:bg-slate-800/30 rounded-2xl space-y-1">
            <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">Max Retries</p>
            <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">3 Attempts</p>
          </div>
        </div>

        <div class="mt-8 pt-8 border-t border-slate-100 dark:border-slate-800 flex items-center justify-between">
          <div class="flex items-center gap-4">
            <div class="w-10 h-10 rounded-xl bg-red-50 dark:bg-red-900/20 flex items-center justify-center text-red-500">
              <LucideFileText class="w-5 h-5" />
            </div>
            <div>
              <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">INV-2023-004.pdf</p>
              <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">2.4 MB • Legal Template</p>
            </div>
          </div>
          <LucideCheckCircle2 class="w-5 h-5 text-blue-500" />
        </div>
      </div>

      <!-- Scenario B -->
      <div 
        @click="selectedScenario = 'B'"
        class="glass p-10 rounded-lg relative cursor-pointer transition-all duration-500 border-2"
        :class="selectedScenario === 'B' ? 'border-blue-500 shadow-2xl shadow-blue-900/10' : 'border-transparent opacity-60 hover:opacity-100'"
        v-motion-slide-visible-bottom
        :delay="100"
      >
        <div class="absolute -top-4 left-10 px-6 py-2 rounded-full text-[9px] font-black uppercase tracking-widest transition-colors shadow-sm"
          :class="selectedScenario === 'B' ? 'bg-blue-500 text-white' : 'bg-slate-100 text-slate-400 dark:bg-slate-800'"
        >
          Skenario B
        </div>

        <div class="flex items-start justify-between mb-8 pt-4">
          <div class="space-y-1">
            <h2 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Admin Kirim Manual via Email</h2>
            <p class="text-[11px] font-bold text-slate-500 leading-relaxed max-w-[280px]">
              Kirim dokumen sebagai lampiran melalui server SMTP internal.
            </p>
          </div>
          <div class="w-12 h-12 rounded-full bg-slate-50 dark:bg-slate-800 flex items-center justify-center text-slate-400 shadow-inner">
            <LucideMail class="w-5 h-5" />
          </div>
        </div>

        <div class="space-y-6">
          <div class="space-y-3">
            <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Recipient Email Address</label>
            <input 
              type="text" 
              placeholder="example@legal.com"
              class="w-full px-5 py-4 bg-slate-50 dark:bg-slate-800/50 border border-slate-100 dark:border-slate-700 rounded-2xl text-[11px] font-bold outline-none focus:ring-2 focus:ring-blue-500/20 transition-all"
            />
          </div>
          <div class="space-y-3">
            <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest px-1">Email Template</label>
            <div class="relative">
              <select class="w-full px-5 py-4 bg-slate-50 dark:bg-slate-800/50 border border-slate-100 dark:border-slate-700 rounded-2xl text-[11px] font-bold appearance-none outline-none">
                <option>Standard Document Distribution</option>
              </select>
              <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
            </div>
          </div>
          <div class="p-6 bg-slate-50 dark:bg-slate-800/30 rounded-2xl space-y-2">
            <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Attachment Summary</p>
            <div class="flex items-center gap-2 text-[10px] font-black text-[#1E3A5F] dark:text-slate-300 uppercase tracking-tight">
              <LucideLink class="w-3 h-3" />
              1 file attached (INV-2023-004.pdf)
            </div>
          </div>
        </div>

        <button 
          v-if="selectedScenario !== 'B'"
          class="w-full mt-10 py-4 border-2 border-slate-200 dark:border-slate-800 text-[10px] font-black text-slate-400 uppercase tracking-widest rounded-2xl hover:bg-slate-50 transition-all"
        >
          Select Skenario B
        </button>
      </div>
    </div>

    <!-- Sticky Bottom Bar -->
    <div class="fixed bottom-10 left-1/2 -translate-x-1/2 w-full max-w-5xl px-6 z-50">
      <div class="glass p-8 rounded-lg shadow-2xl shadow-blue-900/20 border border-white/20 dark:border-white/5 flex items-center justify-between" v-motion-slide-bottom>
        <div class="flex items-center gap-6">
          <div class="w-12 h-12 rounded-2xl bg-slate-50 dark:bg-slate-800 flex items-center justify-center text-[#1E3A5F] dark:text-blue-400 shadow-inner">
            <LucideShieldCheck class="w-6 h-6" />
          </div>
          <div>
            <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Validasi Sistem Lengkap</h4>
            <p class="text-[10px] font-bold text-slate-500 uppercase">Skenario {{ selectedScenario }} siap dijalankan dengan enkripsi AES-256.</p>
          </div>
        </div>
        <div class="flex items-center gap-8">
          <button class="text-[10px] font-black text-slate-400 hover:text-[#1E3A5F] transition-colors uppercase tracking-widest">Simpan Draft</button>
          <button class="px-10 py-4 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-xs font-black uppercase tracking-widest transition-all shadow-xl shadow-blue-900/20 active:scale-95">
            Lanjutkan Distribusi
          </button>
        </div>
      </div>
    </div>

    <!-- Footer Info -->
    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10 opacity-70 hover:opacity-100 transition-opacity">
      <div class="lg:col-span-4 space-y-6">
        <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest flex items-center gap-3">
          <div class="w-1.5 h-1.5 rounded-full bg-blue-500"></div>
          Metode Log
        </h3>
        <div class="space-y-8 pl-1.5 border-l-2 border-slate-100 dark:border-slate-800">
          <div class="relative pl-6">
            <div class="absolute -left-[7px] top-0 w-3 h-3 bg-white dark:bg-slate-900 border-2 border-blue-500 rounded-full shadow-sm"></div>
            <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Config Initiated</p>
            <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest mt-1">Today at 14:20 • Admin Legal</p>
          </div>
          <div class="relative pl-6">
            <div class="absolute -left-[7px] top-0 w-3 h-3 bg-white dark:bg-slate-900 border-2 border-slate-200 dark:border-slate-700 rounded-full shadow-sm"></div>
            <p class="text-[11px] font-black text-slate-400 uppercase tracking-tight">Scenario {{ selectedScenario }} Drafted</p>
            <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest mt-1">Today at 14:22 • Auto-save</p>
          </div>
        </div>
      </div>

      <div class="lg:col-span-8">
        <div class="glass p-10 rounded-lg bg-gradient-to-br from-white/50 to-slate-50/50 dark:from-slate-900/50 dark:to-slate-800/50 relative overflow-hidden group">
          <div class="max-w-xl space-y-4 relative z-10">
            <h2 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Keamanan Berlapis</h2>
            <p class="text-sm font-bold text-slate-500 leading-relaxed">
              Seluruh distribusi digital melalui skenario 'User Download Mandiri' akan direkam dalam audit log sistem dengan identifikasi alamat IP dan metadata perangkat untuk kepatuhan hukum.
            </p>
          </div>
          <div class="absolute right-0 top-1/2 -translate-y-1/2 w-64 h-64 opacity-5 group-hover:scale-110 transition-transform pointer-events-none">
            <LucideShieldCheck class="w-full h-full text-[#1E3A5F]" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideChevronRight, LucideLink, LucideMail, LucideClock, 
  LucideCopy, LucideFileText, LucideCheckCircle2, LucideChevronDown, 
  LucideShieldCheck
} from 'lucide-vue-next'

const selectedScenario = ref('A')

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
@reference "tailwindcss";
.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/5 shadow-sm;
}
</style>

