package ami

import (
	"math"
	"math/rand"
	"reflect"
	"testing"

	"bitbucket.org/rhagenson/demixer/dna"
	"bitbucket.org/rhagenson/demixer/utils"
)

var TestCounts = []struct {
	in       dna.Sequence            // input
	expected map[dna.Nuc]Probability // expected result
}{
	{dna.Sequence([]dna.Nuc("AAAA")), map[dna.Nuc]Probability{
		dna.Nuc('A'): Probability(4.0 / 4.0),
		dna.Nuc('T'): Probability(0.0 / 4.0),
		dna.Nuc('G'): Probability(0.0 / 4.0),
		dna.Nuc('C'): Probability(0.0 / 4.0),
	}},
	{dna.Sequence([]dna.Nuc("AAAT")), map[dna.Nuc]Probability{
		dna.Nuc('A'): Probability(3.0 / 4.0),
		dna.Nuc('T'): Probability(1.0 / 4.0),
		dna.Nuc('G'): Probability(0.0 / 4.0),
		dna.Nuc('C'): Probability(0.0 / 4.0),
	}},
	{dna.Sequence([]dna.Nuc("AAAATTAATT")), map[dna.Nuc]Probability{
		dna.Nuc('A'): Probability(6.0 / 10.0),
		dna.Nuc('T'): Probability(4.0 / 10.0),
		dna.Nuc('G'): Probability(0.0 / 10.0),
		dna.Nuc('C'): Probability(0.0 / 10.0),
	}},
	{dna.Sequence([]dna.Nuc("AACGTACCATTG")), map[dna.Nuc]Probability{
		dna.Nuc('A'): Probability(4.0 / 12.0),
		dna.Nuc('T'): Probability(3.0 / 12.0),
		dna.Nuc('G'): Probability(2.0 / 12.0),
		dna.Nuc('C'): Probability(3.0 / 12.0),
	}},
}

// TestNucProbsReturnsForm checks that NucProbs returns the form that
// is expected throughout the program: a map[Nuc]Probability with proper imports
func TestNucProbsReturnsForm(t *testing.T) {
	expected := "map[dna.Nuc]ami.Probability"
	seq := utils.RandSeq(rand.Intn(math.MaxInt16))
	prnucs := NucProbs(&seq)

	if reflect.TypeOf(prnucs).String() != expected {
		t.Errorf("Expected %s, got %T", expected, prnucs)
	}
}

// TestNucProbsCountsProperly checks whether NucProbs returns the proper
// probabilities given a select few known sequences found in TestCounts above
func TestNucProbsCountsProperly(t *testing.T) {
	for _, tt := range TestCounts {
		probs := NucProbs(&tt.in)
		if !reflect.DeepEqual(probs, tt.expected) {
			t.Errorf("NucProbs(%v): expected %v, received %v", tt.in.ToString(), tt.expected, probs)
		}
	}
}
