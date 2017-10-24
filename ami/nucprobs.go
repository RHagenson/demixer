package ami

import "strings"

// NucProbs takes a Sequence and computes the Probability of each Nuc
func NucProbs(seq *Sequence) map[Nuc]Probability {
	prnucs := make(map[Nuc]Probability)

	for _, v := range A {
		prnucs[v] = Probability(strings.Count((*seq).toString(), string(v)) / len(*seq))
	}

	return prnucs
}
