package main

import "fmt"

func main() {
    a := []string{"Hello", "World"}
    b := []string{"Hello", "World"}
    c := []string{"Hello", "Go!"}
    fmt.Printf("%t\n", equal(a, b))
    fmt.Printf("%t\n", equal(a, c))
    fmt.Printf("%t\n", equal(b, c))
}

func equal(x, y []string) bool {
    if len(x) != len(y) {
        return false
    }

    for i := range x {
        if x[i] != y[i] {
            return false
        }
    }

    return true
}
