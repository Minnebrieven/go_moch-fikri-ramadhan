# Data Structure
## Array
Array adalah data structure yang mengandung sebuah grup dari beberapa element, hanya dapat menyimpan satu tipe data dan ukuran alokasi memorinya tidak berubah (fixed). Berikut cara mendeklarasikan Array
```golang
import (
  "fmt"
  "reflect"
)

func main() {
  var primes [5]int         //Array dengan size = 5 yang mempunyai tipe data integer
  var countries [5]string   //Array dengan size = 5 yang mempunyai tipe data string
  var countries [2]string   //Array dengan size = 2 yang mempunyai tipe data string
  
  countries[0] = "Indonesia"  // Memasukan sebuah value ke element pertama 
  countries[1] = "Singapore" // Memasukan sebuah value ke element kedua
  
  odd_numbers := [5]int{1, 3, 5, 7, 9}   // Menginisialisasi dengan values
  var even_numbers [5]int = [5]int{0, 2, 4} // memasukan value sebagian kedalam element array
  
  
  fmt.Println(reflect.ValueOf(primes).Kind())
  fmt.Println(reflect.ValueOf(countries).Kind())
  
  fmt.Println(countries[0]) // Mengakses value element pertama
  fmt.Println(countries[1]) // Mengakses value element kedua
  
  fmt.Println(odd_numbers)
  fmt.Println(even_numbers)

  
}
```
### Output
```
array
array

Indonesia
Singapore

[1 3 5 7 9]
[0 2 4 0 0]
```


#### Iterasi Array menggunakan Loop
```golang
prima := [5]int{2, 3, 5}

// cara pertama
for index := 0; index < len(prima); index++ {
  fmt.Println(prima[index])
}
// cara kedua
for index, element := range prima {
  fmt.Println(index, "=>", element)
}
for _, value := range prima {
  fmt.Println(value)
}
// cara ketiga
index := 0
for range prima {
  fmt.Println(prima[index])
  index++
}
```
### Output
```
2
3
5
0
0
0 => 2
1 => 3
2 => 5
3 => 0
4 => 0
2
3
5
0
0
2
3
5
0
0
```

## Slice
Slice dan Array hampir sama tetapi terdapat salah satu hal yang membedakan yaitu, Slice alokasi penggunaan memori cenderung dynamis tidak seperti Array tidak dapat berubah (fixed). Untuk mendeklarasikan Slice sama seperti array yang membedakan Slice tidak memasukan ukuran memori yang digunakan didalam bracket []
```golang
import (
  "fmt"
  "reflect"
)

func main() {
  var primes = [5]int{2, 3, 5, 7, 11}     //Array

  // mendeklarasikan sebuah slice dengan mengambil beberapa element dari array
  var part_primes []int = primes[1:4]

  // part_primes = append(part_primes, 10000)
  // menambah data ke slice akan menambah data ke array juga

  fmt.Println(reflect.ValueOf(part_primes).Kind())
  fmt.Println(part_primes)
}
```
### Output
```
slice
[3 5 7]
```

## Map
Map adalah structure data yang menyimpan data dengan bentuk pasangan dari key dan value dimana setiap keynya itu unik, tidak dapat sama/duplicate
<picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://cdn.programiz.com/sites/tutorial2program/files/Map.png">
  <source media="(prefers-color-scheme: light)" srcset="https://cdn.programiz.com/sites/tutorial2program/files/Map.png">
  <img alt="Key and Values Pair" src="https://cdn.programiz.com/sites/tutorial2program/files/Map.png">
</picture>

### Cara mendeklarasikan Map
```golang
func main() {
  // cara pertama
  var gaji map[string]int{}
  fmt.Println(gaji)

  // cara kedua
  var gaji_a = map[string]int{"umam": 1000, "iswanul": 2000}
  fmt.Println(gaji_1)

  // cara ketiga
  gaji_b := map[string]int{}
  fmt.Println(gaji_b)

  // cara keempat menggunakan make()
  var gaji_c = make(map[string]int)
  gaji_c["doe"] = 7000 // memasukan value
  fmt.Println(gaji_c)
}
```
