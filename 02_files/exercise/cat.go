package main

import (
	"fmt"
	"os"
)

// go run cat.go /some/file
// Output: file contents
func main() {
	args := os.Args
	file, error := os.Open(args[1])
	if error != nil {
		fmt.Print("error happend!")
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	b := make([]byte, int(stat.Size()))
	_, err = file.Read(b)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(string(b))
}
