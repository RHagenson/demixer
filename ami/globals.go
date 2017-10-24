package ami

// Global variables
var minK K = 5   // The minimum k value as used in Bauer et al. 2008
var maxK K = 512 // The maximum k value as used in Bauer et al. 2008

// A is the set of valid DNA nucleotides
var A = []Nuc{
	'A',
	'T',
	'G',
	'C',
}
