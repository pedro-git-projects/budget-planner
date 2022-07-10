package main

import (
	"calc/internals"
	"fmt"
)

func main() {
	b1 := internals.NewBill("title", 1, 1, internals.Recurring, 0, "January 2, 2006")
	i := internals.NewBudgetManager(100, *b1)
	fmt.Println(i)
}
