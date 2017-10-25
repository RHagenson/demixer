package lzdistance

// History is the LZ production history of a Sequence following LZ grammar
type History struct {
	elems []interface{}
}

// Length returns the number of elements in the History
func (h *History) Length() int {
	return len(h.elems)
}

// Add an element to the set and returns whther element was added
func (h *History) Add(i interface{}) bool {
	h.elems = append(h.elems, i)
	return h.Contains(i)
}

// Contains whether item(s) are in set or not
func (h *History) Contains(i ...interface{}) bool {
	for _, v1 := range h.elems {
		val1 := v1
		for _, v2 := range i {
			val2 := v2
			if val1 == val2 {
				return true
			}
		}
	}
	return false
}

// NewHistory returns a History value
func NewHistory() History {
	return History{
		elems: make([]interface{}, 0),
	}
}
