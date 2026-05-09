<template>
  <div class="max-w-7xl mx-auto pb-20">
    <Transition name="fade" mode="out-in">
      <!-- STEP 1: FORM -->
      <div v-if="step === 1" class="max-w-3xl mx-auto space-y-10" v-motion-fade>
        <div class="space-y-4 text-center sm:text-left">
          <h1 class="text-3xl font-black text-[#1E3A5F] tracking-tight uppercase">{{ $t('loans.fast_track.step1.title') }}</h1>
          <p class="text-slate-500 font-medium">{{ $t('loans.fast_track.step1.subtitle') }}</p>
        </div>

        <div class="glass rounded-lg p-10 space-y-8 bg-white border border-slate-100 shadow-2xl shadow-slate-200/50">
          <div class="space-y-6">
            <!-- Jenis Keperluan -->
            <div class="space-y-3">
              <label class="text-[10px] font-black text-emerald-600 uppercase tracking-widest flex items-center gap-2">
                {{ $t('loans.fast_track.step1.form.purpose_label') }} <span class="text-red-500">*</span>
              </label>
              <div class="relative">
                <select v-model="form.purpose" class="w-full pl-5 pr-12 py-4 bg-slate-50 border border-slate-200 rounded-2xl text-sm font-bold text-slate-700 appearance-none outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all cursor-pointer">
                  <option value="" disabled>{{ $t('loans.fast_track.step1.form.purpose_placeholder') }}</option>
                  <option>{{ $t('loans.fast_track.step1.form.purpose_options.internal_audit') }}</option>
                  <option>{{ $t('loans.fast_track.step1.form.purpose_options.external_audit') }}</option>
                  <option>{{ $t('loans.fast_track.step1.form.purpose_options.legal_review') }}</option>
                  <option>{{ $t('loans.fast_track.step1.form.purpose_options.annual_tax') }}</option>
                </select>
                <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400 pointer-events-none" />
              </div>
            </div>

            <!-- Departemen Pemilik -->
            <div class="space-y-3">
              <label class="text-[10px] font-black text-emerald-600 uppercase tracking-widest">{{ $t('loans.fast_track.step1.form.dept_label') }}</label>
              <div class="relative">
                <select v-model="form.department" class="w-full pl-5 pr-12 py-4 bg-slate-50 border border-slate-200 rounded-2xl text-sm font-bold text-slate-700 appearance-none outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all cursor-pointer">
                  <option>{{ $t('loans.fast_track.step1.form.dept_options.all') }}</option>
                  <option>{{ $t('loans.fast_track.step1.form.dept_options.accounting') }}</option>
                  <option>{{ $t('loans.fast_track.step1.form.dept_options.finance') }}</option>
                  <option>{{ $t('loans.fast_track.step1.form.dept_options.legal') }}</option>
                </select>
                <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400 pointer-events-none" />
              </div>
            </div>

            <!-- Durasi Peminjaman -->
            <div class="space-y-3">
              <label class="text-[10px] font-black text-emerald-600 uppercase tracking-widest flex items-center gap-2">
                {{ $t('loans.fast_track.step1.form.duration_label') }} <span class="text-red-500">*</span>
              </label>
              <div class="relative max-w-[240px]">
                <input 
                  type="number" 
                  v-model="form.duration"
                  class="w-full pl-5 pr-16 py-4 bg-slate-50 border border-slate-200 rounded-2xl text-sm font-black text-slate-700 outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all"
                />
                <span class="absolute right-6 top-1/2 -translate-y-1/2 text-[9px] font-black text-slate-400 uppercase">{{ $t('loans.fast_track.step1.form.duration_suffix') }}</span>
              </div>
            </div>

            <!-- Catatan Keperluan -->
            <div class="space-y-3">
              <label class="text-[10px] font-black text-emerald-600 uppercase tracking-widest">{{ $t('loans.fast_track.step1.form.notes_label') }}</label>
              <textarea 
                v-model="form.notes"
                :placeholder="$t('loans.fast_track.step1.form.notes_placeholder')" 
                rows="5" 
                class="w-full p-6 bg-slate-50 border border-slate-200 rounded-[2rem] text-sm font-bold text-slate-700 outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all resize-none"
              ></textarea>
            </div>
          </div>

          <!-- Info Box -->
          <div class="p-6 bg-orange-50 border border-orange-100 rounded-2xl flex gap-4 items-start">
            <LucideHistory class="w-5 h-5 text-orange-500 shrink-0 mt-0.5" />
            <p class="text-[11px] font-medium text-orange-800 leading-relaxed">
              {{ $t('loans.fast_track.step1.form.info_box') }}
            </p>
          </div>
        </div>

        <!-- Actions -->
        <div class="flex items-center justify-between px-4">
          <button @click="navigateTo('/documents')" class="text-xs font-black uppercase tracking-[0.2em] text-slate-400 hover:text-slate-600 transition-colors">{{ $t('loans.fast_track.step1.form.btn_cancel') }}</button>
          <button 
            @click="nextStep"
            class="flex items-center gap-4 px-12 py-5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-xs font-black uppercase tracking-[0.2em] shadow-2xl shadow-blue-900/20 transition-all group"
          >
            {{ $t('loans.fast_track.step1.form.btn_next') }}
            <LucideArrowRight class="w-5 h-5 group-hover:translate-x-1 transition-transform" />
          </button>
        </div>
      </div>

      <!-- STEP 2: SELECTION -->
      <div v-else-if="step === 2" class="space-y-10" v-motion-fade>
        <div class="grid grid-cols-1 lg:grid-cols-12 gap-10 items-start">
          <div class="lg:col-span-8 space-y-12">
            <!-- Mandatory Section -->
            <div class="space-y-6">
              <div class="flex items-center justify-between">
                <div class="flex items-center gap-3">
                  <div class="w-1.5 h-8 bg-emerald-500 rounded-full"></div>
                  <h2 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('loans.fast_track.step2.sections.mandatory.title') }}</h2>
                </div>
                <span class="text-[10px] font-black text-emerald-600 bg-emerald-50 px-4 py-1.5 rounded-full border border-emerald-100 uppercase tracking-widest">{{ $t('loans.fast_track.step2.sections.mandatory.selected_count', { count: 3 }) }}</span>
              </div>
              
              <div class="glass rounded-lg overflow-hidden border border-slate-100 bg-white shadow-xl shadow-slate-200/50">
                <table class="w-full text-left">
                  <thead>
                    <tr class="bg-slate-50/50 border-b border-slate-100 text-[9px] font-black text-slate-400 uppercase tracking-widest">
                      <th class="px-6 py-5 w-12 text-center"><input type="checkbox" checked class="rounded" /></th>
                      <th class="px-6 py-5">{{ $t('loans.fast_track.step2.table.title') }}</th>
                      <th class="px-6 py-5">{{ $t('loans.fast_track.step2.table.id') }}</th>
                      <th class="px-6 py-5">{{ $t('loans.fast_track.step2.table.dept') }}</th>
                      <th class="px-6 py-5">{{ $t('loans.fast_track.step2.table.last_update') }}</th>
                      <th class="px-6 py-5">{{ $t('loans.fast_track.step2.table.status') }}</th>
                      <th class="px-6 py-5">{{ $t('loans.fast_track.step2.table.reason') }}</th>
                    </tr>
                  </thead>
                  <tbody class="divide-y divide-slate-50">
                    <tr v-for="doc in mandatoryDocs" :key="doc.id" class="group hover:bg-slate-50/30 transition-colors">
                      <td class="px-6 py-6 text-center"><input type="checkbox" v-model="doc.selected" class="rounded w-4 h-4 text-emerald-500" /></td>
                      <td class="px-6 py-6"><p class="text-sm font-black text-[#1E3A5F] uppercase group-hover:text-primary-600 transition-colors cursor-pointer leading-tight">{{ doc.title }}</p></td>
                      <td class="px-6 py-6"><p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ doc.id }}</p></td>
                      <td class="px-6 py-6"><p class="text-xs font-bold text-slate-500">{{ doc.dept }}</p></td>
                      <td class="px-6 py-6"><p class="text-xs font-bold text-slate-400">{{ doc.date }}</p></td>
                      <td class="px-6 py-6">
                        <div class="flex items-center gap-2">
                          <div class="w-2 h-2 rounded-full bg-green-500"></div>
                          <span class="text-[10px] font-bold text-green-600 uppercase">{{ $t('loans.fast_track.step2.table.available') }}</span>
                        </div>
                      </td>
                      <td class="px-6 py-6">
                        <span class="px-3 py-1 bg-emerald-50 text-emerald-600 rounded-lg text-[8px] font-black uppercase tracking-tight leading-tight block">
                          {{ $t('loans.fast_track.step2.sections.mandatory.reason') }}
                        </span>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>

            <!-- Recommendations Section -->
            <div class="space-y-6">
              <div class="flex items-center gap-3">
                <div class="w-1.5 h-8 bg-orange-500 rounded-full"></div>
                <h2 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('loans.fast_track.step2.sections.recommended.title') }}</h2>
              </div>
              <div class="glass rounded-lg overflow-hidden border border-slate-100 bg-white shadow-xl shadow-slate-200/50">
                <table class="w-full text-left">
                  <thead>
                    <tr class="bg-slate-50/50 border-b border-slate-100 text-[9px] font-black text-slate-400 uppercase tracking-widest">
                      <th class="px-6 py-5 w-12 text-center"><input type="checkbox" class="rounded" /></th>
                      <th class="px-6 py-5">DOC TITLE</th>
                      <th class="px-6 py-5">ID</th>
                      <th class="px-6 py-5">DEPT</th>
                      <th class="px-6 py-5">LAST UPDATE</th>
                      <th class="px-6 py-5">STATUS</th>
                      <th class="px-6 py-5">REASON</th>
                    </tr>
                  </thead>
                  <tbody class="divide-y divide-slate-50 text-slate-600">
                    <tr v-for="doc in recommendedDocs" :key="doc.id" class="group hover:bg-slate-50/30 transition-colors">
                      <td class="px-6 py-6 text-center"><input type="checkbox" v-model="doc.selected" class="rounded w-4 h-4" /></td>
                      <td class="px-6 py-6"><p class="text-sm font-black text-slate-600 uppercase group-hover:text-primary-600 transition-colors cursor-pointer leading-tight">{{ doc.title }}</p></td>
                      <td class="px-6 py-6"><p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ doc.id }}</p></td>
                      <td class="px-6 py-6"><p class="text-xs font-bold text-slate-500">{{ doc.dept }}</p></td>
                      <td class="px-6 py-6"><p class="text-xs font-bold text-slate-400">{{ doc.date }}</p></td>
                      <td class="px-6 py-6">
                        <div class="flex items-center gap-2">
                          <div :class="`w-2 h-2 rounded-full ${doc.status === 'Available' ? 'bg-green-500' : 'bg-orange-500'}`"></div>
                          <span :class="`text-[10px] font-bold uppercase ${doc.status === 'Available' ? 'text-green-600' : 'text-orange-600'}`">{{ doc.status === 'Available' ? $t('loans.fast_track.step2.table.available') : $t('loans.fast_track.step2.table.on_loan') }}</span>
                        </div>
                      </td>
                      <td class="px-6 py-6">
                        <span :class="`px-3 py-1 ${doc.reasonColor} rounded-lg text-[8px] font-black uppercase tracking-tight leading-tight block`">
                          {{ doc.reason }}
                        </span>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>

            <!-- Sensitive Section -->
            <div class="space-y-6">
              <div class="flex items-center gap-4">
                <div class="flex items-center gap-3">
                  <div class="w-1.5 h-8 bg-red-500 rounded-full"></div>
                  <h2 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('loans.fast_track.step2.sections.sensitive.title') }}</h2>
                </div>
                <span class="px-4 py-1.5 bg-red-50 text-red-600 rounded-full text-[9px] font-black uppercase tracking-widest border border-red-100 flex items-center gap-2">
                  <div class="w-1.5 h-1.5 rounded-full bg-red-500"></div>
                  {{ $t('loans.fast_track.step2.sections.sensitive.badge') }}
                </span>
              </div>
              <div class="glass rounded-lg overflow-hidden border border-slate-100 bg-white shadow-xl shadow-slate-200/50">
                <table class="w-full text-left">
                  <thead>
                    <tr class="bg-slate-50/50 border-b border-slate-100 text-[9px] font-black text-slate-400 uppercase tracking-widest">
                      <th class="px-6 py-5 w-12 text-center"><input type="checkbox" class="rounded" /></th>
                      <th class="px-6 py-5">DOC TITLE</th>
                      <th class="px-6 py-5">ID</th>
                      <th class="px-6 py-5">DEPT</th>
                      <th class="px-6 py-5">LAST UPDATE</th>
                      <th class="px-6 py-5">STATUS</th>
                      <th class="px-6 py-5">REASON</th>
                    </tr>
                  </thead>
                  <tbody class="divide-y divide-slate-50">
                    <tr v-for="doc in sensitiveDocs" :key="doc.id" class="group hover:bg-slate-50/30 transition-colors">
                      <td class="px-6 py-6 text-center"><input type="checkbox" v-model="doc.selected" class="rounded w-4 h-4" /></td>
                      <td class="px-6 py-6"><p class="text-sm font-black text-slate-600 uppercase group-hover:text-primary-600 transition-colors cursor-pointer leading-tight">{{ doc.title }}</p></td>
                      <td class="px-6 py-6"><p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ doc.id }}</p></td>
                      <td class="px-6 py-6"><p class="text-xs font-bold text-slate-500">{{ doc.dept }}</p></td>
                      <td class="px-6 py-6"><p class="text-xs font-bold text-slate-400">{{ doc.date }}</p></td>
                      <td class="px-6 py-6">
                        <div class="flex items-center gap-2">
                          <div class="w-2 h-2 rounded-full bg-red-500"></div>
                          <span class="text-[10px] font-bold text-red-600 uppercase">{{ $t('loans.fast_track.step2.table.restricted') }}</span>
                        </div>
                      </td>
                      <td class="px-6 py-6">
                        <span class="px-3 py-1 bg-slate-100 text-slate-500 rounded-lg text-[8px] font-black uppercase tracking-tight leading-tight block">
                          {{ $t('loans.fast_track.step2.sections.sensitive.reason') }}
                        </span>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>

          <!-- Summary Panel -->
          <div class="lg:col-span-4 space-y-8 sticky top-8">
            <div class="glass rounded-lg overflow-hidden bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 flex flex-col">
              <div class="bg-slate-50 px-8 py-6 border-b border-slate-100">
                <h2 class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">{{ $t('loans.fast_track.step2.summary.title') }}</h2>
              </div>
              
              <div class="p-8 space-y-8">
                <div class="space-y-6">
                  <div class="flex justify-between items-center">
                    <span class="text-sm font-bold text-slate-400">{{ $t('loans.fast_track.step2.summary.total_docs') }}</span>
                    <span class="text-sm font-black text-[#1E3A5F]">{{ $t('loans.fast_track.step2.summary.total_selected', { count: 2 }) }}</span>
                  </div>
                  <div class="flex justify-between items-center">
                    <span class="text-sm font-bold text-slate-400">{{ $t('loans.fast_track.step2.summary.duration') }}</span>
                    <span class="text-sm font-black text-[#1E3A5F]">{{ $t('loans.fast_track.step2.summary.duration_value', { count: 3 }) }}</span>
                  </div>
                  <div class="flex justify-between items-center">
                    <span class="text-sm font-bold text-slate-400">{{ $t('loans.fast_track.step2.summary.est_return') }}</span>
                    <span class="text-sm font-black text-emerald-600">18 Maret 2024</span>
                  </div>
                </div>

                <div class="w-full h-px bg-slate-100"></div>

                <!-- Risk Alert -->
                <div class="p-6 bg-amber-50 border border-amber-100 rounded-2xl space-y-3">
                  <div class="flex items-center gap-2 text-amber-600">
                    <LucideAlertTriangle class="w-4 h-4" />
                    <span class="text-[10px] font-black uppercase tracking-widest">{{ $t('loans.fast_track.step2.summary.risk_title') }}</span>
                  </div>
                  <p class="text-[10px] font-medium text-amber-700 leading-relaxed">
                    {{ $t('loans.fast_track.step2.summary.risk_desc') }}
                  </p>
                </div>

                <div class="space-y-3">
                  <button @click="step = 3" class="w-full py-5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 transition-all">
                    {{ $t('loans.fast_track.step2.summary.btn_review', { count: 2 }) }}
                  </button>
                  <button @click="step = 1" class="w-full py-4 bg-white border border-slate-200 text-slate-600 rounded-2xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all text-center block">
                    {{ $t('loans.fast_track.step2.summary.btn_back') }}
                  </button>
                </div>
              </div>
            </div>

            <p class="text-[10px] font-bold text-slate-400 text-center italic">{{ $t('loans.fast_track.step2.summary.hint') }}</p>
          </div>
        </div>
      </div>
      <!-- STEP 3: REVIEW & SUBMIT -->
      <div v-else-if="step === 3" class="space-y-8" v-motion-fade>
        <div class="flex items-center justify-between">
          <h1 class="text-2xl font-black text-[#1E3A5F]">{{ $t('loans.fast_track.step3.title') }}</h1>
          <span class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('loans.fast_track.step3.subtitle') }}</span>
        </div>

        <!-- Request Summary -->
        <div class="glass rounded-3xl p-10 bg-white border border-slate-100 shadow-xl relative overflow-hidden">
          <div class="absolute top-6 right-6 px-4 py-1.5 bg-emerald-50 text-emerald-600 rounded-full text-[10px] font-black uppercase tracking-widest border border-emerald-100">
            {{ $t('loans.fast_track.step3.request_summary.badge') }}
          </div>
          <h2 class="text-xs font-black text-slate-300 uppercase tracking-[0.2em] mb-10">{{ $t('loans.fast_track.step3.request_summary.title') }}</h2>
          <div class="grid grid-cols-2 md:grid-cols-4 gap-12">
            <div class="space-y-2">
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('loans.fast_track.step3.request_summary.purpose') }}</p>
              <p class="text-sm font-black text-[#1E3A5F] uppercase">{{ form.purpose }}</p>
            </div>
            <div class="space-y-2">
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('loans.fast_track.step3.request_summary.duration') }}</p>
              <p class="text-sm font-black text-[#1E3A5F] uppercase">{{ $t('loans.fast_track.step3.request_summary.duration_value', { count: form.duration }) }}</p>
            </div>
            <div class="space-y-2">
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('loans.fast_track.step3.request_summary.applicant') }}</p>
              <p class="text-sm font-black text-[#1E3A5F]">Andi Pratama</p>
            </div>
            <div class="space-y-2">
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('loans.fast_track.step3.request_summary.dept') }}</p>
              <p class="text-sm font-black text-[#1E3A5F]">Internal Audit - HQ</p>
            </div>
          </div>
        </div>

        <!-- Selected Documents -->
        <div class="glass rounded-3xl overflow-hidden bg-white border border-slate-100 shadow-xl">
          <div class="px-10 py-8 border-b border-slate-100">
            <h2 class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('loans.fast_track.step3.docs_table.title') }} <span class="text-slate-400 ml-2">{{ $t('loans.fast_track.step3.docs_table.count', { count: 2 }) }}</span></h2>
          </div>
          <table class="w-full text-left">
            <thead>
              <tr class="bg-slate-50/30 text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100">
                <th class="px-10 py-4">{{ $t('loans.fast_track.step3.docs_table.headers.no') }}</th>
                <th class="px-10 py-4">{{ $t('loans.fast_track.step3.docs_table.headers.id') }}</th>
                <th class="px-10 py-4">{{ $t('loans.fast_track.step3.docs_table.headers.title') }}</th>
                <th class="px-10 py-4">{{ $t('loans.fast_track.step3.docs_table.headers.dept') }}</th>
                <th class="px-10 py-4">{{ $t('loans.fast_track.step3.docs_table.headers.sensitivity') }}</th>
                <th class="px-10 py-4">{{ $t('loans.fast_track.step3.docs_table.headers.status') }}</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50">
              <tr v-for="(doc, i) in mandatoryDocs" :key="doc.id" class="text-xs">
                <td class="px-10 py-6 font-bold text-slate-400">{{ i + 1 }}</td>
                <td class="px-10 py-6 font-bold text-slate-400 uppercase">{{ doc.id }}</td>
                <td class="px-10 py-6 font-black text-[#1E3A5F] uppercase tracking-tight">{{ doc.title }}</td>
                <td class="px-10 py-6 font-bold text-slate-500 uppercase">{{ doc.dept }}</td>
                <td class="px-10 py-6">
                  <span :class="`px-3 py-1 rounded-lg text-[9px] font-black uppercase tracking-widest ${i === 0 ? 'bg-red-50 text-red-500 border border-red-100' : 'bg-amber-50 text-amber-500 border border-amber-100'}`">
                    {{ i === 0 ? 'CONFIDENTIAL' : 'INTERNAL' }}
                  </span>
                </td>
                <td class="px-10 py-6">
                  <div class="flex items-center gap-2 text-green-600">
                    <div class="w-1.5 h-1.5 rounded-full bg-green-500"></div>
                    <span class="font-black uppercase tracking-widest">{{ $t('loans.fast_track.step3.docs_table.available') }}</span>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-10">
          <!-- Policy Section -->
          <div class="glass rounded-3xl p-10 bg-white border border-slate-100 space-y-8">
            <h2 class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('loans.fast_track.step3.policy.title') }}</h2>
            <div class="p-6 bg-blue-50/50 border border-blue-100 rounded-2xl">
              <p class="text-[11px] font-medium text-slate-600 leading-loose" v-html="$t('loans.fast_track.step3.policy.desc')"></p>
            </div>
            <label class="flex items-start gap-4 cursor-pointer group">
              <div class="relative mt-1">
                <input type="checkbox" v-model="form.agree" class="peer appearance-none w-5 h-5 border-2 border-slate-200 rounded-md checked:bg-primary-600 checked:border-primary-600 transition-all" />
                <LucideCheck class="absolute inset-0 w-5 h-5 text-white scale-0 peer-checked:scale-100 transition-transform p-1" />
              </div>
              <span class="text-xs font-bold text-slate-500 group-hover:text-slate-700 transition-colors leading-relaxed">
                {{ $t('loans.fast_track.step3.policy.agree_label') }}
              </span>
            </label>
          </div>

          <!-- Notes for Approver -->
          <div class="glass rounded-3xl p-10 bg-white border border-slate-100 space-y-8">
            <h2 class="text-sm font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('loans.fast_track.step3.approver.title') }}</h2>
            <div class="space-y-4">
              <textarea 
                v-model="form.approverNotes"
                :placeholder="$t('loans.fast_track.step3.approver.placeholder')" 
                rows="6" 
                class="w-full p-6 bg-slate-50 border border-slate-200 rounded-3xl text-sm font-bold text-slate-700 outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all resize-none"
              ></textarea>
              <p class="text-[10px] font-bold text-slate-400 italic">
                <LucideInfo class="inline-block w-3 h-3 mr-1" />
                {{ $t('loans.fast_track.step3.approver.hint') }}
              </p>
            </div>
          </div>
        </div>

        <!-- Footer Actions -->
        <div class="flex items-center justify-between pt-8 border-t border-slate-100">
          <button @click="step = 2" class="flex items-center gap-2 text-xs font-black text-slate-400 hover:text-slate-600 uppercase tracking-widest group">
            <LucideArrowLeft class="w-4 h-4 group-hover:-translate-x-1 transition-transform" />
            {{ $t('loans.fast_track.step3.footer.btn_back') }}
          </button>
          <button 
            @click="submitRequest"
            :disabled="!form.agree"
            class="px-12 py-5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-xs font-black uppercase tracking-[0.2em] shadow-2xl shadow-blue-900/20 transition-all disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {{ $t('loans.fast_track.step3.footer.btn_submit') }}
          </button>
        </div>
      </div>
      <!-- SUCCESS VIEW -->
      <div v-else class="max-w-4xl mx-auto py-10" v-motion-pop>
        <div class="glass rounded-[3rem] p-16 bg-white border border-slate-100 shadow-2xl shadow-slate-200/50 flex flex-col items-center text-center space-y-12 relative overflow-hidden">
          <!-- Success Header -->
          <div class="space-y-6">
            <div class="w-24 h-24 bg-emerald-500 rounded-full flex items-center justify-center text-white mx-auto shadow-xl shadow-emerald-500/20">
              <LucideCheck class="w-12 h-12" />
            </div>
            <div class="space-y-2">
              <h1 class="text-4xl font-black text-[#1E3A5F] tracking-tight">{{ $t('loans.fast_track.success.title') }}</h1>
              <div class="inline-block px-4 py-1.5 bg-slate-50 border border-slate-200 rounded-lg text-xs font-bold text-slate-400 uppercase tracking-[0.2em] mb-4">
                {{ $t('loans.fast_track.success.id_label', { id: 'LOAN-2026-03-00125' }) }}
              </div>
              <div class="flex items-center justify-center gap-2 text-emerald-600 font-bold">
                <LucideHistory class="w-4 h-4" />
                <span class="text-sm">{{ $t('loans.fast_track.success.status_waiting') }}</span>
              </div>
            </div>
          </div>

          <!-- Timeline Stepper -->
          <div class="w-full max-w-2xl pt-10 pb-10">
            <div class="relative flex justify-between">
              <!-- Background Lines -->
              <div class="absolute top-5 left-0 w-full h-0.5 bg-slate-100"></div>
              <div class="absolute top-5 left-0 w-1/3 h-0.5 bg-emerald-500"></div>

              <!-- Step 1: Submitted -->
              <div class="relative z-10 flex flex-col items-center gap-3">
                <div class="w-10 h-10 rounded-full bg-emerald-500 text-white flex items-center justify-center shadow-lg shadow-emerald-500/20">
                  <LucideCheck class="w-5 h-5" />
                </div>
                <div class="text-center space-y-1">
                  <p class="text-[10px] font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('loans.fast_track.success.timeline.step1') }}</p>
                  <p class="text-[9px] font-bold text-slate-400">14 Mar, 09:30</p>
                </div>
              </div>

              <!-- Step 2: Level 1 -->
              <div class="relative z-10 flex flex-col items-center gap-3">
                <div class="w-10 h-10 rounded-full bg-white border-2 border-emerald-500 text-emerald-500 flex items-center justify-center shadow-sm">
                  <LucideUser class="w-5 h-5" />
                </div>
                <div class="text-center space-y-1">
                  <p class="text-[10px] font-black text-emerald-600 uppercase tracking-tight">{{ $t('loans.fast_track.success.timeline.step2') }}</p>
                  <p class="text-[9px] font-black text-emerald-400 uppercase tracking-widest animate-pulse">{{ $t('loans.fast_track.success.timeline.step2_status') }}</p>
                </div>
              </div>

              <!-- Step 3: Level 2 -->
              <div class="relative z-10 flex flex-col items-center gap-3 opacity-40">
                <div class="w-10 h-10 rounded-full bg-white border-2 border-slate-200 text-slate-400 flex items-center justify-center">
                  <LucideGavel class="w-5 h-5" />
                </div>
                <div class="text-center space-y-1">
                  <p class="text-[10px] font-black text-slate-500 uppercase tracking-tight">{{ $t('loans.fast_track.success.timeline.step3') }}</p>
                  <p class="text-[9px] font-black text-slate-300 uppercase tracking-widest">{{ $t('loans.fast_track.success.timeline.step3_status') }}</p>
                </div>
              </div>

              <!-- Step 4: Ready -->
              <div class="relative z-10 flex flex-col items-center gap-3 opacity-40">
                <div class="w-10 h-10 rounded-full bg-white border-2 border-slate-200 text-slate-400 flex items-center justify-center">
                  <LucideBox class="w-5 h-5" />
                </div>
                <div class="text-center space-y-1">
                  <p class="text-[10px] font-black text-slate-500 uppercase tracking-tight">{{ $t('loans.fast_track.success.timeline.step4') }}</p>
                  <p class="text-[9px] font-black text-slate-300 uppercase tracking-widest">{{ $t('loans.fast_track.success.timeline.step4_status') }}</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Actions -->
          <div class="flex flex-col sm:flex-row gap-4 w-full justify-center pt-8 border-t border-slate-50">
            <button class="flex items-center justify-center gap-4 px-10 py-5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-sm font-black uppercase tracking-widest shadow-2xl shadow-blue-900/20 transition-all group">
              <LucideRadar class="w-5 h-5 group-hover:scale-110 transition-transform" />
              {{ $t('loans.fast_track.success.actions.btn_tracking') }}
            </button>
            <button @click="navigateTo('/dashboard')" class="flex items-center justify-center gap-4 px-10 py-5 bg-white border border-slate-200 text-slate-600 rounded-2xl text-sm font-black uppercase tracking-widest hover:bg-slate-50 transition-all group">
              <LucideHome class="w-5 h-5 group-hover:scale-110 transition-transform" />
              {{ $t('loans.fast_track.success.actions.btn_dashboard') }}
            </button>
          </div>
        </div>

        <!-- Support Footer -->
        <div class="flex items-center justify-center gap-2 pt-10 text-slate-400">
          <LucideHelpCircle class="w-4 h-4" />
          <p class="text-[11px] font-bold">{{ $t('loans.fast_track.success.footer_help') }}</p>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { 
  LucideChevronDown, LucideHistory, LucideArrowRight, LucideAlertTriangle, LucideShieldCheck,
  LucideArrowLeft, LucideCheck, LucideInfo, LucideUser, LucideGavel, LucideBox, LucideRadar, LucideHome, LucideHelpCircle
} from 'lucide-vue-next'

