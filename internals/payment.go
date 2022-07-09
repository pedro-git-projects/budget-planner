package internals

import "fmt"

type status int

const (
	paid status = iota
	pending
)

func (s status) String() string {
	switch s {
	case 0:
		return "paid"
	case 1:
		return "pending"
	default:
		return fmt.Sprintf("%d", s)
	}
}
