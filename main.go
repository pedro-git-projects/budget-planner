package main

import (
	"fmt"

	"github.com/pedro-git-projects/budget-planner/internals"
)

func main() {
	b1 := internals.NewBill("title", "1.10", internals.Recurring, 0, "January 2, 2006")
	b2 := internals.NewBill("title", "1", internals.Recurring, 0, "January 2, 2006")
	i := internals.NewBudgetManager("100", *b1, *b2)
	fmt.Println(i)
}
