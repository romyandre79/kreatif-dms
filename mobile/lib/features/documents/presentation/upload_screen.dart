import 'package:flutter/material.dart';
import 'package:lucide_icons/lucide_icons.dart';
import 'package:kreatif_dms/core/theme/app_theme.dart';

class UploadScreen extends StatefulWidget {
  const UploadScreen({super.key});

  @override
  State<UploadScreen> createState() => _UploadScreenState();
}

class _UploadScreenState extends State<UploadScreen> {
  final _titleController = TextEditingController();
  String _selectedCategory = 'Invoice';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Upload Dokumen', style: TextStyle(fontWeight: FontWeight.bold)),
        backgroundColor: Colors.transparent,
        elevation: 0,
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(24.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Upload Area
            GestureDetector(
              onTap: () {},
              child: Container(
                width: double.infinity,
                height: 180,
                decoration: BoxDecoration(
                  color: AppTheme.cardBg,
                  borderRadius: BorderRadius.circular(24),
                  border: Border.all(color: AppTheme.primary.withOpacity(0.3), width: 2),
                ),
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    const Icon(LucideIcons.uploadCloud, size: 48, color: AppTheme.primary),
                    const SizedBox(height: 16),
                    Text(
                      'Pilih file atau ambil foto',
                      style: Theme.of(context).textTheme.bodyMedium,
                    ),
                    const SizedBox(height: 4),
                    const Text(
                      'PDF, JPG, PNG (Max 10MB)',
                      style: TextStyle(color: AppTheme.textSecondary, fontSize: 10),
                    ),
                  ],
                ),
              ),
            ),
            
            const SizedBox(height: 32),
            
            // Metadata Form
            const Text('Metadata Dokumen', style: TextStyle(fontWeight: FontWeight.bold)),
            const SizedBox(height: 16),
            
            TextField(
              controller: _titleController,
              decoration: const InputDecoration(
                labelText: 'Judul Dokumen',
                hintText: 'Contoh: Invoice PT ABC Jan 2024',
              ),
            ),
            const SizedBox(height: 16),
            
            DropdownButtonFormField<String>(
              value: _selectedCategory,
              decoration: const InputDecoration(labelText: 'Kategori'),
              items: ['Invoice', 'Contract', 'HR Document', 'Legal'].map((cat) {
                return DropdownMenuItem(value: cat, child: Text(cat));
              }).toList(),
              onChanged: (val) {
                setState(() {
                  _selectedCategory = val!;
                });
              },
            ),
            const SizedBox(height: 16),
            
            const TextField(
              maxLines: 3,
              decoration: InputDecoration(
                labelText: 'Deskripsi (Opsional)',
                hintText: 'Tambahkan catatan tentang dokumen ini...',
              ),
            ),
            
            const SizedBox(height: 40),
            
            ElevatedButton(
              onPressed: () {},
              child: const Text('Simpan & Upload'),
            ),
          ],
        ),
      ),
    );
  }
}
