package main

func main() {
	hasil := tambah(1, 4)
	println(hasil)
	println(tambah(4, 4))
}

func tambah(val1 int, val2 int) int {
	return val1 + val2
}
