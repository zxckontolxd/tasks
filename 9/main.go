package main

import (
	"fmt"

)

func main() {
	uint8Channel := make(chan uint8)
	float64Channel := make(chan float64)

	go func () {
		for i := range uint8Channel {
			ret := float64(i) * float64(i) * float64(i)
			float64Channel <- ret
		}
		close(float64Channel)
	}()

	go func () {
		for i := range float64Channel {
			fmt.Println(i)
		}
	}()

	for i := uint8(1); i < 10; i++ {
		uint8Channel <- i
	}
	close(uint8Channel)
}