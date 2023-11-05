package main

import "fmt"

type TEST struct {
	id   string
	name string
}

type ID struct {
	id string
}

func main() {

	test := map[ID]TEST{
		{"test"}: {id: "1234", name: "Nope"},
	}

	fmt.Println(test)

	colors := map[string]string{
		"red":    "#ff0000",
		"blue":   "#00ff00",
		"yellow": "#0000ff",
		"white":  "ffffff",
		"black":  "000000",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["white"] = "#fff"
	// colors["black"] = "#000"

	//delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
