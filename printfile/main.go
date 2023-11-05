package main

import (
	"fmt"
	"os"
)

func main() {

	c, err := os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	str := string(c)

	fmt.Println(str)

}
