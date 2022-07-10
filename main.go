package main

import (
	"calc/internals"
)

func main() {
	b1 := internals.NewBill("title", 1, 1, internals.Recurring, 0, "January 2, 2006")
	internals.NewBudgetManager(internals.ToReal(100), *b1)
}
