package hash

import (
	"math/rand"
	"sync"
)

type Hasher interface {
	Hash(x interface{}) uint32
}

// StringHasher uses tabulation hashing to hash strings.
//
// This implementation uses four tables, each with 256 random uint32 values. For more
// information on tabulation hashing, see https://en.wikipedia.org/wiki/Tabulation_hashing
type StringHasher struct {
	tableCount int
	tableSize  int
	// 4 tables, each storing a random number for the possible ascii values
	tables [][]uint32
}

// New initializes and returns a new *StringHasher object.
func NewStringHasher() *StringHasher {
	hasher := &StringHasher{
		tableCount: 4,
		tableSize:  256,
	}

	hasher.tables = make([][]uint32, hasher.tableCount)
	for t := range hasher.tables {
		hasher.tables[t] = make([]uint32, hasher.tableSize)
	}

	hasher.makeTables()
	return hasher
}

// Hash uses tabulation hashing to turn a string into a uint32 value.
func (h *StringHasher) Hash(x interface{}) uint32 {
	if len(x.(string)) > h.tableCount {
		return h.Hash(x.(string)[0:h.tableCount]) ^
			h.Hash(x.(string)[h.tableCount:])
	} else {
		result := h.tables[0][x.(string)[0]]
		xlen := len(x.(string))
		for i := 1; i < h.tableCount && i < xlen; i++ {
			result ^= h.tables[i][x.(string)[i]]
		}
		return result
	}
}

// makeTables assigns to each table index a random uint32 value.
func (h *StringHasher) makeTables() {
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
