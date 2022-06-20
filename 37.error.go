package main

import "errors"

func main() {
	hasil, err := Pembagian(33, 0)
	if err == nil {
		println("Hasil", hasil)
	} else {
		println("Error", err.Error())
	}
}

func Pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("tidak bisa dibagi dengan 0")
	} else {
		return nilai / pembagi, nil
	}
}
