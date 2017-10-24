package ami

import "bitbucket.org/rhagenson/demixer/dna"

// Entry is made up of a K and IkValue pair
type Entry struct {
	k  K
	ik IkValue
}

// Profile is a collection of AMIEntry pairs
type Profile struct {
	entries []Entry
}

// NewProfile generates the AMIProfile of a given sequence in as concurrent a fashion
// as possible
func NewProfile(seq *dna.Sequence) Profile {
	combs := GenerateCombinations(minK, maxK, dna.ValidNucs)
	nks := Nk(*seq, combs)

	// Spin up Pk calculations
	pks := Pk(nks)

	// Spin up Ik calculations
	prnucs := NucProbs(seq)
	iks := Ik(pks, prnucs)

	// TODO Make AMI profile by ordering the iks map in increasing K values
	entries := make([]Entry, int(maxK-minK))

	for index := range entries {
		entries[index].k = K(index)
		entries[index].ik = iks[K(index)]
	}

	ami := Profile{entries}

	return ami
}
