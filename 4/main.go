package main

import "fmt"

func difference(slice1 []string, slice2 []string) []string {
    m := make(map[string]struct{})
    for _, s := range slice2 {
        m[s] = struct{}{}
    }
	
    var ret []string
    for _, s := range slice1 {
        if _, found := m[s]; !found {
            ret = append(ret, s)
        }
    }
    return ret
}

func main() {
    slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
    slice2 := []string{"banana", "date", "fig"}
    
    result := difference(slice1, slice2)
    fmt.Println(result)
}