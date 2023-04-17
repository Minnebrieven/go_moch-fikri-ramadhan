# Summary Clean and Hexagonal Architecture
Tujuan dari arsitektur ini adalah memisahkan setiap concern (perhatian) dengan menggunakan layer agar dapat mengembangkan aplikasi yang modular, scalable, maintainable dan testable.

## Syarat sebelum membangun aplikasi dengan menggunakan Clean Architecture
1. Independent of Framework. 

    Tidak bergantung pada framework
2. Testable. 

    Dapat di test tanpa tools/hal lainnya.
3. Independent of UI. 

    Tidak bergantung pada UI. UI dapat diubah tanpa mengubah seluruh system yang ada
    
4. Independent of Database.
    
    Tidak bergantung pada database. Database dapat diubah dengan mudah tanpa membuat system berhenti.
    
5. Independent of External.

    Tidak bergantung pada tools external atau ThirdParty apapun sehingga system tidak berhenti berjalan ketika tools external atau ThirdParty tidak support lagi

## Benefit
1. Struktur Standart, sehingga ketika memulai project kita tidak perlu lagi mempelajari arsitektur yang lainnya.
2. Mempercepat pengembangan aplikasi dalam jangka waktu yang panjang.
3. Mocking dependencies menjadi hal yang biasa dalam unit testing.
4. Mudah dalam mengganti dari prototype ke solusi yang sebenarnya.

## Layer pada Clean Architecture
Pada umumnya terdapat 3 layer klasik yang biasa digunakan dalam Clean Architecture yaitu:
1. Use Case/Domain Layer.

Mengandung file-file logika bisnis.

2. Controller/Handler/Presentation Layer.

Berisi handler atau controller yang menangani user interaction yang menyediakan data ke layar

4. Drivers/Data Layer.

Berisi file yang mengatur data aplikasi, mengambil data dan mengatur data cache.
