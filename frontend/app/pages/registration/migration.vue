<template>
  <div class="h-screen flex flex-col bg-slate-50 dark:bg-slate-950 overflow-hidden">
    <!-- Header -->
    <header v-if="state !== 'success'" class="bg-white dark:bg-slate-900 border-b border-slate-200 dark:border-slate-800 px-8 py-5 flex items-center justify-between shadow-sm z-10" v-motion-fade>
      <div class="flex items-center gap-8">
        <div class="flex flex-col gap-1">
          <div v-if="['review', 'mandatory', 'duplicate_check', 'location_assignment', 'print_qr'].includes(state)" class="flex items-center gap-2 text-[9px] font-black text-slate-400 uppercase tracking-widest mb-1">
            {{ $t('registration.migration.ocr_review.breadcrumbs') }}
          </div>
          <h1 class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
            <template v-if="state === 'print_qr'">{{ $t('registration.migration.print_qr.title') }}</template>
            <template v-else-if="state === 'duplicate_check'">{{ $t('registration.migration.duplicate_check.title') }}</template>
            <template v-else-if="state === 'location_assignment'">{{ $t('registration.migration.location_assignment.title') }}</template>
            <template v-else-if="state === 'review'">{{ $t('registration.migration.ocr_review.title') }}</template>
            <template v-else-if="state === 'mandatory'">{{ $t('registration.migration.mandatory.title') }}</template>
            <template v-else>{{ $t('registration.migration.title') }}</template>
          </h1>
          <p class="text-xs font-bold text-slate-500">
            <template v-if="state === 'print_qr'">{{ $t('registration.migration.print_qr.subtitle') }}</template>
            <template v-else-if="state === 'duplicate_check'">{{ $t('registration.migration.duplicate_check.subtitle') }}</template>
            <template v-else-if="state === 'location_assignment'">{{ $t('registration.migration.location_assignment.subtitle') }}</template>
            <template v-else-if="state === 'review'">{{ $t('registration.migration.ocr_review.subtitle') }}</template>
            <template v-else-if="state === 'mandatory'">{{ $t('registration.migration.mandatory.subtitle') }}</template>
          </p>
        </div>

        <div v-if="state === 'initial'" class="flex items-center gap-10 ml-4">
          <div class="h-10 w-px bg-slate-100 dark:bg-slate-800"></div>
          <div class="space-y-1">
            <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.meta.filename') }}</p>
            <p class="text-xs font-black text-[#1E3A5F] dark:text-white">legacy_archive_2023_01.pdf</p>
          </div>
          <div class="space-y-1">
            <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.meta.pages') }}</p>
            <p class="text-xs font-black text-[#1E3A5F] dark:text-white">24 {{ $t('registration.migration.meta.pages').toLowerCase() }}</p>
          </div>
        </div>
      </div>

      <div class="flex items-center gap-6">
        <!-- Search Bar for Print State -->
        <div v-if="state === 'print_qr'" class="flex items-center gap-4">
          <div class="relative w-80">
            <LucideSearch class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
            <input type="text" :placeholder="$t('registration.migration.print_qr.search_placeholder')" class="w-full pl-11 pr-5 py-2.5 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all" />
          </div>
          <button class="w-10 h-10 bg-red-50 text-red-500 rounded-xl flex items-center justify-center relative shadow-sm">
            <LucideBell class="w-5 h-5" />
            <span class="absolute top-2.5 right-2.5 w-2 h-2 bg-red-600 rounded-full border-2 border-red-50"></span>
          </button>
        </div>

        <template v-if="state === 'initial'">
          <button class="px-6 py-3 bg-slate-100 dark:bg-slate-800 text-[#1E3A5F] dark:text-white rounded-xl text-xs font-black uppercase tracking-widest hover:bg-slate-200 transition-all flex items-center gap-2">
            <LucideMaximize2 class="w-4 h-4" />
            {{ $t('registration.migration.actions.preview') }}
          </button>
          <button @click="state = 'review'" class="px-8 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-2 active:scale-95">
            <LucideSparkles class="w-4 h-4" />
            {{ $t('registration.migration.actions.ocr_suggest') }}
          </button>
        </template>
        <template v-else-if="state === 'review'">
          <button class="px-8 py-3 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-800 text-slate-600 dark:text-slate-300 rounded-xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all">
            {{ $t('registration.migration.ocr_review.btn_save') }}
          </button>
          <button @click="state = 'mandatory'" class="px-10 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all active:scale-95">
            {{ $t('registration.migration.ocr_review.btn_next') }}
          </button>
        </template>
        <template v-else-if="state === 'location_assignment'">
          <button class="px-8 py-3 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-800 text-slate-600 dark:text-slate-300 rounded-xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all">
            {{ $t('registration.migration.location_assignment.btn_recalculate') }}
          </button>
          <button @click="state = 'success'" class="px-10 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all active:scale-95">
            {{ $t('registration.migration.location_assignment.btn_save') }}
          </button>
        </template>
      </div>
    </header>

    <!-- Main Content -->
    <main class="flex-grow flex overflow-hidden">
      <!-- Split Screen Layout -->
      <div v-if="['initial', 'review', 'mandatory'].includes(state)" class="flex flex-grow overflow-hidden">
        <!-- Left: PDF Viewer Mock -->
        <div class="w-[60%] bg-slate-200 dark:bg-slate-900 flex flex-col border-r border-slate-200 dark:border-slate-800" v-motion-slide-left>
          <div class="bg-slate-800 text-white px-6 py-3 flex items-center justify-between border-b border-slate-700">
            <div class="flex items-center gap-6">
              <button class="p-1 hover:bg-slate-700 rounded transition-colors"><LucideMenu class="w-5 h-5" /></button>
              <div class="flex items-center gap-3">
                <span class="text-xs font-bold uppercase tracking-widest opacity-60">Page</span>
                <div class="flex items-center gap-1">
                  <input type="text" value="1" class="w-10 h-7 bg-slate-900 border border-slate-600 rounded text-center text-xs font-black" />
                  <span class="text-xs font-bold opacity-60">/ 12</span>
                </div>
              </div>
            </div>
            <div class="flex items-center gap-8">
              <div class="flex items-center gap-4 bg-slate-900/50 px-4 py-1.5 rounded-lg border border-slate-700">
                <button class="p-1 hover:bg-slate-700 rounded transition-colors"><LucideMinus class="w-4 h-4" /></button>
                <span class="text-[11px] font-black w-12 text-center uppercase tracking-widest">100%</span>
                <button class="p-1 hover:bg-slate-700 rounded transition-colors"><LucidePlus class="w-4 h-4" /></button>
              </div>
              <div class="h-6 w-px bg-slate-700"></div>
              <div class="flex items-center gap-4">
                <button class="p-1 hover:bg-slate-700 rounded transition-colors"><LucidePrinter class="w-5 h-5" /></button>
                <button class="p-1 hover:bg-slate-700 rounded transition-colors"><LucideDownload class="w-5 h-5" /></button>
                <button class="p-1 hover:bg-slate-700 rounded transition-colors"><LucideExpand class="w-5 h-5" /></button>
              </div>
            </div>
          </div>

          <div class="flex-grow overflow-auto p-12 flex justify-center bg-slate-600/20 relative">
            <div class="bg-white shadow-2xl min-w-[700px] h-fit p-16 relative">
              <div class="space-y-12">
                <div class="flex justify-between">
                  <div class="w-32 h-10 bg-slate-100 rounded"></div>
                  <div class="text-right">
                    <div class="w-24 h-4 bg-slate-100 rounded ml-auto mb-2"></div>
                    <div class="w-16 h-3 bg-slate-50 rounded ml-auto"></div>
                  </div>
                </div>
                <div class="text-center py-10 space-y-4">
                  <h2 class="text-2xl font-black text-slate-800 tracking-widest border-b-2 border-slate-800 inline-block px-10 pb-2">SURAT PERJANJIAN KERJASAMA</h2>
                </div>
                <div class="space-y-8">
                  <div class="w-full h-4 bg-slate-50 rounded"></div>
                  <div class="w-4/5 h-4 bg-slate-50 rounded"></div>
                </div>
              </div>
              <template v-if="state === 'review'">
                <div class="absolute top-[215px] left-1/2 -translate-x-1/2 w-[500px] h-12 bg-primary-500/10 border-2 border-primary-500/50 rounded"></div>
              </template>
            </div>
          </div>
        </div>

        <div class="flex-grow bg-white dark:bg-slate-900 overflow-y-auto custom-scrollbar shadow-2xl shadow-slate-900/5" v-motion-slide-right>
          <div v-if="state === 'initial'" class="p-12 space-y-10 max-w-2xl mx-auto">
            <h3 class="text-xs font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.form.title') }}</h3>
            <div class="space-y-8">
              <div class="space-y-3">
                <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.form.fields.type') }}</label>
                <select class="w-full px-5 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-700 rounded-2xl text-sm font-bold outline-none"><option value="" disabled selected>{{ $t('registration.migration.form.fields.type_placeholder') }}</option></select>
              </div>
              <div class="space-y-3">
                <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.form.fields.number') }}</label>
                <input type="text" class="w-full px-5 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-700 rounded-2xl text-sm font-bold outline-none" />
              </div>
            </div>
          </div>

          <div v-else-if="state === 'review'" class="p-12 space-y-10 max-w-2xl mx-auto">
            <div class="flex items-center justify-between">
              <h3 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('registration.migration.ocr_review.verify_title') }}</h3>
              <span class="px-3 py-1 bg-green-50 text-green-500 rounded-md text-[9px] font-black uppercase">{{ $t('registration.migration.ocr_review.accuracy', { value: 94 }) }}</span>
            </div>
            <div class="space-y-8">
              <div class="space-y-3">
                <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.ocr_review.fields.title') }}</label>
                <input type="text" value="SURAT PERJANJIAN KERJASAMA" class="w-full px-5 py-4 bg-slate-50 border border-primary-500/20 rounded-2xl text-sm font-black text-[#1E3A5F]" />
              </div>
            </div>
          </div>

          <div v-else-if="state === 'mandatory'" class="p-12 space-y-10 max-w-2xl mx-auto">
            <div class="flex items-center justify-between">
              <h3 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase border-l-4 border-primary-500 pl-4">{{ $t('registration.migration.mandatory.form.title') }}</h3>
              <span class="px-3 py-1 bg-green-50 text-green-600 rounded-md text-[8px] font-black uppercase tracking-widest">{{ $t('registration.migration.mandatory.form.required_badge') }}</span>
            </div>
            <div class="space-y-8">
              <div class="space-y-3">
                <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.mandatory.form.fields.dept') }} <span class="text-red-500">*</span></label>
                <select class="w-full px-5 py-4 bg-white dark:bg-slate-800 border border-slate-200 rounded-2xl text-sm font-bold outline-none"><option value="" disabled selected>{{ $t('registration.migration.mandatory.form.fields.placeholders.dept') }}</option></select>
              </div>
              <div class="flex gap-4 pt-6">
                <button @click="state = 'review'" class="flex-grow py-4 bg-white border border-slate-200 rounded-2xl text-xs font-black uppercase text-slate-500">{{ $t('registration.migration.mandatory.form.btn_reset') }}</button>
                <button @click="state = 'duplicate_check'" class="flex-[2] py-4 bg-[#1E3A5F] text-white rounded-2xl text-xs font-black uppercase shadow-xl shadow-blue-900/20">{{ $t('registration.migration.mandatory.form.btn_validate') }}</button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Dashboard Layouts -->
      <div v-else class="flex-grow overflow-y-auto bg-slate-50 dark:bg-slate-950 p-10 custom-scrollbar" v-motion-fade>
        <div class="max-w-6xl mx-auto space-y-10 pb-20">
          
          <!-- State: Duplicate Check -->
          <div v-if="state === 'duplicate_check'" class="space-y-10">
            <div class="bg-green-50/50 dark:bg-green-900/10 border border-green-100 rounded-lg p-8 flex items-center justify-between">
              <div class="space-y-1">
                <h3 class="text-sm font-black text-green-800 dark:text-green-300 uppercase">{{ $t('registration.migration.duplicate_check.alert.title') }}</h3>
                <p class="text-xs font-medium text-green-700/70 dark:text-green-400">{{ $t('registration.migration.duplicate_check.alert.desc') }}</p>
              </div>
              <button class="px-8 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase">{{ $t('registration.migration.duplicate_check.alert.btn_resolve') }}</button>
            </div>
            <div class="grid grid-cols-1 lg:grid-cols-3 gap-10">
              <div class="lg:col-span-2 space-y-8">
                <div class="glass p-10 rounded-lg space-y-8">
                  <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase">{{ $t('registration.migration.duplicate_check.similar.title') }}</h4>
                  <div v-for="rec in similarRecords" :key="rec.id" class="p-8 bg-slate-50 border border-slate-100 rounded-lg flex items-center justify-between group transition-all">
                    <div class="flex items-center gap-8">
                      <div class="w-16 h-16 rounded-full border-4 border-green-500 flex items-center justify-center text-xs font-black">{{ rec.score }}%</div>
                      <div>
                        <p class="text-sm font-black text-[#1E3A5F] uppercase">{{ rec.filename }}</p>
                        <p class="text-[10px] font-bold text-slate-400 uppercase">Found in: {{ rec.path }}</p>
                      </div>
                    </div>
                    <LucideEye class="w-6 h-6 text-slate-300 group-hover:text-primary-500 transition-colors" />
                  </div>
                </div>
              </div>
              <div class="space-y-8">
                <button @click="state = 'location_assignment'" class="w-full py-5 bg-[#1E3A5F] text-white rounded-lg text-xs font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40">{{ $t('registration.migration.duplicate_check.footer.btn_next') }}</button>
              </div>
            </div>
          </div>

          <!-- State: Location Assignment -->
          <div v-else-if="state === 'location_assignment'" class="space-y-10">
            <div class="flex items-center gap-10 border-b border-slate-100 pb-2">
              <button class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase border-b-4 border-primary-500 pb-4">{{ $t('registration.migration.location_assignment.tabs.auto') }}</button>
              <button class="text-xs font-black text-slate-400 uppercase pb-4">{{ $t('registration.migration.location_assignment.tabs.manual') }}</button>
            </div>
            <div class="grid grid-cols-1 lg:grid-cols-3 gap-10">
              <div class="lg:col-span-2 space-y-8">
                <div class="glass p-10 rounded-lg space-y-10">
                  <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase">{{ $t('registration.migration.location_assignment.strategy.title') }}</h4>
                  <div class="grid grid-cols-2 gap-8">
                    <div class="p-8 bg-slate-50 dark:bg-slate-900/50 rounded-lg space-y-2">
                      <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.location_assignment.strategy.target') }}</p>
                      <p class="text-lg font-black text-[#1E3A5F] dark:text-white">{{ $t('registration.migration.location_assignment.strategy.archives') }}</p>
                    </div>
                  </div>
                </div>
                <div class="glass rounded-lg overflow-hidden">
                  <table class="w-full text-left">
                    <tbody class="divide-y divide-slate-50 dark:divide-slate-800/50">
                      <tr v-for="loc in locations" :key="loc.id">
                        <td class="p-8 px-10"><p class="text-sm font-black uppercase text-[#1E3A5F] dark:text-white">{{ loc.code }}</p></td>
                        <td class="p-8"><div class="w-48 h-2 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden"><div :style="{ width: loc.percent + '%' }" class="h-full bg-green-500 rounded-full"></div></div></td>
                        <td class="p-8 text-right px-10 font-black text-[10px] text-[#1E3A5F] dark:text-white uppercase tracking-widest">Backup • Restore</td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div class="space-y-8">
                <div class="bg-[#1E3A5F] p-10 rounded-lg text-white flex flex-col items-center">
                  <div class="w-32 h-32 rounded-full border-8 border-white/10 flex items-center justify-center text-3xl font-black">78%</div>
                </div>
              </div>
            </div>
          </div>

          <!-- State: Success -->
          <div v-else-if="state === 'success'" class="bg-white dark:bg-slate-900 rounded-[3rem] p-0 overflow-hidden shadow-2xl border border-slate-100 dark:border-slate-800">
            <div class="bg-green-50/30 dark:bg-green-900/5 p-16 flex flex-col items-center text-center space-y-6">
              <div class="w-20 h-20 rounded-full bg-green-500 flex items-center justify-center text-white shadow-xl shadow-green-500/20"><LucideCheckCircle2 class="w-10 h-10" /></div>
              <div class="space-y-2">
                <h2 class="text-4xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">{{ $t('registration.migration.success.title') }}</h2>
                <p class="text-sm font-medium text-slate-500 max-w-xl mx-auto">{{ $t('registration.migration.success.subtitle') }}</p>
              </div>
            </div>
            <div class="p-16 grid grid-cols-1 lg:grid-cols-3 gap-16">
              <div class="lg:col-span-2 space-y-12">
                <div class="p-10 bg-slate-50 dark:bg-slate-800/50 rounded-lg space-y-8">
                  <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.success.renaming.final') }}</p>
                  <div class="flex items-center justify-between p-5 bg-white dark:bg-slate-900 border border-slate-100 dark:border-slate-700 rounded-2xl shadow-sm">
                    <span class="text-sm font-black text-primary-500">DMS-INV-2024-00892-CORP.pdf</span>
                    <LucideCopy class="w-4 h-4 text-slate-400 cursor-pointer" />
                  </div>
                </div>
              </div>
              <div class="space-y-12">
                <button @click="state = 'print_qr'" class="w-full py-5 bg-[#1E3A5F] text-white rounded-lg text-xs font-black uppercase shadow-2xl shadow-blue-900/40 transition-all">
                  {{ $t('registration.migration.success.buttons.print') }}
                </button>
                <button @click="state = 'initial'" class="w-full py-5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-lg text-xs font-black uppercase text-[#1E3A5F] flex items-center justify-center gap-3">
                  <LucidePlus class="w-4 h-4" /> {{ $t('registration.migration.success.buttons.more') }}
                </button>
              </div>
            </div>
          </div>

          <!-- State: Print QR -->
          <div v-else-if="state === 'print_qr'" class="space-y-10" v-motion-slide-visible-bottom>
            <div class="grid grid-cols-1 lg:grid-cols-3 gap-10">
              <div class="lg:col-span-2 space-y-8">
                <!-- Label Configuration -->
                <div class="glass p-10 rounded-lg space-y-10">
                  <div class="flex items-center gap-3">
                    <div class="w-10 h-10 rounded-xl bg-primary-50 dark:bg-primary-900/20 text-primary-500 flex items-center justify-center shadow-sm">
                      <LucideSettings2 class="w-5 h-5" />
                    </div>
                    <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('registration.migration.print_qr.config.title') }}</h3>
                  </div>

                  <div class="space-y-10">
                    <div class="space-y-4">
                      <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.print_qr.config.type.label') }}</label>
                      <div class="flex gap-4">
                        <button class="flex-grow flex items-center justify-center gap-3 py-4 px-6 bg-white dark:bg-slate-800 border-2 border-[#1E3A5F] rounded-2xl text-xs font-black text-[#1E3A5F] dark:text-white shadow-lg shadow-blue-900/5">
                          <LucideFileText class="w-4 h-4" />
                          {{ $t('registration.migration.print_qr.config.type.doc') }}
                        </button>
                        <button class="flex-grow flex items-center justify-center gap-3 py-4 px-6 bg-white dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-black text-slate-400 hover:bg-slate-50 transition-all">
                          <LucideLayoutGrid class="w-4 h-4" />
                          {{ $t('registration.migration.print_qr.config.type.rack') }}
                        </button>
                      </div>
                    </div>

                    <div class="grid grid-cols-2 gap-8">
                      <div class="space-y-4">
                        <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.print_qr.config.mode.label') }}</label>
                        <div class="relative">
                          <select class="w-full pl-5 pr-12 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold appearance-none outline-none focus:ring-4 focus:ring-primary-500/10 transition-all">
                            <option>{{ $t('registration.migration.print_qr.config.mode.single') }}</option>
                          </select>
                          <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
                        </div>
                      </div>
                      <div class="space-y-4">
                        <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.print_qr.config.copies') }}</label>
                        <div class="flex items-center bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-2xl overflow-hidden">
                          <button class="p-4 hover:bg-slate-100 dark:hover:bg-slate-700 transition-colors text-slate-400"><LucideMinus class="w-4 h-4" /></button>
                          <input type="text" value="1" class="w-full text-center bg-transparent text-sm font-black outline-none" />
                          <button class="p-4 hover:bg-slate-100 dark:hover:bg-slate-700 transition-colors text-slate-400"><LucidePlus class="w-4 h-4" /></button>
                        </div>
                      </div>
                    </div>

                    <div class="space-y-4">
                      <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.print_qr.config.printer.label') }}</label>
                      <div class="relative">
                        <LucidePrinter class="absolute left-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
                        <select class="w-full pl-12 pr-12 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold appearance-none outline-none">
                          <option>Zebra ZT411 - Warehouse Floor 1</option>
                        </select>
                        <LucideChevronDown class="absolute right-5 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
                      </div>
                      <p class="text-[9px] font-black text-green-500 uppercase tracking-widest flex items-center gap-1.5 px-1">
                        <span class="w-1.5 h-1.5 bg-green-500 rounded-full animate-pulse"></span>
                        {{ $t('registration.migration.print_qr.config.printer.status') }}
                      </p>
                    </div>
                  </div>
                </div>

                <!-- Data Source -->
                <div class="glass p-10 rounded-lg space-y-6">
                  <div class="flex items-center gap-3">
                    <div class="w-10 h-10 rounded-xl bg-primary-50 dark:bg-primary-900/20 text-primary-500 flex items-center justify-center shadow-sm">
                      <LucideDatabase class="w-5 h-5" />
                    </div>
                    <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('registration.migration.print_qr.source.title') }}</h3>
                  </div>
                  <div class="space-y-3">
                    <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">{{ $t('registration.migration.print_qr.source.label') }}</label>
                    <div class="flex gap-4">
                      <input type="text" :placeholder="$t('registration.migration.print_qr.source.placeholder')" class="flex-grow px-5 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-800 rounded-2xl text-xs font-bold outline-none" />
                      <button class="px-8 py-4 bg-slate-800 text-white rounded-2xl text-xs font-black uppercase tracking-widest hover:bg-slate-900 transition-all">
                        {{ $t('registration.migration.print_qr.source.btn_fetch') }}
                      </button>
                    </div>
                  </div>
                </div>
              </div>

              <!-- Label Preview Column -->
              <div class="space-y-8">
                <div class="glass p-10 rounded-lg space-y-10">
                  <div class="flex items-center justify-between">
                    <div class="flex items-center gap-3">
                      <LucideEye class="w-5 h-5 text-primary-500" />
                      <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">{{ $t('registration.migration.print_qr.preview.title') }}</h3>
                    </div>
                    <button class="text-[9px] font-black text-primary-500 uppercase tracking-widest hover:underline">{{ $t('registration.migration.print_qr.preview.settings') }}</button>
                  </div>

                  <!-- Label Render -->
                  <div class="aspect-[4/3] bg-white border border-slate-100 shadow-inner rounded-2xl p-8 flex flex-col relative overflow-hidden group">
                    <div class="flex-grow flex items-start justify-between">
                      <div class="space-y-1">
                        <p class="text-[8px] font-black text-slate-300 uppercase tracking-widest">{{ $t('registration.migration.print_qr.preview.type_label') }}</p>
                        <p class="text-xl font-black text-slate-800 tracking-tight">DOC-2023-0842</p>
                      </div>
                      <div class="w-16 h-16 bg-slate-800 rounded-lg p-2 flex items-center justify-center">
                        <LucideQrCode class="w-full h-full text-white" />
                      </div>
                    </div>
                    <div class="mt-auto flex items-end justify-between border-t border-slate-50 pt-6">
                      <div class="space-y-1">
                        <p class="text-[8px] font-black text-slate-300 uppercase tracking-widest">{{ $t('registration.migration.print_qr.preview.loc_label') }}</p>
                        <p class="text-2xl font-black text-slate-800 tracking-tighter">RACK-A1-04</p>
                      </div>
                      <div class="text-right">
                        <p class="text-[8px] font-black text-slate-300 uppercase tracking-widest">{{ $t('registration.migration.print_qr.preview.area_label') }}</p>
                        <p class="text-[10px] font-black text-slate-500 uppercase">Zone B14</p>
                      </div>
                    </div>
                    <!-- Label Size Marker -->
                    <div class="absolute bottom-2 left-1/2 -translate-x-1/2 text-[7px] font-black text-slate-300 uppercase tracking-widest opacity-0 group-hover:opacity-100 transition-opacity">
                      {{ $t('registration.migration.print_qr.preview.thermal_hint') }}
                    </div>
                  </div>

                  <div class="p-6 bg-primary-50/30 dark:bg-primary-900/10 border border-primary-100/50 dark:border-primary-800/30 rounded-2xl flex gap-4">
                    <LucideInfo class="w-4 h-4 text-primary-500 flex-shrink-0" />
                    <p class="text-[10px] font-bold text-slate-500 leading-relaxed italic">
                      {{ $t('registration.migration.print_qr.preview.info') }}
                    </p>
                  </div>

                  <button class="w-full py-5 bg-[#1E3A5F] text-white rounded-lg text-xs font-black uppercase tracking-widest shadow-2xl shadow-blue-900/40 hover:bg-[#152943] transition-all flex items-center justify-center gap-3">
                    <LucidePrinter class="w-5 h-5" />
                    {{ $t('registration.migration.print_qr.preview.btn_print') }}
                  </button>
                </div>
              </div>
            </div>

            <!-- Recent Print Jobs -->
            <div class="space-y-6">
              <div class="flex items-center justify-between">
                <h3 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('registration.migration.print_qr.history.title') }}</h3>
                <button class="text-[10px] font-black text-primary-500 uppercase tracking-widest hover:underline">{{ $t('registration.migration.print_qr.history.view_all') }}</button>
              </div>
              <div class="glass rounded-lg overflow-hidden">
                <table class="w-full text-left">
                  <thead class="text-[9px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">
                    <tr>
                      <th class="p-8 px-10">{{ $t('registration.migration.print_qr.history.cols.time') }}</th>
                      <th class="p-8">{{ $t('registration.migration.print_qr.history.cols.id') }}</th>
                      <th class="p-8">{{ $t('registration.migration.print_qr.history.cols.printer') }}</th>
                      <th class="p-8 text-center">{{ $t('registration.migration.print_qr.history.cols.status') }}</th>
                      <th class="p-8 text-right px-10">{{ $t('registration.migration.print_qr.history.cols.action') }}</th>
                    </tr>
                  </thead>
                  <tbody class="divide-y divide-slate-50 dark:divide-slate-800/50">
                    <tr v-for="job in printHistory" :key="job.id" class="hover:bg-slate-50/50 dark:hover:bg-slate-800/20 transition-all">
                      <td class="p-8 px-10 text-xs font-bold text-slate-500 uppercase">{{ job.time }}</td>
                      <td class="p-8 text-sm font-black text-[#1E3A5F] dark:text-white uppercase">{{ job.labelId }}</td>
                      <td class="p-8 text-xs font-bold text-slate-500">{{ job.printer }}</td>
                      <td class="p-8 text-center">
                        <span class="px-3 py-1 bg-green-50 text-green-500 rounded-md text-[8px] font-black uppercase tracking-widest border border-green-100">SUCCESS</span>
                      </td>
                      <td class="p-8 text-right px-10">
                        <button class="text-[9px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest hover:underline">{{ $t('registration.migration.print_qr.history.actions.reprint') }}</button>
                      </td>
                    </tr>
                  </tbody>
                </table>
                <div class="p-6 text-center border-t border-slate-100 dark:border-slate-800">
                  <button class="text-[9px] font-black text-slate-400 uppercase tracking-widest hover:text-slate-600 transition-colors">
                    {{ $t('registration.migration.print_qr.history.view_logs') }}
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Fixed Floating Buttons -->
    <div v-if="state !== 'success'" class="fixed bottom-10 right-10 flex flex-col gap-4 z-20">
      <button class="w-14 h-14 bg-white dark:bg-slate-900 text-[#1E3A5F] dark:text-white rounded-full shadow-2xl flex items-center justify-center hover:scale-110 active:scale-95 transition-all ring-1 ring-slate-200 dark:ring-slate-800">
        <LucideHelpCircle class="w-6 h-6" />
      </button>
      <button class="w-14 h-14 bg-[#1E3A5F] text-white rounded-full shadow-2xl flex items-center justify-center hover:scale-110 active:scale-95 transition-all shadow-blue-900/40">
        <LucideSave class="w-6 h-6" />
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  LucideCheckCircle2, LucideMaximize2, LucideSparkles, LucideMenu, 
  LucideMinus, LucidePlus, LucidePrinter, LucideDownload, LucideExpand,
  LucideChevronDown, LucideCalendar, LucidePenTool, LucideAlertTriangle,
  LucideRefreshCcw, LucidePencil, LucideUser, LucideBuilding2, LucideMapPin,
  LucideLightbulb, LucideHelpCircle, LucideSave, LucideTag, LucideCalendarCheck,
  LucideEye, LucideInfo, LucideCopy, LucideMap, LucideQrCode, LucideSearch,
  LucideBell, LucideSettings2, LucideFileText, LucideLayoutGrid, LucideDatabase
} from 'lucide-vue-next'

