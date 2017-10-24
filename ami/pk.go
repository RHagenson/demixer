package ami

// Pk computes the probability of seeing I and J exactly k positions away from
// one another
// TODO A later improvement could be to sum the NK values as they
// are produced in Nk() then add a totalnks function argument here
func Pk(nks map[Combination]NkValue) map[Combination]Probability {
	pks := make(map[Combination]Probability)

	totalnks := SumNk(nks)

	for key, val := range nks {
		// The right side cannot be wrapped in a single
		// PkValue cast as that would result in integer division
		pks[key] = Probability(val) / Probability(totalnks[key.K()])
	}

	return pks
}
