package hash

import (
	"math/rand"
	"sync"
	"time"
)

// Hasher is a generic form of a hash function.
type Hasher interface {
	// TODO: This should be returning an uint32.
	Hash(x interface{}) int
	Reseed()
}

// MakeRand takes a nonzero value and replaces it with a pseudorandom number.
func MakeRand(x *uint32) {
	*x ^= *x << 13
	*x ^= *x >> 17
	*x ^= *x << 5
}

// StringHasher uses tabulation hashing to hash strings.
//
// This implementation uses four tables, each with 256 random uint32 values. For more
// information on tabulation hashing, see https://en.wikipedia.org/wiki/Tabulation_hashing
type StringHasher struct {
	tableCount int
	tableSize  int

	shuffleSeed uint32

	// 4 tables, each storing a random number for the possible ascii values
	tables [][]uint32
}

// New initializes and returns a new *StringHasher object.
func NewStringHasher() *StringHasher {
	hasher := &StringHasher{
		tableCount:  4,
		tableSize:   256,
		shuffleSeed: uint32(time.Now().Hour()),
	}

	hasher.GenerateTables()

	return hasher
}

// Hash uses tabulation hashing to turn a string into a uint32 value.
func (h *StringHasher) Hash(x interface{}) int {
	if len(x.(string)) > h.tableCount {
		return h.Hash(x.(string)[0:h.tableCount]) ^
			h.Hash(x.(string)[h.tableCount:])
	} else {
		result := h.tables[0][x.(string)[0]]
		xlen := len(x.(string))
		for i := 1; i < h.tableCount && i < xlen; i++ {
			result ^= h.tables[i][x.(string)[i]]
		}
		return int(result)
	}
}

// TODO: All Hasher implementations should have a Shuffle-like function.
// However, it likely needs a new name. Regen? Reseed?

// shuffle rearranges the values stored in each table.
func (h *StringHasher) Reseed() {
	// Fisher-Yates shuffle
	for t := range h.tables {
		for i := 0; i < h.tableSize; i++ {
			target := int(h.shuffleSeed) % h.tableSize
			h.tables[t][i], h.tables[t][target] = h.tables[t][target], h.tables[t][i]
			MakeRand(&h.shuffleSeed)
		}
	}
}

// GenerateTables creates a set of tables for tabulation hashing.
func (h *StringHasher) GenerateTables() {
	h.tables = make([][]uint32, h.tableCount)
	for t := range h.tables {
		h.tables[t] = make([]uint32, h.tableSize)
	}

	wg := &sync.WaitGroup{}
	wg.Add(h.tableCount)

	for t := range h.tables {
		go func(t int) {
			defer wg.Done()
			for i := range h.tables[t] {
				h.tables[t][i] = rand.Uint32()
			}
		}(t)
	}
	wg.Wait()
}

type IntHasher struct {
	int
	factor uint32
}

func NewIntHasher() *IntHasher {
	return &IntHasher{factor: rand.Uint32()}
}

func (i *IntHasher) Hash(x interface{}) int {
	return int(i.factor) * i.int
}

func (i *IntHasher) Reseed() {
	MakeRand(&i.factor)
}
