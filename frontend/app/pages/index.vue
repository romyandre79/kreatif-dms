<template>
  <div class="min-h-screen bg-[#020617] text-white selection:bg-primary-500/30 overflow-x-hidden">
    <!-- Navigation -->
    <nav class="fixed top-0 w-full z-50 px-6 py-4 flex items-center justify-between backdrop-blur-md bg-slate-950/20 border-b border-white/5">
      <div class="flex items-center gap-3">
        <div class="w-10 h-10 rounded-xl bg-primary-500 flex items-center justify-center">
          <LucideFileStack class="w-6 h-6 text-white" />
        </div>
        <span class="text-xl font-bold tracking-tighter">Kreatif <span class="text-primary-500">DMS</span></span>
      </div>
      
      <div class="hidden md:flex items-center gap-8">
        <a href="#features" class="text-sm font-medium text-slate-400 hover:text-white transition-colors">Features</a>
        <a href="#solutions" class="text-sm font-medium text-slate-400 hover:text-white transition-colors">Solutions</a>
        <a href="#security" class="text-sm font-medium text-slate-400 hover:text-white transition-colors">Security</a>
      </div>

      <div class="flex items-center gap-4">
        <NuxtLink to="/login" class="text-sm font-semibold hover:text-primary-400 transition-colors">Sign In</NuxtLink>
        <NuxtLink to="/dashboard" class="px-5 py-2.5 rounded-full bg-white text-slate-950 text-sm font-bold hover:bg-slate-200 transition-all">
          Get Started
        </NuxtLink>
      </div>
    </nav>

    <!-- Hero Section -->
    <section class="relative pt-32 pb-20 px-6 lg:pt-48 lg:pb-32">
      <!-- Background Decor -->
      <div class="absolute top-0 left-1/2 -translate-x-1/2 w-full max-w-7xl h-[600px] bg-primary-500/10 blur-[120px] rounded-full -z-10"></div>
      <div class="absolute top-20 right-0 w-96 h-96 bg-blue-500/10 blur-[100px] rounded-full -z-10"></div>

      <div class="max-w-7xl mx-auto grid lg:grid-cols-2 gap-16 items-center">
        <div v-motion-slide-visible-left>
          <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-primary-500/10 border border-primary-500/20 text-primary-400 text-xs font-bold uppercase tracking-widest mb-6">
            <span class="w-2 h-2 rounded-full bg-primary-500 animate-pulse"></span>
            Next Generation Document Management
          </div>
          <h1 class="text-5xl lg:text-7xl font-extrabold tracking-tight leading-[1.1] mb-8">
            Manage your <span class="bg-gradient-to-r from-primary-400 to-blue-400 bg-clip-text text-transparent">Documents</span> with Intelligence.
          </h1>
          <p class="text-lg text-slate-400 max-w-xl mb-10 leading-relaxed">
            Experience the future of DMS. AI-powered OCR, full-text search, and military-grade encryption all in one seamless, premium interface.
          </p>
          <div class="flex flex-wrap gap-4">
            <NuxtLink to="/dashboard" class="px-8 py-4 rounded-2xl bg-primary-500 text-white font-bold text-lg hover:bg-primary-400 shadow-xl shadow-primary-500/30 transition-all flex items-center gap-2">
              Start Free Trial <LucideArrowRight class="w-5 h-5" />
            </NuxtLink>
            <button class="px-8 py-4 rounded-2xl bg-white/5 border border-white/10 font-bold text-lg hover:bg-white/10 transition-all">
              Watch Demo
            </button>
          </div>
          
          <div class="mt-12 flex items-center gap-6 text-slate-500">
            <div class="flex -space-x-3">
              <div v-for="i in 4" :key="i" class="w-10 h-10 rounded-full border-2 border-[#020617] bg-slate-800"></div>
            </div>
            <p class="text-sm font-medium">Trusted by <span class="text-white font-bold">500+</span> companies worldwide</p>
          </div>
        </div>

        <div v-motion-slide-visible-right class="relative">
          <div class="glass p-2 rounded-[2.5rem] shadow-2xl shadow-primary-500/10 relative z-10 overflow-hidden group">
            <img 
              src="/kreatif_dms_hero.png" 
              alt="Kreatif DMS Dashboard" 
              class="w-full rounded-[2rem] transition-transform duration-700 group-hover:scale-105"
            />
            <div class="absolute inset-0 bg-gradient-to-t from-[#020617]/50 to-transparent"></div>
          </div>
          <!-- Decorative bits -->
          <div class="absolute -bottom-6 -left-6 glass p-4 rounded-2xl shadow-xl z-20 flex items-center gap-3" v-motion-pop-visible>
             <div class="w-10 h-10 rounded-full bg-green-500/20 flex items-center justify-center">
                <LucideShieldCheck class="w-6 h-6 text-green-500" />
             </div>
             <div>
                <p class="text-xs font-bold text-slate-400">Security</p>
                <p class="text-sm font-bold">AES-256 Enabled</p>
             </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Features Section -->
    <section id="features" class="py-24 px-6 relative">
      <div class="max-w-7xl mx-auto">
        <div class="text-center max-w-2xl mx-auto mb-20">
          <h2 class="text-4xl font-bold mb-4">Powerful features for modern teams</h2>
          <p class="text-slate-400">Built from the ground up to provide the fastest, most secure, and most intelligent document experience.</p>
        </div>

        <div class="grid md:grid-cols-3 gap-8">
          <div v-for="(f, i) in features" :key="i" v-motion-slide-visible-bottom class="glass p-8 rounded-3xl group hover:border-primary-500/30 transition-all">
            <div :class="`w-14 h-14 rounded-2xl flex items-center justify-center mb-6 bg-slate-900 group-hover:bg-primary-500 transition-colors`">
              <component :is="f.icon" class="w-8 h-8 text-primary-400 group-hover:text-white" />
            </div>
            <h3 class="text-xl font-bold mb-4">{{ f.title }}</h3>
            <p class="text-slate-400 leading-relaxed">{{ f.desc }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Stats Section -->
    <section class="py-20 border-y border-white/5 bg-white/[0.02]">
      <div class="max-w-7xl mx-auto px-6 grid grid-cols-2 lg:grid-cols-4 gap-12 text-center">
        <div v-for="(s, i) in stats" :key="i">
          <p class="text-4xl font-extrabold mb-2">{{ s.value }}</p>
          <p class="text-sm font-bold text-primary-500 uppercase tracking-widest">{{ s.label }}</p>
        </div>
      </div>
    </section>

    <!-- Footer -->
    <footer class="py-20 px-6 border-t border-white/5">
      <div class="max-w-7xl mx-auto grid md:grid-cols-4 gap-12 mb-20">
        <div class="col-span-2">
          <div class="flex items-center gap-3 mb-6">
            <div class="w-8 h-8 rounded-lg bg-primary-500 flex items-center justify-center">
              <LucideFileStack class="w-5 h-5 text-white" />
            </div>
            <span class="text-lg font-bold tracking-tighter">Kreatif <span class="text-primary-500">DMS</span></span>
          </div>
          <p class="text-slate-500 max-w-sm">
            Empowering organizations to handle their knowledge with precision, speed, and absolute security.
          </p>
        </div>
        <div>
          <h4 class="font-bold mb-6">Product</h4>
          <ul class="space-y-4 text-sm text-slate-500">
            <li><a href="#" class="hover:text-white transition-colors">Features</a></li>
            <li><a href="#" class="hover:text-white transition-colors">Security</a></li>
            <li><a href="#" class="hover:text-white transition-colors">Pricing</a></li>
          </ul>
        </div>
        <div>
          <h4 class="font-bold mb-6">Company</h4>
          <ul class="space-y-4 text-sm text-slate-500">
            <li><a href="#" class="hover:text-white transition-colors">About Us</a></li>
            <li><a href="#" class="hover:text-white transition-colors">Careers</a></li>
            <li><a href="#" class="hover:text-white transition-colors">Contact</a></li>
          </ul>
        </div>
      </div>
      <div class="max-w-7xl mx-auto flex flex-col md:flex-row justify-between items-center pt-8 border-t border-white/5 text-xs text-slate-600 font-bold">
        <p>© 2024 Kreatif Digital. All rights reserved.</p>
        <div class="flex gap-8 mt-4 md:mt-0">
          <a href="#">Privacy Policy</a>
          <a href="#">Terms of Service</a>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { 
  LucideFileStack, 
  LucideArrowRight, 
  LucideScanText, 
  LucideShieldCheck, 
  LucideZap, 
  LucideSearch, 
  LucideCpu, 
  LucideCloud 
} from 'lucide-vue-next'

definePageMeta({
  layout: false
})

const features = [
  { 
    title: 'AI-Powered OCR', 
    desc: 'Extract text from any document with 99.9% accuracy using our advanced neural network engine.', 
    icon: LucideScanText 
  },
  { 
    title: 'Fast Search', 
    desc: 'Instantly find any document by its content, metadata, or tags with sub-second latency.', 
    icon: LucideSearch 
  },
  { 
    title: 'Military Security', 
    desc: 'All documents are encrypted with AES-256 and stored in secure hierarchical compartments.', 
    icon: LucideShieldCheck 
  },
  { 
    title: 'Smart Summaries', 
    desc: 'Get instant insights from long documents using built-in AI summarization (Gemini/OpenAI).', 
    icon: LucideCpu 
  },
  { 
    title: 'Cloud Storage', 
    desc: 'Scalable MinIO storage ensures your documents are always accessible and highly available.', 
    icon: LucideCloud 
  },
  { 
    title: 'Instant Preview', 
    desc: 'View documents instantly with dynamic watermarking for unauthorized download protection.', 
    icon: LucideZap 
  }
]

const stats = [
  { label: 'Documents Processed', value: '10M+' },
  { label: 'Uptime Guarantee', value: '99.9%' },
  { label: 'Happy Companies', value: '500+' },
  { label: 'AI Responses/sec', value: '250' }
]
</script>

<style scoped>
.glass {
  background: rgba(15, 23, 42, 0.5);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.05);
}
</style>
