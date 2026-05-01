<template>
  <div class="space-y-8">
    <div class="flex items-center justify-between">
      <div class="space-y-1">
        <h2 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('admin.user_management.title') }}</h2>
        <p class="text-xs font-bold text-slate-400 uppercase tracking-widest">{{ $t('admin.user_management.subtitle') }}</p>
      </div>
      <div class="flex items-center gap-3">
        <button @click="openModal()" class="flex items-center gap-2 px-6 py-2.5 bg-[#1E3A5F] text-white rounded-xl shadow-lg shadow-blue-900/20 hover:bg-[#152943] transition-all font-bold text-sm active:scale-95">
          <LucideUserPlus class="w-4 h-4" />
          {{ $t('admin.user_management.add_btn') }}
        </button>
      </div>
    </div>

    <!-- Filters & Stats -->
    <div class="grid grid-cols-1 lg:grid-cols-4 gap-6">
      <div class="lg:col-span-3 bg-white/80 dark:bg-slate-900/50 backdrop-blur-xl border border-slate-200 dark:border-slate-800 rounded-3xl p-6 flex flex-wrap items-center gap-4">
        <div class="relative flex-1 min-w-[200px]">
          <LucideSearch class="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" />
          <input 
            type="text" 
            v-model="filters.search"
            :placeholder="$t('admin.user_management.search_placeholder')"
            class="w-full pl-11 pr-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold focus:outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all"
          />
        </div>
        <select v-model="filters.role_id" class="px-5 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10">
          <option value="">{{ $t('admin.user_management.filter_role') }}</option>
          <option v-for="r in rolesList" :key="r.id" :value="r.id">{{ r.name }}</option>
        </select>
        <select v-model="filters.status" class="px-5 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10">
          <option value="">{{ $t('admin.user_management.filter_status') }}</option>
          <option value="active">{{ $t('admin.user_management.status.active') }}</option>
          <option value="inactive">{{ $t('admin.user_management.status.inactive') }}</option>
          <option value="pending">{{ $t('admin.user_management.status.pending') }}</option>
        </select>
      </div>
      <div class="bg-primary-500 rounded-3xl p-6 text-white flex items-center justify-between shadow-xl shadow-primary-500/20">
        <div>
          <p class="text-[10px] font-black uppercase tracking-widest opacity-80">{{ $t('admin.user_management.total_label') }}</p>
          <p class="text-3xl font-black">{{ users.length }}</p>
        </div>
        <div class="w-12 h-12 bg-white/20 rounded-2xl flex items-center justify-center">
          <LucideUsers class="w-6 h-6" />
        </div>
      </div>
    </div>

    <!-- Table -->
    <div class="bg-white/80 dark:bg-slate-900/50 backdrop-blur-xl border border-slate-200 dark:border-slate-800 rounded-3xl overflow-hidden relative min-h-[200px]">
      <div v-if="loading" class="absolute inset-0 bg-white/50 dark:bg-slate-900/50 backdrop-blur-[2px] z-10 flex items-center justify-center">
        <div class="flex flex-col items-center gap-3">
          <div class="w-10 h-10 border-4 border-primary-500/20 border-t-primary-500 rounded-full animate-spin"></div>
          <p class="text-[10px] font-black text-primary-500 uppercase tracking-widest">Loading Users...</p>
        </div>
      </div>

      <div class="overflow-x-auto">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">
              <th class="px-8 py-5">{{ $t('admin.user_management.table.name_identity') }}</th>
              <th class="px-6 py-5">{{ $t('admin.user_management.table.email_role') }}</th>
              <th class="px-6 py-5 text-center">{{ $t('admin.user_management.table.status') }}</th>
              <th class="px-6 py-5">{{ $t('admin.user_management.table.last_login') }}</th>
              <th class="px-8 py-5 text-right">{{ $t('admin.user_management.table.actions') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50 dark:divide-slate-800/50">
            <tr v-for="user in filteredUsers" :key="user.id" class="group hover:bg-slate-50 dark:hover:bg-white/5 transition-colors">
              <td class="px-8 py-5">
                <div class="flex items-center gap-4 text-left">
                  <div class="w-10 h-10 rounded-full bg-slate-100 dark:bg-slate-800 flex items-center justify-center text-primary-500 font-black text-xs overflow-hidden">
                    <img v-if="user.avatar_url" :src="user.avatar_url" class="w-full h-full object-cover">
                    <span v-else-if="user.full_name">{{ user.full_name.split(' ').map(n => n[0]).join('').substring(0,2).toUpperCase() }}</span>
                    <LucideUser v-else class="w-5 h-5 opacity-40" />
                  </div>
                  <div>
                    <p class="text-[11px] font-black text-slate-900 dark:text-white uppercase tracking-tight">{{ user.full_name }}</p>
                    <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ user.username || user.email.split('@')[0] }}</p>
                  </div>
                </div>
              </td>
              <td class="px-6 py-5">
                <p class="text-[11px] font-bold text-slate-600 dark:text-slate-300">{{ user.email }}</p>
                <span class="inline-block px-2 py-0.5 bg-blue-500/10 text-blue-500 rounded text-[9px] font-black uppercase mt-1">
                  {{ getRoleName(user.role_id) }}
                </span>
              </td>
              <td class="px-6 py-5">
                <div class="flex justify-center">
                  <span :class="getStatusBadge(user.status)" class="px-3 py-1 rounded-full text-[9px] font-black uppercase tracking-widest">
                    {{ $t(`admin.user_management.status.${user.status}`) }}
                  </span>
                </div>
              </td>
              <td class="px-6 py-5">
                <p class="text-[10px] font-bold text-slate-500">{{ user.last_login || 'Never' }}</p>
              </td>
              <td class="px-8 py-5 text-right">
                <div class="flex items-center justify-end gap-2">
                  <button @click="openModal(user)" class="p-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl text-slate-400 hover:text-blue-500 transition-colors">
                    <LucideEdit3 class="w-4 h-4" />
                  </button>
                  <button @click="deleteUser(user.id)" class="p-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl text-slate-400 hover:text-red-500 transition-colors">
                    <LucideTrash2 class="w-4 h-4" />
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- User Modal (Centered) -->
    <Transition name="fade-scale">
      <div v-if="showModal" class="fixed inset-0 z-[150] flex items-center justify-center p-6">
        <div class="absolute inset-0 bg-slate-950/60 backdrop-blur-md" @click="showModal = false"></div>
        <div class="relative w-full max-w-2xl bg-white dark:bg-slate-900 rounded-[3rem] shadow-2xl overflow-hidden flex flex-col max-h-[90vh]">
          <header class="p-10 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between">
            <div class="space-y-1">
              <h3 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">
                {{ isEdit ? $t('admin.user_management.modal.title_edit') : $t('admin.user_management.modal.title_add') }}
              </h3>
              <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ isEdit ? 'Update user profile and access' : 'Create new user account' }}</p>
            </div>
            <button @click="showModal = false" class="p-3 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-2xl transition-all">
              <LucideX class="w-6 h-6 text-slate-400" />
            </button>
          </header>

          <form @submit.prevent="saveUser" class="flex-1 overflow-y-auto p-10 space-y-10 custom-scrollbar">
            <!-- Profile Photos -->
            <div class="flex items-center gap-8 p-8 bg-slate-50 dark:bg-slate-800/50 rounded-[2.5rem] border border-slate-100 dark:border-slate-800">
              <!-- Avatar -->
              <div class="relative group cursor-pointer" @click="$refs.avatarInput.click()">
                <div class="w-24 h-24 rounded-full bg-white dark:bg-slate-900 flex items-center justify-center border-2 border-primary-500/20 group-hover:border-primary-500 transition-all overflow-hidden shadow-xl ring-4 ring-white dark:ring-slate-800">
                  <img v-if="formData.avatar_url" :src="formData.avatar_url" class="w-full h-full object-cover">
                  <LucideUser v-else class="w-10 h-10 text-slate-300" />
                </div>
                <div class="absolute inset-0 bg-primary-500/80 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity rounded-full">
                  <LucideCamera class="w-6 h-6 text-white" />
                </div>
                <input type="file" ref="avatarInput" class="hidden" accept="image/*" @change="handleAvatarUpload" />
                <p class="mt-3 text-[9px] font-black text-center text-slate-400 uppercase tracking-widest">Avatar</p>
              </div>

              <!-- Signature -->
              <div class="relative group cursor-pointer flex-grow" @click="$refs.signInput.click()">
                <div class="h-24 w-full bg-white dark:bg-slate-900 rounded-2xl flex items-center justify-center border-2 border-dashed border-slate-200 dark:border-slate-700 group-hover:border-primary-500 transition-all overflow-hidden p-4">
                  <img v-if="formData.signature_url" :src="formData.signature_url" class="max-h-full object-contain">
                  <LucidePenTool v-else class="w-8 h-8 text-slate-300" />
                </div>
                <div class="absolute inset-0 bg-primary-500/80 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity rounded-2xl">
                  <LucideUpload class="w-6 h-6 text-white" />
                </div>
                <input type="file" ref="signInput" class="hidden" accept="image/*" @change="handleSignUpload" />
                <p class="mt-3 text-[9px] font-black text-left text-slate-400 uppercase tracking-widest px-2">E-Signature</p>
              </div>
            </div>

            <!-- Account Info -->
            <div class="space-y-8">
              <div class="flex items-center gap-3">
                <span class="w-1.5 h-6 bg-primary-500 rounded-full"></span>
                <h4 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">
                  {{ $t('admin.user_management.modal.section_account') }}
                </h4>
              </div>
              <div class="grid grid-cols-2 gap-8">
                <div class="space-y-2">
                  <label class="text-[10px] font-black text-slate-500 uppercase px-1">{{ $t('admin.user_management.modal.label_name') }}</label>
                  <input v-model="formData.full_name" type="text" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all" required />
                </div>
                <div class="space-y-2">
                  <label class="text-[10px] font-black text-slate-500 uppercase px-1">{{ $t('admin.user_management.modal.label_username') }}</label>
                  <input v-model="formData.username" type="text" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all" required />
                </div>
              </div>
              <div class="space-y-2">
                <label class="text-[10px] font-black text-slate-500 uppercase px-1">{{ $t('admin.user_management.modal.label_email') }}</label>
                <input v-model="formData.email" type="email" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all" required />
              </div>
            </div>

            <!-- Access & Security -->
            <div class="space-y-8">
              <div class="flex items-center gap-3">
                <span class="w-1.5 h-6 bg-primary-500 rounded-full"></span>
                <h4 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">
                  {{ $t('admin.user_management.modal.section_access') }}
                </h4>
              </div>
              <div class="grid grid-cols-2 gap-8">
                <div class="space-y-2">
                  <label class="text-[10px] font-black text-slate-500 uppercase px-1">{{ $t('admin.user_management.modal.label_role') }}</label>
                  <select v-model="formData.role_id" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10" required>
                    <option v-for="r in rolesList" :key="r.id" :value="r.id">{{ r.name }}</option>
                  </select>
                </div>
                <div class="space-y-2">
                  <label class="text-[10px] font-black text-slate-500 uppercase px-1">Departemen</label>
                  <select v-model="formData.department_id" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10">
                    <option value="">Pilih Departemen</option>
                    <option v-for="d in deptList" :key="d.id" :value="d.id">{{ d.name }}</option>
                  </select>
                </div>
              </div>
              <div class="grid grid-cols-2 gap-8">
                <div class="space-y-2">
                  <label class="text-[10px] font-black text-slate-500 uppercase px-1">{{ $t('admin.user_management.modal.label_status') }}</label>
                  <select v-model="formData.status" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10">
                    <option value="active">{{ $t('admin.user_management.status.active') }}</option>
                    <option value="inactive">{{ $t('admin.user_management.status.inactive') }}</option>
                    <option value="pending">{{ $t('admin.user_management.status.pending') }}</option>
                  </select>
                </div>
                <div class="space-y-2">
                  <label class="text-[10px] font-black text-slate-500 uppercase px-1">
                    {{ isEdit ? $t('admin.user_management.modal.label_password_edit') : $t('admin.user_management.modal.label_password') }}
                  </label>
                  <div class="relative">
                    <input :type="showPassword ? 'text' : 'password'" v-model="formData.password" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all" :required="!isEdit" />
                    <button type="button" @click="showPassword = !showPassword" class="absolute right-6 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-600 transition-colors">
                      <LucideEye v-if="!showPassword" class="w-5 h-5" />
                      <LucideEyeOff v-else class="w-5 h-5" />
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <!-- Preferences -->
            <div class="space-y-8">
              <div class="flex items-center gap-3">
                <span class="w-1.5 h-6 bg-primary-500 rounded-full"></span>
                <h4 class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">
                  {{ $t('admin.user_management.modal.section_prefs') }}
                </h4>
              </div>
              <div class="flex items-center justify-between p-8 bg-slate-50 dark:bg-slate-800 rounded-[2.5rem] border border-slate-100 dark:border-slate-800">
                <div class="space-y-1">
                  <div class="flex items-center gap-3">
                    <LucideShieldCheck class="w-5 h-5 text-primary-500" />
                    <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ $t('admin.user_management.modal.label_mfa') }} (2FA)</p>
                  </div>
                  <p class="text-[10px] font-bold text-slate-400 uppercase">{{ $t('admin.user_management.modal.label_mfa_desc') }}</p>
                </div>
                <button type="button" @click="formData.is_mfa_enabled = !formData.is_mfa_enabled" class="w-12 h-6 rounded-full transition-all duration-300 relative p-1" :class="formData.is_mfa_enabled ? 'bg-primary-500' : 'bg-slate-300 dark:bg-slate-700'">
                  <div class="w-4 h-4 bg-white rounded-full transition-transform" :class="{ 'translate-x-6': formData.is_mfa_enabled }"></div>
                </button>
              </div>
            </div>
          </form>

          <footer class="p-10 border-t border-slate-100 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-900/50 flex items-center justify-end gap-6">
            <button @click="showModal = false" class="px-8 py-4 text-[11px] font-black text-slate-400 hover:text-slate-600 uppercase tracking-widest transition-colors active:scale-95">
              {{ $t('admin.user_management.modal.btn_cancel') }}
            </button>
            <button @click="saveUser" :disabled="saving" class="px-12 py-4 bg-[#1E3A5F] text-white rounded-2xl text-[11px] font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 active:scale-95 disabled:opacity-50">
              <LucideSave class="w-5 h-5" />
              {{ saving ? $t('admin.user_management.modal.saving') : $t('admin.user_management.modal.btn_save') }}
            </button>
          </footer>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useApi } from '@/composables/useApi'
