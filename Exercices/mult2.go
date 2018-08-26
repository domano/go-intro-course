package main

import "fmt"
import "strconv"
import "os"

func main() {
  max, err := strconv.Atoi(os.Args[1])
  if err != nil {
		panic("Not an integer!")
	}
  for i := 1; i <= max; i++ {
    for j := 1; j <= max; j++ {
      fmt.Printf("%d\t",i*j)
    }
    fmt.Println("\n")
  }
}
