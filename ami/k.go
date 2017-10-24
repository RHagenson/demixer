package ami

// K is a more abstract represention of distance between nucleotides
type K uint16

var minK K = 5   // The minimum k value as used in Bauer et al. 2008
var maxK K = 512 // The maximum k value as used in Bauer et al. 2008
