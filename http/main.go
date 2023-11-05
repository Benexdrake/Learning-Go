package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1/")

	if err != nil {

		fmt.Println("Error")
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	sb := string(body)
	fmt.Println(sb)

	fmt.Println()

}
