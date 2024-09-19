package adt

type Stack[T interface{}] interface {
	Peek() (T, error)
	Push(T)
	Pop() (T, error)
	Size() int
	ForEach(func(T))
}
