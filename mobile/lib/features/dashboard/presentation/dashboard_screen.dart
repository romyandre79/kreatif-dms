import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:lucide_icons/lucide_icons.dart';
import 'package:kreatif_dms/core/theme/app_theme.dart';
import 'package:kreatif_dms/features/auth/presentation/auth_provider.dart';
import 'package:kreatif_dms/features/dashboard/presentation/profile_screen.dart';
import 'package:kreatif_dms/features/documents/presentation/document_explorer_screen.dart';
import 'package:kreatif_dms/features/documents/presentation/loan_screen.dart';
import 'package:kreatif_dms/features/documents/presentation/upload_screen.dart';
import 'package:kreatif_dms/features/scanner/presentation/scanner_screen.dart';

class DashboardScreen extends StatefulWidget {
  const DashboardScreen({super.key});

  @override
  State<DashboardScreen> createState() => _DashboardScreenState();
}

class _DashboardScreenState extends State<DashboardScreen> {
  int _currentIndex = 0;

  final List<Widget> _pages = [
    const _HomeView(),
    const DocumentExplorerScreen(),
    const ProfileScreen(),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      bottomNavigationBar: BottomNavigationBar(
        backgroundColor: AppTheme.cardBg,
        selectedItemColor: AppTheme.primary,
        unselectedItemColor: AppTheme.textSecondary,
        currentIndex: _currentIndex,
        onTap: (index) {
          setState(() {
            _currentIndex = index;
          });
        },
        items: const [
          BottomNavigationBarItem(icon: Icon(LucideIcons.layoutDashboard), label: 'Home'),
          BottomNavigationBarItem(icon: Icon(LucideIcons.file), label: 'Documents'),
          BottomNavigationBarItem(icon: Icon(LucideIcons.user), label: 'Profile'),
        ],
      ),
      body: _pages[_currentIndex],
    );
  }
}

class _HomeView extends ConsumerWidget {
  const _HomeView();

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final user = ref.watch(authProvider).user;

