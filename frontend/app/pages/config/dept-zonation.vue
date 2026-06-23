<template>
  <div v-motion-fade class="space-y-8 pb-20 p-6">
    <!-- Page Header -->
    <div class="space-y-1">
      <h1 id="page-title" class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
        Departemen Zoning Control Center
      </h1>
      <p class="text-xs font-bold text-slate-500 max-w-2xl uppercase tracking-tighter">
        Atur dan pantau zonasi penempatan arsip fisik berdasarkan departemen menggunakan peta layout interaktif 8x8.
      </p>
    </div>

    <!-- Main Grid Dashboard -->
    <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
      
      <!-- LEFT PANEL: Departments List & Capacities -->
      <div class="lg:col-span-1 flex flex-col space-y-4 bg-white dark:bg-slate-900 border border-slate-200/60 dark:border-slate-800/40 rounded-3xl p-6 shadow-sm">
        <div class="flex items-center justify-between">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Daftar Departemen</h3>
          <span id="dept-count-badge" class="px-2 py-0.5 bg-[#1E3A5F]/10 text-[#1E3A5F] dark:bg-blue-900/30 dark:text-blue-300 rounded text-[8px] font-black uppercase">
            {{ departments.length }} Depts
          </span>
        </div>

        <div class="flex-1 space-y-3 overflow-y-auto pr-1 custom-scrollbar min-h-[400px] max-h-[580px]">
          <div 
            v-for="dept in departmentsWithCounts" 
            :key="dept.id"
            :class="[
              'p-4 rounded-2xl border cursor-pointer transition-all flex flex-col gap-3',
              selectedDeptId === dept.id 
                ? 'bg-white dark:bg-slate-800 border-primary-500 shadow-md ring-1 ring-primary-500' 
                : 'bg-slate-50/50 dark:bg-slate-900/50 border-slate-200/60 dark:border-slate-800/40 hover:border-primary-300'
            ]"
            @click="selectDepartment(dept)"
          >
            <!-- Header of card -->
            <div class="flex justify-between items-start">
              <div class="min-w-0 flex-1 pr-2">
                <h4 class="text-xs font-black text-[#1E3A5F] dark:text-white uppercase truncate">{{ dept.name }}</h4>
                <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest mt-0.5">
                  ID: {{ dept.id.substring(0, 8).toUpperCase() }}
                </p>
              </div>
              <span
