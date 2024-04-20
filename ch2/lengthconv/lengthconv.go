package lengthconv

import "fmt"

type Feet float64
type Meters float64

func (m Meters) String() string { return fmt.Sprintf("%g m", m) }
func (f Feet) String() string { return fmt.Sprintf("%g ft", f) }
