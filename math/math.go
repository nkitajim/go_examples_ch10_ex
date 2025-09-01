package math

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

// Add: return a + b
// [See]: https://www.mathsisfun.com/numbers/addition.html
func Add [T Number](a T, b T) T {
	return a + b
}
