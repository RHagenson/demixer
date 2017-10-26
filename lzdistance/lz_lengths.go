package lzdistance

import (
	"crypto/md5"
	"encoding/hex"

	"bitbucket.org/rhagenson/demixer/dna"
)

// LZ78Length implements an LZ78 search and returns complexity length
func LZ78Length(seq dna.Sequence) int {
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

	return eh.Length()
}

// LZ76Length implements an LZ76 search and returns complexity length
func LZ76Length(s dna.Sequence) int {
	c, l, i, k, maxk := 1, 1, 0, 1, 1
	n := len(s) - 1

	for true {
		if s[i+k-1] != s[l+k-1] {
			if k > maxk {
				maxk = k
			}
			i++
			if i == l {
				c++
				l += maxk
				if (l + 1) > n {
					break
				} else {
					i = 0
					k = 1
					maxk = 1
				}
			} else {
				k = 1
			}
		} else {
			k++
			if (l + k) > n {
				c++
				break
			}
		}
	}
	return c
}
