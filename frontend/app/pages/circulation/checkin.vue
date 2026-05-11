<template>
  <div class="max-w-[1600px] mx-auto space-y-8 pb-20">
    <!-- Return Intake Information Header -->
    <div class="glass p-10 rounded-lg bg-white border border-slate-100 shadow-2xl shadow-slate-200/50" v-motion-slide-top>
      <div class="flex items-center justify-between mb-10">
        <div class="flex items-center gap-4">
          <div class="p-3 bg-blue-50 text-blue-600 rounded-2xl">
            <LucideInbox class="w-6 h-6" />
          </div>
          <h2 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('circulation.checkin.title') }}</h2>
        </div>
        <span class="px-4 py-2 bg-orange-50 text-orange-600 border border-orange-100 rounded-xl text-[10px] font-black uppercase tracking-widest flex items-center gap-2">
          <LucideAlertTriangle class="w-3.5 h-3.5" /> ! {{ $t('circulation.checkin.late_return') }}
        </span>
      </div>

      <div class="grid grid-cols-5 gap-12">
        <div class="space-y-1">
          <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ $t('circulation.checkin.req_no') }}</p>
          <p class="text-sm font-black text-[#1E3A5F] tracking-tight uppercase">REQ-2023-0892</p>
        </div>
        <div class="space-y-1">
          <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ $t('circulation.checkin.requester_name') }}</p>
          <p class="text-sm font-black text-[#1E3A5F]">Andi Pratama</p>
          <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">Legal Department</p>
        </div>
        <div class="space-y-1">
          <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ $t('circulation.checkin.due_date') }}</p>
          <p class="text-sm font-black text-[#1E3A5F] uppercase">20 Oct 2023</p>
        </div>
        <div class="space-y-1">
          <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ $t('circulation.checkin.return_date') }}</p>
          <p class="text-sm font-black text-[#1E3A5F] uppercase">22 Oct 2023</p>
        </div>
        <div class="space-y-1">
          <p class="text-[10px] font-black text-slate-300 uppercase tracking-widest">{{ $t('circulation.checkin.late_status') }}</p>
          <p class="text-sm font-black text-red-500 uppercase tracking-tight">{{ $t('circulation.checkin.overdue_days', { count: 2 }) }}</p>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-12 gap-8 items-start">
      <!-- Left Section: Document Checklist -->
      <div class="col-span-8 glass p-10 rounded-lg bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-8" v-motion-slide-left>
        <div class="flex items-center justify-between">
          <h2 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('circulation.checkin.checklist_title') }}</h2>
          <span class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('circulation.checkin.checklist_hint') }}</span>
        </div>

        <div class="overflow-hidden border-t border-slate-50">
          <table class="w-full text-left">
            <thead>
              <tr class="text-[9px] font-black text-slate-300 uppercase tracking-widest">
                <th class="py-6">{{ $t('circulation.checkin.table.doc_title') }}</th>
                <th class="py-6">{{ $t('circulation.checkin.table.system_id') }}</th>
                <th class="py-6">{{ $t('circulation.checkin.table.expected') }}</th>
                <th class="py-6">{{ $t('circulation.checkin.table.actual_condition') }}</th>
                <th class="py-6">{{ $t('circulation.checkin.table.notes') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50">
              <tr v-for="doc in documents" :key="doc.id" class="group">
                <td class="py-6 pr-6">
                  <p class="text-xs font-black text-[#1E3A5F] leading-relaxed">{{ doc.title }}</p>
                </td>
                <td class="py-6">
                  <p class="text-[10px] font-black text-slate-400 uppercase tracking-tight">{{ doc.id }}</p>
                </td>
                <td class="py-6">
                  <span class="px-3 py-1 bg-emerald-50 text-emerald-600 border border-emerald-100 rounded-lg text-[9px] font-black uppercase tracking-widest">
                    {{ $t('circulation.checkin.table.status_complete') }}
                  </span>
                </td>
                <td class="py-6">
                  <div class="relative w-32">
                    <select v-model="doc.condition" :class="`w-full appearance-none bg-white border rounded-xl px-4 py-2.5 text-xs font-bold outline-none transition-all ${doc.condition === 'Rusak' ? 'border-red-500 text-red-500 ring-4 ring-red-500/5' : 'border-slate-200 text-[#1E3A5F] focus:border-primary-500'}`">
                      <option>{{ $t('circulation.checkin.conditions.complete') }}</option>
                      <option>{{ $t('circulation.checkin.conditions.damaged') }}</option>
                      <option>{{ $t('circulation.checkin.conditions.missing') }}</option>
                    </select>
                    <LucideChevronDown class="absolute right-3 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300 pointer-events-none" />
                  </div>
                </td>
                <td class="py-6">
                  <div class="space-y-1">
                    <input type="text" v-model="doc.note" :placeholder="doc.condition === 'Rusak' ? 'Ada noda air...' : $t('circulation.checkin.table.note_placeholder')"
                           :class="`w-full bg-slate-50 border rounded-xl px-4 py-2 text-[11px] font-medium outline-none transition-all ${doc.condition === 'Rusak' ? 'border-red-200 focus:border-red-500' : 'border-slate-100 focus:border-primary-500'}`" />
                    <p v-if="doc.condition === 'Rusak'" class="text-[8px] font-black text-red-500 uppercase tracking-widest">{{ $t('circulation.checkin.table.damage_reason_req') }}</p>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Right Sidebar: Evidence & Actions -->
      <div class="col-span-4 space-y-8" v-motion-slide-right>
        <!-- Condition Evidence -->
        <div class="glass p-8 rounded-lg bg-white border border-slate-100 shadow-xl shadow-slate-200/50 space-y-8">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3 text-slate-400 uppercase tracking-widest text-[10px] font-black">
              <LucideCamera class="w-4 h-4 text-primary-500" />
              {{ $t('circulation.checkin.evidence.title') }}
            </div>
            <span class="px-2 py-0.5 bg-red-50 text-red-500 border border-red-100 rounded text-[8px] font-black uppercase tracking-widest">{{ $t('circulation.checkin.evidence.required') }}</span>
          </div>

          <div class="space-y-6">
            <div class="aspect-video border-2 border-dashed border-slate-100 rounded-lg flex flex-col items-center justify-center text-center p-6 space-y-3 bg-slate-50/50 group hover:border-primary-200 transition-all cursor-pointer">
              <div class="w-12 h-12 bg-white rounded-2xl flex items-center justify-center shadow-sm text-slate-300 group-hover:text-primary-500 transition-colors">
                <LucideUploadCloud class="w-6 h-6" />
              </div>
              <div class="space-y-1">
                <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('circulation.checkin.evidence.upload_label') }}</p>
                <p class="text-[9px] font-bold text-slate-300 uppercase">{{ $t('circulation.checkin.evidence.upload_hint') }}</p>
              </div>
            </div>

            <div class="grid grid-cols-2 gap-4">
              <div class="relative group aspect-square bg-slate-50 rounded-2xl overflow-hidden border border-slate-100 shadow-sm">
                <div class="absolute inset-0 bg-gradient-to-t from-black/20 to-transparent"></div>
                <p class="absolute bottom-3 left-3 text-[9px] font-black text-white uppercase tracking-widest">IMG_001.jpg</p>
              </div>
              <div class="relative group aspect-square bg-slate-50 rounded-2xl overflow-hidden border border-slate-100 shadow-sm">
                <div class="absolute inset-0 bg-gradient-to-t from-black/20 to-transparent"></div>
                <p class="absolute bottom-3 left-3 text-[9px] font-black text-white uppercase tracking-widest">IMG_002.jpg</p>
              </div>
            </div>

            <div class="space-y-4">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest flex items-center gap-2">
                {{ $t('circulation.checkin.evidence.admin_note') }} <span class="text-red-500">*</span>
              </label>
              <textarea :placeholder="$t('circulation.checkin.evidence.admin_note_placeholder')" 
                        rows="4" 
                        class="w-full p-6 bg-slate-50 border border-slate-100 rounded-3xl text-sm font-medium text-slate-600 outline-none focus:ring-4 focus:ring-primary-500/5 focus:border-primary-500 transition-all resize-none"></textarea>
            </div>
          </div>
        </div>

        <!-- Sticky Actions -->
        <div class="glass p-10 bg-white rounded-lg border border-slate-100 shadow-2xl shadow-slate-200/50 space-y-6">
          <button @click="showSuccessModal = true" class="w-full py-5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-[11px] font-black uppercase tracking-[0.2em] shadow-xl shadow-blue-900/20 transition-all flex items-center justify-center gap-3">
            <LucideCheckCircle2 class="w-4 h-4" /> {{ $t('circulation.checkin.footer.btn_complete') }}
          </button>
          <button class="w-full py-4 bg-white border border-slate-200 text-slate-400 hover:bg-slate-50 rounded-2xl text-[10px] font-black uppercase tracking-[0.2em] transition-all flex items-center justify-center gap-3 group">
            <LucideAlertTriangle class="w-4 h-4 group-hover:text-orange-500 transition-colors" /> {{ $t('circulation.checkin.footer.btn_save_issue') }}
          </button>
          <p class="text-[9px] font-bold text-slate-300 text-center uppercase leading-loose tracking-widest">
            {{ $t('circulation.checkin.footer.hint') }}
          </p>
        </div>
      </div>
    </div>

    <!-- Check-in Success Modal -->
    <Transition name="scale">
      <div v-if="showSuccessModal" class="fixed inset-0 z-[100] flex items-center justify-center p-6 bg-[#1E3A5F]/40 backdrop-blur-sm">
        <div class="glass max-w-md w-full bg-white rounded-lg p-12 shadow-2xl border border-white flex flex-col items-center text-center space-y-10" v-motion-pop>
          <!-- Success Icon -->
          <div class="w-24 h-24 bg-emerald-50 rounded-full flex items-center justify-center text-emerald-500 shadow-xl shadow-emerald-500/20">
            <LucideCheckCircle2 class="w-12 h-12" />
          </div>

          <!-- Message -->
          <div class="space-y-4">
            <h2 class="text-2xl font-black text-[#1E3A5F]">{{ $t('circulation.checkin.modal.title') }}</h2>
            <p class="text-xs font-bold text-slate-400 leading-relaxed px-4" v-html="$t('circulation.checkin.modal.desc')"></p>
          </div>

          <!-- Actions -->
          <div class="w-full space-y-4">
            <button @click="navigateTo('/dashboard')" class="w-full py-5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-sm font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 transition-all">
              {{ $t('circulation.checkin.modal.btn_dashboard') }}
            </button>
            <button class="w-full py-4 bg-white border border-slate-200 text-slate-400 hover:bg-slate-50 rounded-2xl text-xs font-black uppercase tracking-widest transition-all flex items-center justify-center gap-3">
              <LucidePrinter class="w-4 h-4" /> {{ $t('circulation.checkin.modal.btn_print') }}
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
  LucideInbox, LucideAlertTriangle, LucideChevronDown, 
  LucideCamera, LucideUploadCloud, LucideCheckCircle2, LucidePrinter
} from 'lucide-vue-next'

const showSuccessModal = ref(false)
const documents = ref([
  {
    title: 'Akta Pendirian Perusahaan PT. AKR',
    id: 'DOC-LEGAL-001',
    condition: 'Lengkap',
    note: ''
  },
  {
    title: 'Sertifikat HGB No. 452/Kebayoran',
    id: 'DOC-LEGAL-002',
    condition: 'Rusak',
    note: 'Ada noda air di pojok kanan'
  },
  {
    title: 'Surat Izin Usaha Perdagangan (SIUP)',
    id: 'DOC-LEGAL-104',
    condition: 'Lengkap',
    note: ''
  }
])
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
</style>
