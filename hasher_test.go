package dictionary_test

import (
	"testing"

	"github.com/mlposey/dictionary"
)

// *StringHasher.Hash should create unique results for different permutations
// of the same string.
func TestStringHasher_Hash(t *testing.T) {
	var hasher dictionary.Hasher = dictionary.NewStringHasher()

	abc := hasher.Hash("abc")
	bac := hasher.Hash("bac")

	if abc == 0 || bac == 0 {
		t.Error("0 as hash value in StringHasher")
	}

	if abc == bac {
		t.Error("collisions in StringHasher")
	}
}

func TestStringHasher_Reseed(t *testing.T) {
	var hasher dictionary.Hasher = dictionary.NewStringHasher()

	hasher.(*dictionary.StringHasher).Reseed()
}
