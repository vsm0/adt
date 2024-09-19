package stack

import "testing"

func TestStack(t *testing.T) {
	s := New[int]();

	t.Logf("Pushing 10 values...")

	for i := 0; i < 10; i++ {
		s.Push(5)
	}

	t.Logf("Reading values...")

	s.ForEach(func(v int) {
		t.Logf("%v", v)
	})

	t.Logf("Popping 10 values...")

	for s.Size() > 0 {
		if _, err := s.Pop(); err != nil {
			t.Fatalf("%v", err)
		}
	}
}