import { useI18n } from 'vue-i18n'
import { 
  LucideUsers, LucideSearch, LucideUserPlus, LucideEdit3, 
  LucideTrash2, LucideX, LucideSave, LucideEye, LucideEyeOff,
  LucideUser, LucideCamera, LucidePenTool, LucideUpload, LucideShieldCheck
} from 'lucide-vue-next'

const { $api } = useApi()
const { t } = useI18n()
const config = useRuntimeConfig()

const showModal = ref(false)
const isEdit = ref(false)
const saving = ref(false)
const loading = ref(false)
const showPassword = ref(false)

const rolesList = ref([])
const deptList = ref([])
const users = ref([])

const filters = ref({
  search: '',
  role_id: '',
  status: ''
})

const formData = ref({
  id: '',
  full_name: '',
  username: '',
  email: '',
  role_id: '',
  department_id: '',
  status: 'active',
  password: '',
  avatar_url: '',
  signature_url: '',
  is_mfa_enabled: false
})

const fetchUsers = async () => {
  loading.value = true
  try {
    const res = await $api(`${config.public.apiBase}/users`)
    if (res && res.data) {
      users.value = res.data
    }
  } catch (err) {
    console.error('Failed to fetch users:', err)
  } finally {
    loading.value = false
  }
}

const fetchRoles = async () => {
  try {
    const res = await $api(`${config.public.apiBase}/master/roles`)
    if (res && res.data) {
      rolesList.value = res.data
    }
  } catch (err) {
    console.error('Failed to fetch roles:', err)
  }
}

