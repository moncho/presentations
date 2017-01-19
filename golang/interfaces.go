package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type person string

// START OMIT
type Greeter interface {
	Greet(person)
}
type Kicker interface {
	Kick(person)
}

type Bouncer struct {
	name string
}

func (g Bouncer) Greet(somebody person) {
	fmt.Printf("%s says: welcome %s", g.name, somebody)
}

func (g Bouncer) Kick(somebody person) {
	fmt.Printf("%s says: get the $$$$ out of here, you %s", g.name, somebody)
}

func (g Bouncer) isInGoodMood() bool {
	return rand.Intn(2) != 0
}

//END OMIT

// START 2OMIT

func in(g Greeter, somebody person) {
	g.Greet(somebody)
}
func nope(g Kicker, somebody person) {
	g.Kick(somebody)
}

func main() {
	b := Bouncer{name: "Gopher"}
	somebody := person("FunnyGuy")
	if b.isInGoodMood() {
		in(b, somebody)
	} else {
		nope(b, somebody)
	}
}

//END 2OMIT
