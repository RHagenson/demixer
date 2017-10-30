package ami

import (
	"math"
	"math/rand"
	"reflect"
	"testing"

	"bitbucket.org/rhagenson/demixer/dna"
	"bitbucket.org/rhagenson/demixer/utils"
)

const localMinK K = 5
const localMaxK K = 10

func TestNkValueCanBeUint16(t *testing.T) {
	nk := *new(NkValue)
	_ = (uint16)(nk)
}

func TestNkReturnsForm(t *testing.T) {
	expected := "map[ami.Combination]ami.NkValue"
	seq := utils.RandSeq(rand.Intn(math.MaxInt16))
	combs := GenerateCombinations(localMinK, localMaxK, dna.ValidNucs)
	nks := Nk(seq, combs)

	if reflect.TypeOf(nks).String() != expected {
		t.Errorf("Expected %s, got %T", expected, reflect.TypeOf(nks).String())
	}
}

func TestNkReturnsAllCombinations(t *testing.T) {
	seq := utils.RandSeq(rand.Intn(math.MaxInt16))
	combs := GenerateCombinations(localMinK, localMaxK, dna.ValidNucs)
	nks := Nk(seq, combs)

	for _, combo := range combs {
		if _, ok := nks[combo]; !ok {
			t.Errorf("Tried finding %T in Nk() output and failed.", combo)
		}
	}
}

func TestSumNkReturnsForm(t *testing.T) {
	expected := "map[ami.K]ami.NkValue"
	seq := utils.RandSeq(rand.Intn(math.MaxInt16))
	combs := GenerateCombinations(localMinK, localMaxK, dna.ValidNucs)
	nks := Nk(seq, combs)
	sumnks := SumNk(nks)

	if reflect.TypeOf(sumnks).String() != expected {
		t.Errorf("Expected %s, got %T", expected, reflect.TypeOf(sumnks).String())
	}
}

func TestSumNkCountsProperly(t *testing.T) {
	// NOTE: The Sequence below must be all the same nucleotide
	seq := dna.Sequence([]dna.Nuc("AAAAAAAAAAAAAAAAA"))
	combs := GenerateCombinations(localMinK, localMaxK, dna.ValidNucs)
	nks := Nk(seq, combs)
	sumnks := SumNk(nks)

	for key, val := range sumnks {
		if int(val) != len(seq)-int(key) {
			t.Errorf("SumNk is not counting properly. %v should satify property "+
				"sumnks[key] == len(seq) - int(key) given all nucleotides are "+
				"the same.", sumnks)
			break
		}
	}
}