:class="[
                'px-1.5 py-0.5 rounded text-[8px] font-black uppercase tracking-wider shrink-0',
                dept.rackCount > 0 
                  ? 'bg-emerald-500/10 text-emerald-500 dark:bg-emerald-950/30' 
                  : 'bg-slate-200 text-slate-500 dark:bg-slate-800'
              ]">
                {{ dept.rackCount > 0 ? 'Active' : 'Standard' }}
              </span>
            </div>

            <!-- Details -->
            <div class="flex items-center justify-between text-[10px] pt-2 border-t border-slate-100 dark:border-slate-800/60">
              <span class="font-bold text-slate-500">Zoned Racks:</span>
              <span class="font-black text-[#1E3A5F] dark:text-white">{{ dept.rackCount }} / 64</span>
            </div>

            <!-- Color Bar indicator -->
            <div class="h-1 w-full rounded-full overflow-hidden bg-slate-100 dark:bg-slate-800">
              <div 
                class="h-full transition-all duration-500" 
                :style="{ 
                  width: `${(dept.rackCount / 64) * 100}%`,
                  backgroundColor: getDeptColorHex(dept.name)
                }"
              />
            </div>
          </div>
        </div>

        <!-- [+ NEW DEPARTMENT] Button at the bottom -->
        <button 
          id="new-dept-btn"
          class="w-full py-4 bg-slate-50 dark:bg-slate-850 hover:bg-slate-100 dark:hover:bg-slate-800 border border-slate-200 dark:border-slate-800 text-[#1E3A5F] dark:text-white rounded-xl text-[10px] font-black uppercase tracking-widest transition-all flex items-center justify-center gap-2 cursor-pointer"
          @click="openNewDeptModal"
        >
          <LucidePlus class="w-4.5 h-4.5" />
          New Department
        </button>
      </div>

      <!-- MIDDLE PANEL: Interactive 8x8 Layout Map -->
      <div class="lg:col-span-2 space-y-4">
        <!-- Toolbar controls -->
        <div class="flex flex-wrap items-center justify-between gap-4 bg-white dark:bg-slate-900 p-4 border border-slate-200/60 dark:border-slate-800/40 rounded-2xl shadow-sm">
          <div class="flex items-center gap-3">
            <label for="floor-selector" class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Layout Area:</label>
            <select 
              id="floor-selector"
              v-model="activeFloorId" 
              class="bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl px-3 py-1.5 text-xs font-black text-[#1E3A5F] dark:text-white outline-none focus:ring-2 focus:ring-primary-500/20 cursor-pointer"
            >
              <option v-for="fl in floorsList" :key="fl.id" :value="fl.id">{{ fl.name }} ({{ fl.code }})</option>
              <option value="__unassigned__">Belum Ditentukan</option>
            </select>
          </div>

          <!-- Zoom tool -->
          <div class="flex items-center gap-1.5 bg-slate-50 dark:bg-slate-800/50 p-1 rounded-xl">
            <button 
              id="zoom-out-btn"
              class="w-7 h-7 flex items-center justify-center rounded-lg hover:bg-white dark:hover:bg-slate-700 text-slate-500 hover:text-[#1E3A5F] dark:hover:text-white transition-colors cursor-pointer" 
              title="Zoom Out"
              @click="zoomOut"
            >
              <LucideMinus class="w-4 h-4" />
            </button>
            <span class="text-[9px] font-black text-slate-400 px-2 select-none">{{ Math.round(zoom * 100) }}%</span>
            <button 
              id="zoom-in-btn"
              class="w-7 h-7 flex items-center justify-center rounded-lg hover:bg-white dark:hover:bg-slate-700 text-slate-500 hover:text-[#1E3A5F] dark:hover:text-white transition-colors cursor-pointer" 
              title="Zoom In"
              @click="zoomIn"
            >
              <LucidePlus class="w-4 h-4" />
            </button>
            <div class="h-4 w-px bg-slate-200 dark:bg-slate-700 mx-1"/>
            <button 
              id="zoom-reset-btn"
              class="w-7 h-7 flex items-center justify-center rounded-lg hover:bg-white dark:hover:bg-slate-700 text-[9px] font-black text-slate-500 hover:text-[#1E3A5F] dark:hover:text-white transition-colors cursor-pointer" 
              title="Reset Zoom"
              @click="zoomReset"
            >
              RESET
            </button>
          </div>
        </div>

        <!-- Dynamic Legends from Database Departments -->
        <div class="bg-white dark:bg-slate-900 border border-slate-200/60 dark:border-slate-800/40 rounded-2xl p-4 shadow-sm flex flex-wrap gap-4 items-center justify-center">
          <div v-for="dept in departments" :key="dept.id" class="flex items-center gap-2">
            <span 
              class="w-3.5 h-3.5 rounded border inline-block"
              :style="{ 
                backgroundColor: getDeptColorHex(dept.name) + '20',
                borderColor: getDeptColorHex(dept.name) + '80'
              }"
            />
            <span class="text-[9px] font-black text-slate-500 uppercase tracking-wider">{{ dept.name }}</span>
          </div>
          <div class="flex items-center gap-2">
            <div class="w-3.5 h-3.5 rounded bg-slate-200 dark:bg-slate-800 border border-slate-300 dark:border-slate-700 flex items-center justify-center text-[7px] text-slate-500">
              <LucideLock class="w-2.5 h-2.5" />
            </div>
            <span class="text-[9px] font-black text-slate-500 uppercase tracking-wider">LOCKED (GREY)</span>
          </div>
        </div>

        <!-- 8x8 Layout Map Grid Container (with bottom capacity bar) -->
        <div class="bg-white dark:bg-slate-900 border border-slate-200/60 dark:border-slate-800/40 rounded-3xl shadow-sm overflow-hidden flex flex-col min-h-[550px]">
          <!-- Grid Area -->
          <div class="flex-grow p-6 flex items-center justify-center relative">
            <!-- Selection info overlay -->
            <div v-if="selectedRacks.length > 0" class="absolute top-4 left-4 z-10 px-3 py-1.5 bg-primary-500/90 text-white rounded-xl shadow-md text-[9px] font-black uppercase tracking-wider flex items-center gap-2">
              <span>{{ selectedRacks.length }} Rak Dipilih</span>
              <button class="hover:underline cursor-pointer" @click="clearSelection">Batal</button>
            </div>

            <div 
              class="transition-all duration-300 ease-out origin-center"
              :style="{ transform: `scale(${zoom})` }"
            >
              <!-- 8x8 Grid Map -->
              <div class="grid grid-cols-8 gap-2.5 p-4 border border-slate-100 dark:border-slate-800 rounded-2xl bg-slate-50/50 dark:bg-slate-950/30">
                <template v-for="(row, rIdx) in gridRows" :key="rIdx">
                  <div 
                    v-for="cell in row" 
                    :id="`rack-cell-${cell.label}`"
                    :key="cell.label"
                    :class="[
                      'w-14 h-14 rounded-xl border flex flex-col items-center justify-center relative cursor-pointer select-none transition-all duration-300 group',
                      isSelected(cell.rack) 
                        ? 'ring-2 ring-primary-500 ring-offset-2 dark:ring-offset-slate-900 scale-95 shadow-md z-10 border-primary-500' 
                        : '',
                      getCellStyle(cell.rack).bg
                    ]"
                    @mousedown="handleMouseDown(cell)"
                    @mouseenter="handleMouseEnter(cell)"
                  >
                    <!-- Coordinates or Actual Rack Name Label -->
                    <span :class="['text-[9px] font-black uppercase truncate max-w-full px-1', getCellStyle(cell.rack).text]">
                      {{ getCellLabel(cell) }}
                    </span>

                    <!-- Small badge/icon inside -->
                    <div class="absolute bottom-1 right-1 flex gap-0.5 items-center">
                      <LucideLock v-if="cell.rack?.is_full_override" class="w-2.5 h-2.5 text-slate-500" />
                      <span v-else-if="cell.rack" class="text-[7px] font-black text-slate-400/80">
                        {{ Math.round(((cell.rack.current_docs_count || 0) / (cell.rack.max_docs_capacity || 100)) * 100) }}%
                      </span>
                    </div>

                    <!-- Hover tooltip -->
                    <div class="absolute bottom-full left-1/2 -translate-x-1/2 mb-2 px-2.5 py-1.5 bg-[#0D121F] text-white text-[8px] font-bold rounded-lg opacity-0 pointer-events-none group-hover:opacity-150 transition-opacity whitespace-nowrap shadow-xl z-50">
                      <p class="font-black uppercase text-primary-400">{{ cell.rack?.name || 'Empty Cell' }}</p>
                      <p class="mt-0.5">Dept: {{ getDeptName(cell.rack?.department_id) }}</p>
                      <p v-if="cell.rack?.is_full_override" class="text-red-400">STATUS: LOCKED</p>
                      <p v-else-if="cell.rack">Filled: {{ cell.rack.current_docs_count || 0 }}/{{ cell.rack.max_docs_capacity || 100 }} docs</p>
                    </div>
                  </div>
                </template>
              </div>
            </div>
          </div>

          <!-- Bottom Stats Bar -->
          <div class="bg-slate-50 dark:bg-slate-950 px-6 py-4 border-t border-slate-100 dark:border-slate-800 flex justify-between items-center text-[10px] font-black text-slate-400 uppercase tracking-widest">
            <span id="storage-capacity-label">TOTAL STORAGE CAPACITY: {{ totalCapacity.toLocaleString() }} Units</span>
            <span id="allocated-percentage-label">ALLOCATED: {{ allocatedPercent }}%</span>
          </div>
        </div>
      </div>

      <!-- RIGHT PANEL: Zone Rule Editor & Logs -->
      <div class="lg:col-span-1 space-y-6">
        <div class="flex items-center justify-between px-2">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Zone Rule Editor</h3>
        </div>

        <div class="glass p-6 rounded-3xl space-y-6">
          <div class="space-y-4">
            <!-- Action type selector tab -->
            <div>
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest mb-2 block">Action Type</label>
              <div class="grid grid-cols-2 gap-2 bg-slate-100 dark:bg-slate-800/80 p-1 rounded-xl">
                <button 
                  id="action-reserve-btn"
                  :class="[
                    'py-2 text-[9px] font-black uppercase tracking-wider rounded-lg transition-all cursor-pointer',
                    formAction === 'reserve' ? 'bg-[#1E3A5F] text-white shadow-sm' : 'text-slate-500 hover:text-slate-700 dark:hover:text-slate-300'
                  ]"
                  @click="formAction = 'reserve'"
                >
                  Reserve Zone
                </button>
                <button 
                  id="action-release-btn"
                  :class="[
                    'py-2 text-[9px] font-black uppercase tracking-wider rounded-lg transition-all cursor-pointer',
                    formAction === 'release' ? 'bg-[#1E3A5F] text-white shadow-sm' : 'text-slate-500 hover:text-slate-700 dark:hover:text-slate-300'
                  ]"
                  @click="formAction = 'release'"
                >
                  Release Zone
                </button>
              </div>
            </div>

            <!-- Target Department (shown only when action is Reserve) -->
            <div v-if="formAction === 'reserve'" class="space-y-2">
              <label for="target-dept-select" class="text-[9px] font-black text-slate-400 uppercase tracking-widest block">Target Department</label>
              <select 
                id="target-dept-select"
                v-model="targetDeptId" 
                class="w-full bg-slate-50 dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl px-4 py-3 text-xs font-black text-[#1E3A5F] dark:text-white outline-none focus:ring-2 focus:ring-primary-500 cursor-pointer"
              >
                <option value="" disabled>Pilih Departemen...</option>
                <option v-for="dept in departments" :key="dept.id" :value="dept.id">
                  {{ dept.name }}
                </option>
              </select>
            </div>

            <!-- Effective Date Picker -->
            <div class="space-y-2">
              <label for="effective-date-picker" class="text-[9px] font-black text-slate-400 uppercase tracking-widest block">Tanggal Efektif</label>
              <input 
                id="effective-date-picker"
                v-model="effectiveDate" 
                type="date" 
                class="w-full bg-slate-50 dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl px-4 py-3 text-xs font-black text-[#1E3A5F] dark:text-white outline-none focus:ring-2 focus:ring-primary-500 cursor-pointer" 
              >
            </div>

            <!-- Conflict warning callout box -->
            <div v-if="hasConflict" class="p-4 bg-amber-50 dark:bg-amber-950/20 border-l-4 border-amber-500 text-amber-700 dark:text-amber-300 flex items-start gap-3 rounded-xl animate-pulse">
              <LucideAlertTriangle class="w-5 h-5 shrink-0 mt-0.5" />
              <div class="space-y-1">
                <p class="text-[10px] font-black uppercase tracking-widest">Peringatan Konflik Zonasi</p>
                <p class="text-[9px] font-medium leading-relaxed">Beberapa rak yang dipilih sedang dalam status Full Override (Terkunci). Menerapkan zonasi baru pada rak ini memerlukan persetujuan tambahan atau status lock harus dibuka terlebih dahulu.</p>
              </div>
            </div>

            <!-- Help text when no selection -->
            <p v-if="selectedRacks.length === 0" class="text-[9.5px] font-bold text-slate-400 text-center uppercase tracking-widest">
              Pilih satu atau beberapa rak di peta layout untuk menerapkan aturan
            </p>

            <!-- Buttons -->
            <div class="pt-4 space-y-2">
              <button 
                id="apply-zoning-btn"
                :disabled="isSaving || selectedRacks.length === 0"
                class="w-full py-4 bg-[#1E3A5F] disabled:opacity-50 text-white rounded-xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] disabled:cursor-not-allowed transition-all flex items-center justify-center gap-2 cursor-pointer"
                @click="applyZoning"
              >
                <LucideSave v-if="!isSaving" class="w-4 h-4" />
                <LucideLoader2 v-else class="w-4 h-4 animate-spin" />
                Apply Zoning Rule
              </button>
              
              <button 
                id="clear-selection-btn"
                :disabled="selectedRacks.length === 0"
                class="w-full py-3 bg-transparent text-slate-500 hover:text-slate-800 dark:hover:text-slate-300 disabled:opacity-30 rounded-xl text-[10px] font-black uppercase tracking-widest transition-all cursor-pointer"
                @click="clearSelection"
              >
                Reset Pilihan
              </button>
            </div>
          </div>
        </div>

        <!-- Logs / Active Schedule panel -->
        <div class="space-y-3">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-2">Zonation Schedule Log</h3>
          <div class="bg-white dark:bg-slate-900 border border-slate-200/60 dark:border-slate-800/40 rounded-3xl p-6 shadow-sm space-y-4">
            <div class="space-y-3 max-h-[300px] overflow-y-auto pr-2 custom-scrollbar">
              <div 
                v-for="sched in formattedLogs" 
                :key="sched.id"
                class="p-3 bg-slate-50 dark:bg-slate-800/50 rounded-xl border border-slate-100 dark:border-slate-700/50 flex flex-col gap-2 transition-all hover:scale-[1.01]"
              >
                <div class="flex justify-between items-center">
                  <span class="text-[9px] font-black text-slate-700 dark:text-white uppercase tracking-tight truncate max-w-[120px]">{{ sched.name }}</span>
                  <span
                    :class="[
                      'px-1.5 py-0.5 rounded text-[8px] font-black uppercase shrink-0',
                      sched.status === 'Active' ? 'bg-emerald-500/10 text-emerald-500' : 'bg-amber-500/10 text-amber-500'
                    ]"
                  >{{ sched.status }}</span>
                </div>
                <div class="flex justify-between text-[9px]">
                  <span class="font-bold text-slate-400 uppercase">Target:</span>
                  <span class="font-black text-slate-600 dark:text-slate-300 truncate max-w-[110px]">{{ sched.dept }}</span>
                </div>
                <div class="flex justify-between text-[9px] items-center">
                  <span class="font-bold text-slate-400 uppercase">Date:</span>
                  <span class="font-black text-slate-500">{{ sched.date }}</span>
                </div>
                <div class="flex justify-between text-[9px] items-center">
                  <span class="font-bold text-slate-400 uppercase">By:</span>
                  <span class="font-black text-slate-500">{{ sched.user }}</span>
                </div>
                <!-- Progress bar -->
                <div class="w-full bg-slate-200 dark:bg-slate-700 h-1 rounded-full overflow-hidden">
                  <div class="bg-primary-500 h-full" :style="{ width: `${sched.progress}%` }"/>
                </div>
              </div>
            </div>
          </div>
        </div>

      </div>
    </div>

    <!-- New Department Modal -->
    <Transition name="fade">
      <div v-if="showNewDeptModal" class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm z-[150] flex items-center justify-center p-4">
        <div v-motion-pop class="bg-white dark:bg-slate-900 rounded-3xl border border-slate-200 dark:border-slate-800 w-full max-w-md shadow-2xl p-6 space-y-6">
          <div class="flex justify-between items-center pb-4 border-b border-slate-100 dark:border-slate-800">
            <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">Tambah Departemen Baru</h3>
            <button class="text-slate-400 hover:text-slate-700 dark:hover:text-slate-300 cursor-pointer" @click="closeNewDeptModal">
              <LucideX class="w-5 h-5" />
            </button>
          </div>

          <form class="space-y-4" @submit.prevent="saveNewDepartment">
            <!-- Name -->
            <div class="space-y-2">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Nama Departemen</label>
              <input 
                v-model="newDeptForm.name"
                type="text" 
                required
                class="w-full bg-slate-50 dark:bg-slate-950 border border-slate-200 dark:border-slate-800 rounded-xl px-4 py-3 text-xs font-black text-[#1E3A5F] dark:text-white outline-none focus:ring-2 focus:ring-primary-500"
                placeholder="Contoh: Quality Assurance"
              >
            </div>

            <!-- Branch -->
            <div class="space-y-2">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Cabang (Branch)</label>
              <select 
                v-model="newDeptForm.branch_id"
                required
                class="w-full bg-slate-50 dark:bg-slate-950 border border-slate-200 dark:border-slate-800 rounded-xl px-4 py-3 text-xs font-black text-[#1E3A5F] dark:text-white outline-none focus:ring-2 focus:ring-primary-500 cursor-pointer"
              >
                <option value="" disabled>Pilih Cabang...</option>
                <option v-for="branch in branches" :key="branch.id" :value="branch.id">
                  {{ branch.name }}
                </option>
              </select>
            </div>

            <!-- Head of Department (Optional) -->
            <div class="space-y-2">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Kepala Departemen (Head)</label>
              <select 
                v-model="newDeptForm.head_id"
                class="w-full bg-slate-50 dark:bg-slate-950 border border-slate-200 dark:border-slate-800 rounded-xl px-4 py-3 text-xs font-black text-[#1E3A5F] dark:text-white outline-none focus:ring-2 focus:ring-primary-500 cursor-pointer"
              >
                <option :value="null">Belum Ditentukan</option>
                <option v-for="user in users" :key="user.id" :value="user.id">
                  {{ user.full_name }}
                </option>
              </select>
            </div>

            <!-- Actions -->
            <div class="pt-4 flex gap-3">
              <button 
                type="button"
                class="flex-1 py-3.5 border border-slate-200 dark:border-slate-850 rounded-xl text-[10px] font-black uppercase tracking-widest text-slate-500 hover:bg-slate-50 dark:hover:bg-slate-800 transition-all cursor-pointer"
                @click="closeNewDeptModal"
              >
                Batal
              </button>
              <button 
                type="submit"
                :disabled="isSavingDept"
                class="flex-1 py-3.5 bg-[#1E3A5F] disabled:opacity-50 text-white rounded-xl text-[10px] font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center justify-center gap-2 cursor-pointer"
              >
                <LucideLoader2 v-if="isSavingDept" class="w-4 h-4 animate-spin" />
                Simpan
              </button>
            </div>
          </form>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { 
  LucideMinus, LucidePlus, LucideLock, LucideSave, 
  LucideLoader2, LucideAlertTriangle, LucideX 
} from 'lucide-vue-next'
import { useApi } from '@/composables/useApi'

