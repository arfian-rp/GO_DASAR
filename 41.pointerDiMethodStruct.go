package main

import "fmt"

type Man struct {
	name string
}

func (man Man) changeName1() {
	man.name = "Teddy"
}

func (man *Man) changeName2() {
	man.name = "Teddy"
}

func main() {
	gamal := Man{name: "Gamal"}
	gamal.changeName1() //ga ngaruh karena pass by value (data di duplikat)
	fmt.Println(gamal)

	gamal.changeName2() //pass by reference
	fmt.Println(gamal)  //berubah
}
