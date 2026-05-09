<template>
  <div class="space-y-8">
    <div class="flex items-center justify-between">
      <div class="space-y-1 text-left">
        <h2 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ t('admin.modules.title') }}</h2>
        <p class="text-xs font-bold text-slate-400 uppercase tracking-widest">{{ t('admin.modules.subtitle') }}</p>
      </div>
      <div class="flex items-center gap-3">
        <button @click="openModal()" class="flex items-center gap-2 px-6 py-2.5 bg-[#1E3A5F] text-white rounded-xl shadow-lg shadow-blue-900/20 hover:bg-[#152943] transition-all font-bold text-sm active:scale-95">
          <LucidePlusCircle class="w-4 h-4" />
          {{ t('admin.modules.add_btn') }}
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
            :placeholder="t('admin.modules.search_placeholder')"
            class="w-full pl-11 pr-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold focus:outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all"
          />
        </div>
        <select v-model="filters.category" class="px-5 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10">
          <option value="">Semua Kategori</option>
          <option v-for="cat in categories" :key="cat" :value="cat">{{ cat }}</option>
        </select>
      </div>
      <div class="bg-primary-500 rounded-3xl p-6 text-white flex items-center justify-between shadow-xl shadow-primary-500/20">
        <div>
          <p class="text-[10px] font-black uppercase tracking-widest opacity-80">{{ t('admin.modules.total_label') }}</p>
          <p class="text-3xl font-black">{{ modules.length }}</p>
        </div>
        <div class="w-12 h-12 bg-white/20 rounded-2xl flex items-center justify-center">
          <LucideLayers class="w-6 h-6" />
        </div>
      </div>
    </div>

    <!-- Table -->
    <div class="bg-white/80 dark:bg-slate-900/50 backdrop-blur-xl border border-slate-200 dark:border-slate-800 rounded-3xl overflow-hidden">
      <div class="overflow-x-auto">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">
              <th class="px-8 py-5">{{ t('admin.modules.table.module_id') }}</th>
              <th class="px-6 py-5">{{ t('admin.modules.table.parent_path') }}</th>
              <th class="px-6 py-5">{{ t('admin.modules.table.permissions') }}</th>
              <th class="px-6 py-5 text-center">{{ t('admin.modules.table.sort_order') }}</th>
              <th class="px-8 py-5 text-right">{{ t('admin.modules.table.actions') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-50 dark:divide-slate-800/50">
            <tr v-if="loading" v-for="i in 5" :key="i" class="animate-pulse">
              <td colspan="5" class="px-8 py-6">
                <div class="h-8 bg-slate-100 dark:bg-slate-800 rounded-xl w-full"></div>
              </td>
            </tr>
            <tr v-else v-for="mod in filteredModules" :key="mod.id" class="group hover:bg-slate-50 dark:hover:bg-white/5 transition-colors">
              <td class="px-8 py-5">
                <div class="flex items-center gap-4 text-left">
                  <div class="w-10 h-10 rounded-xl bg-slate-100 dark:bg-slate-800 flex items-center justify-center text-primary-500 shadow-sm">
                    <component :is="getIcon(mod.icon)" class="w-5 h-5" />
                  </div>
                  <div>
                    <p class="text-[11px] font-black text-slate-900 dark:text-white uppercase tracking-tight">{{ mod.name }}</p>
                    <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest">{{ mod.id }}</p>
                  </div>
                </div>
              </td>
              <td class="px-6 py-5">
                <p class="text-[10px] font-black text-primary-500 uppercase tracking-widest">{{ mod.parent_id || 'ROOT' }}</p>
                <p class="text-[10px] font-bold text-slate-500 dark:text-slate-400 mt-0.5">{{ mod.path || '-' }}</p>
              </td>
              <td class="px-6 py-5">
                <div class="flex flex-wrap gap-1.5">
                  <span v-for="action in mod.allowed_actions" :key="action" class="px-2 py-0.5 bg-slate-100 dark:bg-slate-800 text-slate-500 dark:text-slate-400 rounded-md text-[8px] font-black uppercase tracking-tighter">
                    {{ action }}
                  </span>
                </div>
              </td>
              <td class="px-6 py-5 text-center">
                <span class="text-[11px] font-black text-[#1E3A5F] dark:text-slate-300 bg-slate-100 dark:bg-slate-800 px-3 py-1 rounded-lg">
                  {{ mod.sort_order }}
                </span>
              </td>
              <td class="px-8 py-5 text-right">
                <div class="flex items-center justify-end gap-2">
                  <button @click="openModal(mod)" class="p-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl text-slate-400 hover:text-blue-500 transition-colors">
                    <LucideEdit3 class="w-4 h-4" />
                  </button>
                  <button @click="deleteModule(mod.id)" class="p-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl text-slate-400 hover:text-red-500 transition-colors">
                    <LucideTrash2 class="w-4 h-4" />
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Module Modal (Centered) -->
    <Transition name="modal-scale">
      <div v-if="showModal" class="fixed inset-0 z-[110] flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-slate-950/60 backdrop-blur-md" @click="showModal = false"></div>
        <div class="relative w-full max-w-2xl bg-white dark:bg-slate-900 max-h-[90vh] shadow-2xl rounded-lg flex flex-col overflow-hidden border border-white/20 dark:border-slate-800">
          <header class="p-8 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between bg-slate-50/50 dark:bg-slate-800/30">
            <div class="flex items-center gap-4">
              <div class="w-12 h-12 rounded-2xl bg-primary-500/10 flex items-center justify-center text-primary-500">
                <LucideLayers v-if="!isEdit" class="w-6 h-6" />
                <LucideEdit3 v-else class="w-6 h-6" />
              </div>
              <div>
                <h3 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ isEdit ? t('admin.modules.modal.title_edit') : t('admin.modules.modal.title_add') }}</h3>
                <p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{{ t('admin.modules.modal.subtitle') }}</p>
              </div>
            </div>
            <button @click="showModal = false" class="p-3 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-2xl transition-all text-slate-400 hover:text-red-500 active:scale-90">
              <LucideX class="w-6 h-6" />
            </button>
          </header>

          <form @submit.prevent="saveModule" class="flex-1 overflow-y-auto p-8 space-y-4 custom-scrollbar">
            <!-- Core Info -->
            <div class="space-y-6">
              <div class="flex items-center gap-3">
                <span class="w-1.5 h-6 bg-primary-500 rounded-full"></span>
                <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ t('admin.modules.modal.section_core') }}</h4>
              </div>
              <div class="grid grid-cols-2 gap-8">
                <div class="space-y-2 text-left">
                  <label class="text-[10px] font-black text-slate-500 uppercase px-1">{{ t('admin.modules.modal.label_id') }}</label>
                  <input v-model="formData.id" type="text" :disabled="isEdit" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all disabled:opacity-50" placeholder="e.g. dashboard_v2" required />
                </div>
                <div class="space-y-2 text-left">
                  <label class="text-[10px] font-black text-slate-500 uppercase px-1">{{ t('admin.modules.modal.label_name') }}</label>
                  <input v-model="formData.name" type="text" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all" placeholder="Nama yang muncul di menu" required />
                </div>
              </div>
              <div class="grid grid-cols-2 gap-8">
                <div class="space-y-2 text-left">
                  <label class="text-[10px] font-black text-slate-500 uppercase px-1">{{ t('admin.modules.modal.label_parent') }}</label>
                  <select v-model="formData.parent_id" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 appearance-none bg-[url('data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20fill%3D%22none%22%20viewBox%3D%220%200%2024%2024%22%20stroke%3D%22%2394a3b8%22%3E%3Cpath%20stroke-linecap%3D%22round%22%20stroke-linejoin%3D%22round%22%20stroke-width%3D%222%22%20d%3D%22M19%209l-7%207-7-7%22%2F%3E%3C%2Fsvg%3E')] bg-[length:1.25rem_1.25rem] bg-[right_1.5rem_center] bg-no-repeat">
                    <option value="">{{ t('admin.modules.modal.no_parent') }}</option>
                    <option v-for="m in modules.filter(mod => !mod.parent_id)" :key="m.id" :value="m.id">
                      {{ m.name }}
                    </option>
                  </select>
                </div>
                <div class="space-y-2 text-left">
                  <label class="text-[10px] font-black text-slate-500 uppercase px-1">{{ t('admin.modules.modal.label_sort') }}</label>
                  <input v-model.number="formData.sort_order" type="number" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10" required />
                </div>
              </div>
            </div>

            <!-- Permissions Schema -->
            <div class="space-y-6">
              <div class="flex items-center gap-3">
                <span class="w-1.5 h-6 bg-primary-500 rounded-full"></span>
                <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ t('admin.modules.modal.section_perms') }}</h4>
              </div>
              <div class="grid grid-cols-2 sm:grid-cols-4 gap-4">
                <div v-for="action in allActions" :key="action" 
                  @click="toggleAction(action)"
                  class="group flex flex-col items-center justify-center gap-3 p-4 bg-slate-50 dark:bg-slate-800 border-2 rounded-3xl cursor-pointer transition-all active:scale-95"
                  :class="(formData.allowed_actions || []).includes(action) ? 'border-primary-500 bg-primary-500/5 shadow-lg shadow-primary-500/5' : 'border-transparent opacity-60 hover:opacity-100 hover:bg-slate-100 dark:hover:bg-slate-700'">
                  <div class="w-8 h-8 rounded-full border-2 flex items-center justify-center transition-all duration-300"
                    :class="(formData.allowed_actions || []).includes(action) ? 'bg-primary-500 border-primary-500 text-white' : 'border-slate-300 dark:border-slate-600 text-transparent group-hover:border-slate-400'">
                    <LucideCheck class="w-4 h-4" />
                  </div>
                  <span class="text-[9px] font-black uppercase tracking-widest" :class="(formData.allowed_actions || []).includes(action) ? 'text-primary-500' : 'text-slate-500'">{{ action }}</span>
                </div>
              </div>
            </div>

            <!-- Path & UI -->
            <div class="space-y-6">
              <div class="flex items-center gap-3">
                <span class="w-1.5 h-6 bg-primary-500 rounded-full"></span>
                <h4 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ t('admin.modules.modal.section_nav') }}</h4>
              </div>
              <div class="grid grid-cols-2 gap-8">
                <div class="space-y-2 text-left">
                  <label class="text-[10px] font-black text-slate-500 uppercase px-1">{{ t('admin.modules.modal.label_path') }}</label>
                  <input v-model="formData.path" type="text" class="w-full px-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10 transition-all" placeholder="/admin/example" />
                </div>
                <div class="space-y-2 text-left">
                  <label class="text-[10px] font-black text-slate-500 uppercase px-1">{{ t('admin.modules.modal.label_icon') }}</label>
                  <div class="relative">
                    <div class="absolute left-6 top-1/2 -translate-y-1/2 w-5 h-5 flex items-center justify-center text-primary-500">
                      <component :is="getIcon(formData.icon)" class="w-full h-full" />
                    </div>
                    <input v-model="formData.icon" type="text" class="w-full pl-14 pr-6 py-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-2xl text-xs font-bold outline-none focus:ring-4 focus:ring-primary-500/10" placeholder="LucideLayers" />
                  </div>
                </div>
              </div>
            </div>
          </form>

          <footer class="p-8 border-t border-slate-100 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-900/50 flex items-center justify-end gap-4">
            <button @click="showModal = false" class="px-8 py-4 text-[10px] font-black text-slate-400 hover:text-slate-600 uppercase tracking-widest transition-all active:scale-95">{{ t('admin.modules.modal.btn_cancel') }}</button>
            <button @click="saveModule" :disabled="saving" class="px-10 py-4 bg-[#1E3A5F] text-white rounded-2xl text-[10px] font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all flex items-center gap-3 active:scale-95 disabled:opacity-50">
              <LucideSave v-if="!saving" class="w-4 h-4" />
              <LucideLoader2 v-else class="w-4 h-4 animate-spin" />
              {{ saving ? t('admin.modules.modal.saving') : t('admin.modules.modal.btn_save') }}
            </button>
          </footer>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRuntimeConfig } from '#app'
