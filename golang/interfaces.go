package main

import "fmt"

// START OMIT
type person string

//Greeter interface for greeters
type Greeter interface {
	Greet(person)
}

func greet(g Greeter, somebody person) {
	g.Greet(somebody)
}

func main() {
	greet(PoliteGreeter{name: "Gopher"}, "world")
}

//END OMIT

//PoliteGreeter says hi to people
type PoliteGreeter struct {
	name string
}

//Greet says hi to the given name
func (g PoliteGreeter) Greet(somebody person) {
	fmt.Printf("Hi %s, my name is %s", somebody, g.name)
}
