package ami

import (
	"math"
	"math/rand"
	"reflect"
	"testing"

	"bitbucket.org/rhagenson/demixer/dna"
	"bitbucket.org/rhagenson/demixer/utils"
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
