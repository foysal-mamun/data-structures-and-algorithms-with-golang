package stacks

import (
	"errors"
)

type Element struct {
	next  *Element
	Value interface{}
}

type List struct {
	node *Element
	len  int
}

// List constractor
func NewList() *List {

	l := new(List)
	l.node = nil
	l.len = 0
	return l
}

func (l *List) HasNext() bool {
	return l.node != nil
}

func (i *List) Next() *Element {
	c := i.node
	i.node = c.next
	return c
}

func (l *List) IsEmpty() bool {
	return l.node == nil
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Pop() (interface{}, error) {

	if l.node == nil {
		return "", errors.New("list empty")
	}
	e := l.node
	l.node = e.next
	l.len--
	return e.Value, nil
}

func (l *List) Push(v interface{}) *Element {
	e := &Element{Value: v}
	e.next = l.node
	l.node = e
	l.len++

	return e
}
