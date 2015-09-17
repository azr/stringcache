package stringcache

import "sync"

//String cache map implementation
//
//It can be created with Map{} or using NewMap(size)
//
//Setting Size to 0 or initializing with Map{} has the same effect,
//an infinitely growing map.
//
//With a size (> 0) the map is allocated with size and will be unlinked uppon limit reach.
//Meaning old cached strings will remain in memory until they are totally unliked.
//
//As the underlying type is map[string]string this might take more than
//2*sizeof(string) space (with string being a golang struct).
//
type Map struct {
	c map[string]string
	sync.Mutex
	size int
}

func NewMap(size int) *Map {
	if size < 0 {
		panic("stringcache.NewMap's size should be more than 0")
	}
	return &Map{
		c:    make(map[string]string, size),
		size: size,
	}
}

func (m *Map) Len() int {
	return len(m.c)
}

func (m *Map) Get(s string) string {
	m.Lock()
	defer m.Unlock()
	c, found := m.c[s]
	if found {
		return c
	}

	if m.size != 0 && len(m.c) == m.size {
		m.c = nil
	}
	if m.c == nil {
		m.c = make(map[string]string)
	}
	m.c[s] = s
	return m.c[s]
}

var _ Getter = &Map{}
