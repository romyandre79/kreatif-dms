import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:lucide_icons/lucide_icons.dart';
import 'package:qr_flutter/qr_flutter.dart';
import 'package:url_launcher/url_launcher.dart';
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
                        child: (user?.avatarUrl != null && user!.avatarUrl!.startsWith('http'))
                          ? ClipOval(
                              child: Image.network(
                                user.avatarUrl!,
                                fit: BoxFit.cover,
                                width: 100,
                                height: 100,
                                errorBuilder: (context, error, stackTrace) => 
                                  const Icon(LucideIcons.user, size: 50, color: AppTheme.primary),
                              ),
                            )
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
              subtitle: user?.email ?? '-', 
              showChevron: false,
            ),
            _ProfileMenuItem(
              icon: LucideIcons.building, 
              title: 'Unit Kerja', 
              subtitle: '${user?.companyName ?? '-'} \n${user?.branchName ?? '-'}', 
              showChevron: false,
            ),
            _ProfileMenuItem(
              icon: LucideIcons.users, 
              title: 'Departemen', 
              subtitle: user?.departmentName ?? '-', 
              showChevron: false,
            ),
            
            const SizedBox(height: 32),
            
            // Section: Keamanan
            _buildSectionHeader('KEAMANAN'),
            _ProfileMenuItem(
              icon: LucideIcons.shield, 
              title: 'Ubah PIN Peminjaman', 
              onTap: () => _showPINDialog(context, ref)
            ),
            _ProfileMenuItem(
              icon: LucideIcons.lock, 
              title: 'Ubah Password', 
              onTap: () => _showPasswordDialog(context, ref)
            ),
            _ProfileMenuItem(
              icon: LucideIcons.smartphone, 
              title: 'Autentikasi Dua Faktor (MFA)', 
              subtitle: (user?.isMfaEnabled ?? false) ? 'Aktif' : 'Nonaktif', 
              onTap: () => _showMFADialog(context, ref)
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

  void _showPINDialog(BuildContext context, WidgetRef ref) {
    final controller = TextEditingController();
    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        backgroundColor: AppTheme.cardBg,
        title: const Text('Set PIN Peminjaman'),
        content: TextField(
          controller: controller,
          keyboardType: TextInputType.number,
          maxLength: 6,
          obscureText: true,
          decoration: const InputDecoration(
            labelText: 'PIN Baru (6 Digit)',
            hintText: 'Masukkan 6 angka',
          ),
        ),
        actions: [
          TextButton(onPressed: () => Navigator.pop(context), child: const Text('Batal')),
          ElevatedButton(
            onPressed: () async {
              if (controller.text.length != 6) return;
              try {
                await ref.read(authProvider.notifier).setPIN(controller.text);
                if (context.mounted) {
                  Navigator.pop(context);
                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(content: Text('PIN berhasil diperbarui')),
                  );
                }
              } catch (e) {
                if (context.mounted) {
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(content: Text(e.toString()), backgroundColor: Colors.redAccent),
                  );
                }
              }
            },
            child: const Text('Simpan'),
          ),
        ],
      ),
    );
  }

  void _showPasswordDialog(BuildContext context, WidgetRef ref) {
    final oldController = TextEditingController();
    final newController = TextEditingController();
    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        backgroundColor: AppTheme.cardBg,
        title: const Text('Ubah Password'),
        content: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            TextField(
              controller: oldController,
              obscureText: true,
              decoration: const InputDecoration(labelText: 'Password Lama'),
            ),
            TextField(
              controller: newController,
              obscureText: true,
              decoration: const InputDecoration(labelText: 'Password Baru'),
            ),
          ],
        ),
        actions: [
          TextButton(onPressed: () => Navigator.pop(context), child: const Text('Batal')),
          ElevatedButton(
            onPressed: () async {
              try {
                await ref.read(authProvider.notifier).changePassword(oldController.text, newController.text);
                if (context.mounted) {
                  Navigator.pop(context);
                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(content: Text('Password berhasil diperbarui')),
                  );
                }
              } catch (e) {
                if (context.mounted) {
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(content: Text(e.toString()), backgroundColor: Colors.redAccent),
                  );
                }
              }
            },
            child: const Text('Simpan'),
          ),
        ],
      ),
    );
  }

  void _showMFADialog(BuildContext context, WidgetRef ref) {
    final user = ref.read(authProvider).user;
    if (user?.isMfaEnabled ?? false) {
      showDialog(
        context: context,
        builder: (context) => AlertDialog(
          backgroundColor: AppTheme.cardBg,
          title: const Text('Status MFA'),
          content: const Text('Autentikasi Dua Faktor (MFA) sudah aktif di akun Anda.'),
          actions: [
            TextButton(onPressed: () => Navigator.pop(context), child: const Text('Tutup')),
          ],
        ),
      );
      return;
    }

    showDialog(
      context: context,
      builder: (context) => _MFASetupDialog(ref: ref),
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

class _MFASetupDialog extends StatefulWidget {
  final WidgetRef ref;
  const _MFASetupDialog({required this.ref});

  @override
  State<_MFASetupDialog> createState() => _MFASetupDialogState();
}

class _MFASetupDialogState extends State<_MFASetupDialog> {
  String? secret;
  String? url;
  bool loading = false;
  final codeController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      backgroundColor: AppTheme.cardBg,
      title: const Text('Setup MFA'),
      content: SizedBox(
        width: MediaQuery.of(context).size.width,
        child: secret == null 
        ? Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              const Text('Aktifkan MFA untuk keamanan tambahan. Klik tombol di bawah untuk mulai.'),
              const SizedBox(height: 16),
              SizedBox(
                width: double.infinity,
                child: ElevatedButton(
                  onPressed: loading ? null : () async {
                    setState(() => loading = true);
                    try {
                      final res = await widget.ref.read(authProvider.notifier).setupMFA();
                      setState(() {
                        secret = res['secret'];
                        url = res['url'];
                        loading = false;
                      });
                    } catch (e) {
                      setState(() => loading = false);
                      if (mounted) {
                        ScaffoldMessenger.of(context).showSnackBar(
                          SnackBar(content: Text(e.toString()), backgroundColor: Colors.redAccent),
                        );
                      }
                    }
                  },
                  child: loading 
                    ? const SizedBox(height: 20, width: 20, child: CircularProgressIndicator(strokeWidth: 2)) 
                    : const Text('Mulai Setup'),
                ),
              ),
            ],
          ) 
        : SingleChildScrollView(
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                const Text('Scan QR Code ini di aplikasi Authenticator Anda:'),
                const SizedBox(height: 16),
                Container(
                  padding: const EdgeInsets.all(12),
                  decoration: BoxDecoration(
                    color: Colors.white,
                    borderRadius: BorderRadius.circular(12),
                  ),
                  child: QrImageView(
                    data: url ?? '',
                    version: QrVersions.auto,
                    size: 200.0,
                  ),
                ),
                const SizedBox(height: 16),
                const Text('Atau masukkan secret key secara manual:'),
                const SizedBox(height: 8),
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Expanded(
                      child: SelectableText(
                        secret!, 
                        style: const TextStyle(fontSize: 14, fontWeight: FontWeight.bold, color: AppTheme.primary),
                        textAlign: TextAlign.center,
                      ),
                    ),
                    IconButton(
                      icon: const Icon(LucideIcons.copy, size: 16),
                      onPressed: () {
                        Clipboard.setData(ClipboardData(text: secret!));
                        ScaffoldMessenger.of(context).showSnackBar(
                          const SnackBar(content: Text('Secret key disalin!')),
                        );
                      },
                    ),
                  ],
                ),
                const SizedBox(height: 16),
                ElevatedButton.icon(
                  onPressed: () async {
                    final uri = Uri.parse(url ?? '');
                    if (await canLaunchUrl(uri)) {
                      await launchUrl(uri);
                    } else {
                      if (mounted) {
                        ScaffoldMessenger.of(context).showSnackBar(
                          const SnackBar(content: Text('Gagal membuka aplikasi Autentikator')),
                        );
                      }
                    }
                  },
                  icon: const Icon(LucideIcons.externalLink, size: 18),
                  label: const Text('Buka Aplikasi Autentikator'),
                  style: ElevatedButton.styleFrom(
                    backgroundColor: Colors.blueGrey,
                    foregroundColor: Colors.white,
                  ),
                ),
                const SizedBox(height: 32),
                const Text('Masukkan 6-digit kode verifikasi:'),
                const SizedBox(height: 8),
                TextField(
                  controller: codeController,
                  keyboardType: TextInputType.number,
                  maxLength: 6,
                  textAlign: TextAlign.center,
                  style: const TextStyle(fontSize: 24, fontWeight: FontWeight.bold, letterSpacing: 8),
                  decoration: const InputDecoration(
                    hintText: '000000',
                    counterText: '',
                  ),
                ),
              ],
            ),
          ),
      ),
      actions: [
        TextButton(onPressed: () => Navigator.pop(context), child: const Text('Batal')),
        if (secret != null)
          ElevatedButton(
            onPressed: loading ? null : () async {
              setState(() => loading = true);
              try {
                await widget.ref.read(authProvider.notifier).verifyMFA(secret!, codeController.text);
                if (mounted) {
                  Navigator.pop(context);
                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(content: Text('MFA berhasil diaktifkan')),
                  );
                }
              } catch (e) {
                setState(() => loading = false);
                if (mounted) {
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(content: Text(e.toString()), backgroundColor: Colors.redAccent),
                  );
                }
              }
            },
            child: loading 
              ? const SizedBox(height: 20, width: 20, child: CircularProgressIndicator(strokeWidth: 2, color: Colors.white))
              : const Text('Verifikasi & Aktifkan'),
          ),
      ],
    );
  }
}


class _ProfileMenuItem extends StatelessWidget {
  final IconData icon;
  final String title;
  final String? subtitle;
  final VoidCallback? onTap;
  final bool showChevron;

  const _ProfileMenuItem({
    required this.icon,
    required this.title,
    this.subtitle,
    this.onTap,
    this.showChevron = true,
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
      trailing: showChevron ? const Icon(LucideIcons.chevronRight, size: 16, color: AppTheme.textSecondary) : null,
      onTap: onTap,
    );
  }
}
