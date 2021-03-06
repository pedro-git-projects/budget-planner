package internals

import (
	"fmt"
	"strconv"

	"github.com/bojanz/currency"
)

type budgetManager struct {
	totalBalance            currency.Amount
	balanceWithoutRecurring currency.Amount
	dailyLimit              currency.Amount
	bills                   []bill
}

// NewBudgetManager returns a new BudgetManager instance holding all the specified information
func NewBudgetManager(totalBalance string, bills ...bill) *budgetManager {
	c, err := currency.NewAmount(totalBalance, code)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return nil
	}

	b := &budgetManager{
		totalBalance:            c,
		balanceWithoutRecurring: c,
		dailyLimit:              c,
		bills:                   bills,
	}

	b.setBalanceWithoutRecurring()
	b.setDailyLimit()
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
func (b budgetManager) getRecurringBillsCost() currency.Amount {
	r := b.getRecurringBills()
	sum, _ := currency.NewAmount("0", code)
	if len(r) > 0 {
		for i := 0; i < len(r); i++ {
			sum, _ = sum.Add(r[i].cost)
		}
		return sum
	}
	return sum
}

// setBalanceWithoutRecurring sets the balanceWithoutRecurring field of the budgetManager struct
func (b *budgetManager) setBalanceWithoutRecurring() {
	c := b.getRecurringBillsCost()
	diff, _ := b.totalBalance.Sub(c)
	b.balanceWithoutRecurring = diff
}

// setBalanceWithoutRecurring sets the dailyLimit field of the budgetManager struct rounding down
func (b *budgetManager) setDailyLimit() {
	n := currentMonthNumberOfDays()
	balanceS := b.totalBalance.Number()               // currency -> string
	balanceF, err := strconv.ParseFloat(balanceS, 64) // string -> float
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}
	div := balanceF / n
	limitString := fmt.Sprintf("%f", div)               // float -> string
	limit, err := currency.NewAmount(limitString, code) // string -> currency
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}
	b.dailyLimit = limit.RoundTo(2, currency.RoundDown)
}

// getPaidSum returns the sum of all paid bills
func (b budgetManager) getPaidSum() currency.Amount {
	sum, _ := currency.NewAmount("0", code)
	for i := 0; i < len(b.bills); i++ {
		if b.bills[i].status == Paid {
			sum, _ = sum.Add(b.bills[i].cost)
		}
	}
	return sum
}

// getBillIdByTitle returns the ids and positions of all bills with a certain title
func (b budgetManager) getBillPostionByTitle(t string) []int {
	var positions []int
	for i, bill := range b.bills {
		if bill.title == t {
			positions = append(positions, i)
		}
	}
	return positions
}

// payBillByTitle pays the bill with the corresponding title if it unique and informs the ID otherwise
func (b *budgetManager) payBillByTitle(t string) {
	positions := b.getBillPostionByTitle(t)
	if len(positions) > 1 {
		fmt.Sprintln(fmt.Sprintf("There is more than one bill with the title %s:\n", t))
		for i := 0; i < len(positions); i++ {
			fmt.Sprintln(fmt.Sprintf("%v of ID: %v", b.bills[positions[i]], b.bills[positions[i]].getId()))
		}
		fmt.Sprintln("Please select the desired bill by ID instead")
	}
	if len(positions) == 1 {
		err := b.bills[positions[0]].PayBill()
		if err == nil {
			b.totalBalance, _ = b.totalBalance.Sub(b.bills[positions[0]].cost)
			fmt.Sprintln("Bill paid successfully")
		} else {
			fmt.Sprintln("Bill already paid")
		}
	}
}
