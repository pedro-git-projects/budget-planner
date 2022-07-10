package internals

import (
	"reflect"
	"testing"
)

// TODO: Convert to table testing

// TestGetRecurringBills checks if the recurring bills, and only those bills, are returned
func TestGetRecurringBills(t *testing.T) {
	var s []bill
	b1 := NewBill("b1", 1, 1, Recurring, Paid, "January 2, 2006")
	b2 := NewBill("b2", 1, 1, Recurring, Paid, "January 2, 2006")
	b3 := NewBill("b3", 1, 1, OneTime, Paid, "January 2, 2006")

	m := NewBudgetManager(10000, *b1, *b2, *b3)
	got := m.getRecurringBills()
	expected := append(s, *b1, *b2)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v but got %v\n", expected, got)
	}
}

// TestGetRecurringBills checks if the slice comes back empty when there are no recurring bills
func TestNoRecurringBills(t *testing.T) {
	var s []bill
	b1 := NewBill("b1", 1, 1, OneTime, Paid, "January 2, 2006")
	b2 := NewBill("b2", 1, 1, OneTime, Paid, "January 2, 2006")
	b3 := NewBill("b3", 1, 1, OneTime, Paid, "January 2, 2006")

	m := NewBudgetManager(10000, *b1, *b2, *b3)
	got := m.getRecurringBills()
	expected := s

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v but got %v\n", expected, got)
	}
}

// TestRecurringBillsCost checks the sum of those bills that are recurring
func TestRecurringBillsCost(t *testing.T) {
	b1 := NewBill("b1", 1, 120, Recurring, Paid, "January 2, 2006")
	b2 := NewBill("b2", 1, 180, Recurring, Paid, "January 2, 2006")
	b3 := NewBill("b3", 1, 350, OneTime, Paid, "January 2, 2006")
	m := NewBudgetManager(10000, *b1, *b2, *b3)

	got := m.getRecurringBillsCost()
	want := ToReal(300)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v but got %v\n", want, got)
	}
}
