// basename removes directory components and a .suffix
// e.g. a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package main

import (
    "os"
    "fmt"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s <path>\n", os.Args[0])
        return
    }

    // b := basename1(os.Args[1]);
    b := comma(os.Args[1]);
    fmt.Println(b);
}

func comma(s string) string {
    n := len(s)
    if n <= 3 {
        return s
    }

    return comma(s[:n-3]) + "," + s[n-3:]
}
