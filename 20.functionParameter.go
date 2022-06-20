package main

import "fmt"

func main() {
	sayHelloTo("Budee", "Santoso")
}

func sayHelloTo(fname string, lname string) {
	fmt.Println("Hello", fname, lname)
}
