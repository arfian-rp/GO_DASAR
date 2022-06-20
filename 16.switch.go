package main

func main() {
	name := "Budee"

	switch name {
	case "Budee":
		println("Hello Budee")
	case "Benee":
		println("Hello Benee")
	default:
		println("Hello !Budee&&!Benee")
	}

	//switch dengan short statement
	switch length := len(name); length > 5 {
	case true:
		println("nama terlalu panjang")
	default:
	}

	//switch tanpa kondisi (mirip if)
	switch {
	case len(name) > 5:
		println("nama terlalu panjang")
	default:
		println("nama terlalu pendek")
	}
}
