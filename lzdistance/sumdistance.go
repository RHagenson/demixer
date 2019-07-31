package lzdistance

import (
	"bytes"

	"github.com/rhagenson/demixer/dna"
)

// SumDistance is method three of computing LZ-based distance between
// two sequences. ds(S,Q) = c(S,Q) - c(S) + c(Q,S) - c(Q)
func SumDistance(s, q dna.Sequence) Distance {
	if bytes.Equal(s.Bytes(), q.Bytes()) {
		return Distance(0)
	}

	cs, cq, csq, cqs := LZ76Factors(s, q)
	return SumDistanceFromFactors(cs, cq, csq, cqs)
}

// SumDistanceFromFactors implements method three of computing
// LZ-based distance given the corresponding factors
func SumDistanceFromFactors(cs, cq, csq, cqs int) Distance {
	return Distance((csq - cs) + (cqs - cq))
}
