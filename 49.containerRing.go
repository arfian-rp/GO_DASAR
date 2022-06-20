package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

/*
container/ring => implementasi structur data circular list.
circular list => data ring. dimana diakhir element akan kembali ke element awal
*/

func main() {
	data := ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "value " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
