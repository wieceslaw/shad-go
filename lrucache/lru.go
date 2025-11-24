//go:build !solution

package lrucache

type lruCache struct {
	cap  int
	dict map[int]int
}

func (c *lruCache) Get(key int) (int, bool) {
	i, ok := c.dict[key]
	return i, ok
}

func (c *lruCache) Set(key, value int) {
	c.dict[key] = value
}

func (c *lruCache) Range(f func(key, value int) bool) {

}

func (c *lruCache) Clear() {
	for k := range c.dict {
		delete(c.dict, k)
	}
}

func New(cap int) Cache {
	return &lruCache{cap, make(map[int]int)}
}
