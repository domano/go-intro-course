# 04 Projektstruktur

## Gopath
The GOPATH is the base path go uses to find packages.
Since go 1.8 the default is set to `~/go`

### Create a workspace
```shell
mkdir hello_world
cd hello_world
export GOPATH=`pwd`
mkdir -p src/hello
```

### Handy alias
``
alias gopath="export GOPATH=\`pwd\`; export PATH=\`pwd\`/bin:$PATH"
``

### Compile and run
```shell
go install hello
bin/hello
```

## Go Toolchain
```shell
go build <package>    # Just build in the current directory
go test <packages>    # Execute tests (File pattern:*_test.go) for package
go get <packages>     # Get and install all dependencies
go get -t <package>   # Get and install all dependencies incl. the ones needed for the tests
go vet <package>      # Statical code analysis
go fmt <package>      # Uniform code format
go generate <package> # Code generation

```

### go get
Go's dependency concept

* source code dependencies
* No binary library concept (since 1.8 .so files can be imported though)
* Automated download and build via `go get package/name`

Example:

    export GOPATH=`pwd`
    go get github.com/sDoe/servelocal

Improts are references to code repositories:

    import "github.com/gorilla/handlers"

## Packages
* Every go file starts with a package declaration
* Every directory contains just one package
* The `main` package is used as the entry point
* Import of packages
```go
import(
    "fmt"
    "math"
)
```
* Packages can be fully qualified, z.B.  `import github.com/tarent/loginsrv`
* Upped case identifiers get exported and can be accessed outside of packages (~ public).

### init()

The `init()`-function of a package gets called when the package gets imported (before the call to ```main()```), even if the imported package is not used.
```
package foo

func init() {
    // do some initial stuff here
}
```

```
package main

import (
    _ foo
)
```