definePageMeta({
  layout: 'default'
})

const { $api } = useApi()

// State
const departments = ref([])
const racks = ref([])
const branches = ref([])
const users = ref([])
const floorsList = ref([])
const selectedDeptId = ref('')
const selectedRacks = ref([])
const activeFloorId = ref('')
const zoom = ref(1.0)
const formAction = ref('reserve') // 'reserve' | 'release'
const targetDeptId = ref('')
const effectiveDate = ref(new Date().toISOString().split('T')[0])
const isSaving = ref(false)

// New Department form state
const showNewDeptModal = ref(false)
const isSavingDept = ref(false)
const newDeptForm = ref({
  name: '',
  branch_id: '',
  head_id: null
})

// Drag-select helpers
const isSelecting = ref(false)

// Active zonation logs log
const zonationLogs = ref([])

const formattedLogs = computed(() => {
  return zonationLogs.value.map(log => {
    let isRelease = false
    try {
      const details = typeof log.details === 'string' ? JSON.parse(log.details) : log.details
      const itDept = departments.value.find(d => d.name.toLowerCase().includes('it'))
      if (details && (details.department_id === itDept?.id || details.DepartmentID === itDept?.id)) {
        isRelease = true
      }
    } catch {
      // ignore
    }

    const actionName = isRelease ? 'RELEASE' : 'RESERVE'
    return {
      id: log.id,
      name: `${actionName} - ${log.rack_name || 'RAK'}`,
      dept: log.department_name || 'Released (IT)',
      date: new Date(log.created_at).toISOString().split('T')[0],
      status: 'Active',
      progress: 100,
      user: log.user_name || 'System'
    }
  })
})