const step = ref(1)

const form = reactive({
  purpose: '',
  department: 'Semua',
  duration: 3,
  notes: '',
  agree: false,
  approverNotes: ''
})

const mandatoryDocs = ref([
  { id: 'DOC-2024-FIN-001', title: 'Laporan Laba Rugi Q4 2023', dept: 'Finance', date: '12 Jan 2024', selected: true },
  { id: 'DOC-2024-LGL-042', title: 'Akta Perubahan Anggaran Dasar No. 12', dept: 'Legal', date: '05 Feb 2024', selected: true }
])

const recommendedDocs = ref([
  { id: 'PROC-2022-005', title: 'SOP Pengadaan Barang 2022', dept: 'Procurement', date: '10 Dec 2022', status: 'Available', reason: 'HISTORIS AUDIT SEBELUMNYA', reasonColor: 'bg-orange-50 text-orange-600', selected: false },
  { id: 'PAY-2023-882', title: 'Bukti Transfer Vendor Global', dept: 'Finance', date: '22 Jan 2024', status: 'On Loan', reason: 'TERKAIT LAPORAN Q4', reasonColor: 'bg-orange-50 text-orange-600', selected: false }
])

const sensitiveDocs = ref([
  { id: 'LGL-2023-015', title: 'Master Service Agreement - Tech Corp', dept: 'Legal', date: '14 Aug 2023', selected: false }
])

const nextStep = () => {
  if (form.purpose) {
    step.value = 2
  } else {
    alert('Silakan pilih jenis keperluan.')
  }
}

const submitRequest = () => {
  step.value = 4
}
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
.fade-enter-active, .fade-leave-active { transition: opacity 0.3s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; }
</style>
