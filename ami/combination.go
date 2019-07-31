package ami

import "github.com/rhagenson/demixer/dna"

// Combination is a permuation of k, I, and J values
// where k is the distance between I and J, while I and J are both valid
// nucleotides
type Combination struct {
	i dna.Nuc
	j dna.Nuc
	k K
}

// Nuc1 returns the nucleotide that is internally first in the pair
func (comb *Combination) Nuc1() dna.Nuc {
	return comb.i
}

// Nuc2 returns the nucleotide that is internally second in the pair
func (comb *Combination) Nuc2() dna.Nuc {
	return comb.j
}

// K returns the distance between the nucleotide pair
func (comb *Combination) K() K {
	return comb.k
}

// NewCombination generates a new Combination from values
func NewCombination(i, j dna.Nuc, k K) Combination {
	return Combination{
		i: i,
		j: j,
		k: k,
	}
}

// GenerateCombinations generates every permutation of k, I, and J values
// TODO: The length of the sequence Combinations are being generated for
// should also be passed so that only Combinations that are short enough
// to be possible are generated.
func GenerateCombinations(minK, maxK K, nucs []dna.Nuc) []Combination {
	combs := make([]Combination, int(maxK-minK)*len(nucs)*len(nucs))

	for k := minK; k <= maxK; k++ {
		for idx1, i := range nucs {
			for idx2, j := range nucs {
				combs[int(k-minK)+idx1+idx2] = NewCombination(i, j, k)
			}
		}
	}

	return combs
}

// PruneCombinations removes any combinations that are not possible
// in the Sequence. Primarily removing any Combinations with too large
// a K value
func PruneCombinations(seq dna.Sequence, combs []Combination) []Combination {
	for idx, val := range combs {
		// Too long of K
		if int(val.K()) > len(seq) {
			combs = append(combs[:idx], combs[idx+1:]...)
		}
	}
	return combs
}
