package main

import "fmt"

func main() {
	//var colors map[string]string

	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// delete(colors, "white")

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#008000",
		"blue":  "#0000FF",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
