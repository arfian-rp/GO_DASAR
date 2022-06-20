package main

import (
	"container/list"
	"fmt"
)

/*
container/list => implementasi structur data double linked list di golang.
linked list tidak memiliki index, jadi hanya bisa diiterasi
*/

func main() {
	data := list.New()
	data.PushBack("teddy")
	data.PushBack("yudi")
	data.PushBack("jaka")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