const fetchDepts = async () => {
  try {
    const res = await $api(`${config.public.apiBase}/master/departments`)
    if (res && res.data) {
      deptList.value = res.data
    }
  } catch (err) {
    console.error('Failed to fetch departments:', err)
  }
}

onMounted(() => {
  fetchUsers()
  fetchRoles()
  fetchDepts()
})

const handleAvatarUpload = (e) => {
  const file = e.target.files[0]
  if (file) {
    const reader = new FileReader()
    reader.onload = (event) => {
      formData.value.avatar_url = event.target.result
    }
    reader.readAsDataURL(file)
  }
}

const handleSignUpload = (e) => {
  const file = e.target.files[0]
  if (file) {
    const reader = new FileReader()
    reader.onload = (event) => {
      formData.value.signature_url = event.target.result
    }
    reader.readAsDataURL(file)
  }
}

const filteredUsers = computed(() => {
  return users.value.filter(u => {
    const matchSearch = !filters.value.search || 
      u.full_name.toLowerCase().includes(filters.value.search.toLowerCase()) ||
      u.email.toLowerCase().includes(filters.value.search.toLowerCase())
    const matchRole = !filters.value.role_id || u.role_id == filters.value.role_id
    const matchStatus = !filters.value.status || u.status === filters.value.status
    return matchSearch && matchRole && matchStatus
  })
})

