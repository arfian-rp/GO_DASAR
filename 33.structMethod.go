package main

//struct method => seakan-akan membuat method di struct

type Orang struct {
	Name string
}

func (me Orang) sayHelloTo(teman Orang) {
	println("Hello", teman.Name, "my name is", me.Name)
}

func main() {
	rudy := Orang{Name: "Rudy"}
	julian := Orang{Name: "Julian"}
	rudy.sayHelloTo(julian) //seakan-akan memiliki method
}