// Load database state
const loadDepartments = async () => {
  try {
    const res = await $api('/master/departments')
    departments.value = res.data || []
  } catch (error) {
    console.error('Failed to load departments', error)
  }
}

const loadRacks = async () => {
  try {
    const res = await $api('/master/racks')
    racks.value = res.data || []
  } catch (error) {
    console.error('Failed to load racks', error)
  }
}

const loadBranchesAndUsers = async () => {
  try {
    const [branchesRes, usersRes] = await Promise.all([
      $api('/master/branches-all'),
      $api('/users')
    ])
    branches.value = branchesRes.data || []
    users.value = usersRes.data || []
  } catch (error) {
    console.error('Failed to load branches and users', error)
  }
}

const loadZonationLogs = async () => {
  try {
    const res = await $api('/master/racks/zonation-logs')
    zonationLogs.value = res.data || []
  } catch (error) {
    console.error('Failed to load zonation logs', error)
  }
}

const loadFloors = async () => {
  try {
    const res = await $api('/master/floors')
    floorsList.value = res.data || []
    // Default to first floor
    if (floorsList.value.length > 0 && !activeFloorId.value) {
      activeFloorId.value = floorsList.value[0].id
    }
  } catch (error) {
    console.error('Failed to load floors', error)
  }
}

