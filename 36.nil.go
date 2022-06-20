package main

import "fmt"

func main() {
	/*
		biasanya kalau di bahasa pemrograman lain object yg belum diinisialisasikan nilainya null/nill.
		tetapi golang otomatis akan dibuatkan default valuenya.
	*/
	var person map[string]string = nil
	fmt.Println(person)

	//fungsi nil adalah unduk validasi data misal:
	hewans := NewMap("")
	if hewans == nil { //daripada mengecek apakah terdapat value a didalam key b
		fmt.Println("data kosong")
	} else {
		fmt.Println(hewans)
	}
}

func NewMap(name string) map[string]string {
	if name != "" {
		return map[string]string{"name": name}
	}
	return nil
}
