package internals

import "testing"

// TestPaymentString is a table test for the String overload for the status type
func TestPaymentString(t *testing.T) {
	var stringCases = []struct {
		s    status
		want string
	}{
		{0, "paid"},
		{1, "pending"},
	}

	for _, tests := range stringCases {
		got := tests.s.String()

		if got != tests.want {
			t.Errorf("expected '%s' but got '%s'`", tests.want, got)
		}
	}
}
