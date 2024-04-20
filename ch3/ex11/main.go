
package main

import (
    "fmt"
    "bytes"
    "os"
    "strings"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s <path>\n", os.Args[0])
        return
    }

    b := comma3(os.Args[1]);
    fmt.Println(b);
}

func comma3(s string) string {
    var buf bytes.Buffer
    n := strings.Index(s, ".") 
    if n > 0 {
        s = s[:n]
    }

    if s[0] == '+' || s[0] == '-' {
        buf.WriteByte(s[0])
        s = s[1:]
    }

    if len(s) <= 3 {
        return s
    }

    rem := len(s) % 3;
    if rem > 0 {
        buf.WriteString(s[:rem])
    }
    for i := rem; i < len(s); i += 3 {
        if i != 0 {
            buf.WriteByte(',')
        }
        buf.WriteString(s[i:i+3])
    }

    return buf.String()
}
