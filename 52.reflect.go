package main

import (
	"fmt"
	"reflect"
)

/*
reflect => fitur dimana kita bisa melihat structur kode saat aplikasi sedang berjalan.
*/

type Sample struct {
	Name string `required:"true" max:"10"` //struct tag
	Age  int
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

func main() {
	sample := Sample{Name: "Adi", Age: 22}
	fmt.Println(reflect.TypeOf(sample))
	fmt.Println(reflect.TypeOf(sample).Field(0))
	fmt.Println(reflect.TypeOf(sample).Field(0).Tag.Get("required"))
	fmt.Println(reflect.TypeOf(sample).Field(1))

	fmt.Println(IsValid(Sample{Age: 12}))
	fmt.Println(IsValid(Sample{Age: 12, Name: "yanto"}))
}
