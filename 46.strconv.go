package main

import "strconv"

func main() {
	valBool, _ := strconv.ParseBool("true")
	println(valBool)
	strBool := strconv.FormatBool(true)
	println(strBool)

	//Atoi & Itoa
	println(strconv.Itoa(12))
	println(strconv.Atoi("33"))
}
