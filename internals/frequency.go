package internals

import "fmt"

type Frequency int

const (
	Recurring Frequency = iota
	OneTime
)

/*
	String overloads the String reciever for frequency
	printing the correct name for the type when it is correct
	and simply printing the actual integer in case it is not
*/
func (f Frequency) String() string {
	switch f {
	case 0:
		return "recurring"
	case 1:
		return "one-time"
	default:
		return fmt.Sprintf("%d", int(f))
	}
}
