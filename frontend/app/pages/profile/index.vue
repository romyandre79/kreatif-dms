<template>
  <div class="max-w-4xl mx-auto space-y-8 pb-20">
    <PageHeader 
      :title="$t('profile.title')"
      :subtitle="$t('profile.subtitle')"
    />

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
      <!-- Left: Photo & Basics -->
      <div class="lg:col-span-1 space-y-8">
        <div class="glass rounded-3xl p-8 flex flex-col items-center text-center" v-motion-slide-visible-bottom>
          <div class="relative group cursor-pointer mb-6">
            <div class="w-32 h-32 rounded-full overflow-hidden border-4 border-white dark:border-slate-800 shadow-xl bg-slate-100 dark:bg-slate-800 flex items-center justify-center">
              <img v-if="userPhoto" :src="userPhoto" class="w-full h-full object-cover" />
              <LucideUser v-else class="w-12 h-12 text-slate-400" />
            </div>
            <label class="absolute inset-0 flex items-center justify-center bg-black/40 text-white rounded-full opacity-0 group-hover:opacity-100 transition-opacity cursor-pointer">
              <LucideCamera class="w-6 h-6" />
              <input type="file" @change="onPhotoChange" class="hidden" accept="image/*" />
            </label>
          </div>
          <h2 class="font-black text-xl text-slate-900 dark:text-white">{{ auth.user?.full_name }}</h2>
          <p class="text-xs font-bold text-primary-500 uppercase tracking-widest mt-1">{{ auth.user?.role }}</p>
          <p class="text-xs text-slate-400 mt-4">{{ auth.user?.email }}</p>
        </div>

        <div class="glass rounded-2xl p-6" v-motion-slide-visible-bottom>
          <h3 class="text-[10px] font-black uppercase tracking-widest text-slate-400 mb-4">{{ $t('profile.security.title') }}</h3>
          <div class="space-y-3">
            <div class="p-4 rounded-xl border-2 border-slate-50 dark:border-slate-900 bg-white dark:bg-slate-950/50">
              <div class="flex items-center justify-between mb-4">
                <div class="flex items-center gap-2">
                  <LucideShieldCheck class="w-4 h-4" :class="auth.user?.is_mfa_enabled ? 'text-green-500' : 'text-slate-400'" />
                  <span class="text-[10px] font-black uppercase tracking-widest text-slate-600 dark:text-slate-400">Two-Factor Auth</span>
                </div>
                <div v-if="auth.user?.is_mfa_enabled" class="flex items-center gap-1">
                  <div class="w-1.5 h-1.5 bg-green-500 rounded-full animate-pulse"></div>
                  <span class="text-[8px] font-black uppercase text-green-500">ACTIVE</span>
                </div>
              </div>
              <button 
                @click="initMfaSetup"
                class="w-full py-3 rounded-xl text-[10px] font-black uppercase tracking-widest transition-all"
                :class="auth.user?.is_mfa_enabled 
                  ? 'bg-red-50 dark:bg-red-900/10 text-red-600 dark:text-red-400 hover:bg-red-100' 
                  : 'bg-primary-500 text-white shadow-lg shadow-primary-500/20 hover:bg-primary-600'"
              >
                {{ auth.user?.is_mfa_enabled ? 'Disable 2FA' : 'Setup 2FA' }}
              </button>
            </div>

            <div class="p-4 rounded-xl border-2 border-slate-50 dark:border-slate-900 bg-white dark:bg-slate-950/50">
              <div class="flex items-center justify-between mb-4">
                <div class="flex items-center gap-2">
                  <LucideLock class="w-4 h-4" :class="auth.user?.pin_status === 'SET' ? 'text-blue-500' : 'text-slate-400'" />
                  <span class="text-[10px] font-black uppercase tracking-widest text-slate-600 dark:text-slate-400">Digital PIN</span>
                </div>
                <div v-if="auth.user?.pin_status === 'SET'" class="flex items-center gap-1">
                  <div class="w-1.5 h-1.5 bg-blue-500 rounded-full"></div>
                  <span class="text-[8px] font-black uppercase text-blue-500">ACTIVE</span>
                </div>
              </div>
              <button 
                @click="showPinModal = true"
                class="w-full py-3 rounded-xl text-[10px] font-black uppercase tracking-widest transition-all"
                :class="auth.user?.pin_status === 'SET'
                  ? 'bg-blue-50 dark:bg-blue-900/10 text-blue-600 dark:text-blue-400 hover:bg-blue-100' 
                  : 'bg-primary-500 text-white shadow-lg shadow-primary-500/20 hover:bg-primary-600'"
              >
                {{ auth.user?.pin_status === 'SET' ? 'Change PIN' : 'Set Digital PIN' }}
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Right: Detailed Info & Signature -->
      <div class="lg:col-span-2 space-y-8">
        <!-- Personal Info -->
        <div class="glass rounded-3xl p-8" v-motion-slide-visible-bottom>
          <h3 class="font-black text-lg text-slate-800 dark:text-white mb-8">{{ $t('profile.personal.title') }}</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="space-y-2">
              <label class="text-[10px] font-black uppercase tracking-widest text-slate-400 ml-1">{{ $t('profile.personal.name_label') }}</label>
              <input 
                v-model="profileData.fullName"
                type="text" 
                class="w-full px-4 py-3 rounded-xl border border-slate-200 dark:border-slate-800 bg-white dark:bg-slate-900 text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 outline-none transition-all"
              />
            </div>
            <div class="space-y-2">
              <label class="text-[10px] font-black uppercase tracking-widest text-slate-400 ml-1">{{ $t('profile.personal.email_label') }}</label>
              <input 
                :value="auth.user?.email"
                disabled
                type="email" 
                class="w-full px-4 py-3 rounded-xl border border-slate-100 dark:border-slate-800 bg-slate-50 dark:bg-slate-950 text-sm font-bold text-slate-400 cursor-not-allowed"
              />
            </div>
          </div>
        </div>

        <!-- Digital Signature -->
        <div class="glass rounded-3xl p-8" v-motion-slide-visible-bottom>
          <div class="flex items-center justify-between mb-8">
            <h3 class="font-black text-lg text-slate-800 dark:text-white">{{ $t('profile.signature.title') }}</h3>
            <span class="bg-blue-50 dark:bg-blue-900/20 text-blue-600 dark:text-blue-400 text-[9px] font-black uppercase px-2 py-1 rounded-md">{{ $t('profile.signature.used_for') }}</span>
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div class="space-y-4">
              <p class="text-xs font-bold text-slate-500 leading-relaxed">
                {{ $t('profile.signature.hint') }}
              </p>
              <SignaturePad v-model="profileData.signature" />
            </div>
            
            <div class="flex flex-col items-center justify-center p-8 bg-slate-50/50 dark:bg-slate-900/50 rounded-2xl border-2 border-dashed border-slate-200 dark:border-slate-800">
              <p class="text-[10px] font-black uppercase tracking-widest text-slate-400 mb-6">{{ $t('profile.signature.preview') }}</p>
              <div class="w-full aspect-[2/1] bg-white dark:bg-slate-950 rounded-xl border border-slate-100 dark:border-slate-800 flex items-center justify-center overflow-hidden">
                <img v-if="profileData.signature" :src="profileData.signature" class="max-w-full max-h-full object-contain p-4" />
                <LucideShieldCheck v-else class="w-12 h-12 text-slate-100 dark:text-slate-900" />
              </div>
            </div>
          </div>
        </div>

        <!-- Actions -->
        <div class="flex justify-end gap-4" v-motion-fade>
          <button class="px-8 py-3.5 text-xs font-black uppercase tracking-widest text-slate-400 hover:text-slate-600 transition-colors">
            {{ $t('profile.actions.reset') }}
          </button>
          <button 
            @click="handleSave"
            :disabled="saving"
            class="px-10 py-3.5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-xl font-black text-xs uppercase tracking-widest shadow-xl shadow-blue-900/20 transition-all disabled:opacity-50"
          >
            <span v-if="saving">{{ $t('profile.actions.saving') }}</span>
            <span v-else>{{ $t('profile.actions.save') }}</span>
          </button>
        </div>
      </div>
    </div>
    
    <!-- MFA Setup Modal -->
    <Modal v-model="showMfaModal" title="Setup Two-Factor Authentication" size="md">
      <div class="space-y-8 py-4">
        <!-- Step 1: Info -->
        <div v-if="mfaStep === 1" class="space-y-6" v-motion-fade>
          <div class="p-4 bg-primary-50 dark:bg-primary-900/10 rounded-2xl border border-primary-100 dark:border-primary-800/50 flex gap-4">
            <div class="w-10 h-10 bg-primary-500 rounded-xl flex items-center justify-center flex-shrink-0 shadow-lg shadow-primary-500/20">
              <LucideSmartphone class="w-5 h-5 text-white" />
            </div>
            <div class="space-y-1">
              <p class="text-sm font-bold text-slate-800 dark:text-white">Protect your account</p>
              <p class="text-xs text-slate-500 leading-relaxed">2FA adds an extra layer of security. You'll need a code from your authenticator app to log in.</p>
            </div>
          </div>
          
          <div class="space-y-4">
            <p class="text-[10px] font-black uppercase tracking-widest text-slate-400">HOW IT WORKS:</p>
            <ol class="space-y-3">
              <li class="flex items-center gap-3 text-xs font-bold text-slate-600 dark:text-slate-400">
                <span class="w-5 h-5 rounded-full bg-slate-100 dark:bg-slate-800 flex items-center justify-center text-[10px]">1</span>
                Scan the QR code with Google Authenticator or Authy
              </li>
              <li class="flex items-center gap-3 text-xs font-bold text-slate-600 dark:text-slate-400">
                <span class="w-5 h-5 rounded-full bg-slate-100 dark:bg-slate-800 flex items-center justify-center text-[10px]">2</span>
                Enter the 6-digit verification code
              </li>
            </ol>
          </div>
        </div>

        <!-- Step 2: QR Code -->
        <div v-if="mfaStep === 2" class="space-y-8 flex flex-col items-center" v-motion-pop>
          <div class="p-4 bg-white rounded-3xl shadow-2xl border border-slate-100">
            <QrcodeVue :value="mfaData.url" :size="200" level="H" />
          </div>
          
          <div class="text-center space-y-2">
            <p class="text-[10px] font-black uppercase tracking-widest text-slate-400">Unable to scan?</p>
            <div class="px-4 py-2 bg-slate-50 dark:bg-slate-900 rounded-lg border border-slate-100 dark:border-slate-800">
              <code class="text-sm font-black text-primary-600 dark:text-primary-400 tracking-wider">{{ mfaData.secret }}</code>
            </div>
          </div>

          <div class="w-full space-y-4">
            <label class="text-[10px] font-black uppercase tracking-widest text-slate-400 block text-center">Enter Verification Code</label>
            <input 
              v-model="mfaVerifyCode"
              type="text" 
              maxlength="6"
              placeholder="000 000"
              class="w-full text-center text-3xl tracking-[1rem] font-black py-4 rounded-2xl border-2 border-slate-100 dark:border-slate-800 bg-slate-50 dark:bg-slate-900 focus:border-primary-500 focus:ring-4 focus:ring-primary-500/10 outline-none transition-all"
            />
          </div>
        </div>

        <!-- Step 3: Success -->
        <div v-if="mfaStep === 3" class="space-y-6 text-center py-8" v-motion-pop>
          <div class="w-20 h-20 bg-green-500 rounded-full flex items-center justify-center mx-auto shadow-xl shadow-green-500/30">
            <LucideCheck class="w-10 h-10 text-white" />
          </div>
          <div class="space-y-2">
            <h4 class="text-2xl font-black text-slate-900 dark:text-white">2FA Enabled!</h4>
            <p class="text-sm text-slate-500">Your account is now more secure. Next time you login, you'll be asked for a code.</p>
          </div>
        </div>
      </div>

      <template #footer>
        <div class="flex justify-between w-full">
          <button 
            v-if="mfaStep < 3"
            @click="showMfaModal = false"
            class="px-6 py-2.5 text-xs font-black uppercase tracking-widest text-slate-400 hover:text-slate-600 transition-colors"
          >
            Cancel
          </button>
          
          <button 
            v-if="mfaStep === 1"
            @click="generateMfaSecret"
            :disabled="loadingMfa"
            class="px-8 py-2.5 bg-primary-500 hover:bg-primary-600 text-white rounded-xl font-black text-xs uppercase tracking-widest shadow-lg shadow-primary-500/20 transition-all flex items-center gap-2"
          >
            <LucideLoader2 v-if="loadingMfa" class="w-4 h-4 animate-spin" />
            Start Setup
          </button>

          <button 
            v-if="mfaStep === 2"
            @click="verifyMfaSetup"
            :disabled="loadingMfa || mfaVerifyCode.length < 6"
            class="px-8 py-2.5 bg-primary-500 hover:bg-primary-600 text-white rounded-xl font-black text-xs uppercase tracking-widest shadow-lg shadow-primary-500/20 transition-all flex items-center gap-2"
          >
            <LucideLoader2 v-if="loadingMfa" class="w-4 h-4 animate-spin" />
            Verify & Activate
          </button>

          <button 
            v-if="mfaStep === 3"
            @click="showMfaModal = false"
            class="w-full py-2.5 bg-slate-900 dark:bg-white text-white dark:text-slate-900 rounded-xl font-black text-xs uppercase tracking-widest transition-all"
          >
            Great, thanks!
          </button>
        </div>
      </template>
    </Modal>

    <!-- PIN Setup Modal -->
    <Modal v-model="showPinModal" title="Set Digital PIN" size="sm">
      <div class="space-y-6 py-4">
        <div class="text-center space-y-2">
          <LucideLock class="w-12 h-12 text-primary-500 mx-auto" />
          <h4 class="text-lg font-black text-slate-800 dark:text-white">Secure Your Actions</h4>
          <p class="text-xs text-slate-400 leading-relaxed px-4">
            Create a 6-digit PIN to authorize sensitive operations like document approvals and exports.
          </p>
        </div>

        <div class="space-y-4">
          <label class="text-[10px] font-black uppercase tracking-widest text-slate-400 block text-center">New 6-Digit PIN</label>
          <input 
            v-model="pinData.pin"
            type="password" 
            maxlength="6"
            placeholder="......"
            class="w-full text-center text-3xl tracking-[1.5rem] font-black py-4 rounded-2xl border-2 border-slate-100 dark:border-slate-800 bg-slate-50 dark:bg-slate-900 focus:border-primary-500 focus:ring-4 focus:ring-primary-500/10 outline-none transition-all"
          />
          
          <label class="text-[10px] font-black uppercase tracking-widest text-slate-400 block text-center mt-6">Confirm PIN</label>
          <input 
            v-model="pinData.confirm"
            type="password" 
            maxlength="6"
            placeholder="......"
            class="w-full text-center text-3xl tracking-[1.5rem] font-black py-4 rounded-2xl border-2 border-slate-100 dark:border-slate-800 bg-slate-50 dark:bg-slate-900 focus:border-primary-500 focus:ring-4 focus:ring-primary-500/10 outline-none transition-all"
          />
        </div>
      </div>

      <template #footer>
        <div class="flex justify-between w-full">
          <button @click="showPinModal = false" class="px-6 py-2.5 text-xs font-black uppercase tracking-widest text-slate-400 hover:text-slate-600 transition-colors">
            Cancel
          </button>
          <button 
            @click="handleSavePin"
            :disabled="loadingPin || pinData.pin.length !== 6 || pinData.pin !== pinData.confirm"
            class="px-8 py-2.5 bg-primary-500 hover:bg-primary-600 text-white rounded-xl font-black text-xs uppercase tracking-widest shadow-lg shadow-primary-500/20 transition-all flex items-center gap-2"
          >
            <LucideLoader2 v-if="loadingPin" class="w-4 h-4 animate-spin" />
            Save PIN
          </button>
        </div>
      </template>
    </Modal>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { 
  LucideUser, 
  LucideCamera, 
  LucideShieldCheck,
  LucideSmartphone,
  LucideCheck,
  LucideLoader2,
  LucideLock
} from 'lucide-vue-next'
import { useAuthStore } from '~/stores/auth'
import SignaturePad from '~/components/SignaturePad.vue'
import PageHeader from '~/components/PageHeader.vue'
import QrcodeVue from 'qrcode.vue'

