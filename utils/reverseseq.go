package utils

import (
	"bitbucket.org/rhagenson/demixer/dna"
)

// ReverseSeq takes a Sequence and returns its reverse equivalent
func ReverseSeq(seq dna.Sequence) dna.Sequence {
	r := seq.Bytes()
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return dna.Sequence(string(r))
}
