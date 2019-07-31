package lzdistance

import (
	"math"
	"math/rand"
	"reflect"
	"testing"

	"github.com/rhagenson/demixer/dna"
)

// TestNormSumDistanceSymmetry check whether d(S,Q) == d(Q,S)
func TestNormSumDistanceSymmetry(t *testing.T) {
	// Generate two sequences at random in parallel
	seqchan := make(chan dna.Sequence, 2)
	go func(seq chan dna.Sequence) {
		seq <- randSeq(rand.Intn(math.MaxInt16))
	}(seqchan)
	go func(seq chan dna.Sequence) {
		seq <- randSeq(rand.Intn(math.MaxInt16))
	}(seqchan)

	// Gather the sequences
	s := <-seqchan
	q := <-seqchan

	// Generate forward and reverse distance in parallel
	distchan := make(chan Distance, 2)
	go func(dist chan Distance) {
		dist <- NormSumDistance(s, q)
	}(distchan)
	go func(dist chan Distance) {
		dist <- NormSumDistance(q, s)
	}(distchan)

	// Gather the distances and check for equality
	if <-distchan != <-distchan {
		t.Error()
	}
}

func TestNormSumDistanceGreaterThanZero(t *testing.T) {
	// Generate two sequences at random in parallel
	seqchan := make(chan dna.Sequence, 2)
	go func(seq chan dna.Sequence) {
		seq <- randSeq(rand.Intn(math.MaxInt16))
	}(seqchan)
	go func(seq chan dna.Sequence) {
		seq <- randSeq(rand.Intn(math.MaxInt16))
	}(seqchan)

	// Gather the sequences
	s := <-seqchan
	q := <-seqchan

	// Generate forward and reverse distance in parallel
	dist := NormSumDistance(q, s)

	// Check that s is not the same as q then if their distance is non-zero
	if !reflect.DeepEqual(s, q) {
		if dist == Distance(0) {
			t.Error()
		}
	} else {
		if dist != Distance(0) {
			t.Error()
		}
	}
}

func TestNormSumDistanceOfOneSequenceIsZero(t *testing.T) {
	// Generate one random sequence
	s := randSeq(rand.Intn(math.MaxInt16))

	// Generate one distance measure
	dist := NormSumDistance(s, s)

	// Gather the distances and check for equality
	if dist != Distance(0) {
		t.Errorf("Expected 0, but received %v", dist)
	}
}
