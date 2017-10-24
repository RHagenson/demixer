package ami

import "bitbucket.org/rhagenson/demixer/dna"

// NkValue is a wrapper around an underlying usigned integer
type NkValue uint16

// Nk generates a map from Combination to number of occurences in the Sequence
func Nk(seq dna.Sequence, combs []Combination) map[Combination]NkValue {
	nks := make(map[Combination]NkValue)

	for _, comb := range combs {
		var nkval NkValue

		for i, v := range seq {
			if v == comb.Nuc1() {
				if seq[K(i)+comb.K()] == comb.Nuc2() {
					nkval = nkval + 1
				}
			}
		}
		nks[comb] = nkval
	}

	return nks
}

// SumNk sums out all Combinations and NkValues over their K factor
func SumNk(nks map[Combination]NkValue) map[K]NkValue {
	sumNk := make(map[K]NkValue)

	for key, val := range nks {
		sumNk[key.K()] = sumNk[key.K()] + val
	}

	return sumNk
}
