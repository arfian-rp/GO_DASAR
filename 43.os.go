package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("error", err.Error())
	}
	fmt.Println(hostname)
}
