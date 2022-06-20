package main

import "fmt"

func main() {
	//map tidak ada nilai max, tetapi index(key) tidak boleh sama, kalau key sama artinya mengubah
	person := map[string]string{
		"name":    "rudy",
		"address": "cina",
	}
	person["title"] = "human" //menambah/mengubah

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(len(person))
	delete(person, "title")
	fmt.Println(person)

	//atau
	hewan := make(map[string]string)
	hewan["satu"] = "harimau"
	fmt.Println(hewan)
}
