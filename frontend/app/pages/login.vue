<!-- HMR Test -->
<template>
  <div class="min-h-screen flex flex-col md:flex-row bg-white dark:bg-slate-950 relative">
    <!-- Language Switcher (Fixed) -->
    <div class="absolute top-6 right-6 z-50">
      <LanguageSwitcher />
    </div>

    <!-- Left Side: Branding & Info -->
    <div class="md:w-5/12 bg-[#2D5A8E] p-12 flex flex-col justify-between text-white relative overflow-hidden">
      <!-- Background Decor -->
      <div class="absolute top-[-10%] left-[-10%] w-80 h-80 bg-white/10 blur-[80px] rounded-full"></div>
      <div class="absolute bottom-[-10%] right-[-10%] w-60 h-60 bg-primary-500/20 blur-[60px] rounded-full"></div>

      <div class="z-10" v-motion-slide-left>
        <div class="flex items-center gap-3 mb-20">
          <div class="w-10 h-10 rounded-lg bg-white/20 flex items-center justify-center backdrop-blur-sm border border-white/30 overflow-hidden p-2">
            <img src="/logo.png" alt="Logo" class="w-full h-full object-contain" />
          </div>
          <span class="text-2xl font-black tracking-tighter uppercase">{{ config.public.appName }}</span>
        </div>

        <h1 class="text-5xl font-bold leading-tight mb-6">{{ $t('login.left_panel.title') }}</h1>
        <div class="inline-flex px-3 py-1 bg-white/20 border border-white/30 rounded-md text-[10px] font-bold uppercase tracking-widest mb-8">
          {{ $t('login.left_panel.edition') }}
        </div>
        <p class="text-xl text-white/70 max-w-sm leading-relaxed">
          {{ $t('login.left_panel.subtitle') }}
        </p>
      </div>

      <div class="z-10 flex justify-between items-end text-[10px] font-medium text-white/50" v-motion-fade>
        <p>© 2024 {{ config.public.copyright }}. All rights reserved.</p>
        <p>Version 0.0.1</p>
      </div>
    </div>

    <!-- Right Side: Login Form -->
    <div class="flex-1 flex flex-col justify-center items-center p-8 py-12 lg:p-12 bg-white dark:bg-slate-950 md:w-7/12">
      <div class="w-full">
        <h2 class="text-3xl md:text-5xl font-extrabold text-slate-900 dark:text-white mb-2 md:mb-4">
          {{ mode === 'login' ? $t('login.title') : $t('register.title') }}
        </h2>
        <p class="text-lg md:text-xl text-slate-500 dark:text-slate-400 mb-8 md:mb-14">
          {{ mode === 'login' ? $t('login.subtitle') : $t('register.subtitle') }}
        </p>

        <!-- Auth Tabs (Login Only) -->
        <div v-if="mode === 'login'" class="flex border-b border-slate-200 dark:border-slate-800 mb-8">
          <button 
            @click="authType = 'sso'"
            class="flex-1 pb-3 text-sm font-semibold transition-all relative"
            :class="authType === 'sso' ? 'text-primary-600 dark:text-primary-400' : 'text-slate-400 hover:text-slate-600'"
          >
            {{ $t('login.sso_tab') }}
            <div v-if="authType === 'sso'" class="absolute bottom-0 left-0 w-full h-0.5 bg-primary-500" v-motion-pop></div>
          </button>
          <button 
            @click="authType = 'local'"
            class="flex-1 pb-3 text-sm font-semibold transition-all relative"
            :class="authType === 'local' ? 'text-primary-600 dark:text-primary-400' : 'text-slate-400 hover:text-slate-600'"
          >
            {{ $t('login.local_tab') }}
            <div v-if="authType === 'local'" class="absolute bottom-0 left-0 w-full h-0.5 bg-primary-500" v-motion-pop></div>
          </button>
        </div>

        <!-- Form -->
        <form @submit.prevent="handleSubmit" class="space-y-6 md:space-y-10">
          <div v-if="mode === 'register'" class="space-y-4" v-motion-slide-visible-top>
            <label class="text-lg font-bold text-slate-700 dark:text-slate-300">
              {{ $t('register.name_label') }}
            </label>
            <div class="relative group">
              <LucideUser class="absolute left-6 top-1/2 -translate-y-1/2 w-7 h-7 text-slate-400 group-focus-within:text-primary-500 transition-colors" />
              <input 
                v-model="form.fullName"
                type="text" 
                :placeholder="$t('register.name_placeholder')"
                class="w-full pl-14 md:pl-16 pr-6 py-4 md:py-5.5 rounded-[1.5rem] md:rounded-[2rem] border-2 border-slate-200 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-900/50 focus:outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all text-base md:text-lg"
                required
              />
            </div>
          </div>

          <div class="space-y-4">
            <label class="text-lg font-bold text-slate-700 dark:text-slate-300">
              {{ authType === 'sso' ? $t('login.username_label') : $t('login.email_label') }}
            </label>
            <div class="relative group">
              <LucideMail class="absolute left-6 top-1/2 -translate-y-1/2 w-7 h-7 text-slate-400 group-focus-within:text-primary-500 transition-colors" />
              <input 
                v-model="form.email"
                type="text" 
                :placeholder="authType === 'sso' ? $t('login.username_placeholder') : $t('login.email_placeholder')"
                class="w-full pl-14 md:pl-16 pr-6 py-4 md:py-5.5 rounded-[1.5rem] md:rounded-[2rem] border-2 border-slate-200 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-900/50 focus:outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all text-base md:text-lg"
                required
              />
            </div>
          </div>

          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <label class="text-lg font-bold text-slate-700 dark:text-slate-300">{{ $t('login.password_label') }}</label>
              <button 
                v-if="mode === 'login'" 
                type="button"
                @click="showResetModal = true"
                class="text-base font-bold text-slate-500 hover:text-primary-500 transition-colors"
              >
                {{ $t('login.forgot_password') }}
              </button>
            </div>
            <div class="relative group">
              <LucideLock class="absolute left-6 top-1/2 -translate-y-1/2 w-7 h-7 text-slate-400 group-focus-within:text-primary-500 transition-colors" />
              <input 
                v-model="form.password"
                :type="showPassword ? 'text' : 'password'" 
                :placeholder="$t('login.password_placeholder')"
                class="w-full pl-14 md:pl-16 pr-14 md:pr-16 py-4 md:py-5.5 rounded-[1.5rem] md:rounded-[2rem] border-2 border-slate-200 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-900/50 focus:outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all text-base md:text-lg"
                required
              />
              <button 
                type="button"
                @click="showPassword = !showPassword"
                class="absolute right-6 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-600 transition-colors"
              >
                <LucideEye v-if="!showPassword" class="w-6 h-6" />
                <LucideEyeOff v-else class="w-6 h-6" />
              </button>
            </div>
          </div>

          <!-- Confirm Password (Register Only) -->
          <div v-if="mode === 'register'" class="space-y-4" v-motion-slide-visible-top>
            <label class="text-lg font-bold text-slate-700 dark:text-slate-300">
              {{ $t('register.confirm_password_label') }}
            </label>
            <div class="relative group">
              <LucideLock class="absolute left-6 top-1/2 -translate-y-1/2 w-7 h-7 text-slate-400 group-focus-within:text-primary-500 transition-colors" />
              <input 
                v-model="form.confirmPassword"
                type="password" 
                :placeholder="$t('register.confirm_password_placeholder')"
                class="w-full pl-14 md:pl-16 pr-6 py-4 md:py-5.5 rounded-[1.5rem] md:rounded-[2rem] border-2 border-slate-200 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-900/50 focus:outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all text-base md:text-lg"
                required
              />
            </div>
          </div>

          <!-- Alerts -->
          <div class="space-y-4">
            <Alert 
              v-model="showError" 
              :message="error" 
              type="error" 
            />
            <Alert 
              v-model="showSuccess" 
              :message="successMessage" 
              type="success" 
            />
          </div>

          <button 
            type="submit"
            :disabled="loading"
            class="w-full py-5 md:py-6 bg-[#1E3A5F] hover:bg-[#152943] disabled:bg-slate-400 text-white font-black text-lg md:text-xl rounded-[1.5rem] md:rounded-[2rem] transition-all shadow-2xl shadow-blue-900/30 flex items-center justify-center gap-3"
          >
            <LucideLoader2 v-if="loading" class="w-6 h-6 animate-spin" />
            <span v-else>{{ mode === 'login' ? $t('login.sign_in_button') : $t('register.submit_button') }}</span>
            <LucideArrowRight v-if="!loading" class="w-6 h-6" />
          </button>
        </form>

        <p v-if="mode === 'login' && authType === 'local'" class="mt-8 text-center text-base text-slate-500">
          {{ $t('login.no_account') }} 
          <button @click="mode = 'register'" class="font-bold text-primary-600 hover:text-primary-700 hover:underline transition-colors">
            {{ $t('login.register_now') }}
          </button>
        </p>

        <p v-if="mode === 'register'" class="mt-8 text-center text-base text-slate-500">
          {{ $t('register.already_have_account') }} 
          <button @click="mode = 'login'; authType = 'local'" class="font-bold text-primary-600 hover:text-primary-700 hover:underline transition-colors">
            {{ $t('register.login_now') }}
          </button>
        </p>

        <p class="mt-12 text-center text-sm text-slate-500">
          {{ $t('login.trouble_signing_in') }} <a href="#" class="font-bold text-slate-900 dark:text-white hover:underline">{{ $t('login.contact_support') }}</a>
        </p>

        <!-- Reset Password Modal -->
        <Modal v-model="showResetModal" :title="$t('reset.title')">
          <div class="space-y-6">
            <div class="p-4 bg-blue-50 dark:bg-blue-900/20 border border-blue-100 dark:border-blue-800 rounded-xl flex gap-3 text-blue-700 dark:text-blue-300 text-sm leading-relaxed">
              <LucideInfo class="w-5 h-5 flex-shrink-0 mt-0.5" />
              <p>{{ $t('reset.info') }}</p>
            </div>

            <div class="space-y-2">
              <label class="text-sm font-bold text-slate-700 dark:text-slate-300">{{ $t('reset.email_label') }}</label>
              <input 
                v-model="resetEmail"
                type="email" 
                :placeholder="$t('reset.email_placeholder')"
                class="w-full px-4 py-3 rounded-xl border border-slate-200 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-900/50 focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500 transition-all"
              />
            </div>
          </div>

          <template #footer>
            <button 
              @click="showResetModal = false"
              class="px-6 py-2.5 text-slate-600 dark:text-slate-400 font-bold text-sm hover:text-slate-900 dark:hover:text-white transition-colors"
            >
              {{ $t('reset.cancel_button') }}
            </button>
            <button 
              @click="handleResetPassword"
              :disabled="loading"
              class="px-6 py-2.5 bg-[#1E3A5F] hover:bg-[#152943] text-white font-bold text-sm rounded-xl transition-all shadow-lg shadow-blue-900/20 flex items-center gap-2"
            >
              <LucideLoader2 v-if="loading" class="w-4 h-4 animate-spin" />
              {{ $t('reset.submit_button') }}
            </button>
          </template>
        </Modal>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideFileStack, 
  LucideMail, 
  LucideLock, 
  LucideEye, 
  LucideEyeOff, 
  LucideArrowRight,
  LucideLoader2,
  LucideAlertCircle,
  LucideUser,
  LucideInfo
} from 'lucide-vue-next'
import { useAuthStore } from '~/stores/auth'

