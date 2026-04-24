<template>
  <div class="max-w-4xl mx-auto space-y-8 pb-20">
    <div v-motion-fade>
      <h1 class="text-3xl font-bold text-slate-900 dark:text-white">My Profile</h1>
      <p class="text-slate-500 dark:text-slate-400 mt-1">Manage your identity and digital signatures.</p>
    </div>

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
          <h3 class="text-[10px] font-black uppercase tracking-widest text-slate-400 mb-4">Account Security</h3>
          <button class="w-full py-3 bg-slate-50 dark:bg-slate-900/50 rounded-xl text-xs font-black uppercase tracking-widest text-slate-600 dark:text-slate-400 hover:bg-primary-50 dark:hover:bg-primary-900/20 hover:text-primary-600 transition-all">
            Change Password
          </button>
        </div>
      </div>

      <!-- Right: Detailed Info & Signature -->
      <div class="lg:col-span-2 space-y-8">
        <!-- Personal Info -->
        <div class="glass rounded-3xl p-8" v-motion-slide-visible-bottom>
          <h3 class="font-black text-lg text-slate-800 dark:text-white mb-8">Personal Information</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="space-y-2">
              <label class="text-[10px] font-black uppercase tracking-widest text-slate-400 ml-1">Full Name</label>
              <input 
                v-model="profileData.fullName"
                type="text" 
                class="w-full px-4 py-3 rounded-xl border border-slate-200 dark:border-slate-800 bg-white dark:bg-slate-900 text-sm font-bold focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 outline-none transition-all"
              />
            </div>
            <div class="space-y-2">
              <label class="text-[10px] font-black uppercase tracking-widest text-slate-400 ml-1">Email Address</label>
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
            <h3 class="font-black text-lg text-slate-800 dark:text-white">Digital Signature</h3>
            <span class="bg-blue-50 dark:bg-blue-900/20 text-blue-600 dark:text-blue-400 text-[9px] font-black uppercase px-2 py-1 rounded-md">Used for Approvals</span>
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div class="space-y-4">
              <p class="text-xs font-bold text-slate-500 leading-relaxed">
                Your signature will be used to authorize document loans and approvals. Ensure it is clear and legible.
              </p>
              <SignaturePad v-model="profileData.signature" />
            </div>
            
            <div class="flex flex-col items-center justify-center p-8 bg-slate-50/50 dark:bg-slate-900/50 rounded-2xl border-2 border-dashed border-slate-200 dark:border-slate-800">
              <p class="text-[10px] font-black uppercase tracking-widest text-slate-400 mb-6">Signature Preview</p>
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
            Reset Changes
          </button>
          <button 
            @click="handleSave"
            :disabled="saving"
            class="px-10 py-3.5 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-xl font-black text-xs uppercase tracking-widest shadow-xl shadow-blue-900/20 transition-all disabled:opacity-50"
          >
            <span v-if="saving">Saving Profile...</span>
            <span v-else>Save All Changes</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { LucideUser, LucideCamera, LucideShieldCheck } from 'lucide-vue-next'
import { useAuthStore } from '~/stores/auth'
import SignaturePad from '~/components/SignaturePad.vue'

const auth = useAuthStore()
const saving = ref(false)
const userPhoto = ref(auth.user?.avatar_url || null)

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
    alert('Profile updated successfully!')
  }, 1000)
}
</script>
