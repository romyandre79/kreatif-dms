<template>
  <div class="space-y-8 pb-20 max-w-[1600px] mx-auto">
    <div class="flex flex-col lg:flex-row gap-8">
      <!-- Left Sidebar -->
      <div class="lg:w-[380px] space-y-6">
        <!-- Compliance Summary -->
        <div class="bg-white dark:bg-[#0D121F] rounded-lg p-8 shadow-sm border border-slate-100 dark:border-slate-800" v-motion-slide-visible-bottom>
          <p class="text-[10px] font-black uppercase tracking-[0.2em] text-slate-400 mb-6">RINGKASAN KEPATUHAN</p>
          <div class="grid grid-cols-2 gap-8 mb-8">
            <div class="space-y-1">
              <p class="text-4xl font-black text-[#1E3A5F] dark:text-white">{{ stats.score }}%</p>
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-tight leading-tight">Compliance Score</p>
            </div>
            <div class="space-y-1">
              <p class="text-4xl font-black text-red-500">{{ stats.lockouts }}</p>
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-tight leading-tight">Lockout Incidents</p>
            </div>
          </div>
          <button class="w-full py-4 px-6 bg-slate-50 dark:bg-slate-800/40 rounded-2xl flex items-center justify-between group transition-all hover:bg-primary-500/5">
            <span class="text-[11px] font-black uppercase tracking-widest text-slate-600 dark:text-slate-300 group-hover:text-primary-500">{{ stats.nonCompliant }} Akun Non-Compliant</span>
            <LucideArrowRight class="w-4 h-4 text-slate-400 group-hover:text-primary-500 transition-transform group-hover:translate-x-1" />
          </button>
        </div>

        <!-- Rules Enforcement -->
        <div class="bg-white dark:bg-[#0D121F] rounded-lg p-8 shadow-sm border border-slate-100 dark:border-slate-800" v-motion-slide-visible-bottom>
          <div class="flex items-center justify-between mb-8">
            <div class="flex items-center gap-3">
              <LucideShieldCheck class="w-5 h-5 text-primary-500" />
              <h3 class="font-black text-sm uppercase tracking-widest text-slate-800 dark:text-white">Aturan Penegakan PIN</h3>
            </div>
            <LucideLoader2 v-if="savingSettings" class="w-4 h-4 animate-spin text-primary-500" />
          </div>

          <div class="space-y-8">
            <div v-for="(rule, key) in pinRules" :key="key" class="bg-slate-50/50 dark:bg-slate-900/30 p-5 rounded-2xl border border-slate-100/50 dark:border-slate-800/50">
              <div class="flex items-start justify-between gap-4">
                <div class="space-y-1">
                  <p class="text-[11px] font-black uppercase tracking-tight text-slate-800 dark:text-slate-200">{{ $t(`admin.pin.rules.${key}.label`) }}</p>
                  <p class="text-[10px] text-slate-400 leading-relaxed font-medium">{{ $t(`admin.pin.rules.${key}.desc`) }}</p>
                </div>
                <Switch 
                  :model-value="rule.enabled" 
                  @update:model-value="val => updateSecuritySetting(rule.settingKey, val.toString(), 'boolean')" 
                />
              </div>
            </div>

            <div class="grid grid-cols-2 gap-4">
              <div v-for="config in pinConfigs" :key="config.key" class="space-y-2">
                <p class="text-[9px] font-black uppercase tracking-widest text-slate-400 ml-1">{{ $t(`admin.pin.rules.config.${config.key}`) }}</p>
                <div class="px-4 py-3 bg-slate-50 dark:bg-slate-800 rounded-xl border border-slate-100 dark:border-slate-700 text-xs font-bold text-slate-700 dark:text-slate-300">
                  {{ config.value }} {{ config.unit }}
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- PIN Preview / Info -->
        <div class="bg-[#EDF1F7] dark:bg-[#151B2B] rounded-lg p-8 shadow-sm border border-slate-200/50 dark:border-slate-800" v-motion-slide-visible-bottom>
          <p class="text-[10px] font-black uppercase tracking-[0.2em] text-center text-slate-400 mb-8">KEAMANAN BERLAPIS</p>
          <div class="space-y-6">
            <div class="flex items-center gap-4 p-4 bg-white dark:bg-[#1E2538] rounded-2xl shadow-sm">
              <div class="w-10 h-10 bg-green-500/10 rounded-xl flex items-center justify-center">
                <LucideFingerprint class="w-5 h-5 text-green-500" />
              </div>
              <div>
                <p class="text-[11px] font-black text-slate-800 dark:text-white uppercase tracking-tight">Biometric Ready</p>
                <p class="text-[9px] font-bold text-slate-400">Integrated with System</p>
              </div>
            </div>
            <div class="flex items-center gap-4 p-4 bg-white dark:bg-[#1E2538] rounded-2xl shadow-sm">
              <div class="w-10 h-10 bg-blue-500/10 rounded-xl flex items-center justify-center">
                <LucideHistory class="w-5 h-5 text-blue-500" />
              </div>
              <div>
                <p class="text-[11px] font-black text-slate-800 dark:text-white uppercase tracking-tight">Audit Trail</p>
                <p class="text-[9px] font-bold text-slate-400">Every Access Logged</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Main Content -->
      <div class="flex-1">
        <div class="bg-white dark:bg-[#0D121F] rounded-lg shadow-sm border border-slate-100 dark:border-slate-800 overflow-hidden" v-motion-slide-visible-bottom>
          <div class="p-10 border-b border-slate-100 dark:border-slate-800">
            <div class="flex items-start justify-between mb-8">
              <div>
                <h2 class="text-xl font-black text-slate-900 dark:text-white uppercase tracking-tight mb-1">Pendaftaran PIN Pengguna</h2>
                <p class="text-xs font-bold text-slate-400">Kelola status keamanan PIN untuk semua staf operasional</p>
              </div>
              <div class="flex items-center gap-4">
                <div class="relative w-64">
                  <LucideSearch class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-300" />
                  <input v-model="searchQuery" type="text" placeholder="Cari nama atau peran..." 
                         class="w-full pl-10 pr-4 py-2.5 bg-slate-50 dark:bg-slate-900 border-none rounded-xl text-xs font-bold placeholder:text-slate-300 focus:ring-2 focus:ring-primary-500/10" />
                </div>
                <button @click="fetchUsers" class="p-2.5 bg-slate-50 dark:bg-slate-900 rounded-xl text-slate-400 hover:text-primary-500 transition-colors">
                  <LucideRefreshCcw :class="{'animate-spin': loading}" class="w-5 h-5" />
                </button>
              </div>
            </div>

            <div class="overflow-x-auto">
              <table class="w-full text-left border-collapse">
                <thead>
                  <tr>
                    <th class="pb-6 text-[10px] font-black uppercase tracking-[0.2em] text-slate-400/80">NAMA PENGGUNA</th>
                    <th class="pb-6 text-[10px] font-black uppercase tracking-[0.2em] text-slate-400/80">PERAN</th>
                    <th class="pb-6 text-[10px] font-black uppercase tracking-[0.2em] text-slate-400/80">STATUS PIN</th>
                    <th class="pb-6 text-[10px] font-black uppercase tracking-[0.2em] text-slate-400/80">PEMBARUAN</th>
                    <th class="pb-6 text-[10px] font-black uppercase tracking-[0.2em] text-slate-400/80 text-center">GAGAL</th>
                    <th class="pb-6 text-[10px] font-black uppercase tracking-[0.2em] text-slate-400/80 text-right">AKSI</th>
                  </tr>
                </thead>
                <tbody v-if="!loading">
                  <tr v-for="user in filteredUsers" :key="user.id" class="group border-t border-slate-50 dark:border-slate-800/50">
                    <td class="py-6">
                      <div class="flex items-center gap-4">
                        <div class="w-12 h-12 rounded-full overflow-hidden shadow-sm bg-slate-800 ring-2 ring-white dark:ring-slate-800">
                          <img v-if="user.avatar_url?.Valid" :src="user.avatar_url.String" class="w-full h-full object-cover" />
                          <div v-else class="w-full h-full bg-slate-800 flex items-center justify-center font-black text-white text-xs">
                            {{ user.full_name.charAt(0) }}
                          </div>
                        </div>
                        <div>
                          <p class="text-[13px] font-black text-slate-800 dark:text-white uppercase tracking-tight leading-none mb-1">{{ user.full_name }}</p>
                          <p class="text-[10px] font-bold text-slate-400 group-hover:text-primary-500 transition-colors lowercase">{{ user.email }}</p>
                        </div>
                      </div>
                    </td>
                    <td class="py-6">
                      <div class="px-5 py-2 bg-slate-50/80 dark:bg-slate-800/40 rounded-xl inline-block">
                        <span class="text-[9px] font-black uppercase tracking-widest text-slate-500 dark:text-slate-400 leading-none">
                          {{ user.role_name.String || user.role_name }}
                        </span>
                      </div>
                    </td>
                    <td class="py-6">
                      <div class="inline-flex items-center px-5 py-1.5 rounded-full border text-[9px] font-black uppercase tracking-[0.1em]"
                           :class="getStatusUI(getStatus(user))">
                        {{ getStatus(user) === 'SET' ? 'ACTIVE' : getStatus(user) }}
                      </div>
                    </td>
                    <td class="py-6">
                      <p class="text-[11px] font-black text-slate-600 dark:text-slate-400 uppercase tracking-tight">
                        {{ formatDate(user.pin_updated_at) }}
                      </p>
                    </td>
                    <td class="py-6 text-center">
                      <p class="text-[14px] font-black" :class="user.pin_failed_attempts > 0 ? 'text-red-500' : 'text-slate-900 dark:text-white'">
                        {{ user.pin_failed_attempts }}
                      </p>
                    </td>
                    <td class="py-6 text-right">
                      <button @click="openPinModal(user)" class="p-2.5 bg-slate-50 dark:bg-slate-800 rounded-xl text-slate-400 hover:text-primary-500 hover:bg-primary-500/5 transition-all">
                        <LucideEdit class="w-4 h-4" />
                      </button>
                    </td>
                  </tr>
                  <tr v-if="filteredUsers.length === 0">
                    <td colspan="6" class="py-32 text-center">
                      <div class="w-20 h-20 bg-slate-50 dark:bg-slate-800/50 rounded-lg flex items-center justify-center mx-auto mb-6">
                        <LucideShieldCheck class="w-8 h-8 text-slate-200 dark:text-slate-700" />
                      </div>
                      <p class="text-[10px] font-black uppercase tracking-[0.2em] text-slate-400">No authorized users found</p>
                    </td>
                  </tr>
                </tbody>
                <tbody v-else>
                  <tr v-for="i in 4" :key="i" class="border-t border-slate-50">
                    <td colspan="6" class="py-6">
                      <div class="flex items-center gap-4 animate-pulse">
                        <div class="w-12 h-12 rounded-full bg-slate-100 dark:bg-slate-800"></div>
                        <div class="space-y-2">
                          <div class="h-3 w-40 bg-slate-100 dark:bg-slate-800 rounded"></div>
                          <div class="h-2 w-24 bg-slate-50 dark:bg-slate-900 rounded"></div>
                        </div>
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
            
            <div class="mt-8 pt-8 border-t border-slate-50 dark:border-slate-800 flex items-center justify-between">
              <p class="text-[10px] font-black uppercase tracking-[0.2em] text-slate-400">
                MENAMPILKAN {{ filteredUsers.length }} DARI {{ totalUsers }} PENGGUNA TERPILIH
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- PIN Input Modal -->
    <Transition name="fade">
      <div v-if="showPinModal" class="fixed inset-0 z-[100] flex items-center justify-center p-6 bg-slate-900/60 backdrop-blur-sm">
        <div class="bg-[#EDF1F7] dark:bg-[#151B2B] w-full max-w-md rounded-lg p-10 shadow-2xl border border-white/20 relative" v-motion-pop>
          <button @click="showPinModal = false" class="absolute right-8 top-8 p-2 text-slate-400 hover:text-slate-600 dark:hover:text-white transition-colors">
            <LucideX class="w-6 h-6" />
          </button>

          <div class="text-center mb-10">
            <div class="w-20 h-20 rounded-full overflow-hidden mx-auto mb-6 ring-4 ring-white dark:ring-slate-800 shadow-xl">
              <img v-if="selectedUser.avatar_url?.Valid" :src="selectedUser.avatar_url.String" class="w-full h-full object-cover" />
              <div v-else class="w-full h-full bg-slate-800 flex items-center justify-center font-black text-white text-xl">
                {{ selectedUser.full_name.charAt(0) }}
              </div>
            </div>
            <h3 class="text-xl font-black text-slate-800 dark:text-white uppercase tracking-tight mb-1">Set User PIN</h3>
            <p class="text-xs font-bold text-slate-400">{{ selectedUser.full_name }}</p>
          </div>

          <!-- PIN Dots -->
          <div class="flex justify-center gap-3 mb-12">
            <div v-for="i in 6" :key="i" 
                 class="w-12 h-16 rounded-2xl border-2 flex items-center justify-center transition-all bg-white dark:bg-[#1E2538]"
                 :class="i <= pinInput.length ? 'border-primary-500 shadow-lg shadow-primary-500/10' : 'border-slate-200 dark:border-slate-700'">
              <div v-if="i <= pinInput.length" class="w-3 h-3 rounded-full bg-primary-600 dark:bg-primary-400"></div>
            </div>
          </div>

          <!-- Keypad -->
          <div class="grid grid-cols-3 gap-4 mb-8">
            <button v-for="n in 9" :key="n" @click="handleKey(n.toString())"
                    class="h-16 rounded-2xl bg-white dark:bg-[#1E2538] flex items-center justify-center text-2xl font-black text-slate-700 dark:text-white hover:bg-slate-50 dark:hover:bg-primary-500/10 transition-all shadow-sm active:scale-95">
              {{ n }}
            </button>
            <button @click="handleDelete" class="h-16 rounded-2xl flex items-center justify-center text-slate-400 hover:text-red-500 transition-colors">
              <LucideDelete class="w-8 h-8" />
            </button>
            <button @click="handleKey('0')" class="h-16 rounded-2xl bg-white dark:bg-[#1E2538] flex items-center justify-center text-2xl font-black text-slate-700 dark:text-white hover:bg-slate-50 dark:hover:bg-primary-500/10 transition-all shadow-sm active:scale-95">
              0
            </button>
            <button @click="saveUserPin" :disabled="pinInput.length < 6 || savingPin"
                    class="h-16 rounded-2xl bg-primary-500 flex items-center justify-center text-white hover:bg-primary-600 transition-all shadow-lg shadow-primary-500/20 active:scale-95 disabled:opacity-50 disabled:grayscale">
              <LucideLoader2 v-if="savingPin" class="w-8 h-8 animate-spin" />
              <LucideCheckCircle2 v-else class="w-8 h-8" />
            </button>
          </div>

          <p v-if="errorMessage" class="text-center text-[10px] font-bold text-red-500 uppercase tracking-widest mt-4">
            {{ errorMessage }}
          </p>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { 
  LucideShieldCheck, 
  LucideInfo, 
  LucideSearch, 
  LucideFilter, 
  LucideRefreshCcw, 
  LucideArrowRight,
  LucideDelete,
  LucideCheckCircle2,
  LucideLoader2,
  LucideEdit,
  LucideX,
  LucideFingerprint,
  LucideHistory
} from 'lucide-vue-next'
import PageHeader from '~/components/PageHeader.vue'
import Switch from '~/components/Switch.vue'

