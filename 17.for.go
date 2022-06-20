package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		println("Print ke ", i)
		i++
	}

	i = 1

	angka := [5]int{}
	for i <= 5 {
		angka[i-1] = i
		i++
	}
	fmt.Println(angka)

	i = 0

	//for dengan statement
	//for init statement(pertama sebelum for dieksekusi); kondisi; post statement(dieksekusi tiap akhir perulangan)
	for batas := len(angka) - 1; i <= batas; i++ {
		println(angka[i])
	}

	//for range
	iniSlice := []int{4, 6, 3, 1, 3}
	for indexOrKeyInMap, value := range iniSlice {
		println(indexOrKeyInMap, "->", value)
	}

}
