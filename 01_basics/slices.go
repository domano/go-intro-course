package main

import "fmt"

func main() {

	colors := []string{"black", "red", "blue"}
	colors = append(colors, "green")
	colors = colors[1 : len(colors)-1]

	fmt.Println(len(colors))
	fmt.Println(cap(colors))
	fmt.Println(colors)
}
