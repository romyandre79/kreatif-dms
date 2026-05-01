<template>
  <div class="flex flex-col h-[calc(100vh-8rem)] bg-slate-50/50 dark:bg-slate-950/50 rounded-[2.5rem] overflow-hidden border border-slate-200 dark:border-slate-800">
    <!-- Top Alert Banner (Dirty State) -->
    <Transition name="slide-down">
      <div v-if="isDirty" class="bg-amber-500/10 border-b border-amber-500/20 px-8 py-3 flex items-center justify-between backdrop-blur-md relative z-[20]">
        <div class="flex items-center gap-3">
          <LucideAlertTriangle class="w-4 h-4 text-amber-500" />
          <span class="text-[11px] font-black text-amber-600 dark:text-amber-400 uppercase tracking-widest">
            PERINGATAN: Terdapat perubahan yang belum dipublikasikan pada skema "{{ selectedRole?.name }}".
          </span>
        </div>
        <div class="flex items-center gap-6">
          <button @click="resetChanges" class="text-[10px] font-black text-slate-400 hover:text-slate-600 uppercase tracking-widest transition-colors">Batalkan Semua</button>
          <button @click="saveChanges" :disabled="saving" class="px-6 py-2 bg-amber-500 text-white rounded-xl text-[10px] font-black uppercase tracking-widest shadow-lg shadow-amber-500/20 hover:bg-amber-600 transition-all active:scale-95">
             {{ saving ? 'Menyimpan...' : 'Tinjau & Simpan' }}
          </button>
        </div>
      </div>
    </Transition>

    <div class="flex flex-1 overflow-hidden">
      <!-- Left Sidebar: Hierarchy -->
      <aside class="w-80 border-r border-slate-200 dark:border-slate-800 bg-white/50 dark:bg-slate-900/50 backdrop-blur-xl flex flex-col">
        <div class="p-8 space-y-6">
          <div class="space-y-1">
            <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">Hierarchy Peran</h3>
          </div>
          
          <div class="relative group">
            <LucideSearch class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 group-focus-within:text-primary-500 transition-colors" />
            <input 
              type="text" 
              v-model="searchRole"
              placeholder="Filter peran..."
              class="w-full pl-11 pr-4 py-3 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold focus:outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all"
            />
          </div>
        </div>

        <nav class="flex-1 overflow-y-auto px-4 pb-8 custom-scrollbar space-y-1">
          <div v-if="loadingRoles" class="p-8 text-center animate-pulse">
            <div class="h-4 bg-slate-200 dark:bg-slate-800 rounded-full w-3/4 mx-auto mb-4"></div>
            <div class="h-4 bg-slate-200 dark:bg-slate-800 rounded-full w-1/2 mx-auto"></div>
          </div>
          <button 
            v-for="role in filteredRoles" 
            :key="role.id"
            @click="selectRole(role)"
            class="w-full flex items-center justify-between px-5 py-4 rounded-2xl transition-all duration-300 group relative"
            :class="[
              selectedRole?.id === role.id 
                ? 'bg-[#1E3A5F] text-white shadow-xl shadow-blue-900/20' 
                : 'text-slate-500 dark:text-slate-400 hover:bg-white dark:hover:bg-slate-800 hover:shadow-md'
            ]"
          >
            <div class="flex items-center gap-4 text-left">
              <LucideShield class="w-5 h-5" :class="selectedRole?.id === role.id ? 'text-blue-400' : 'text-slate-400'" />
              <div>
                <p class="text-[11px] font-black uppercase tracking-tight">{{ role.name }}</p>
                <p v-if="role.disabled" class="text-[9px] font-bold opacity-60 uppercase">(Disabled)</p>
              </div>
            </div>
            <div v-if="selectedRole?.id === role.id" class="w-2 h-2 rounded-full bg-blue-400 animate-pulse"></div>
            <LucideLock v-if="role.disabled" class="w-3 h-3 opacity-40" />
          </button>
        </nav>
      </aside>

      <!-- Center Content: Matrix -->
      <main class="flex-1 bg-white dark:bg-slate-900 flex flex-col overflow-hidden">
        <header class="px-10 py-8 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between bg-slate-50/30 dark:bg-slate-900/30">
          <div class="flex items-center gap-4 text-left">
            <h2 class="text-lg font-black text-[#1E3A5F] dark:text-white uppercase tracking-tighter">Matriks Izin: {{ selectedRole?.name }}</h2>
          </div>
          <div class="flex items-center gap-3 bg-slate-100 dark:bg-slate-800 p-1.5 rounded-xl">
            <span class="text-[10px] font-black text-slate-500 uppercase px-3 tracking-widest">Full Access</span>
            <button 
              @click="toggleFullAccess"
              class="w-10 h-5 rounded-full transition-all duration-500 relative p-1"
              :class="isFullAccess ? 'bg-primary-500' : 'bg-slate-300 dark:bg-slate-700'"
            >
              <div class="w-3 h-3 bg-white rounded-full shadow-sm transition-transform duration-500" :class="{ 'translate-x-5': isFullAccess }"></div>
            </button>
          </div>
        </header>

        <div class="flex-1 overflow-auto custom-scrollbar">
          <div v-if="loadingPermissions" class="p-20 text-center text-slate-300 font-black uppercase tracking-widest text-xs animate-pulse">
            Sinkronisasi matriks izin...
          </div>
          <table v-else class="w-full text-left border-collapse">
            <thead>
              <tr class="text-[9px] font-black text-slate-400 uppercase tracking-[0.2em] border-b border-slate-50 dark:border-slate-800 sticky top-0 bg-white dark:bg-slate-900 z-10">
                <th class="p-8 pl-10 w-64">Modul</th>
                <th v-for="action in actions" :key="action" class="p-4 text-center">{{ action }}</th>
              </tr>
            </thead>
            <tbody v-for="group in permissionGroups" :key="group.name">
              <tr class="bg-slate-50/50 dark:bg-slate-800/30">
                <td colspan="10" class="px-10 py-3 text-[9px] font-black text-primary-500 uppercase tracking-widest border-y border-slate-100 dark:border-slate-800">
                  {{ group.name }}
                </td>
              </tr>
              <tr v-for="mod in group.modules" :key="mod.id" class="group hover:bg-slate-50/50 dark:hover:bg-slate-800/20 transition-colors">
                <td class="px-10 py-6">
                  <span class="text-[11px] font-black text-[#1E3A5F] dark:text-slate-200 uppercase tracking-tight group-hover:text-primary-500 transition-colors">{{ mod.name }}</span>
                </td>
                <td v-for="action in actions" :key="action" class="p-4 text-center">
                  <label v-if="mod.allowed_actions.includes(action)" class="relative inline-flex items-center justify-center cursor-pointer group/check">
                    <input 
                      type="checkbox" 
                      :checked="isFullAccess || hasPermission(mod.id, action)"
                      @change="togglePermission(mod.id, action)"
                      class="sr-only peer"
                    >
                    <div class="w-6 h-6 border-2 border-slate-200 dark:border-slate-700 rounded-lg peer-checked:bg-primary-500 peer-checked:border-primary-500 transition-all duration-300 flex items-center justify-center shadow-sm">
                      <LucideCheck class="w-3.5 h-3.5 text-white scale-0 peer-checked:scale-100 transition-transform duration-300 stroke-[4px]" />
                    </div>
                  </label>
                  <div v-else class="w-1.5 h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full mx-auto"></div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </main>

      <!-- Right Panel: Info -->
      <aside class="w-[400px] bg-slate-50/50 dark:bg-slate-900/50 border-l border-slate-200 dark:border-slate-800 flex flex-col p-4 overflow-y-auto custom-scrollbar gap-10">
        <!-- Role Header -->
        <div class="flex items-start gap-6">
          <div class="w-16 h-16 rounded-[1.5rem] bg-white dark:bg-slate-800 shadow-xl border border-slate-100 dark:border-slate-700 flex items-center justify-center text-primary-500">
            <LucideShieldCheck class="w-8 h-8" />
          </div>
          <div class="space-y-1 text-left">
            <h2 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ selectedRole?.name }}</h2>
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">User Authority</p>
          </div>
        </div>

        <!-- Description -->
        <div class="space-y-3 text-left">
          <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Deskripsi</label>
          <p class="text-[12px] font-medium text-slate-500 dark:text-slate-400 leading-relaxed italic">
            {{ selectedRole?.description || 'Tidak ada deskripsi untuk peran ini.' }}
          </p>
        </div>

        <!-- Stats Grid -->
        <div class="grid grid-cols-2 gap-6">
          <div class="p-6 bg-white dark:bg-slate-800/50 rounded-[2.5rem] shadow-sm border border-slate-100 dark:border-slate-700 space-y-2 text-left">
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest">User Aktif</p>
            <p class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tighter">{{ selectedRole?.user_count || 0 }}</p>
          </div>
        </div>

      </aside>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRuntimeConfig } from '#app'
