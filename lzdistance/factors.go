package lzdistance

import "bitbucket.org/rhagenson/demixer/dna"

// LZFactors takes two Sequences and computes the four possible
// factorizations c(S), c(Q), c(S,Q), and c(Q,S)
func LZFactors(s, q dna.Sequence) (cs, cq, csq, cqs int) {
	cs = Eh(s).Length()
	cq = Eh(q).Length()
	csq = Eh(s.AppendSeq(q)).Length()
	cqs = Eh(q.AppendSeq(s)).Length()

	return
}
