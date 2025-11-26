//go:build !solution

package lrucache

import (
	"container/list"
)

type Pair struct {
	key   int
	value int
}

type LruCache struct {
	cap   int
	dict  map[int]*list.Element
	order *list.List
}

func (c *LruCache) Get(key int) (int, bool) {
	el, ok := c.dict[key]
	if !ok {
		return 0, false
	}

	c.order.Remove(el)
	delete(c.dict, key)

	val := el.Value.(Pair).value
	c.dict[key] = c.order.PushFront(Pair{key, val})

	return val, ok
}

func (c *LruCache) Set(key, value int) {
	if c.cap == 0 {
		return
	}

	if len(c.dict) >= c.cap {
		el := c.order.Back()
		key := el.Value.(Pair).key
		delete(c.dict, key)
		c.order.Remove(el)
	}

	el, ok := c.dict[key]
	if ok {
		c.order.Remove(el)
		delete(c.dict, key)
	}

	c.dict[key] = c.order.PushFront(Pair{key, value})
}

func (c *LruCache) Range(f func(key, value int) bool) {
	l := c.order
	for e := l.Back(); e != nil; e = e.Prev() {
		v := e.Value.(Pair)
		if !f(v.key, v.value) {
			return
		}
	}
}

func (c *LruCache) Clear() {
	for c.order.Len() != 0 {
		c.order.Remove(c.order.Front())
	}
	for k := range c.dict {
		delete(c.dict, k)
	}
}

func New(cap int) Cache {
	return &LruCache{
		cap,
		make(map[int]*list.Element, cap),
		list.New(),
	}
}
