<template>
  <div class="space-y-8">
    <div class="flex items-center justify-between">
      <div class="space-y-1">
        <h2 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Manajemen User</h2>
        <p class="text-xs font-bold text-slate-400 uppercase tracking-widest">Kelola identitas, akses, dan keamanan pengguna sistem.</p>
      </div>
      <div class="flex items-center gap-3">
        <button @click="openModal()" class="flex items-center gap-2 px-6 py-2.5 bg-[#1E3A5F] text-white rounded-xl shadow-lg shadow-blue-900/20 hover:bg-[#152943] transition-all font-bold text-sm active:scale-95">
          <LucideUserPlus class="w-4 h-4" />
          Tambah User Baru
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
            placeholder="Cari user (nama, email...)"
            class="w-full pl-11 pr-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold focus:outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all"
          />
        </div>
        <select v-model="filters.role" class="px-5 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10">
          <option value="">Semua Role</option>
          <option v-for="r in roles" :key="r" :value="r">{{ r }}</option>
        </select>
        <select v-model="filters.status" class="px-5 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10">
          <option value="">Semua Status</option>
          <option value="active">Aktif</option>
          <option value="inactive">Non-Aktif</option>
          <option value="pending">Pending</option>
        </select>
      </div>
      <div class="bg-primary-500 rounded-3xl p-6 text-white flex items-center justify-between shadow-xl shadow-primary-500/20">
        <div>
          <p class="text-[10px] font-black uppercase tracking-widest opacity-80">Total Users</p>
          <p class="text-3xl font-black">{{ users.length }}</p>
        </div>
        <div class="w-12 h-12 bg-white/20 rounded-2xl flex items-center justify-center">
          <LucideUsers class="w-6 h-6" />
        </div>
      </div>
    </div>

    <!-- Table -->
    <div class="bg-white/80 dark:bg-slate-900/50 backdrop-blur-xl border border-slate-200 dark:border-slate-800 rounded-3xl overflow-hidden">
      <div class="overflow-x-auto">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">
              <th class="px-8 py-5">Nama & Identitas</th>
              <th class="px-6 py-5">Email & Role</th>
              <th class="px-6 py-5 text-center">Status</th>
              <th class="px-6 py-5">Last Login</th>
              <th class="px-8 py-5 text-right">Aksi</th>
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
                    <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ user.username }}</p>
                  </div>
                </div>
              </td>
              <td class="px-6 py-5">
                <p class="text-[11px] font-bold text-slate-600 dark:text-slate-300">{{ user.email }}</p>
                <span class="inline-block px-2 py-0.5 bg-blue-500/10 text-blue-500 rounded text-[9px] font-black uppercase mt-1">{{ user.role }}</span>
              </td>
              <td class="px-6 py-5">
                <div class="flex justify-center">
                  <span :class="getStatusBadge(user.status)" class="px-3 py-1 rounded-full text-[9px] font-black uppercase tracking-widest">
                    {{ user.status }}
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

    <!-- User Modal (Sidebar) -->
    <Transition name="slide-right">
      <div v-if="showModal" class="fixed inset-0 z-50 flex justify-end">
        <div class="absolute inset-0 bg-slate-950/40 backdrop-blur-sm" @click="showModal = false"></div>
        <div class="relative w-full max-w-xl bg-white dark:bg-slate-900 h-full shadow-2xl flex flex-col">
          <header class="p-8 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between">
            <h3 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ isEdit ? 'Edit User' : 'Tambah User Baru' }}</h3>
            <button @click="showModal = false" class="p-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl transition-colors">
              <LucideX class="w-6 h-6 text-slate-400" />
            </button>
          </header>

          <form @submit.prevent="saveUser" class="flex-1 overflow-y-auto p-8 space-y-8 custom-scrollbar">
            <!-- Account Info -->
            <div class="space-y-6">
              <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-50 dark:border-slate-800 pb-2">Informasi Akun</h4>
              <div class="grid grid-cols-2 gap-6">
                <div class="space-y-2">
                  <label class="text-[10px] font-black text-slate-500 uppercase px-1">Full Name</label>
                  <input v-model="formData.full_name" type="text" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all" required />
                </div>
                <div class="space-y-2">
                  <label class="text-[10px] font-black text-slate-500 uppercase px-1">Username</label>
                  <input v-model="formData.username" type="text" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all" required />
                </div>
              </div>
              <div class="space-y-2">
                <label class="text-[10px] font-black text-slate-500 uppercase px-1">Email Address</label>
                <input v-model="formData.email" type="email" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all" required />
              </div>
            </div>

            <!-- Access & Security -->
            <div class="space-y-6">
              <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-50 dark:border-slate-800 pb-2">Akses & Keamanan</h4>
              <div class="grid grid-cols-2 gap-6">
                <div class="space-y-2">
                  <label class="text-[10px] font-black text-slate-500 uppercase px-1">Role / Jabatan</label>
                  <select v-model="formData.role" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10" required>
                    <option v-for="r in roles" :key="r" :value="r">{{ r }}</option>
                  </select>
                </div>
                <div class="space-y-2">
                  <label class="text-[10px] font-black text-slate-500 uppercase px-1">Status Akun</label>
                  <select v-model="formData.status" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10">
                    <option value="active">Aktif</option>
                    <option value="inactive">Non-Aktif</option>
                  </select>
                </div>
              </div>
              <div class="space-y-2">
                <label class="text-[10px] font-black text-slate-500 uppercase px-1">{{ isEdit ? 'Ganti Password (Biarkan kosong jika tidak diubah)' : 'Password' }}</label>
                <div class="relative">
                  <input :type="showPassword ? 'text' : 'password'" v-model="formData.password" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all" :required="!isEdit" />
                  <button type="button" @click="showPassword = !showPassword" class="absolute right-5 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-600 transition-colors">
                    <LucideEye v-if="!showPassword" class="w-4 h-4" />
                    <LucideEyeOff v-else class="w-4 h-4" />
                  </button>
                </div>
              </div>
            </div>

            <!-- Preferences -->
            <div class="space-y-6">
              <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-50 dark:border-slate-800 pb-2">Preferensi Sistem</h4>
              <div class="flex items-center justify-between p-6 bg-slate-50 dark:bg-slate-800 rounded-[2rem] border border-slate-100 dark:border-slate-700">
                <div class="space-y-1">
                  <p class="text-[11px] font-black uppercase tracking-tight">Two-Factor Authentication</p>
                  <p class="text-[9px] font-bold text-slate-400 uppercase">Wajibkan verifikasi OTP saat login.</p>
                </div>
                <button type="button" @click="formData.mfa = !formData.mfa" class="w-10 h-5 rounded-full transition-all duration-300 relative p-1" :class="formData.mfa ? 'bg-primary-500' : 'bg-slate-300'">
                  <div class="w-3 h-3 bg-white rounded-full transition-transform" :class="{ 'translate-x-5': formData.mfa }"></div>
                </button>
              </div>
            </div>
          </form>

          <footer class="p-8 border-t border-slate-100 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-900/50 flex items-center justify-end gap-4">
            <button @click="showModal = false" class="px-8 py-3 text-[10px] font-black text-slate-400 hover:text-slate-600 uppercase tracking-widest transition-colors">Batal</button>
            <button @click="saveUser" :disabled="saving" class="px-10 py-3 bg-[#1E3A5F] text-white rounded-xl text-[10px] font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 active:scale-95">
              <LucideSave class="w-4 h-4" />
              {{ saving ? 'Menyimpan...' : 'Simpan User' }}
            </button>
          </footer>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { 
  LucideUsers, LucideSearch, LucideUserPlus, LucideEdit3, 
  LucideTrash2, LucideX, LucideSave, LucideEye, LucideEyeOff
} from 'lucide-vue-next'

