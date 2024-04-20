package main

import (
    "fmt"
    "os"
    "strings"
)

func areAnagrams(s1, s2 string) bool {
    if len(s1) == 0 || len(s2) == 0 {
        return false
    } else if len(s1) != len(s2) {
        return false
    }

    for i := 0; i < len(s1); i++ {
        // fmt.Printf("s1[%d:%d] = %s\n", i, i+1, s1[i:i+1])
        // fmt.Printf("s2 = %s\n", s2)
        idx := strings.Index(s2, s1[i:i+1])
        if idx == -1 {
            // fmt.Printf("Here 1!")
            return false
        }

        // fmt.Printf("Before trim s2 = %s\n", s2)
        s2 = strings.Replace(s2, s1[i:i+1], "", 1)
        // fmt.Printf("After trim s2 = %s\n", s2)
    }

    if len(s2) != 0 {
        return false
        // fmt.Printf("Here 2!")
    }

    return true
}

func main() {
    if len(os.Args) != 3 {
        fmt.Fprintf(os.Stderr, "Usage: %s <string1> <string2>\n", os.Args[0])
        return
    }

    v := areAnagrams(os.Args[1], os.Args[2]);
    if v == true {
        fmt.Println("True")
    } else {
        fmt.Println("False")
    }
}
