package hash

import "testing"

// IsPowerOfTwo returns true if potentialPower is a power of two.
func IsPowerOfTwo(potentialPower int) bool {
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
	nonPower := 234
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
	truePower := 4
	dict := NewDictionary(truePower)

	if dict.tableSize != truePower {
		t.Error("Valid NewDictionary size was changed")
	}
}

func TestDictionary_Insert(t *testing.T) {
	dict := NewDictionary(30)
	dict.Insert("bob", 3)
	dict.Insert("bbo", 5)
	dict.Insert("obb", 2)
	if dict.Size != 3 {
		t.Error("Error inserting values in Dictionary")
	}
}

func TestDictionary_Get(t *testing.T) {
	dict := NewDictionary()
	dict.Insert("bob", 3)
	dict.Insert("bbo", 5)
	dict.Insert("obb", 2)

	if dict.Get("bob").(int) != 3 ||
		dict.Get("bbo").(int) != 5 ||
		dict.Get("obb").(int) != 2 {
		t.Error("Error retrieving values from Dictionary")
	}
}

func TestDictionary_Remove(t *testing.T) {
	dict := NewDictionary()
	dict.Insert("bob", 5)

	oldSize := dict.Size

	if err := dict.Remove("bob"); err != nil || dict.Size != oldSize-1 {
		t.Error("Failed removing object from Dictionary")
	}
}
