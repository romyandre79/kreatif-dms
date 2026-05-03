import 'dart:ui';
import 'package:flutter/material.dart';
import 'package:flutter_animate/flutter_animate.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:lucide_icons/lucide_icons.dart';
import 'package:kreatif_dms/core/theme/app_theme.dart';
import 'package:kreatif_dms/features/auth/presentation/auth_provider.dart';
import 'package:kreatif_dms/features/dashboard/presentation/dashboard_screen.dart';

class LoginScreen extends ConsumerStatefulWidget {
  const LoginScreen({super.key});

  @override
  ConsumerState<LoginScreen> createState() => _LoginScreenState();
}

class _LoginScreenState extends ConsumerState<LoginScreen> {
  final _usernameController = TextEditingController();
  final _passwordController = TextEditingController();
  bool _obscurePassword = true;

  void _handleLogin() async {
    final identifier = _usernameController.text.trim();
    final password = _passwordController.text.trim();

    if (identifier.isEmpty || password.isEmpty) {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Username dan Password wajib diisi')),
      );
      return;
    }

    await ref.read(authProvider.notifier).login(identifier, password);
  }

  @override
  Widget build(BuildContext context) {
    final authState = ref.watch(authProvider);

    // Listen for errors
    ref.listen(authProvider, (previous, next) {
      if (next.error != null) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(next.error!),
            backgroundColor: Colors.redAccent,
          ),
        );
      }
      if (next.user != null) {
        // Handled reactively by main.dart
      }
    });

    return Scaffold(
      body: Stack(
        children: [
          // Background Gradients
          Positioned(
            top: -100,
            right: -50,
            child: Container(
              width: 300,
              height: 300,
              decoration: BoxDecoration(
                shape: BoxShape.circle,
                color: AppTheme.primary.withOpacity(0.2),
              ),
            ).animate().fadeIn(duration: 800.ms).scale(begin: const Offset(0.5, 0.5)),
          ),
          Positioned(
            bottom: -50,
            left: -50,
            child: Container(
              width: 250,
              height: 250,
              decoration: BoxDecoration(
                shape: BoxShape.circle,
                color: AppTheme.secondary.withOpacity(0.15),
              ),
            ).animate().fadeIn(duration: 1000.ms).scale(begin: const Offset(0.5, 0.5)),
          ),

          SafeArea(
            child: SingleChildScrollView(
              padding: const EdgeInsets.symmetric(horizontal: 24.0),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  const SizedBox(height: 60),
                  // Logo Placeholder
                  Container(
                    width: 64,
                    height: 64,
                    decoration: BoxDecoration(
                      color: AppTheme.primary,
                      borderRadius: BorderRadius.circular(16),
                      boxShadow: [
                        BoxShadow(
                          color: AppTheme.primary.withOpacity(0.4),
                          blurRadius: 20,
                          offset: const Offset(0, 8),
                        ),
                      ],
                    ),
                    child: const Icon(LucideIcons.fileStack, color: Colors.white, size: 32),
                  ).animate().fadeIn(duration: 600.ms).moveY(begin: 20, end: 0),
                  
                  const SizedBox(height: 32),
                  Text(
                    'Selamat Datang',
                    style: Theme.of(context).textTheme.headlineMedium,
                  ).animate().fadeIn(delay: 200.ms).moveX(begin: -20, end: 0),
                  
                  const SizedBox(height: 8),
                  Text(
                    'Masuk untuk mengelola dokumen Anda',
                    style: Theme.of(context).textTheme.bodyMedium,
                  ).animate().fadeIn(delay: 400.ms),

                  const SizedBox(height: 48),
                  
                  // Form
                  Column(
                    children: [
                      TextField(
                        controller: _usernameController,
                        decoration: const InputDecoration(
                          hintText: 'Email / Username',
                          prefixIcon: Icon(LucideIcons.user, size: 20),
                        ),
                      ).animate().fadeIn(delay: 600.ms),
                      
                      const SizedBox(height: 16),
                      
                      TextField(
                        controller: _passwordController,
                        obscureText: _obscurePassword,
                        decoration: InputDecoration(
                          hintText: 'Password',
                          prefixIcon: const Icon(LucideIcons.lock, size: 20),
                          suffixIcon: IconButton(
                            icon: Icon(
                              _obscurePassword ? LucideIcons.eye : LucideIcons.eyeOff,
                              size: 20,
                            ),
                            onPressed: () {
                              setState(() {
                                _obscurePassword = !_obscurePassword;
                              });
                            },
                          ),
                        ),
                      ).animate().fadeIn(delay: 800.ms),
                      
                      const SizedBox(height: 12),
                      Align(
                        alignment: Alignment.centerRight,
                        child: TextButton(
                          onPressed: () {},
                          child: const Text('Lupa Password?'),
                        ),
                      ).animate().fadeIn(delay: 900.ms),
                      
                      const SizedBox(height: 32),
                      
                      ElevatedButton(
                        onPressed: authState.isLoading ? null : _handleLogin,
                        child: authState.isLoading
                            ? const SizedBox(
                                height: 24,
                                width: 24,
                                child: CircularProgressIndicator(
                                  strokeWidth: 2,
                                  color: Colors.white,
                                ),
                              )
                            : const Text('Login'),
                      ).animate().fadeIn(delay: 1000.ms).scale(),
                    ],
                  ),
                  
                  const SizedBox(height: 100),
                  Center(
                    child: Text(
                      '© 2024 Kreatif DMS - Premium Edition',
                      style: Theme.of(context).textTheme.bodySmall,
                    ),
                  ),
                  const SizedBox(height: 24),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}
