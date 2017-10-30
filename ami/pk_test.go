package ami

import (
	"math"
	"math/rand"
	"reflect"
	"testing"

	"bitbucket.org/rhagenson/demixer/dna"
	"bitbucket.org/rhagenson/demixer/utils"
)

func TestPkReturnsForm(t *testing.T) {
	expected := "map[ami.Combination]ami.Probability"
	seq := utils.RandSeq(rand.Intn(math.MaxInt16))
	combs := GenerateCombinations(localMinK, localMaxK, dna.ValidNucs)
	nks := Nk(seq, combs)
	pks := Pk(nks)

	if reflect.TypeOf(pks).String() != expected {
		t.Errorf("Expected %s, got %T", expected, reflect.TypeOf(pks).String())
	}
}

func TestPkReturnsAllCombinations(t *testing.T) {
	seq := utils.RandSeq(rand.Intn(math.MaxInt16))
	combs := GenerateCombinations(localMinK, localMaxK, dna.ValidNucs)
	nks := Nk(seq, combs)
	pks := Pk(nks)

	for _, combo := range combs {
		if _, ok := pks[combo]; !ok {
			t.Errorf("Tried finding %t in Pk() output and failed.", combo)
		}
	}
}
