package hash

import (
	"github.com/satori/go.uuid"
	"math/rand"
	"testing"
)

// TODO: Write tests for NewDictionary

// Test for *Dictionary.Insert
//
// Ensure that string keys containing the same set of characters do not
// cause an infinite loop but are inserted resulting in a size increment of 1.
//
// TODO: Write tests for int and Hashable types
func TestDictionary_Insert_String(t *testing.T) {
	dict := NewDictionary(30)
	dict.Insert("bob", 3)
	dict.Insert("bbo", 5)
	dict.Insert("obb", 2)
	if dict.Size != 3 {
		t.Error("Error inserting string in Dictionary")
	}
}

// Test for *Dictionary.Insert
//
// Ensure that the proper Hasher is used when the key is of type int32.
func TestDictionary_Insert_Int32(t *testing.T) {
	dict := NewDictionary()
	var key int32 = 3
	dict.Insert(key, 2)
	if dict.Get(key).(int) != 2 {
		t.Error("Error inserting int32 in Dictionary")
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

func BenchmarkDictionary_Insert(b *testing.B) {
	d := NewDictionary(b.N)
	for i := 0; i < b.N; i++ {
		d.Insert(uuid.NewV4().String(), rand.Int())
	}
}

func BenchmarkMap_Insert(b *testing.B) {
	m := make(map[string]int, b.N)
	for i := 0; i < b.N; i++ {
		m[uuid.NewV4().String()] = rand.Int()
	}
}

func BenchmarkDictionary_Get(b *testing.B) {
	var keys []string
	d := NewDictionary(b.N)
	for i := 0; i < b.N; i++ {
		d.Insert(uuid.NewV4().String(), rand.Int())
	}

	var vals []int

	for _, key := range keys {
		vals = append(vals, d.Get(key).(int))
	}
}

func BenchmarkMap_Get(b *testing.B) {
	var keys []string
	m := make(map[string]int, b.N)
	for i := 0; i < b.N; i++ {
		m[uuid.NewV4().String()] = rand.Int()
	}

	var vals []int

	for _, key := range keys {
		vals = append(vals, m[key])
	}
}
