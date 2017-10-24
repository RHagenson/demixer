package dna

// Nuc is a more abstract represention of a single nucleotide byte
type Nuc byte

// ValidNucs is the set of valid DNA nucleotides
var ValidNucs = []Nuc{
	'A',
	'T',
	'G',
	'C',
}
