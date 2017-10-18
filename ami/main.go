// The AMI demixer is one built on Average Mutual Information (AMI) of genetics
// sequences, as defined by Bauer et al. 2008
// The measure of AMI is based on the idea that each species has a unique
// mutual information across its sequences when considering a profile
// of how many positions, k, after a given nucleotide, X, does another
// nucleotide, Y, follow.
package main

import (
	"math"
	"strings"
	"sync"
)

// A Sequence is a type of string containing only A, T, G, and C letters
type Sequence string

const minK uint16 = 5             // The minimum k value as used in Bauer et al. 2008
const maxK uint16 = math.MaxInt16 // Maximum k value to consider
var validNucs = [...]rune{'A', 'T', 'G', 'C'}

// verifySequence verifies that a given Sequence is valid
// I.e. that it is a DNA sequence, all capital letters, and other properties
// Fails early at the first failed check
func verifySequence(seq Sequence) bool {
	return checkSequenceCapitalization(seq) &&
		checkSequenceATGC(seq)
}

// checkSequenceCapitalization verifies that all letters are capitalized
func checkSequenceCapitalization(seq Sequence) bool {
	if string(seq) != strings.ToUpper(string(seq)) {
		return false
	}
	return true
}

// checkSequenceATGC verifies that seq only contains A, T, G, C characters
// TODO: Could eventually refactor to use CountNucs function
func checkSequenceATGC(seq Sequence) bool {
	if len(string(seq)) != (strings.Count(string(seq), "A") +
		strings.Count(string(seq), "T") +
		strings.Count(string(seq), "G") +
		strings.Count(string(seq), "C")) {
		return false
	}
	return true
}

// PrNuc computes the probability of each nucleotide by counting the
// number of times the nucleotide occurs in the sequence
func PrNuc(seq Sequence) map[rune]float64 {
	counts := CountNucs(seq)
	probs := map[rune]float64{
		'A': 0,
		'T': 0,
		'C': 0,
		'G': 0,
	}

	// For each rune, get count and calculate probability
	for r := range counts {
		probs[r] = float64(counts[r] / uint64(len(string(seq))))
	}

	return probs
}

// CountNucs counts the number of occurences of each valid
// DNA nucleotide (A, T, G, C) in seq and returns
// a map of rune to unsigned count
func CountNucs(seq Sequence) map[rune]uint64 {
	counts := map[rune]uint64{
		'A': 0,
		'T': 0,
		'G': 0,
		'C': 0,
	}

	for _, v := range seq {
		counts[v] = counts[v] + 1
	}

	return counts
}

// A NucNucKMap is a wrapper around sync.Map
// which uses a key generator to map from I, J, k combinations
// to a thread-safe map
type NucNucKMap struct {
	smap sync.Map
}

// Add value to thread-safe map via I, J, and k combination
// equal to val
func (n *NucNucKMap) Add(I, J rune, k uint16, val float64) {
	n.smap.Store(generateKey(I, J, k), val)
}

// Get value from thread-safe map via I, J, and k combination
// returns a uint16
func (n *NucNucKMap) Get(I, J rune, k uint16) float64 {
	if val, found := n.smap.Load(generateKey(I, J, k)); !found {
		panic("Attempted to retrieve non-existent NucNucMap key.")
	} else {
		return val.(float64)
	}
}

// generateKey is used for NucNucKMap to map I, J, k combinations into
// map keys
func generateKey(i, j rune, k uint16) rune {
	return (i + j + rune(k))
}

// NkSequence calculates the number of times a given nucleotide, I, is followed
// by another nucleotide, J, exactly k positions away.
// Only k values up to maxK are checked
func NkSequence(seq Sequence) *NucNucKMap {
	smap := new(NucNucKMap)

	for k := minK; k < maxK && k < uint16(len(string(seq))); k++ {
		// Spawn goroutine for each k value
		go func(k uint16, seq Sequence) {
			for _, i := range validNucs {
				for _, j := range validNucs {
					// Spawn goroutine for each I, J combination
					go func(i, j rune) {
						smap.Add(i, j, k, Nk(k, seq, i, j))
					}(i, j)
				}
			}
		}(k, seq)
	}
	return smap
}

// Nk counts the number of times J follows I k positions ahead in seq
func Nk(k uint16, seq Sequence, I rune, J rune) float64 {
	count := float64(0)

	for i, v := range seq {
		// Break if there is not enough seq left for k positions ahead
		if uint16(i)+k > uint16(len(string(seq))) {
			break
		}
		// Only consider indexes of I
		if v != I {
			continue
		} else {
			if rune(seq[uint16(i)+k]) == J {
				count = count + 1
			}
		}
	}
	return count
}

// PkSequence calculates the probability that I is followed by J exactly
// k positions away
// TODO: Currently Pk calls NkSequence for each I, J, k combination; this should
// be calculated once and only once
func PkSequence(seq Sequence) *NucNucKMap {
	smap := new(NucNucKMap)

	for k := minK; k < maxK && k < uint16(len(string(seq))); k++ {
		// Spawn goroutine for each k value
		go func(k uint16, seq Sequence) {
			for _, i := range validNucs {
				for _, j := range validNucs {
					// Spawn goroutine for each I, J combination
					go func(i, j rune) {
						smap.Add(i, j, k, Pk(k, seq, i, j))
					}(i, j)
				}
			}
		}(k, seq)
	}
	return smap
}

// Pk computes the probability of J following I exactly k positions
// ahead in seq
func Pk(k uint16, seq Sequence, I rune, J rune) float64 {
	nk := Nk(k, seq, I, J)
	nks := NkSequence(seq)
	total := float64(0)

	for k := minK; k < maxK && k < uint16(len(string(seq))); k++ {
		// Spawn goroutine for each k value
		go func(k uint16, seq Sequence) {
			for _, i := range validNucs {
				for _, j := range validNucs {
					// Spawn goroutine for each I, J combination
					go func(i, j rune) {
						total += total + nks.Get(i, j, k)
					}(i, j)
				}
			}
		}(k, seq)
	}
	return nk / total
}

// Ik computes the average mutual information (AMI) for a given k and given
// sequence
func Ik(k uint16, seq Sequence, I, J rune) float64 {
	pks := PkSequence(seq)
	total := float64(0)

	// Spawn goroutine for each k value
	go func(k uint16, seq Sequence) {
		for _, i := range validNucs {
			for _, j := range validNucs {
				// Spawn goroutine for each I, J combination
				go func(i, j rune) {
					total += total + pks.Get(i, j, k)
				}(i, j)
			}
		}
	}(k, seq)

	return total * math.Log10(Pk(k, seq, I, J)/(PrNuc(seq)[I]*PrNuc(seq)[J]))
}

// IkSequence generates the full AMI profile for a given Sequence
// Mapping from k value to Ik value
func IkSequence(seq Sequence) *sync.Map {
	iks := new(sync.Map)

	for k := minK; k < maxK && k < uint16(len(string(seq))); k++ {
		// Spawn goroutine for each k value
		go func(k uint16, seq Sequence) {
			for _, i := range validNucs {
				for _, j := range validNucs {
					// Spawn goroutine for each I, J combination
					go func(i, j rune) {
						iks.Store(k, Ik(k, seq, i, j))
					}(i, j)
				}
			}
		}(k, seq)
	}

	return iks
}

func main() {
	return
}
