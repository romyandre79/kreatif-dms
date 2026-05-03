import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:lucide_icons/lucide_icons.dart';
import 'package:kreatif_dms/core/theme/app_theme.dart';
import 'package:kreatif_dms/features/notifications/data/notification_repository.dart';

class NotificationScreen extends ConsumerWidget {
  const NotificationScreen({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final notificationsAsync = ref.watch(notificationsProvider);

    return Scaffold(
      appBar: AppBar(
        title: const Text('Notifikasi', style: TextStyle(fontWeight: FontWeight.bold)),
        actions: [
          TextButton(
            onPressed: () async {
              await ref.read(notificationRepositoryProvider).markAllAsRead();
              ref.refresh(notificationsProvider);
              ref.refresh(unreadCountProvider);
            },
            child: const Text('Baca Semua'),
          ),
        ],
      ),
      body: notificationsAsync.when(
        data: (notifications) {
          if (notifications.isEmpty) {
            return const Center(
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Icon(LucideIcons.bellOff, size: 64, color: AppTheme.textSecondary),
                  SizedBox(height: 16),
                  Text('Belum ada notifikasi', style: TextStyle(color: AppTheme.textSecondary)),
                ],
              ),
            );
          }

          return RefreshIndicator(
            onRefresh: () async {
              ref.refresh(notificationsProvider);
              ref.refresh(unreadCountProvider);
            },
            child: ListView.builder(
              padding: const EdgeInsets.all(16),
              itemCount: notifications.length,
              itemBuilder: (context, index) {
                final notif = notifications[index];
                return _NotificationItem(notif: notif);
              },
            ),
          );
        },
        loading: () => const Center(child: CircularProgressIndicator()),
        error: (err, stack) => Center(child: Text('Error: $err')),
      ),
    );
  }
}

class _NotificationItem extends ConsumerWidget {
  final dynamic notif;
  const _NotificationItem({required this.notif});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    Color iconColor;
    IconData icon;

    switch (notif.type) {
      case 'SUCCESS':
        iconColor = Colors.green;
        icon = LucideIcons.checkCircle2;
        break;
      case 'ERROR':
        iconColor = Colors.red;
        icon = LucideIcons.alertCircle;
        break;
      case 'WARNING':
        iconColor = Colors.orange;
        icon = LucideIcons.alertTriangle;
        break;
      default:
        iconColor = AppTheme.primary;
        icon = LucideIcons.info;
    }

    return GestureDetector(
      onTap: () async {
        if (!notif.isRead) {
          await ref.read(notificationRepositoryProvider).markAsRead(notif.id);
          ref.refresh(notificationsProvider);
          ref.refresh(unreadCountProvider);
        }
      },
      child: Container(
        margin: const EdgeInsets.only(bottom: 12),
        padding: const EdgeInsets.all(16),
        decoration: BoxDecoration(
          color: notif.isRead ? AppTheme.cardBg.withOpacity(0.5) : AppTheme.cardBg,
          borderRadius: BorderRadius.circular(16),
          border: notif.isRead ? null : Border.all(color: AppTheme.primary.withOpacity(0.3)),
        ),
        child: Row(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Container(
              padding: const EdgeInsets.all(10),
              decoration: BoxDecoration(
                color: iconColor.withOpacity(0.1),
                borderRadius: BorderRadius.circular(12),
              ),
              child: Icon(icon, color: iconColor, size: 24),
            ),
            const SizedBox(width: 16),
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Expanded(
                        child: Text(
                          notif.title,
                          style: TextStyle(
                            fontWeight: notif.isRead ? FontWeight.w500 : FontWeight.bold,
                            fontSize: 16,
                          ),
                        ),
                      ),
                      Text(
                        _formatDate(notif.createdAt),
                        style: const TextStyle(color: AppTheme.textSecondary, fontSize: 11),
                      ),
                    ],
                  ),
                  const SizedBox(height: 4),
                  Text(
                    notif.body,
                    style: TextStyle(
                      color: notif.isRead ? AppTheme.textSecondary : AppTheme.textPrimary,
                      fontSize: 14,
                    ),
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }

  String _formatDate(DateTime dt) {
    final now = DateTime.now();
    final diff = now.difference(dt);
    if (diff.inMinutes < 60) return '${diff.inMinutes}m';
    if (diff.inHours < 24) return '${diff.inHours}j';
    return '${dt.day}/${dt.month}';
  }
}