// Compute departments with assigned rack counts
const departmentsWithCounts = computed(() => {
  return departments.value.map(dept => {
    const count = racks.value.filter(r => r.department_id === dept.id).length
    return {
      ...dept,
      rackCount: count
    }
  })
})

// Highlight department-specific cells in layout map
const selectDepartment = (dept) => {
  if (selectedDeptId.value === dept.id) {
    // Deselect if clicked again
    selectedDeptId.value = ''
    selectedRacks.value = []
  } else {
    selectedDeptId.value = dept.id
    // Highlight all racks assigned to this department
    selectedRacks.value = racks.value.filter(r => r.department_id === dept.id)
  }
}

// Dynamic department hex colors for indicators
const getDeptColorHex = (name) => {
  const lowercaseName = (name || '').toLowerCase()
  if (lowercaseName.includes('legal')) return '#1E3A5F' // Navy
  if (lowercaseName.includes('finance') || lowercaseName.includes('tax')) return '#0EA5E9' // Sky Blue
  if (lowercaseName.includes('procurement')) return '#D97706' // Brown/Amber
  if (lowercaseName.includes('human') || lowercaseName.includes('hr')) return '#A855F7' // Purple
  if (lowercaseName.includes('it') || lowercaseName.includes('infra')) return '#10B981' // Emerald/Green
  return '#6366F1' // Fallback Indigo
}

