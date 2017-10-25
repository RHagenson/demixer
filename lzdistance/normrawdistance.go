package lzdistance

import (
	"math"

	"bitbucket.org/rhagenson/demixer/dna"
)

// NormRawDistance is method two of computing LZ-based distance between
// two sequences. d*(S,Q) = d(S,Q) / MAX{c(S), c(Q)}
func NormRawDistance(s, q dna.Sequence) Distance {
	cs, cq, csq, cqs := LZFactors(s, q)
	return NormRawDistanceFromFactors(cs, cq, csq, cqs)
}

// NormRawDistanceFromFactors implements method two of computing
// LZ-based distance given the corresponding factors
func NormRawDistanceFromFactors(cs, cq, csq, cqs int) Distance {
	return RawDistanceFromFactors(cs, cq, csq, cqs) /
		Distance(math.Max(float64(cs), float64(cq)))
}
