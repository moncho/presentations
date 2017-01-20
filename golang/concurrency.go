package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var wg sync.WaitGroup

//START 2OMIT

type ball struct {
	hits int
}

func main() {
	court := make(chan *ball)
	b := &ball{0}
	wg.Add(2)

	go player("Nadal", court)
	go player("Federer", court)

	court <- b
	wg.Wait()
	fmt.Printf("What a game!! The ball was hit %d times\n", b.hits)

}

//END 2OMIT

//START OMIT

func player(name string, court chan *ball) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Println(name, "won!!")
			return
		}

		if rand.Intn(10) == 0 {
			fmt.Println(name, "missed the ball")
			close(court)
			return
		}

		fmt.Println(name, "hit the ball")
		ball.hits = ball.hits + 1
		court <- ball
	}
}

//END OMIT
