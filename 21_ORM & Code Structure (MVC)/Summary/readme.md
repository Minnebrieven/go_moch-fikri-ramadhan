# ORM & MVC Summary
# ORM
ORM adalah kepanjangan dari Object-Relational Mapping. ORM adalah sebuah teknik mengkonversi data dengan tipe yang incompatible dengan menggunakan OOP.

### GORM 
GORM adalah salah satu contoh ORM yang dapat digunakan di bahasa pemrograman GOlang. Dengan memanfaatkan type struct, GORM dapat membuat sebuah table berdasarkan struct tersebut.
- contoh struct 
```golang
type User struct {
  gorm.Model
  ID      uint64
  Name    string
  Age     int32
  Gender  string
}
```

### Kelebihan dan Kekurangan ORM
#### Kelebihan
1. Mengurangi Query yang repetitif.
2. Otomatis mengambil data menjadi object yang siap dipakai
3. Screening Data yang simple
4. Mempunyai fitur cache Query

#### Kekurangan
1. Cost overhead process
2. Memuat data relasi yang tidak dibutuhkan
3. Query yang complex sulit/panjang untuk ditulis dengan ORM
4. fungsi SQL yang spesifik mungkin tidak bekerja dengan beberapa vendor (contoh Force Index)

# MVC
MVC singkatan dari Model, View and Controller. MVC ini adalah sebuah metode mengorganisasi kode yang populer. Salah satu hal yang membuat metode organisasi ini populer adalah setiap section code yang dibuat mempunyai purpose (tujuan), dan setiap purpose (tujuan) berbeda-beda.

### Kenapa Harus Menggunakan MVC?
Untuk menciptakan sebuah aplikasi yang modular, memudahkan proses pengembangan ketika terjadi error dan lebih sedikit confict ketika proses versioning terjadi.
