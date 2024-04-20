package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the polulation count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
	pc[byte(x>>(1*8))] +
	pc[byte(x>>(2*8))] +
	pc[byte(x>>(3*8))] +
	pc[byte(x>>(4*8))] +
	pc[byte(x>>(5*8))] +
	pc[byte(x>>(6*8))] +
	pc[byte(x>>(7*8))] +
	pc[byte(x>>(8*8))])
}

// PopCount2 uses a for loop instead of a single expression as done
// PopCount
func PopCount2(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}

	return sum
}

// PopCount3 counts bits by shifting its argument through 64 bit positions, testing the rightmost bit each time.
func PopCount3(x uint64) int {
	sum := 0;
	for i := 0; i < 64; i++ {
		if ((x >> i) & 0x01) == 0x01 {
			sum++
		}
	}
	return sum
}

func PopCount4(x uint64) int {
	sum := 0
	for x != 0 {
		sum++
		x = x & (x-1)
	}
	return sum
}
