import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:lucide_icons/lucide_icons.dart';
import 'package:file_picker/file_picker.dart';
import 'package:image_picker/image_picker.dart';
import 'package:kreatif_dms/core/theme/app_theme.dart';
import 'package:kreatif_dms/features/documents/presentation/document_provider.dart';
import 'package:kreatif_dms/features/documents/data/document_repository.dart';
import 'package:kreatif_dms/features/documents/domain/document_models.dart';
import 'package:kreatif_dms/features/auth/presentation/auth_provider.dart';
import 'dart:io';

class UploadScreen extends ConsumerStatefulWidget {
  const UploadScreen({super.key});

  @override
  ConsumerState<UploadScreen> createState() => _UploadScreenState();
}

class _UploadScreenState extends ConsumerState<UploadScreen> {
  final _titleController = TextEditingController();
  final _descController = TextEditingController();
  
  File? _selectedFile;
  String? _selectedFileName;
  bool _isUploading = false;

  // Topology selection
  TopologyNode? _selectedCompany;
  TopologyNode? _selectedBranch;
  TopologyNode? _selectedDepartment;
  TopologyNode? _selectedRack;
  TopologyNode? _selectedBox;
  TopologyNode? _selectedOrdner;

  Future<void> _pickFile() async {
    final FilePickerResult? result = await FilePicker.pickFiles(
      type: FileType.any,
    );

    if (result != null) {
      setState(() {
        _selectedFile = File(result.files.single.path!);
        _selectedFileName = result.files.single.name;
        if (_titleController.text.isEmpty) {
          _titleController.text = _selectedFileName!.split('.').first;
        }
      });
    }
  }

  Future<void> _takePhoto() async {
    final picker = ImagePicker();
    final photo = await picker.pickImage(source: ImageSource.camera);

    if (photo != null) {
      setState(() {
        _selectedFile = File(photo.path);
        _selectedFileName = photo.name;
        if (_titleController.text.isEmpty) {
          _titleController.text = 'Foto_${DateTime.now().millisecondsSinceEpoch}';
        }
      });
    }
  }

  Future<void> _handleUpload() async {
    if (_selectedFile == null) {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Pilih file terlebih dahulu')),
      );
      return;
    }

