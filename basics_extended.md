# Weitere Grundlagen

## defer
------------------------------
* `defer`: Delays execution until the end of the current function
* Execution is guaranteed (== finally{})

```go
import "os"

func main() {
	file, err := os.Create("/tmp/hello")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()

	file.WriteString("Hello World\n")
}
```

## Error Handling in Go
* Error as a return value

```go
	if err != nil {
		return nil, err
	}
```

### Wrapping errors

* Enrich an error with contextual information
* Get a stack trace, allowing follow the call stack of an error 

```go
package main

import "fmt"
import "github.com/pkg/errors"

func main() {
        err := errors.New("error")
        err = errors.Wrap(err, "open failed")
        err = errors.Wrap(err, "read config failed")

        fmt.Println(err) // read config failed: open failed: error
}
```

### Specific errors

In some cases it can be useful to return specific error types (if you want to check which error occured for example).
Usually this happens using global variables of type `error`.

See https://golang.org/pkg/io/#pkg-variables

Alternatives:

https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully

https://www.youtube.com/watch?v=lsBF58Q-DnY

## Panic, Recover
* `panic`: Causes a panic (something like an exception)
* `recover`: Catches a panic

```go
func travel() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r, "..don't panic!")
		}
	}()
	panic("I lost my towel")
}
```

## Type conversion

In go types can be converted into compatible types. This happens explicitly to avoid invisible bugs.

```go
i := 42
f := float64(i)
u := uint(f)
```

This also can be done with custom types
```go
s := "string"
type myString string
var ms myString
ms = myString(s)
```


## Empty interface: `interface{}`

* Entspricht `Object` in Java oder `void *` in C.
* Variablen vom Typ `interface{}` können beliebige Werte aufnehmen


### Type assertion

Auf den konkreten Wert einer `interface{}` Variablen kann über eine Type Assertion zugegriffen werden.

```go
var o interface{} = 42
i := o.(int)
i++
```

Wenn eine Type Assertion fehl schlägt wir ein `panic()`ausgelöst.

Um ein Panic zum umgehen kann eine Type Assertion inkl. Test ausgeführt werden:

```go
i, ok := o.(int)
```

### Type switches
* `interface{}` entspricht dem generischen Typ (Object bzw. void*).
* Mit type switches lässt sich elegant in einem Schritt prüfen und konvertieren.

```go
untypedList := []interface{}{"Hallo", 42, false}

for _, item := range untypedList {
	switch i := item.(type) {
	case string:
		fmt.Println("String: " + i)
	case int:
		i++
		fmt.Println("Int: " + strconv.Itoa(i))
	default:
		fmt.Println(i)
	}
}
```
