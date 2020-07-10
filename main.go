package main

import "fmt"

func main() {
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// fmt.Println(colors)
	// delete(colors, "white")
	// fmt.Println(colors)

	// var colors map[string]string

	colors := map[string]string{
		"red":    "#ff0000",
		"green":  "#008000",
		"blue":   "#0000FF",
		"yellow": "#FFFF00",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("The hex code for", color, "is", hex)
	}
}
