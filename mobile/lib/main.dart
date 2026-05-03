import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:kreatif_dms/core/theme/app_theme.dart';
import 'package:kreatif_dms/features/auth/presentation/auth_provider.dart';
import 'package:kreatif_dms/features/auth/presentation/login_screen.dart';
import 'package:kreatif_dms/features/dashboard/presentation/dashboard_screen.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  runApp(
    const ProviderScope(
      child: KreatifDMSApp(),
    ),
  );
}

class KreatifDMSApp extends ConsumerWidget {
  const KreatifDMSApp({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final authState = ref.watch(authProvider);

    return MaterialApp(
      title: 'Kreatif DMS',
      debugShowCheckedModeBanner: false,
      theme: AppTheme.darkTheme,
      home: authState.user != null ? const DashboardScreen() : const LoginScreen(),
    );
  }
}
