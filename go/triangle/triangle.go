package triangle


type Kind string

const (
    NaT = "not a triangle" 
    Equ = "equilateral"
    Iso = "isosceles"
    Sca = "scalene"
)

func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if a <= 0 || b <= 0 || c <= 0 || a + b < c || a + c < b || b + c < a {
		return NaT
	}
	if a == b && b == c && a == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}
	return k
}