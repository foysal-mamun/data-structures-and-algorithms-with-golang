package stacks_test

import (
	"fmt"
	"github.com/foysal-mamun/data-structures-and-algorithms-with-golang/stacks"
	"strings"
)

func ExampleList() {

	l := stacks.NewList()

	for _, cmd := range strings.Split("to be or not to - be - - that - - - is", " ") {

		if cmd == "-" {
			l.Pop()
		} else {
			l.Push(cmd)
		}
	}

	for l.HasNext() {
		fmt.Println(l.Next().Value)
	}

	// Output:
	// is
	// to
}

func ExampleStack() {
	s := stacks.NewStack()
	s.Push(1)

	fmt.Println(s.IsEmpty())

	// Output:
	// true
}
