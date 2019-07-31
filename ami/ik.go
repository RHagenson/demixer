package ami

import (
	"math"

	"github.com/rhagenson/demixer/dna"
)

// IkValue is the measure of AMI at a given K value
type IkValue float64

// Ik generates a map from K to IkValue
func Ik(pks map[Combination]Probability, prnucs map[dna.Nuc]Probability) map[K]IkValue {
	iks := make(map[K]IkValue)

	for key, val := range pks {
		prnuc1 := prnucs[key.Nuc1()]
		prnuc2 := prnucs[key.Nuc2()]
		extra := IkValue(val) * IkValue(math.Log10(float64(val/(prnuc1*prnuc2))))

		iks[key.K()] = iks[key.K()] + extra
	}

	return iks
}
