package main

import (
	"fmt"
	"sort"
)

type Siswa struct {
	Name string
	Age  int
}

type SiswaSlice []Siswa

func (siswa SiswaSlice) Len() int {
	return len(siswa)
}

func (siswa SiswaSlice) Less(i, j int) bool {
	return siswa[i].Age < siswa[j].Age
}

func (siswa SiswaSlice) Swap(i, j int) {
	siswa[i], siswa[j] = siswa[j], siswa[i]
}

func main() {
	siswas := []Siswa{
		{"Adi", 8},
		{"Ari", 10},
		{"Edo", 9},
	}
	sort.Sort(SiswaSlice(siswas))
	fmt.Println(siswas)
}
