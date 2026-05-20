<template>
  <Transition
    enter-active-class="transition duration-500 ease-out"
    enter-from-class="translate-y-20 opacity-0"
    enter-to-class="translate-y-0 opacity-100"
    leave-active-class="transition duration-300 ease-in"
    leave-from-class="translate-y-0 opacity-100"
    leave-to-class="translate-y-20 opacity-0"
  >
    <div v-if="isVisible" class="fixed bottom-12 left-1/2 -translate-x-1/2 z-[100] lg:ml-[140px]">
      <div class="bg-[#1E3A5F] text-white rounded-full px-8 py-4 flex items-center gap-10 shadow-[0_20px_50px_rgba(30,58,95,0.4)] border border-white/5 backdrop-blur-md ring-1 ring-white/10">
        <div class="flex items-center gap-4 border-r border-white/10 pr-10">
          <div class="relative">
            <LucideShoppingCart class="w-7 h-7 text-blue-100" />
            <span class="absolute -top-2 -right-2 w-5 h-5 bg-red-500 text-[10px] font-black flex items-center justify-center rounded-full border-2 border-[#1E3A5F] ring-1 ring-red-400/50 shadow-lg shadow-red-500/20">
              {{ cartStore.count }}
            </span>
          </div>
          <div class="leading-tight">
            <p class="text-[11px] font-black uppercase tracking-tight">{{ cartStore.count }} {{ cartStore.count > 1 ? 'items' : 'item' }} {{ $t('documents.explorer.cart_suffix', 'in cart') }}</p>
            <p class="text-[9px] font-bold text-blue-300 uppercase tracking-widest mt-0.5">{{ $t('documents.explorer.ready_loan', 'Physical Loan Queue') }}</p>
          </div>
        </div>
        
        <div class="flex items-center gap-4">
          <button 
            @click="navigateTo('/loans/cart')" 
            class="bg-white text-[#1E3A5F] px-8 py-3 rounded-full text-xs font-black uppercase tracking-widest hover:scale-105 active:scale-95 transition-all shadow-lg"
          >
            {{ $t('documents.explorer.view_cart') }}
          </button>
          <button 
            @click="cartStore.clearCart()"
            class="p-3 text-white/40 hover:text-red-400 transition-colors"
            title="Clear all"
          >
            <LucideX class="w-5 h-5" />
          </button>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup>
import { useCartStore } from '~/stores/cart'
import { LucideShoppingCart, LucideX } from 'lucide-vue-next'
import { useRoute } from 'vue-router'

const cartStore = useCartStore()
const route = useRoute()

const isVisible = computed(() => {
  const hideOn = ['/loans/cart', '/loans/checkout']
  return cartStore.count > 0 && !hideOn.includes(route.path)
})
</script>
