# Concurrent Programming
Sequential program adalah sebuah program dimana sebuah task harus diselesaikan terlebih dahulu sebelum menjalankan task selanjutnya dijalankan. Parallel program adalah sebuah program dimana beberapa task berjalan secara bersamaan. Sedangkan Concurrent program adalah sebuah program dimana beberapa task dapat dijalankan "Independently" dapat berjalan sendiri dan dapat muncul secara bersamaan


### Dalam Bahasa golang terdapat 3 fitur concurenccy yaitu :
- Goroutines
Goroutines adalah function atau method yang ringan dan independen yang berjalan bersamaan.
- Channels
Channel adalah sebuah objek yang dapat digunakan untuk mengirim dan menerima data antara goroutine.
- Select
memungkinkan pemrogram untuk mengontrol data antara beberapa operasi channel yang siap digunakan.

### Race Condition
Race condition adalah kondisi di mana beberapa proses atau thread bersaing untuk mengakses dan memodifikasi suatu memori yang sama secara bersamaan. Hal ini menyebabkan program menjadi unsyncronize. Untuk mencegah race condition, program perlu menggunakan teknik sinkronisasi seperti locking/blocking dengan waitgroup, channel blocking atau mutex
