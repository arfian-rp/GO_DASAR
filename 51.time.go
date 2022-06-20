package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Local())
	fmt.Println(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	fmt.Println(time.Parse(time.RFC3339, "2009-01-02T15:22:07Z"))
}
