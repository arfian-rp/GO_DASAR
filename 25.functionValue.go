package main

func main() {
	goodbye := getGoodbye
	println(goodbye("Tyson"))
}

func getGoodbye(name string) string {
	return "Goodbye " + name
}
