// Fetch prints the content found at a URL.
package main

import (
    "fmt"
    "net/http"
    "io"
    "os"
)

func main() {
    for _, url := range os.Args[1:] {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        dst := os.Stdout
        if _, err := io.Copy(dst, resp.Body); err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }
    }
}
