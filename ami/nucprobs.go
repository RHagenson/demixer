package ami

import (
	"strings"

	"github.com/rhagenson/demixer/dna"
)

// NucProbs takes a Sequence and computes the Probability of each Nuc
func NucProbs(seq *dna.Sequence) map[dna.Nuc]Probability {
	prnucs := make(map[dna.Nuc]Probability)

	for _, v := range dna.ValidNucs {
		prnucs[v] = Probability(float64(strings.Count((*seq).ToString(), string(v))) / float64(len(*seq)))
	}

	return prnucs
}
