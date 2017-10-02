# Objects

## Methods
* Methods are functions with a _Receiver_-Variable
* Receivers can only be of a type contained in the current package

```go
type Point [2]int

func (p Point) Add(pointToAdd Point) Point {
	p[0] = p[0] + pointToAdd[0]
	p[1] = p[1] + pointToAdd[1]
	return p
}
```

## Objekte
* For Methods to be able to manipulate the state of an object, the receiver type needs to be a pointer.
* Usually structs are used for something similiar to classes.

```go
type Item struct {
	Name string
	pos  Point
}

func (item *Item) MoveTo(vector Point) {
	item.pos = item.pos.Add(vector)
}
```

## Constructors
* In go there are no explizit constructors
* The convention is to use a function like `NewTypname() *Typename` for construction.
* A lot of times the default values of struct attributes are enough and a constructor is not necessary.

```go
func NewItem(name string) *Item {
	return &Item{
		Name: name,
	}
}

func (item *Item) MoveInDirection(vector Point, time int) {
	for i := 0; i < time; i++ {
		item.pos = item.pos.Add(vector)
	}
}
```

## Embedding
* Structs can embed other struct types
* Calls can be delegated to an embedded struct automatically.

```go
type Starship struct {
	Item
	Speed int
}

func NewStarship(name string, speed int) *Starship {
	ship := &Starship{
		Speed: speed,
	}
	ship.Item.Name = name
	return ship
}

func travel() {
  herzAusGold := NewStarship("Enterprise", 42)
  herzAusGold.MoveTo(Point{1, 1})
}
```

## Method override
* Methods of embedded types can be overridden

```go
func (ship *Starship) MoveInDirection(vector Point, time int) {
	for i := 0; i < ship.Speed; i++ {
		ship.Item.MoveInDirection(vector, time)
	}
}
```

## Interfaces
* There is no inheritance, but there are interfaces.
* Interfaces are defined by behaviour, not by state.
* Interfaces in go are implemented with duck typing: If it looks like a duck, it is a duck!

```go

type myInt int

func (i myInt) String() string {
	return strconv.Itoa(int(i))
}

type Printable interface {
	String() string
}

func Test_Stringer(t *testing.T) {
	printable := []Printable{
		NewItem("Atom"),
		NewStarship("HerzAusGold", 42),
		myInt(42),
		os.ModeAppend | os.ModeSocket,
	}

	for _, p := range printable {
		fmt.Println(p.String())
	}
}
```

## Exercises

### Exercise: object oriented Key-Value store
Change the KV-store so that it uses a Store interface internally, abstracting all operations on the internal storage.

### Exercise: Testing with Mocks
* Test the Store
* Test code calling the store, but use a mock instead of the real store
