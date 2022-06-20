package main

import "fmt"

func main() {
	/*
		slice => potongan data array.
		slice mirip dengan array, yg membedakan adalah ukuran slice bisa berubah.
		slice dan array selalu terkoneksi, dimana slice adalah data yg mengakses sebagian atau seluruh data di array.

		detail:
		slice memiliki 3 data, yaitu pointer, length dan capacity.
		pointer => petunjuk data pertama di array para slice.
		length => panjang slice.main()
		capacity => kapasitas slice, length tidak boleh lebih dari capacity.

		membuat slice dari array:
		array[low:high] => low - sebelum high
		array[low:] => low - akhir
		array[:high] => 0 - sebelum high
		array[:] => 0 - akhir
	*/

	month := [...]string{
		"jan",
		"feb",
		"mar",
		"apr",
		"mei",
		"jun",
		"jul",
		"agu",
		"sep",
		"okt",
		"nov",
		"des",
	}

	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// saat merubah array/slice, slice/array ikut berubah
	month[5] = "dyuba"
	fmt.Println(slice1)
	slice1[1] = "juney"
	fmt.Println(month)

	//append(). kalau capacity penuh maka akan membuat array baru
	slice2 := month[10:]
	fmt.Println(slice2)
	slice3 := append(slice2, "Tyson")
	fmt.Println(slice3)
	slice3[1] = "!des"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(month)

	//make([]typeData,length,capacity)
	newSlice := make([]string, 2, 5)
	newSlice[0] = "sapi"
	newSlice[1] = "kamving"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy(to, from)
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
	newSlice[0] = "Mebk" //ga ngaruh ke copySlide
	fmt.Println(copySlice)

	//array vs slice
	iniArray := [...]int{1, 2, 3, 4}
	iniSlice := []int{1, 2, 3, 4}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
