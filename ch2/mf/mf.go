// Mf converts its numerical argument to meters and feet.
package main

import (
	"fmt"
	"os"
	"strconv"
	"ch2/lengthconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := lengthconv.Feet(t)
		m := lengthconv.Meters(t)
		fmt.Printf("%s = %s, %s = %s\n", f, lengthconv.FtToM(f), m, lengthconv.MToFt(m))
	}
}
