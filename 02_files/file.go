package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("/tmp/hello")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()

	file.WriteString("Hello irbid\n")
}