// Zoom controls
const zoomIn = () => {
  if (zoom.value < 1.4) {
    zoom.value = parseFloat((zoom.value + 0.1).toFixed(1))
  }
}

const zoomOut = () => {
  if (zoom.value > 0.7) {
    zoom.value = parseFloat((zoom.value - 0.1).toFixed(1))
  }
}

const zoomReset = () => {
  zoom.value = 1.0
}

// Parse coordinates safely supporting float, string, or pgtype.Numeric object
const parseCoordinate = (val) => {
  if (val === null || val === undefined) return null
  if (typeof val === 'number') return val
  if (typeof val === 'string') return parseFloat(val)
  if (typeof val === 'object') {
    if (val.Valid === false) return null
    if (val.Int !== undefined && val.Exp !== undefined) {
      return val.Int * Math.pow(10, val.Exp)
    }
  }
  return null
}

// Parse floor_id from pgtype.UUID object
const parseFloorId = (val) => {
  if (!val) return null
  if (typeof val === 'string') return val
  if (typeof val === 'object') {
    if (val.Valid === false) return null
    return val.Bytes || val.String || null
  }
  return null
}

// 8x8 Grid Generator (dynamically maps coordinate-less racks to empty slots)
const filteredRacksByFloor = computed(() => {
  if (activeFloorId.value === '__unassigned__') {
    return racks.value.filter(r => {
      const fid = parseFloorId(r.floor_id)
      return !fid
    })
  }
  return racks.value.filter(r => {
    const fid = parseFloorId(r.floor_id)
    return fid === activeFloorId.value
  })
})

