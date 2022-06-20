package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`a([a-z])i`)
	fmt.Println(regex.MatchString("adi"))
	fmt.Println(regex.MatchString("agi"))
	fmt.Println(regex.MatchString("ari"))

	fmt.Println(regex.FindAllString("adi awi ari agi a1o aji ahi qwd", -1)) // adalah max, -1 untuk semua(tidak ada max)
}
