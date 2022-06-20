package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Trim("    hello    ", " "))
	fmt.Println(strings.ToLower("HELLO"))
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.Contains("hello-world", "-"))
	fmt.Println(strings.Split("hello-world", "-"))
	fmt.Println(strings.ReplaceAll("hello-world", "-", "_"))
	fmt.Println(strings.ToTitle("hello world"))
}
