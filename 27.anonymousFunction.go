package main

import "strings"

func main() {
	filter := func(s string) string {
		return strings.ToUpper(s)
	}

	sayHelloWithFilter2("yuda", filter)
}

func sayHelloWithFilter2(name string, filter func(string) string) {
	println("Hello,", filter(name))
}
