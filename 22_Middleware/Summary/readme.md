# Middleware Summary
middleware adalah sebuah entitas yang terhubung ke sebuah request/response server processing. Middleware dapat mengimplementasikan berbagai macam fungsionalitas diseputar HTTP Request yang akan masuk untuk server dan Response yang akan keluar nanti.

## Implementasi Middleware
1. LoggingMiddleware.
2. SessionMiddleware
3. AuthMiddleware
4. CustomMiddleware

## Log Middleware
fungsinya mencatat HTTP Request yang masuk selain itu sebagai jejak pengguna user. Tidak hanya itu log ini bisa menjadi sumber dara untuk analisis

## Auth Middleware
fungsinya mengidentifikasi user/pengguna selain itu berguna untuk mengamankan data dan informasi. Framework Echo sendiri memiliki beberapa Auth Middleware. dua diantaranya Basic Auth dan JWTAuth

1. Basic Auth


Basic Auth adalah sebuah teknik autentikasi http request dimana username dan password diletakan di request header

2. JWT Auth


JWT Auth ini adalah teknik yang paling banyak digunakan karena keaamannya. teknik autentikasi ini memerlukan kunci (secret-key) untuk proses decode dan encode token yang telah dibuat
