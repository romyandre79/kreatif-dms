export const formatBytes = (bytes: number, decimals = 2) => {
  if (!+bytes) return '0 Bytes'

  const k = 1024
  const dm = decimals < 0 ? 0 : decimals
  const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']

  const i = Math.floor(Math.log(bytes) / Math.log(k))

  return `${parseFloat((bytes / Math.pow(k, i)).toFixed(dm))} ${sizes[i]}`
}

export const timeAgo = (date: string | Date) => {
  const d = new Date(date)
  const now = new Date()
  const diff = now.getTime() - d.getTime()
  const seconds = Math.floor(diff / 1000)
  const minutes = Math.floor(seconds / 60)
  const hours = Math.floor(minutes / 60)
  const days = Math.floor(hours / 24)

  if (days > 0) return `${days} days ago`
  if (hours > 0) return `${hours} hours ago`
  if (minutes > 0) return `${minutes} minutes ago`
  return 'Just now'
}

export const parseMarkdown = (text: string) => {
  if (!text) return ''
  return text
    .replace(/\\n/g, '\n') // Handle literal \n from DB
    .replace(/^### (.*$)/gim, '<h3 class="text-lg font-black mt-4 mb-2 text-[#1E3A5F] dark:text-white">$1</h3>')
    .replace(/^## (.*$)/gim, '<h2 class="text-xl font-black mt-6 mb-3 text-[#1E3A5F] dark:text-white">$1</h2>')
    .replace(/^# (.*$)/gim, '<h1 class="text-2xl font-black mt-8 mb-4 text-[#1E3A5F] dark:text-white">$1</h1>')
    .replace(/\*\*(.*)\*\*/gim, '<b class="font-black">$1</b>')
    .replace(/\*(.*)\*/gim, '<i>$1</i>')
    .replace(/^- (.*$)/gim, '<li class="ml-4 list-disc text-slate-600 dark:text-slate-400">$1</li>')
    .replace(/\n/gim, '<br>')
}