const config = useRuntimeConfig()
definePageMeta({
  layout: false
})

const { t } = useI18n()

useHead({
  title: t('login.title')
})

const authType = ref('sso')
const mode = ref('login')
const showPassword = ref(false)
const loading = ref(false)
const error = ref('')
const showError = ref(false)
const successMessage = ref('')
const showSuccess = ref(false)
const showResetModal = ref(false)
const resetEmail = ref('')

const form = reactive({
  fullName: '',
  email: '',
  password: '',
  confirmPassword: ''
})

const handleSubmit = async () => {
  if (mode.value === 'login') {
    await handleLogin()
  } else {
    await handleRegister()
  }
}

const handleRegister = async () => {
  if (form.password !== form.confirmPassword) {
    error.value = t('register.errors.password_mismatch')
    showError.value = true
    return
  }

  loading.value = true
  error.value = ''
  showError.value = false
  successMessage.value = ''
  showSuccess.value = false

  try {
    const config = useRuntimeConfig()
    await $fetch(`${config.public.apiBase}/auth/register`, {
      method: 'POST',
      body: {
        full_name: form.fullName,
        email: form.email,
        password: form.password
      }
    })
    
    successMessage.value = t('register.success_message')
    showSuccess.value = true
    
    // Switch to login after success
    setTimeout(() => {
      mode.value = 'login'
      authType.value = 'local'
    }, 3000)
    
  } catch (err) {
    error.value = err.data?.message || t('register.errors.failed')
    showError.value = true
  } finally {
    loading.value = false
  }
}
const handleLogin = async () => {
  loading.value = true
  error.value = ''
  showError.value = false
  
  try {
    const config = useRuntimeConfig()
    const res = await $fetch(`${config.public.apiBase}/auth/login`, {
      method: 'POST',
      body: {
        identifier: form.email,
        password: form.password,
        auth_type: authType.value
      }
    })
    
    const auth = useAuthStore()
    auth.setTokens(res.data.access_token, res.data.refresh_token)
    auth.setUser({
      id: res.data.user_id,
      full_name: res.data.full_name,
      email: form.email,
      role: res.data.role,
      avatar_url: res.data.avatar_url,
      signature_url: res.data.signature_url
    })
    
    navigateTo('/dashboard')
    
  } catch (err) {
    error.value = err.data?.message || t('login.errors.invalid_credentials')
    showError.value = true
  } finally {
    loading.value = false
  }
}

const handleResetPassword = async () => {
  if (!resetEmail.value) return

  loading.value = true
  try {
    const config = useRuntimeConfig()
    await $fetch(`${config.public.apiBase}/auth/forgot-password`, {
      method: 'POST',
      body: {
        email: resetEmail.value
      }
    })
    
    successMessage.value = t('reset.success_message')
    showSuccess.value = true
    showResetModal.value = false
    resetEmail.value = ''
  } catch (err) {
    error.value = err.data?.message || 'Failed to send reset link.'
    showError.value = true
  } finally {
    loading.value = false
  }
}
</script>
