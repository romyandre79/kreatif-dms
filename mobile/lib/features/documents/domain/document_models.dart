class TopologyNode {
  final String id;
  final String name;
  final String type;
  final List<TopologyNode>? children;

  TopologyNode({
    required this.id,
    required this.name,
    required this.type,
    this.children,
  });

  factory TopologyNode.fromJson(Map<String, dynamic> json) {
    return TopologyNode(
      id: json['id'] ?? '',
      name: json['name'],
      type: json['type'],
      children: json['children'] != null
          ? (json['children'] as List)
              .map((i) => TopologyNode.fromJson(i))
              .toList()
          : null,
    );
  }
}

class Document {
  final String id;
  final String title;
  final String? description;
  final String category;
  final String status;
  final Map<String, dynamic> metadata;
  final String? fileUrl;
  final DateTime createdAt;

  Document({
    required this.id,
    required this.title,
    this.description,
    required this.category,
    required this.status,
    required this.metadata,
    this.fileUrl,
    required this.createdAt,
  });

  factory Document.fromJson(Map<String, dynamic> json) {
    return Document(
      id: json['id'],
      title: json['title'] ?? json['name'] ?? 'Untitled',
      description: json['description'],
      category: json['category'] ?? 'General',
      status: json['status'] ?? 'active',
      metadata: json['metadata'] ?? {},
      fileUrl: json['file_url'],
      createdAt: DateTime.parse(json['created_at'] ?? DateTime.now().toIso8601String()),
    );
  }
}
