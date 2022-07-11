package internals

import (
	"fmt"
	"time"

	"github.com/bojanz/currency"
)

type bill struct {
	title     string
	id        int
	cost      currency.Amount
	frequency Frequency
	status    Status
	due       time.Time
}

// NewBill returns a new bill instance holding all the specified information
func NewBill(title string, cost string, frequency Frequency, status Status, due string) *bill {
	c, err := currency.NewAmount(cost, code)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return nil
	}
	b := new(bill)
	b.title = title
	b.id = autoIncrement.ID()
	b.cost = c
	b.frequency = frequency
	b.status = status
	b.due, _ = time.Parse(layoutUS, due)
	return b
}

// String overloads the default string reciever function for the bill type
func (b bill) String() string {
	return fmt.Sprintf("title: %s\namount due ON or BEFORE %v: %v\npayment status: %v \nthis bill is %s", b.title, b.due.Format(layoutUS), b.cost, b.status, b.frequency)
}
