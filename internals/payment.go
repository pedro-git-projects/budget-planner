package internals

import "fmt"

type Status int

const (
	Paid Status = iota
	Pending
)

func (s Status) String() string {
	switch s {
	case 0:
		return "paid"
	case 1:
		return "pending"
	default:
		return fmt.Sprintf("%d", s)
	}
}
