package lzdistance

import (
  "testing"
  "bitbucket.org/rhagenson/demixer/dna"
)

var ehTests = []struct {
        in        dna.Sequence // input
        expected int // expected result
}{
        {dna.Sequence([]dna.Nuc{'A', 'A', 'A', 'A'}), 2}, // [A, AA]
        {dna.Sequence([]dna.Nuc{'A', 'A', 'A', 'T'}), 3}, // [A, AA, T]
        {dna.Sequence([]dna.Nuc{'A', 'A', 'T', 'T'}), 3}, // [A, AT, T]
        {dna.Sequence([]dna.Nuc{'A', 'T', 'A', 'G'}), 3}, // [A, T, AG]
        {dna.Sequence([]dna.Nuc{'A', 'T', 'G', 'C'}), 4}, // [A, T, G, C]
}

func TestEhReturnsValid(t *testing.T) {
  for _, tt := range ehTests {
                length := Eh(tt.in).Length()
                if length != tt.expected {
                        t.Errorf("Eh(%v): expected %v, received %v", tt.in, tt.expected, length)
                }
        }
}
