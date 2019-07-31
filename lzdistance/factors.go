package lzdistance

import "github.com/rhagenson/demixer/dna"

// LZ76Factors takes two Sequences and computes the four possible
// factorizations c(S), c(Q), c(S,Q), and c(Q,S)
func LZ76Factors(s, q dna.Sequence) (cs, cq, csq, cqs int) {
	cs = LZ76Length(s)
	cq = LZ76Length(q)
	csq = LZ76Length(s.AppendSeq(q))
	cqs = LZ76Length(q.AppendSeq(s))

	return
}

// LZ78Factors takes two Sequences and computes the four possible
// factorizations c(S), c(Q), c(S,Q), and c(Q,S)
func LZ78Factors(s, q dna.Sequence) (cs, cq, csq, cqs int) {
	cs = LZ78Length(s)
	cq = LZ78Length(q)
	csq = LZ78Length(s.AppendSeq(q))
	cqs = LZ78Length(q.AppendSeq(s))

	return
}
