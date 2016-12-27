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

// Test for *Dictionary.Insert
//
// Ensure that string keys containing the same set of characters do not
// cause an infinite loop but are inserted resulting in a size increment of 1.
//
// TODO: Write tests for int and Hashable types
func TestDictionary_Insert(t *testing.T) {
	dict := NewDictionary(30)
	dict.Insert("bob", 3)
	dict.Insert("bbo", 5)
	dict.Insert("obb", 2)
	if dict.Size != 3 {
		t.Error("Error inserting values in Dictionary")
	}
}

// Test for *Dictionary.Insert
//
// Ensure that inserting into a full dictionary triggers a resize.
func TestDictionary_Insert_Full(t *testing.T) {
	dict := NewDictionary(2)
	dict.Insert("bob", 3)
	dict.Insert("bbo", 5)
	dict.Insert("obb", 2)

	if dict.Size != 3 {
		t.Error("Failed to resize full Dictionary")
	}
}

// Test for *Dictionary.Get
//
// Ensure that string keys containing the same set of characters can be
// retrieved as distinct objects.
func TestDictionary_Get_Exists(t *testing.T) {
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

// Test for *Dictionary.Get
//
// Ensure that calling Get with a nonexistent key returns nil.
func TestDictionary_Get_NoExists(t *testing.T) {
	dict := NewDictionary()
	dict.Insert("bob", 3)

	if dict.Get("obb") != nil {
		t.Error("Error retrieving nonexistent value from Dictionary")
	}
}

// Test for *Dictionary.Remove
//
// Ensure that objects which exist can be removed and that the size is reduced
// by 1 as a result.
func TestDictionary_Remove_Exists(t *testing.T) {
	dict := NewDictionary()
	dict.Insert("bob", 5)

	oldSize := dict.Size

	if err := dict.Remove("bob"); err != nil || dict.Size != oldSize-1 {
		t.Error("Failed removing object from Dictionary")
	}
}

// Test for *Dictionary.Remove
//
// Ensure that calling Remove with a nonexistent key returns an error.
func TestDictionary_Remove_NoExists(t *testing.T) {
	dict := NewDictionary()
	dict.Insert("bob", 5)

	if err := dict.Remove("obb"); err == nil {
		t.Error("Failed removing nonexistent object from Dictionary")
	}
}
