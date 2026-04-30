<template>
  <div ref="containerRef" class="relative w-full" v-on-click-outside="close">
    <div 
      @click="toggle"
      class="w-full bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl px-4 py-2 text-xs font-bold cursor-pointer flex items-center justify-between group hover:border-blue-500/50 transition-all"
      :class="{ 'ring-2 ring-blue-500/20 border-blue-500': isOpen }"
    >
      <span v-if="selectedLabel" class="text-slate-700 dark:text-slate-200 truncate">{{ selectedLabel }}</span>
      <span v-else class="text-slate-400 uppercase tracking-widest text-[10px]">{{ placeholder }}</span>
      <LucideChevronDown class="w-4 h-4 text-slate-400 group-hover:text-blue-500 transition-colors" :class="{ 'rotate-180': isOpen }" />
    </div>

    <Teleport to="body">
      <Transition
        enter-active-class="transition duration-200 ease-out"
        enter-from-class="transform scale-95 opacity-0"
        enter-to-class="transform scale-100 opacity-100"
        leave-active-class="transition duration-150 ease-in"
        leave-from-class="transform scale-100 opacity-100"
        leave-to-class="transform scale-95 opacity-0"
      >
        <div 
          v-if="isOpen" 
          :style="dropdownStyle"
          class="fixed bg-white dark:bg-slate-900 rounded-2xl shadow-2xl border border-slate-200 dark:border-slate-800 z-[9999] overflow-hidden flex flex-col max-h-72 pointer-events-auto"
        >
          <!-- Search Input -->
          <div class="p-3 border-b border-slate-50 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-900/50">
            <div class="relative group">
              <LucideSearch class="absolute left-3 top-1/2 -translate-y-1/2 w-3.5 h-3.5 text-slate-400 group-focus-within:text-blue-500 transition-colors" />
              <input 
                ref="searchInput"
                type="text" 
                v-model="searchQuery"
                placeholder="Cari..." 
                class="w-full bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-lg pl-9 pr-4 py-2 text-[11px] font-bold outline-none focus:ring-2 focus:ring-blue-500/20 transition-all"
                @keydown.esc="close"
              />
            </div>
          </div>

          <!-- Options List -->
          <div class="overflow-y-auto flex-grow custom-scrollbar py-2">
            <div 
              v-for="option in displayedOptions" 
              :key="option.id"
              @click="select(option)"
              class="px-4 py-2.5 text-[11px] font-bold text-slate-600 dark:text-slate-300 hover:bg-blue-50 dark:hover:bg-blue-900/20 hover:text-blue-600 dark:hover:text-blue-400 cursor-pointer transition-colors flex items-center justify-between group"
              :class="{ 'bg-blue-50/50 dark:bg-blue-900/10 text-blue-600': modelValue === option.id }"
            >
              <span class="truncate">{{ option.name }} <span v-if="option.subtext" class="text-[9px] text-slate-400 font-normal ml-1">({{ option.subtext }})</span></span>
              <LucideCheck v-if="modelValue === option.id" class="w-3.5 h-3.5" />
            </div>

            <div v-if="displayedOptions.length === 0" class="px-4 py-10 text-center">
              <p class="text-[10px] font-black uppercase tracking-widest text-slate-400 opacity-50">Tidak ada data</p>
            </div>
            
            <!-- Limit indicator if not searching -->
            <div v-if="!searchQuery && options.length > 5 && displayedOptions.length === 5" class="px-4 py-2 border-t border-slate-50 dark:border-slate-800 mt-2">
              <p class="text-[9px] font-black text-slate-400 uppercase tracking-tighter italic text-center">Menampilkan 5 data awal. Gunakan pencarian untuk data lainnya.</p>
            </div>
          </div>
        </div>
      </Transition>
    </Teleport>
  </div>
</template>

<script setup>
import { ref, computed, nextTick, onMounted, onUnmounted, watch } from 'vue'
import { vOnClickOutside } from '@vueuse/components'
import { LucideChevronDown, LucideSearch, LucideCheck } from 'lucide-vue-next'

const props = defineProps({
  modelValue: [String, Number, null],
  options: {
    type: Array,
    default: () => []
  },
  placeholder: {
    type: String,
    default: 'Pilih data...'
  }
})

const emit = defineEmits(['update:modelValue', 'change'])

const isOpen = ref(false)
const searchQuery = ref('')
const searchInput = ref(null)
const containerRef = ref(null)
const dropdownStyle = ref({})

const selectedLabel = computed(() => {
  const option = props.options.find(o => o.id === props.modelValue)
  return option ? option.name : ''
})

const filteredOptions = computed(() => {
  if (!searchQuery.value) return props.options
  const q = searchQuery.value.toLowerCase()
  return props.options.filter(o => 
    (o.name || '').toLowerCase().includes(q) || 
    (o.subtext || '').toLowerCase().includes(q)
  )
})

const displayedOptions = computed(() => {
  if (!searchQuery.value) {
    return filteredOptions.value.slice(0, 5)
  }
  return filteredOptions.value
})

const updatePosition = () => {
  if (!containerRef.value || !isOpen.value) return
  
  const rect = containerRef.value.getBoundingClientRect()
  
  dropdownStyle.value = {
    position: 'fixed',
    top: `${rect.bottom + 8}px`,
    left: `${rect.left}px`,
    width: `${rect.width}px`,
    transformOrigin: 'top'
  }
}

const toggle = () => {
  isOpen.value = !isOpen.value
  if (isOpen.value) {
    nextTick(() => {
      updatePosition()
      searchInput.value?.focus()
    })
  }
}

const close = () => {
  isOpen.value = false
  searchQuery.value = ''
}

const select = (option) => {
  emit('update:modelValue', option.id)
  emit('change', option)
  close()
}

// Handle scroll and resize to keep dropdown aligned
onMounted(() => {
  window.addEventListener('scroll', updatePosition, true)
  window.addEventListener('resize', updatePosition)
})

onUnmounted(() => {
  window.removeEventListener('scroll', updatePosition, true)
  window.removeEventListener('resize', updatePosition)
})

// Update position if window size changes or if parent scrolls
watch(isOpen, (newVal) => {
  if (newVal) {
    nextTick(updatePosition)
  }
})
</script>

<style scoped>
@reference "../assets/css/main.css";

.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  @apply bg-transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  @apply bg-slate-200 dark:bg-slate-800 rounded-full;
}
</style>