    return SafeArea(
      child: SingleChildScrollView(
        padding: const EdgeInsets.all(24.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Header
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      'Halo,',
                      style: Theme.of(context).textTheme.bodyMedium,
                    ),
                    Text(
                      user?.fullName ?? 'User',
                      style: Theme.of(context).textTheme.headlineMedium?.copyWith(fontSize: 24),
                    ),
                  ],
                ),
                CircleAvatar(
                  radius: 24,
                  backgroundColor: AppTheme.primary.withOpacity(0.2),
                  child: const Icon(LucideIcons.user, color: AppTheme.primary),
                ),
              ],
            ),
            const SizedBox(height: 32),

            // Search Bar
            Container(
              padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
              decoration: BoxDecoration(
                color: AppTheme.cardBg,
                borderRadius: BorderRadius.circular(16),
              ),
              child: Row(
                children: [
                  const Icon(LucideIcons.search, color: AppTheme.textSecondary, size: 20),
                  const SizedBox(width: 12),
                  Text(
                    'Cari dokumen atau boks...',
                    style: Theme.of(context).textTheme.bodyMedium,
                  ),
                ],
              ),
            ),
            const SizedBox(height: 32),

            // Quick Actions
            Text(
              'Aksi Cepat',
              style: Theme.of(context).textTheme.bodyLarge?.copyWith(fontWeight: FontWeight.bold),
            ),
            const SizedBox(height: 16),
            Row(
              children: [
                _QuickActionItem(
                  icon: LucideIcons.scanLine,
                  label: 'Scan QR',
                  color: AppTheme.primary,
                  onTap: () {
                    Navigator.of(context).push(MaterialPageRoute(builder: (_) => const ScannerScreen()));
                  },
                ),
                const SizedBox(width: 16),
                _QuickActionItem(
                  icon: LucideIcons.upload,
                  label: 'Upload',
                  color: AppTheme.secondary,
                  onTap: () {
                    Navigator.of(context).push(MaterialPageRoute(builder: (_) => const UploadScreen()));
                  },
                ),
                const SizedBox(width: 16),
                _QuickActionItem(
                  icon: LucideIcons.bookOpen,
                  label: 'Pinjam',
                  color: Colors.orangeAccent,
                  onTap: () {
                    Navigator.of(context).push(MaterialPageRoute(builder: (_) => const LoanScreen()));
                  },
                ),
              ],
            ),
            const SizedBox(height: 32),

            // Stats Cards
            Container(
              width: double.infinity,
              padding: const EdgeInsets.all(20),
              decoration: BoxDecoration(
                gradient: LinearGradient(
                  colors: [AppTheme.primary, AppTheme.primary.withOpacity(0.7)],
                  begin: Alignment.topLeft,
                  end: Alignment.bottomRight,
                ),
                borderRadius: BorderRadius.circular(24),
                boxShadow: [
                  BoxShadow(
                    color: AppTheme.primary.withOpacity(0.3),
                    blurRadius: 20,
                    offset: const Offset(0, 10),
                  ),
                ],
              ),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  const Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Icon(LucideIcons.database, color: Colors.white, size: 28),
                      Icon(LucideIcons.chevronRight, color: Colors.white, size: 20),
                    ],
                  ),
                  const SizedBox(height: 20),
                  const Text(
                    'Total Dokumen',
                    style: TextStyle(color: Colors.white70, fontSize: 16),
                  ),
                  const Text(
                    '1,284',
                    style: TextStyle(
                      color: Colors.white,
                      fontSize: 32,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                  const SizedBox(height: 8),
                  Container(
                    padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 4),
                    decoration: BoxDecoration(
                      color: Colors.white24,
                      borderRadius: BorderRadius.circular(8),
                    ),
                    child: const Text(
                      '+12 hari ini',
                      style: TextStyle(color: Colors.white, fontSize: 12),
                    ),
                  ),
                ],
              ),
            ),
            const SizedBox(height: 32),

            // Recent Documents
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Text(
                  'Dokumen Terbaru',
                  style: Theme.of(context).textTheme.bodyLarge?.copyWith(fontWeight: FontWeight.bold),
                ),
                TextButton(onPressed: () {}, child: const Text('Lihat Semua')),
              ],
            ),
            const SizedBox(height: 8),
            const _RecentDocItem(
              title: 'Invoice-2024-001.pdf',
              subtitle: 'Finance • 2 jam yang lalu',
              icon: LucideIcons.fileText,
            ),
            const _RecentDocItem(
              title: 'Contract-Draft-v2.pdf',
              subtitle: 'Legal • 5 jam yang lalu',
              icon: LucideIcons.fileCode,
            ),
          ],
        ),
      ),
    );
  }
}

class _QuickActionItem extends StatelessWidget {
  final IconData icon;
  final String label;
  final Color color;
  final VoidCallback onTap;

  const _QuickActionItem({
    required this.icon,
    required this.label,
    required this.color,
    required this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    return Expanded(
      child: GestureDetector(
        onTap: onTap,
        child: Container(
          padding: const EdgeInsets.symmetric(vertical: 20),
          decoration: BoxDecoration(
            color: AppTheme.cardBg,
            borderRadius: BorderRadius.circular(20),
          ),
          child: Column(
            children: [
              Icon(icon, color: color, size: 28),
              const SizedBox(height: 12),
              Text(
                label,
                style: const TextStyle(fontSize: 12, fontWeight: FontWeight.w500),
              ),
            ],
          ),
        ),
      ),
    );
  }
}

class _RecentDocItem extends StatelessWidget {
  final String title;
  final String subtitle;
  final IconData icon;

  const _RecentDocItem({
    required this.title,
    required this.subtitle,
    required this.icon,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.only(bottom: 12),
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: AppTheme.cardBg.withOpacity(0.5),
        borderRadius: BorderRadius.circular(16),
      ),
      child: Row(
        children: [
          Container(
            padding: const EdgeInsets.all(10),
            decoration: BoxDecoration(
              color: AppTheme.primary.withOpacity(0.1),
              borderRadius: BorderRadius.circular(12),
            ),
            child: Icon(icon, color: AppTheme.primary, size: 24),
          ),
          const SizedBox(width: 16),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  title,
                  style: const TextStyle(fontWeight: FontWeight.w600),
                ),
                Text(
                  subtitle,
                  style: const TextStyle(color: AppTheme.textSecondary, fontSize: 12),
                ),
              ],
            ),
          ),
          const Icon(LucideIcons.moreVertical, color: AppTheme.textSecondary, size: 20),
        ],
      ),
    );
  }
}
