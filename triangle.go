package main

import "fmt"

type Kind [] string

const (
    // Pick values for the following identifiers used by the test program.
	NaT = "Nat"
    Equ = "Equ"
    Iso = "Iso"
    Sca = "Sca"
)

func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if a < 0 || b < 0 || c < 0 {
		k = NaT
		fmt.Println(k)
		return k
	} else if a == b && b == c {
		fmt.Println(Equ)
		return Equ
	} else if a != b && b != c {
		fmt.Println(Sca)
		return Sca
	} else {
		return NaT
	}
	// triangleType := Kind {
	
	// }
}

func main() {
	KindFromSides(-2,2,2)
	KindFromSides(2,2,2)
}