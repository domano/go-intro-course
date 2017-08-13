package main

import "fmt"

func main() {

	colors := [5]string{"black", "red", "blue", "green", "white"}

	bunt := colors[1 : len(colors)-1]
	fmt.Println(bunt)

	bunt = append(bunt, "orange")
	fmt.Println(bunt)
	fmt.Println(colors)
}
