package lzdistance

import (
	"math/rand"

	"bitbucket.org/rhagenson/demixer/dna"
)

// randSeq is a test utility to produce random Sequences based on ValidNucs
func randSeq(length int) dna.Sequence {
	seq := make([]dna.Nuc, length)

	for i := 0; i < length; i++ {
		seq[i] = dna.ValidNucs[rand.Intn(len(dna.ValidNucs))]
	}
	return dna.Sequence(seq)
}
