package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	originalSlice := make([]int, 10)

	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100)
	}

	fmt.Println("Randomed slice: ", originalSlice)

	onlyEven := sliceExample(originalSlice)
	fmt.Println("Only even slice: ", onlyEven)

	addedSlice := addElement(originalSlice, 42)
	fmt.Println("Added slice: ", addedSlice)

	copyedSlice := copySlice(originalSlice)
	fmt.Println("Copyed slice: ", copyedSlice)

	removedSlice := removeElement(originalSlice, 3)
	fmt.Println("Removed slice: ", removedSlice)

}

func sliceExample(originalSlice []int) []int {
	ret := make([]int, 0)
	for i := range originalSlice {
		if originalSlice[i] % 2 == 0 {
			ret = append(ret, originalSlice[i])
		}
	}
	return ret
}

func addElement(originalSlice []int, num int) []int {
	ret := append(originalSlice, num)
	return ret
}

func copySlice(original []int) []int {
    copied := make([]int, len(original))
    copy(copied, original)
    return copied
}

func removeElement(slice []int, index int) []int {
    if index < 0 || index >= len(slice) {
        return slice
    }
    return append(slice[:index], slice[index+1:]...)
}