package utils

import (
	"reflect"
	"testing"
)

// TODO: Add test for concurrent ID increment

// TestID tests if the generated id's are incremental
func TestID(t *testing.T) {
	var i Increment
	var s []int
	x := i.ID()
	y := i.ID()
	z := i.ID()
	got := append(s, x, y, z)
	want := []int{0, 1, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected: %v got: %v", got, want)
	}
}
