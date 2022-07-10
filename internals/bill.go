package internals

import (
	"fmt"
	"time"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

type bill struct {
	title     string
	id        int
	cost      Real
	frequency Frequency
	status    Status
	due       time.Time
}

// NewBill returns a new bill instance holding all the specified information
func NewBill(title string, cost float64, frequency Frequency, status Status, due string) *bill {
	b := new(bill)
	b.title = title
	b.id = autoIncrement.ID()
	b.cost = ToReal(cost)
	b.frequency = frequency
	b.status = status
	b.due, _ = time.Parse(layoutUS, due)
	return b
}

// String overloads the default string reciever function for the bill type
func (b bill) String() string {
	return fmt.Sprintf("title: %s\namount due ON or BEFORE %v: %v\npayment status: %v \nthis bill is %s", b.title, b.due.Format(layoutUS), b.cost, b.status, b.frequency)
}
