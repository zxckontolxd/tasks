package main

import (
    "fmt"
    "math/rand"
    "time"
)

func generator(ch chan<- int) {
    for {
        num := rand.Intn(100)
        ch <- num
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    ch := make(chan int)

    go generator(ch)

    for num := range ch {
        fmt.Println("Received:", num)
    }
}