const gridRows = computed(() => {
  const rows = []
  
  // Create an empty 8x8 grid
  const grid = Array.from({ length: 8 }, () => Array(8).fill(null))
  
  // Racks without valid grid coordinates
  const unmapped = []
  
  filteredRacksByFloor.value.forEach(rack => {
    const px = parseCoordinate(rack.map_pos_x)
    const py = parseCoordinate(rack.map_pos_y)
    
    if (px >= 1 && px <= 8 && py >= 1 && py <= 8) {
      if (!grid[py - 1][px - 1]) {
        grid[py - 1][px - 1] = rack
      } else {
        unmapped.push(rack)
      }
    } else {
      unmapped.push(rack)
    }
  })
  
  // Dynamically place unmapped racks into empty grid cells
  let unmappedIdx = 0
  for (let y = 0; y < 8; y++) {
    for (let x = 0; x < 8; x++) {
      if (!grid[y][x] && unmappedIdx < unmapped.length) {
        const rack = unmapped[unmappedIdx++]
        rack.map_pos_x = x + 1
        rack.map_pos_y = y + 1
        grid[y][x] = rack
      }
    }
  }
  
  // Construct 8x8 grid rows
  for (let y = 1; y <= 8; y++) {
    const cols = []
    for (let x = 1; x <= 8; x++) {
      const rack = grid[y - 1][x - 1]
      cols.push({
        x,
        y,
        label: rack ? rack.name : `${String.fromCharCode(64 + x)}${y}`,
        rack
      })
    }
    rows.push(cols)
  }
  
  return rows
})

const getCellLabel = (cell) => {
  if (!cell.rack) return cell.label
  return cell.rack.name
    .replace('RACK-', '')
    .replace('Rack ', '')
    .replace('Department ', 'Dept ')
}

// Toggle cell selection
const isSelected = (rack) => {
  if (!rack) return false
  return selectedRacks.value.some(r => r.id === rack.id)
}

const handleMouseDown = (cell) => {
  if (!cell.rack) return
  isSelecting.value = true
  
  if (isSelected(cell.rack)) {
    selectedRacks.value = selectedRacks.value.filter(r => r.id !== cell.rack.id)
  } else {
    selectedRacks.value.push(cell.rack)
  }
}

const handleMouseEnter = (cell) => {
  if (!isSelecting.value || !cell.rack) return
  if (!isSelected(cell.rack)) {
    selectedRacks.value.push(cell.rack)
  }
}

const handleMouseUp = () => {
  isSelecting.value = false
}

const clearSelection = () => {
  selectedRacks.value = []
  selectedDeptId.value = ''
}

// Stats computations
const totalCapacity = computed(() => {
  return racks.value.reduce((acc, r) => acc + (r.max_docs_capacity || 100), 0)
})

const allocatedPercent = computed(() => {
  if (racks.value.length === 0) return 0
  const itDept = departments.value.find(d => d.name.toLowerCase().includes('it'))
  const defaultDeptId = itDept ? itDept.id : null
  const assigned = racks.value.filter(r => r.department_id && r.department_id !== defaultDeptId).length
  return Math.round((assigned / racks.value.length) * 100)
})

// Conflict detector
const hasConflict = computed(() => {
  return selectedRacks.value.some(r => r.is_full_override)
})

// Department names helper
const getDeptName = (deptId) => {
  if (!deptId) return 'Unassigned'
  const dept = departments.value.find(d => d.id === deptId)
  return dept ? dept.name : 'Unknown'
}

