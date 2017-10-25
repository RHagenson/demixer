package dna

import "strings"

// Sequence is simply a string of nucleotides
type Sequence []Nuc

// AppendSeq concatenates two Sequences together and returns the result
func (seq Sequence) AppendSeq(other Sequence) Sequence {
	return append(seq, other...)
}

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

	for _, v := range ValidNucs {
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
	return strings.ToUpper((*seq).ToString()) == (*seq).ToString()
}

// ToString returns the string representation of the DNA Sequence
func (seq *Sequence) ToString() string {
	bytes := make([]byte, len(*seq))
	for i, v := range *seq {
		bytes[i] = byte(v)
	}
	return string(bytes)
}

// Bytes outputs the []byte equivalent of the Sequence
func (seq *Sequence) Bytes() []byte {
	bytes := make([]byte, len(*seq))

	for i, v := range *seq {
		bytes[i] = byte(v)
	}

	return bytes
}
