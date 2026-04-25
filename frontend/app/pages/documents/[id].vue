<template>
  <div class="max-w-7xl mx-auto space-y-8 pb-20" :class="{ 'opacity-0 pointer-events-none': isEnlarged }">
    <!-- Top Navigation & Header -->
    <div class="flex items-center justify-between" v-motion-fade>
      <div class="space-y-4">
        <button 
          @click="navigateTo('/documents')" 
          class="flex items-center gap-2 text-xs font-black text-slate-400 hover:text-primary-600 transition-colors uppercase tracking-widest group"
        >
          <LucideArrowLeft class="w-4 h-4 group-hover:-translate-x-1 transition-transform" />
          {{ $t('documents.detail.btn_back') }}
        </button>
        <div class="space-y-1">
          <h1 class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">KONTRAK VENDOR PT XYZ</h1>
          <p class="text-sm font-bold text-slate-400 uppercase tracking-widest">LGL-AKR-001-2026-005.pdf</p>
        </div>
      </div>
    </div>

    <!-- Standard PDF Viewer -->
    <div class="glass rounded-[2.5rem] overflow-hidden shadow-2xl shadow-slate-200/50 dark:shadow-none border border-white" v-motion-fade>
      <div class="bg-[#525659] p-2 flex items-center justify-between text-white border-b border-white/10">
        <div class="flex items-center gap-4">
          <button class="p-2 hover:bg-white/10 rounded-lg transition-colors"><LucideMenu class="w-5 h-5" /></button>
          <div class="flex items-center gap-2">
            <span class="text-xs font-bold opacity-80 uppercase tracking-widest">{{ $t('documents.detail.viewer.page') }}</span>
            <input type="text" value="1" class="w-10 bg-[#323639] border-none rounded px-2 py-1 text-xs font-bold text-center focus:ring-1 focus:ring-blue-400" />
            <span class="text-xs font-bold opacity-80">/ 12</span>
          </div>
        </div>
        <div class="flex items-center gap-6">
          <div class="flex items-center gap-4 bg-[#323639] px-4 py-1.5 rounded-full border border-white/5 shadow-inner">
            <button class="hover:text-blue-400"><LucideMinus class="w-4 h-4" /></button>
            <span class="text-xs font-bold w-12 text-center tracking-tighter">100%</span>
            <button class="hover:text-blue-400"><LucidePlus class="w-4 h-4" /></button>
          </div>
          <div class="flex items-center gap-4 border-l border-white/10 pl-6">
            <button class="hover:text-blue-400 transition-colors"><LucidePrinter class="w-5 h-5" /></button>
            <button class="hover:text-blue-400 transition-colors"><LucideDownload class="w-5 h-5" /></button>
            <button @click="isEnlarged = true" class="hover:text-blue-400 transition-colors"><LucideMaximize2 class="w-5 h-5" /></button>
          </div>
        </div>
      </div>
      
      <!-- Mock PDF Content -->
      <div class="bg-[#8E9194] p-10 min-h-[600px] flex justify-center">
        <div class="bg-white w-full max-w-3xl shadow-2xl p-16 space-y-12 relative overflow-hidden">
          <div class="flex justify-between items-start">
            <div class="w-32 h-10 bg-slate-100 rounded-lg animate-pulse"></div>
            <div class="text-right">
              <div class="w-32 h-4 bg-slate-900 ml-auto mb-1"></div>
              <p class="text-[10px] text-slate-400 font-bold uppercase tracking-widest">Jakarta, 15 Februari 2026</p>
            </div>
          </div>
          <div class="space-y-4 pt-10">
            <p class="text-[10px] text-slate-400 font-black uppercase tracking-[0.3em]">CONFIDENTIAL DOCUMENT</p>
            <h2 class="text-3xl font-black text-[#1E3A5F] border-b-2 border-[#1E3A5F]/20 pb-4 tracking-tight">SURAT PERJANJIAN KERJASAMA</h2>
          </div>
          <div class="space-y-6">
            <p class="text-sm font-medium text-slate-600 leading-relaxed italic">Pada hari ini, Senin tanggal lima belas bulan Februari tahun dua ribu dua puluh enam, yang bertanda tangan di bawah ini:</p>
            <div class="grid grid-cols-[100px_1fr] gap-x-4 gap-y-4 text-sm">
              <span class="font-bold text-slate-400 uppercase tracking-widest text-[10px] pt-1">Nama</span>
              <div class="h-4 bg-slate-900 w-48 rounded-sm"></div>
              <span class="font-bold text-slate-400 uppercase tracking-widest text-[10px] pt-1">Jabatan</span>
              <div class="h-4 bg-slate-900 w-32 rounded-sm"></div>
              <span class="font-bold text-slate-400 uppercase tracking-widest text-[10px] pt-1">Alamat</span>
              <div class="h-4 bg-slate-900 w-80 rounded-sm"></div>
            </div>
          </div>
          <div class="absolute bottom-10 right-10 opacity-10 rotate-[-15deg] pointer-events-none">
            <p class="text-5xl font-black uppercase tracking-[0.5em] text-slate-900">14:33:27 WIB</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Actions Bar -->
    <div class="flex flex-wrap items-center justify-between gap-6 py-4" v-motion-fade>
      <div class="flex items-center gap-6">
        <div class="space-y-1">
          <h3 class="text-xl font-black text-[#1E3A5F] tracking-tight uppercase">KONTRAK VENDOR PT XYZ</h3>
          <p class="text-xs font-bold text-slate-400 uppercase tracking-widest">LGL-AKR-001-2026-005</p>
        </div>
        <span class="px-4 py-1.5 bg-green-50 text-green-600 rounded-full text-[10px] font-black uppercase tracking-widest border border-green-100 flex items-center gap-2">
          <div class="w-1.5 h-1.5 rounded-full bg-green-500"></div>
          {{ $t('documents.detail.actions.available') }}
        </span>
      </div>

      <div class="flex items-center gap-3">
        <button class="flex items-center gap-3 px-8 py-3.5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 transition-all">
          <LucideShoppingCart class="w-4 h-4" />
          {{ $t('documents.detail.actions.add_to_cart') }}
        </button>
        <button class="flex items-center gap-3 px-8 py-3.5 bg-white border border-slate-200 rounded-xl text-xs font-black uppercase tracking-widest text-slate-600 hover:bg-slate-50 transition-all">
          <LucideFileDown class="w-4 h-4" />
          {{ $t('documents.detail.actions.request_digital') }}
        </button>
        <button class="p-3.5 bg-white border border-slate-200 rounded-xl text-slate-400 hover:text-slate-600 transition-all">
          <LucideMoreHorizontal class="w-5 h-5" />
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10 items-start">
      <div class="lg:col-span-7 space-y-10">
        <div class="flex flex-wrap gap-2" v-motion-fade>
          <span v-for="tag in ['VENDOR', 'CLEANING', 'KONTRAK', 'KERJA', 'GEDUNG']" :key="tag" class="px-4 py-1.5 bg-white border border-slate-100 text-slate-400 text-[10px] font-black rounded-lg uppercase tracking-widest hover:border-primary-500 hover:text-primary-600 transition-colors cursor-pointer">#{{ tag }}</span>
        </div>
        <div class="space-y-6" v-motion-fade>
          <div class="flex items-center gap-3">
            <div class="w-1.5 h-6 bg-primary-500 rounded-full"></div>
            <h2 class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">{{ $t('documents.detail.metadata.title') }}</h2>
          </div>
          <div class="glass rounded-[2rem] overflow-hidden border border-slate-100">
            <table class="w-full text-sm">
              <tbody class="divide-y divide-slate-50">
                <tr v-for="(val, label) in metadata" :key="label" class="group">
                  <td class="px-8 py-4 w-1/3 bg-slate-50/50 text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ label }}</td>
                  <td class="px-8 py-4 font-bold text-[#1E3A5F]">{{ val }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
        <div class="space-y-6" v-motion-fade>
          <div class="flex items-center gap-3">
            <div class="w-1.5 h-6 bg-primary-500 rounded-full"></div>
            <h2 class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">{{ $t('documents.detail.location.title') }}</h2>
          </div>
          <div class="bg-slate-50/50 border border-slate-100 rounded-[2rem] p-8 flex items-center justify-between relative overflow-hidden group">
            <div class="space-y-6 relative z-10">
              <div class="grid grid-cols-2 gap-x-12 gap-y-6">
                <div class="flex items-start gap-4">
                  <LucideHome class="w-5 h-5 text-slate-400 mt-0.5" />
                  <div class="space-y-1">
                    <p class="text-xs font-bold text-slate-700">Gedung A / Ruang 2</p>
                    <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('documents.detail.location.building_room') }}</p>
                  </div>
                </div>
                <div class="flex items-start gap-4">
                  <LucideLayers class="w-5 h-5 text-slate-400 mt-0.5" />
                  <div class="space-y-1">
                    <p class="text-xs font-bold text-slate-700">Lemari L3 / Rak R2</p>
                    <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ $t('documents.detail.location.cabinet_shelf') }}</p>
                  </div>
                </div>
              </div>
              <div class="flex items-center gap-4 bg-white/60 p-4 rounded-2xl border border-white/80 shadow-sm w-fit">
                <LucideBox class="w-5 h-5 text-primary-600" />
                <p class="text-sm font-black text-primary-600 uppercase tracking-tight">Box B5 - Folder 04</p>
              </div>
            </div>
            <div class="relative w-24 h-24 bg-white rounded-2xl shadow-xl flex items-center justify-center p-2 group-hover:scale-105 transition-transform">
              <LucideQrCode class="w-full h-full text-slate-300" />
              <div class="absolute bottom-1 right-1 bg-primary-500 text-white p-1 rounded-md"><LucideLocateFixed class="w-3 h-3" /></div>
              <span class="absolute -top-3 -right-3 bg-white px-2 py-1 rounded-md text-[8px] font-black text-slate-400 shadow-sm border border-slate-50 uppercase tracking-widest">{{ $t('documents.detail.location.scan') }}</span>
            </div>
            <LucideMapPin class="absolute top-4 right-4 w-5 h-5 text-slate-200 group-hover:text-primary-200 transition-colors" />
          </div>
        </div>
        <div class="space-y-6" v-motion-fade>
          <div class="flex items-center gap-3">
            <div class="w-1.5 h-6 bg-primary-500 rounded-full"></div>
            <h2 class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">{{ $t('documents.detail.related.title') }}</h2>
          </div>
          <div class="space-y-4">
            <div v-for="rel in relatedDocs" :key="rel.name" class="group flex items-center justify-between p-5 bg-white border border-slate-100 rounded-3xl hover:border-primary-200 hover:shadow-lg hover:shadow-primary-500/5 transition-all cursor-pointer">
              <div class="flex items-center gap-5">
                <div :class="`w-12 h-12 rounded-2xl flex items-center justify-center ${rel.bg}`"><LucideFileText :class="`w-6 h-6 ${rel.color}`" /></div>
                <div class="space-y-1">
                  <p class="text-sm font-black text-[#1E3A5F] group-hover:text-primary-600 transition-colors uppercase tracking-tight">{{ rel.name }}</p>
                  <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ rel.desc }}</p>
                </div>
              </div>
              <LucideChevronRight class="w-5 h-5 text-slate-300 group-hover:translate-x-1 transition-transform" />
            </div>
          </div>
        </div>
      </div>
      <div class="lg:col-span-5 space-y-8" v-motion-fade>
        <div class="flex items-center gap-3">
          <div class="w-1.5 h-6 bg-primary-500 rounded-full"></div>
          <h2 class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">{{ $t('documents.detail.loan_history.title') }}</h2>
        </div>
        <div class="relative pl-12 space-y-12">
          <div class="absolute left-6 top-4 bottom-4 w-0.5 bg-slate-100"></div>
          <div v-for="(loan, index) in timeline" :key="index" class="relative group">
            <div class="absolute -left-6 top-1.5 w-6 h-6 rounded-full bg-white border-4 border-slate-100 group-hover:border-primary-200 transition-colors z-10"></div>
            <div class="space-y-4">
              <div class="flex items-center gap-4">
                <div class="w-12 h-12 rounded-2xl bg-slate-100 border-2 border-white shadow-sm overflow-hidden flex-shrink-0"><img :src="`https://i.pravatar.cc/150?u=${loan.name}`" class="w-full h-full object-cover" /></div>
                <div>
                  <p class="text-sm font-black text-[#1E3A5F]">{{ loan.name }}</p>
                  <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">Dept. {{ loan.dept }}</p>
                </div>
              </div>
              <div class="bg-slate-50/50 p-4 rounded-2xl border border-slate-100 group-hover:bg-primary-50/30 group-hover:border-primary-100 transition-all">
                <p class="text-[10px] font-bold text-slate-500 leading-relaxed uppercase tracking-widest" v-html="$t('documents.detail.loan_history.returned', { date: loan.date }) + ' • ' + $t('documents.detail.loan_history.duration', { days: loan.duration })"></p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- ENLARGED MODE OVERLAY -->
    <Transition
      enter-active-class="transition duration-300 ease-out"
      enter-from-class="opacity-0 scale-95"
      enter-to-class="opacity-100 scale-100"
      leave-active-class="transition duration-200 ease-in"
      leave-from-class="opacity-100 scale-100"
      leave-to-class="opacity-0 scale-95"
    >
      <div v-if="isEnlarged" class="fixed inset-0 z-[100] bg-white flex flex-col overflow-hidden">
        <!-- PDF Viewer Header (Persistent) -->
        <div class="bg-[#323639] p-2 flex items-center justify-between text-white border-b border-white/10 shrink-0">
          <div class="flex items-center gap-4">
            <button class="p-2 hover:bg-white/10 rounded-lg transition-colors"><LucideMenu class="w-5 h-5" /></button>
            <div class="flex items-center gap-2">
              <span class="text-xs font-bold opacity-80 uppercase tracking-widest">{{ $t('documents.detail.viewer.page') }}</span>
              <input type="text" value="1" class="w-10 bg-[#323639] border-none rounded px-2 py-1 text-xs font-bold text-center focus:ring-1 focus:ring-blue-400" />
              <span class="text-xs font-bold opacity-80">/ 12</span>
            </div>
          </div>
          <div class="flex items-center gap-6">
            <div class="flex items-center gap-4 bg-[#525659] px-4 py-1.5 rounded-full border border-white/5 shadow-inner">
              <button class="hover:text-blue-400"><LucideMinus class="w-4 h-4" /></button>
              <span class="text-xs font-bold w-12 text-center tracking-tighter">100%</span>
              <button class="hover:text-blue-400"><LucidePlus class="w-4 h-4" /></button>
            </div>
            <div class="flex items-center gap-4 border-l border-white/10 pl-6">
              <button class="hover:text-blue-400 transition-colors"><LucidePrinter class="w-5 h-5" /></button>
              <button class="hover:text-blue-400 transition-colors"><LucideDownload class="w-5 h-5" /></button>
              <button @click="isEnlarged = false" class="hover:text-blue-400 transition-colors text-blue-400"><LucideMinimize2 class="w-5 h-5" /></button>
            </div>
          </div>
        </div>

        <div class="flex-1 flex overflow-hidden">
          <!-- Left: PDF Content -->
          <div class="flex-1 bg-[#525659] overflow-y-auto p-10 flex justify-center custom-scrollbar">
            <div class="bg-white w-full max-w-4xl shadow-2xl p-20 space-y-16 relative overflow-hidden h-fit">
              <div class="flex justify-between items-start">
                <div class="w-40 h-12 bg-slate-100 rounded-lg animate-pulse"></div>
                <div class="text-right">
                  <div class="w-40 h-5 bg-slate-900 ml-auto mb-2"></div>
                  <p class="text-xs text-slate-400 font-bold uppercase tracking-widest">Jakarta, 15 Februari 2026</p>
                </div>
              </div>

              <div class="space-y-6 pt-10">
                <p class="text-xs text-slate-400 font-black uppercase tracking-[0.4em]">CONFIDENTIAL DOCUMENT</p>
                <h2 class="text-4xl font-black text-[#1E3A5F] border-b-4 border-[#1E3A5F]/20 pb-6 tracking-tight">SURAT PERJANJIAN KERJASAMA</h2>
              </div>

              <div class="space-y-8">
                <p class="text-base font-medium text-slate-600 leading-relaxed italic">
                  Pada hari ini, Senin tanggal lima belas bulan Februari tahun dua ribu dua puluh enam, yang bertanda tangan di bawah ini:
                </p>

                <div class="grid grid-cols-[120px_1fr] gap-x-6 gap-y-6 text-base">
                  <span class="font-bold text-slate-400 uppercase tracking-widest text-xs pt-1">Nama</span>
                  <div class="h-5 bg-slate-900 w-64 rounded-sm"></div>
                  
                  <span class="font-bold text-slate-400 uppercase tracking-widest text-xs pt-1">Jabatan</span>
                  <div class="h-5 bg-slate-900 w-48 rounded-sm"></div>
                  
                  <span class="font-bold text-slate-400 uppercase tracking-widest text-xs pt-1">Alamat</span>
                  <div class="h-5 bg-slate-900 w-full max-w-lg rounded-sm"></div>
                </div>

                <p class="text-sm font-bold text-slate-800 leading-relaxed pt-4">
                  Dalam hal ini bertindak untuk dan atas nama PT XYZ INDONESIA, selanjutnya disebut sebagai <span class="text-primary-600">PIHAK PERTAMA</span>.
                </p>

                <div class="w-full h-0.5 bg-slate-100 my-8"></div>

                <div class="space-y-4">
                  <p class="text-sm font-black text-slate-900 uppercase tracking-widest">PASAL 1 - LINGKUP PEKERJAAN</p>
                  <p class="text-sm text-slate-600 leading-relaxed">
                    Pekerjaan yang dimaksud dalam kontrak ini adalah pengadaaan jasa pemeliharaan gedung (cleaning services) untuk periode tahun anggaran 2026.
                  </p>
                </div>

                <div class="p-8 bg-slate-50 rounded-2xl border border-slate-100 italic text-slate-500 text-sm leading-loose">
                  "The parties hereby agree to maintain strict confidentiality regarding the terms of this agreement and any trade secrets shared during the duration of the contract..."
                </div>
              </div>

              <div class="pt-20 grid grid-cols-2 gap-20">
                <div class="text-center space-y-20">
                  <div class="h-px w-32 bg-slate-200 mx-auto"></div>
                  <p class="text-[10px] font-black uppercase tracking-widest text-slate-400">PIHAK PERTAMA</p>
                </div>
                <div class="text-center space-y-20">
                  <div class="h-px w-32 bg-slate-200 mx-auto"></div>
                  <p class="text-[10px] font-black uppercase tracking-widest text-slate-400">PIHAK KEDUA</p>
                </div>
              </div>

              <!-- Watermark -->
              <div class="absolute inset-0 flex items-center justify-center pointer-events-none opacity-[0.03] -rotate-12">
                <p class="text-[120px] font-black uppercase tracking-[0.2em]">BUDI QARTONO - NIK: 2026001 - 2026-03-20 14:33:27 WIB</p>
              </div>
            </div>
          </div>

          <!-- Right: Info Panel -->
          <div class="w-[450px] border-l border-slate-200 flex flex-col bg-white overflow-hidden shadow-2xl">
            <div class="flex-1 overflow-y-auto p-8 space-y-10 custom-scrollbar">
              <!-- Header Info -->
              <div class="flex items-start justify-between">
                <span class="px-3 py-1 bg-green-50 text-green-600 rounded-full text-[10px] font-black uppercase tracking-widest border border-green-100 flex items-center gap-1.5">
                  <div class="w-1.5 h-1.5 rounded-full bg-green-500"></div>
                  {{ $t('documents.detail.actions.available') }}
                </span>
                <button class="text-slate-300 hover:text-slate-600"><LucideMoreHorizontal class="w-5 h-5" /></button>
              </div>

              <div class="space-y-1">
                <h3 class="text-xl font-black text-[#1E3A5F] tracking-tight uppercase leading-tight">KONTRAK VENDOR PT XYZ</h3>
                <p class="text-xs font-bold text-slate-400 uppercase tracking-widest">LGL-AKR-001-2026-005</p>
              </div>

              <!-- Tags -->
              <div class="flex flex-wrap gap-1.5">
                <span v-for="tag in ['VENDOR', 'CLEANING', 'KONTRAK', 'KERJA', 'GEDUNG']" :key="tag" class="px-3 py-1 bg-slate-50 border border-slate-100 text-slate-400 text-[9px] font-black rounded-lg uppercase tracking-widest">#{{ tag }}</span>
              </div>

              <!-- Metadata -->
              <div class="space-y-4">
                <h4 class="text-[10px] font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('documents.detail.metadata.title') }}</h4>
                <div class="space-y-3">
                  <div v-for="(val, label) in metadata" :key="label" class="flex justify-between items-center text-xs">
                    <span class="font-bold text-slate-400">{{ label }}</span>
                    <span class="font-black text-[#1E3A5F]">{{ val }}</span>
                  </div>
                </div>
              </div>

              <!-- Location -->
              <div class="space-y-4">
                <h4 class="text-[10px] font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('documents.detail.location.title') }}</h4>
                <div class="bg-slate-50/50 p-5 rounded-2xl border border-slate-100 flex items-center justify-between">
                  <div class="space-y-3">
                    <div class="flex items-center gap-2 text-[10px] font-bold text-slate-600">
                      <LucideHome class="w-3.5 h-3.5 text-slate-400" />
                      Gedung A / Ruang 2
                    </div>
                    <div class="flex items-center gap-2 text-[10px] font-bold text-slate-600">
                      <LucideLayers class="w-3.5 h-3.5 text-slate-400" />
                      Lemari L3 / Rak R2
                    </div>
                    <div class="flex items-center gap-2 text-[10px] font-black text-primary-600 uppercase">
                      <LucideBox class="w-3.5 h-3.5" />
                      Box B5 - Folder 04
                    </div>
                  </div>
                  <LucideQrCode class="w-12 h-12 text-slate-200" />
                </div>
              </div>

              <!-- Related -->
              <div class="space-y-4">
                <h4 class="text-[10px] font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('documents.detail.related.title') }}</h4>
                <div class="space-y-2">
                  <div v-for="rel in relatedDocs" :key="rel.name" class="flex items-center justify-between p-3 bg-white border border-slate-100 rounded-2xl hover:border-primary-200 transition-all cursor-pointer group">
                    <div class="flex items-center gap-3">
                      <div :class="`w-8 h-8 rounded-xl flex items-center justify-center ${rel.bg}`"><LucideFileText :class="`w-4 h-4 ${rel.color}`" /></div>
                      <div>
                        <p class="text-[10px] font-black text-[#1E3A5F] group-hover:text-primary-600 uppercase tracking-tight">{{ rel.name }}</p>
                        <p class="text-[8px] font-bold text-slate-400 uppercase tracking-widest">{{ rel.desc }}</p>
                      </div>
                    </div>
                    <LucideChevronRight class="w-3 h-3 text-slate-300" />
                  </div>
                </div>
              </div>

              <!-- Loan History -->
              <div class="space-y-4">
                <h4 class="text-[10px] font-black text-slate-300 uppercase tracking-[0.2em]">{{ $t('documents.detail.loan_history.title') }}</h4>
                <div class="space-y-6 relative pl-6">
                  <div class="absolute left-[7px] top-2 bottom-2 w-0.5 bg-slate-100"></div>
                  <div v-for="loan in timeline.slice(0, 2)" :key="loan.name" class="relative">
                    <div class="absolute -left-6 top-1 w-3.5 h-3.5 rounded-full bg-white border-2 border-slate-200 z-10"></div>
                    <p class="text-[10px] font-black text-[#1E3A5F]">{{ loan.name }} <span class="text-slate-400 font-bold ml-1">(Dept. {{ loan.dept }})</span></p>
                    <p class="text-[8px] font-bold text-slate-400 uppercase tracking-widest mt-1" v-html="$t('documents.detail.loan_history.returned', { date: loan.date }) + ' • ' + $t('documents.detail.loan_history.duration', { days: loan.duration })"></p>
                  </div>
                </div>
              </div>
            </div>

            <!-- Footer Actions -->
            <div class="p-8 bg-slate-50 border-t border-slate-100 space-y-3 shrink-0">
              <button class="w-full flex items-center justify-center gap-3 py-4 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-[10px] font-black uppercase tracking-[0.15em] shadow-xl shadow-blue-900/20 transition-all group">
                <LucideShoppingCart class="w-4 h-4 group-hover:scale-110 transition-transform" />
                {{ $t('documents.detail.actions.add_to_cart') }}
              </button>
              <button class="w-full flex items-center justify-center gap-3 py-4 bg-white border border-slate-200 text-slate-600 rounded-2xl text-[10px] font-black uppercase tracking-[0.15em] hover:bg-slate-50 transition-all">
                <LucideFileDown class="w-4 h-4" />
                {{ $t('documents.detail.actions.request_digital') }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
const { t } = useI18n()
import { ref } from 'vue'
import { 
  LucideArrowLeft, LucideMenu, LucideMinus, LucidePlus, LucidePrinter, LucideDownload, LucideMaximize2, LucideMinimize2,
  LucideShoppingCart, LucideFileDown, LucideMoreHorizontal, LucideHome, LucideLayers, LucideBox,
  LucideQrCode, LucideLocateFixed, LucideMapPin, LucideChevronRight, LucideFileText
} from 'lucide-vue-next'

const isEnlarged = ref(false)

const metadata = computed(() => ({
  [t('documents.detail.metadata.labels.doc_type')]: 'Kontrak',
  [t('documents.detail.metadata.labels.company')]: 'PT AMS',
  [t('documents.detail.metadata.labels.department')]: 'Legal',
  [t('documents.detail.metadata.labels.year')]: '2026',
  [t('documents.detail.metadata.labels.date_issued')]: '15 Feb 2026',
  [t('documents.detail.metadata.labels.retention')]: '7 Years (Permanent)',
  [t('documents.detail.metadata.labels.version')]: 'v2.1 Final'
}))

const relatedDocs = [
  { name: 'SLA MAINTENANCE 2026', desc: 'Attached - Linked', color: 'text-orange-500', bg: 'bg-orange-50' },
  { name: 'QUOTATION VENDOR XYZ', desc: 'Reference - Digital Only', color: 'text-blue-500', bg: 'bg-blue-50' }
]

const timeline = [
  { name: 'Robby Wijaya', dept: 'General', date: '02 Jan 2026', duration: 5 },
  { name: 'Ani Suryani', dept: 'Finance', date: '15 Nov 2025', duration: 2 },
  { name: 'Budi Pratama', dept: 'Legal', date: '18 Aug 2025', duration: 1 }
]
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
</style>