const { t } = useI18n()
const config = useRuntimeConfig()
const { $api } = useApi()

const loading = ref(true)
const allUsers = ref([])
const searchQuery = ref('')
const currentPage = ref(1)
const itemsPerPage = 10

// PIN Modal State
const showPinModal = ref(false)
const selectedUser = ref(null)
const pinInput = ref('')
const savingPin = ref(false)
const errorMessage = ref('')

// Settings State
const savingSettings = ref(false)
const rawSettings = ref([])

// Target roles for PIN management
const allowedRoles = ['superadmin', 'manajer', 'admin doc controller', 'kepala doc controller']

const filteredUsers = computed(() => {
  const users = allUsers.value.filter(u => {
    const roleName = (u.role_name?.String || u.role_name || '').toLowerCase()
    const roleMatch = allowedRoles.some(r => roleName.includes(r))
    
    const searchMatch = u.full_name.toLowerCase().includes(searchQuery.value.toLowerCase()) || 
                       u.email.toLowerCase().includes(searchQuery.value.toLowerCase())
    return roleMatch && searchMatch
  })

  // Apply pagination
  const start = (currentPage.value - 1) * itemsPerPage
  return users.slice(start, start + itemsPerPage)
})

const totalUsers = computed(() => {
  return allUsers.value.filter(u => {
    const roleName = (u.role_name?.String || u.role_name || '').toLowerCase()
    return allowedRoles.some(r => roleName.includes(r))
  }).length
})