const showModal = ref(false)
const isEdit = ref(false)
const saving = ref(false)
const showPassword = ref(false)

const roles = ['Admin', 'User Umum', 'Section Head', 'Head', 'Management']

const filters = ref({
  search: '',
  role: '',
  status: ''
})

const users = ref([
  { id: 1, full_name: 'Budi Santoso', username: 'budi.s', email: 'budi@kreatif.com', role: 'Admin', status: 'active', last_login: 'Today, 08:30', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Budi' },
  { id: 2, full_name: 'Siti Aminah', username: 'siti.a', email: 'siti@kreatif.com', role: 'Head', status: 'active', last_login: 'Yesterday, 14:20', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Siti' },
  { id: 3, full_name: 'Andi Wijaya', username: 'andi.w', email: 'andi@kreatif.com', role: 'Section Head', status: 'inactive', last_login: '2 days ago', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Andi' },
  { id: 4, full_name: 'Dewi Lestari', username: 'dewi.l', email: 'dewi@kreatif.com', role: 'User Umum', status: 'active', last_login: 'Today, 09:45', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Dewi' },
  { id: 5, full_name: 'Iwan Setiawan', username: 'iwan.s', email: 'iwan@kreatif.com', role: 'Management', status: 'pending', last_login: null, avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Iwan' },
])

const formData = ref({
  full_name: '',
  username: '',
  email: '',
  role: 'User Umum',
  status: 'active',
  password: '',
  mfa: false
})

const filteredUsers = computed(() => {
  return users.value.filter(u => {
    const matchSearch = !filters.value.search || 
      u.full_name.toLowerCase().includes(filters.value.search.toLowerCase()) ||
      u.email.toLowerCase().includes(filters.value.search.toLowerCase())
    const matchRole = !filters.value.role || u.role === filters.value.role
    const matchStatus = !filters.value.status || u.status === filters.value.status
    return matchSearch && matchRole && matchStatus
  })
})

const getStatusBadge = (status) => {
  switch (status) {
    case 'active': return 'bg-green-500/10 text-green-500 border border-green-500/20'
    case 'inactive': return 'bg-red-500/10 text-red-500 border border-red-500/20'
    case 'pending': return 'bg-amber-500/10 text-amber-500 border border-amber-500/20'
    default: return 'bg-slate-500/10 text-slate-500'
  }
}

const openModal = (user = null) => {
  if (user) {
    isEdit.value = true
    formData.value = { ...user, password: '' }
  } else {
    isEdit.value = false
    formData.value = {
      full_name: '',
      username: '',
      email: '',
      role: 'User Umum',
      status: 'active',
      password: '',
      mfa: false
    }
  }
  showModal.value = true
}

const saveUser = async () => {
  saving.value = true
  // Simulate API call
  await new Promise(resolve => setTimeout(resolve, 1000))
  saving.value = false
  showModal.value = false
}

const deleteUser = (id) => {
  if (confirm('Apakah Anda yakin ingin menghapus user ini?')) {
    users.value = users.value.filter(u => u.id !== id)
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

.slide-right-enter-active, .slide-right-leave-active {
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}
.slide-right-enter-from, .slide-right-leave-to {
  transform: translateX(100%);
}
</style>
