package main

func main() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break //menghentikan perulangan
		}
		if i == 3 {
			continue //lanjut perulangan selanjutnya
		}
		println(i)

	}
}
