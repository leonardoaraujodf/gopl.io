// basename removes directory components and a .suffix
// e.g. a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package main

import (
    "os"
    "fmt"
    "strings"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s <path>\n", os.Args[0])
        return
    }

    // b := basename1(os.Args[1]);
    b := basename2(os.Args[1]);
    fmt.Println(b);
}

func basename1(s string) string {
    // Discard last '/' and everything before.
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '/' {
            s = s[i+1:]
            break
        }
    }
    // Preserve everything before last '.'.
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '.' {
            s = s[:i]
        }
    }

    return s
}

func basename2(s string) string {
    slash := strings.LastIndex(s, "/") // -1 if "/" not found
    s = s[slash+1:]
    if dot := strings.LastIndex(s, "."); dot >= 0 {
        s = s[:dot]
    }
    return s
}
