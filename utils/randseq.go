package utils

import (
	"math/rand"
	"time"

	"github.com/rhagenson/demixer/dna"
)

// RandSeq is a test utility to produce random Sequences based on ValidNucs
func RandSeq(length int) dna.Sequence {
	seq := make([]dna.Nuc, length)

	for i := 0; i < length; i++ {
		seq[i] = dna.ValidNucs[rand.Intn(len(dna.ValidNucs))]
	}
	return dna.Sequence(seq)
}

// WeightedRandSeq is a test utility to produce random Sequences based on
// ValidNucs given weights for each. weights have the same number of elements as
// ValidNucs and the sum of weights should be 100
func WeightedRandSeq(length int, weights map[dna.Nuc]int) dna.Sequence {
	if len(dna.ValidNucs) != len(weights) {
		panic("Nucs and Weights are not the same length " +
			"so cannot be treated in parallel")
	}
	totalweight := func(w map[dna.Nuc]int) int {
		sum := int(0)
		for _, v := range w {
			sum = sum + v
		}
		return sum
	}(weights)

	if totalweight != 100 {
		panic("Weights did not sum to 100.")
	}

	seq := make([]dna.Nuc, length)

	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(int(totalweight))

		for idx, val := range dna.ValidNucs {
			r -= weights[val]
			if r <= 0 {
				seq[i] = dna.ValidNucs[idx]
				break
			}
		}
	}

	return dna.Sequence(seq)
}
