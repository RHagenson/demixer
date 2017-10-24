package lzdistance

import "bitbucket.org/rhagenson/demixer/dna"

// SumDistance is method three of computing LZ-based distance between
// two sequences. ds(S,Q) = c(S,Q) - c(S) + c(Q,S) - c(Q)
func SumDistance(s, q *dna.Sequence) Distance {
	cs, cq, csq, cqs := LZFactors(s, q)
	return SumDistanceFromFactors(cs, cq, csq, cqs)
}

// SumDistanceFromFactors implements method three of computing
// LZ-based distance given the corresponding factors
func SumDistanceFromFactors(cs, cq, csq, cqs int) Distance {
	return Distance(csq - cs + cqs - cq)
}
