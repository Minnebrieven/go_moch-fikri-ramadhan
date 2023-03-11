# Soal Clean Code
## Tugas 1
Kode berikut ini dituliskan tanpa mengikuti kebiasaan penulisan yang disarankan. 
- Berapa banyak kekurangan dalam penulisan kode tersebut? 

  Total terdapat 10 kesalahan
- Bagian mana saja terjadi kekurangan tersebut? 

  2 Penamaan struct, 4 field_data dari struct dan 2 method yang tidak mengikuti konvensi umum serta 1 penamaan variable yang tidak menggambarkan isi variable tersebut dan return value memberikan nilai kosong disarankan untuk return value "nil" atau error
- Tuliskan alasan dari setiap kekurangan tersebut. Alasan bisa diberikan dalam bentuk komentar pada kode yang disediakan berikut.
```golang
package main

type user struct {    //penamaan struct yang tidak mengikuti konvensi umum fix=>User
  id       int    //penamaan field yang tidak mengikuti konvensi umum fix=>ID
  username int    //penamaan field yang tidak mengikuti konvensi umum fix=>Username
  password int    //penamaan field yang tidak mengikuti konvensi umum fix=>Password
}

type userservice struct {  //penamaan struct yang tidak mengikuti konvensi umum fix=>UserService
  t []user        //penamaan field yang tidak mendeskripsikan isi field tersebut fix=>users []User
}

func (u userservice) getallusers() []user {  //penamaan method yang tidak mengikuti konvensi umum fix=>GetAllUser()
  return u.t
}

func (u userservice) getuserbyid(id int) user {  //penamaan method yang tidak mengikuti konvensi umum fix=>GetUserById()
  for _, r := range u.t { //penamaan variable r yang tidak mendeskripsikan isi variable tersebut
    if id == r.id {
      return r    
    }
  }

  return user{} //pada method ini jika case data tidak ditemukan, method ini mengembalikan nilai kosong. disarankan untuk mengembalikan nilai "nil" atau error
}
```
## Tugas 2
Kode berikut ini dituliskan tanpa mengikuti kebiasaan penulisan yang disarankan. Ubahlah penulisan kode berikut menjadi lebih terbaca dan lebih rapi!
```golang
package main

type kendaraan struct {
 totalroda       int
 kecepatanperjam int
}

type mobil struct {
 kendaraan
}

func (m *mobil) berjalan() {
 m.tambahkecepatan(10)
}

func (m *mobil) tambahkecepatan(kecepatanbaru int) {
 m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru
}

func main() {
 mobilcepat := mobil{}
 mobilcepat.berjalan()
 mobilcepat.berjalan()
 mobilcepat.berjalan()

 mobillamban := mobil{}
 mobillamban.berjalan()
}
```
#### Berikut Refactoring dari Code diatas

```golang
package main

type Vehicle struct {
  TotalWheels  int
  SpeedPerHour int
}

type Car struct {
  Vehicle
}

func (c *Car) Move() {
  c.IncreseSpeed(10)
}

func (c *Car) IncreseSpeed(newSpeed int) {
  c.SpeedPerHour += newSpeed //current speed plus the new speed
}

func main() {
  fastCar := Car{}
  fastCar.Move()
  fastCar.Move()
  fastCar.Move()

  slowCar := Car{}
  slowCar.Move()
}

```