const getRoleName = (roleId) => {
  const role = rolesList.value.find(r => r.id === roleId)
  return role ? role.name : 'Unknown'
}

const getStatusBadge = (status) => {
  switch (status) {
    case 'active': return 'bg-green-500/10 text-green-500 border border-green-500/20'
    case 'inactive': return 'bg-red-500/10 text-red-500 border border-red-500/20'
    case 'pending': return 'bg-amber-500/10 text-amber-500 border border-amber-500/20'
    case 'approved': return 'bg-blue-500/10 text-blue-500 border border-blue-500/20'
    default: return 'bg-slate-500/10 text-slate-500'
  }
}

const openModal = (user = null) => {
  if (user) {
    isEdit.value = true
    formData.value = { 
      id: user.id,
      full_name: user.full_name,
      username: user.username || user.email.split('@')[0],
      email: user.email,
      role_id: user.role_id,
      department_id: user.department_id || '',
      status: user.status || (user.is_active ? 'active' : 'inactive'),
      password: '',
      avatar_url: user.avatar_url || '',
      signature_url: user.signature_url || '',
      is_mfa_enabled: user.is_mfa_enabled || false
    }
  } else {
    isEdit.value = false
    formData.value = {
      id: '',
      full_name: '',
      username: '',
      email: '',
      role_id: rolesList.value.length > 0 ? rolesList.value[0].id : '',
      department_id: '',
      status: 'active',
      password: '',
      avatar_url: '',
      signature_url: '',
      is_mfa_enabled: false
    }
  }
  showModal.value = true
}

