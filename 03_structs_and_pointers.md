
## 03 structs and pointers

### structs
type definition:
```go
type Person struct {
	Name  string
	Given string
	Age   int
}
```

Usage:
```go
person := Person{
	Name:  "Doe",
	Given: "John",
	Age:   42,
}
fmt.Println(person)

person.Given = "Mike"
fmt.Println(person)
```

### Pointer
* Parameters are always passed _by value_. (Even structs)
* type safe pointers can be used
```go
a := 41

var b *int
b = &a

*b++

fmt.Println(a) // 42
```

### pointers and structs
strcut references:
```go
// copy by value
person1 := Person{
	Name: "Doe",
}
    
person2 := person1
person2.Name = "Mike"

fmt.Println(person1.Name) // Doe

// copy by reference
person3 := &person1
person3.Name = "Mike"

fmt.Println(person1.Name) // Mike
```

### pointers, slices and maps
Types created by `make()` are always pointer types.

```go
colors1 := []string{"red", "blue"}

colors2 := colors1
colors2[0] = "black"
colors2[1] = "white"
    
fmt.Println(colors1) // black, white
```

## Exercise: key-value store
Write a small program called `kv`, capable of saving and reading Key-Value pairs in/from a file.

Set value:
```shell
kv name=Doe vorname=John alter=42
```

Get value:
```shell
kv name vorname
> name=Doe
> vorname=John
```

Get all values:
```shell
kv
> name=Doe
> vorname=John
> alter=42
```

