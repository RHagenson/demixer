package ami

import (
	"math"

	"bitbucket.org/rhagenson/demixer/dna"
)

// Entry is made up of a K and IkValue pair
type Entry struct {
	k  K
	ik IkValue
}

// Profile is a collection of AMIEntry pairs
type Profile struct {
	entries []Entry
}

// NewProfile generates the AMIProfile of a given sequence in as concurrent
// a fashion as possible
// TODO: Current Profile has NaN entries. I suspect in Ik() there is a division
// by near-zero error
func NewProfile(seq *dna.Sequence) Profile {
	combs := GenerateCombinations(minK, maxK, dna.ValidNucs)
	nks := Nk(*seq, combs)

	// Spin up Pk calculations
	pks := Pk(nks)

	// Spin up Ik calculations
	prnucs := NucProbs(seq)
	iks := Ik(pks, prnucs)

	entries := make([]Entry, 0)

	for k := minK; k < K(len(*seq)); k++ {
		if val, ok := iks[K(k)]; ok {
			if math.IsNaN(float64(val)) {
				entries = append(entries, Entry{
					k:  K(k),
					ik: IkValue(0),
				})
			} else {
				entries = append(entries, Entry{
					k:  K(k),
					ik: val,
				})
			}
		}
	}

	ami := Profile{entries}

	return ami
}
