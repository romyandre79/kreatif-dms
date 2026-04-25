<template>
  <div class="max-w-7xl mx-auto space-y-8 pb-20">
    <!-- Top Navigation -->
    <div class="flex items-center justify-between" v-motion-fade>
      <div class="space-y-4">
        <button 
          @click="navigateTo('/loans/cart')" 
          class="flex items-center gap-2 text-xs font-black text-slate-400 hover:text-primary-600 transition-colors uppercase tracking-widest group"
        >
          <LucideArrowLeft class="w-4 h-4 group-hover:-translate-x-1 transition-transform" />
          {{ $t('documents.detail.btn_back') }}
        </button>
        <div class="space-y-1">
          <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">{{ $t('loans.checkout.title') }}</h1>
          <p class="text-sm font-bold text-slate-400 uppercase tracking-widest">{{ $t('loans.checkout.subtitle') }}</p>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8 items-start">
      <!-- Left Column: Details -->
      <div class="lg:col-span-8 space-y-8">
        <!-- Document Review Table -->
        <div class="glass rounded-3xl overflow-hidden border border-slate-100 bg-white">
          <div class="px-8 py-5 border-b border-slate-100 bg-slate-50/30">
            <h2 class="text-xs font-black text-[#1E3A5F] uppercase tracking-[0.2em]">{{ $t('loans.checkout.review.title', { count: selectedDocs.length }) }}</h2>
          </div>
          <table class="w-full text-left">
            <thead>
              <tr class="text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100 bg-slate-50/10">
                <th class="px-8 py-4">{{ $t('loans.checkout.review.table.doc_no') }}</th>
                <th class="px-8 py-4">{{ $t('loans.checkout.review.table.doc_title') }}</th>
                <th class="px-8 py-4">{{ $t('loans.checkout.review.table.category') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50">
              <tr v-for="doc in selectedDocs" :key="doc.id" class="text-xs">
                <td class="px-8 py-4 font-bold text-slate-400">{{ doc.id }}</td>
                <td class="px-8 py-4 font-black text-[#1E3A5F] uppercase">{{ doc.title }}</td>
                <td class="px-8 py-4">
                  <span class="px-3 py-1 bg-slate-100 text-slate-500 rounded-lg font-bold">{{ doc.category }}</span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Loan Details Form -->
        <div class="glass rounded-3xl p-10 space-y-10 bg-white border border-slate-100">
          <h2 class="text-xs font-black text-[#1E3A5F] uppercase tracking-[0.2em]">{{ $t('loans.checkout.form.title') }}</h2>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-10">
            <!-- Pickup Method -->
            <div class="space-y-4">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('loans.checkout.form.method_label') }}</label>
              <div class="flex items-center gap-4">
                <label class="flex items-center gap-3 p-4 bg-slate-50 rounded-2xl border-2 border-primary-500/20 cursor-pointer flex-1">
                  <input type="radio" checked name="method" class="w-5 h-5 text-primary-600 focus:ring-primary-500/10" />
                  <span class="text-sm font-bold text-slate-700">{{ $t('loans.checkout.form.method_courier') }}</span>
                </label>
              </div>
            </div>

            <!-- Warehouse Location -->
            <div class="space-y-4">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('loans.checkout.form.warehouse_label') }}</label>
              <div class="p-4 bg-primary-50 border border-primary-100 rounded-2xl flex items-center gap-3">
                <LucideMapPin class="w-5 h-5 text-primary-600" />
                <span class="text-sm font-black text-primary-900 tracking-tight">{{ $t('loans.checkout.form.warehouse_default') }}</span>
              </div>
            </div>

            <!-- Pickup Date -->
            <div class="space-y-4">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('loans.checkout.form.pickup_date_label') }}</label>
              <div class="relative">
                <input 
                  type="date" 
                  class="w-full pl-4 pr-10 py-4 bg-slate-50 border border-slate-200 rounded-2xl text-sm font-bold text-slate-700 focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 outline-none transition-all"
                />
              </div>
            </div>

            <!-- Time Slot -->
            <div class="space-y-4">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('loans.checkout.form.time_slot_label') }}</label>
              <div class="relative">
                <select class="w-full pl-4 pr-10 py-4 bg-slate-50 border border-slate-200 rounded-2xl text-sm font-bold text-slate-700 appearance-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 outline-none transition-all">
                  <option>{{ $t('loans.checkout.form.time_slot_placeholder') }}</option>
                  <option>09:00 - 11:00</option>
                  <option>13:00 - 15:00</option>
                </select>
                <LucideChevronDown class="absolute right-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400 pointer-events-none" />
              </div>
            </div>
          </div>

          <!-- Notes -->
          <div class="space-y-4">
            <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('loans.checkout.form.notes_label') }}</label>
            <textarea 
              :placeholder="$t('loans.checkout.form.notes_placeholder')" 
              rows="4" 
              class="w-full p-6 bg-slate-50 border border-slate-200 rounded-3xl text-sm font-bold text-slate-700 outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all resize-none"
            ></textarea>
          </div>
        </div>
      </div>

      <!-- Right Column: Summary -->
      <div class="lg:col-span-4 space-y-6" v-motion-slide-visible-bottom>
        <div class="glass rounded-[2.5rem] overflow-hidden bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 flex flex-col">
          <div class="bg-[#1E3A5F] px-8 py-6">
            <h2 class="text-lg font-black text-white uppercase tracking-tight">{{ $t('loans.checkout.summary.title') }}</h2>
          </div>
          
          <div class="p-8 flex-1 space-y-8">
            <div class="space-y-6">
              <div class="flex justify-between items-center">
                <span class="text-sm font-bold text-slate-400">{{ $t('loans.checkout.summary.total_docs') }}</span>
                <span class="text-sm font-black text-[#1E3A5F]">{{ $t('loans.checkout.summary.total_files', { count: 3 }) }}</span>
              </div>
              <div class="flex justify-between items-center">
                <span class="text-sm font-bold text-slate-400">{{ $t('loans.checkout.summary.duration') }}</span>
                <span class="text-sm font-black text-[#1E3A5F]">{{ $t('loans.checkout.summary.duration_value', { count: 7 }) }}</span>
              </div>
              <div class="flex justify-between items-start">
                <span class="text-sm font-bold text-slate-400">{{ $t('loans.checkout.summary.est_return') }}</span>
                <span class="text-sm font-black text-primary-600 text-right leading-tight" v-html="$t('loans.checkout.summary.est_return_pending')"></span>
              </div>
            </div>

            <div class="w-full h-px bg-slate-100"></div>

            <div class="flex items-center justify-between">
              <span class="text-[10px] font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('loans.checkout.summary.security_level') }}</span>
              <span class="px-3 py-1 bg-amber-50 text-amber-600 rounded-full text-[10px] font-black uppercase tracking-widest border border-amber-100">CONFIDENTIAL</span>
            </div>

            <div class="pt-4 space-y-4">
              <button class="w-full flex items-center justify-center gap-3 py-5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 transition-all group">
                <LucideClipboardCheck class="w-5 h-5 group-hover:scale-110 transition-transform" />
                {{ $t('loans.checkout.summary.btn_confirm') }}
              </button>
              <button @click="navigateTo('/loans/cart')" class="w-full py-4 text-slate-400 hover:text-slate-600 text-xs font-black uppercase tracking-widest flex items-center justify-center gap-2">
                <LucideArrowLeft class="w-4 h-4" />
                {{ $t('loans.checkout.summary.btn_back') }}
              </button>
            </div>
          </div>
        </div>

        <!-- Info Box -->
        <div class="p-6 bg-blue-50 border-l-4 border-blue-500 rounded-r-2xl flex gap-4">
          <LucideInfo class="w-6 h-6 text-blue-500 shrink-0" />
          <p class="text-[11px] font-bold text-blue-800 leading-relaxed">
            {{ $t('loans.checkout.info_box') }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideArrowLeft, 
  LucideMapPin, 
  LucideChevronDown, 
  LucideClipboardCheck, 
  LucideInfo 
} from 'lucide-vue-next'

const selectedDocs = [
  { id: 'DOC-2023-XYZ-001', title: 'Laporan Keuangan Q3 2023', category: 'Financial' },
  { id: 'DOC-2023-XYZ-004', title: 'Kontrak Vendor IT - Phase 1', category: 'Legal' },
  { id: 'DOC-2023-ABC-012', title: 'SOP Pengarsipan Digital', category: 'Internal' }
]
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
</style>
