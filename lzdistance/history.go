package lzdistance

import set "github.com/deckarep/golang-set"

// History is the LZ production history of a Sequence following LZ grammar
type History struct {
	set.Set
}

// Length returns the number of elements in the History
func (h History) Length() int {
	return h.Cardinality()
}
