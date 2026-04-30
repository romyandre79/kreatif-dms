<template>
  <div class="space-y-8">
    <PageHeader 
      title="Integration Status Monitor" 
      description="Real-time health telemetry for key document management services."
    >
      <template #actions>
        <div class="flex items-center gap-3">
          <button @click="downloadReport" class="flex items-center gap-2 px-4 py-2.5 bg-slate-100 dark:bg-slate-800 hover:bg-slate-200 dark:hover:bg-slate-700 text-slate-700 dark:text-slate-300 rounded-xl border border-slate-200 dark:border-slate-700 transition-all font-bold text-sm">
            <LucideDownload class="w-4 h-4" />
            Download Integration Report
          </button>
          <button @click="openTypeModal" class="flex items-center gap-2 px-6 py-2.5 bg-[#1E3A5F] text-white rounded-xl shadow-lg shadow-blue-900/20 hover:bg-[#152943] transition-all font-bold text-sm active:scale-95">
            <LucidePlus class="w-4 h-4" />
            Add Integration
          </button>
          <button @click="refreshAll" class="flex items-center gap-2 px-4 py-2.5 bg-primary-600 hover:bg-primary-500 text-white rounded-xl shadow-lg shadow-primary-500/20 transition-all font-bold text-sm">
            <LucideRefreshCw class="w-4 h-4" :class="{ 'animate-spin': refreshing }" />
            Refresh All
          </button>
        </div>
      </template>
    </PageHeader>

    <!-- Top Stats Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-4 gap-6">
      <div v-for="stat in topStats" :key="stat.label" class="relative group">
        <div :class="`absolute inset-x-0 -bottom-px h-1 transition-all duration-300 ${stat.color}`"></div>
        <div class="bg-white/80 dark:bg-slate-900/50 backdrop-blur-xl border border-slate-200 dark:border-slate-800 p-6 rounded-2xl">
          <div class="flex items-center justify-between mb-4">
            <span class="text-xs font-black text-slate-400 uppercase tracking-widest">{{ stat.label }}</span>
            <component :is="stat.icon" :class="`w-5 h-5 ${stat.iconColor}`" />
          </div>
          <div class="flex items-end gap-2">
            <span class="text-3xl font-black text-slate-900 dark:text-white tracking-tight">{{ stat.value }}</span>
            <span v-if="stat.suffix" class="text-sm font-bold text-slate-400 mb-1.5">{{ stat.suffix }}</span>
          </div>
          <div class="mt-4 flex items-center gap-2">
            <div :class="`w-2 h-2 rounded-full ${stat.statusColor} animate-pulse`"></div>
            <span :class="`text-xs font-bold uppercase tracking-tighter ${stat.statusTextColor}`">{{ stat.statusText }}</span>
          </div>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
      <!-- Main Integration Table -->
      <div class="lg:col-span-8 bg-white/80 dark:bg-slate-900/50 backdrop-blur-xl border border-slate-200 dark:border-slate-800 rounded-3xl overflow-hidden flex flex-col">
        <div class="p-6 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between bg-slate-50/50 dark:bg-slate-950/20">
          <div class="flex items-center gap-3">
            <h3 class="text-lg font-black text-slate-900 dark:text-white uppercase tracking-tight">Integration Service Table</h3>
            <div class="flex items-center gap-2 px-2.5 py-1 bg-green-500/10 rounded-full border border-green-500/20">
              <div class="w-1.5 h-1.5 rounded-full bg-green-500 animate-pulse"></div>
              <span class="text-[10px] font-black text-green-500 uppercase tracking-widest">Live Telemetry Active</span>
            </div>
          </div>
        </div>
        
        <div class="overflow-x-auto">
          <table class="w-full text-left border-collapse">
            <thead>
              <tr class="text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-100 dark:border-slate-800">
                <th class="px-8 py-4">Service Name</th>
                <th class="px-6 py-4">Host Address</th>
                <th class="px-6 py-4">Status</th>
                <th class="px-6 py-4">Latency</th>
                <th class="px-6 py-4 text-right">Actions</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-50 dark:divide-slate-800/50">
              <tr v-for="service in services" :key="service.name" class="group hover:bg-slate-50 dark:hover:bg-white/5 transition-colors">
                <td class="px-8 py-5">
                  <div class="flex items-center gap-4">
                    <div :class="`w-10 h-10 rounded-xl flex items-center justify-center border transition-all duration-300 ${service.status === 'online' ? 'bg-white dark:bg-slate-800 border-slate-100 dark:border-slate-700 shadow-sm group-hover:scale-110' : 'bg-red-500/5 border-red-500/20'}`">
                      <component :is="service.icon" :class="`w-5 h-5 ${service.status === 'online' ? 'text-primary-500' : 'text-red-500'}`" />
                    </div>
                    <div>
                      <p class="text-sm font-black text-slate-900 dark:text-white leading-tight">{{ service.name }}</p>
                      <p class="text-[10px] font-bold text-slate-400 uppercase tracking-tighter mt-0.5">{{ service.category }}</p>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-5">
                  <span class="text-xs font-mono font-medium text-slate-500 dark:text-slate-400">{{ service.host }}</span>
                </td>
                <td class="px-6 py-5">
                  <div v-if="service.status === 'online'" class="inline-flex items-center px-3 py-1 bg-green-500/10 text-green-600 dark:text-green-400 rounded-full border border-green-500/20 text-[10px] font-black uppercase tracking-widest">
                    Selesai
                  </div>
                  <div v-else-if="service.status === 'warning'" class="inline-flex flex-col items-center px-3 py-1 bg-amber-500/10 text-amber-600 dark:text-amber-400 rounded-lg border border-amber-500/20 text-[10px] font-black uppercase tracking-widest leading-tight text-center">
                    Latency <span>Spike</span>
                  </div>
                  <div v-else class="inline-flex items-center px-3 py-1 bg-red-500/10 text-red-600 dark:text-red-400 rounded-full border border-red-500/20 text-[10px] font-black uppercase tracking-widest">
                    Gagal
                  </div>
                </td>
                <td class="px-6 py-5">
                  <div class="flex items-center gap-3">
                    <div class="flex flex-col">
                      <span class="text-[10px] font-black text-slate-900 dark:text-white">{{ service.lastLatency }}ms</span>
                      <div class="w-16 h-4 flex items-end gap-0.5 mt-1">
                        <div v-for="(v, i) in service.latencyHistory" :key="i" 
                          class="w-1.5 rounded-full transition-all duration-500" 
                          :class="[v > 200 ? 'bg-red-500' : v > 100 ? 'bg-amber-500' : 'bg-green-500']"
                          :style="`height: ${Math.min(v/5, 100)}%`"
                        ></div>
                      </div>
                    </div>
                    <div :class="`w-1.5 h-1.5 rounded-full ${service.status === 'online' ? 'bg-green-500' : 'bg-red-500'}`"></div>
                  </div>
                </td>
                <td class="px-6 py-5 text-right">
                  <div class="flex items-center justify-end gap-2">
                    <button @click="openModal(service.raw)" class="p-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-lg text-slate-400 hover:text-blue-500 transition-colors">
                      <LucideEdit3 class="w-4 h-4" />
                    </button>
                    <button @click="deleteNode(service.id)" class="p-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-lg text-slate-400 hover:text-red-500 transition-colors">
                      <LucideTrash2 class="w-4 h-4" />
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Right Column: Alert Feed & Trends -->
      <div class="lg:col-span-4 space-y-8">
        <!-- Live Alert Feed -->
        <div class="bg-white/80 dark:bg-slate-900/50 backdrop-blur-xl border border-slate-200 dark:border-slate-800 rounded-3xl p-6">
          <div class="flex items-center justify-between mb-6">
            <div class="flex items-center gap-2">
              <LucideActivity class="w-4 h-4 text-red-500 animate-pulse" />
              <h3 class="text-sm font-black text-slate-900 dark:text-white uppercase tracking-widest">Live Alert Feed</h3>
            </div>
            <span class="text-[10px] font-mono font-bold text-slate-400">T+ 0.5s</span>
          </div>
          
          <div class="space-y-4">
            <div v-for="alert in alerts" :key="alert.id" 
              class="relative pl-6 py-4 pr-4 rounded-2xl overflow-hidden transition-all hover:scale-[1.02] cursor-default"
              :class="alert.bg"
            >
              <div :class="`absolute left-0 top-0 bottom-0 w-1.5 ${alert.accent}`"></div>
              <div class="flex gap-4">
                <LucideAlertTriangle v-if="alert.type === 'error'" class="w-5 h-5 text-red-500 flex-shrink-0 mt-0.5" />
                <LucideTimer v-else-if="alert.type === 'warning'" class="w-5 h-5 text-amber-500 flex-shrink-0 mt-0.5" />
                <LucideCheckCircle2 v-else class="w-5 h-5 text-blue-500 flex-shrink-0 mt-0.5" />
                
                <div class="flex-1 min-w-0">
                  <h4 class="text-xs font-black uppercase tracking-widest mb-1" :class="alert.titleColor">{{ alert.title }}</h4>
                  <p class="text-xs font-medium text-slate-500 dark:text-slate-400 leading-relaxed">{{ alert.message }}</p>
                  <div class="mt-3 flex items-center justify-between">
                    <span class="text-[10px] font-bold text-slate-400 tracking-widest">{{ alert.time }}</span>
                    <span class="text-[10px] font-mono font-bold text-slate-300 dark:text-slate-600 tracking-tighter">{{ alert.code }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="bg-white/80 dark:bg-slate-900/50 backdrop-blur-xl border border-slate-200 dark:border-slate-800 rounded-3xl p-6">
          <h3 class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-6">Latency Trend (Recent Checks)</h3>
          <div class="h-40 flex items-end justify-between gap-1 px-2">
            <div v-for="(v, i) in globalTrend" :key="i" 
              class="flex-1 rounded-t-sm transition-all duration-700 hover:opacity-80" 
              :class="[v > 70 ? 'bg-red-500/40 border-t-2 border-red-500' : i === globalTrend.length - 1 ? 'bg-primary-500 shadow-[0_0_15px_rgba(59,130,246,0.5)]' : 'bg-slate-700/40']"
              :style="`height: ${Math.min(v, 100)}%`"
            ></div>
          </div>
        </div>
      </div>
    </div>

    <!-- Type Selection Modal -->
    <Transition name="fade">
      <div v-if="showTypeModal" class="fixed inset-0 z-[60] flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-slate-950/40 backdrop-blur-sm" @click="showTypeModal = false"></div>
        <div class="relative bg-white dark:bg-slate-900 w-full max-w-2xl rounded-[2.5rem] shadow-2xl p-10 border border-slate-200 dark:border-slate-800" v-motion-pop>
          <div class="text-center space-y-2 mb-10">
            <h3 class="text-2xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Select Service Type</h3>
            <p class="text-slate-400 font-bold text-xs uppercase tracking-widest">Choose the type of integration you want to add</p>
          </div>

          <div class="grid grid-cols-2 sm:grid-cols-3 gap-6">
            <button v-for="type in integrationTypes" :key="type.id" 
              @click="selectType(type.id)"
              class="group p-6 rounded-3xl border border-slate-100 dark:border-slate-800 bg-slate-50/50 dark:bg-slate-800/30 hover:bg-white dark:hover:bg-slate-800 hover:border-blue-200 hover:shadow-xl hover:shadow-blue-900/5 transition-all text-center space-y-4"
            >
              <div class="w-12 h-12 rounded-2xl bg-white dark:bg-slate-700 shadow-sm border border-slate-50 dark:border-slate-600 flex items-center justify-center mx-auto text-slate-400 group-hover:text-blue-500 transition-colors">
                <component :is="type.icon" class="w-6 h-6" />
              </div>
              <p class="text-[10px] font-black text-slate-500 group-hover:text-[#1E3A5F] uppercase tracking-widest transition-colors">{{ type.name }}</p>
            </button>
          </div>
          
          <button @click="showTypeModal = false" class="mt-10 w-full py-4 text-[10px] font-black text-slate-400 hover:text-slate-600 uppercase tracking-widest transition-colors">
            Cancel
          </button>
        </div>
      </div>
    </Transition>

    <!-- Configuration Modal (Dynamic) -->
    <Transition name="fade">
      <div v-if="showModal" class="fixed inset-0 z-50 flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-slate-950/60 backdrop-blur-sm" @click="showModal = false"></div>
        <div class="relative bg-white dark:bg-slate-950 w-full max-w-6xl rounded-[2.5rem] shadow-2xl overflow-hidden border border-slate-200 dark:border-slate-800" v-motion-pop>
          <!-- Modal Header -->
          <div class="px-10 py-8 bg-slate-50 dark:bg-slate-900/50 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between">
            <div class="space-y-1">
              <h2 class="text-2xl font-black text-[#1E3A5F] dark:text-white tracking-tight uppercase">
                {{ isEdit ? 'Edit Integration' : 'Add New Integration' }}
              </h2>
              <p class="text-slate-500 font-bold text-xs uppercase tracking-widest">Service Node & Connection Parameters</p>
            </div>
            <div class="flex items-center gap-4">
              <button 
                v-if="formData.service_type === 'LDAP'" 
                @click="viewLogs"
                class="flex items-center gap-2 px-4 py-2 text-[10px] font-black uppercase tracking-widest text-slate-500 hover:text-[#1E3A5F] transition-colors border border-transparent hover:border-slate-200 rounded-xl"
              >
                <LucideHistory class="w-4 h-4" />
                View Error Log
              </button>
              <button 
                v-if="formData.service_type === 'LDAP'" 
                @click="testConnection"
                :disabled="testing"
                class="flex items-center gap-2 px-4 py-2 text-[10px] font-black uppercase tracking-widest text-slate-500 hover:text-[#1E3A5F] transition-colors border border-transparent hover:border-slate-200 rounded-xl disabled:opacity-50"
              >
                <LucideCheckSquare class="w-4 h-4" :class="{ 'animate-spin': testing }" />
                {{ testing ? 'Testing...' : 'Test Connection' }}
              </button>
              <button @click="saveNode" :disabled="loading" class="px-8 py-3 bg-[#1E3A5F] text-white rounded-xl text-xs font-black uppercase tracking-widest flex items-center gap-3 shadow-xl shadow-blue-900/20 hover:bg-[#152943] transition-all disabled:opacity-50">
                <LucideSave class="w-4 h-4" />
                {{ loading ? 'Saving...' : 'Save Configuration' }}
              </button>
              <button @click="showModal = false" class="p-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl transition-colors">
                <LucidePlus class="w-6 h-6 rotate-45 text-slate-400" />
              </button>
            </div>
          </div>

          <!-- Modal Body -->
          <div class="p-8 max-h-[70vh] overflow-y-auto custom-scrollbar bg-slate-50/30">
            <!-- General Node Info -->
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8 mb-8">
              <div class="space-y-2">
                <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-1">Node Name</label>
                <input v-model="formData.name" type="text" class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-blue-500/10 focus:border-[#1E3A5F] transition-all" placeholder="e.g. Primary LDAP">
              </div>
              <div class="space-y-2">
                <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-1">Service Type</label>
                <select v-model="formData.service_type" class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-blue-500/10 focus:border-[#1E3A5F] transition-all">
                  <option v-for="type in integrationTypes" :key="type.id" :value="type.id">
                    {{ type.name }}
                  </option>
                </select>
              </div>
              <div class="space-y-2">
                <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-1">Connection String (DSN/URL)</label>
                <input v-model="formData.config.connection_string" type="text" class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-blue-500/10 focus:border-[#1E3A5F] transition-all" placeholder="e.g. ldap://10.20.30.144:389">
              </div>
            </div>

            <div v-if="formData.service_type === 'LDAP'" class="grid grid-cols-1 gap-8">
                    <div class="flex items-center gap-3">
                      <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">LDAP ACTIVE</p>
                      <button @click="formData.is_active = !formData.is_active" :class="`w-12 h-6 rounded-full transition-all relative ${formData.is_active ? 'bg-[#1E3A5F]' : 'bg-slate-300'}`">
                        <div :class="`absolute top-1 w-4 h-4 bg-white rounded-full transition-all ${formData.is_active ? 'left-7' : 'left-1'}`"></div>
                      </button>
                    </div>
                    <div class="space-y-3">
                      <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Provider Type</label>
                      <div class="grid grid-cols-2 gap-4">
                        <button @click="formData.driver = 'active_directory'" :class="`py-3 rounded-xl text-xs font-black uppercase tracking-widest border transition-all flex items-center justify-center gap-2 ${formData.driver === 'active_directory' ? 'bg-blue-50 border-blue-200 text-blue-600' : 'bg-white border-slate-100 text-slate-400'}`">
                          <LucideServer class="w-4 h-4" /> Active Directory
                        </button>
                        <button @click="formData.driver = 'openldap'" :class="`py-3 rounded-xl text-xs font-black uppercase tracking-widest border transition-all flex items-center justify-center gap-2 ${formData.driver === 'openldap' ? 'bg-blue-50 border-blue-200 text-blue-600' : 'bg-white border-slate-100 text-slate-400'}`">
                          <LucideHash class="w-4 h-4" /> OpenLDAP
                        </button>
                      </div>
                    </div>
                    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                      <div class="space-y-2">
                        <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Timeout (s)</label>
                        <input v-model="formData.config.timeout" type="text" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-bold" placeholder="30">
                      </div>
                      <div class="space-y-2">
                        <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Bind DN</label>
                        <input v-model="formData.config.bind_dn" type="text" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-bold" placeholder="cn=admin,dc=example,dc=com">
                      </div>
                      <div class="space-y-2">
                        <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Password</label>
                        <div class="relative">
                          <input :type="showPassword ? 'text' : 'password'" v-model="formData.config.bind_password" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-bold pr-14">
                          <button @click="showPassword = !showPassword" class="absolute right-4 top-1/2 -translate-y-1/2 text-slate-400">
                            <LucideEye v-if="!showPassword" class="w-4 h-4" />
                            <LucideEyeOff v-else class="w-4 h-4" />
                          </button>
                        </div>
                      </div>
                    </div>
            </div>

            <div v-else-if="formData.service_type === 'AI'" class="grid grid-cols-1 gap-8">
              <div class="space-y-6">
                <div class="space-y-3">
                  <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">AI Provider</label>
                  <div class="grid grid-cols-2 gap-4">
                    <button @click="formData.driver = 'gemini'; formData.config.model = 'gemini-1.5-flash'" :class="`py-3 rounded-xl text-xs font-black uppercase tracking-widest border transition-all flex items-center justify-center gap-2 ${formData.driver === 'gemini' ? 'bg-blue-50 border-blue-200 text-blue-600' : 'bg-white border-slate-100 text-slate-400'}`">
                      <LucideCpu class="w-4 h-4" /> Google Gemini
                    </button>
                    <button @click="formData.driver = 'openai'; formData.config.model = 'gpt-4o'" :class="`py-3 rounded-xl text-xs font-black uppercase tracking-widest border transition-all flex items-center justify-center gap-2 ${formData.driver === 'openai' ? 'bg-blue-50 border-blue-200 text-blue-600' : 'bg-white border-slate-100 text-slate-400'}`">
                      <LucideCpu class="w-4 h-4" /> OpenAI
                    </button>
                  </div>
                </div>

                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div class="space-y-2">
                    <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">API Key</label>
                    <div class="relative">
                      <input :type="showPassword ? 'text' : 'password'" v-model="formData.config.token" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-bold pr-14" placeholder="Enter API Key">
                      <button @click="showPassword = !showPassword" class="absolute right-4 top-1/2 -translate-y-1/2 text-slate-400">
                        <LucideEye v-if="!showPassword" class="w-4 h-4" />
                        <LucideEyeOff v-else class="w-4 h-4" />
                      </button>
                    </div>
                  </div>
                  <div class="space-y-2">
                    <div class="flex items-center justify-between">
                      <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Model Name</label>
                      <button @click="fetchDynamicModels" :disabled="fetchingModels" class="text-[9px] font-black text-blue-500 uppercase tracking-widest hover:text-blue-600 disabled:opacity-50 flex items-center gap-1">
                        <LucideRefreshCw class="w-3 h-3" :class="{ 'animate-spin': fetchingModels }" />
                        Fetch from API
                      </button>
                    </div>
                    <select v-model="formData.config.model" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-bold focus:ring-2 focus:ring-blue-500/20 transition-all">
                      <option v-for="m in aiModels[formData.driver] || []" :key="m.id" :value="m.id">
                        {{ m.name }}
                      </option>
                    </select>
                  </div>
                </div>

                <div class="space-y-2">
                  <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">System Instruction / Prompt</label>
                  <textarea v-model="formData.config.system_prompt" rows="4" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-medium leading-relaxed resize-none" placeholder="Default system prompt for AI responses..."></textarea>
                </div>
              </div>
            </div>

            <div v-else-if="formData.service_type === 'SEARCH'" class="grid grid-cols-1 gap-8">
              <div class="space-y-6">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div class="space-y-2">
                    <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Auth Method</label>
                    <div class="grid grid-cols-2 gap-4">
                      <button @click="formData.driver = 'basic'" :class="`py-3 rounded-xl text-xs font-black uppercase tracking-widest border transition-all ${formData.driver === 'basic' ? 'bg-blue-50 border-blue-200 text-blue-600' : 'bg-white border-slate-100 text-slate-400'}`">Basic Auth</button>
                      <button @click="formData.driver = 'apikey'" :class="`py-3 rounded-xl text-xs font-black uppercase tracking-widest border transition-all ${formData.driver === 'apikey' ? 'bg-blue-50 border-blue-200 text-blue-600' : 'bg-white border-slate-100 text-slate-400'}`">API Key</button>
                    </div>
                  </div>
                  <div class="space-y-2">
                    <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Index Name</label>
                    <input type="text" v-model="formData.config.index_name" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-bold" placeholder="e.g. documents">
                  </div>
                </div>

                <div v-if="formData.driver === 'basic'" class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div class="space-y-2">
                    <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Username</label>
                    <input type="text" v-model="formData.config.username" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-bold" placeholder="elastic">
                  </div>
                  <div class="space-y-2">
                    <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Password</label>
                    <div class="relative">
                      <input :type="showPassword ? 'text' : 'password'" v-model="formData.config.password" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-bold pr-14">
                      <button @click="showPassword = !showPassword" class="absolute right-4 top-1/2 -translate-y-1/2 text-slate-400">
                        <LucideEye v-if="!showPassword" class="w-4 h-4" />
                        <LucideEyeOff v-else class="w-4 h-4" />
                      </button>
                    </div>
                  </div>
                </div>

                <div v-if="formData.driver === 'apikey'" class="space-y-2">
                  <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Elastic API Key</label>
                  <input type="text" v-model="formData.config.api_key" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-bold" placeholder="Enter Base64 API Key">
                </div>
              </div>
            </div>

            <div v-else-if="formData.service_type === 'S3' || formData.service_type === 'STORAGE'" class="grid grid-cols-1 gap-8">
              <div class="space-y-6">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div class="space-y-2">
                    <label class="text-[10px] font-black text-slate-400 uppercase tracking-widest px-1">Bucket Name</label>
                    <input v-model="formData.config.bucket" type="text" class="w-full px-5 py-3.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-xl text-sm font-bold focus:ring-4 focus:ring-blue-500/10 focus:border-[#1E3A5F] transition-all" placeholder="e.g. krdms-vault">
                  </div>
                  <div class="flex items-center gap-3 pt-8">
                    <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Use SSL (HTTPS)</p>
                    <button @click="formData.config.use_ssl = !formData.config.use_ssl" :class="`w-12 h-6 rounded-full transition-all relative ${formData.config.use_ssl ? 'bg-[#1E3A5F]' : 'bg-slate-300'}`">
                      <div :class="`absolute top-1 w-4 h-4 bg-white rounded-full transition-all ${formData.config.use_ssl ? 'left-7' : 'left-1'}`"></div>
                    </button>
                  </div>
                </div>

                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div class="space-y-2">
                    <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Access Key</label>
                    <input type="text" v-model="formData.config.access_key" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-bold" placeholder="Enter Access Key">
                  </div>
                  <div class="space-y-2">
                    <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Secret Key</label>
                    <div class="relative">
                      <input :type="showPassword ? 'text' : 'password'" v-model="formData.config.secret_key" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-bold pr-14" placeholder="Enter Secret Key">
                      <button @click="showPassword = !showPassword" class="absolute right-4 top-1/2 -translate-y-1/2 text-slate-400">
                        <LucideEye v-if="!showPassword" class="w-4 h-4" />
                        <LucideEyeOff v-else class="w-4 h-4" />
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <div v-else-if="formData.service_type === 'WHATSAPP'" class="grid grid-cols-1 gap-8">
              <div class="space-y-6">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div class="space-y-2">
                    <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">WhatsApp Provider</label>
                    <select v-model="formData.driver" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-bold focus:ring-2 focus:ring-blue-500/20 transition-all">
                      <option value="fonnte">Fonnte.com</option>
                      <option value="wablas">Wablas.com</option>
                      <option value="generic">Generic Webhook / API</option>
                    </select>
                  </div>
                </div>

                <div v-if="formData.driver === 'fonnte' || formData.driver === 'wablas'" class="space-y-6">
                  <div class="p-6 bg-blue-50 dark:bg-blue-900/20 rounded-2xl border border-blue-100 dark:border-blue-800/30 flex items-start gap-4">
                    <div class="p-3 bg-white dark:bg-slate-800 rounded-xl shadow-sm">
                      <LucideMail class="w-6 h-6 text-blue-600" />
                    </div>
                    <div class="space-y-1">
                      <h4 class="text-sm font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">{{ formData.driver === 'fonnte' ? 'Fonnte' : 'Wablas' }} Integration</h4>
                      <p class="text-xs text-slate-500 leading-relaxed">Enter your API Token to enable automated WhatsApp notifications via {{ formData.driver === 'fonnte' ? 'Fonnte' : 'Wablas' }} gateway.</p>
                    </div>
                  </div>

                  <div class="space-y-2">
                    <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">API Token / API Key</label>
                    <div class="relative">
                      <input :type="showPassword ? 'text' : 'password'" v-model="formData.config.token" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-bold pr-14" placeholder="Enter your API Token">
                      <button @click="showPassword = !showPassword" class="absolute right-4 top-1/2 -translate-y-1/2 text-slate-400">
                        <LucideEye v-if="!showPassword" class="w-4 h-4" />
                        <LucideEyeOff v-else class="w-4 h-4" />
                      </button>
                    </div>
                  </div>
                </div>

                <div v-if="formData.driver === 'generic'" class="space-y-6">
                  <div class="space-y-2">
                    <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Custom API Endpoint</label>
                    <input type="text" v-model="formData.endpoint" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-bold" placeholder="https://api.yourprovider.com/send">
                  </div>
                  <div class="space-y-2">
                    <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Auth Header (Optional)</label>
                    <input type="text" v-model="formData.config.auth_header" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-bold" placeholder="e.g. Bearer your_token">
                  </div>
                </div>
              </div>
            </div>

            <div v-else-if="formData.service_type === 'SMTP'" class="grid grid-cols-1 gap-8">
              <div class="space-y-6">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div class="space-y-2">
                    <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Sender Email</label>
                    <input type="email" v-model="formData.config.from_email" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-bold" placeholder="noreply@company.com">
                  </div>
                  <div class="flex items-center gap-3 pt-8">
                    <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Enable Authentication</p>
                    <button @click="formData.config.auth = !formData.config.auth" :class="`w-12 h-6 rounded-full transition-all relative ${formData.config.auth ? 'bg-[#1E3A5F]' : 'bg-slate-300'}`">
                      <div :class="`absolute top-1 w-4 h-4 bg-white rounded-full transition-all ${formData.config.auth ? 'left-7' : 'left-1'}`"></div>
                    </button>
                  </div>
                </div>

                <div v-if="formData.config.auth" class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div class="space-y-2">
                    <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">SMTP Username</label>
                    <input type="text" v-model="formData.config.user" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-bold" placeholder="Enter username">
                  </div>
                  <div class="space-y-2">
                    <label class="text-[9px] font-black text-slate-400 uppercase tracking-widest">SMTP Password</label>
                    <div class="relative">
                      <input :type="showPassword ? 'text' : 'password'" v-model="formData.config.pass" class="w-full px-5 py-3.5 bg-slate-50 dark:bg-slate-800 border-none rounded-xl text-sm font-bold pr-14" placeholder="Enter password">
                      <button @click="showPassword = !showPassword" class="absolute right-4 top-1/2 -translate-y-1/2 text-slate-400">
                        <LucideEye v-if="!showPassword" class="w-4 h-4" />
                        <LucideEyeOff v-else class="w-4 h-4" />
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <!-- Modal Footer Stats -->
          <div class="px-10 py-6 bg-slate-50 dark:bg-slate-900 border-t border-slate-100 dark:border-slate-800 flex items-center justify-between">
            <div class="flex items-center gap-10">
              <div class="space-y-0.5">
                <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest">Last Sync Activity</p>
                <p class="text-[11px] font-black text-[#1E3A5F] dark:text-white tracking-tight">{{ syncStats.lastSync || 'Never' }}</p>
              </div>
              <div v-if="formData.service_type === 'LDAP'" class="flex items-center gap-6">
                <div class="text-center">
                  <p class="text-lg font-black text-[#1E3A5F] dark:text-white">{{ syncStats.users.toLocaleString() }}</p>
                  <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">Users Synced</p>
                </div>
                <div class="text-center border-l border-slate-200 dark:border-slate-800 pl-6">
                  <p class="text-lg font-black text-[#1E3A5F] dark:text-white">{{ syncStats.groups.toLocaleString() }}</p>
                  <p class="text-[8px] font-black text-slate-400 uppercase tracking-widest">Groups Map</p>
                </div>
              </div>
            </div>
            <div class="flex items-center gap-3">
              <button @click="showModal = false" class="px-8 py-3 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-slate-500 rounded-xl text-xs font-black uppercase tracking-widest hover:bg-slate-50 transition-all">
                Cancel
              </button>
            </div>
          </div>
        </div>
      </div>
    </Transition>

    <!-- Sync Logs Modal -->
    <Transition name="fade">
      <div v-if="showLogsModal" class="fixed inset-0 z-[70] flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-slate-950/60 backdrop-blur-sm" @click="showLogsModal = false"></div>
        <div class="relative bg-white dark:bg-slate-900 w-full max-w-4xl rounded-[2.5rem] shadow-2xl overflow-hidden border border-slate-200 dark:border-slate-800" v-motion-pop>
          <div class="px-10 py-8 bg-slate-50 dark:bg-slate-950/50 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between">
            <div class="space-y-1">
              <h3 class="text-xl font-black text-[#1E3A5F] dark:text-white uppercase tracking-tight">Sync & Connection Logs</h3>
              <p class="text-slate-400 font-bold text-[10px] uppercase tracking-widest">Audit trail for service authentication events</p>
            </div>
            <button @click="showLogsModal = false" class="p-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-xl transition-colors">
              <LucidePlus class="w-5 h-5 rotate-45 text-slate-400" />
            </button>
          </div>
          
          <div class="p-0 max-h-[60vh] overflow-y-auto custom-scrollbar">
            <table class="w-full text-left border-collapse">
              <thead>
                <tr class="text-[9px] font-black text-slate-400 uppercase tracking-widest bg-slate-50/50 dark:bg-slate-950/20 border-b border-slate-100 dark:border-slate-800">
                  <th class="px-8 py-4">Timestamp</th>
                  <th class="px-6 py-4">Status</th>
                  <th class="px-6 py-4">Users</th>
                  <th class="px-6 py-4">Groups</th>
                  <th class="px-6 py-4">Details</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-slate-50 dark:divide-slate-800/50">
                <tr v-for="log in syncLogs" :key="log.id" class="text-xs">
                  <td class="px-8 py-4 font-bold text-slate-500 whitespace-nowrap">
                    {{ formatDate(new Date(log.created_at)) }}
                  </td>
                  <td class="px-6 py-4">
                    <span :class="`px-2 py-1 rounded-full text-[9px] font-black uppercase tracking-widest ${log.status === 'success' ? 'bg-green-500/10 text-green-500' : 'bg-red-500/10 text-red-500'}`">
                      {{ log.status }}
                    </span>
                  </td>
                  <td class="px-6 py-4 font-black text-slate-700 dark:text-slate-300">{{ log.users_synced || 0 }}</td>
                  <td class="px-6 py-4 font-black text-slate-700 dark:text-slate-300">{{ log.groups_synced || 0 }}</td>
                  <td class="px-6 py-4">
                    <p v-if="log.status === 'failed'" class="text-red-500 font-medium truncate max-w-[200px]" :title="decodeError(log.error_details)">
                      {{ decodeError(log.error_details) }}
                    </p>
                    <p v-else class="text-slate-400">Normal connection</p>
                  </td>
                </tr>
                <tr v-if="syncLogs.length === 0">
                  <td colspan="5" class="px-8 py-20 text-center">
                    <div class="flex flex-col items-center gap-4">
                      <LucideHistory class="w-12 h-12 text-slate-200" />
                      <p class="text-sm font-bold text-slate-400 uppercase tracking-widest">No logs found for this node</p>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
          
          <div class="px-10 py-6 bg-slate-50 dark:bg-slate-950/50 border-t border-slate-100 dark:border-slate-800 text-right">
            <button @click="showLogsModal = false" class="px-6 py-2.5 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl text-[10px] font-black uppercase tracking-widest text-slate-600 dark:text-slate-400 hover:bg-slate-50 transition-all">
              Close Logs
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>


<script setup>
import { ref, onMounted, onUnmounted, computed } from 'vue'
import { useRuntimeConfig } from '#app'
import { 
  LucideDownload, LucideRefreshCw, LucideActivity, LucideServer, 
  LucideDatabase, LucideHardDrive, LucideScan, LucidePrinter, 
  LucideHash, LucideMail, LucideAlertTriangle, LucideCheckCircle2, 
  LucideTimer, LucidePlus, LucideEdit3, LucideTrash2, LucideSave, 
  LucideEye, LucideEyeOff, LucideCheckSquare, LucideNetwork, 
  LucideShare2, LucideFilter, LucideHistory, LucidePlay, LucideSettings,
  LucideSearch, LucideCpu
} from 'lucide-vue-next'
import { useApi } from '~/composables/useApi'

const refreshing = ref(false)
const nodes = ref([])
const { $api } = useApi()
const config = useRuntimeConfig()

// Modal & Form State
const showTypeModal = ref(false)
const showModal = ref(false)
const isEdit = ref(false)
const loading = ref(false)
const testing = ref(false)
const showPassword = ref(false)
const showLogsModal = ref(false)
const syncLogs = ref([])
const fetchingModels = ref(false)
const totalSyncCount = ref(0)

const syncStats = ref({
  users: 0,
  groups: 0,
  lastSync: ''
})

const formatDate = (date) => {
  return new Intl.DateTimeFormat('id-ID', {
    day: '2-digit',
    month: 'short',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
    timeZoneName: 'short'
  }).format(date)
}

const integrationTypes = [
  { id: 'LDAP', name: 'LDAP / AD', icon: LucideNetwork },
  { id: 'SEARCH', name: 'Search Engine', icon: LucideSearch },
  { id: 'WHATSAPP', name: 'WhatsApp', icon: LucideMail },
  { id: 'OCR', name: 'OCR Engine', icon: LucideScan },
  { id: 'STORAGE', name: 'Storage', icon: LucideHardDrive },
  { id: 'SMTP', name: 'SMTP / Mail', icon: LucidePrinter },
  { id: 'AI', name: 'AI Provider', icon: LucideCpu }
]

const aiModels = ref({
  gemini: [
    { id: 'gemini-1.5-pro', name: 'Gemini 1.5 Pro (Powerful)' },
    { id: 'gemini-1.5-flash', name: 'Gemini 1.5 Flash (Fast)' },
    { id: 'gemini-1.0-pro', name: 'Gemini 1.0 Pro' }
  ],
  openai: [
    { id: 'gpt-4o', name: 'GPT-4o (Omni)' },
    { id: 'gpt-4-turbo', name: 'GPT-4 Turbo' },
    { id: 'gpt-3.5-turbo', name: 'GPT-3.5 Turbo' }
  ]
})

const initialForm = {
  name: '',
  service_type: 'LDAP',
  driver: 'active_directory',
  endpoint: '',
  is_active: true,
  is_critical: true,
  config: {
    connection_string: '',
    timeout: '30',
    bind_dn: '',
    bind_password: '',
    anonymous_bind: false,
    sync_lang: 'id',
    inclusion_filter: '',
    auto_disable: false,
    default_role: 'viewer',
    sync_frequency: '6 Jam',
    // AI Config
    token: '',
    model: 'gemini-1.5-flash',
    system_prompt: 'You are a helpful assistant for Kreatif DMS.',
    // SEARCH Config
    username: '',
    password: '',
    api_key: '',
    index_name: 'documents',
    // STORAGE Config
    access_key: '',
    secret_key: '',
    bucket: '',
    use_ssl: true,
    // WHATSAPP Config
    // token is already defined for AI, we can reuse it or define separately
    // but in DB it's stored in config_json.
    // SMTP Config
    from_email: 'noreply@kreatif-dms.com',
    auth: true,
    user: '',
    pass: ''
  }
}

const formData = ref(JSON.parse(JSON.stringify(initialForm)))

const fetchNodes = async () => {
  try {
    refreshing.value = true
    const res = await $api(`${config.public.apiBase}/master/integration/status`)
    if (res && res.data) {
      nodes.value = res.data.nodes || []
      totalSyncCount.value = res.data.sync_count || 0
    }
  } catch (err) {
    console.error('Failed to fetch integration status:', err)
  } finally {
    refreshing.value = false
  }
}

const openTypeModal = () => {
  showTypeModal.value = true
}

const selectType = (typeId) => {
  showTypeModal.value = false
  openModal({ service_type: typeId })
}

const openModal = (node = null) => {
  if (node && node.id) {
    isEdit.value = true
    
    // Decode base64 config_json if necessary
    let configData = node.config_json || {}
    if (typeof configData === 'string') {
      try {
        configData = JSON.parse(atob(configData))
      } catch (e) {
        console.error('Failed to decode config_json:', e)
        configData = {}
      }
    }

    // Handle pgtype.Text or plain string for driver
    console.log('DEBUG: node.driver RAW:', node.driver)
    let driverValue = ''
    if (node.driver) {
      driverValue = typeof node.driver === 'object' ? (node.driver.String || '') : node.driver
    }

    // Set smart defaults if driver is empty based on service type
    if (!driverValue || driverValue === '') {
      if (node.service_type === 'WHATSAPP') driverValue = 'fonnte'
      else if (node.service_type === 'AI') driverValue = 'gemini'
      else if (node.service_type === 'SEARCH') driverValue = 'basic'
      else if (node.service_type === 'LDAP') driverValue = 'active_directory'
    }

    // Map legacy driver names
    if (driverValue === 'wagateway') driverValue = 'fonnte'
    
    console.log('DEBUG: driverValue FINAL:', driverValue)

    formData.value = {
      id: node.id,
      name: node.name,
      service_type: node.service_type,
      driver: driverValue,
      endpoint: node.endpoint,
      is_active: node.is_active,
      is_critical: node.is_critical,
      config: { 
        ...JSON.parse(JSON.stringify(initialForm.config)), 
        ...configData,
        // Map connection_string from endpoint if missing in config
        connection_string: configData.connection_string || node.endpoint || '',
        // Map password if source uses bind_pass
        bind_password: configData.bind_password || configData.bind_pass || ''
      }
    }
    console.log('DECODED CONFIG:', configData)
    console.log('FINAL FORM CONFIG:', formData.value.config)

    // Load stats if exist
    syncStats.value = {
      users: configData.users_count || 0,
      groups: configData.groups_count || 0,
      lastSync: node.last_check_at ? formatDate(new Date(node.last_check_at)) : ''
    }
  } else {
    isEdit.value = false
    syncStats.value = { users: 0, groups: 0, lastSync: '' }
    const baseForm = JSON.parse(JSON.stringify(initialForm))
    if (node && node.service_type) {
      baseForm.service_type = node.service_type
    }
    formData.value = baseForm
  }
  showModal.value = true
}

const testConnection = async () => {
  testing.value = true
  try {
    const res = await $api(`${config.public.apiBase}/master/integration/test-connection`, {
      method: 'POST',
      body: {
        service_type: formData.value.service_type,
        endpoint: formData.value.endpoint,
        config_json: formData.value.config
      }
    })
    
    if (res && res.data) {
      // Update stats from backend
      syncStats.value.users = res.data.users || 0
      syncStats.value.groups = res.data.groups || 0
      syncStats.value.lastSync = formatDate(new Date())
      
      // Update formData for persistence
      formData.value.config.users_count = res.data.users
      formData.value.config.groups_count = res.data.groups
      
      alert('Connection Successful!')
    }
  } catch (err) {
    console.error('Test Connection Error:', err)
    alert('Connection Failed: ' + (err.message || 'Unknown error'))
  } finally {
    testing.value = false
  }
}

const fetchDynamicModels = async () => {
  if (!formData.value.config.token) return
  
  try {
    fetchingModels.value = true
    const res = await $api(`${config.public.apiBase}/master/integration/ai-models`, {
      params: {
        driver: formData.value.driver,
        api_key: formData.value.config.token
      }
    })
    
    if (res && res.data) {
      aiModels.value[formData.value.driver] = res.data
      // If current model not in new list, pick first
      if (res.data.length > 0 && !res.data.find(m => m.id === formData.value.config.model)) {
        formData.value.config.model = res.data[0].id
      }
    }
  } catch (err) {
    console.error('Failed to fetch models:', err)
    alert('Failed to fetch models from API. Please check your API Key.')
  } finally {
    fetchingModels.value = false
  }
}

const saveNode = async () => {
  loading.value = true
  try {
    const payload = {
      name: formData.value.name,
      service_type: formData.value.service_type,
      driver: formData.value.driver,
      endpoint: formData.value.endpoint,
      is_active: formData.value.is_active,
      is_critical: formData.value.is_critical,
      config_json: formData.value.config
    }

    if (isEdit.value) {
      await $api(`${config.public.apiBase}/master/integration/nodes/${formData.value.id}`, {
        method: 'PUT',
        body: payload
      })
    } else {
      await $api(`${config.public.apiBase}/master/integration/nodes`, {
        method: 'POST',
        body: payload
      })
    }
    
    showModal.value = false
    fetchNodes()
  } catch (err) {
    alert('Failed to save integration node')
  } finally {
    loading.value = false
  }
}

const viewLogs = async () => {
  try {
    const res = await $api(`${config.public.apiBase}/master/integration/sync-logs`)
    if (res && res.data) {
      // Filter logs for the current provider/type if needed, 
      // but the API currently returns all. We'll show all for now.
      syncLogs.value = res.data
    }
    showLogsModal.value = true
  } catch (err) {
    console.error('Failed to fetch sync logs:', err)
    alert('Failed to fetch logs')
  }
}

const decodeError = (details) => {
  if (!details) return ''
  try {
    // If it's base64 encoded string
    if (typeof details === 'string') {
      const decoded = JSON.parse(atob(details))
      return decoded.error || JSON.stringify(decoded)
    }
    // If it's already an object
    return details.error || JSON.stringify(details)
  } catch (e) {
    return String(details)
  }
}

const deleteNode = async (id) => {
  if (!confirm('Are you sure you want to delete this integration node?')) return
  try {
    await $api(`${config.public.apiBase}/master/integration/nodes/${id}`, {
      method: 'DELETE'
    })
    fetchNodes()
  } catch (err) {
    alert('Failed to delete integration node')
  }
}

const getIcon = (type) => {
  switch (type) {
    case 'LDAP': return LucideNetwork
    case 'DATABASE': return LucideDatabase
    case 'STORAGE': return LucideHardDrive
    case 'OCR': return LucideScan
    case 'PRINTER': return LucidePrinter
    case 'WHATSAPP': return LucideMail
    default: return LucideHash
  }
}

const services = computed(() => {
  return nodes.value.map(node => ({
    id: node.id,
    name: node.name,
    category: node.service_type,
    host: node.endpoint,
    status: node.status || 'offline',
    icon: getIcon(node.service_type),
    latencyHistory: node.latency_history || [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    lastLatency: node.last_latency || 0,
    lastError: node.last_error,
    raw: node
  }))
})

const topStats = computed(() => {
  const onlineCount = nodes.value.filter(n => n.status === 'online').length
  const totalCount = nodes.value.length
  const avgLatency = nodes.value.reduce((acc, curr) => acc + (curr.last_latency || 0), 0) / (totalCount || 1)
  const criticalIssues = nodes.value.filter(n => n.status === 'offline' && n.is_critical).length

  return [
    { label: 'Overall Health', value: totalCount ? ((onlineCount/totalCount)*100).toFixed(1) : '0', suffix: '%', icon: LucideCheckCircle2, iconColor: 'text-green-500', color: 'bg-green-500', statusColor: 'bg-green-500', statusTextColor: 'text-green-500', statusText: `${onlineCount}/${totalCount} Services Online` },
    { label: 'Avg Latency', value: avgLatency.toFixed(0), suffix: 'ms', icon: LucideTimer, iconColor: 'text-blue-500', color: 'bg-blue-500', statusColor: 'bg-blue-500', statusTextColor: 'text-blue-500', statusText: 'Global Average' },
    { label: 'Critical Alerts', value: String(criticalIssues).padStart(2, '0'), suffix: '', icon: LucideAlertTriangle, iconColor: 'text-red-500', color: 'bg-red-500', statusColor: 'bg-red-500', statusTextColor: 'text-red-500', statusText: criticalIssues > 0 ? 'Action Required' : 'All Critical Services OK' },
    { label: 'Sync Operations', value: totalSyncCount.value > 999 ? (totalSyncCount.value/1000).toFixed(1) + 'k' : totalSyncCount.value.toString(), suffix: '', icon: LucideRefreshCw, iconColor: 'text-indigo-500', color: 'bg-indigo-500', statusColor: 'bg-indigo-500', statusTextColor: 'text-indigo-500', statusText: 'Lifetime Activity' },
  ]
})

const alerts = computed(() => {
  const result = []
  nodes.value.forEach(node => {
    if (node.status === 'offline') {
      result.push({
        id: node.id,
        type: 'error',
        title: 'Service Offline',
        message: `${node.name} is currently unreachable at ${node.endpoint}.`,
        time: node.last_check_at ? new Date(node.last_check_at).toLocaleTimeString() : 'N/A',
        code: 'ERR_CONN_TIMEOUT',
        bg: 'bg-red-500/5 dark:bg-red-500/10',
        accent: 'bg-red-500',
        titleColor: 'text-red-600 dark:text-red-400'
      })
    } else if (node.last_latency > 100) {
      result.push({
        id: node.id,
        type: 'warning',
        title: 'High Latency',
        message: `${node.name} response time is elevated (${node.last_latency}ms).`,
        time: node.last_check_at ? new Date(node.last_check_at).toLocaleTimeString() : 'N/A',
        code: 'WRN_LATENCY_SPIKE',
        bg: 'bg-amber-500/5 dark:bg-amber-500/10',
        accent: 'bg-amber-500',
        titleColor: 'text-amber-600 dark:text-amber-400'
      })
    }
  })
  if (result.length === 0 && nodes.value.length > 0) {
    result.push({
      id: 'all-ok',
      type: 'info',
      title: 'System Healthy',
      message: 'All integration nodes are responding within normal parameters.',
      time: new Date().toLocaleTimeString(),
      code: 'INF_HEALTH_OK',
      bg: 'bg-blue-500/5 dark:bg-blue-500/10',
      accent: 'bg-blue-500',
      titleColor: 'text-blue-600 dark:text-blue-400'
    })
  }
  return result
})

const globalTrend = computed(() => {
  if (!nodes.value.length) return Array(15).fill(0)
  const maxLength = Math.max(...nodes.value.map(n => (n.latency_history || []).length), 0)
  if (maxLength === 0) return Array(15).fill(0)
  const trend = []
  for (let i = 0; i < maxLength; i++) {
    let sum = 0
    let count = 0
    nodes.value.forEach(node => {
      const history = node.latency_history || []
      const offset = history.length - maxLength + i
      if (offset >= 0 && offset < history.length) {
        sum += history[offset]
        count++
      }
    })
    trend.push(count > 0 ? (sum / count) / 5 : 0)
  }
  return trend
})

const downloadReport = async () => {
  try {
    const res = await $api(`${config.public.apiBase}/master/integration/report`, {
      method: 'GET'
    })
    // Note: fetchWithAuth returns response._data, which for blob might need adjustment 
    // but typically our backend returns data field.
    const url = window.URL.createObjectURL(new Blob([res.data]))
    const link = document.createElement('a')
    link.href = url
    link.setAttribute('download', `integration_report_${new Date().getTime()}.csv`)
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
  } catch (err) {
    console.error('Failed to download report:', err)
  }
}

const refreshAll = () => fetchNodes()

onMounted(() => {
  fetchNodes()
  const timer = setInterval(fetchNodes, 30000)
  onUnmounted(() => clearInterval(timer))
})

definePageMeta({
  layout: 'default'
})
</script>

<style scoped>
/* Glassmorphism subtle polish */
.bg-white\/80 {
  backdrop-filter: blur(24px) saturate(180%);
}

button {
  cursor: pointer;
}
</style>
