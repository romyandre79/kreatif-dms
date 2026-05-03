import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:kreatif_dms/features/auth/data/auth_repository.dart';
import 'package:kreatif_dms/features/auth/domain/auth_models.dart';

class AuthState {
  final User? user;
  final String? accessToken;
  final bool isLoading;
  final String? error;

  AuthState({this.user, this.accessToken, this.isLoading = false, this.error});

  AuthState copyWith({User? user, String? accessToken, bool? isLoading, String? error}) {
    return AuthState(
      user: user ?? this.user,
      accessToken: accessToken ?? this.accessToken,
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
    final token = await _storage.read(key: 'jwt_token');
    if (token != null) {
      state = state.copyWith(accessToken: token);
      await loadProfile();
    }
  }

  Future<void> loadProfile() async {
    state = state.copyWith(isLoading: true);
    try {
      final user = await _repository.getProfile();
      state = state.copyWith(isLoading: false, user: user);
    } catch (e) {
      state = state.copyWith(isLoading: false, error: e.toString());
      // If token is invalid, logout
      if (e.toString().contains('Unauthorized')) {
        logout();
      }
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
        
        state = state.copyWith(accessToken: response.accessToken);
        
        // After login, fetch the full profile to get detailed info (Department, etc)
        await loadProfile();
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

  Future<void> setPIN(String pin) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      await _repository.setPIN(pin);
      state = state.copyWith(isLoading: false);
    } catch (e) {
      state = state.copyWith(isLoading: false, error: e.toString());
      rethrow;
    }
  }

  Future<void> changePassword(String oldPassword, String newPassword) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      await _repository.changePassword(oldPassword, newPassword);
      state = state.copyWith(isLoading: false);
    } catch (e) {
      state = state.copyWith(isLoading: false, error: e.toString());
      rethrow;
    }
  }

  Future<Map<String, dynamic>> setupMFA() async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final res = await _repository.setupMFA();
      state = state.copyWith(isLoading: false);
      return res;
    } catch (e) {
      state = state.copyWith(isLoading: false, error: e.toString());
      rethrow;
    }
  }

  Future<void> verifyMFA(String secret, String code) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      await _repository.verifyMFA(secret, code);
      await loadProfile(); // Refresh profile to see MFA enabled status
      state = state.copyWith(isLoading: false);
    } catch (e) {
      state = state.copyWith(isLoading: false, error: e.toString());
      rethrow;
    }
  }
}
