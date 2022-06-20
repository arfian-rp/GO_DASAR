package main

type Person struct {
	Name string
}

type HasName interface {
	GetName() string
}

func (p Person) GetName() string { //karena memiliki func GetName otomatis Person mengikuti kontrak HasName
	return p.Name
}

func sayHello(hasName HasName) {
	println("hello", hasName.GetName())
}

func main() {
	rudy := Person{"Rudy"}
	sayHello(rudy)
}
