package lzdistance

import (
	"bytes"

	"github.com/rhagenson/demixer/dna"
)

// AverageLZDistance is a dangerous method that averages the Distance
// values of the four methods of computing LZ-based Sequence Distance
// The methods have bounds [0, +Inf) and [0,1] depending on being
// normalized or not. This method exists because realistically
// the upper-bound on non-normalized methods is far less than +Inf due
// to being limited by unique k-mers given the small DNA alphabet
func AverageLZDistance(s, q dna.Sequence) Distance {
	if bytes.Equal(s.Bytes(), q.Bytes()) {
		return Distance(0)
	}

	cs, cq, csq, cqs := LZ76Factors(s, q)
	raw := RawDistanceFromFactors(cs, cq, csq, cqs)
	normraw := NormRawDistanceFromFactors(cs, cq, csq, cqs)
	sum := SumDistanceFromFactors(cs, cq, csq, cqs)
	normsum := NormSumDistanceFromFactors(cs, cq, csq, cqs)

	return Distance(raw+normraw+sum+normsum) / 4
}

// AverageNormLZDistance is a "higher-confidence" method to conclude Distance
// between two Sequences by taking the average of the two [0,1] bounded methods
// of LZ-based Sequence Distance
func AverageNormLZDistance(s, q dna.Sequence) Distance {
	if bytes.Equal(s.Bytes(), q.Bytes()) {
		return Distance(0)
	}

	cs, cq, csq, cqs := LZ76Factors(s, q)
	normraw := NormRawDistanceFromFactors(cs, cq, csq, cqs)
	normsum := NormSumDistanceFromFactors(cs, cq, csq, cqs)

	return Distance(normraw+normsum) / 2
}
