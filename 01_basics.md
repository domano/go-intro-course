
# 01 Basics

## Hello World
```go
// file: hello.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, 世界")
}
```

Single file execution:
```go
go run hello.go
```

or using the [Golang Playground](https://play.golang.org/)

## Data types & variables
* Dynamic Memory Management
* Strongly typed
* All variables have a save initial value
Declaration and allocation:
```go
var i int
i = 42
```

Declaration with inference:
```go
i := 42
```

### built-in data types
* int, uint and float types in varying sizes
* bool
* complex, rune
* fixed length arrays `[4]int{42, 43, 44, 45}`
* slices (abstraction of interfaces): `[]int{42, 43, 44, 45}`
* string (as byte slice)
* maps `map[string]int{"a": 42, "b": 43}`
* structs

type conversion:
```go
a := 42
b := uint64(a)
```

### maps
```go
m := make(map[KeyType]ValueType)
```

Example:
```go
person := map[string]string{
	"name":    "Doe",
	"given": "John",
}
fmt.Println(person)
fmt.Println(len(person))

person["given"] = "Mike"
fmt.Println(person)
delete(person, "given")

_, exist := person["given"]
fmt.Println(exist) // false
```

### Arrays
Arrays have a fixed size. They are pretty inflexible and are rarely used directly
```go
colors := [5]string{"black", "red", "blue", "green", "white"}
fmt.Println(len(colors))
fmt.Println(colors[0])
```

### Slices
slices are a flexible list type. Internally data is put into an array,
so that slices actually are a reference to an array with a starting and an end index.

```go
colors := []string{"black", "red", "blue"}
colors = append(colors, "green")
colors = colors[1 : len(colors)-1]

fmt.Println(len(colors)) // 2
fmt.Println(cap(colors)) // 5
fmt.Println(colors)      // [red blue]
```

More verbose:
```go
slicename := make(type, len, cap)

// example:
colors := make([]string, 0, 5)
```

## Strings
strings are byte slices:
```go
	b := []byte{72, 97, 108, 108, 111, 32, 87, 101, 108, 116}
	fmt.Println(string(b))
```

Accessing sub strings:
```go
	fmt.Println(s[0:5])
```

The package [strings](https://golang.org/pkg/strings/) contains handy functions:
```go
    fmt.Println(strings.HasPrefix(s, "Hallo"))
	fmt.Println(strings.ToLower(s))
```


The package [strconv](https://golang.org/pkg/strconv/) contains functions for string conversion:
```go
	var i int
	i, err := strconv.Atoi("42")
	if err != nil {
		panic("not an integer")
	}
	fmt.Println(i)
```
    
## Call by Value

By default allocations in go happen _by value_ (instead of _by reference_).
__Beware__: maps and slices (including strings) themselves are references to an underlying data structure. Therefore they are called _by reference_.

## Text output
```go
import (
	"fmt"
	"os"
)

func main() {
	var i = 42

	// direct output
	fmt.Print(i)
	fmt.Print("\n")

	// output with newline
	fmt.Println(i)

	// format values
	fmt.Printf("the answer is %v\n", i)

	// formating as string
	s := fmt.Sprintf("the answer is %v\n", i)
	fmt.Print(s)

	// write to writer, e.g. stderr
	fmt.Fprintf(os.Stderr, "the answer is %v\n", i)
}
```

## Functions
Normal functions:
```go
func name(parameter1 type, parameter2 type) (returnParam1 type, returnParam2 type) {
 ..
}
```

Functions can be assigned to a variable:
```go
var hello = func(name string) {
	fmt.Println("Hello " + name)
}

var executer = func(name string, f func(name string)) {
	f(name)
}

executer("Marvin", hello)
```

## Control structures

### if {}
```go
if 2 > 1 {
	fmt.Println("1>2")
}

if data, err := readFromDatabase(); err != nil {
	fmt.Println("error reading data")
} else {
	fmt.Println(data)
}
```

### for {}
```go
colors := []string{"black", "red", "blue", "green", "white"}

// clasic for
for i := 0; i < len(colors); i++ {
	fmt.Printf("%v: %v\n", i, colors[i])
}

// iterate
for i, color := range colors {
	fmt.Printf("%v: %v\n", i, color)
}

// while
i := 0
for i < len(colors) {
	fmt.Printf("%v: %v\n", i, colors[i])
	i++
}

// while true
j := 0
for {
	if j >= len(colors) {
		break
	}
	fmt.Printf("%v: %v\n", j, colors[j])
	j++
}
```
    
### switch {}
```go
color := "nothing"
switch color {
case "green":
	fmt.Printf("Green")
case "red":
	fmt.Printf("Red")
default:
	fmt.Printf("Black")
}
```

## Exercise: multiplication table

__Input__: n, max number 

__Output__: formatted table multiplicating all smaller numbers (>0)

__Example__: go run mult.go 4

```
  1   2   3   4
  2   4   6   8
  3   6   9  12
  4   8  12  16
```


__Hint__: `os.Args` contains the call parameters 


