package stacks

import (
	"fmt"
	"testing"
)

func TestListConstructor(t *testing.T) {
	l := NewList()
	if l.len != 0 {
		t.Errorf("Expected: 0; but got %d", l.len)
	}
}

func TestListLen(t *testing.T) {
	l := NewList()
	l.Push(12)
	l.Push(1)
	len := l.Len()
	if len != 2 {
		t.Errorf("Expected: 0; but got %d", len)
	}
}

func TestListIsEmpty(t *testing.T) {
	l := NewList()
	l.Push(12)
	if l.IsEmpty() {
		t.Errorf("list empty %v", l.IsEmpty())
	}
}

func TestListPush(t *testing.T) {
	l := NewList()
	e := l.Push(12)
	if e.Value != 12 {
		t.Errorf("Expected Value: 12; but got %d", e.Value)
	}
}

func TestListPop(t *testing.T) {
	l := NewList()
	l.Push(12)
	l.Push(1)
	l.Push(3)
	v, err := l.Pop()
	v, err = l.Pop()
	if err != nil {
		t.Errorf(err.Error())
	}
	if v != 1 {
		t.Errorf("Expected: 1; but got %d; len %d", v, l.Len())
	}
}

func TestListIterator(t *testing.T) {
	l := NewList()
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)

	for l.HasNext() {
		fmt.Println(l.Next().Value)
	}
}
