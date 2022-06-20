package main

func main() {
	name := "yudass"

	if name == "yudi" {
		println("Hello Yudi")
	} else if name == "yuda" {
		println("Hello Yuda")
	} else {
		println("Hello !Yudi&&!Yuda")
	}

	//if dengan short statement
	if length := len(name); length > 5 {
		println("nama terlalu panjang")
	}
}
