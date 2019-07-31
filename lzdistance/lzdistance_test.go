package lzdistance

import (
	"math"
	"math/rand"
	"reflect"
	"testing"

	"github.com/rhagenson/demixer/dna"
	"github.com/rhagenson/demixer/utils"
)

// TestAverageLZDistanceSymmetry check whether d(S,Q) == d(Q,S)
func TestAverageLZDistanceSymmetry(t *testing.T) {
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

	// Generate forward and reverse distance in parallel
	distchan := make(chan Distance, 2)
	go func(dist chan Distance) {
		dist <- AverageLZDistance(s, q)
	}(distchan)
	go func(dist chan Distance) {
		dist <- AverageLZDistance(q, s)
	}(distchan)

	// Gather the distances and check for equality
	if <-distchan != <-distchan {
		t.Error()
	}
}

func TestAverageLZDistanceGreaterThanZero(t *testing.T) {
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

	// Generate forward and reverse distance in parallel
	dist := AverageLZDistance(q, s)

	// Check that s is not the same as q then if their distance is non-zero
	if !reflect.DeepEqual(s, q) {
		if dist == Distance(0) {
			t.Error()
		}
	}
}

func TestAverageLZDistanceOfOneSequenceIsZero(t *testing.T) {
	// Generate one random sequence
	s := utils.RandSeq(rand.Intn(math.MaxInt16))

	// Generate one distance measure
	dist := AverageLZDistance(s, s)

	// Gather the distances and check for equality
	if dist != Distance(0) {
		t.Errorf("Expected 0, but received %v", dist)
	}
}

// TestAverageNormLZDistanceSymmetry check whether d(S,Q) == d(Q,S)
func TestAverageNormLZDistanceSymmetry(t *testing.T) {
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

	// Generate forward and reverse distance in parallel
	distchan := make(chan Distance, 2)
	go func(dist chan Distance) {
		dist <- AverageNormLZDistance(s, q)
	}(distchan)
	go func(dist chan Distance) {
		dist <- AverageNormLZDistance(q, s)
	}(distchan)

	// Gather the distances and check for equality
	if <-distchan != <-distchan {
		t.Error()
	}
}

func TestAverageNormLZDistanceGreaterThanZero(t *testing.T) {
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

	// Generate forward and reverse distance in parallel
	dist := AverageNormLZDistance(q, s)

	// Check that s is not the same as q then if their distance is non-zero
	if !reflect.DeepEqual(s, q) {
		if dist == Distance(0) {
			t.Error()
		}
	}
}

func TestAverageNormLZDistanceOfOneSequenceIsZero(t *testing.T) {
	// Generate one random sequence
	s := utils.RandSeq(rand.Intn(math.MaxInt16))

	// Generate one distance measure
	dist := AverageNormLZDistance(s, s)

	// Gather the distances and check for equality
	if dist != Distance(0) {
		t.Errorf("Expected 0, but received %v", dist)
	}
}

// TestAverageNormLZDistanceSymmetry check whether d(S,Q) == d(Q,S)
func TestAverageNormLZDistanceTriangleInequality(t *testing.T) {
	// Generate three sequences at random in parallel
	seqchan := make(chan dna.Sequence, 3)
	go func(seq chan dna.Sequence) {
		seq <- utils.RandSeq(rand.Intn(math.MaxInt16))
	}(seqchan)
	go func(seq chan dna.Sequence) {
		seq <- utils.RandSeq(rand.Intn(math.MaxInt16))
	}(seqchan)
	go func(seq chan dna.Sequence) {
		seq <- utils.RandSeq(rand.Intn(math.MaxInt16))
	}(seqchan)

	// Gather the sequences
	s := <-seqchan
	q := <-seqchan
	mid := <-seqchan

	// Generate triangulation distances
	StoQ := AverageNormLZDistance(s, q)
	StoMid := AverageNormLZDistance(s, mid)
	MidToQ := AverageNormLZDistance(mid, q)

	// Gather the distances and check for equality
	if StoQ > (StoMid + MidToQ) {
		t.Error()
	}
}
