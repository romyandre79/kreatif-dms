<template>
  <div class="relative w-full rounded-lg overflow-hidden border border-slate-200 dark:border-slate-700 bg-slate-50 dark:bg-slate-900 shadow-inner group">
    <!-- Image Wrapper -->
    <div 
      class="relative w-full flex items-center justify-center min-h-[300px] cursor-crosshair overflow-hidden"
      @click="handleMapClick"
    >
      <template v-if="floorPlanUrl">
        <img 
          :src="floorPlanUrl" 
          alt="Floor Plan" 
          class="w-full h-auto max-h-[600px] object-contain transition-transform duration-300" 
          ref="mapImageRef"
        />
        
        <!-- Target Reticle Overlay on Hover -->
        <div v-if="isEditing" class="absolute inset-0 pointer-events-none opacity-0 group-hover:opacity-100 transition-opacity bg-black/10">
           <div class="absolute inset-0 flex items-center justify-center">
             <span class="bg-primary-500 text-white text-[10px] font-bold px-2 py-1 rounded shadow-lg uppercase tracking-widest">
                Klik untuk set posisi
             </span>
           </div>
        </div>

        <!-- Render Existing Racks -->
        <div 
          v-for="rack in activeRacks" 
          :key="rack.id"
          class="absolute w-4 h-4 -ml-2 -mt-2 bg-primary-500 rounded-full shadow-lg ring-2 ring-white dark:ring-slate-800 cursor-pointer hover:scale-125 transition-transform z-10"
          :style="{ left: `${rack.map_pos_x}%`, top: `${rack.map_pos_y}%` }"
          @click.stop="emit('rack-click', rack)"
        >
          <div v-if="rack.id === activeRackId" class="absolute inset-0 bg-primary-500 rounded-full animate-ping opacity-75"></div>
          
          <!-- Tooltip -->
          <div class="absolute bottom-full left-1/2 -translate-x-1/2 mb-2 opacity-0 hover:opacity-100 bg-slate-800 text-white text-[10px] font-bold px-2 py-1 rounded whitespace-nowrap pointer-events-none transition-opacity">
            {{ rack.name || 'Rak' }}
          </div>
        </div>
      </template>

      <!-- Empty State / Upload Prompt -->
      <div v-else class="text-center p-10 flex flex-col items-center gap-4">
        <div class="w-16 h-16 rounded-2xl bg-slate-100 dark:bg-slate-800 flex items-center justify-center text-slate-400">
          <LucideImagePlus class="w-8 h-8" />
        </div>
        <div class="space-y-1">
          <p class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">Denah Belum Tersedia</p>
          <p class="text-[11px] text-slate-500 font-medium">Unggah gambar layout/denah gudang untuk mengaktifkan pemetaan internal.</p>
        </div>
        <button 
          v-if="canEditLayout"
          @click.stop="triggerUpload" 
          class="mt-2 px-6 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest hover:bg-[#152943] transition-all flex items-center gap-2"
        >
          <LucideUpload class="w-4 h-4" />
          Unggah Denah
        </button>
        <input type="file" ref="fileInputRef" class="hidden" accept="image/*" @change="handleFileUpload" />
      </div>
    </div>

    <!-- Toolbar overlay -->
    <div v-if="floorPlanUrl && canEditLayout" class="absolute bottom-4 right-4 flex gap-2">
      <button 
        @click="triggerUpload" 
        class="p-3 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl shadow-lg text-slate-600 dark:text-slate-300 hover:text-primary-500 transition-colors"
        title="Ganti Denah"
      >
        <LucideImage class="w-4 h-4" />
      </button>
      <input type="file" ref="fileInputRef" class="hidden" accept="image/*" @change="handleFileUpload" />
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { LucideImagePlus, LucideUpload, LucideImage } from 'lucide-vue-next'

const props = defineProps({
  floorPlanUrl: { type: String, default: null },
  racks: { type: Array, default: () => [] },
  activeRackId: { type: [String, Number], default: null },
  isEditing: { type: Boolean, default: false },
  canEditLayout: { type: Boolean, default: false }
})

const emit = defineEmits(['update:floorPlan', 'map-click', 'rack-click'])

const mapImageRef = ref(null)
const fileInputRef = ref(null)

// Filter out racks that don't have coordinates yet (or are invalid)
const activeRacks = computed(() => {
  return props.racks.filter(r => r.map_pos_x != null && r.map_pos_y != null)
})

const handleMapClick = (e) => {
  if (!props.isEditing || !props.floorPlanUrl || !mapImageRef.value) return;

  // Calculate percentage relative to the image element
  const rect = mapImageRef.value.getBoundingClientRect()
  
  // Calculate relative position 
  let x = e.clientX - rect.left
  let y = e.clientY - rect.top

  // Clamp values to ensure they are within the image bounds
  x = Math.max(0, Math.min(x, rect.width))
  y = Math.max(0, Math.min(y, rect.height))

  // Convert to percentages (0 to 100)
  const percentX = (x / rect.width) * 100
  const percentY = (y / rect.height) * 100

  // Emit with 2 decimal places precision
  emit('map-click', { 
    x: Number(percentX.toFixed(2)), 
    y: Number(percentY.toFixed(2)) 
  })
}

const triggerUpload = () => {
  if (fileInputRef.value) {
    fileInputRef.value.click()
  }
}

const handleFileUpload = (e) => {
  const file = e.target.files?.[0]
  if (!file) return

  // For this prototype, we'll convert the image to a Base64 data URL
  // In production, this should upload to the server and return a URL
  const reader = new FileReader()
  reader.onload = (event) => {
    emit('update:floorPlan', event.target.result)
  }
  reader.readAsDataURL(file)
  
  // Reset input
  e.target.value = ''
}
</script>
