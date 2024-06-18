package main

import "fmt"

type RotationType int

const (
    RotateLeft RotationType = iota
    RotateRight
)

func main() {
    s := []int{1, 2, 3}
    fmt.Println(s)
    rotate(RotateLeft, 1, s)
    fmt.Println(s)
    rotate(RotateRight, 1, s)
    fmt.Println(s)
}

func rotate(r RotationType, numPos int, s []int) {
    result := make([]int, len(s), cap(s))
    if (r == RotateLeft) {
        result = s[numPos:]
        result = append(result, s[0:numPos]...)
    } else {
        result = s[len(s)-numPos:]
        result = append(result, s[:len(s)-numPos]...)
    }

    copy(s, result)
}
