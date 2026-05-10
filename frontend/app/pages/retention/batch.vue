<template>
  <div class="max-w-[1600px] mx-auto space-y-8 pb-20">
    <!-- Header -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6" v-motion-fade>
      <div class="space-y-2">
        <h1 class="text-3xl font-black text-[#1E3A5F] tracking-tight uppercase">{{ $t('retention.batch.title') }}</h1>
        <p class="text-slate-500 font-medium">{{ $t('retention.batch.subtitle') }}</p>
      </div>
      <div class="flex items-center gap-4">
        <button @click="navigateTo('/retention/approaching')" class="px-6 py-4 bg-white border border-slate-200 text-slate-400 rounded-2xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all flex items-center gap-2">
          <LucideX class="w-4 h-4" /> {{ $t('retention.batch.btn_cancel') }}
        </button>
        <button @click="navigateTo('/retention/export')" class="px-8 py-4 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-xs font-black uppercase tracking-[0.2em] shadow-xl shadow-blue-900/20 transition-all flex items-center gap-3 group">
          <LucideFileDown class="w-4 h-4 group-hover:translate-y-1 transition-transform" /> {{ $t('retention.batch.btn_export') }}
        </button>
      </div>
    </div>

    <!-- Info Cards -->
    <div class="grid grid-cols-3 gap-8" v-motion-slide-visible-bottom>
      <div class="glass p-10 rounded-[3rem] bg-white border border-slate-100 shadow-xl shadow-slate-200/50 flex items-center gap-8 relative overflow-hidden group">
        <div class="p-5 bg-blue-50 text-primary-500 rounded-3xl relative z-10 group-hover:bg-primary-500 group-hover:text-white transition-all duration-500">
          <LucideHash class="w-8 h-8" />
        </div>
        <div class="space-y-1 relative z-10">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('retention.batch.cards.identifier') }}</p>
          <p class="text-2xl font-black text-[#1E3A5F] tracking-tight">BATCH-2024-001</p>
        </div>
        <LucideHash class="absolute -right-4 -bottom-4 w-32 h-32 text-slate-50 opacity-50 pointer-events-none group-hover:text-primary-50 group-hover:opacity-100 transition-all duration-700" />
      </div>

      <div class="glass p-10 rounded-[3rem] bg-white border border-slate-100 shadow-xl shadow-slate-200/50 flex items-center gap-8 relative overflow-hidden group">
        <div class="p-5 bg-emerald-50 text-emerald-500 rounded-3xl relative z-10 group-hover:bg-emerald-500 group-hover:text-white transition-all duration-500">
          <LucideFileStack class="w-8 h-8" />
        </div>
        <div class="space-y-1 relative z-10">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('retention.batch.cards.selected') }}</p>
          <p class="text-2xl font-black text-[#1E3A5F] tracking-tight">15 Documents</p>
        </div>
        <LucideFileStack class="absolute -right-4 -bottom-4 w-32 h-32 text-slate-50 opacity-50 pointer-events-none group-hover:text-emerald-50 group-hover:opacity-100 transition-all duration-700" />
      </div>

      <div class="glass p-10 rounded-[3rem] bg-white border border-slate-100 shadow-xl shadow-slate-200/50 flex items-center gap-8 relative overflow-hidden group">
        <div class="p-5 bg-orange-50 text-orange-500 rounded-3xl relative z-10 group-hover:bg-orange-500 group-hover:text-white transition-all duration-500">
          <LucideCalendar class="w-8 h-8" />
        </div>
        <div class="space-y-1 relative z-10">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('retention.batch.cards.creation') }}</p>
          <p class="text-2xl font-black text-[#1E3A5F] tracking-tight">Oct 24, 2023</p>
        </div>
        <LucideCalendar class="absolute -right-4 -bottom-4 w-32 h-32 text-slate-50 opacity-50 pointer-events-none group-hover:text-orange-50 group-hover:opacity-100 transition-all duration-700" />
      </div>
    </div>

    <!-- Document List -->
    <div class="glass rounded-[3rem] bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 overflow-hidden" v-motion-slide-visible-bottom>
      <div class="p-10 border-b border-slate-50 flex items-center justify-between">
        <h2 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('retention.batch.table.title') }}</h2>
        <button class="flex items-center gap-2 text-[11px] font-black text-primary-500 uppercase tracking-widest hover:text-primary-600 transition-colors">
          <LucidePlusCircle class="w-4 h-4" /> {{ $t('retention.batch.table.add_more') }}
        </button>
      </div>

      <div class="overflow-x-auto">
        <table class="w-full text-left">
          <thead>
            <tr class="text-[10px] font-black text-slate-300 uppercase tracking-widest border-b border-slate-50 bg-slate-50/30">
              <th class="px-10 py-6">{{ $t('retention.batch.table.id') }}</th>
              <th class="px-6 py-6">{{ $t('retention.batch.table.doc_title') }}</th>
              <th class="px-6 py-6">{{ $t('retention.batch.table.category') }}</th>
              <th class="px-6 py-6">{{ $t('retention.batch.table.date') }}</th>
              <th class="px-10 py-6 text-right">{{ $t('retention.batch.table.actions') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50">
            <tr v-for="doc in items" :key="doc.id" class="group hover:bg-slate-50/50 transition-colors">
              <td class="px-10 py-8 text-xs font-black text-slate-400 uppercase tracking-tight">{{ doc.id }}</td>
              <td class="px-6 py-8">
                <p class="text-sm font-black text-[#1E3A5F] tracking-tight group-hover:text-primary-600 transition-colors">{{ doc.title }}</p>
              </td>
              <td class="px-6 py-8">
                <span :class="`px-3 py-1 rounded-lg text-[9px] font-black uppercase tracking-widest ${doc.catColor}`">
                  {{ doc.category }}
                </span>
              </td>
              <td class="px-6 py-8 text-xs font-bold text-slate-500 uppercase tracking-tight">{{ doc.date }}</td>
              <td class="px-10 py-8 text-right">
                <button class="p-3 text-slate-200 hover:text-red-500 hover:bg-red-50 rounded-2xl transition-all">
                  <LucideTrash2 class="w-5 h-5" />
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div class="p-8 bg-slate-50/30 text-center border-t border-slate-50">
        <button class="text-[10px] font-black text-slate-300 uppercase tracking-[0.2em] hover:text-slate-500 transition-colors">
          {{ $t('retention.batch.table.show_more', { count: 12 }) }}
        </button>
      </div>
    </div>

    <!-- Evaluation Notes -->
    <div class="glass p-10 rounded-[3rem] bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-8" v-motion-slide-visible-bottom>
      <div class="flex items-center gap-4">
        <div class="p-3 bg-slate-50 text-[#1E3A5F] rounded-2xl">
          <LucideMessageSquareText class="w-6 h-6" />
        </div>
        <h2 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('retention.batch.notes.title') }}</h2>
      </div>

      <div class="space-y-4">
        <textarea v-model="notes" 
                  maxlength="1000"
                  :placeholder="$t('retention.batch.notes.placeholder')" 
                  rows="6" 
                  class="w-full p-10 bg-slate-50/50 border border-slate-100 rounded-lg text-sm font-medium text-slate-600 outline-none focus:ring-4 focus:ring-primary-500/5 focus:border-primary-500 transition-all resize-none"></textarea>
        
        <div class="flex items-center justify-between px-2">
          <p class="text-[10px] font-bold text-slate-300 uppercase tracking-widest">
            {{ $t('retention.batch.notes.hint') }}
          </p>
          <p :class="`text-[10px] font-black uppercase tracking-widest ${notes.length > 900 ? 'text-red-500' : 'text-slate-300'}`">
            {{ $t('retention.batch.notes.char_limit', { count: notes.length }) }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideX, LucideFileDown, LucideHash, LucideFileStack, 
  LucideCalendar, LucidePlusCircle, LucideTrash2, 
  LucideMessageSquareText 
} from 'lucide-vue-next'

const notes = ref('')
const items = ref([
  {
    id: 'REF-2023-X92',
    title: 'Financial Audit Report Q4',
    category: 'FINANCIAL',
    catColor: 'bg-blue-50 text-blue-500 border border-blue-100',
    date: 'Dec 15, 2023'
  },
  {
    id: 'REF-2023-A41',
    title: 'Client Acquisition Contracts',
    category: 'LEGAL',
    catColor: 'bg-purple-50 text-purple-500 border border-purple-100',
    date: 'Jan 02, 2024'
  },
  {
    id: 'REF-2023-K12',
    title: 'Staff Payroll Summary FY23',
    category: 'HR',
    catColor: 'bg-yellow-50 text-orange-500 border border-orange-100',
    date: 'Dec 20, 2023'
  }
])
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
</style>
