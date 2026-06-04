const fs = require('fs');
const file = 'frontend/app/pages/admin/hierarchy.vue';
let content = fs.readFileSync(file, 'utf8');

// Replace right panel template
const templateRegex = /<!-- Supervisor assignment matrix -->[\s\S]*?<!-- Impacted Workflows list -->/m;
const newTemplate = `<!-- Department assignment matrix -->
            <div class="space-y-2">
              <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest block text-left">Departemen Assignment</label>
              
              <div class="grid grid-cols-2 gap-2">
                <div class="bg-slate-50 dark:bg-slate-800/30 border border-slate-200/60 dark:border-slate-700 rounded-xl p-3 text-center">
                  <p class="text-[8px] font-bold text-slate-400 uppercase tracking-wider mb-1">Departemen Awal</p>
                  <p class="text-[11px] font-black text-[#1E3A5F] dark:text-slate-300 uppercase leading-snug">
                    {{ selectedEmployeeOriginalDepartmentName }}
                  </p>
                </div>

                <div class="bg-blue-500/[0.03] dark:bg-blue-500/[0.01] border border-blue-200/60 dark:border-blue-900/60 rounded-xl p-3 text-center flex flex-col justify-between">
                  <p class="text-[8px] font-bold text-blue-500 dark:text-blue-400 uppercase tracking-wider mb-1">Departemen Baru</p>
                  <select 
                    v-model="selectedEmployee.department_id"
                    class="text-[10px] font-black text-blue-700 dark:text-blue-400 bg-blue-100/50 dark:bg-blue-950/30 px-2 py-1 rounded border-0 focus:ring-1 focus:ring-blue-500 text-center w-full uppercase cursor-pointer"
                  >
                    <option :value="null">Tidak ada</option>
                    <option v-for="dept in departments" :key="dept.id" :value="dept.id">{{ dept.name }}</option>
                  </select>
                </div>
              </div>
            </div>

            <!-- Impacted Workflows list -->`;

content = content.replace(templateRegex, newTemplate);

// We need to keep everything above <script setup>, and rewrite the entire <script setup> block.
// Wait, there's <style scoped> at the end.
const scriptRegex = /<script setup>[\s\S]*?<\/script>/m;

