package main

import "fmt"

func main() {
	//constant => variable yg tidak bisa diubah (immutable)
	const url = "https://arfian-id.web.app"
	//multiple const
	const (
		firstName = "rudie"
		lastname  = "hermavan"
	)
	fmt.Println("url = ", url)
	fmt.Println(firstName, " ", lastname)
}
