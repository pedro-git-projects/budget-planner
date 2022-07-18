package internals

import (
	"reflect"
	"testing"

	"github.com/bojanz/currency"
)

// TODO: Convert to table testing

// TestGetRecurringBills checks if the recurring bills, and only those bills, are returned
func TestGetRecurringBills(t *testing.T) {
	var s []bill
	b1 := NewBill("b1", "1", Recurring, Paid, "January 2, 2006")
	b2 := NewBill("b2", "1", Recurring, Paid, "January 2, 2006")
	b3 := NewBill("b3", "1", OneTime, Paid, "January 2, 2006")

	m := NewBudgetManager("10000", *b1, *b2, *b3)
	got := m.getRecurringBills()
	expected := append(s, *b1, *b2)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v but got %v\n", expected, got)
	}
}

// TestGetRecurringBills checks if the slice comes back empty when there are no recurring bills
func TestNoRecurringBills(t *testing.T) {
	var s []bill
	b1 := NewBill("b1", "1", OneTime, Paid, "January 2, 2006")
	b2 := NewBill("b2", "1", OneTime, Paid, "January 2, 2006")
	b3 := NewBill("b3", "1", OneTime, Paid, "January 2, 2006")

	m := NewBudgetManager("10000", *b1, *b2, *b3)
	got := m.getRecurringBills()
	expected := s

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v but got %v\n", expected, got)
	}
}

// TestRecurringBillsCost checks the sum of those bills that are recurring
func TestRecurringBillsCost(t *testing.T) {
	b1 := NewBill("b1", "120", Recurring, Paid, "January 2, 2006")
	b2 := NewBill("b2", "180", Recurring, Paid, "January 2, 2006")
	b3 := NewBill("b3", "350", OneTime, Paid, "January 2, 2006")
	m := NewBudgetManager("10000", *b1, *b2, *b3)

	got := m.getRecurringBillsCost()
	want, _ := currency.NewAmount("300", code)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v but got %v\n", want, got)
	}
}

// TestGetPaidSum tests if the sum getPaidSum method correctly sums paid values
func TestGetPaidSum(t *testing.T) {
	b1 := NewBill("b1", "120", Recurring, Paid, "January 2, 2006")
	b2 := NewBill("b2", "180", Recurring, Paid, "January 2, 2006")
	b3 := NewBill("b3", "350", OneTime, Pending, "January 2, 2006")

	m := NewBudgetManager("10000", *b1, *b2, *b3)
	got := m.getPaidSum()
	want, _ := currency.NewAmount("300", code)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v but got %v\n", want, got)
	}
}

// TestGetBillIdPositionTitle guarantees we're returning the correct positon in the slice
func TestGetBillIdPositionTitle(t *testing.T) {
	b1 := NewBill("b1", "120", Recurring, Paid, "January 2, 2006")  // 0
	b2 := NewBill("b2", "180", Recurring, Paid, "January 2, 2006")  // 1
	b3 := NewBill("b3", "350", OneTime, Pending, "January 2, 2006") // 2
	b4 := NewBill("b1", "350", OneTime, Pending, "January 2, 2006") // 3

	m := NewBudgetManager("10000", *b1, *b2, *b3, *b4)
	got := m.getBillPostionByTitle("b1")
	want := []int{0, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v but got %v\n", want, got)
	}
}

func TestPayBillByTitle(t *testing.T) {
	// Case 1: title is unique
	b1 := NewBill("b1", "120", Recurring, Paid, "January 2, 2006")
	b2 := NewBill("b2", "180", Recurring, Paid, "January 2, 2006")
	b3 := NewBill("b3", "350", OneTime, Pending, "January 2, 2006")
	m := NewBudgetManager("1000", *b1, *b2, *b3)
	m.payBillByTitle("b3")
	got := m.bills[2].status
	want := Paid

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v but got %v\n", want, got)
	}

	// Case 2: there is more than one title
	c1 := NewBill("c1", "120", Recurring, Paid, "January 2, 2006")
	c2 := NewBill("c2", "180", Recurring, Paid, "January 2, 2006")
	c3 := NewBill("c1", "350", OneTime, Pending, "January 2, 2006")
	m1 := NewBudgetManager("1000", *c1, *c2, *c3)
	m1.payBillByTitle("c1")
	got1 := m1.bills[2].status
	want1 := Pending

	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("expected %v but got %v\n", want, got)
	}

	// Case 3: paying the same bill multiple times
	m.payBillByTitle("b1")
	m.payBillByTitle("b1")
	m.payBillByTitle("b1")
	got2 := m.totalBalance
	want2, _ := currency.NewAmount("650", code)

	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("expected %v but got %v\n", want, got)
	}
}
