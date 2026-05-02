DASHBOARD_HTML = """
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Kreatif DMS | OCR Intelligence</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <script src="https://unpkg.com/lucide@latest"></script>
    <link href="https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@200;300;400;500;600;700;800&display=swap" rel="stylesheet">
    <style>
        body { font-family: 'Plus Jakarta Sans', sans-serif; background-color: #020617; color: #f8fafc; }
        .glass { background: rgba(15, 23, 42, 0.6); backdrop-filter: blur(12px); border: 1px solid rgba(51, 65, 85, 0.5); }
        .custom-scrollbar::-webkit-scrollbar { width: 4px; height: 4px; }
        .custom-scrollbar::-webkit-scrollbar-track { background: rgba(15, 23, 42, 0.1); }
        .custom-scrollbar::-webkit-scrollbar-thumb { background: rgba(56, 189, 248, 0.2); border-radius: 10px; }
        .custom-scrollbar::-webkit-scrollbar-thumb:hover { background: rgba(56, 189, 248, 0.5); }
        @keyframes scan { 0% { transform: translateY(-100%); opacity: 0; } 50% { opacity: 1; } 100% { transform: translateY(100%); opacity: 0; } }
        .scan-line { height: 2px; background: linear-gradient(90deg, transparent, #0ea5e9, transparent); animation: scan 2s linear infinite; }
    </style>
</head>
<body class="min-h-screen p-4 lg:p-8 overflow-x-hidden overflow-y-auto flex flex-col gap-6 custom-scrollbar">
    <!-- Header -->
    <header class="flex items-center justify-between glass p-6 rounded-[2.5rem] shrink-0">
        <div class="flex items-center gap-4">
            <div class="w-12 h-12 bg-sky-500 rounded-2xl flex items-center justify-center shadow-lg shadow-sky-500/20">
                <i data-lucide="scan-text" class="text-white w-7 h-7"></i>
            </div>
            <div>
                <h1 class="text-xl font-black tracking-tighter text-white">KREATIF <span class="text-sky-500">DMS</span></h1>
                <p class="text-[10px] font-bold text-slate-500 uppercase tracking-[0.2em]">Intelligent OCR Engine v1.0</p>
            </div>
        </div>
        <div class="flex items-center gap-3">
            <div id="aiBadge" class="hidden flex items-center gap-2 px-4 py-2 bg-indigo-500/10 border border-indigo-500/20 rounded-xl">
                <div class="w-2 h-2 bg-indigo-500 rounded-full animate-pulse"></div>
                <span class="text-[10px] font-black text-indigo-400 uppercase tracking-widest">Gemini AI Active</span>
            </div>
            <div class="flex items-center gap-2 px-4 py-2 bg-sky-500/10 border border-sky-500/20 rounded-xl">
                <div class="w-2 h-2 bg-sky-500 rounded-full animate-pulse"></div>
                <span class="text-[10px] font-black text-sky-400 uppercase tracking-widest">Live System</span>
            </div>
        </div>
    </header>

    <main class="flex-1 grid grid-cols-1 lg:grid-cols-12 gap-6">
        <!-- Left: Engine Console -->
        <section class="lg:col-span-4 flex flex-col gap-6 min-h-0">
            <div class="glass p-6 rounded-[2.5rem] flex flex-col gap-4">
                <div class="flex items-center justify-between">
                    <h2 class="text-xs font-black uppercase tracking-widest text-slate-400">Engine Playground</h2>
                    <i data-lucide="terminal" class="w-4 h-4 text-sky-500"></i>
                </div>
                
                <div id="dropzone" class="relative group cursor-pointer">
                    <input type="file" id="fileInput" class="hidden" accept="image/*,.pdf" onchange="processFile()">
                    <div onclick="document.getElementById('fileInput').click()" 
                         class="border-2 border-dashed border-slate-800 rounded-[2rem] p-8 text-center transition-all group-hover:border-sky-500/50 group-hover:bg-sky-500/5">
                        <i data-lucide="upload-cloud" class="w-10 h-10 text-slate-600 mx-auto mb-4 group-hover:text-sky-500 group-hover:scale-110 transition-transform"></i>
                        <p class="text-xs font-bold text-slate-400 group-hover:text-slate-200">DROP DOCUMENT OR CLICK TO ANALYZE</p>
                        <p class="text-[9px] text-slate-600 mt-2">Supports JPG, PNG, PDF</p>
                    </div>
                </div>

                <div id="loading" class="hidden space-y-4">
                    <div class="flex items-center justify-between text-[10px] font-black uppercase">
                        <span class="text-sky-400">Processing Document...</span>
                        <span id="progressText" class="text-slate-500">65%</span>
                    </div>
                    <div class="h-1.5 bg-slate-900 rounded-full overflow-hidden">
                        <div class="h-full bg-sky-500 w-[65%] shadow-[0_0_10px_rgba(14,165,233,0.5)]"></div>
                    </div>
                </div>
            </div>

            <!-- Real-time Console -->
            <div class="glass p-6 rounded-[2.5rem] flex-1 flex flex-col gap-4 min-h-0 overflow-hidden">
                <div class="flex items-center justify-between">
                    <h2 class="text-xs font-black uppercase tracking-widest text-slate-400">Live Console</h2>
                    <button onclick="clearConsole()" class="text-[10px] font-black text-slate-600 hover:text-sky-400 transition-colors uppercase">Clear</button>
                </div>
                <div id="console" class="flex-1 font-mono text-[11px] space-y-2 overflow-y-auto custom-scrollbar pr-2">
                    <div class="text-sky-500/50 italic">// Initializing Kreatif DMS OCR System...</div>
                    <div class="text-slate-500">System ready. Awaiting document input.</div>
                </div>
            </div>
        </section>

        <!-- Right: Stats & History -->
        <section class="lg:col-span-8 flex flex-col gap-6 min-h-0">
            <!-- Stats -->
            <div class="grid grid-cols-3 gap-6 shrink-0">
                <div class="glass p-6 rounded-[2rem]">
                    <p class="text-[10px] font-black text-slate-500 uppercase tracking-widest mb-1">Total Processed</p>
                    <p id="statTotal" class="text-3xl font-black text-white">0</p>
                </div>
                <div class="glass p-6 rounded-[2rem]">
                    <p class="text-[10px] font-black text-slate-500 uppercase tracking-widest mb-1">Avg Duration</p>
                    <p id="statDur" class="text-3xl font-black text-sky-400">0s</p>
                </div>
                <div class="glass p-6 rounded-[2rem]">
                    <p class="text-[10px] font-black text-slate-500 uppercase tracking-widest mb-1">Success Rate</p>
                    <p id="statRate" class="text-3xl font-black text-green-500">0%</p>
                </div>
            </div>

            <!-- Table -->
            <div class="glass p-8 rounded-[2.5rem] flex-1 min-h-0 flex flex-col overflow-hidden">
                <div class="flex items-center justify-between mb-6">
                    <h2 class="text-xs font-black uppercase tracking-widest text-slate-400">Processing History</h2>
                    <div class="flex gap-2">
                        <div class="w-3 h-3 bg-green-500/20 border border-green-500/30 rounded-sm"></div>
                        <div class="w-3 h-3 bg-red-500/20 border border-red-500/30 rounded-sm"></div>
                    </div>
                </div>
                <div class="flex-1 overflow-auto custom-scrollbar">
                    <table class="w-full text-left border-collapse">
                        <thead class="sticky top-0 bg-[#0f172a] z-10">
                            <tr class="text-[10px] font-black uppercase tracking-widest text-slate-500 border-b border-slate-800/50">
                                <th class="pb-4 px-6">Filename</th>
                                <th class="pb-4 px-6">Size</th>
                                <th class="pb-4 px-6">Doc Type (AI)</th>
                                <th class="pb-4 px-6">Duration</th>
                                <th class="pb-4 px-6">Accuracy</th>
                                <th class="pb-4 px-6">Status</th>
                                <th class="pb-4 px-6">Action</th>
                            </tr>
                        </thead>
                        <tbody id="historyBody">
                            <!-- Rows loaded here -->
                        </tbody>
                    </table>
                </div>
            </div>
        </section>
    </main>

    <!-- Viewer Modal -->
    <div id="viewerModal" class="fixed inset-0 z-50 hidden">
        <div class="absolute inset-0 bg-slate-950/90 backdrop-blur-xl"></div>
        <div class="absolute inset-6 lg:inset-12 glass rounded-[3rem] overflow-hidden flex flex-col animate-in fade-in zoom-in duration-300">
            <!-- Modal Header -->
            <div class="p-8 border-b border-slate-800/50 flex items-center justify-between">
                <div>
                    <h3 id="modalTitle" class="text-xl font-black text-white">document_preview.pdf</h3>
                    <p id="modalMeta" class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mt-1">2026-05-01 12:34:56 | Processing Success</p>
                </div>
                <button onclick="closeViewer()" class="p-3 bg-slate-900 border border-slate-800 rounded-2xl hover:bg-red-500/10 hover:border-red-500/50 hover:text-red-500 transition-all">
                    <i data-lucide="x" class="w-6 h-6"></i>
                </button>
            </div>

            <!-- Modal Content -->
            <div class="flex-1 p-8 grid grid-cols-1 lg:grid-cols-12 gap-8 overflow-y-auto custom-scrollbar select-none">
                <!-- Left: Document View (Taking 8 columns) -->
                <div class="lg:col-span-8 glass rounded-[2rem] relative overflow-hidden bg-slate-950/50 flex flex-col h-[500px] lg:h-full shrink-0">
                    <div class="p-4 border-b border-slate-800/50 flex items-center justify-between">
                        <div class="flex items-center gap-4">
                            <button onclick="prevPage()" class="p-2 bg-slate-900 border border-slate-800 rounded-lg hover:bg-slate-800 transition-all" title="Previous Page">
                                <i data-lucide="chevron-left" class="w-4 h-4"></i>
                            </button>
                            <span class="text-xs font-black uppercase tracking-widest"><span id="currentPageNum">1</span> / <span id="totalPageNum">1</span></span>
                            <button onclick="nextPage()" class="p-2 bg-slate-900 border border-slate-800 rounded-lg hover:bg-slate-800 transition-all" title="Next Page">
                                <i data-lucide="chevron-right" class="w-4 h-4"></i>
                            </button>
                        </div>
                        
                        <div class="flex items-center gap-2">
                            <button onclick="zoomOut()" class="p-2 bg-slate-900 border border-slate-800 rounded-lg hover:bg-slate-800 transition-all" title="Zoom Out">
                                <i data-lucide="zoom-out" class="w-4 h-4 text-slate-400"></i>
                            </button>
                            <span id="zoomLevel" class="text-[10px] font-black text-sky-400 uppercase w-12 text-center">100%</span>
                            <button onclick="zoomIn()" class="p-2 bg-slate-900 border border-slate-800 rounded-lg hover:bg-slate-800 transition-all" title="Zoom In">
                                <i data-lucide="zoom-in" class="w-4 h-4 text-slate-400"></i>
                            </button>
                            <div class="w-px h-4 bg-slate-800 mx-1"></div>
                            <button onclick="resetZoom()" class="p-2 bg-slate-900 border border-slate-800 rounded-lg hover:bg-slate-800 transition-all" title="Zoom as Page">
                                <i data-lucide="maximize" class="w-4 h-4 text-slate-400"></i>
                            </button>
                        </div>
                    </div>
                    <div class="flex-1 overflow-auto custom-scrollbar p-8 flex items-start justify-center bg-slate-950/20 cursor-grab active:cursor-grabbing" id="viewerScrollContainer">
                        <div id="zoomWrapper" class="inline-block">
                            <div id="imageContainer" class="relative transition-transform duration-200 ease-out origin-top">
                                <img id="originalImage" class="max-w-none shadow-2xl rounded-sm" src="" alt="Document">
                                <div id="overlay" class="absolute inset-0 z-10 pointer-events-none"></div>
                            </div>
                        </div>
                    </div>
                </div>

                <!-- Right: Analysis & Data (Taking 4 columns) -->
                <div class="lg:col-span-4 flex flex-col gap-6 overflow-y-auto custom-scrollbar pr-2 min-h-0">
                    <!-- AI Insights -->
                    <div id="aiPanel" class="glass p-6 rounded-[2rem] bg-indigo-500/5 border-indigo-500/20 flex flex-col gap-4 shrink-0">
                        <div class="flex items-center justify-between">
                            <h4 class="text-[10px] font-black uppercase tracking-widest text-indigo-400">AI Smart Insights</h4>
                            <span id="docTypeTag" class="px-2 py-1 bg-indigo-500/20 text-indigo-300 rounded text-[9px] font-black uppercase tracking-widest">Invoice</span>
                        </div>
                        <p id="aiSummary" class="text-sm font-medium text-slate-300 italic">"Detected invoice from Kreatif Studio for Cloud Services."</p>
                        <div id="aiEntities" class="flex flex-wrap gap-2">
                            <!-- Entities injected here -->
                        </div>
                    </div>

                    <!-- Cleaned Text -->
                    <div class="glass p-6 rounded-[2rem] flex flex-col gap-4 shrink-0 bg-emerald-500/5 border-emerald-500/20">
                        <div class="flex items-center justify-between">
                            <h4 class="text-[10px] font-black uppercase tracking-widest text-emerald-400">Cleaned Text Recovery</h4>
                            <button onclick="copyCleanedText()" class="p-1.5 hover:bg-emerald-500/20 text-emerald-500 rounded-lg transition-all">
                                <i data-lucide="copy" class="w-3.5 h-3.5"></i>
                            </button>
                        </div>
                        <div id="cleanedText" class="text-sm font-medium text-slate-400 leading-relaxed">
                            Awaiting AI refinement...
                        </div>
                    </div>

                    <!-- OCR Raw Data -->
                    <div class="glass p-6 rounded-[2.5rem] flex flex-col gap-4 shrink-0">
                        <div class="flex items-center justify-between">
                            <h4 class="text-[10px] font-black uppercase tracking-widest text-slate-500">Detected OCR Fragments</h4>
                            <button onclick="toggleHighlight()" id="highlightAllBtn" class="text-[9px] font-black uppercase tracking-widest px-3 py-1 border border-slate-700 rounded-lg hover:bg-slate-700 transition-all">Highlight All</button>
                        </div>
                        <div id="textDetails" class="space-y-3">
                            <!-- Words injected here -->
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <script>
        lucide.createIcons();

        let isHighlighted = false;
        let currentWords = [];
        let currentPage = 1;
        let previewPaths = [];
        let zoomScale = 1.0;
        const ZOOM_STEP = 0.1;
        const MIN_ZOOM = 0.1;
        const MAX_ZOOM = 5.0;

        function zoomIn() {
            if (zoomScale < MAX_ZOOM) {
                zoomScale += ZOOM_STEP;
                applyZoom();
            }
        }

        function zoomOut() {
            if (zoomScale > MIN_ZOOM) {
                zoomScale -= ZOOM_STEP;
                applyZoom();
            }
        }

        function resetZoom() {
            const container = document.getElementById('viewerScrollContainer');
            const img = document.getElementById('originalImage');
            if (img.complete && img.naturalWidth > 0) {
                const availableWidth = container.clientWidth - 64; // padding p-8
                zoomScale = availableWidth / img.naturalWidth;
                if (zoomScale > 1.2) zoomScale = 1.0; // Don't over-scale small images too much
            } else {
                zoomScale = 1.0;
            }
            applyZoom();
        }

        function applyZoom() {
            const container = document.getElementById('imageContainer');
            const wrapper = document.getElementById('zoomWrapper');
            const img = document.getElementById('originalImage');
            
            container.style.transform = `scale(${zoomScale})`;
            document.getElementById('zoomLevel').innerText = Math.round(zoomScale * 100) + '%';
            
            if (img.complete && img.naturalWidth > 0) {
                wrapper.style.width = (img.naturalWidth * zoomScale) + 'px';
                wrapper.style.height = (img.naturalHeight * zoomScale) + 'px';
            }
        }

        // Mouse wheel zoom
        const scrollContainer = document.getElementById('viewerScrollContainer');
        scrollContainer.addEventListener('wheel', (e) => {
            if (e.ctrlKey) {
                e.preventDefault();
                if (e.deltaY < 0) zoomIn();
                else zoomOut();
            }
        }, { passive: false });

        // Drag to scroll (Panning)
        let isDragging = false;
        let startX, startY, scrollLeft, scrollTop;

        scrollContainer.addEventListener('mousedown', (e) => {
            if (e.button !== 0) return; // Only left click
            isDragging = true;
            startX = e.pageX - scrollContainer.offsetLeft;
            startY = e.pageY - scrollContainer.offsetTop;
            scrollLeft = scrollContainer.scrollLeft;
            scrollTop = scrollContainer.scrollTop;
        });

        scrollContainer.addEventListener('mouseleave', () => { isDragging = false; });
        scrollContainer.addEventListener('mouseup', () => { isDragging = false; });

        scrollContainer.addEventListener('mousemove', (e) => {
            if (!isDragging) return;
            e.preventDefault();
            const x = e.pageX - scrollContainer.offsetLeft;
            const y = e.pageY - scrollContainer.offsetTop;
            const walkX = (x - startX) * 1.5;
            const walkY = (y - startY) * 1.5;
            scrollContainer.scrollLeft = scrollLeft - walkX;
            scrollContainer.scrollTop = scrollTop - walkY;
        });

        function prevPage() {
            if (currentPage > 1) {
                currentPage--;
                updatePageView();
            }
        }

        function nextPage() {
            if (currentPage < previewPaths.length) {
                currentPage++;
                updatePageView();
            }
        }

        function updatePageView() {
            const img = document.getElementById('originalImage');
            img.src = previewPaths[currentPage - 1];
            document.getElementById('currentPageNum').innerText = currentPage;
            
            // Wait for image to load to re-render overlays and reset zoom if first time
            img.onload = () => {
                renderTextDetails(currentWords);
                // Optional: resetZoom() on first page load or every page change?
                // resetZoom(); 
            };
        }

        // Drag and Drop support
        const dropzone = document.getElementById('dropzone');
        dropzone.ondragover = (e) => {
            e.preventDefault();
            dropzone.classList.add('border-sky-500', 'bg-sky-500/5');
        };
        dropzone.ondragleave = () => {
            dropzone.classList.remove('border-sky-500', 'bg-sky-500/5');
        };
        dropzone.ondrop = (e) => {
            e.preventDefault();
            dropzone.classList.remove('border-sky-500', 'bg-sky-500/5');
            const file = e.dataTransfer.files[0];
            if (file) {
                const input = document.getElementById('fileInput');
                const dataTransfer = new DataTransfer();
                dataTransfer.items.add(file);
                input.files = dataTransfer.files;
                processFile();
            }
        };

        async function processFile() {
            const input = document.getElementById('fileInput');
            const file = input.files[0];
            if (!file) return;

            const loading = document.getElementById('loading');
            const dropzoneUI = document.getElementById('dropzone');
            
            loading.classList.remove('hidden');
            dropzoneUI.classList.add('opacity-50', 'pointer-events-none');

            const formData = new FormData();
            formData.append('file', file);

            try {
                addLog(`[UI] Uploading ${file.name} (${(file.size/1024).toFixed(1)} KB)...`, 'text-sky-400');
                const resp = await fetch('/ocr/process', {
                    method: 'POST',
                    body: formData
                });
                
                if (!resp.ok) {
                    const errorData = await resp.json().catch(() => ({detail: resp.statusText}));
                    throw new Error(errorData.detail || `Server error: ${resp.status}`);
                }
                
                const data = await resp.json();
                addLog(`[UI] Extraction complete!`, 'text-green-400');
                loadDashboardData();
            } catch (err) {
                addLog(`[UI] Error: ${err.message}`, 'text-red-500');
                console.error("Upload error:", err);
                alert("Upload Failed: " + err.message);
            } finally {
                loading.classList.add('hidden');
                dropzoneUI.classList.remove('opacity-50', 'pointer-events-none');
                input.value = '';
            }
        }

        async function loadDashboardData() {
            try {
                const resp = await fetch('/ocr/stats');
                const data = await resp.json();
                
                document.getElementById('statTotal').innerText = data.stats.total;
                document.getElementById('statDur').innerText = data.stats.avg_time.toFixed(2) + 's';
                document.getElementById('statRate').innerText = data.stats.success_rate.toFixed(1) + '%';
                
                const historyBody = document.getElementById('historyBody');
                historyBody.innerHTML = '';
                
                data.history.forEach(req => {
                    const statusColor = req.status === "Success" ? "text-green-500 bg-green-500/10" : "text-red-500 bg-red-500/10";
                    let aiData = null;
                    try {
                        aiData = (req.ai_analysis && typeof req.ai_analysis === 'string') 
                                 ? JSON.parse(req.ai_analysis) 
                                 : req.ai_analysis;
                    } catch(e) {}
                    const docType = aiData ? (aiData.doc_type || '-') : '-';

                    const tr = document.createElement('tr');
                    tr.className = "border-b border-slate-800/50 hover:bg-slate-800/30 transition-colors";
                    tr.innerHTML = `
                        <td class="py-4 px-6 font-bold text-slate-200 text-sm">${req.filename}</td>
                        <td class="py-4 px-6 text-slate-400 text-[10px]">${req.size}</td>
                        <td class="py-4 px-6 font-black text-[10px] text-sky-400 uppercase tracking-tighter">${docType}</td>
                        <td class="py-4 px-6 font-black text-xs text-sky-400">${req.duration.toFixed(2)}s</td>
                        <td class="py-4 px-6 font-black text-xs text-sky-500">${(req.accuracy * 100).toFixed(1)}%</td>
                        <td class="py-4 px-6">
                            <span class="px-3 py-1 rounded-full text-[9px] font-black uppercase tracking-widest ${statusColor}">
                                ${req.status}
                            </span>
                        </td>
                        <td class="py-4 px-6">
                            <button onclick="viewResult(${req.id})" class="p-1.5 bg-sky-500/10 hover:bg-sky-500/20 text-sky-500 rounded-lg transition-colors">
                                <i data-lucide="eye" class="w-3.5 h-3.5"></i>
                            </button>
                        </td>
                    `;
                    historyBody.appendChild(tr);
                });
                lucide.createIcons();
            } catch (err) {
                console.error("Failed to load history:", err);
            }
        }

        async function viewResult(jobId) {
            try {
                const response = await fetch(`/ocr/result/${jobId}`);
                const data = await response.json();
                
                if (data) {
                    document.getElementById('modalTitle').innerText = data.filename;
                    document.getElementById('modalMeta').innerText = `${data.timestamp} | Status: ${data.status}`;
                    
                    previewPaths = data.preview_paths || [data.preview_path];
                    currentPage = 1;
                    document.getElementById('totalPageNum').innerText = previewPaths.length;
                    
                    const img = document.getElementById('originalImage');
                    img.src = previewPaths[0];
                    document.getElementById('currentPageNum').innerText = "1";
                    
                    console.log("[JS] Result data received:", data);
                    currentWords = typeof data.words_json === 'string' ? JSON.parse(data.words_json) : data.words_json;
                    console.log("[JS] currentWords count:", currentWords.length);
                    
                    // AI Panel - Move outside onload to ensure it shows up immediately
                    const aiPanel = document.getElementById('aiPanel');
                    const aiBadge = document.getElementById('aiBadge');
                    const cleanedTextDiv = document.getElementById('cleanedText');
                    
                    let ai = null;
                    try {
                        ai = typeof data.ai_analysis === 'string' ? JSON.parse(data.ai_analysis) : data.ai_analysis;
                    } catch(e) {}

                    if (ai) {
                        aiPanel.classList.remove('hidden');
                        aiBadge.classList.remove('hidden');
                        document.getElementById('docTypeTag').innerText = ai.doc_type || 'Document';
                        document.getElementById('aiSummary').innerText = `"${ai.summary || ''}"`;
                        cleanedTextDiv.innerText = ai.cleaned_text || 'No cleaned text available.';
                        
                        const entitiesDiv = document.getElementById('aiEntities');
                        entitiesDiv.innerHTML = '';
                        if (ai.entities) {
                            Object.entries(ai.entities).forEach(([key, val]) => {
                                const span = document.createElement('div');
                                span.className = 'px-3 py-1 bg-slate-800/50 rounded-lg border border-slate-700/50 flex items-center gap-2 whitespace-nowrap';
                                span.innerHTML = `<span class="text-[9px] font-black text-slate-500 uppercase">${key}:</span><span class="text-[10px] font-bold text-sky-400">${val}</span>`;
                                entitiesDiv.appendChild(span);
                            });
                        }
                    } else {
                        aiPanel.classList.add('hidden');
                        aiBadge.classList.add('hidden');
                        cleanedTextDiv.innerText = 'No AI analysis available.';
                    }

                    renderTextDetails(currentWords); // Render text panel immediately
                    
                    img.onload = () => {
                        renderTextDetails(currentWords);
                        resetZoom(); // Fit to page on first view
                    };

                    document.getElementById('viewerModal').classList.remove('hidden');
                    lucide.createIcons();
                }
            } catch (err) {
                console.error("Failed to fetch result:", err);
            }
        }

        function renderTextDetails(words) {
            console.log("[JS] Rendering details for page:", currentPage);
            const container = document.getElementById('textDetails');
            const overlay = document.getElementById('overlay');
            container.innerHTML = '';
            overlay.innerHTML = '';

            const pageWords = words.filter(w => w.page == currentPage);
            console.log("[JS] pageWords count:", pageWords.length);

            pageWords.forEach((w, i) => {
                const div = document.createElement('div');
                div.id = `text-item-${i}`;
                div.className = 'group p-4 rounded-xl border border-slate-800/40 hover:border-sky-500/50 hover:bg-sky-500/5 transition-all cursor-pointer';
                div.innerHTML = `
                    <div class="flex items-center justify-between mb-1">
                        <span class="text-[9px] font-black text-slate-600 uppercase">Conf: ${ (w.confidence * 100).toFixed(1) }%</span>
                    </div>
                    <p class="text-sm font-medium text-slate-300 group-hover:text-white">${w.text}</p>
                `;
                
                div.onclick = () => {
                    isHighlighted = false;
                    document.getElementById('highlightAllBtn').innerText = 'Highlight All';
                    highlightBox(w.box, i, true);
                    div.scrollIntoView({ behavior: 'smooth', block: 'center' });
                };
                
                div.onmouseenter = () => highlightBox(w.box, i, true);
                div.onmouseleave = () => { if(!isHighlighted) overlay.innerHTML = ''; };
                container.appendChild(div);
                
                if (isHighlighted) highlightBox(w.box, i);
            });
        }

        function highlightBox(box, index, single = false) {
            const overlay = document.getElementById('overlay');
            const rect = document.createElement('div');
            // Ensure pointer events are active on the box even if parent has them disabled
            rect.className = 'absolute border-2 border-sky-500/40 bg-sky-500/10 transition-all cursor-pointer hover:bg-sky-500/20 hover:border-sky-400 pointer-events-auto';
            if (single) {
                overlay.innerHTML = '';
                rect.classList.add('border-sky-400', 'bg-sky-500/30', 'z-20', 'ring-4', 'ring-sky-500/20');
            }
            rect.style.left = `${box.x}px`;
            rect.style.top = `${box.y}px`;
            rect.style.width = `${box.w}px`;
            rect.style.height = `${box.h}px`;
            
            rect.onclick = (e) => {
                e.stopPropagation();
                const item = document.getElementById(`text-item-${index}`);
                if (item) {
                    item.scrollIntoView({ behavior: 'smooth', block: 'center' });
                    item.classList.add('bg-sky-500/20', 'border-sky-500');
                    setTimeout(() => item.classList.remove('bg-sky-500/20', 'border-sky-500'), 2000);
                    highlightBox(box, index, true);
                }
            };

            overlay.appendChild(rect);
        }

        function toggleHighlight() {
            isHighlighted = !isHighlighted;
            document.getElementById('highlightAllBtn').innerText = isHighlighted ? 'Hide All' : 'Highlight All';
            renderTextDetails(currentWords);
        }

        function closeViewer() {
            document.getElementById('viewerModal').classList.add('hidden');
        }

        function copyCleanedText() {
            const text = document.getElementById('cleanedText').innerText;
            navigator.clipboard.writeText(text);
            alert('Cleaned text copied!');
        }

        // WebSocket
        const consoleDiv = document.getElementById('console');
        const ws_protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
        const ws = new WebSocket(`${ws_protocol}//${window.location.host}/ws/logs`);
        ws.onmessage = (event) => {
            if (event.data === "REFRESH_HISTORY") {
                loadDashboardData();
                return;
            }
            const div = document.createElement('div');
            div.className = 'text-slate-400';
            div.innerHTML = `<span class="text-slate-600 mr-2">>>></span>${event.data}`;
            consoleDiv.appendChild(div);
            consoleDiv.scrollTop = consoleDiv.scrollHeight;
        };

        function addLog(msg, color) {
            const div = document.createElement('div');
            div.className = color;
            div.innerHTML = `<span class="text-slate-600 mr-2">>>></span>${msg}`;
            consoleDiv.appendChild(div);
            consoleDiv.scrollTop = consoleDiv.scrollHeight;
        }

        function clearConsole() { consoleDiv.innerHTML = ''; }

        loadDashboardData();
    </script>
</body>
</html>
"""