const auth = useAuthStore()
const { t } = useI18n()
const saving = ref(false)
const userPhoto = ref(auth.user?.avatar_url || null)

// MFA Setup State
const showMfaModal = ref(false)
const mfaStep = ref(1)
const loadingMfa = ref(false)
const mfaVerifyCode = ref('')
const mfaData = ref({
  secret: '',
  url: ''
})

// PIN Setup State
const showPinModal = ref(false)
const loadingPin = ref(false)
const pinData = ref({
  pin: '',
  confirm: ''
})

const profileData = ref({
  fullName: auth.user?.full_name || '',
  signature: auth.user?.signature_url || null
})

const onPhotoChange = (e) => {
  const file = e.target.files[0]
  if (!file) return
  
  const reader = new FileReader()
  reader.onload = (event) => {
    userPhoto.value = event.target.result
  }
  reader.readAsDataURL(file)
}

const handleSave = async () => {
  saving.value = true
  // In a real app, you would send this to the backend
  // For now, we simulate success
  setTimeout(() => {
    auth.setUser({
      ...auth.user,
      full_name: profileData.value.fullName,
      avatar_url: userPhoto.value,
      signature_url: profileData.value.signature
    })
    saving.value = false
    alert(t('profile.actions.success'))
  }, 1000)
}

