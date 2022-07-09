package internals

import (
	"fmt"
	"time"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

type Bill struct {
	title     string
	id        uint
	cost      Real
	frequency frequency
	status    status
	due       time.Time
}

func (b Bill) String() string {
	return fmt.Sprintf("title: %s\n amount due ON or BEFORE %v: %v\n payment status: %v \n this bill is %s", b.title, b.due.Format(layoutUS), b.cost, b.status, b.frequency)
}
