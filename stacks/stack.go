package stacks

type Stack struct {
	value []interface{}
	n     int
}

func NewStack() *Stack {
	s := new(Stack)
	s.value = make([]interface{}, 0, 0)
	s.n = 0

	return s
}

func (s *Stack) IsEmpty() bool {
	return s.n == 0
}

func (s *Stack) resize(size int) {

	newSlice := make([]interface{}, size+1)

	i := 0
	for _, v := range s.value {
		newSlice[i] = v
		i += 1
	}
	s.value = newSlice
}

func (s *Stack) Push(v interface{}) {
	if s.n == len(s.value) {
		s.resize(len(s.value) * 3 / 2)
	}
	s.value[s.n] = v
	s.n += 1
}