// MFA Methods
const initMfaSetup = () => {
  if (auth.user?.is_mfa_enabled) {
    // Logic to disable MFA could go here
    if (confirm('Are you sure you want to disable 2FA? This will reduce your account security.')) {
      disableMfa()
    }
  } else {
    mfaStep.value = 1
    showMfaModal.value = true
  }
}

const generateMfaSecret = async () => {
  loadingMfa.value = true
  try {
    const config = useRuntimeConfig()
    const { useApi } = await import('~/composables/useApi')
    const api = useApi()
    
    const res = await api.post('/auth/mfa/setup')
    mfaData.value = {
      secret: res.data.secret,
      url: res.data.url
    }
    mfaStep.value = 2
  } catch (err) {
    alert('Failed to initiate MFA setup')
  } finally {
    loadingMfa.value = false
  }
}

const verifyMfaSetup = async () => {
  loadingMfa.value = true
  try {
    const config = useRuntimeConfig()
    const { useApi } = await import('~/composables/useApi')
    const api = useApi()
    
    await api.post('/auth/mfa/verify', {
      secret: mfaData.value.secret,
      code: mfaVerifyCode.value
    })
    
    // Update local auth store
    auth.setUser({
      ...auth.user,
      is_mfa_enabled: true
    })
    
    mfaStep.value = 3
  } catch (err) {
    alert(err.response?._data?.message || 'Invalid verification code')
  } finally {
    loadingMfa.value = false
  }
}

const disableMfa = async () => {
  // We'd need a backend endpoint for this too, but for now we'll just mock it or skip
  alert('MFA disabled (Mock)')
  auth.setUser({ ...auth.user, is_mfa_enabled: false })
}

const handleSavePin = async () => {
  loadingPin.value = true
  try {
    const { useApi } = await import('~/composables/useApi')
    const api = useApi()
    
    await api.post('/auth/pin', {
      pin: pinData.value.pin
    })
    
    auth.setUser({
      ...auth.user,
      pin_status: 'SET'
    })
    
    showPinModal.value = false
    pinData.value = { pin: '', confirm: '' }
    alert('Digital PIN updated successfully')
  } catch (err) {
    alert(err.response?._data?.message || 'Failed to update PIN')
  } finally {
    loadingPin.value = false
  }
}
</script>