const getStatus = (user) => {
  if (user.pin_locked_until?.Valid && new Date(user.pin_locked_until.Time) > new Date()) {
    return 'LOCKED'
  }
  return user.pin_status?.String || 'NOT SET'
}

const getStatusUI = (status) => {
  switch (status) {
    case 'SET':
    case 'ACTIVE':
      return 'bg-green-50 text-green-600 border-green-100'
    case 'LOCKED':
      return 'bg-red-50 text-red-600 border-red-100'
    case 'EXPIRED':
      return 'bg-orange-50 text-orange-600 border-orange-100'
    default:
      return 'bg-slate-50 text-slate-400 border-slate-100'
  }
}

const stats = computed(() => {
  const users = allUsers.value.filter(u => {
    const roleName = (u.role_name?.String || u.role_name || '').toLowerCase()
    return allowedRoles.some(r => roleName.includes(r))
  })
  
  const secured = users.filter(u => getStatus(u) === 'SET' || getStatus(u) === 'ACTIVE').length
  const total = users.length
  const score = total > 0 ? Math.round((secured / total) * 100) : 0
  const lockouts = users.filter(u => getStatus(u) === 'LOCKED').length
  const nonCompliant = total - secured
  
  return { score, total, lockouts, nonCompliant }
})

const pinRules = computed(() => ({
  mandatory: { 
    enabled: getSettingValue('pin_approval_mandatory') === 'true',
    settingKey: 'pin_approval_mandatory'
  },
  security: { 
    enabled: getSettingValue('pin_security_download_print') === 'true',
    settingKey: 'pin_security_download_print'
  },
  verification: { 
    enabled: getSettingValue('pin_archive_verification') === 'true',
    settingKey: 'pin_archive_verification'
  }
}))

