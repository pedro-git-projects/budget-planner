package internals

import "fmt"

// Real represents the Brazilian Real amount in terms of cents
type Real int64

// ToReal converts a float to Real
func ToReal(f float64) Real {
	return Real((f * 100) + 0.5)
}

/*
	float64 overloads the float64 casting reciever function for Real
	the difference in implementation is that the value will be safely
 	rounded to the nearest cent
*/
func (r Real) float64() float64 {
	x := float64(r)
	x = x / 100
	return x
}

/*
	Multiply safely multiplies Real values by floats
	rounding to the nearest cent
*/
func (r Real) Multiply(f float64) Real {
	x := (float64(r) * f) + 0.5
	return Real(x)
}

// String overloads the String reciever for Real
func (r Real) String() string {
	x := float64(r)
	x = x / 100
	return fmt.Sprintf("R$%.2f", x)
}
