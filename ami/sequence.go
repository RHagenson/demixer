package ami

import "strings"

// Sequence is simply a string of nucleotides
type Sequence []Nuc

// Validate checks whether the given Sequence is a valid DNA string
func (seq *Sequence) Validate() bool {
	if !seq.validateLength() {
		return false
	}
	if !seq.validateCapitalization() {
		return false
	}

	return true
}

// validateLength checks that the Sequence is only made up of valid nucleotides
func (seq *Sequence) validateLength() bool {
	count := 0

	for _, v := range A {
		for _, pos := range *seq {
			if v == pos {
				count++
			}
		}
	}

	return len(*seq) == count
}

// validateCapitalization checks that the Seq is made up of only capitals
func (seq *Sequence) validateCapitalization() bool {
	return strings.ToUpper((*seq).toString()) == (*seq).toString()
}

func (seq *Sequence) toString() string {
	bytes := make([]byte, len(*seq))
	for i, v := range *seq {
		bytes[i] = byte(v)
	}
	return string(bytes)
}
