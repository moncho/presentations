package main

import "fmt"

//PoliteGreeter says hi to people
// START OMIT
type PoliteGreeter struct {
	name string
}

//Greet says hi to the given name
func (g PoliteGreeter) Greet(name string) {
	fmt.Printf("Hi %s, my name is %s", name, g.name)
}

// END OMIT
func main() {
	greeter := PoliteGreeter{name: "Gopher"}
	greeter.Greet("world")
}
