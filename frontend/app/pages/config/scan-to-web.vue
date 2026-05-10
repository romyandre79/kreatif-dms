<template>
  <div class="space-y-10 pb-20" v-motion-fade>
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div>
        <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">
          Scan to Web Launcher
        </h1>
        <p class="text-sm text-slate-500 mt-1 uppercase font-bold tracking-widest text-[10px]">
          Direct scanning from local devices to document management
        </p>
      </div>
      
      <div class="flex items-center gap-4">
        <div :class="bridgeStatus === 'connected' ? 'bg-emerald-500/10 text-emerald-500 border-emerald-500/20' : 'bg-red-500/10 text-red-500 border-red-500/20'"
             class="px-4 py-2 rounded-xl border text-[10px] font-black uppercase tracking-widest flex items-center gap-2">
          <div :class="bridgeStatus === 'connected' ? 'bg-emerald-500 animate-pulse' : 'bg-red-500'" class="w-2 h-2 rounded-full"></div>
          Bridge: {{ bridgeStatus }}
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Left Column: Controls -->
      <div class="lg:col-span-4 space-y-8">
        <div class="glass p-8 rounded-lg space-y-8 border border-white/10 shadow-sm bg-white/50 dark:bg-slate-900/50 backdrop-blur-xl">
          <div class="flex items-center gap-4 border-b border-slate-100 dark:border-slate-800 pb-6">
            <div class="w-10 h-10 rounded-xl bg-primary-500/10 flex items-center justify-center">
              <LucidePrinter class="w-5 h-5 text-primary-500" />
            </div>
            <h3 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">Scanner Source</h3>
          </div>

          <div class="space-y-6">
            <!-- Driver Type -->
            <div class="space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Driver Type</label>
              <div class="grid grid-cols-2 gap-4">
                <button v-for="t in [{id:'local', label:'Local (Bridge)'}, {id:'network', label:'Network'}]" :key="t.id"
                        @click="driverType = t.id"
                        :class="driverType === t.id ? 'bg-primary-500 text-white shadow-lg shadow-primary-500/20' : 'bg-slate-100 dark:bg-slate-800 text-slate-500'"
                        class="py-3 rounded-xl text-[10px] font-black uppercase tracking-widest transition-all">
                  {{ t.label }}
                </button>
              </div>
            </div>

            <div class="space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Select Scanner Device</label>
              <select v-model="selectedScannerID" 
                      class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-medium focus:ring-2 focus:ring-primary-500 transition-all appearance-none">
                <option value="" disabled>Choose a registered device...</option>
                <option v-for="s in registeredScanners" :key="s.id" :value="s.endpoint">
                  {{ s.name }} ({{ s.service_type === 'SCANNER_LOCAL' ? 'Local' : 'Network' }})
                </option>
              </select>
              <div v-if="driverType === 'local'" class="flex items-center gap-2 mt-2">
                <div :class="localScanners.some(ls => ls.id === selectedScannerID) ? 'bg-emerald-500' : 'bg-slate-300'" class="w-2 h-2 rounded-full"></div>
                <span class="text-[9px] font-bold uppercase tracking-widest text-slate-400">
                  {{ localScanners.some(ls => ls.id === selectedScannerID) ? 'Hardware Detected' : 'Hardware Not Found' }}
                </span>
              </div>
            </div>

            <!-- Scan Settings -->
            <div class="pt-6 border-t border-slate-100 dark:border-slate-800 space-y-6">
              <div class="grid grid-cols-2 gap-4">
                <div class="space-y-2">
                  <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Resolution (DPI)</label>
                  <select v-model="scanSettings.dpi" class="w-full px-4 py-3 bg-slate-100 dark:bg-slate-800 border-none rounded-xl text-xs font-bold">
                    <option :value="150">150 DPI</option>
                    <option :value="300">300 DPI</option>
                    <option :value="600">600 DPI</option>
                  </select>
                </div>
                <div class="space-y-2">
                  <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Color Mode</label>
                  <select v-model="scanSettings.color_mode" class="w-full px-4 py-3 bg-slate-100 dark:bg-slate-800 border-none rounded-xl text-xs font-bold">
                    <option value="color">Full Color</option>
                    <option value="grayscale">Grayscale</option>
                    <option value="bw">Black & White</option>
                  </select>
                </div>
              </div>

              <!-- Auto Mode Toggle -->
              <div class="flex items-center justify-between p-4 bg-primary-500/5 rounded-2xl border border-primary-500/10">
                <div class="flex items-center gap-3">
                  <LucideZap class="w-4 h-4 text-primary-500" />
                  <div>
                    <p class="text-[10px] font-black text-[#1E3A5F] dark:text-white uppercase tracking-widest">Auto-Read Mode</p>
                    <p class="text-[8px] text-slate-400 font-bold uppercase tracking-widest">Auto upload & OCR</p>
                  </div>
                </div>
                <button @click="autoMode = !autoMode" 
                        :class="autoMode ? 'bg-primary-500' : 'bg-slate-300 dark:bg-slate-700'"
                        class="w-10 h-5 rounded-full relative transition-all">
                  <div :class="autoMode ? 'translate-x-5' : 'translate-x-1'" class="absolute top-1 w-3 h-3 bg-white rounded-full transition-all"></div>
                </button>
              </div>
            </div>

            <!-- Action Button -->
            <button @click="startScan" 
                    :disabled="scanning || (driverType === 'local' && !selectedScannerID) || (driverType === 'network' && !networkEndpoint)"
                    class="w-full py-5 bg-primary-500 text-white rounded-lg text-sm font-black uppercase tracking-[0.2em] shadow-2xl shadow-primary-500/40 hover:scale-[1.02] active:scale-[0.98] transition-all disabled:opacity-50 disabled:hover:scale-100 flex items-center justify-center gap-4">
              <LucideLoader2 v-if="scanning" class="w-5 h-5 animate-spin" />
              <LucideZap v-else class="w-5 h-5" />
              {{ scanning ? 'Scanning...' : 'Start Scanning' }}
            </button>
          </div>
        </div>
      </div>

      <!-- Right Column: Preview & History -->
      <div class="lg:col-span-8 space-y-8">
        <!-- Preview Window -->
        <div class="glass rounded-[3rem] overflow-hidden border border-white/10 shadow-2xl bg-white dark:bg-slate-900 min-h-[500px] flex flex-col">
          <div class="p-6 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between">
            <div class="flex items-center gap-3">
              <div class="w-2 h-2 rounded-full bg-primary-500"></div>
              <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Document Preview</h3>
            </div>
            <div v-if="lastResult" class="flex gap-4">
              <button @click="uploadResult" class="px-4 py-2 bg-emerald-500 text-white rounded-xl text-[10px] font-black uppercase tracking-widest shadow-lg shadow-emerald-500/20">
                Process Document
              </button>
              <button @click="lastResult = null" class="px-4 py-2 bg-slate-100 dark:bg-slate-800 text-slate-500 rounded-xl text-[10px] font-black uppercase tracking-widest">
                Discard
              </button>
            </div>
          </div>
          
          <div class="flex-1 bg-slate-50 dark:bg-[#05080F] flex items-center justify-center p-8">
            <div v-if="lastResult" class="max-w-full max-h-full shadow-2xl rounded-lg overflow-hidden border-4 border-white dark:border-slate-800">
              <img :src="`data:image/${lastResult.format};base64,${lastResult.data}`" class="max-w-full h-auto" />
            </div>
            <div v-else-if="scanning" class="text-center space-y-4">
              <div class="w-20 h-20 border-4 border-primary-500/20 border-t-primary-500 rounded-full animate-spin mx-auto"></div>
              <p class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] animate-pulse">Capturing image from hardware...</p>
            </div>
            <div v-else class="text-center space-y-6 opacity-20">
              <LucideCameraOff class="w-24 h-24 mx-auto text-slate-400" />
              <p class="text-xs font-black text-slate-400 uppercase tracking-widest">No document scanned yet</p>
            </div>
          </div>
        </div>

        <!-- Recent Activity from this device -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div v-for="i in 2" :key="i" class="glass p-6 rounded-3xl border border-white/10 bg-white/30 dark:bg-slate-900/30">
            <div class="flex items-center gap-4">
              <div class="w-12 h-12 rounded-2xl bg-slate-100 dark:bg-slate-800 flex items-center justify-center">
                <LucideHistory class="w-6 h-6 text-slate-400" />
              </div>
              <div>
                <p class="text-[10px] font-black text-slate-500 uppercase tracking-widest">Recent Activity</p>
                <p class="text-xs font-bold text-[#1E3A5F] dark:text-white">Waiting for first scan...</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watch } from 'vue'
