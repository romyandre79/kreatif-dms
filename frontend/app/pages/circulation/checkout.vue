<template>
  <div class="max-w-[1600px] mx-auto grid grid-cols-12 gap-8 items-start pb-20">
    <!-- Left Sidebar: Pickup Queue -->
    <aside class="col-span-3 space-y-6" v-motion-slide-left>
      <div class="flex items-center justify-between px-4">
        <h2 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('circulation.checkout.pickup_queue') }}</h2>
        <span class="px-3 py-1 bg-blue-50 text-blue-600 rounded-lg text-[10px] font-black uppercase tracking-widest border border-blue-100">
          {{ $t('circulation.checkout.pending_count', { count: 4 }) }}
        </span>
      </div>

      <div class="space-y-4">
        <div v-for="req in queue" :key="req.no" 
             @click="selectedRequest = req"
             :class="`group p-6 bg-white rounded-lg border transition-all cursor-pointer shadow-xl shadow-slate-200/20 ${selectedRequest?.no === req.no ? 'border-primary-500 ring-4 ring-primary-500/10' : 'border-slate-100 hover:border-slate-200'}`">
          <div class="flex items-center justify-between mb-4">
            <p class="text-[11px] font-black text-primary-600 uppercase tracking-tight">{{ req.no }}</p>
            <span class="text-[9px] font-black text-emerald-500 uppercase tracking-widest bg-emerald-50 px-2 py-0.5 rounded-md border border-emerald-100">{{ $t('circulation.checkout.status_ready') }}</span>
          </div>
          <div class="space-y-1 mb-4">
            <h4 class="text-sm font-black text-[#1E3A5F] group-hover:text-primary-600 transition-colors">{{ req.name }}</h4>
            <p class="text-[10px] font-bold text-slate-400 uppercase">{{ req.dept }}</p>
          </div>
          <div class="flex items-center gap-2 text-slate-400">
            <LucideClock class="w-3.5 h-3.5" />
            <span class="text-[10px] font-bold uppercase tracking-tight">{{ $t('circulation.checkout.pickup_time', { time: req.time }) }}</span>
          </div>
        </div>
      </div>
    </aside>

    <!-- Main Content: Verification & Handover -->
    <div class="col-span-6 space-y-8" v-motion-fade>
      <!-- Pickup Verification Card -->
      <div class="glass p-10 rounded-[3rem] bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-10">
        <div class="flex items-center gap-4">
          <div class="p-3 bg-blue-50 text-blue-600 rounded-2xl">
            <LucideShieldCheck class="w-6 h-6" />
          </div>
          <h2 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('circulation.checkout.verification_title') }}</h2>
        </div>

        <div class="grid grid-cols-2 gap-y-10 gap-x-12">
          <div class="space-y-1">
            <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ $t('circulation.checkout.req_no') }}</p>
            <p class="text-sm font-black text-[#1E3A5F] tracking-tight uppercase">{{ selectedRequest?.no }}</p>
          </div>
          <div class="space-y-1">
            <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ $t('circulation.checkout.pickup_code') }}</p>
            <p class="text-sm font-black text-primary-600 tracking-widest uppercase">AK-9921-X</p>
          </div>
          <div class="space-y-1">
            <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ $t('circulation.checkout.requester_name') }}</p>
            <p class="text-sm font-black text-[#1E3A5F]">{{ selectedRequest?.name }}</p>
          </div>
          <div class="space-y-1">
            <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ $t('circulation.checkout.dept') }}</p>
            <p class="text-sm font-black text-[#1E3A5F] uppercase">{{ selectedRequest?.dept }}</p>
          </div>
          <div class="col-span-2 space-y-1">
            <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ $t('circulation.checkout.pickup_window') }}</p>
            <p class="text-sm font-black text-[#1E3A5F] uppercase">August 24, 2023 | 14:00 - 15:00 WIB</p>
          </div>
        </div>

        <div class="p-8 bg-blue-50/50 border border-blue-100 rounded-lg space-y-6">
          <p class="text-[11px] font-black text-blue-700 uppercase tracking-widest">{{ $t('circulation.checkout.checklist_title') }}</p>
          <div class="space-y-4">
            <label class="flex items-center gap-4 cursor-pointer group">
              <input type="checkbox" class="w-5 h-5 rounded border-2 border-blue-200 text-blue-600 focus:ring-blue-500/10" />
              <span class="text-xs font-bold text-slate-600 group-hover:text-blue-900 transition-colors">{{ $t('circulation.checkout.checklist_id') }}</span>
            </label>
            <label class="flex items-center gap-4 cursor-pointer group">
              <input type="checkbox" class="w-5 h-5 rounded border-2 border-blue-200 text-blue-600 focus:ring-blue-500/10" />
              <span class="text-xs font-bold text-slate-600 group-hover:text-blue-900 transition-colors">{{ $t('circulation.checkout.checklist_code') }}</span>
            </label>
          </div>
        </div>
      </div>

      <!-- Document Handover Table -->
      <div class="glass p-10 rounded-[3rem] bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-10">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-4">
            <div class="p-3 bg-slate-50 text-slate-400 rounded-2xl">
              <LucideFileStack class="w-6 h-6" />
            </div>
            <h2 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('circulation.checkout.handover_title') }}</h2>
          </div>
          <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.checkout.selected_count', { count: 3 }) }}</span>
        </div>

        <div class="overflow-hidden border-t border-slate-50">
          <table class="w-full text-left">
            <thead>
              <tr class="text-[9px] font-black text-slate-300 uppercase tracking-widest">
                <th class="py-6">{{ $t('circulation.checkout.table.doc_id') }}</th>
                <th class="py-6">{{ $t('circulation.checkout.table.doc_name') }}</th>
                <th class="py-6 text-center">{{ $t('circulation.checkout.table.handover') }}</th>
                <th class="py-6">{{ $t('circulation.checkout.table.condition') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50">
              <tr v-for="doc in selectedRequest?.docs" :key="doc.id" class="group">
                <td class="py-6 text-xs font-black text-slate-400 uppercase tracking-tight">{{ doc.id }}</td>
                <td class="py-6">
                  <p class="text-xs font-black text-[#1E3A5F]">{{ doc.name }}</p>
                  <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest mt-1">{{ doc.type }} • {{ doc.pages }} Pages</p>
                </td>
                <td class="py-6 text-center">
                  <input type="checkbox" checked class="w-5 h-5 rounded border-2 border-emerald-200 text-emerald-500" />
                </td>
                <td class="py-6">
                  <input type="text" :placeholder="$t('circulation.checkout.table.note_placeholder')" 
                         class="w-full bg-slate-50 border border-slate-100 rounded-xl px-4 py-2 text-[11px] font-medium outline-none focus:border-primary-500 transition-all" />
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <!-- Right Sidebar: Scan & Signature -->
    <div class="col-span-3 space-y-8" v-motion-slide-right>
      <!-- Scan & Confirm -->
      <div class="glass p-8 rounded-lg bg-white border border-slate-100 shadow-xl shadow-slate-200/50 space-y-6">
        <div class="flex items-center gap-3 text-slate-400 uppercase tracking-widest text-[10px] font-black">
          <LucideScanBarcode class="w-4 h-4 text-primary-500" />
          {{ $t('circulation.checkout.scan_confirm') }}
        </div>
        <div class="aspect-video border-2 border-dashed border-slate-100 rounded-lg flex flex-col items-center justify-center text-center p-6 space-y-3 bg-slate-50/50 group hover:border-primary-200 transition-all cursor-pointer">
          <div class="w-12 h-12 bg-white rounded-2xl flex items-center justify-center shadow-sm text-slate-300 group-hover:text-primary-500 transition-colors">
            <LucideQrCode class="w-6 h-6" />
          </div>
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.checkout.scan_hint') }}</p>
        </div>
        <p class="text-[9px] font-bold text-slate-300 text-center uppercase leading-loose">
          {{ $t('circulation.checkout.scan_desc') }}
        </p>
      </div>

      <!-- Signature & Completion -->
      <div class="glass p-8 rounded-lg bg-white border border-slate-100 shadow-xl shadow-slate-200/50 space-y-8">
        <div class="flex items-center gap-3 text-slate-400 uppercase tracking-widest text-[10px] font-black">
          <LucidePenTool class="w-4 h-4 text-primary-500" />
          {{ $t('circulation.checkout.signature_title') }}
        </div>
        
        <!-- Signature Pad Placeholder -->
        <div class="relative aspect-square bg-slate-50 border border-slate-100 rounded-lg overflow-hidden flex flex-col items-center justify-center group">
          <LucideSignature class="w-16 h-16 text-slate-100 group-hover:text-slate-200 transition-colors" />
          <p class="text-[10px] font-black text-slate-200 uppercase tracking-[0.2em] mt-4">{{ $t('circulation.checkout.signature_area') }}</p>
          <button class="absolute bottom-6 right-6 text-[9px] font-black text-primary-500 uppercase tracking-widest bg-white px-3 py-1.5 rounded-lg border border-primary-100 shadow-sm">{{ $t('circulation.checkout.btn_clear') }}</button>
        </div>

        <label class="flex items-start gap-4 cursor-pointer group">
          <input type="checkbox" class="w-5 h-5 rounded border-2 border-slate-200 text-[#1E3A5F] focus:ring-blue-900/10 mt-1" />
          <span class="text-[10px] font-bold text-slate-500 leading-relaxed group-hover:text-slate-900 transition-colors">
            {{ $t('circulation.checkout.confirm_checkbox') }}
          </span>
        </label>

        <div class="space-y-3">
          <button class="w-full py-5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-[11px] font-black uppercase tracking-[0.2em] shadow-xl shadow-blue-900/20 transition-all flex items-center justify-center gap-3">
            {{ $t('circulation.checkout.btn_complete') }} <LucideCheckCircle2 class="w-4 h-4" />
          </button>
          <button class="w-full py-4 bg-white border border-slate-200 text-slate-400 hover:bg-slate-50 rounded-2xl text-[10px] font-black uppercase tracking-[0.2em] transition-all">
            {{ $t('circulation.checkout.btn_cancel') }}
          </button>
        </div>
      </div>

      <!-- Handover Stats -->
      <div class="px-8 space-y-4">
        <div class="flex items-center justify-between text-[10px] font-black uppercase tracking-widest">
          <span class="text-slate-400">{{ $t('circulation.checkout.stats.today_total') }}</span>
          <span class="text-[#1E3A5F]">24</span>
        </div>
        <div class="space-y-2">
          <div class="flex items-center justify-between text-[10px] font-black uppercase tracking-widest">
            <span class="text-slate-400">{{ $t('circulation.checkout.stats.success_rate') }}</span>
            <span class="text-emerald-500">98.5%</span>
          </div>
          <div class="h-1.5 bg-slate-100 rounded-full overflow-hidden">
            <div class="h-full bg-emerald-500 w-[98.5%] rounded-full shadow-lg shadow-emerald-500/20"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideClock, LucideShieldCheck, LucideFileStack, 
  LucideScanBarcode, LucideQrCode, LucidePenTool, 
  LucideSignature, LucideCheckCircle2 
} from 'lucide-vue-next'

const queue = ref([
  {
    no: '#REQ-2023-0892',
    name: 'Budi Santoso',
    dept: 'Finance & Tax Dept',
    time: '14:00 - 15:00',
    docs: [
      { id: 'DOC-AF102', name: 'Tax Invoice - Q2 Vendor Group', type: 'Original Copy', pages: 12, notePlaceholder: 'e.g. Good condition' },
      { id: 'DOC-AF105', name: 'Lease Agreement - Building C', type: 'Original Copy', pages: 45, notePlaceholder: 'e.g. Good condition' },
      { id: 'DOC-AF109', name: 'PO #9812 - IT Infrastructure', type: 'Duplicate Copy', pages: 2, notePlaceholder: 'Slightly creased corner' }
    ]
  },
  {
    no: '#REQ-2023-0895',
    name: 'Siti Aminah',
    dept: 'Legal Affairs',
    time: '15:30 - 16:30',
    docs: []
  },
  {
    no: '#REQ-2023-0901',
    name: 'Robert Downey',
    dept: 'Human Resources',
    time: 'Tomorrow',
    docs: []
  }
])

const selectedRequest = ref(queue.value[0])
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
</style>