import * as LucideIcons from 'lucide-vue-next'
import { 
  LucideLayers, LucideSearch, LucidePlusCircle, LucideEdit3, 
  LucideTrash2, LucideX, LucideSave, LucideCheck, LucideLoader2
} from 'lucide-vue-next'
import { useI18n } from 'vue-i18n'
import { useApi } from '@/composables/useApi'

const { t } = useI18n()
const { $api } = useApi()
const config = useRuntimeConfig()

const loading = ref(false)
const saving = ref(false)
const showModal = ref(false)
const isEdit = ref(false)

const allActions = ['VIEW', 'CREATE', 'EDIT', 'DELETE', 'EXPORT', 'DOWNLOAD', 'PRINT']
const modules = ref([])
const categories = ref([])

const filters = ref({
  search: '',
  category: ''
})

const formData = ref({
  id: '',
  name: '',
  category: '',
  path: '',
  icon: 'LucideLayers',
  allowed_actions: ['VIEW'],
  sort_order: 0,
  parent_id: ''
})

const getIcon = (name) => {
  return LucideIcons[name] || LucideIcons.LucideLayers
}

const filteredModules = computed(() => {
  return modules.value.filter(m => {
    const matchSearch = !filters.value.search || 
      m.id.toLowerCase().includes(filters.value.search.toLowerCase()) ||
      m.name.toLowerCase().includes(filters.value.search.toLowerCase()) ||
      m.category.toLowerCase().includes(filters.value.search.toLowerCase())
    const matchCat = !filters.value.category || m.category === filters.value.category
    return matchSearch && matchCat
  })
})

