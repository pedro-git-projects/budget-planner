package internals

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/bojanz/currency"
)

// TestNewBill ensures that the function is equivalent to populating a struct directly
func TestNewBill(t *testing.T) {
	got := NewBill("title", "1", Recurring, Paid, "January 2, 2006")

	date := "January 2, 2006"
	d, _ := time.Parse(layoutUS, date)
	c, _ := currency.NewAmount("1", "BRL")
	want := bill{
		title:     "title",
		id:        0,
		cost:      c,
		frequency: Recurring,
		status:    Paid,
		due:       d,
	}
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("expected %v, but got %v\n", got, want)
	}
}

// TestBillString guarantees the Bill struct is formated as expected
func TestBillString(t *testing.T) {
	date := "July 5, 2022"
	d, _ := time.Parse(layoutUS, date)
	c, _ := currency.NewAmount("230", "BRL")

	b := bill{
		title:     "Jiu-Jitsu",
		id:        1,
		cost:      c,
		frequency: Recurring,
		status:    Paid,
		due:       d,
	}

	got := b.String()
	s := fmt.Sprintf("title: Jiu-Jitsu\namount due ON or BEFORE July 5, 2022: 230 BRL\npayment status: paid \nthis bill is recurring")
	want := s
	if got != want {
		t.Errorf("\nexpected %v \ngot %v", want, got)
	}
}
