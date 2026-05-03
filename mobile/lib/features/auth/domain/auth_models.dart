class User {
  final String id;
  final String fullName;
  final String role;
  final String? avatarUrl;
  final String? signatureUrl;

  User({
    required this.id,
    required this.fullName,
    required this.role,
    this.avatarUrl,
    this.signatureUrl,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json['user_id'],
      fullName: json['full_name'],
      role: json['role'],
      avatarUrl: json['avatar_url'],
      signatureUrl: json['signature_url'],
    );
  }
}

class LoginResponse {
  final String? accessToken;
  final String? refreshToken;
  final User? user;
  final bool mfaRequired;

  LoginResponse({
    this.accessToken,
    this.refreshToken,
    this.user,
    this.mfaRequired = false,
  });

  factory LoginResponse.fromJson(Map<String, dynamic> json) {
    return LoginResponse(
      accessToken: json['access_token'],
      refreshToken: json['refresh_token'],
      user: json['access_token'] != null ? User.fromJson(json) : null,
      mfaRequired: json['mfa_required'] ?? false,
    );
  }
}
