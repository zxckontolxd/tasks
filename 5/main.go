package main

import "fmt"

func intersect(slice1, slice2 []int) (bool, []int) {
    m := make(map[int]struct{})
    for _, s := range slice1 {
        m[s] = struct{}{}
    }

    var intersection []int
    for _, s := range slice2 {
        if _, found := m[s]; found {
            intersection = append(intersection, s)
        }
    }

    return len(intersection) > 0, intersection
}

func main() {
    a := []int{65, 3, 58, 678, 64}
    b := []int{64, 2, 3, 43}

    hasIntersection, intersectedValues := intersect(a, b)
    fmt.Println(hasIntersection, intersectedValues)
}