package main

import "strings"

type Filter func(string) string

func main() {
	sayHelloWithFilter("yuda", ToUpper)
}

func sayHelloWithFilter(name string, filter Filter) {
	println("Hello,", filter(name))
}

func ToUpper(name string) string {
	return strings.ToUpper(name)
}
