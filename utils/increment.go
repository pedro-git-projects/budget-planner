package utils

import "sync"

// Increment is go routine safe auto incrementing type
type Increment struct {
	sync.Mutex
	id int
}

/*
	ID auto increments an ID when called upon an increment variable
	it is designed to be used in constructors
*/
func (a *Increment) ID() (id int) {
	a.Lock()
	defer a.Unlock()

	id = a.id
	a.id++
	return
}
