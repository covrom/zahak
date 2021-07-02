package engine

type CachedEval struct {
	Gate       uint32   // 4
	HashMove   Move     // 4
	Eval       int16    // 2
	StaticEval int16    // 2
	Age        uint16   // 2
	Depth      int8     // 1
	Type       NodeType // 1
}

func (c *CachedEval) Update(gate uint32, hashmove Move, eval int16, staticEval int16, depth int8, nodeType NodeType, age uint16) {
	c.Gate = gate
	c.HashMove = hashmove
	c.Eval = eval
	c.StaticEval = staticEval
	c.Depth = depth
	c.Type = nodeType
	c.Age = age
}

type NodeType uint8

const (
	Exact      NodeType = 1 << iota // PV-Node
	UpperBound                      // All-Node
	LowerBound                      // Cut-Node
)

var oldAge = uint16(5)

var EmptyEval = CachedEval{0, EmptyMove, 0, 0, 0, 0, 0}
var CACHE_ENTRY_SIZE = uint32(4 + 4 + 2 + 2 + 2 + 1 + 1)

type Cache struct {
	items    []CachedEval
	size     uint32
	consumed int
}

const DEFAULT_CACHE_SIZE = uint32(10)
const MAX_CACHE_SIZE = uint32(8000)

func (c *Cache) Consumed() int {
	return int((float64(c.consumed) / float64(len(c.items))) * 1000)
}

func (c *Cache) extractHashAndGate(hash uint64) (uint32, uint32) {
	gate := uint32(0x00000000FFFFFFFFF & hash)
	key := uint32(hash>>32) % uint32(len(c.items))
	return key, gate
}

func (c *Cache) Set(hash uint64, hashmove Move, eval int16, staticEval int16, depth int8, nodeType NodeType, age uint16) {
	key, gate := c.extractHashAndGate(hash)
	oldValue := c.items[key]
	if oldValue != EmptyEval {
		if gate == oldValue.Gate {
			c.items[key].Update(gate, hashmove, eval, staticEval, depth, nodeType, age)
			return
		}
		if age-oldValue.Age >= oldAge {
			c.items[key].Update(gate, hashmove, eval, staticEval, depth, nodeType, age)
			return
		}
		if oldValue.Depth > depth {
			return
		}
		if oldValue.Type == Exact || nodeType != Exact {
			return
		} else if nodeType == Exact {
			c.items[key].Update(gate, hashmove, staticEval, eval, depth, nodeType, age)
			return
		}
		c.items[key].Update(gate, hashmove, staticEval, eval, depth, nodeType, age)
	} else {
		c.consumed += 1
		c.items[key].Update(gate, hashmove, staticEval, eval, depth, nodeType, age)
	}
}

func (c *Cache) Size() uint32 {
	return c.size
}

func (c *Cache) Get(hash uint64) (Move, int16, int16, int8, NodeType, bool) {
	key, gate := c.extractHashAndGate(hash)
	item := c.items[key]
	if item.Gate == gate {
		return item.HashMove, item.Eval, item.StaticEval, item.Depth, item.Type, true
	}
	return EmptyMove, 0, 0, 0, 0, false
}

func NewCache(megabytes uint32) *Cache {
	if megabytes > MAX_CACHE_SIZE {
		return nil
	}
	size := int(megabytes * 1024 * 1024 / CACHE_ENTRY_SIZE)

	return &Cache{make([]CachedEval, roundPowerOfTwo(size)), megabytes, 0}
}

func roundPowerOfTwo(size int) int {
	var x = 1
	for (x << 1) <= size {
		x <<= 1
	}
	return x
}
