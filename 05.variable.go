package main

import "fmt"

func main() {
	var name string = "rudy"
	//bisa jg
	age := 12 // := hanya untuk deklarasi awal
	var alamat = "jl. blmjadi"
	//multiple var
	var (
		firstName = "rudie"
		lastname  = "hermavan"
	)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(alamat)
	fmt.Println(firstName, " ", lastname)
}
