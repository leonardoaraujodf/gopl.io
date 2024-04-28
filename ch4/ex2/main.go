package main

import(
    "fmt"
    "os"
    "crypto/sha256"
    "crypto/sha512"
)

var usage string = "Usage: %s <sha256|sha384|sha512> <string>\n";

func main() {
    if len(os.Args) != 3 {
        fmt.Printf(usage, os.Args[0])
        os.Exit(1)
    }

    if os.Args[1] == "sha256" {
        c := sha256.Sum256([]byte(os.Args[2]))
        fmt.Printf("%x\n", c)
    } else if os.Args[1] == "sha384" {
        c := sha512.Sum384([]byte(os.Args[2]))
        fmt.Printf("%x\n", c)
    } else if os.Args[1] == "sha512" {
        c := sha512.Sum512([]byte(os.Args[2]))
        fmt.Printf("%x\n", c)
    } else {
        fmt.Printf(usage, os.Args[0])
        os.Exit(1)
    }
}
