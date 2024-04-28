package main

import (
    "crypto/sha256"
    "fmt"
)

func main() {
    c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("X"))
    fmt.Printf("Diff bits: %d\n", countDiffBitsSHA256(&c1, &c2))
}

func countDiffBitsSHA256(c1 *[32]byte, c2 *[32]byte) int {
    res := [32]byte{}
    for i := 0; i < 32; i++ {
        // If two bits differ the XOR operation will result in
        // a bit 1. The var res will have the location of all bits that
        // differ.
        res[i] = c1[i] ^ c2[i]
    }
    return popCount(&res)
}

func popCount(ptr *[32]byte) int {
	sum := 0
    for i := 0; i < 32; i++ {
        x := ptr[i]
        for x != 0 {
            sum++
            x = x & (x-1)
        }
    }
	return sum
}