    if (_selectedCompany == null || _selectedBranch == null || _selectedDepartment == null) {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Pilih Company, Branch, dan Department')),
      );
      return;
    }

    setState(() => _isUploading = true);

    try {
      await ref.read(documentRepositoryProvider).uploadDocument(
        filePath: _selectedFile!.path,
        title: _titleController.text,
        description: _descController.text,
        companyId: _selectedCompany!.id,
        branchId: _selectedBranch!.id,
        departmentId: _selectedDepartment!.id,
        rackId: _selectedRack?.id,
        boxId: _selectedBox?.id,
        ordnerId: _selectedOrdner?.id,
      );

      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Dokumen berhasil diunggah')),
        );
        Navigator.pop(context);
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Gagal unggah: $e'), backgroundColor: Colors.red),
        );
      }
    } finally {
      if (mounted) setState(() => _isUploading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    final topologyAsync = ref.watch(topologyProvider);

    return Scaffold(
      appBar: AppBar(
        title: const Text('Upload Dokumen', style: TextStyle(fontWeight: FontWeight.bold)),
        backgroundColor: Colors.transparent,
        elevation: 0,
      ),
      body: topologyAsync.when(
        data: (companies) {
          // Flatten departments and track parents
          final List<TopologyNode> allDepts = [];
          final Map<String, Map<String, String>> deptMap = {};

          for (final c in companies) {
            for (final b in c.children ?? <TopologyNode>[]) {
              for (final d in b.children ?? <TopologyNode>[]) {
                allDepts.add(d);
                deptMap[d.id] = {'companyId': c.id, 'branchId': b.id};
              }
            }
          }

          // Pre-select user's department if available
          final authState = ref.watch(authProvider);
          if (_selectedDepartment == null && authState.user?.departmentId != null) {
            final userDeptId = authState.user!.departmentId;
            final userDept = allDepts.where((d) => d.id == userDeptId).firstOrNull;
            if (userDept != null) {
              WidgetsBinding.instance.addPostFrameCallback((_) {
                if (mounted && _selectedDepartment == null) {
                  setState(() {
                    _selectedDepartment = userDept;
                    final ctx = deptMap[userDept.id];
                    if (ctx != null) {
                      _selectedCompany = TopologyNode(id: ctx['companyId']!, name: '', type: 'company');
                      _selectedBranch = TopologyNode(id: ctx['branchId']!, name: '', type: 'branch');
                    }
                  });
                }
              });
            }
          }

          return SingleChildScrollView(
            padding: const EdgeInsets.all(24.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
              // Upload Area
              GestureDetector(
                onTap: () {
                  showModalBottomSheet(
                    context: context,
                    builder: (context) => SafeArea(
                      child: Column(
                        mainAxisSize: MainAxisSize.min,
                        children: [
                          ListTile(
                            leading: const Icon(LucideIcons.file),
                            title: const Text('Pilih File'),
                            onTap: () {
                              Navigator.pop(context);
                              _pickFile();
                            },
                          ),
                          ListTile(
                            leading: const Icon(LucideIcons.camera),
                            title: const Text('Ambil Foto'),
                            onTap: () {
                              Navigator.pop(context);
                              _takePhoto();
                            },
                          ),
                        ],
                      ),
                    ),
                  );
                },
                child: Container(
                  width: double.infinity,
                  height: 140,
                  decoration: BoxDecoration(
                    color: AppTheme.cardBg,
                    borderRadius: BorderRadius.circular(24),
                    border: Border.all(
                      color: _selectedFile != null ? Colors.green : AppTheme.primary.withOpacity(0.3), 
                      width: 2,
                    ),
                  ),
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Icon(
                        _selectedFile != null ? LucideIcons.checkCircle : LucideIcons.uploadCloud, 
                        size: 40, 
                        color: _selectedFile != null ? Colors.green : AppTheme.primary,
                      ),
                      const SizedBox(height: 12),
                      Text(
                        _selectedFileName ?? 'Pilih file atau ambil foto',
                        style: Theme.of(context).textTheme.bodyMedium,
                        textAlign: TextAlign.center,
                      ),
                    ],
                  ),
                ),
              ),
              
              const SizedBox(height: 24),
              
              const Text('Informasi Dokumen', style: TextStyle(fontWeight: FontWeight.bold)),
              const SizedBox(height: 16),
              
              TextField(
                controller: _titleController,
                decoration: const InputDecoration(
                  labelText: 'Judul Dokumen',
                  prefixIcon: Icon(LucideIcons.type, size: 20),
                ),
              ),
              const SizedBox(height: 16),
              
              TextField(
                controller: _descController,
                maxLines: 2,
                decoration: const InputDecoration(
                  labelText: 'Deskripsi (Opsional)',
                  prefixIcon: Icon(LucideIcons.alignLeft, size: 20),
                ),
              ),
              
              const SizedBox(height: 24),
              const Text('Lokasi & Penyimpanan', style: TextStyle(fontWeight: FontWeight.bold)),
              const SizedBox(height: 16),

              // Department Dropdown
              DropdownButtonFormField<TopologyNode>(
                value: _selectedDepartment,
                decoration: const InputDecoration(labelText: 'Department'),
                items: allDepts.map((d) => DropdownMenuItem(
                  value: d, 
                  child: Text(d.name),
                )).toList(),
                onChanged: (val) {
                  setState(() {
                    _selectedDepartment = val;
                    // Find parent IDs from the flattened map
                    final ctx = deptMap[val?.id];
                    if (ctx != null) {
                      _selectedCompany = TopologyNode(id: ctx['companyId']!, name: '', type: 'company');
                      _selectedBranch = TopologyNode(id: ctx['branchId']!, name: '', type: 'branch');
                    }
                    _selectedRack = null;
                    _selectedBox = null;
                    _selectedOrdner = null;
                  });
                },
              ),
              const SizedBox(height: 16),

              // Rack Dropdown
              DropdownButtonFormField<TopologyNode>(
                value: _selectedRack,
                decoration: const InputDecoration(labelText: 'Rack (Opsional)'),
                items: _selectedDepartment?.children?.map((node) => DropdownMenuItem(
                  value: node, 
                  child: Text(node.name),
                )).toList() ?? [],
                onChanged: (val) {
                  setState(() {
                    _selectedRack = val;
                    _selectedBox = null;
                    _selectedOrdner = null;
                  });
                },
              ),
              const SizedBox(height: 16),

              // Box Dropdown
              DropdownButtonFormField<TopologyNode>(
                value: _selectedBox,
                decoration: const InputDecoration(labelText: 'Box (Opsional)'),
                items: _selectedRack?.children?.map((node) => DropdownMenuItem(
                  value: node, 
                  child: Text(node.name),
                )).toList() ?? [],
                onChanged: (val) {
                  setState(() {
                    _selectedBox = val;
                    _selectedOrdner = null;
                  });
                },
              ),
              const SizedBox(height: 16),

              // Ordner Dropdown
              DropdownButtonFormField<TopologyNode>(
                value: _selectedOrdner,
                decoration: const InputDecoration(labelText: 'Ordner (Opsional)'),
                items: _selectedBox?.children?.map((node) => DropdownMenuItem(
                  value: node, 
                  child: Text(node.name),
                )).toList() ?? [],
                onChanged: (val) {
                  setState(() {
                    _selectedOrdner = val;
                  });
                },
              ),
              
              const SizedBox(height: 40),
              
              SizedBox(
                width: double.infinity,
                child: ElevatedButton(
                  onPressed: _isUploading ? null : _handleUpload,
                  style: ElevatedButton.styleFrom(
                    padding: const EdgeInsets.symmetric(vertical: 16),
                    shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
                  ),
                  child: _isUploading 
                    ? const SizedBox(height: 20, width: 20, child: CircularProgressIndicator(strokeWidth: 2, color: Colors.white))
                    : const Text('Simpan & Upload', style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold)),
                ),
              ),
              const SizedBox(height: 24),
            ],
          ),
        );
      },
      loading: () => const Center(child: CircularProgressIndicator()),
      error: (err, stack) => Center(child: Text('Error: $err')),
    ));
  }
}