import { 
  LucidePrinter, LucideZap, LucideLoader2, LucideHistory, 
  LucideCameraOff, LucideSettings2 
} from 'lucide-vue-next'
import { useApi } from '@/composables/useApi'

const { $api } = useApi()
const config = useRuntimeConfig()

const bridgeStatus = ref('checking')
const driverType = ref('local')
const selectedScannerID = ref('')
const localScanners = ref([])
const registeredScanners = ref([])
const networkEndpoint = ref('')
const scanning = ref(false)
const lastResult = ref(null)
const autoMode = ref(false)

const scanSettings = ref({
  dpi: 300,
  color_mode: 'color',
  format: 'jpeg'
})

let bridgeCheckInterval = null

const fetchRegisteredScanners = async () => {
  try {
    const res = await $api(`${config.public.apiBase}/master/integration/status`)
    if (res && res.data) {
      registeredScanners.value = (res.data.nodes || []).filter(n => 
        n.service_type === 'SCANNER_LOCAL' || n.service_type === 'SCANNER_NETWORK'
      )
      
      // Auto select first if nothing selected
      if (registeredScanners.value.length > 0 && !selectedScannerID.value) {
        const first = registeredScanners.value[0]
        selectedScannerID.value = first.endpoint
        driverType.value = first.service_type === 'SCANNER_LOCAL' ? 'local' : 'network'
        if (driverType.value === 'network') networkEndpoint.value = first.endpoint
      }
    }
  } catch (err) {
    console.error('Failed to fetch registered scanners:', err)
  }
}

