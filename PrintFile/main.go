package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open((os.Args[1]))
	var message string
	if err == nil {
		bs := make([]byte, 22999)
		file.Read(bs)
		message = string(bs)
	}
	fmt.Println(message)
}
