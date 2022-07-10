package internals

import "testing"

// TestStatusString is a table test for the String overload for the status type
func TestStatusString(t *testing.T) {
	var tests = []struct {
		f    Frequency
		want string
	}{
		{0, "recurring"},
		{1, "one-time"},
		{3, "3"},
	}

	for _, test := range tests {
		got := test.f.String()

		if got != test.want {
			t.Errorf("expected '%s' but got '%s'`", test.want, got)
		}
	}
}
