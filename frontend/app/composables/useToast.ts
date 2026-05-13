import { useToastStore } from '~/stores/toast'

export const useToast = () => {
  const store = useToastStore()

  const success = (message: string) => {
    store.addToast(message, 'success')
  }

  const error = (message: string) => {
    store.addToast(message, 'error')
  }

  const info = (message: string) => {
    store.addToast(message, 'info')
  }

  const warning = (message: string) => {
    store.addToast(message, 'warning')
  }

  return {
    success,
    error,
    info,
    warning
  }
}
