# Summary Intro RESTful API

## API
Application Programming Interface adalah sebuah kumpulan dari fungsi (function) dan prosedur yang memungkinkan program komputer berkomunikasi dengan program atau komputer lainnya.

## REST API
REST API (REpresentational State Transfer Application Programming Interface) adalah salah satu jenis API yang sering digunakan. Salah satu hal yang membuat API ini sering digunakan yaitu REST API ini berkomunikasi dengan protokol HTTP dan menggunakan HTTP method yang sesuai seperti GET, POST, PUT , DELETE dalam mengirimkan data.

## HTTP RESPONSE CODE
Ketika sebuah API mengirimkan data dengan methodnya masing-masing akan menerima response code untuk mengindikasikan apakah API tersebut berhasil dieksekusi atau tidak. 


Berikut berbagai macam HTTP Response Code yang umum digunakan:
- 200 - OK
Response ini berarti API berhasil dieksekusi.
- 201 - Created
Response ini berarti API berhasil dieksekusi dan data yang dikirim berhasil dibuat.
- 400 - Bad Request
Response ini berarti eksekusi API tidak berjalan dengan baik, terdapat kesalahan pada uri yang dieksekusi.
- 404 - Not Found
Response ini berarti eksekusi API tidak dapat ditemukan atau data yang dicari tidak ditemukan.
- 401 - Unauthorized
Response ini berarti eksekusi API tidak diizinkan atau tidak dapat diakses sembarangan.
- 405 - Method Not Allowed
Response ini berarti API dengan method HTTP yang sedang dijalankan tidak diizinkan untuk dieksekusi.
- 500 - Internal Server Error
Response ini berarti terjadi kesalahan di sisi Server.
