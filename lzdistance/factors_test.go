package lzdistance

import (
	"math"
	"math/rand"
	"testing"

	"github.com/rhagenson/demixer/dna"
	"github.com/rhagenson/demixer/utils"
)

// TestZ76FactorsHaveSubadditivity checks that the produced factors hold to:
//    1. c(SQ) - c(S) <= c(Q)
//    2. c(QS) - c(Q) <- c(S)
func TestLZ76FactorsHaveSubadditivity(t *testing.T) {
	// Generate two sequences at random in parallel
	seqchan := make(chan dna.Sequence, 2)
	go func(seq chan dna.Sequence) {
		seq <- utils.RandSeq(rand.Intn(math.MaxInt16))
	}(seqchan)
	go func(seq chan dna.Sequence) {
		seq <- utils.RandSeq(rand.Intn(math.MaxInt16))
	}(seqchan)

	// Gather the sequences
	s := <-seqchan
	q := <-seqchan

	// Generate factors
	cs, cq, csq, cqs := LZ76Factors(s, q)

	// Check c(S,Q) - c(S) <= c(Q)
	if !((csq - cs) <= cq) {
		t.Errorf("Expected (%#v - %#v) <= %#v however this was not true.", csq, cs, cq)
	}
	if !((cqs - cq) <= cs) {
		t.Errorf("Expected (%#v - %#v) <= %#v however this was not true.", cqs, cq, cs)
	}
}

// TestZ76FactorsHaveSubadditivity checks that the produced factors hold to:
//    1. c(SQ) - c(S) <= c(Q)
//    2. c(QS) - c(Q) <- c(S)
func TestLZ78FactorsHaveSubadditivity(t *testing.T) {
	// Generate two sequences at random in parallel
	seqchan := make(chan dna.Sequence, 2)
	go func(seq chan dna.Sequence) {
		seq <- utils.RandSeq(rand.Intn(math.MaxInt16))
	}(seqchan)
	go func(seq chan dna.Sequence) {
		seq <- utils.RandSeq(rand.Intn(math.MaxInt16))
	}(seqchan)

	// Gather the sequences
	s := <-seqchan
	q := <-seqchan

	// Generate factors
	cs, cq, csq, cqs := LZ78Factors(s, q)

	// Check c(S,Q) - c(S) <= c(Q)
	if !((csq - cs) <= cq) {
		t.Errorf("Expected (%#v - %#v) <= %#v however this was not true.", csq, cs, cq)
	}
	if !((cqs - cq) <= cs) {
		t.Errorf("Expected (%#v - %#v) <= %#v however this was not true.", cqs, cq, cs)
	}
}
