import { defineStore } from 'pinia'

export const useCartStore = defineStore('cart', {
  state: () => ({
    items: [] as any[]
  }),
  
  getters: {
    count: (state) => state.items.length,
    itemIds: (state) => state.items.map(item => item.id)
  },
  
  actions: {
    addItem(item: any) {
      if (!this.itemIds.includes(item.id)) {
        this.items.push(item)
      }
    },
    
    removeItem(itemId: string) {
      this.items = this.items.filter(item => item.id !== itemId)
    },
    
    toggleItem(item: any) {
      if (this.itemIds.includes(item.id)) {
        this.removeItem(item.id)
      } else {
        this.addItem(item)
      }
    },
    
    clearCart() {
      this.items = []
    },
    
    isInCart(itemId: string) {
      return this.itemIds.includes(itemId)
    }
  },
  
  persist: true
})
