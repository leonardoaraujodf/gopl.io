package main

import "fmt"

func main() {
    str := []string{"abc", "abc", "abc", "def", "ghi", "jlm"}
    str = eliminateAdjacents(str)
    fmt.Println(str)
}

func eliminateAdjacents(str []string) []string{
    i := 1
    for j := 1; j < len(str); j++ {
        if str[j - 1] != str[j] {
            str[i] = str[j]
            i++
        }
    }

    return str[:i]
}
