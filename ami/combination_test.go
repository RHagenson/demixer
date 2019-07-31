package ami

import (
	"testing"

	"github.com/rhagenson/demixer/dna"
)

// Ensure Combinations can include the same Nuc at both Nuc1 and Nuc2
func TestCombinationCanIncludeSameNuc(t *testing.T) {
	combs := GenerateCombinations(5, 10, dna.ValidNucs)

	foundEquality := false

	for _, val := range combs {
		if val.Nuc1() == val.Nuc2() {
			foundEquality = true
			break
		}
	}

	if !foundEquality {
		t.Error("Combinations including the same Nuc are not being made.")
	}
}

// Ensure Combinations include full K range
func TestCombinationsHaveFullRange(t *testing.T) {
	combs := GenerateCombinations(localMinK, localMaxK, dna.ValidNucs)
	lowerMet := false
	upperMet := false

	for _, val := range combs {
		switch val.K() {
		case K(localMinK):
			lowerMet = true
		case K(localMaxK):
			upperMet = true
		default:
			continue
		}
	}

	if !(lowerMet && upperMet) {
		t.Errorf("K values in GenerateCombinations() are not being treated "+
			"inclusively. It is %t that the lower end is inclusive and "+
			"%t that the upper end is inclusive.", lowerMet, upperMet)
	}
}
