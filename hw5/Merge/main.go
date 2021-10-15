package main

import (
	"fmt"
	"sync"
)

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	c1 := make(chan int, 5)
	c2 := make(chan int, 5)
	for i := 0; i <= 4; i++ {
		c1 <- i
	}
	for i := 0; i <= 4; i++ {
		c2 <- i
	}

	c3 := merge(c1, c2)
	for i := range c3 {
		fmt.Println(i)
	}
}
