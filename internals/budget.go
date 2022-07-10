package internals

import (
	"fmt"
	"reflect"
)

type BudgetManager struct {
	totalBalance            Real
	balanceWithoutRecurring Real
	dailyLimit              Real
	bills                   []bill
}

func NewBudgetManager(totalBalance Real, bills ...bill) *BudgetManager {
	b := new(BudgetManager)
	b.totalBalance = totalBalance
	b.balanceWithoutRecurring = totalBalance
	b.dailyLimit = totalBalance
	b.bills = bills
	fmt.Println(reflect.TypeOf(bills))
	return b
}
