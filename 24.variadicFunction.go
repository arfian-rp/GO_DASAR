package main

func main() {
	//varargs => menerima data lebih dari satu input
	//jika param lebih dari 1, maka varargs ditaruh di akhir parameter
	println(sumAll(1, 2, 3, 4, 5, 6, 6, 7, 8))
	//atau
	slice := []int{1, 2, 4, 5, 6}
	println(sumAll(slice...)) //kalau terlanjur pakai slice
}

func sumAll(numbers ...int) int { //number adalah varargs
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
