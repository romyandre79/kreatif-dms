<template>
  <div class="max-w-7xl mx-auto space-y-8 pb-20">
    <!-- Top Navigation -->
    <div class="flex items-center justify-between" v-motion-fade>
      <div class="space-y-4">
        <button 
          @click="navigateTo('/documents')" 
          class="flex items-center gap-2 text-xs font-black text-slate-400 hover:text-primary-600 transition-colors uppercase tracking-widest group"
        >
          <LucideArrowLeft class="w-4 h-4 group-hover:-translate-x-1 transition-transform" />
          {{ $t('documents.detail.btn_back') }}
        </button>
        <div class="space-y-1">
          <h1 class="text-3xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">{{ $t('loans.cart.title') }} <span class="text-slate-400 font-bold" v-html="$t('loans.cart.items_count', { count: cartItems.length })"></span></h1>
          <p class="text-sm font-bold text-slate-400 uppercase tracking-widest">{{ $t('loans.cart.subtitle') }}</p>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8 items-start">
      <!-- Left Column: Cart Items -->
      <div class="lg:col-span-8 space-y-4">
        <div class="glass rounded-lg overflow-hidden border border-slate-100 bg-white/80 backdrop-blur-xl">
          <div class="px-8 py-5 border-b border-slate-100 flex items-center justify-between bg-slate-50/50">
            <label class="flex items-center gap-3 cursor-pointer group">
              <input type="checkbox" class="w-5 h-5 rounded-md border-2 border-slate-200 text-primary-600 focus:ring-primary-500/20" />
              <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest group-hover:text-slate-600 transition-colors">{{ $t('loans.cart.table.select_all') }}</span>
            </label>
            <button @click="clearCart" class="text-[10px] font-black text-orange-500 uppercase tracking-widest hover:underline">{{ $t('loans.cart.table.clear_cart') }}</button>
          </div>

          <div class="divide-y divide-slate-100">
            <div v-for="item in cartItems" :key="item.id" class="p-8 flex items-center gap-6 group hover:bg-slate-50/50 transition-all">
              <input type="checkbox" class="w-5 h-5 rounded-md border-2 border-slate-200 text-primary-600 focus:ring-primary-500/20" />
              <div :class="`w-14 h-14 rounded-2xl flex items-center justify-center shadow-sm ${item.iconBg}`">
                <LucideFileText :class="`w-7 h-7 ${item.iconColor}`" />
              </div>
              <div class="flex-1 min-w-0">
                <p class="text-sm font-black text-[#1E3A5F] truncate uppercase tracking-tight">{{ item.id }} - {{ item.name }}</p>
                <div class="flex items-center gap-3 mt-1.5">
                  <div class="flex items-center gap-1 text-[10px] font-bold text-slate-400">
                    <LucideBuilding2 class="w-3 h-3" />
                    {{ item.dept }}
                  </div>
                  <span class="px-2 py-0.5 bg-green-50 text-green-600 rounded-md text-[8px] font-black uppercase tracking-widest border border-green-100">{{ $t('loans.cart.table.available') }}</span>
                </div>
              </div>
              <button @click="removeFromCart(item.id)" class="p-2.5 rounded-xl text-slate-300 hover:text-red-500 hover:bg-red-50 transition-all opacity-0 group-hover:opacity-100">
                <LucideTrash2 class="w-5 h-5" />
              </button>
            </div>
            
            <div v-if="cartItems.length === 0" class="py-20 flex flex-col items-center justify-center text-slate-300">
              <LucideShoppingCart class="w-20 h-20 mb-4 opacity-20" />
              <p class="text-xs font-black uppercase tracking-widest">Keranjang Kosong</p>
              <p class="text-[10px] mt-2 font-bold opacity-50 uppercase tracking-tight">Pilih dokumen dari explorer untuk ditambahkan</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Column: Loan Summary -->
      <div class="lg:col-span-4 space-y-6" v-motion-slide-visible-bottom>
        <div class="glass rounded-lg p-8 space-y-8 bg-white border border-slate-100 shadow-xl shadow-slate-200/50">
          <h2 class="text-xl font-black text-[#1E3A5F] uppercase tracking-tight">{{ $t('loans.cart.summary.title') }}</h2>
          
          <div class="space-y-6">
            <div class="flex justify-between items-center text-sm">
              <span class="font-bold text-slate-400">{{ $t('loans.cart.summary.total_docs') }}</span>
              <span class="font-black text-[#1E3A5F]">{{ $t('loans.cart.summary.total_items', { count: cartItems.length }) }}</span>
            </div>

            <div class="space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('loans.cart.summary.duration_label') }}</label>
              <div class="relative">
                <select class="w-full pl-4 pr-10 py-3 bg-slate-50 border border-slate-200 rounded-xl text-sm font-bold text-slate-700 appearance-none focus:outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all">
                  <option>{{ $t('loans.cart.summary.duration_options.7_days') }}</option>
                  <option>{{ $t('loans.cart.summary.duration_options.14_days') }}</option>
                  <option>{{ $t('loans.cart.summary.duration_options.30_days') }}</option>
                </select>
                <LucideChevronDown class="absolute right-4 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 pointer-events-none" />
              </div>
            </div>

            <div class="space-y-3">
              <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ $t('loans.cart.summary.purpose_label') }}</label>
              <textarea 
                :placeholder="$t('loans.cart.summary.purpose_placeholder')" 
                rows="4" 
                class="w-full p-4 bg-slate-50 border border-slate-200 rounded-2xl text-sm font-bold text-slate-700 outline-none focus:ring-4 focus:ring-primary-500/10 focus:border-primary-500 transition-all resize-none"
              ></textarea>
            </div>
          </div>

          <div class="p-4 bg-orange-50 border border-orange-100 rounded-2xl flex gap-3">
            <LucideInfo class="w-5 h-5 text-orange-500 shrink-0 mt-0.5" />
            <p class="text-[10px] font-medium text-orange-700 leading-relaxed">
              {{ $t('loans.cart.summary.info_box') }}
            </p>
          </div>

          <div class="space-y-3">
            <button 
              @click="navigateTo('/loans/checkout')"
              :disabled="cartItems.length === 0"
              class="w-full py-4 bg-[#1E3A5F] hover:bg-[#152943] text-white rounded-2xl text-xs font-black uppercase tracking-widest shadow-xl shadow-blue-900/20 transition-all disabled:opacity-50 disabled:cursor-not-allowed"
            >
              {{ $t('loans.cart.summary.btn_checkout') }}
            </button>
            <button 
              @click="navigateTo('/documents')"
              class="w-full py-4 bg-white text-slate-400 hover:text-slate-600 rounded-2xl text-xs font-black uppercase tracking-widest transition-all"
            >
              {{ $t('loans.cart.summary.btn_continue') }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { 
  LucideArrowLeft, 
  LucideFileText, 
  LucideTrash2, 
  LucideBuilding2, 
  LucideChevronDown, 
  LucideInfo 
} from 'lucide-vue-next'
import { useCartStore } from '~/stores/cart'

const cartStore = useCartStore()
const cartItems = computed(() => cartStore.items)

const removeFromCart = (itemId) => {
  cartStore.removeItem(itemId)
}

const clearCart = () => {
  cartStore.clearCart()
}
</script>

<style scoped>
.glass { background: rgba(255, 255, 255, 0.7); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.3); }
</style>