const fetchModules = async () => {
  loading.value = true
  try {
    const res = await $api(`${config.public.apiBase}/master/modules`)
    if (res && res.data) {
      modules.value = res.data
      const cats = [...new Set(res.data.map(m => m.category))]
      categories.value = cats.sort()
    }
  } catch (err) {
    console.error('Failed to fetch modules:', err)
  } finally {
    loading.value = false
  }
}

const openModal = (mod = null) => {
  if (mod) {
    isEdit.value = true
    formData.value = { 
      ...mod, 
      path: mod.path || '',
      icon: mod.icon || 'LucideLayers',
      parent_id: mod.parent_id || '',
      allowed_actions: mod.allowed_actions || []
    }
  } else {
    isEdit.value = false
    formData.value = {
      id: '',
      name: '',
      category: '',
      path: '',
      icon: 'LucideLayers',
      allowed_actions: ['VIEW'],
      sort_order: (modules.value.length + 1) * 10,
      parent_id: ''
    }
  }
  showModal.value = true
}

const toggleAction = (action) => {
  const index = formData.value.allowed_actions.indexOf(action)
  if (index === -1) {
    formData.value.allowed_actions.push(action)
  } else {
    formData.value.allowed_actions.splice(index, 1)
  }
}

const saveModule = async () => {
  if (!formData.value.id || !formData.value.name) return
  
  saving.value = true
  try {
    const url = isEdit.value 
      ? `${config.public.apiBase}/master/modules/${formData.value.id}`
      : `${config.public.apiBase}/master/modules`
    
    await $api(url, {
      method: isEdit.value ? 'PUT' : 'POST',
      body: formData.value
    })
    
    await fetchModules()
    showModal.value = false
  } catch (err) {
    console.error('Failed to save module:', err)
    alert('Failed to save module: ' + err.message)
  } finally {
    saving.value = false
  }
}

const deleteModule = async (id) => {
  if (confirm(`Apakah Anda yakin ingin menghapus modul "${id}"? Hal ini akan menghapus semua izin terkait!`)) {
    try {
      await $api(`${config.public.apiBase}/master/modules/${id}`, {
        method: 'DELETE'
      })
      await fetchModules()
    } catch (err) {
      console.error('Failed to delete module:', err)
      alert('Failed to delete module: ' + err.message)
    }
  }
}

onMounted(fetchModules)

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

.modal-scale-enter-active, .modal-scale-leave-active {
  transition: all 0.4s cubic-bezier(0.34, 1.56, 0.64, 1);
}
.modal-scale-enter-from, .modal-scale-leave-to {
  opacity: 0;
  transform: scale(0.9) translateY(20px);
}
</style>
