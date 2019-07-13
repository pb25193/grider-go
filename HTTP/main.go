package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// fmt.Println(resp.Body)

	lw := logWriter{}

	io.Copy(&lw, resp.Body)
}

func (lw *logWriter) Write(bs []byte) (int, error) {
	fmt.Println("Just wrote this many bytes:", len(bs))
	return 1, nil
}
