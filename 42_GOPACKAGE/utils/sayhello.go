package utils

import "strings"

/*
Access Modifier di golang
jika diawali huruf besar maka bisa diakses dari luar.
jila diawali huruf kecil maka tidak bisa diakses dari luar.
*/

func SayHello(name string) { //bisa diakses dari luar
	println("Hello,", filter(name))
}

func filter(name string) string { //tidak bisa diakses dari luar
	return strings.ToUpper(name)
}
