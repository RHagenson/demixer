package ami

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"testing"

	"github.com/rhagenson/demixer/dna"
	"github.com/rhagenson/demixer/utils"
)

func TestEntryIsMadeOfKandIk(t *testing.T) {
	entry := *new(Entry)
	_ = entry.k
	_ = entry.ik
}

func TestProfileIsMadeOfEntries(t *testing.T) {
	expected := "[]ami.Entry"
	profile := *new(Profile)
	if reflect.TypeOf(profile.entries).String() != expected {
		t.Error("Profile entries is not a []Entry collection.")
	}
}

func TestProfileCanBeMadeFromEntries(t *testing.T) {
	_ = Profile{make([]Entry, 1)}
}

func TestNewProfileGeneratesProfile(t *testing.T) {
	expected := "ami.Profile"
	seq := utils.RandSeq(rand.Intn(math.MaxInt16))
	profile := NewProfile(&seq)

	if reflect.TypeOf(profile).String() != expected {
		t.Error("NewProfile() did not generate a proper Profile.")
	}
}

func TestNewProfileIsCorrect(t *testing.T) {
	expected := Profile{[]Entry{
		Entry{k: 5, ik: IkValue(0)},
		Entry{k: 6, ik: IkValue(0)},
		Entry{k: 7, ik: IkValue(0)},
		Entry{k: 8, ik: IkValue(1.2041199826559248)},
		Entry{k: 9, ik: IkValue(0)},
		Entry{k: 10, ik: IkValue(0)},
		Entry{k: 11, ik: IkValue(0)},
		Entry{k: 12, ik: IkValue(1.2041199826559248)},
		Entry{k: 13, ik: IkValue(0)},
		Entry{k: 14, ik: IkValue(0)},
		Entry{k: 15, ik: IkValue(0)},
	}}
	seq := dna.Sequence([]dna.Nuc("ATGCATGCATGCATGC"))
	profile := NewProfile(&seq)

	if !reflect.DeepEqual(profile, expected) {
		t.Errorf("NewProfile() did not generate the correct "+
			"known profile in testing. Expected %v, but got %v", expected, profile)
	}
}

// Since AMI takes the average mutual information across K values
// It should not matter whether the combination was A-T or T-A thus
// the forward profile should match the reverse profile
func TestAMIHasSymmetry(t *testing.T) {
	forwardseq := utils.RandSeq(rand.Intn(math.MaxInt16))
	reverseseq := utils.ReverseSeq(forwardseq)
	forwardprofile := NewProfile(&forwardseq)
	reverseprofile := NewProfile(&reverseseq)

	if !checkMatchingLengths(forwardprofile, reverseprofile) {
		t.Error("Forward and reverse profiles had non-matching lengths.")
	}
	if !checkMatchingKs(forwardprofile, reverseprofile) {
		t.Error("Forward and reverse profiles had non-matching K values.")
	}
	if !checkMatchingIks(forwardprofile, reverseprofile) {
		t.Error("Forward and reverse profiles had non-matching Ik values.")
	}
}

func checkMatchingLengths(p1, p2 Profile) bool {
	return len(p1.entries) == len(p2.entries)
}

func checkMatchingKs(p1, p2 Profile) bool {
	for idx, val := range p1.entries {
		if p2.entries[idx].k != val.k {
			return false
		}
	}
	return true
}

func checkMatchingIks(p1, p2 Profile) bool {
	const threshold float64 = 0.0001
	for idx, val := range p1.entries {
		floatik1 := float64(val.ik)
		floatik2 := float64(p2.entries[idx].ik)
		if !utils.FloatsApproxEqual(floatik1, floatik2, threshold) {
			fmt.Printf("Ik values were not within threshold (%[1]v):\n"+
				"p1[%[2]d]: %[3]v\n"+
				"p2[%[2]d]: %[4]v\n",
				threshold, idx, p2.entries[idx].ik, val.ik)
			return false
		}
	}
	return true
}
