# Recursive
Recursive adalah state dari sebuah function yang menyelesaikan masalah dengan cara memanggil dirinya (function) itu sendiri berulang kali. Jika sebuah masalah tidak terlalu besar maka recursive akan menyelesaikan masalah dengan cepat. Dengan menggunakan recursive dapat memudahkan dan memperpendek code yang akan ditulis

#### Contoh Recursive : Faktorial
```golang
  func faktorial(n int) int {
    if n == 1{
      return 1
    } else {
      return n * faktorial(n - 1)
    }
  }
```

# Sorting
Sorting adalah proses mengurutkan data. Kita dapat mengurutkan nomor, kata, pairs dll. 

### Bubble Sort
Bubble sort merupakan salah satu metode pengurutan yang banyak digunakan. Bubble sort ini bekerja dengan melakukan pertukaran data secara terus menerus sampai iterasi yang dilakukan tidak terjadi lagi pertukaran

#### Syntax
```golang
for {
  var swapped bool
  for i := 0; i < len(sortArray); i++ {
    if i == len(sortArray)-1 {
      continue
    }
    leftElement, rightElement := sortArray[i], sortArray[i+1]
    if leftElement[1] > rightElement[1] {
      sortArray[i] = rightElement
      sortArray[i+1] = leftElement
      swapped = true
    }
  }
    if !swapped {
      break
    }
}
```

# Searching
Searching adalah sebuah proses mencari sebuah value yang diberikan didalam sebuah list/array dari value. Salah satu metode searching yang sering dipakai adalah binary search. Metode ini mencari posisi dari sebuah element/value yang berada didalam list/array yang telah terurut(sorted).