const checkBridge = async () => {
  try {
    const res = await fetch('http://localhost:7878/health')
    if (res.ok) {
      if (bridgeStatus.value !== 'connected') {
        bridgeStatus.value = 'connected'
        fetchLocalScanners()
      }
    } else {
      bridgeStatus.value = 'disconnected'
    }
  } catch (err) {
    bridgeStatus.value = 'disconnected'
  }
}

const fetchLocalScanners = async () => {
  if (bridgeStatus.value !== 'connected') return
  try {
    const res = await fetch('http://localhost:7878/scanners')
    const data = await res.json()
    localScanners.value = data
  } catch (err) {
    console.error('Failed to fetch scanners from bridge')
  }
}

// Watch for manual selection change to update driver type
watch(selectedScannerID, (newID) => {
  const found = registeredScanners.value.find(s => s.endpoint === newID)
  if (found) {
    driverType.value = found.service_type === 'SCANNER_LOCAL' ? 'local' : 'network'
    if (driverType.value === 'network') networkEndpoint.value = found.endpoint
  }
})

const startScan = async () => {
  scanning.value = true
  lastResult.value = null
  
  // Find the registered scanner node
  const scannerNode = registeredScanners.value.find(s => s.endpoint === selectedScannerID.value)
  if (!scannerNode) {
    alert('Scanner configuration not found')
    scanning.value = false
    return
  }

  // Use connection_string from config as the base URL (e.g. http://192.168.100.4:7878)
  let baseUrl = scannerNode.config?.connection_string || scannerNode.endpoint
  if (!baseUrl.startsWith('http')) {
    // If it doesn't look like a URL, it might be just an IP or we need a fallback
    if (scannerNode.service_type === 'SCANNER_LOCAL') {
      baseUrl = 'http://localhost:7878'
    } else {
      baseUrl = `http://${baseUrl}`
    }
  }
  
  if (!baseUrl.endsWith('/')) baseUrl += '/'
  const scanUrl = `${baseUrl}scan`
  
  // Device ID is stored in the endpoint column for local scanners
  const deviceID = scannerNode.endpoint

  try {
    const res = await fetch(scanUrl, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        scanner_id: deviceID,
        ...scanSettings.value
      })
    })

    if (!res.ok) throw new Error('Scan failed')
    
    const result = await res.json()
    lastResult.value = result

    // Auto Registration to Backend Integration Nodes
    registerInBackend(result)

    if (autoMode.value) {
      uploadResult()
    }
  } catch (err) {
    alert('Scanning error: ' + err.message)
  } finally {
    scanning.value = false
  }
}

const registerInBackend = async (scanResult) => {
  try {
    const scannerName = driverType.value === 'local' 
      ? localScanners.value.find(s => s.id === selectedScannerID.value)?.name || 'Local Scanner'
      : 'Network Scanner'

    await $api(`${config.public.apiBase}/master/scanners/register`, {
      method: 'POST',
      body: {
        name: scannerName,
        service_type: driverType.value === 'local' ? 'SCANNER_LOCAL' : 'SCANNER_NETWORK',
        endpoint: driverType.value === 'local' ? selectedScannerID.value : networkEndpoint.value,
        config: {
          last_scan_width: scanResult.width,
          last_scan_height: scanResult.height,
          resolution: scanSettings.value.dpi
        }
      }
    })
  } catch (err) {
    console.error('Failed to register scanner in backend', err)
  }
}

const uploadResult = async () => {
  if (!lastResult.value) return
  
  try {
    // Convert Base64 to Blob
    const byteCharacters = atob(lastResult.value.data)
    const byteNumbers = new Array(byteCharacters.length)
    for (let i = 0; i < byteCharacters.length; i++) {
      byteNumbers[i] = byteCharacters.charCodeAt(i)
    }
    const byteArray = new Uint8Array(byteNumbers)
    const blob = new Blob([byteArray], { type: `image/${lastResult.value.format}` })
    
    const formData = new FormData()
    formData.append('file', blob, `scan_${Date.now()}.${lastResult.value.format}`)
    formData.append('title', `Scan - ${new Date().toLocaleString()}`)
    formData.append('description', 'Scanned via Web Launcher')

    await $api(`${config.public.apiBase}/documents`, {
      method: 'POST',
      body: formData
    })

    console.log('Document uploaded successfully')
    lastResult.value = null
  } catch (err) {
    console.error('Upload failed:', err)
    alert('Upload failed: ' + err.message)
  }
}

onMounted(() => {
  fetchRegisteredScanners()
  checkBridge()
  bridgeCheckInterval = setInterval(checkBridge, 5000)
})

onUnmounted(() => {
  if (bridgeCheckInterval) clearInterval(bridgeCheckInterval)
})

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
.glass {
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(20px);
}
.dark .glass {
  background: rgba(15, 23, 42, 0.7);
}
</style>
