package main

import "fmt"

func main() {
	var i32 int32 = 4123
	var i64 int64 = int64(i32)
	var i16 int16 = int16(i32)

	fmt.Println(i64)
	fmt.Println(i16)
	fmt.Println(string("rudei"[0])) //byte ke string
}
