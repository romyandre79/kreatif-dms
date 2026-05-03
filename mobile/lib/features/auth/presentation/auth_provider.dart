import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:kreatif_dms/features/auth/data/auth_repository.dart';
import 'package:kreatif_dms/features/auth/domain/auth_models.dart';

class AuthState {
  final User? user;
  final bool isLoading;
  final String? error;

  AuthState({this.user, this.isLoading = false, this.error});

  AuthState copyWith({User? user, bool? isLoading, String? error}) {
    return AuthState(
      user: user ?? this.user,
      isLoading: isLoading ?? this.isLoading,
      error: error,
    );
  }
}

final authProvider = StateNotifierProvider<AuthNotifier, AuthState>((ref) {
  return AuthNotifier(ref.watch(authRepositoryProvider));
});

class AuthNotifier extends StateNotifier<AuthState> {
  final AuthRepository _repository;
  final _storage = const FlutterSecureStorage();

  AuthNotifier(this._repository) : super(AuthState()) {
    _checkPersistence();
  }

  Future<void> _checkPersistence() async {
    // In a real app, you would fetch the user profile here using the saved token
    // For now, we just check if token exists
    final token = await _storage.read(key: 'jwt_token');
    if (token != null) {
      // Potentially fetch user profile
    }
  }

  Future<void> login(String identifier, String password) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final response = await _repository.login(
        identifier: identifier,
        password: password,
      );

      if (response.mfaRequired) {
        state = state.copyWith(isLoading: false, error: 'MFA Required (Not yet implemented)');
        return;
      }

      if (response.accessToken != null) {
        await _storage.write(key: 'jwt_token', value: response.accessToken);
        await _storage.write(key: 'refresh_token', value: response.refreshToken);
        state = state.copyWith(isLoading: false, user: response.user);
      }
    } catch (e) {
      state = state.copyWith(isLoading: false, error: e.toString());
    }
  }

  Future<void> logout() async {
    await _storage.delete(key: 'jwt_token');
    await _storage.delete(key: 'refresh_token');
    state = AuthState();
  }
}
