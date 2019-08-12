package main

import "fmt"

type test map[string]string

func main() {

	//creating a map
	// var colors map[string] string
	/*colors := map[string]string{
		"red":  "#FF0000",
		"blue": "#0000ff",
	}*/

	//creates an empty map of type string, string
	colors := make(map[string]string)
	//to add values to map
	colors["white"] = "#ffffff"

	//to access the map
	fmt.Println(colors["white"])

	//remove a key --> using an internal func
	delete(colors, "white")

	//to print the map
	fmt.Println("After deleting white key:", colors)

	//creating a new map
	testMap := map[string]string{
		"red":   "#FF0000",
		"blue":  "#0000ff",
		"white": "#fffff",
	}
	printMap(testMap)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex Code for:", color, "is", hex)
	}
}
