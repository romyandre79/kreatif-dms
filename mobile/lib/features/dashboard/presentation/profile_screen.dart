import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:lucide_icons/lucide_icons.dart';
import 'package:kreatif_dms/core/theme/app_theme.dart';
import 'package:kreatif_dms/features/auth/presentation/auth_provider.dart';

class ProfileScreen extends ConsumerWidget {
  const ProfileScreen({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final user = ref.watch(authProvider).user;

    return Scaffold(
      appBar: AppBar(
        title: const Text('Profil Saya', style: TextStyle(fontWeight: FontWeight.bold)),
        backgroundColor: Colors.transparent,
        elevation: 0,
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(24.0),
        child: Column(
          children: [
            // Profile Header
            Center(
              child: Column(
                children: [
                  Stack(
                    children: [
                      CircleAvatar(
                        radius: 50,
                        backgroundColor: AppTheme.primary.withOpacity(0.1),
                        child: user?.avatarUrl != null 
                          ? ClipOval(child: Image.network(user!.avatarUrl!, fit: BoxFit.cover))
                          : const Icon(LucideIcons.user, size: 50, color: AppTheme.primary),
                      ),
                      Positioned(
                        bottom: 0,
                        right: 0,
                        child: Container(
                          padding: const EdgeInsets.all(4),
                          decoration: const BoxDecoration(
                            color: AppTheme.primary,
                            shape: BoxShape.circle,
                          ),
                          child: const Icon(LucideIcons.camera, size: 16, color: Colors.white),
                        ),
                      ),
                    ],
                  ),
                  const SizedBox(height: 16),
                  Text(
                    user?.fullName ?? 'User Name',
                    style: const TextStyle(fontSize: 22, fontWeight: FontWeight.bold),
                  ),
                  const SizedBox(height: 4),
                  Container(
                    padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 4),
                    decoration: BoxDecoration(
                      color: AppTheme.primary.withOpacity(0.1),
                      borderRadius: BorderRadius.circular(20),
                    ),
                    child: Text(
                      user?.role.toUpperCase() ?? 'ROLE',
                      style: const TextStyle(
                        color: AppTheme.primary, 
                        fontSize: 10, 
                        fontWeight: FontWeight.bold, 
                        letterSpacing: 1.2
                      ),
                    ),
                  ),
                ],
              ),
            ),
            
            const SizedBox(height: 40),
            
            // Section: Info Akun
            _buildSectionHeader('INFORMASI AKUN'),
            _ProfileMenuItem(
              icon: LucideIcons.mail, 
              title: 'Email', 
              subtitle: 'user@example.com', // Placeholder for now
              onTap: () {}
            ),
            _ProfileMenuItem(
              icon: LucideIcons.building, 
              title: 'Departemen', 
              subtitle: 'IT Operations', // Placeholder for now
              onTap: () {}
            ),
            
            const SizedBox(height: 32),
            
            // Section: Keamanan
            _buildSectionHeader('KEAMANAN'),
            _ProfileMenuItem(icon: LucideIcons.shield, title: 'Ubah PIN Peminjaman', onTap: () {}),
            _ProfileMenuItem(icon: LucideIcons.lock, title: 'Ubah Password', onTap: () {}),
            _ProfileMenuItem(
              icon: LucideIcons.smartphone, 
              title: 'Autentikasi Dua Faktor (MFA)', 
              subtitle: 'Nonaktif', 
              onTap: () {}
            ),
            
            const SizedBox(height: 32),
            
            // Section: Lainnya
            _buildSectionHeader('LAINNYA'),
            _ProfileMenuItem(icon: LucideIcons.helpCircle, title: 'Pusat Bantuan', onTap: () {}),
            _ProfileMenuItem(icon: LucideIcons.info, title: 'Tentang Aplikasi', subtitle: 'v1.0.0 (Premium)', onTap: () {}),
            
            const SizedBox(height: 40),
            
            // Logout
            SizedBox(
              width: double.infinity,
              child: OutlinedButton.icon(
                onPressed: () => _showLogoutDialog(context, ref),
                icon: const Icon(LucideIcons.logOut, size: 18),
                label: const Text('Keluar dari Akun', style: TextStyle(fontWeight: FontWeight.bold)),
                style: OutlinedButton.styleFrom(
                  foregroundColor: Colors.redAccent,
                  side: const BorderSide(color: Colors.redAccent, width: 1.5),
                  padding: const EdgeInsets.symmetric(vertical: 16),
                  shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
                ),
              ),
            ),
            const SizedBox(height: 40),
          ],
        ),
      ),
    );
  }

  Widget _buildSectionHeader(String title) {
    return Padding(
      padding: const EdgeInsets.only(left: 4, bottom: 12),
      child: Align(
        alignment: Alignment.centerLeft,
        child: Text(
          title,
          style: const TextStyle(
            color: AppTheme.textSecondary,
            fontSize: 10,
            fontWeight: FontWeight.bold,
            letterSpacing: 1.5,
          ),
        ),
      ),
    );
  }

  void _showLogoutDialog(BuildContext context, WidgetRef ref) {
    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        backgroundColor: AppTheme.cardBg,
        title: const Text('Konfirmasi Logout'),
        content: const Text('Apakah Anda yakin ingin keluar dari akun ini?'),
        actions: [
          TextButton(
            onPressed: () => Navigator.pop(context),
            child: const Text('Batal', style: TextStyle(color: AppTheme.textSecondary)),
          ),
          TextButton(
            onPressed: () {
              Navigator.pop(context);
              ref.read(authProvider.notifier).logout();
            },
            child: const Text('Logout', style: TextStyle(color: Colors.redAccent, fontWeight: FontWeight.bold)),
          ),
        ],
      ),
    );
  }
}

class _ProfileMenuItem extends StatelessWidget {
  final IconData icon;
  final String title;
  final String? subtitle;
  final VoidCallback onTap;

  const _ProfileMenuItem({
    required this.icon,
    required this.title,
    this.subtitle,
    required this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    return ListTile(
      contentPadding: const EdgeInsets.symmetric(horizontal: 4),
      leading: Container(
        padding: const EdgeInsets.all(10),
        decoration: BoxDecoration(
          color: AppTheme.cardBg,
          borderRadius: BorderRadius.circular(12),
        ),
        child: Icon(icon, size: 20, color: AppTheme.textSecondary),
      ),
      title: Text(title, style: const TextStyle(fontWeight: FontWeight.w600, fontSize: 14)),
      subtitle: subtitle != null ? Text(subtitle!, style: const TextStyle(fontSize: 11, color: AppTheme.textSecondary)) : null,
      trailing: const Icon(LucideIcons.chevronRight, size: 16, color: AppTheme.textSecondary),
      onTap: onTap,
    );
  }
}
