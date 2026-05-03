import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:lucide_icons/lucide_icons.dart';
import 'package:kreatif_dms/core/theme/app_theme.dart';
import 'package:kreatif_dms/features/documents/domain/document_models.dart';
import 'package:kreatif_dms/features/documents/presentation/document_provider.dart';

class DocumentExplorerScreen extends ConsumerWidget {
  const DocumentExplorerScreen({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final topologyAsync = ref.watch(topologyProvider);

    return Scaffold(
      appBar: AppBar(
        title: const Text('Explorer', style: TextStyle(fontWeight: FontWeight.bold)),
        backgroundColor: Colors.transparent,
        elevation: 0,
        actions: [
          IconButton(onPressed: () => ref.refresh(topologyProvider), icon: const Icon(LucideIcons.refreshCw, size: 20)),
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
            child: topologyAsync.when(
              data: (nodes) => ListView.builder(
                padding: const EdgeInsets.all(16),
                itemCount: nodes.length,
                itemBuilder: (context, index) {
                  final node = nodes[index];
                  return _TopologyItem(node: node);
                },
              ),
              loading: () => const Center(child: CircularProgressIndicator()),
              error: (err, stack) => Center(child: Text('Error: $err')),
            ),
          ),
        ],
      ),
    );
  }
}

class _TopologyItem extends StatelessWidget {
  final TopologyNode node;

  const _TopologyItem({required this.node});

  @override
  Widget build(BuildContext context) {
    final isFolder = node.type != 'document';

    return ListTile(
      contentPadding: const EdgeInsets.symmetric(vertical: 4),
      leading: Container(
        padding: const EdgeInsets.all(10),
        decoration: BoxDecoration(
          color: isFolder ? Colors.amber.withOpacity(0.1) : AppTheme.primary.withOpacity(0.1),
          borderRadius: BorderRadius.circular(12),
        ),
        child: Icon(
          isFolder ? LucideIcons.folder : LucideIcons.fileText, 
          color: isFolder ? Colors.amber : AppTheme.primary, 
          size: 24
        ),
      ),
      title: Text(node.name, style: const TextStyle(fontWeight: FontWeight.bold)),
      subtitle: Text(
        node.type.toUpperCase(), 
        style: const TextStyle(color: AppTheme.textSecondary, fontSize: 10, fontWeight: FontWeight.bold)
      ),
      trailing: isFolder ? const Icon(LucideIcons.chevronRight, color: AppTheme.textSecondary, size: 16) : null,
      onTap: () {
        if (isFolder && node.children != null && node.children!.isNotEmpty) {
          Navigator.of(context).push(
            MaterialPageRoute(
              builder: (_) => _FolderDetailScreen(node: node),
            ),
          );
        }
      },
    );
  }
}

class _FolderDetailScreen extends StatelessWidget {
  final TopologyNode node;

  const _FolderDetailScreen({required this.node});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(node.name),
        backgroundColor: Colors.transparent,
      ),
      body: ListView.builder(
        padding: const EdgeInsets.all(16),
        itemCount: node.children?.length ?? 0,
        itemBuilder: (context, index) {
          final child = node.children![index];
          return _TopologyItem(node: child);
        },
      ),
    );
  }
}
