package internals

import (
	"fmt"
	"testing"
	"time"
)

// TestBillString guarantees the Bill struct is formated as expected
func TestBillString(t *testing.T) {
	date := "July 5, 2022"
	d, _ := time.Parse(layoutUS, date)

	b := bill{
		title:     "Jiu-Jitsu",
		id:        1,
		cost:      ToReal(230.00),
		frequency: Recurring,
		status:    Paid,
		due:       d,
	}

	got := b.String()
	s := fmt.Sprintf("title: Jiu-Jitsu\n amount due ON or BEFORE July 5, 2022: R$230.00\n payment status: paid \n this bill is recurring")
	want := s
	if got != want {
		t.Errorf("expected %v got %v", want, got)
	}
}
