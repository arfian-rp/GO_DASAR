package main

import "fmt"

type Address struct {
	alamat string
}

func main() {
	//secara default golang semua variable di passing by value, bukan by reference
	//pass by value
	add1 := Address{"jl blm jadi"}
	add2 := add1
	add1.alamat = "jl sdh jadi"
	fmt.Println(add1)
	fmt.Println(add2) //berbeda

	//kita bisa membuat pass by reference dengan pointer
	add3 := Address{"jl. blmjd"}
	add4 := &add3 //add4 reference ke add3
	// var add4 *Address  = &add3
	add3.alamat = "jl sdhjd"
	fmt.Println(add3)
	fmt.Println(add4)

	// add4 = &Address{"jl. raya"} //kalau mencoba asign ulang maka add3 tidak berubah
	*add4 = Address{"jl. raya"} //kecuali kalau pake operator * (memaksa semua yg mengacu ke add3 diganti Address{"jl. raya"} )
	fmt.Println(add3)           // berubah
	fmt.Println(add4)

	//function new
	add5 := new(Address) //dibuatkan pointer ke Address data kosong
	add6 := add5
	add6.alamat = "Indonesia"
	fmt.Println(add5) //berubah
	fmt.Println(add6)
}
