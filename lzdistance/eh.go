package lzdistance

import (
	"crypto/md5"
	"encoding/hex"

	"bitbucket.org/rhagenson/demixer/dna"
)

// Eh computes the exhaustive history of a Sequence
// TODO: Since the History is almost completely worthless with
// hashes stored instead of Sequences or []Nuc values it is pretty worthless to
// return here
func Eh(seq dna.Sequence) *History {
	eh := NewHistory()
	temp := make(dna.Sequence, 0)
	for _, v := range seq {
		temp = temp.AppendSeq(dna.Sequence{v})
		hasher := md5.New()
		hasher.Write(temp.Bytes())
		hashsum := hex.EncodeToString(hasher.Sum(nil))
		// Entry not in History already then add it
		if found := eh.Contains(hashsum); !found {
			// Failed to add entry to History should panic
			if added := eh.Add(hashsum); !added {
				panic("Failed to add entry to exhaustive history.")
			}
			// Clear temp collection
			temp = make(dna.Sequence, 0)
		}
	}
	return &eh
}
