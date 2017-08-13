
# 02 Files

## The `os` package
[https://golang.org/pkg/os/](https://golang.org/pkg/os/)
```go
var Args []string     // Commandline arguments

func Create(name string) (*File, error) // create a file
func Open(name string) (*File, error)   // open a file
func Getenv(key string) string          // get environment variable
func Exit(code int)                     // exit with return code
```

### Writing to a file
------------------------------
* Usage of the ```os``` package
* ```defer``` statements are called at the end of a function (similar to javas ```finally{}``` just without the ```try```)

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

### exec example
```go
import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: cmd command [args]\n")
	}
	c := exec.Command(os.Args[1], os.Args[2:]...)

	if out, err := c.Output(); err != nil {
		fmt.Printf("error: %v", err)
	} else {
		fmt.Printf("> %v", string(out))
	}
}
```

## Übung: cat Befehl
Implement base functionality of the `cat` command from linux

## Übung: wc Befehl
Implement base functionality of the `wc` command from linux

## Übung: tac Befehl
Implement base functionality of the `tac` command from linux

__Hint:__ Look at `bufio.Scanner` for file read operations