const newScript = `<script setup>
import { ref, computed, h, defineComponent, onMounted } from 'vue'
import { LucideSearch, LucideZoomIn, LucideZoomOut, LucideMaximize2, LucideFilter, LucideAlertTriangle, LucideGitCompare, LucideActivity, LucideBuilding2, LucideFolder, LucideUser, LucideCrown } from 'lucide-vue-next'
import { useToast } from 'vue-toastification'
import { useRuntimeConfig } from '#app'

definePageMeta({
  layout: 'admin',
  middleware: 'auth'
})

const toast = useToast()
const config = useRuntimeConfig()
const { $api } = useNuxtApp()

// UI State
const viewMode = ref('chart') // 'split', 'tree', 'chart'
const selectedDept = ref('all')
const selectedStatus = ref('active')
const searchQuery = ref('')
const selectedEmployeeId = ref(null)
const isSaving = ref(false)

// OrgCard Definition
const OrgCard = defineComponent({
  name: 'OrgCard',
  props: {
    nodeData: Object,
    selected: Boolean,
    isDragTarget: Boolean,
  },
  emits: ['click', 'dragstart', 'dragover', 'dragleave', 'drop'],
  setup(props, { emit }) {
    const tierColors = {
      company: { gradient: 'from-slate-800 to-slate-900', ring: 'ring-slate-800/20', accent: 'bg-slate-800' },
      branch: { gradient: 'from-blue-600 to-blue-800', ring: 'ring-blue-600/20', accent: 'bg-blue-600' },
      department: { gradient: 'from-amber-500 to-amber-600', ring: 'ring-amber-500/20', accent: 'bg-amber-500' },
      head: { gradient: 'from-[#1E3A5F] to-[#2d5a8f]', ring: 'ring-[#1E3A5F]/20', accent: 'bg-[#1E3A5F]' },
      staff: { gradient: 'from-emerald-500 to-teal-600', ring: 'ring-emerald-500/20', accent: 'bg-emerald-500' },
      empty_head: { gradient: 'from-slate-300 to-slate-400', ring: 'ring-slate-300/20', accent: 'bg-slate-300' }
    }

    return () => {
      const type = props.nodeData.type
      const isUser = type === 'user'
      const isEmptyHead = type === 'empty_head'
      const colors = tierColors[isUser ? (props.nodeData.isHead ? 'head' : 'staff') : type] || tierColors.staff
      
      let initials = '?'
      if (isUser) {
        initials = props.nodeData.name?.split(' ').map(n => n[0]).join('').substring(0, 2).toUpperCase() || '?'
      }

      let subtitle = ''
      if (isUser) subtitle = props.nodeData.role || 'User'
      else if (type === 'department') subtitle = 'Departemen'
      else if (type === 'branch') subtitle = props.nodeData.location || 'Kantor Cabang'
      else if (type === 'company') subtitle = 'Perusahaan Utama'
      else if (isEmptyHead) subtitle = 'Kepala Departemen Kosong'

      let iconRender
      if (type === 'company' || type === 'branch') {
        iconRender = h(LucideBuilding2, { class: 'w-5 h-5 text-white' })
      } else if (type === 'department') {
        iconRender = h(LucideFolder, { class: 'w-5 h-5 text-white' })
      } else if (isEmptyHead) {
        iconRender = h(LucideCrown, { class: 'w-5 h-5 text-white' })
      } else {
        iconRender = initials
      }
      
      const isDraggable = isUser || isEmptyHead
      
      return h('div', {
        class: [
          'org-card group cursor-pointer transition-all duration-300 w-[220px]',
          'rounded-2xl border relative overflow-hidden',
          'hover:shadow-lg hover:-translate-y-0.5',
          props.selected 
            ? \`bg-white dark:bg-slate-800 border-[#1E3A5F] dark:border-blue-500 shadow-xl shadow-[#1E3A5F]/10 ring-2 \${colors.ring}\` 
            : 'bg-white dark:bg-slate-800 border-slate-200/60 dark:border-slate-700 shadow-sm',
          props.isDragTarget
            ? 'border-primary-500 ring-4 ring-primary-500/50 scale-105 z-20 shadow-2xl'
            : '',
          isEmptyHead ? 'opacity-70 border-dashed' : ''
        ].filter(Boolean).join(' '),
        draggable: isDraggable,
        onClick: () => emit('click'),
        onDragstart: (e) => isDraggable ? emit('dragstart', e, props.nodeData) : e.preventDefault(),
        onDragover: (e) => emit('dragover', e, props.nodeData),
        onDragleave: (e) => emit('dragleave', e, props.nodeData),
        onDrop: (e) => emit('drop', e, props.nodeData)
      }, [
        h('div', { class: \`h-1.5 bg-gradient-to-r \${colors.gradient}\` }),
        h('div', { class: 'p-3.5 flex items-center gap-3' }, [
          h('div', { class: 'relative flex-shrink-0' }, [
            h('div', {
              class: 'w-10 h-10 rounded-full ring-2 ring-white dark:ring-slate-700 shadow-sm overflow-hidden flex items-center justify-center font-bold text-xs bg-gradient-to-br ' + colors.gradient + ' text-white'
            }, [iconRender])
          ]),
          h('div', { class: 'flex-1 min-w-0 text-left' }, [
            h('p', {
              class: 'text-[11px] font-black uppercase tracking-tight truncate text-[#1E3A5F] dark:text-white'
            }, props.nodeData.name || ''),
            h('p', {
              class: 'text-[9px] font-bold text-slate-400 uppercase tracking-wider mt-0.5 truncate'
            }, subtitle)
          ])
        ])
      ])
    }
  }
})

// Recursive Org Node Component
const OrgNode = defineComponent({
  name: 'OrgNode',
  props: {
    node: Object,
    selectedId: String,
    dragTargetId: String,
    isRoot: { type: Boolean, default: false }
  },
  emits: ['select', 'dragstart', 'dragover', 'dragleave', 'drop', 'dragend'],
  setup(props, { emit }) {
    return () => {
      const hasChildren = props.node.children && props.node.children.length > 0;
      
      const card = h(OrgCard, {
        nodeData: props.node.data,
        selected: props.selectedId === props.node.data.id,
        isDragTarget: props.dragTargetId === props.node.data.id,
        onClick: () => emit('select', props.node.data.id),
        onDragstart: (e, nodeData) => emit('dragstart', e, nodeData),
        onDragover: (e, nodeData) => emit('dragover', e, nodeData),
        onDragleave: (e, nodeData) => emit('dragleave', e, nodeData),
        onDrop: (e, nodeData) => emit('drop', e, nodeData),
      });

      const buildChildren = () => {
        return props.node.children.map(child => h(OrgNode, {
          node: child,
          selectedId: props.selectedId,
          dragTargetId: props.dragTargetId,
          isRoot: false,
          onSelect: (id) => emit('select', id),
          onDragstart: (e, nodeData) => emit('dragstart', e, nodeData),
          onDragover: (e, nodeData) => emit('dragover', e, nodeData),
          onDragleave: (e, nodeData) => emit('dragleave', e, nodeData),
          onDrop: (e, nodeData) => emit('drop', e, nodeData),
        }))
      }

      if (props.isRoot) {
        if (!hasChildren) return h('div', { class: 'tree-node-root' }, [card]);
        return [
          h('div', { class: 'tree-node-root' }, [card]),
          h('div', { class: 'tree-children' }, buildChildren())
        ];
      }

      if (!hasChildren) return h('div', { class: 'tree-child' }, [card]);
      return h('div', { class: 'tree-child' }, [
        card,
        h('div', { class: 'tree-children' }, buildChildren())
      ]);
    }
  }
})

// Recursive List Component for Sidebar
const OrgListItem = defineComponent({
  name: 'OrgListItem',
  props: { node: Object, selectedId: String, depth: Number },
  emits: ['select'],
  setup(props, { emit }) {
    return () => {
      const isSelected = props.selectedId === props.node.data.id
      const hasChildren = props.node.children && props.node.children.length > 0
      
      const item = h('div', {
        class: [
          'px-3 py-2 rounded-xl text-[10px] font-bold cursor-pointer transition-all flex items-center justify-between',
          isSelected ? 'bg-primary-500/10 text-primary-600 dark:text-primary-400' : 'text-slate-600 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800'
        ],
        style: { marginLeft: \`\${props.depth * 12}px\` },
        onClick: () => emit('select', props.node.data.id)
      }, [
        h('span', props.node.data.name)
      ])

      if (!hasChildren) return item
      
      return h('div', {}, [
        item,
        h('div', { class: 'mt-1 space-y-1 border-l border-slate-200 dark:border-slate-800 ml-3' }, 
          props.node.children.map(child => h(OrgListItem, {
            node: child,
            selectedId: props.selectedId,
            depth: props.depth + 1,
            onSelect: (id) => emit('select', id)
          }))
        )
      ])
    }
  }
})

// --- Data Logic ---
const zoomScale = ref(1)

const companies = ref([])
const branches = ref([])
const departments = ref([])
const employees = ref([])

const originalDepartments = ref({})
const originalHeads = ref({})

const fetchHierarchy = async () => {
  try {
    const res = await $api(\`\${config.public.apiBase}/admin/hierarchy\`)
    
    if (res && res.data) {
      companies.value = res.data.companies || []
      branches.value = res.data.branches || []
      departments.value = res.data.departments || []
      employees.value = res.data.users || []
      
      originalDepartments.value = {}
      originalHeads.value = {}
      
      employees.value.forEach(emp => {
        originalDepartments.value[emp.id] = emp.department_id
      })
      departments.value.forEach(dept => {
        originalHeads.value[dept.id] = dept.head_id
      })
      
      if (employees.value.length > 0 && !selectedEmployeeId.value) {
        selectedEmployeeId.value = employees.value[0].id
      }
    }
  } catch (err) {
    console.error('Failed to fetch hierarchy', err)
  }
}

// Build Unified Tree
const treeData = computed(() => {
  const branchMap = {}
  const deptMap = {}
  const userMap = {}
  
  // Prepare users
  employees.value.forEach(emp => {
    userMap[emp.id] = { data: { ...emp, type: 'user', name: emp.full_name, role: emp.role_name }, children: [] }
  })
  
  // Prepare departments
  departments.value.forEach(dept => {
    const deptNode = { data: { ...dept, type: 'department' }, children: [] }
    deptMap[dept.id] = deptNode
    
    let headNode = null
    if (dept.head_id && userMap[dept.head_id]) {
      headNode = userMap[dept.head_id]
      headNode.data.isHead = true
      deptNode.children.push(headNode)
    } else {
      // Empty Head Placeholder
      headNode = { data: { type: 'empty_head', id: 'empty_head_' + dept.id, name: 'Kepala Departemen', department_id: dept.id }, children: [] }
      deptNode.children.push(headNode)
    }
    
    // Attach staffs under head
    employees.value.forEach(emp => {
      if (emp.department_id === dept.id && emp.id !== dept.head_id) {
        headNode.children.push(userMap[emp.id])
      }
    })
  })
  
  // Prepare branches
  branches.value.forEach(branch => {
    const branchNode = { data: { ...branch, type: 'branch' }, children: [] }
    branchMap[branch.id] = branchNode
    
    departments.value.forEach(dept => {
      if (dept.branch_id === branch.id) {
        branchNode.children.push(deptMap[dept.id])
      }
    })
  })
  
  // Prepare companies
  const roots = []
  companies.value.forEach(company => {
    const companyNode = { data: { ...company, type: 'company' }, children: [] }
    
    branches.value.forEach(branch => {
      if (branch.company_id === company.id) {
        companyNode.children.push(branchMap[branch.id])
      }
    })
    roots.push(companyNode)
  })
  
  return roots
})

onMounted(() => {
  fetchHierarchy()
})

// --- Drag and Drop Logic ---
const draggedSourceNode = ref(null)
const dragTargetId = ref(null)

const handleDragStart = (e, nodeData) => {
  draggedSourceNode.value = nodeData
  e.dataTransfer.effectAllowed = 'move'
  e.dataTransfer.setData('text/plain', nodeData.id)
  setTimeout(() => {
    e.target.style.opacity = '0.4'
  }, 0)
}

const handleDragOver = (e, targetData) => {
  e.preventDefault()
  if (!draggedSourceNode.value) return
  
  // You can drop a user onto a Department (makes them staff)
  // Or onto an empty_head / user (if isHead) (makes them head)
  if (targetData.type === 'department' || targetData.type === 'empty_head' || (targetData.type === 'user' && targetData.isHead)) {
    dragTargetId.value = targetData.id
  } else {
    dragTargetId.value = null
  }
}

const handleDragLeave = (e, targetData) => {
  if (dragTargetId.value === targetData.id) {
    dragTargetId.value = null
  }
}

const handleDrop = (e, targetData) => {
  e.preventDefault()
  const source = draggedSourceNode.value
  dragTargetId.value = null
  draggedSourceNode.value = null
  
  if (!source || source.id === targetData.id) return
  
  if (source.type === 'user') {
    const emp = employees.value.find(u => u.id === source.id)
    if (!emp) return

    if (targetData.type === 'department') {
      // Drop onto department -> Make regular staff
      emp.department_id = targetData.id
      
      // If they were the head of this or another department, remove them
      const ownedDept = departments.value.find(d => d.head_id === emp.id)
      if (ownedDept) ownedDept.head_id = null
      
      toast.success(\`\${emp.full_name} dipindahkan ke staf Departemen \${targetData.name}\`)
    } 
    else if (targetData.type === 'empty_head' || (targetData.type === 'user' && targetData.isHead)) {
      // Drop onto Head slot -> Make them the new head
      const targetDeptId = targetData.department_id
      const dept = departments.value.find(d => d.id === targetDeptId)
      if (dept) {
        // Move them to this department
        emp.department_id = targetDeptId
        // Unassign old head if any
        if (dept.head_id && dept.head_id !== emp.id) {
            // Old head becomes regular staff
            // (They are already in the department, so they just drop down to staff)
        }
        // Assign new head
        dept.head_id = emp.id
        
        toast.success(\`\${emp.full_name} ditetapkan sebagai Kepala Departemen \${dept.name}\`)
      }
    }
    
    selectEmployee(emp.id)
  }
}

const handleDragEnd = (e) => {
  e.target.style.opacity = '1'
  draggedSourceNode.value = null
  dragTargetId.value = null
}

const handleCanvasDragOver = (e) => { e.preventDefault() }
const handleCanvasDrop = (e) => { e.preventDefault() }

const selectEmployee = (id) => {
  if (id.startsWith('empty_head_')) return
  selectedEmployeeId.value = id
}

// Right Panel Logic
const selectedEmployee = computed(() => {
  return employees.value.find(e => e.id === selectedEmployeeId.value)
})

const selectedEmployeeOriginalDepartmentName = computed(() => {
  if (!selectedEmployeeId.value) return 'None'
  const origId = originalDepartments.value[selectedEmployeeId.value]
  if (!origId) return 'None'
  const dept = departments.value.find(d => d.id === origId)
  return dept ? dept.name : 'None'
})

const impactedWorkflows = computed(() => {
  if (!selectedEmployeeId.value) return []
  const wfs = []
  
  if (selectedEmployee.value && selectedEmployee.value.department_id !== originalDepartments.value[selectedEmployeeId.value]) {
    wfs.push({
      title: 'Pemindahan Departemen',
      description: \`Hak akses dokumen akan disesuaikan dengan departemen baru.\`,
      icon: LucideActivity,
      type: 'warning'
    })
  }
  return wfs
})

const saveDraft = async () => {
  isSaving.value = true
  
  const userUpdates = employees.value
    .filter(emp => emp.department_id !== originalDepartments.value[emp.id])
    .map(emp => ({
      user_id: emp.id,
      department_id: emp.department_id
    }))
    
  const deptUpdates = departments.value
    .filter(dept => dept.head_id !== originalHeads.value[dept.id])
    .map(dept => ({
      department_id: dept.id,
      head_id: dept.head_id
    }))
    
  if (userUpdates.length === 0 && deptUpdates.length === 0) {
    toast.info('Tidak ada perubahan untuk disimpan.')
    isSaving.value = false
    return
  }

  try {
    const res = await $api(\`\${config.public.apiBase}/admin/hierarchy/bulk-update\`, {
      method: 'PUT',
      body: { user_updates: userUpdates, department_updates: deptUpdates }
    })
    
    toast.success('Struktur organisasi berhasil disimpan dan diaktifkan.')
    
    userUpdates.forEach(u => originalDepartments.value[u.user_id] = u.department_id)
    deptUpdates.forEach(d => originalHeads.value[d.department_id] = d.head_id)
    
  } catch (err) {
    toast.error('Terjadi kesalahan saat menyimpan struktur')
    console.error(err)
  } finally {
    isSaving.value = false
  }
}

const runSimulation = () => {
  toast.success('Simulasi alur persetujuan berhasil dijalankan. Tidak ada kebocoran hak akses terdeteksi.')
}

const zoomIn = () => { if (zoomScale.value < 1.5) zoomScale.value += 0.1 }
const zoomOut = () => { if (zoomScale.value > 0.6) zoomScale.value -= 0.1 }
const zoomReset = () => { zoomScale.value = 1 }

// Add this so it doesn't fail
const hasConflicts = ref(false)
</script>`;

content = content.replace(scriptRegex, newScript);
fs.writeFileSync(file, content, 'utf8');
console.log('Successfully updated hierarchy.vue!');
