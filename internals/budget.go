package internals

import (
	"fmt"
)

type budgetManager struct {
	totalBalance            Real
	balanceWithoutRecurring Real
	dailyLimit              Real
	bills                   []bill
}

// NewBudgetManager returns a new BudgetManager instance holding all the specified information
func NewBudgetManager(totalBalance float64, bills ...bill) *budgetManager {
	r := ToReal(totalBalance)
	b := new(budgetManager)
	b.totalBalance = r
	b.setBalanceWithoutRecurring()
	b.dailyLimit = r
	b.bills = bills
	return b
}

// String overloads the default string reciever function for the budgetManager type
func (b budgetManager) String() string {
	if len(b.bills) > 0 {
		list, sep := "", " "
		for i := 0; i < len(b.bills); i++ {
			list += b.bills[i].String() + sep + "\n"
		}
		return fmt.Sprintf("totalBalance: %v\nbalance after recurring bills: %v\ndaily budget: %v\nbills:\n%v", b.totalBalance, b.balanceWithoutRecurring, b.dailyLimit, list)
	} else {
		return fmt.Sprintf("totalBalance: %v\nbalance after recurring bills: %v\ndaily budget: %v\n", b.totalBalance, b.balanceWithoutRecurring, b.dailyLimit)
	}
}

// getRecurringBills returns the slice of all non recurring bills if there are any and an empty slice otherwise
func (b budgetManager) getRecurringBills() []bill {
	var r []bill
	if len(b.bills) > 0 {
		for i := 0; i < len(b.bills); i++ {
			if b.bills[i].frequency == Recurring {
				r = append(r, b.bills[i])
			}
		}
		return r
	}
	return r
}

// getRecurringBillsCost returns the sum of the cost of all recurring bills and 0 if there are none
func (b budgetManager) getRecurringBillsCost() Real {
	r := b.getRecurringBills()
	if len(r) > 0 {
		sum := ToReal(0)
		for i := 0; i < len(r); i++ {
			sum += r[i].cost
		}
		return sum
	}
	return 0
}

// setBalanceWithoutRecurring sets the balanceWithoutRecurring field of the budgetManager struct
func (b *budgetManager) setBalanceWithoutRecurring() {
	c := b.getRecurringBillsCost()
	diff := b.totalBalance - c
	b.balanceWithoutRecurring = diff
}
