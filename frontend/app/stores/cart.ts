import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { useToast } from '~/composables/useToast'

export const useCartStore = defineStore('cart', () => {
  const items = ref<any[]>([])
  const toast = useToast()

  const count = computed(() => items.value.length)
  const itemIds = computed(() => items.value.map(item => item.id))

  const isInCart = (itemId: string) => {
    return itemIds.value.includes(itemId)
  }

  const addItem = (item: any) => {
    if (!isInCart(item.id)) {
      items.value.push(item)
      toast.success('Dokumen ditambahkan ke keranjang')
    }
  }

  const removeItem = (itemId: string) => {
    items.value = items.value.filter(item => item.id !== itemId)
    toast.info('Dokumen dihapus dari keranjang')
  }

  const toggleItem = (item: any) => {
    if (isInCart(item.id)) {
      removeItem(item.id)
    } else {
      addItem(item)
    }
  }

  const clearCart = () => {
    items.value = []
    toast.warning('Keranjang dikosongkan')
  }

  return {
    items,
    count,
    itemIds,
    isInCart,
    addItem,
    removeItem,
    toggleItem,
    clearCart
  }
}, {
  persist: true
})
