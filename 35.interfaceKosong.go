package main

/*
jika di java ada class Object, di kotlin ada class Any.
maka di golang kita bisa membuat interface kosong untuk semua tipe menjadi implementasinya
*/

func main() {
	cetak(1)
	cetak("ini string")
	cetak('a')
	//dst

	apaAja := getSomething()
	println(apaAja)
}

func cetak(content interface{}) { //sebagai parameter
	println(content)
}

func getSomething() interface{} { //sebagai return value
	return "string"
}
