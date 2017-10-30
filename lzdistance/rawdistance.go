package lzdistance

import (
	"bytes"
	"math"

	"bitbucket.org/rhagenson/demixer/dna"
)

// RawDistance is method one of computing the LZ-based distance between
// two Sequences. d(S,Q) = MAX{c(S,Q)-c(S), c(Q,S)-c(Q)}
func RawDistance(s, q dna.Sequence) Distance {
	if bytes.Equal(s.Bytes(), q.Bytes()) {
		return Distance(0)
	}

	cs, cq, csq, cqs := LZ76Factors(s, q)
	return RawDistanceFromFactors(cs, cq, csq, cqs)
}

// RawDistanceFromFactors implements method one of computing LZ-based distance
// given the corresponding factors
func RawDistanceFromFactors(cs, cq, csq, cqs int) Distance {
	return Distance(math.Max(float64(csq-cs), float64(cqs-cq)))
}