const pinConfigs = computed(() => [
  { key: 'length', value: getSettingValue('pin_length') || '6', unit: 'digit' },
  { key: 'max_attempts', value: getSettingValue('pin_max_attempts') || '3', unit: 'Kali' },
  { key: 'lockout_duration', value: getSettingValue('pin_lockout_duration') || '15', unit: 'Menit' },
  { key: 'expiry', value: getSettingValue('pin_expiry_days') || '90', unit: 'Hari' }
])

const getSettingValue = (key) => {
  const s = rawSettings.value.find(s => s.key === key)
  return s ? s.value : null
}

const fetchSettings = async () => {
  try {
    const res = await $api(`${config.public.apiBase}/master/settings/Security`)
    if (res && res.data) {
      rawSettings.value = res.data
    }
  } catch (err) {
    console.error('Failed to fetch settings:', err)
  }
}

const updateSecuritySetting = async (key, value, type, desc) => {
  savingSettings.value = true
  try {
    await $api(`${config.public.apiBase}/master/settings/Security`, {
      method: 'POST',
      body: { key, value, type, description: desc }
    })
    await fetchSettings()
  } catch (err) {
    console.error('Failed to update setting:', err)
  } finally {
    savingSettings.value = false
  }
}

const fetchUsers = async () => {
  loading.value = true
  try {
    const res = await $api(`${config.public.apiBase}/users/`)
    if (res && res.data) {
      allUsers.value = res.data
    }
  } catch (err) {
    console.error('Failed to fetch users:', err)
  } finally {
    loading.value = false
  }
}

