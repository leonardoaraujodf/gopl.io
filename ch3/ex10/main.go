package main

import (
    "fmt"
    "bytes"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s <path>\n", os.Args[0])
        return
    }

    b := comma2(os.Args[1]);
    fmt.Println(b);
}
// intsToString is like fmt.Sprintf(values) but adds commas.
// func intsToString(values []int) string {
//     var buf bytes.Buffer
//     buf.WriteByte('[')
//     for i, v := range values {
//         if i > 0 {
//             buf.WriteString(", ")
//         }
//         fmt.Fprintf(&buf, "%d", v)
//     }
//     buf.WriteByte(']')
//     return buf.String()
// }

func comma2(s string) string {
    var buf bytes.Buffer
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
