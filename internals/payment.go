package internals

type status bool

func (s status) String() string {
	if s == true {
		return "paid"
	} else {
		return "pending"
	}
}
