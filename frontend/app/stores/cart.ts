import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { useToast } from '~/composables/useToast'

export const useCartStore = defineStore('cart', () => {
  const items = ref<any[]>([])
  const selectedIds = ref<string[]>([])
  const loanDuration = ref<number>(7)
  const requestedLoanDocIds = ref<string[]>([])
  const toast = useToast()

  const count = computed(() => items.value.length)
  const itemIds = computed(() => items.value.map(item => item.id))
  
  const selectedItems = computed(() => items.value.filter(item => selectedIds.value.includes(item.id)))
  const selectedCount = computed(() => selectedIds.value.length)

  const requestLoans = (ids: string[]) => {
    ids.forEach(id => {
      if (!requestedLoanDocIds.value.includes(id)) {
        requestedLoanDocIds.value.push(id)
      }
    })
  }

  const isInCart = (itemId: string) => {
    return items.value.some(item => item.id === itemId || item.original_id === itemId)
  }

  const isSelected = (itemId: string) => {
    return selectedIds.value.includes(itemId) || items.value.some(item => (item.id === itemId || item.original_id === itemId) && selectedIds.value.includes(item.id))
  }

  const addItem = (item: any) => {
    if (!isInCart(item.id) && !isInCart(item.original_id)) {
      items.value.push(item)
      if (!selectedIds.value.includes(item.id)) {
        selectedIds.value.push(item.id)
      }
      toast.success('Dokumen ditambahkan ke keranjang')
    }
  }

  const removeItem = (itemId: string) => {
    items.value = items.value.filter(item => item.id !== itemId && item.original_id !== itemId)
    selectedIds.value = selectedIds.value.filter(id => id !== itemId && items.value.some(item => item.id === id))
    toast.info('Dokumen dihapus dari keranjang')
  }

  const toggleItem = (item: any) => {
    if (isInCart(item.id)) {
      removeItem(item.id)
    } else {
      addItem(item)
    }
  }

  const toggleSelect = (itemId: string) => {
    if (selectedIds.value.includes(itemId)) {
      selectedIds.value = selectedIds.value.filter(id => id !== itemId)
    } else {
      selectedIds.value.push(itemId)
    }
  }

  const selectAll = () => {
    selectedIds.value = items.value.map(item => item.id)
  }

  const deselectAll = () => {
    selectedIds.value = []
  }

  const clearCart = () => {
    items.value = []
    selectedIds.value = []
  }

  return {
    items,
    selectedIds,
    loanDuration,
    selectedItems,
    selectedCount,
    count,
    itemIds,
    isInCart,
    isSelected,
    addItem,
    removeItem,
    toggleItem,
    toggleSelect,
    selectAll,
    deselectAll,
    clearCart,
    requestedLoanDocIds,
    requestLoans
  }
}, {
  persist: true
})