import { 
  LucideShieldCheck, LucideSearch, LucideLock, LucideFileStack, 
  LucideUser, LucideCheck, LucideShield, LucideUsers,
  LucideLayoutDashboard, LucideShieldAlert, LucideCheckCircle2,
  LucideXCircle, LucideAlertTriangle, LucideInfo
} from 'lucide-vue-next'
import { useApi } from '@/composables/useApi'

const { $api } = useApi()
const config = useRuntimeConfig()

const searchRole = ref('')
const isDirty = ref(false)
const isFullAccess = ref(false)
const saving = ref(false)
const loadingRoles = ref(false)
const loadingPermissions = ref(false)

const actions = ['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT']

const roles = ref([])
const selectedRole = ref(null)
const systemModules = ref([])
const rolePermissions = ref({}) // Format: { roleId: Set('moduleId:action') }

const filteredRoles = computed(() => {
  if (!searchRole.value) return roles.value
  const q = searchRole.value.toLowerCase()
  return roles.value.filter(r => r.name.toLowerCase().includes(q))
})

const permissionGroups = computed(() => {
  const groups = {}
  systemModules.value.forEach(mod => {
    if (!groups[mod.category]) {
      groups[mod.category] = { name: mod.category, modules: [] }
    }
    groups[mod.category].modules.push(mod)
  })
  return Object.values(groups)
})

