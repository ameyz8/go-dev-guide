package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	response, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// byteSlice := make([]byte, 100000)
	// response.Body.Read(byteSlice)
	// fmt.Println(string(byteSlice))

	io.Copy(os.Stdout, response.Body)
}