const saveUser = async () => {
  saving.value = true
  try {
    const payload = {
      full_name: formData.value.full_name,
      email: formData.value.email,
      role_id: parseInt(formData.value.role_id),
      department_id: formData.value.department_id,
      status: formData.value.status,
      avatar_url: formData.value.avatar_url,
      signature_url: formData.value.signature_url,
      is_mfa_enabled: formData.value.is_mfa_enabled
    }

    if (isEdit.value) {
      await $api(`${config.public.apiBase}/users/${formData.value.id}`, {
        method: 'PUT',
        body: payload
      })
    } else {
      await $api(`${config.public.apiBase}/users`, {
        method: 'POST',
        body: { ...payload, password: formData.value.password }
      })
    }
    await fetchUsers()
    showModal.value = false
  } catch (err) {
    console.error('Failed to save user:', err)
    alert('Gagal menyimpan user: ' + (err.data?.message || err.message))
  } finally {
    saving.value = false
  }
}

const deleteUser = async (id) => {
  if (confirm('Apakah Anda yakin ingin menghapus user ini?')) {
    try {
      await $api(`${config.public.apiBase}/users/${id}`, {
        method: 'DELETE'
      })
      await fetchUsers()
    } catch (err) {
      console.error('Failed to delete user:', err)
      alert('Gagal menghapus user: ' + (err.data?.message || err.message))
    }
  }
}

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

.fade-scale-enter-active, .fade-scale-leave-active {
  transition: all 0.4s cubic-bezier(0.34, 1.56, 0.64, 1);
}
.fade-scale-enter-from, .fade-scale-leave-to {
  opacity: 0;
  transform: scale(0.9) translateY(20px);
}
</style>
