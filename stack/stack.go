package stack

import (
	"errors"
	"sync"
)

type Stack[T interface{}] struct {
	mut sync.Mutex
	frames []T
}

func New[T interface{}]() *Stack[T] {
	return &Stack[T]{
		frames: make([]T, 0),
	}
}

func (s *Stack[T]) Size() int {
	s.mut.Lock()
	defer s.mut.Unlock()
	return len(s.frames)
}

func (s *Stack[T]) Peek() (T, error) {
	l := s.Size()
	if l == 0 {
		var result T
		return result, errors.New("stack is empty")
	}
	s.mut.Lock()
	defer s.mut.Unlock()
	return s.frames[l - 1], nil
}

func (s *Stack[T]) Push(frame T) {
	s.mut.Lock()
	s.frames = append(s.frames, frame)
	s.mut.Unlock()
}

func (s *Stack[T]) Pop() (T, error) {
	l := s.Size()
	if l == 0 {
		var result T
		return result, errors.New("stack is empty")
	}
	s.mut.Lock()
	result := s.frames[l - 1]
	s.frames = s.frames[:l - 1]
	s.mut.Unlock()
	return result, nil
}

func (s *Stack[T]) ForEach(fn func(T)) {
	s.mut.Lock()
	for _, frame := range s.frames {
		fn(frame)
	}
	s.mut.Unlock()
}
