package main

import "fmt"

func main() {
	colors := map[string]string{
		"Red":   "#ff0000",
		"Green": "#00ff00",
		"Blue":  "#0000ff",
		"White": "#ffffff",
		"Black": "#000000",
	}

	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
