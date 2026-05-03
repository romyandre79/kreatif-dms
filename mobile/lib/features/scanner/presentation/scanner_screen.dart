import 'package:flutter/material.dart';
import 'package:mobile_scanner/mobile_scanner.dart';
import 'package:lucide_icons/lucide_icons.dart';
import 'package:kreatif_dms/core/theme/app_theme.dart';

class ScannerScreen extends StatefulWidget {
  const ScannerScreen({super.key});

  @override
  State<ScannerScreen> createState() => _ScannerScreenState();
}

class _ScannerScreenState extends State<ScannerScreen> {
  final MobileScannerController cameraController = MobileScannerController();

  @override
  void dispose() {
    cameraController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Scan QR Code', style: TextStyle(fontWeight: FontWeight.bold)),
        backgroundColor: Colors.transparent,
        elevation: 0,
        actions: [
          ValueListenableBuilder<MobileScannerState>(
            valueListenable: cameraController,
            builder: (context, state, child) {
              final IconData icon;
              final Color color;
              
              switch (state.torchState) {
                case TorchState.off:
                  icon = LucideIcons.flashlightOff;
                  color = Colors.grey;
                  break;
                case TorchState.on:
                  icon = LucideIcons.flashlight;
                  color = Colors.yellow;
                  break;
                default:
                  icon = LucideIcons.flashlightOff;
                  color = Colors.grey;
              }
              
              return IconButton(
                icon: Icon(icon, color: color),
                onPressed: () => cameraController.toggleTorch(),
              );
            },
          ),
          ValueListenableBuilder<MobileScannerState>(
            valueListenable: cameraController,
            builder: (context, state, child) {
              final IconData icon;
              
              switch (state.cameraDirection) {
                case CameraFacing.front:
                  icon = LucideIcons.camera;
                  break;
                case CameraFacing.back:
                  icon = LucideIcons.refreshCw;
                  break;
                default:
                  icon = LucideIcons.refreshCw;
              }
              
              return IconButton(
                icon: Icon(icon),
                onPressed: () => cameraController.switchCamera(),
              );
            },
          ),
        ],
      ),
      body: Stack(
        children: [
          MobileScanner(
            controller: cameraController,
            onDetect: (capture) {
              final List<Barcode> barcodes = capture.barcodes;
              for (final barcode in barcodes) {
                debugPrint('Barcode found! ${barcode.rawValue}');
                // Handle scanned barcode
              }
            },
          ),
          // Overlay
          Center(
            child: Container(
              width: 250,
              height: 250,
              decoration: BoxDecoration(
                border: Border.all(color: AppTheme.primary, width: 4),
                borderRadius: BorderRadius.circular(24),
              ),
              child: const Stack(
                children: [
                  Positioned(
                    top: 10,
                    left: 10,
                    child: Icon(LucideIcons.scanLine, color: AppTheme.primary, size: 24),
                  ),
                ],
              ),
            ),
          ),
          Positioned(
            bottom: 80,
            left: 0,
            right: 0,
            child: Center(
              child: Container(
                padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 12),
                decoration: BoxDecoration(
                  color: Colors.black54,
                  borderRadius: BorderRadius.circular(30),
                ),
                child: const Text(
                  'Arahkan kamera ke QR Code di Boks atau Rak',
                  style: TextStyle(color: Colors.white, fontSize: 12),
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
