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
	id        uint
	cost      Real
	frequency Frequency
	status    Status
	due       time.Time
}

func NewBill(title string, id uint, cost Real, frequency Frequency, status Status, due string) *bill {
	b := new(bill)
	b.title = title
	b.id = id
	b.cost = cost
	b.frequency = frequency
	b.status = status
	b.due, _ = time.Parse(layoutUS, due)
	return b
}

func (b bill) String() string {
	return fmt.Sprintf("title: %s\n amount due ON or BEFORE %v: %v\n payment status: %v \n this bill is %s", b.title, b.due.Format(layoutUS), b.cost, b.status, b.frequency)
}
