package hash

import (
	"hash"
	"hash/fnv"
	"math/rand"
	"sync"
)

type Hasher interface {
	Hash(x interface{}) uint32
}

//----------begin StringHasher----------

// StringHasher uses tabulation hashing to hash strings.
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

//----------end StringHasher----------

//----------begin Int32Hasher----------

// Int32Hasher wraps built-in fnv hashing in the Hasher interface.
type Int32Hasher struct {
	fnvHash hash.Hash32
}

// NewInt32Hasher initializes and returns an *Int32Hasher object.
func NewInt32Hasher() *Int32Hasher {
	return &Int32Hasher{fnv.New32()}
}

// Hash uses built-in fnv hashing to hash an int32 value.
func (h *Int32Hasher) Hash(x interface{}) uint32 {
	h.fnvHash.Reset()
	h.fnvHash.Write([]byte(x.(int32)))
	return h.fnvHash.Sum32()
}

//----------end Int32Hasher----------
