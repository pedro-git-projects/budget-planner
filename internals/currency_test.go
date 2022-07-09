package internals

import (
	"testing"
)

/*
	TestMultiply tests if multiplying a float by a Real will result in the correct value
	note that if it was a float multiplication we would've gotten
	0.886275000000000035 instead of 0.89
*/
func TestMultiply(t *testing.T) {
	f1 := 9.0975
	f2 := 0.0975
	f1Real := ToReal(f1)
	got := f1Real.Multiply(f2)
	want := ToReal(0.89)

	if got != want {
		t.Errorf("expected %v but got %v", want, got)
	}
}

// TestRealString guarantees the currency struct is formated as expected
func TestRealString(t *testing.T) {
	f := 0.886275000000000035
	r := ToReal(f)
	want := "R$0.89"
	got := r.String()

	if got != want {
		t.Errorf("expected %v but got %v", want, got)
	}
}
