# Advance Function
Terdapat 3 Function Lanjutan yaitu Variadic function, Anonymous function dan Closure function. 


#### Variadic funtion 


digunakan ketika jumlah parameter yang akan diinput tidak diketehui maka Variadic funtion ini akan menginput list dari input parameter tersebut. 

#### Variadic function Declaration
```golang
func sum(numbers ...int) int {

  var total int = 0
  for _, number := range numbers {
    total += number
  }
  return total
}

```

#### Anonymous function 
adalah function yang tidak mempunyai nama, biasanya digunakan ketika ingin membuat function dalam satu baris atau mengelompokkon sintax-sintax di dalam function. 

#### Anonymous function Declaration
```golang
func main() {
  func() {
    fmt.Println("hello world!")
  }()
```

#### Closure function 
adalah function anonim yang mereferensikan variabel yang dideklarasikan di luar fungsi itu sendiri. Function ini tidak menggunakan variable yang diteruskan melalui parameter tetapi variable akan tersedia ketika dideklarasikan

#### Closure function Declaration
```golang
func newCounter() func() int {
  count := 0
  return func() int {
    count += 1
    return count
  }
}
```



# Pointer

Pointer adalah variable yang menyimpan memory address dari variable lain. Ketika data yang direferensikan berubah maka data di pointer akan ikut berubah

#### Pointer Declaration
```golang
  //normal variable
  var int a = 10  
  
  //pointer variable
  var p *int = &a
```

# Struct
Struct adalah user-defined type yang berisi kumpulan nama field/properties atau function

#### Struct Declaration
```golang
type Karyawa struct {
  Nama  string
  Umur  int
  Gaji  float64
}
```
