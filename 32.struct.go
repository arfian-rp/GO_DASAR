package main

import "fmt"

type Customer struct { //mirip class kalau di oop
	Name, Address string //bisa digabung kalau sama tipedatanya
	Age           int
}

func main() {
	/*
		struct => template data yg digunakan menggabungkan nol atau lebih tipedata lainnya dalam satu kesatuan.
		struct => kumpulan dari field
	*/
	var unDre Customer
	unDre.Name = "undre"
	unDre.Address = "english"
	unDre.Age = 22
	fmt.Println(unDre)

	//struct literals
	rudi := Customer{ //recomended
		Name:    "Rudi",
		Address: "Indonesia",
		Age:     44,
	}
	fmt.Println(rudi)

	budi := Customer{"Budi",
		"Indonesia", 66}
	fmt.Println(budi)
}
