package main

func main() {
	//type assertions => kemampuan merubah tipedata menjad tipedata yg diinginkan.

	result := random()
	resultString := result.(string) //kalau yakin adalah string
	println(resultString)

	// resultInt := result.(int) //panic
	// println(resultInt)

	//type assertions dibantu switch
	switch value := result.(type) {
	case string:
		println("String", value)
		//code program apabila string
	case int:
		println("Int", value)
		//code program apabila int
		//dan strusnya
	}
}

func random() interface{} {
	return "STRING"
}