// Retrieve custom styling details for cells
const getCellStyle = (rack) => {
  if (!rack) {
    return {
      bg: 'bg-slate-50 dark:bg-slate-900 border-slate-200 dark:border-slate-800 text-slate-300 dark:text-slate-700',
      text: 'text-slate-300 dark:text-slate-700'
    }
  }

  if (rack.is_full_override) {
    return {
      bg: 'bg-slate-200 dark:bg-slate-800/80 border-slate-300 dark:border-slate-700 text-slate-500',
      text: 'text-slate-500'
    }
  }

  const dept = departments.value.find(d => d.id === rack.department_id)
  if (!dept) {
    return {
      bg: 'bg-slate-100 dark:bg-slate-800/40 border-slate-200 dark:border-slate-800 text-slate-400',
      text: 'text-slate-400'
    }
  }

  const name = dept.name.toLowerCase()
  if (name.includes('legal')) {
    return {
      bg: 'bg-blue-900/90 dark:bg-blue-950/90 text-white border-blue-700/50',
      text: 'text-blue-100'
    }
  } else if (name.includes('finance') || name.includes('tax')) {
    return {
      bg: 'bg-sky-500/20 dark:bg-sky-500/10 border-sky-400/50 text-sky-700 dark:text-sky-300',
      text: 'text-sky-800 dark:text-sky-300'
    }
  } else if (name.includes('procurement')) {
    return {
      bg: 'bg-amber-500/20 dark:bg-amber-600/10 border-amber-500/50 text-amber-800 dark:text-amber-300',
      text: 'text-amber-850 dark:text-amber-300'
    }
  } else if (name.includes('human') || name.includes('hr')) {
    return {
      bg: 'bg-purple-500/20 dark:bg-purple-600/10 border-purple-500/50 text-purple-800 dark:text-purple-300',
      text: 'text-purple-800 dark:text-purple-300'
    }
  } else if (name.includes('it') || name.includes('infra')) {
    return {
      bg: 'bg-emerald-500/20 dark:bg-emerald-600/10 border-emerald-500/50 text-emerald-800 dark:text-emerald-300',
      text: 'text-emerald-800 dark:text-emerald-300'
    }
  }

  return {
    bg: 'bg-indigo-500/20 dark:bg-indigo-600/10 border-indigo-500/50 text-indigo-800 dark:text-indigo-300',
    text: 'text-indigo-800 dark:text-indigo-300'
  }
}

// Apply Zoning Changes
const applyZoning = async () => {
  if (selectedRacks.value.length === 0) return
  if (!targetDeptId.value && formAction.value === 'reserve') {
    alert('Silakan pilih departemen tujuan!')
    return
  }

  isSaving.value = true
  try {
    const itDept = departments.value.find(d => d.name.toLowerCase().includes('it'))
    const defaultDeptId = itDept ? itDept.id : departments.value[0]?.id
    const deptId = formAction.value === 'reserve' ? targetDeptId.value : defaultDeptId

    if (!deptId) {
      throw new Error('Departemen default tidak ditemukan.')
    }

    const promises = selectedRacks.value.map(rack => {
      return $api(`/master/racks/${rack.id}`, {
        method: 'PUT',
        body: {
          department_id: deptId,
          name: rack.name,
          location_detail: rack.location_detail || ''
        }
      })
    })

    await Promise.all(promises)

    await loadZonationLogs()

    alert(`Aturan zonasi berhasil diterapkan ke ${selectedRacks.value.length} rak!`)
    clearSelection()
    await loadRacks()
  } catch (error) {
    alert('Gagal menerapkan zonasi: ' + (error.data?.message || error.message))
  } finally {
    isSaving.value = false
  }
}

// Modal actions
const openNewDeptModal = () => {
  newDeptForm.value = {
    name: '',
    branch_id: branches.value.length > 0 ? branches.value[0].id : '',
    head_id: null
  }
  showNewDeptModal.value = true
}

const closeNewDeptModal = () => {
  showNewDeptModal.value = false
}

const saveNewDepartment = async () => {
  if (!newDeptForm.value.name.trim()) return
  isSavingDept.value = true
  try {
    await $api('/master/departments', {
      method: 'POST',
      body: {
        name: newDeptForm.value.name,
        branch_id: newDeptForm.value.branch_id,
        head_id: newDeptForm.value.head_id
      }
    })
    alert('Departemen baru berhasil ditambahkan!')
    closeNewDeptModal()
    await loadDepartments()
  } catch (error) {
    alert('Gagal menambahkan departemen: ' + (error.data?.message || error.message))
  } finally {
    isSavingDept.value = false
  }
}

// Window event listeners
onMounted(async () => {
  await Promise.all([loadDepartments(), loadRacks(), loadBranchesAndUsers(), loadZonationLogs(), loadFloors()])
  if (import.meta.client) {
    window.addEventListener('mouseup', handleMouseUp)
  }
})
</script>

<style scoped>
@reference "../../assets/css/main.css";

.glass {
  @apply bg-white/70 dark:bg-slate-900/70 backdrop-blur-xl border border-white/20 dark:border-white/10 shadow-sm;
}

.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  @apply bg-transparent;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  @apply bg-slate-200 dark:bg-slate-800 rounded-full;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  @apply bg-slate-300 dark:bg-slate-700;
}
</style>