const openPinModal = (user) => {
  selectedUser.value = user
  pinInput.value = ''
  errorMessage.value = ''
  showPinModal.value = true
}

const handleKey = (key) => {
  if (pinInput.value.length < 6) {
    pinInput.value += key
  }
}

const handleDelete = () => {
  pinInput.value = pinInput.value.slice(0, -1)
}

const saveUserPin = async () => {
  if (pinInput.value.length < 6) return
  
  savingPin.value = true
  errorMessage.value = ''
  try {
    await $api(`${config.public.apiBase}/auth/user-pin`, {
      method: 'POST',
      body: {
        user_id: selectedUser.value.id,
        pin: pinInput.value
      }
    })
    showPinModal.value = false
    await fetchUsers()
  } catch (err) {
    errorMessage.value = err.response?._data?.message || 'Failed to set PIN'
  } finally {
    savingPin.value = false
  }
}

onMounted(() => {
  fetchUsers()
  fetchSettings()
})

const formatDate = (dateStr) => {
  if (!dateStr || !dateStr.Valid) return 'Never'
  const date = new Date(dateStr.Time)
  const options = { day: '2-digit', month: 'short', year: 'numeric' }
  return date.toLocaleDateString('id-ID', options)
}
</script>

<style scoped>
.fade-enter-active, .fade-leave-active { transition: opacity 0.3s; }
.fade-enter-from, .fade-leave-to { opacity: 0; }
</style>
