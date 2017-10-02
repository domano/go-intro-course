
# Testing go code

## Tests
* All Files ending in `_test.go` contains testing code
* `go test <package>`
* Test are functions with names matching the following pattern: `func Test_*(t *testing.T)`

```go
package foo

import "testing"

func Test_Simple(t *testing.T) {

	t.Logf("This Test fails")
	t.Fail()

}
```

## stretchr/testify
* Simple library offering assertions
* Install test dependencies with `go get -t`

```go
package foo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_With_Testify(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, 1)
}
```

## Test multiple packages
Multiple packages can be tested at once: `go test package/...`

Example:
```shell
go test -cover github.com/tarent/loginsrv/...
ok  	github.com/tarent/loginsrv	1.012s	coverage: 69.0% of statements
ok  	github.com/tarent/loginsrv/caddy	0.004s	coverage: 89.1% of statements
?   	github.com/tarent/loginsrv/caddy/demo	[no test files]
ok  	github.com/tarent/loginsrv/htpasswd	2.022s	coverage: 98.8% of statements
ok  	github.com/tarent/loginsrv/httpupstream	0.014s	coverage: 97.6% of statements
ok  	github.com/tarent/loginsrv/logging	0.005s	coverage: 94.0% of statements
ok  	github.com/tarent/loginsrv/login	0.009s	coverage: 95.0% of statements
ok  	github.com/tarent/loginsrv/model	0.012s	coverage: 100.0% of statements
ok  	github.com/tarent/loginsrv/oauth2	0.019s	coverage: 92.5% of statements
ok  	github.com/tarent/loginsrv/osiam	0.005s	coverage: 89.9% of statements

```

## check coverage
Show coverage with `go tool cover`

```shell
go test -cover -coverprofile cover.out  github.com/tarent/loginsrv/login
go tool cover -html=cover.out
```

## Data Driven Tests
```
func Test_Operations(t *testing.T) {
	tests := []struct {
		title    string
		op       func(int, int) int
		a        int
		b        int
		expected int
	}{
		{"add", add, 0, 0, 0},
		{"add", add, 0, 1, 1},
		{"add", add, 1, 2, 3},
		{"add", add, 2, 2, 4},
		{"add", add, 42, 42, 42}, // this one fails
		{"sub", sub, 0, 0, 0},
		{"sub", sub, 0, 1, -1},
		{"sub", sub, 42, 42, 42}, // this one fails
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v(%v, %v) == %v", test.title, test.a, test.b, test.expected)
		t.Run(testName, func(t *testing.T) {
			actual := test.op(test.a, test.b)
			if actual != test.expected {
				t.Logf("expected %v, but got %v", test.expected, actual)
				t.Fail()
			}
		})
	}
}
```

## Mocking

Mocking can not be dynamic. Mock-Frameworks generate mocks at development time and they are usually checked in with the rest of your code. [GoMock](https://github.com/golang/mock) is a good example for a mock framework.

```
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := NewMockMyInterface(ctrl)

	mock.EXPECT(). ...
```
    
Code generation in go:
```
//go:generate go get github.com/golang/mock/mockgen
//go:generate $GOPATH/bin/mockgen -self_package objects -package objects -destination $GOPATH/src/objects/mocks_test.go objects Flyable
func ...
```

```
$ go generate .
```

## Benchmarks

### Benchmarks: Basic idea
* Benchmarks are located in the test files, just like normal test functions
* Their signature matches the following pattern: `func Benchmark_*(b *testing.B)`
* The benchmark code gets executed in a loop `b.N` times.
* Depending on execution time go changes `b.N` (e.g. 100, 10000, 1000000).

### Benchmarks: Helper functions
To avoid benchmark code from being part of the benchmark itself, benchmark timers can be manually controlled:

* `b.ResetTimer()` - Reset the timer
* `b.StartTimer()` - (Re)Start the timer
* `b.StopTimer()` - Stop the timer

### Benchmarks: Example

```go
func Benchmark_Creation_Of_Goroutines(b *testing.B) {
	fmt.Printf("testing with %v goroutines\n", b.N)
	doneChannel := make(chan bool)

	for i := 0; i < b.N; i++ {
		go doInBackground(doneChannel)
	}

	for i := 0; i < b.N; i++ {
		<-doneChannel
	}
}
```

### Benchmark execution

Benchmarks can be executed with `go test -bench <regex>`

Example:
```shell
go test -bench '.*' goroutine_lots_of_test.go
```

## Exercise: console calculator

* Write a simple calculator
* Test it! Coverage should be 90% or higher

## Exercise: Benchmarking of the Key-Value store
* Measure the time needed by the key-value store for a _write->save->read_ cycle.
