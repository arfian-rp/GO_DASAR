package main

func main() {
	//closure => kemampuan function berinteraksi dengan data-data disekitarnya
	counter := 0
	inc := func() {
		println("inc")
		counter++
	}
	inc()
	inc()
	inc()
	println(counter)
}
