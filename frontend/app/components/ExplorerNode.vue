<template>
  <div class="space-y-1">
    <button 
      @click="toggle"
      class="w-full flex items-center justify-between px-3 py-2 rounded-lg text-sm font-bold transition-colors group"
      :class="[
        node.active ? 'text-primary-600 bg-primary-50' : 'text-slate-600 hover:bg-slate-50',
        node.expanded ? 'text-primary-600' : ''
      ]"
    >
      <div class="flex items-center gap-2 overflow-hidden">
        <LucideChevronDown 
          v-if="node.children && node.children.length" 
          class="w-3.5 h-3.5 shrink-0 transition-transform" 
          :class="{ '-rotate-90': !node.expanded }" 
        />
        <div v-else class="w-3.5 shrink-0"></div>
        
        <LucideFolder v-if="!node.expanded" class="w-4 h-4 text-amber-400 fill-amber-400 shrink-0" />
        <LucideFolderOpen v-else class="w-4 h-4 text-amber-400 fill-amber-400 shrink-0" />
        
        <span class="uppercase tracking-tight truncate">{{ node.name }}</span>
      </div>
      <span v-if="node.count" class="text-[10px] text-slate-400 bg-slate-100 px-1.5 py-0.5 rounded-md group-hover:bg-white transition-colors">
        {{ node.count }}
      </span>
    </button>

    <!-- Recursive Children -->
    <div v-if="node.expanded && node.children && node.children.length" class="ml-4 pl-4 border-l border-slate-100 space-y-1">
      <ExplorerNode 
        v-for="child in node.children" 
        :key="child.id" 
        :node="child" 
        :parent-path="currentPath"
        @select="$emit('select', $event)"
      />
    </div>
  </div>
</template>

<script setup>
import { LucideChevronDown, LucideFolder, LucideFolderOpen } from 'lucide-vue-next'

const props = defineProps({
  node: {
    type: Object,
    required: true
  },
  parentPath: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['select'])

const currentPath = computed(() => [...props.parentPath, { 
  name: props.node.name, 
  id: props.node.id, 
  type: props.node.type,
  node: props.node // Keep reference for easier clicking
}])

const toggle = () => {
  if (props.node.children && props.node.children.length) {
    props.node.expanded = !props.node.expanded
  } else {
    emit('select', { node: props.node, path: currentPath.value })
  }
}
</script>
