# Time & Space Complexity

Dengan menggunakan konsep Time Complexity kita dapat mempermudah estimasi waktu yang digunakan suatu program beroperasi

Time Complexity biasanya di deskripsikan menggunakan Big-O notation. Terdapat berbagai macam Big-O notation yang mendeskripsikan estimasi waktu suatu program beroperasi diantaranya :

## O(1) - Constant Time
  Pada notation ini program akan berjalan hanya 1 kali
```golang
func dominan(n int) int {
	var hasil int = 17
	hasil += n
	return hasil
}

func main() {
	fmt.Println(dominan(10))
}
```
#### Output
```
27
```

## O(n) - Linear Time
  Program akan berjalan sesuai jumlah nilai *n* yang dimasukan. Jika *n* diinput dengan nilai 10  => O(n) = O(10) .maka program akan berjalan 10 kali.

```golang
func dominan(n int) int {
	var programBerjalan int
	for i := 0; i < n; i++ {
		programBerjalan += 1
		fmt.Println("Program Berjalan => ", i, " kali")
	}
	return programBerjalan
}

func main() {
	fmt.Println("Total Program Berjalan => ", dominan(10))
}
```
#### Output
```
Program Berjalan =>  0  kali
Program Berjalan =>  1  kali
Program Berjalan =>  2  kali
Program Berjalan =>  3  kali
Program Berjalan =>  4  kali
Program Berjalan =>  5  kali
Program Berjalan =>  6  kali
Program Berjalan =>  7  kali
Program Berjalan =>  8  kali
Program Berjalan =>  9  kali
Total Program Berjalan =>  10
```

## O(n + m) - Linear Time
  Program akan berjalan sesuai jumlah nilai *n* dan *m* yang dimasukan. Jika *n* dan *m* diinput dengan nilai 15 dan 20  => O(n + m) = O(15 + 20) .maka program akan berjalan 35 kali.

```golang
func dominan(n, m int) int {
	var programBerjalan int
	for i := 0; i < n; i++ {
		programBerjalan += 1
		fmt.Println("Program 1 Berjalan => ", i, " kali")
	}
  for j := 0; j < m; j++ {
		programBerjalan += 1
		fmt.Println("Program 2 Berjalan => ", j, " kali")
	}
	return programBerjalan
}

func main() {
	fmt.Println("Total Program Berjalan => ", dominan(15, 20))
}
```
#### Output
```
Program 1 Berjalan =>  0  kali
Program 1 Berjalan =>  1  kali
Program 1 Berjalan =>  2  kali
Program 1 Berjalan =>  3  kali
Program 1 Berjalan =>  4  kali
Program 1 Berjalan =>  5  kali
Program 1 Berjalan =>  6  kali
Program 1 Berjalan =>  7  kali
Program 1 Berjalan =>  8  kali
Program 1 Berjalan =>  9  kali
Program 1 Berjalan =>  10  kali
Program 1 Berjalan =>  11  kali
Program 1 Berjalan =>  12  kali
Program 1 Berjalan =>  13  kali
Program 1 Berjalan =>  14  kali

Program 2 Berjalan =>  0  kali
Program 2 Berjalan =>  1  kali
Program 2 Berjalan =>  2  kali
Program 2 Berjalan =>  3  kali
Program 2 Berjalan =>  4  kali
Program 2 Berjalan =>  5  kali
Program 2 Berjalan =>  6  kali
Program 2 Berjalan =>  7  kali
Program 2 Berjalan =>  8  kali
Program 2 Berjalan =>  9  kali
Program 2 Berjalan =>  10  kali
Program 2 Berjalan =>  11  kali
Program 2 Berjalan =>  12  kali
Program 2 Berjalan =>  13  kali
Program 2 Berjalan =>  14  kali
Program 2 Berjalan =>  15  kali
Program 2 Berjalan =>  16  kali
Program 2 Berjalan =>  17  kali
Program 2 Berjalan =>  18  kali
Program 2 Berjalan =>  19  kali
Total Program Berjalan =>  35
```

## O(log n) - Logarithmic Time
Pada notation ini nilai *n* berkurang setengahnya (dibagi 2) setiap perulangan (loop/iteration). Jika n = 2<sup>x</sup>, maka log n = x ~> log 32 = 5
```golang
func logaritma(n int) int {
	var programBerjalan int
	for n > 1 {
		n /= 2
		fmt.Println("Program Berjalan")
		programBerjalan += 1
	}
	return programBerjalan
}

func main() {
	fmt.Println("Total Program Berjalan = ", logaritma(35))
}
```
#### Output
```
Program Berjalan
Program Berjalan
Program Berjalan
Program Berjalan
Program Berjalan
Total Program Berjalan =  5
```

## O(n<sup>2</sup>) - Quadratic Time
Pada notation ini program berjalan sebanyak *n*<sup>2</sup>. jika nilai *n* = 5, maka nilai *n* = 25.
```golang
func quadratic(n int) int {
	var programBerjalan int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Println("Program Berjalan => ", programBerjalan, " kali")
			programBerjalan += 1
		}
	}
	return programBerjalan
}

func main() {
	fmt.Println("Total Program Berjalan = ", quadratic(5))
}
```
#### Output
```
Program Berjalan =>  0  kali
Program Berjalan =>  1  kali
Program Berjalan =>  2  kali
Program Berjalan =>  3  kali
Program Berjalan =>  4  kali
Program Berjalan =>  5  kali
Program Berjalan =>  6  kali
Program Berjalan =>  7  kali
Program Berjalan =>  8  kali
Program Berjalan =>  9  kali
Program Berjalan =>  10  kali
Program Berjalan =>  11  kali
Program Berjalan =>  12  kali
Program Berjalan =>  13  kali
Program Berjalan =>  14  kali
Program Berjalan =>  15  kali
Program Berjalan =>  16  kali
Program Berjalan =>  17  kali
Program Berjalan =>  18  kali
Program Berjalan =>  19  kali
Program Berjalan =>  20  kali
Program Berjalan =>  21  kali
Program Berjalan =>  22  kali
Program Berjalan =>  23  kali
Program Berjalan =>  24  kali
Total Program Berjalan =  25
```