const activePermissionsCount = computed(() => {
  if (!selectedRole.value) return 0
  return (rolePermissions.value[selectedRole.value.id] || new Set()).size
})

const fetchRoles = async () => {
  loadingRoles.value = true
  try {
    const res = await $api(`${config.public.apiBase}/master/roles`)
    if (res && res.data) {
      roles.value = res.data
      if (roles.value.length > 0 && !selectedRole.value) {
        selectRole(roles.value[0])
      }
    }
  } catch (err) {
    console.error('Failed to fetch roles:', err)
  } finally {
    loadingRoles.value = false
  }
}

const fetchModules = async () => {
  try {
    const res = await $api(`${config.public.apiBase}/master/modules`)
    if (res && res.data) {
      systemModules.value = res.data
    }
  } catch (err) {
    console.error('Failed to fetch modules:', err)
  }
}

const fetchRolePermissions = async (roleId) => {
  loadingPermissions.value = true
  try {
    const res = await $api(`${config.public.apiBase}/master/roles/${roleId}/permissions`)
    if (res && res.data) {
      const perms = new Set()
      res.data.forEach(p => {
        perms.add(`${p.module_id}:${p.action}`)
      })
      rolePermissions.value[roleId] = perms
      isDirty.value = false
    }
  } catch (err) {
    console.error('Failed to fetch permissions:', err)
  } finally {
    loadingPermissions.value = false
  }
}

const selectRole = (role) => {
  selectedRole.value = role
  
  // Auto set Full Access for superadmin
  if (role.name.toLowerCase() === 'superadmin') {
    isFullAccess.value = true
  } else {
    isFullAccess.value = false
  }

  if (!rolePermissions.value[role.id]) {
    fetchRolePermissions(role.id)
  } else {
    isDirty.value = false
  }
}

const hasPermission = (modId, action) => {
  if (!selectedRole.value) return false
  const perms = rolePermissions.value[selectedRole.value.id]
  return perms ? perms.has(`${modId}:${action}`) : false
}

const togglePermission = (modId, action) => {
  if (!selectedRole.value) return
  const roleId = selectedRole.value.id
  if (!rolePermissions.value[roleId]) {
    rolePermissions.value[roleId] = new Set()
  }
  
  const key = `${modId}:${action}`
  const perms = rolePermissions.value[roleId]
  
  if (perms.has(key)) {
    perms.delete(key)
  } else {
    perms.add(key)
  }
  
  isDirty.value = true
}

const toggleFullAccess = () => {
  isFullAccess.value = !isFullAccess.value
  isDirty.value = true
  
  if (isFullAccess.value && selectedRole.value) {
    const roleId = selectedRole.value.id
    const perms = new Set()
    systemModules.value.forEach(mod => {
      mod.allowed_actions.forEach(action => {
        perms.add(`${mod.id}:${action}`)
      })
    })
    rolePermissions.value[roleId] = perms
  }
}

const saveChanges = async () => {
  if (!selectedRole.value) return
  
  saving.value = true
  try {
    const roleId = selectedRole.value.id
    const perms = Array.from(rolePermissions.value[roleId] || []).map(p => {
      const [modId, action] = p.split(':')
      return { module_id: modId, action }
    })
    
    await $api(`${config.public.apiBase}/master/roles/${roleId}/permissions`, {
      method: 'POST',
      body: perms
    })
    
    isDirty.value = false
    // Show success toast or something
  } catch (err) {
    console.error('Failed to save permissions:', err)
  } finally {
    saving.value = false
  }
}

const resetChanges = () => {
  if (selectedRole.value) {
    fetchRolePermissions(selectedRole.value.id)
  }
}

onMounted(() => {
  fetchRoles()
  fetchModules()
})

const riskLevel = computed(() => {
  if (selectedRole.value?.name.toLowerCase() === 'admin' || selectedRole.value?.name.toLowerCase() === 'superadmin') return 5
  return 2
})

const riskText = computed(() => riskLevel.value >= 4 ? 'KRITIS' : 'NORMAL')
const riskColor = computed(() => riskLevel.value >= 4 ? 'text-red-500' : 'text-blue-500')
const riskBg = computed(() => riskLevel.value >= 4 ? 'bg-red-500' : 'bg-blue-500')

const restrictedActions = [
  'Hapus Audit Trail Utama',
  'Modifikasi Lisensi Vendor'
]

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: rgba(148, 163, 184, 0.2);
  border-radius: 10px;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: rgba(148, 163, 184, 0.4);
}

.slide-down-enter-active, .slide-down-leave-active {
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}
.slide-down-enter-from, .slide-down-leave-to {
  transform: translateY(-100%);
  opacity: 0;
}
</style>