const state = ref('initial') // 'initial', 'review', 'mandatory', 'duplicate_check', 'location_assignment', 'success', 'print_qr'

definePageMeta({
  layout: false
})

const similarRecords = [
  { id: 1, score: 92, filename: 'LGC-2019-PRJ-0042.pdf', date: '12 Oct 2019', path: 'Procurement / 2019 / Vendor Contracts' },
  { id: 2, score: 78, filename: 'PRJ-CONTRACT-V4.docx', date: '05 Jan 2020', path: 'Projects / Heritage / Archive_Bin' }
]

const locations = [
  { id: 1, code: 'R- 04/B- 022', level: 'Level 2, Aisle B', percent: 92, type: 'PDF-OCR', isFull: false },
  { id: 2, code: 'R- 04/B- 022', level: 'Level 2, Aisle B', percent: 45, type: 'TIF-RAW', isFull: false },
  { id: 3, code: 'R- 04/B- 022', level: 'Level 3, Aisle B', percent: 100, type: 'OVERFLOW', isFull: true }
]

const printHistory = [
  { id: 1, time: 'Oct 24, 14:20', labelId: 'DOC-2023-0841', printer: 'Zebra ZT411', status: 'SUCCESS' },
  { id: 2, time: 'Oct 24, 13:55', labelId: 'RACK-B2-01', printer: 'Zebra ZT411', status: 'SUCCESS' }
]
</script>

<style scoped>
@reference "../../assets/css/main.css";

.custom-scrollbar::-webkit-scrollbar { width: 6px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #E2E8F0; border-radius: 10px; }
.dark .custom-scrollbar::-webkit-scrollbar-thumb { background: #1E293B; }

.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}

h1, h2, h3, h4 {
  text-shadow: 0 2px 4px rgba(30, 58, 95, 0.05);
}
</style>
