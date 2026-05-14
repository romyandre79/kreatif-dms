<template>
  <div class="space-y-4">
    <div class="relative border-2 border-dashed border-slate-200 dark:border-slate-800 rounded-2xl bg-slate-50 dark:bg-slate-900/50 overflow-hidden group">
      <canvas 
        ref="canvas"
        @mousedown="startDrawing"
        @mousemove="draw"
        @mouseup="stopDrawing"
        @mouseleave="stopDrawing"
        @touchstart.passive="startDrawing"
        @touchmove.prevent="draw"
        @touchend.passive="stopDrawing"
        class="w-full h-48 cursor-crosshair touch-none"
      ></canvas>
      
      <div v-if="isEmpty" class="absolute inset-0 flex flex-col items-center justify-center pointer-events-none text-slate-400 group-hover:text-slate-500 transition-colors">
        <LucidePenTool class="w-8 h-8 mb-2" />
        <p class="text-xs font-bold uppercase tracking-widest">Draw your signature here</p>
      </div>

      <button 
        @click="clear"
        type="button"
        class="absolute bottom-4 right-4 p-2 bg-white dark:bg-slate-800 rounded-lg shadow-sm border border-slate-100 dark:border-slate-700 text-slate-400 hover:text-red-500 transition-all"
        title="Clear"
      >
        <LucideTrash2 class="w-4 h-4" />
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { LucidePenTool, LucideTrash2 } from 'lucide-vue-next'

const canvas = ref(null)
const ctx = ref(null)
const isDrawing = ref(false)
const isEmpty = ref(true)

const props = defineProps({
  modelValue: String
})

const emit = defineEmits(['update:modelValue'])

let resizeObserver = null

const initCanvas = () => {
  const c = canvas.value
  if (!c) return
  
  ctx.value = c.getContext('2d')
  
  // Set canvas resolution to match display size
  const rect = c.getBoundingClientRect()
  c.width = rect.width
  c.height = rect.height
  
  ctx.value.strokeStyle = '#1E3A5F'
  ctx.value.lineWidth = 2
  ctx.value.lineJoin = 'round'
  ctx.value.lineCap = 'round'

  // Restore if modelValue exists
  if (props.modelValue && props.modelValue.startsWith('data:image')) {
    const img = new Image()
    img.onload = () => {
      ctx.value.drawImage(img, 0, 0)
      isEmpty.value = false
    }
    img.src = props.modelValue
  }
}

onMounted(() => {
  initCanvas()
  
  // Re-init on resize (important for modals/layout changes)
  resizeObserver = new ResizeObserver(() => {
    // We only re-init if the canvas is actually visible and has size
    if (canvas.value && canvas.value.offsetWidth > 0) {
      initCanvas()
    }
  })
  if (canvas.value) {
    resizeObserver.observe(canvas.value)
  }
})

onUnmounted(() => {
  if (resizeObserver) {
    resizeObserver.disconnect()
  }
})

const startDrawing = (e) => {
  isDrawing.value = true
  isEmpty.value = false
  const pos = getPos(e)
  ctx.value.beginPath()
  ctx.value.moveTo(pos.x, pos.y)
}

const draw = (e) => {
  if (!isDrawing.value) return
  const pos = getPos(e)
  ctx.value.lineTo(pos.x, pos.y)
  ctx.value.stroke()
}

const stopDrawing = () => {
  if (!isDrawing.value) return
  isDrawing.value = false
  save()
}

const getPos = (e) => {
  const c = canvas.value
  const rect = c.getBoundingClientRect()
  const clientX = e.touches ? e.touches[0].clientX : e.clientX
  const clientY = e.touches ? e.touches[0].clientY : e.clientY
  return {
    x: clientX - rect.left,
    y: clientY - rect.top
  }
}

const clear = () => {
  const c = canvas.value
  ctx.value.clearRect(0, 0, c.width, c.height)
  isEmpty.value = true
  emit('update:modelValue', null)
}

const save = () => {
  const dataURL = canvas.value.toDataURL('image/png')
  emit('update:modelValue', dataURL)
}

defineExpose({ clear, save })
</script>
