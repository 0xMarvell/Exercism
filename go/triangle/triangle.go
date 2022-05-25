// Package triangle determines the type of the triangle based of the length of its sides.
package triangle

type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides checks the kind of traingle based on the length of its sides.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case a <= 0:
		return NaT
	case b <= 0:
		return NaT
	case c <= 0:
		return NaT
	case a+b < c || b+c < a || a+c < b:
		return NaT
	case a == b && b == c && c == a:
		return Equ
	case a != b && b != c && c != a:
		return Sca
	default:
		return Iso
	}
}
