package main

func main() {
	println(fatorial(5))
}

func fatorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * fatorial(value-1)
	}
}
