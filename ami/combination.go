package ami

// Combination is a permuation of k, I, and J values
// where k is the distance between I and J, while I and J are both valid
// nucleotides
type Combination struct {
	i Nuc
	j Nuc
	k K
}

// Nuc1 returns the nucleotide that is internally first in the pair
func (comb *Combination) Nuc1() Nuc {
	return comb.i
}

// Nuc2 returns the nucleotide that is internally second in the pair
func (comb *Combination) Nuc2() Nuc {
	return comb.j
}

// K returns the distance between the nucleotide pair
func (comb *Combination) K() K {
	return comb.k
}

// NewCombination generates a new Combination from values
func NewCombination(i, j Nuc, k K) Combination {
	return Combination{
		i: i,
		j: j,
		k: k,
	}
}

// GenerateCombinations generates every permutation of k, I, and J values
// Outputting these permuations on a dedicated channel
func GenerateCombinations(minK, maxK K, nucs []Nuc) []Combination {
	combs := make([]Combination, int(maxK-minK)*len(nucs)*len(nucs))

	for k := minK; k < maxK; k++ {
		for idx1, i := range nucs {
			for idx2, j := range nucs {
				combs[int(k-minK)+idx1+idx2] = NewCombination(i, j, k)
			}
		}
	}

	return combs
}
