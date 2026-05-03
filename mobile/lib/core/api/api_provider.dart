import 'dart:io';
import 'package:dio/dio.dart';
import 'package:dio/io.dart'; // Import for IOHttpClientAdapter
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

final apiProvider = Provider<Dio>((ref) {
  // Use 10.0.2.2 for Android emulator to access host machine's localhost
  // Use localhost for iOS simulator or web
  final baseUrl = Platform.isAndroid 
      ? 'https://10.0.2.2:8080/api/v1' 
      : 'https://localhost:8080/api/v1';

  final dio = Dio(
    BaseOptions(
      baseUrl: baseUrl,
      connectTimeout: const Duration(seconds: 10),
      receiveTimeout: const Duration(seconds: 10),
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
    ),
  );

  // Allow self-signed certificates for development
  (dio.httpClientAdapter as IOHttpClientAdapter).createHttpClient = () {
    final client = HttpClient();
    client.badCertificateCallback = (X509Certificate cert, String host, int port) => true;
    return client;
  };

  dio.interceptors.add(InterceptorsWrapper(
    onRequest: (options, handler) async {
      const storage = FlutterSecureStorage();
      final token = await storage.read(key: 'jwt_token');
      if (token != null) {
        options.headers['Authorization'] = 'Bearer $token';
      }
      return handler.next(options);
    },
    onError: (DioException e, handler) {
      if (e.response?.statusCode == 401) {
        // Handle logout
      }
      return handler.next(e);
    },
  ));

  return dio;
});
