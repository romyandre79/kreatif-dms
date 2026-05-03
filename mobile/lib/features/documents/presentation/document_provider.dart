import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:kreatif_dms/features/documents/data/document_repository.dart';
import 'package:kreatif_dms/features/documents/domain/document_models.dart';

final topologyProvider = FutureProvider<List<TopologyNode>>((ref) async {
  return ref.watch(documentRepositoryProvider).getTopology();
});

final recentDocumentsProvider = FutureProvider<List<Document>>((ref) async {
  return ref.watch(documentRepositoryProvider).getRecentDocuments();
});
