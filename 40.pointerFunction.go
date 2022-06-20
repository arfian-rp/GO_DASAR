package main

import "fmt"

type Siswa struct {
	name string
	age  int
}

func changeNameSiswaToJaka1(siswa Siswa) {
	siswa.name = "Jaka"
}

func changeNameSiswaToJaka2(siswa *Siswa) {
	siswa.name = "Jaka"
}

func main() {
	adi := Siswa{name: "Adi", age: 12}
	changeNameSiswaToJaka1(adi) //ga ngaruh karena pass by value (data di duplikat)
	fmt.Println(adi)
	changeNameSiswaToJaka2(&adi) //pass by reference
	fmt.Println(adi)             //berubah
}
