import 'package:flutter/material.dart';
import 'package:lucide_icons/lucide_icons.dart';
import 'package:kreatif_dms/core/theme/app_theme.dart';

class DocumentExplorerScreen extends StatelessWidget {
  const DocumentExplorerScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Explorer', style: TextStyle(fontWeight: FontWeight.bold)),
        backgroundColor: Colors.transparent,
        elevation: 0,
        actions: [
          IconButton(onPressed: () {}, icon: const Icon(LucideIcons.filter, size: 20)),
        ],
      ),
      body: Column(
        children: [
          // Breadcrumbs
          Container(
            padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
            child: Row(
              children: [
                const Icon(LucideIcons.home, size: 14, color: AppTheme.primary),
                const Icon(LucideIcons.chevronRight, size: 14, color: AppTheme.textSecondary),
                Text('Root', style: TextStyle(color: AppTheme.primary, fontSize: 12, fontWeight: FontWeight.bold)),
              ],
            ),
          ),
          
          Expanded(
            child: ListView(
              padding: const EdgeInsets.all(16),
              children: [
                _FolderItem(title: 'Finance & Accounting', subtitle: '428 Documents', icon: LucideIcons.folder),
                _FolderItem(title: 'Human Resources', subtitle: '156 Documents', icon: LucideIcons.folder),
                _FolderItem(title: 'Legal & Compliance', subtitle: '89 Documents', icon: LucideIcons.folder),
                _FolderItem(title: 'Marketing', subtitle: '210 Documents', icon: LucideIcons.folder),
                const SizedBox(height: 24),
                Text('Recent Files', style: Theme.of(context).textTheme.bodyLarge?.copyWith(fontWeight: FontWeight.bold)),
                const SizedBox(height: 16),
                _FileItem(title: 'Q3-Financial-Report.pdf', size: '2.4 MB', date: '2 days ago'),
                _FileItem(title: 'Employee-Handbook-2024.pdf', size: '1.8 MB', date: '5 days ago'),
              ],
            ),
          ),
        ],
      ),
    );
  }
}

class _FolderItem extends StatelessWidget {
  final String title;
  final String subtitle;
  final IconData icon;

  const _FolderItem({required this.title, required this.subtitle, required this.icon});

  @override
  Widget build(BuildContext context) {
    return ListTile(
      contentPadding: const EdgeInsets.symmetric(vertical: 4),
      leading: Container(
        padding: const EdgeInsets.all(10),
        decoration: BoxDecoration(
          color: Colors.amber.withOpacity(0.1),
          borderRadius: BorderRadius.circular(12),
        ),
        child: Icon(icon, color: Colors.amber, size: 24),
      ),
      title: Text(title, style: const TextStyle(fontWeight: FontWeight.bold)),
      subtitle: Text(subtitle, style: const TextStyle(color: AppTheme.textSecondary, fontSize: 12)),
      trailing: const Icon(LucideIcons.chevronRight, color: AppTheme.textSecondary, size: 16),
      onTap: () {},
    );
  }
}

class _FileItem extends StatelessWidget {
  final String title;
  final String size;
  final String date;

  const _FileItem({required this.title, required this.size, required this.date});

  @override
  Widget build(BuildContext context) {
    return ListTile(
      contentPadding: const EdgeInsets.symmetric(vertical: 4),
      leading: Container(
        padding: const EdgeInsets.all(10),
        decoration: BoxDecoration(
          color: AppTheme.primary.withOpacity(0.1),
          borderRadius: BorderRadius.circular(12),
        ),
        child: const Icon(LucideIcons.fileText, color: AppTheme.primary, size: 24),
      ),
      title: Text(title, style: const TextStyle(fontWeight: FontWeight.bold)),
      subtitle: Text('$size • $date', style: const TextStyle(color: AppTheme.textSecondary, fontSize: 12)),
      trailing: const Icon(LucideIcons.moreVertical, color: AppTheme.textSecondary, size: 16),
      onTap: () {},
    );
  }
}
