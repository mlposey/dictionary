package hash

import "testing"

// IsPowerOfTwo returns true if potentialPower is a power of two.
func IsPowerOfTwo(potentialPower uint) bool {
	return (potentialPower & (potentialPower - 1)) == 0
}

// Test for NewDictionary
//
// Ensure that the default table size is a power of two.
func TestNewDictionary_Default(t *testing.T) {
	dict := NewDictionary()
	if !IsPowerOfTwo(dict.tableSize) {
		t.Error("Default dictionary size is not a power of two")
	}
}

// Test for NewDictionary
//
// Ensure that provided table sizes which are not powers of two
// become powers of two.
func TestNewDictionary_ToPower(t *testing.T) {
	var nonPower uint = 234
	dict := NewDictionary(nonPower)

	potentialPower := dict.tableSize

	// Make sure the table size became a power of two.
	if !IsPowerOfTwo(potentialPower) {
		t.Error("Dictionary table size not a power of two")
	}
}

// Test for NewDictionary
//
// Ensure that, if a provided table size is a power of two, it is not changed.
func TestNewDictionary_KeepSize(t *testing.T) {
	var truePower uint = 4
	dict := NewDictionary(truePower)

	if dict.tableSize != truePower {
		t.Error("Valid NewDictionary size was changed")
	}
}
