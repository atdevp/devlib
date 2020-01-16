package lru

import (
	"container/list"
	"errors"
)

type LRU struct {
	Size     int
	LinkList *list.List
	M        map[interface{}]*list.Element
}

type entry struct {
	k interface{}
	v interface{}
}

func NewLRU(size int) (*LRU, error) {
	if size <= 0 {
		return nil, errors.New("size must > 0")
	}

	c := &LRU{
		Size:     size,
		LinkList: list.New(),
		M:        make(map[interface{}]*list.Element),
	}
	return c, nil
}

func (c *LRU) Exist(k interface{}) bool {
	_, exist := c.M[k]
	return exist
}

func (c *LRU) Set(k, v interface{}) {
	// existance
	if e, ok := c.M[k]; ok {
		c.LinkList.MoveToFront(e)
		e.Value.(*entry).v = v
	}

	// unexistance
	ent := &entry{k: k, v: v}
	e := c.LinkList.PushFront(ent)
	c.M[k] = e
	if c.LinkList.Len() > c.Size {
		c.RemoveOldest()
	}
}

func (c *LRU) RemoveOldest() {
	e := c.LinkList.Back()
	if e != nil {
		c.removeElement(e)
	}
}

func (c *LRU) removeElement(e *list.Element) {
	c.LinkList.Remove(e)

	ent := e.Value.(*entry)
	delete(c.M, ent.k)
}

func (c *LRU) Get(k interface{}) (interface{}, bool) {

	e, ok := c.M[k]
	if !ok {
		return nil, false
	}

	c.LinkList.MoveToFront(e)
	if e.Value.(*entry) == nil {
		return nil, false
	}
	return e.Value.(*entry).v, true
}

func (c *LRU) Length() int {
	return c.LinkList.Len()
}

func (c *LRU) Delete(k interface{}) {
	e, ok := c.M[k]
	if !ok {
		return
	}
	c.removeElement(e)
}
