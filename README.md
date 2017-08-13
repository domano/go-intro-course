# Golang Schulung

## Grundlagen

### [01 basics](01_basics.md)
- Hello World
- data types & variables
- strings
- Call by Value & by Reference
- Text output
- Functions
- Control structures
- Exercise: multiplication table

### [02 files](02_files.md)
- The `os` package
- Write into a file
- Exec Example
- Exercise: wc command
- Exercise: cat command
- Exercise: tac command

### [03 structs and pointers](03_structs_and_pointers.md)
- Structs
- Pointers
- Pointer and Structs
- Pointer and Slices/Maps
- Exercise: Key-Value store

### [04 Projectstrcture](04_projectstructure.md)
- Gopath
- Go Toolchain
- Packages

### [05 testing](05_testing.md)
- Tests
- stretchr/testify
- Testing multiple packages
- Coverage
- Data Driven Tests
- Benchmarks
- Exercise: Calculator with tests
- Exercise: Benchmarking the Key-Value store

### [06 objects and interfaces](06_objects_and_interfaces.md)
- Methods
- Objects
- Embedding
- Überschreiben
- Interfaces
- Mocking
- Übung: Key-Value Objekt Orientiert
- Übung: Testen mit Mocks

### [08 http](08_web_development_teil1.md)
- Http Basics
- Templating
- Nette Features im Http Package
- `http.Client`
- Package `http/httptest`

### [09 http extended](09_web_development_teil2.md)
- HTTP2 mit Go
- Graceful Shutdown
- Middleware
- Context
- Web Frameworks & alternative Router
- Web Sockets
- Databases mit `jinzhu/gorm`

## Erweiterte Themen

### [10 Nebenlaufigkeit](10_nebenlaufigkeit.md)
- Goroutinen
- Channel
- Buffered Channel
- Das `sync` Package

### [11 Native Development](11_native_development.md)
- CGO
- Laden von Shared Libraries
- Cross Compiling

### [12 Golang im Docker Container](12_go_docker.md)

### [13 Libs](13_libs.md)
- awesome-go
- Package `time`
- Argumente und Umgebungsvariablen
- Logging
- Bolt

### [Debugging und Profiling](14_debugging_profiling.md)
- Race detection
- Debugging mit delve
- Profiling
- GODEBUG

### [Dependency Management](15_dependency_management.md)
- `vendor` Verzeichnis
- glide
- Alternativen

# Gute Videos
- [Simplicity is Complicated](https://www.youtube.com/watch?v=rFejpH_tAHM), Rob Pike
- [Concurrency Is Not Parallelism](https://www.youtube.com/watch?v=oV9rvDllKEg), Rob Pike
- [Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs), Rob Pike
- [Advanced Go Concurrency Patterns](https://www.youtube.com/watch?v=QDDwwePbDtw), Sameer Ajmani
- [Dont Just Check Errors Handle Them Gracefully](https://www.youtube.com/watch?v=lsBF58Q-DnY), Dave Cheney
