package internals

import "testing"

// TestPaymentString is a table test for the String overload for the status type
func TestPaymentString(t *testing.T) {
	var stringCases = []struct {
		s    status
		want string
	}{
		{true, "paid"},
		{false, "pending"},
	}

	for _, tests := range stringCases {
		got := tests.s.String()

		if got != tests.want {
			t.Errorf("expected '%s' but got '%s'`", tests.want, got)
		}
	}
}
