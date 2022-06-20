package main

import "flag"

//go run 44.flag.go -host=localhost -user=root -pass=root

func main() {
	host := flag.String("host", "defaultHost", "Put your host") //return pointer
	user := flag.String("user", "defaultUser", "Put your user")
	pass := flag.String("pass", "defaultPass", "Put your pass")

	flag.Parse() //parse

	println("host:", *host)
	println("user:", *user)
	println("pass:", *pass)
}
