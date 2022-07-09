package internals

import "fmt"

type frequency int

const (
	recurring frequency = iota
	oneTime
)

/*
	String overloads the String reciever for frequency
	printing the correct name for the type when it is correct
	and simply printing the actual integer in case it is not
*/
func (f frequency) String() string {
	switch f {
	case 0:
		return "recurring"
	case 1:
		return "one-time"
	default:
		return fmt.Sprintf("%d", int(f))
	}
}
