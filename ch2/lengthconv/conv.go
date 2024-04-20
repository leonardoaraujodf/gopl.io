package lengthconv

// FtToM converts length in meters to feet.
func FtToM(f Feet) Meters { return Meters(f * 0.3048) }
// MtoFt converts length in feet to meters.
func MToFt(m Meters) Feet{ return Feet(m / 0.3048) }
