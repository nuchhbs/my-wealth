import 'dart:convert';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;

class ApiClient {
  static String get baseUrl => dotenv.env['API_BASE_URL'] ?? 'http://localhost:8080';

  Future<dynamic> get(String path) async {
    final res = await http.get(Uri.parse('$baseUrl$path'));
    if (res.statusCode != 200) throw Exception('API error: ${res.statusCode}');
    return jsonDecode(res.body);
  }
}
