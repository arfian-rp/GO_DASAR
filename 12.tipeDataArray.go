package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "herruw"
	names[1] = "magweye"

	fmt.Println(names)
	println(names[0])
	println(names[1])
	// println(names[2]) error

	// atau
	// var values = [3]int{
	// 	1, 2, 3,
	// }
	values := [3]int{
		1, 2, 3,
	}
	println(values[2])
	println(len(values)) //get panjang array

	// gunakan [...] apabila tidaktau jumlahnya
	months := [...]string{
		"jan",
	}
	println(months[0])
}
