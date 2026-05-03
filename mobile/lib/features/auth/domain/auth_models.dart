class User {
  final String id;
  final String fullName;
  final String email;
  final String role;
  final String? avatarUrl;
  final String? signatureUrl;
  final String? departmentId;
  final String? branchId;
  final String? companyId;
  final String? departmentName;
  final String? branchName;
  final String? companyName;
  final bool isMfaEnabled;

  User({
    required this.id,
    required this.fullName,
    required this.email,
    required this.role,
    this.avatarUrl,
    this.signatureUrl,
    this.departmentId,
    this.branchId,
    this.companyId,
    this.departmentName,
    this.branchName,
    this.companyName,
    this.isMfaEnabled = false,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    String? getValue(dynamic field) {
      if (field == null) return null;
      if (field is Map && field.containsKey('String')) {
        final val = field['String'].toString();
        return val.isEmpty ? null : val;
      }
      final val = field.toString();
      return val.isEmpty ? null : val;
    }

    bool getBool(dynamic field) {
      if (field == null) return false;
      if (field is Map && field.containsKey('Bool')) {
        return field['Bool'] == true;
      }
      return field == true;
    }

    return User(
      id: json['id'] ?? json['user_id'],
      fullName: json['full_name'],
      email: json['email'] ?? '',
      role: json['role_name'] ?? json['role'],
      avatarUrl: getValue(json['avatar_url']),
      signatureUrl: getValue(json['signature_url']),
      departmentId: getValue(json['department_id']),
      branchId: getValue(json['branch_id']),
      companyId: getValue(json['company_id']),
      departmentName: getValue(json['department_name']),
      branchName: getValue(json['branch_name']),
      companyName: getValue(json['company_name']),
      isMfaEnabled: getBool(json['is_mfa_enabled']),
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
