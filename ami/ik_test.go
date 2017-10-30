package ami

import (
	"math"
	"math/rand"
	"reflect"
	"testing"

	"bitbucket.org/rhagenson/demixer/dna"
	"bitbucket.org/rhagenson/demixer/utils"
)

func TestIkReturnsForm(t *testing.T) {
	expected := "map[ami.K]ami.IkValue"
	seq := utils.RandSeq(rand.Intn(math.MaxInt16))
	combs := GenerateCombinations(5, 10, dna.ValidNucs)
	prnucs := NucProbs(&seq)
	nks := Nk(seq, combs)
	pks := Pk(nks)
	iks := Ik(pks, prnucs)

	if reflect.TypeOf(iks).String() != expected {
		t.Errorf("Expected %s, got %T", expected, prnucs)
	}
}
