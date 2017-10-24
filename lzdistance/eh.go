package lzdistance

import "bitbucket.org/rhagenson/demixer/dna"

// Eh computes the exhaustive history of a Sequence
func Eh(seq dna.Sequence) History {
	eh := *new(History)

	temp := make([]dna.Nuc, 0)

	for _, v := range seq {
		if !eh.Contains(append(temp, v)) {
			eh.Add(v)
			temp = make([]dna.Nuc, 0)
		} else {
			temp = append(temp, v)
		}
	}

	return eh
}
