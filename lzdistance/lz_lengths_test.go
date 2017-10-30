package lzdistance

import (
	"testing"

	"bitbucket.org/rhagenson/demixer/dna"
)

var LZ78Tests = []struct {
	in       dna.Sequence // input
	expected int          // expected result
}{
	{dna.Sequence([]dna.Nuc("AAAA")), 2},         // [A, AA]
	{dna.Sequence([]dna.Nuc("AAAT")), 3},         // [A, AA, T]
	{dna.Sequence([]dna.Nuc("AATT")), 3},         // [A, AT, T]
	{dna.Sequence([]dna.Nuc("ATAG")), 3},         // [A, T, AG]
	{dna.Sequence([]dna.Nuc("ATGC")), 4},         // [A, T, G, C]
	{dna.Sequence([]dna.Nuc("AACGTACCATTG")), 7}, // [A, AC, G, T, ACC, AT, TG]
	{dna.Sequence([]dna.Nuc("CTAGGGACTTAT")), 8}, // [C, T, A, G, GG, AC, TT, AT]
	{dna.Sequence([]dna.Nuc("ACGGTCACCAA")), 7},  // [A, C, G, GT, CA, CC, AA]
}

var LZ76Tests = []struct {
	in       dna.Sequence // input
	expected int          // expected LZ78 result
}{
	{dna.Sequence([]dna.Nuc("AAAA")), 2},         // [A, AA]
	{dna.Sequence([]dna.Nuc("AAAT")), 2},         // [A, AA, T]
	{dna.Sequence([]dna.Nuc("AATT")), 3},         // [A, AT, T]
	{dna.Sequence([]dna.Nuc("ATAG")), 3},         // [A, T, AG]
	{dna.Sequence([]dna.Nuc("ATGC")), 4},         // [A, T, G, C]
	{dna.Sequence([]dna.Nuc("AACGTACCATTG")), 7}, // [A, AC, G, T, ACC, AT, TG]
	{dna.Sequence([]dna.Nuc("CTAGGGACTTAT")), 7}, // [C, T, A, G, AC, AT]
	{dna.Sequence([]dna.Nuc("ACGGTCACCAA")), 7},  // [A, C, G, GT, CA, CC, AA]
}

func TestLZ78Length(t *testing.T) {
	for _, tt := range LZ78Tests {
		length := LZ78Length(tt.in)
		if length != tt.expected {
			t.Errorf("LZ78Length(%v): expected %v, received %v", tt.in.ToString(), tt.expected, length)
		}
	}
}

func TestLZ76Length(t *testing.T) {
	for _, tt := range LZ76Tests {
		length := LZ76Length(tt.in)
		if length != tt.expected {
			t.Errorf("LZ76Length(%v): expected %v, received %v", tt.in.ToString(), tt.expected, length)
		}
	}
}
