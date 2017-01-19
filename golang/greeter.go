package main

import "fmt"

//PoliteGreeter says hi to people
// START OMIT
type PoliteGreeter struct {
	name string
}

type person string

//Greet says hi to the given name
func (g PoliteGreeter) Greet(p person) {
	fmt.Printf("Hi %s, my name is %s", p, g.name)
}

// END OMIT
func main() {
	greeter := PoliteGreeter{name: "Gopher"}
	greeter.Greet("world")
}
