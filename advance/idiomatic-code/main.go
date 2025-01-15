package main

import (
	"fmt"
	"net"
	"sync"
)

//struct initialization

type Vector struct {
	x int
	y int
	// X, Y int //also works
}

func main() {
	Vector1 := Vector{10, 20}       //wrong way
	Vector2 := Vector{x: 10, y: 20} //ideomatic way

	fmt.Println(Vector1)
	fmt.Println(Vector2)

}

// mutex grouping
// wrong way
// type Server struct {
// 	listenAddr string
// 	isRunning  bool
// 	peers      map[string]net.Conn
// 	mu         sync.RWMutex
// }

type Server struct {
	listenAddr string
	isRunning  bool

	peerLock sync.RWMutex //peers are below mutex, ie. mutex is protecting peers
	peers    map[string]net.Conn

	otherResourceLock sync.RWMutex
	otherResource     map[string]string
}

// group the constants and keep don't capitalize all the letters
const (
	Scalar       = 3.45
	X            = 3
	valueOfConst = 34
)

// variable grouping
func foo() int {

	x := 100
	y := 2
	foo := "foo"

	fmt.Println(foo)
	return x + y
}

// correct way
func fooImproved() int {
	var (
		x   = 3
		y   = 3
		foo = "dsaf"
	)
	fmt.Println(foo)
	return x + y
}

// functions that panic
// if functions panics then the func name should be preceeded by must -> mustParseIntFromString
func parseIntFromString(x string) (int, error) {

	//logic
	panic("oops")
	return 10, nil
}

//interface declaration and naming

// wrong name	//interface name should end with er ie storager/ storer/
//this interface has too much func
// type Storage interface {
// 	Get()
// 	Put()
// 	Delete()
// 	Patch()
// }

// correct Way
type Getter interface {
	Get()
}
type Putter interface {
	Put()
} //and on and on
type Storer interface {
	Getter
	Putter
	//and so on
}

// function grouping
// function sequence importand functions first ie. top of the file
// same for variables/ types/ structs
// constants always at top
func veryImportantFuncExported() {}

func veryImportantFunc() {}

func simpleUtil() {}

//less imp func go down here ...

// http handler naming
// must start with handle
func handleGetUserById() {}

func handleResizeImage() {}

// enums
type Suit byte

const (
	SuitHearts Suit = iota
	SuitClubs
	SuitDiamonds
	SuitSpades
)

// constructors should be right below the type/struct
type Order struct {
	size float64
}

//if the package name is order then the method name should be new
//then it would look like order.New because order.NewOrder dosne't look good

// if the package is not order then the name should be NewOrder
// the name should be new(nameoftheType)
func newOrder(size float64) *Order {
	//logic here
	return &Order{
		size: size,
	}
}
