package main

import (
    "fmt"
	"sync"
)

func mergeChannels(channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
	mergedChannel := make(chan int)
	merge := func(c <- chan int) {
		for n := range c {
			mergedChannel <- n
		}
		wg.Done()
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go merge(c)
	}

	go func() {
		wg.Wait()
		close(mergedChannel)
	}()

	return mergedChannel
}

func main() {
    c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go func() {
		for _, n := range []int{1, 3, 5} {
			c1 <- n
		}
		close(c1)
	}()

	go func() {
		for _, n := range []int{2, 4, 6} {
			c2 <- n
		}
		close(c2)
	}()

	go func() {
		for _, n := range []int{7, 8, 9} {
			c3 <- n
		}
		close(c3)
	}()

	result := mergeChannels(c1, c2, c3)
	for v := range result {
		fmt.Println(v)
	}
}