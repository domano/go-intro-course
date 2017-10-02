package main

import (
	"fmt"
)

func main() {


	myFunction("a", "b", "c")
	
}

func myFunction(param ...string) {
	fmt.Println(len(param))
}
