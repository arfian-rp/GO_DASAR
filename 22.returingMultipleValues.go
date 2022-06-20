package main

func main() {
	_, lastName := getFullName()  //_ artinya ignore
	firstName, _ := getFullName() //_ artinya ignore

	println(firstName, lastName)
}

func getFullName() (string, string) {
	return "Mike", "Tyson"
}
