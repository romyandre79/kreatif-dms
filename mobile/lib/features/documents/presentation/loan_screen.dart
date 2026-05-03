import 'package:flutter/material.dart';
import 'package:lucide_icons/lucide_icons.dart';
import 'package:kreatif_dms/core/theme/app_theme.dart';

class LoanScreen extends StatelessWidget {
  const LoanScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Peminjaman', style: TextStyle(fontWeight: FontWeight.bold)),
        backgroundColor: Colors.transparent,
        elevation: 0,
      ),
      body: DefaultTabController(
        length: 2,
        child: Column(
          children: [
            const TabBar(
              tabs: [
                Tab(text: 'Aktif'),
                Tab(text: 'Riwayat'),
              ],
              indicatorColor: AppTheme.primary,
              labelColor: AppTheme.primary,
              unselectedLabelColor: AppTheme.textSecondary,
              labelStyle: TextStyle(fontWeight: FontWeight.bold),
            ),
            Expanded(
              child: TabBarView(
                children: [
                  _ActiveLoansList(),
                  const Center(child: Text('Belum ada riwayat peminjaman')),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}

class _ActiveLoansList extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return ListView(
      padding: const EdgeInsets.all(24),
      children: [
        _LoanItem(
          title: 'Financial Audit Q3 2023',
          id: 'LN-2023-45',
          date: 'Sep 20, 2023',
          returnDate: 'Sep 30, 2023',
          status: 'ACTIVE',
        ),
        _LoanItem(
          title: 'Project Alpha Requirements',
          id: 'LN-2023-52',
          date: 'Oct 01, 2023',
          returnDate: 'Oct 05, 2023',
          status: 'PENDING',
        ),
      ],
    );
  }
}

class _LoanItem extends StatelessWidget {
  final String title;
  final String id;
  final String date;
  final String returnDate;
  final String status;

  const _LoanItem({
    required this.title,
    required this.id,
    required this.date,
    required this.returnDate,
    required this.status,
  });

  @override
  Widget build(BuildContext context) {
    final bool isPending = status == 'PENDING';

    return Container(
      margin: const EdgeInsets.only(bottom: 16),
      padding: const EdgeInsets.all(20),
      decoration: BoxDecoration(
        color: AppTheme.cardBg,
        borderRadius: BorderRadius.circular(20),
        border: isPending ? Border.all(color: Colors.orangeAccent.withOpacity(0.3)) : null,
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Text(id, style: const TextStyle(color: AppTheme.textSecondary, fontSize: 10, fontWeight: FontWeight.bold)),
              Container(
                padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 4),
                decoration: BoxDecoration(
                  color: isPending ? Colors.orangeAccent.withOpacity(0.1) : Colors.greenAccent.withOpacity(0.1),
                  borderRadius: BorderRadius.circular(8),
                ),
                child: Text(
                  status,
                  style: TextStyle(
                    color: isPending ? Colors.orangeAccent : Colors.greenAccent,
                    fontSize: 10,
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ),
            ],
          ),
          const SizedBox(height: 12),
          Text(title, style: const TextStyle(fontSize: 16, fontWeight: FontWeight.bold)),
          const SizedBox(height: 16),
          Row(
            children: [
              const Icon(LucideIcons.calendar, size: 14, color: AppTheme.textSecondary),
              const SizedBox(width: 8),
              Text('Pinjam: $date', style: const TextStyle(color: AppTheme.textSecondary, fontSize: 12)),
            ],
          ),
          const SizedBox(height: 4),
          Row(
            children: [
              const Icon(LucideIcons.clock, size: 14, color: AppTheme.textSecondary),
              const SizedBox(width: 8),
              Text('Kembali: $returnDate', style: const TextStyle(color: AppTheme.textSecondary, fontSize: 12)),
            ],
          ),
        ],
      ),
    );
  }
}
