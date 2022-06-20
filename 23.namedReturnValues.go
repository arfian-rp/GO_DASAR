package main

func main() {
	f, m, l := getCompleteName()
	println(f, m, l)
}

func getCompleteName() (first, middle, last string) { //kalau semua tipe samam bisa didefiniskan di-akhir
	first = "Alexandre"
	middle = "Rudolph"
	last = "Pierre"
	return //ga wajib return satu-satu
}
