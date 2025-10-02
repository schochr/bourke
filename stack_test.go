package bourke

import (
	"math"
	"testing"
)

func Test_Stack_Empty(t *testing.T) {
	stack := NewStack[int](64)

	if stack.Size() != 0 {
		t.Errorf("actual: %d, expected: %d", stack.Size(), 0)
	}

	_, peekError := stack.Peek()
	_, popError := stack.Pop()

	if stack.Size() != 0 {
		t.Errorf("actual: %d, expected: %d", stack.Size(), 0)
	}
	if peekError.Error() != "empty" {
		t.Errorf("peek: error should be defined/returned")
	}
	if popError.Error() != "empty" {
		t.Errorf("pop: error should be defined/returned")
	}
}

func Test_Stack_Full(t *testing.T) {
	stack := &stack[int]{make([]int, 64), 64}
	stack.index = math.MaxInt - 1

	err := stack.Push(1)

	if err.Error() != "full" {
		t.Errorf("push MUST return error 'full'")
	}
}

func Test_Stack_MixedOperations(t *testing.T) {
	stack := NewStack[int](64)

	stack.Push(2)
	if stack.Size() != 1 {
		t.Errorf("actual: %d, expected: %d", stack.Size(), 1)
	}
	stack.Push(4)
	if stack.Size() != 2 {
		t.Errorf("actual: %d, expected: %d", stack.Size(), 2)
	}
	stack.Push(8)
	if stack.Size() != 3 {
		t.Errorf("actual: %d, expected: %d", stack.Size(), 3)
	}
	stack.Push(16)
	if stack.Size() != 4 {
		t.Errorf("actual: %d, expected: %d", stack.Size(), 4)
	}

	peek1, _ := stack.Peek()

	pop1, _ := stack.Pop()
	if stack.Size() != 3 {
		t.Errorf("actual: %d, expected: %d", stack.Size(), 3)
	}
	pop2, _ := stack.Pop()
	if stack.Size() != 2 {
		t.Errorf("actual: %d, expected: %d", stack.Size(), 2)
	}
	pop3, _ := stack.Pop()
	if stack.Size() != 1 {
		t.Errorf("actual: %d, expected: %d", stack.Size(), 1)
	}
	pop4, _ := stack.Pop()
	if stack.Size() != 0 {
		t.Errorf("actual: %d, expected: %d", stack.Size(), 0)
	}

	_, popError := stack.Pop()
	if stack.Size() != 0 {
		t.Errorf("actual: %d, expected: %d", stack.Size(), 0)
	}
	_, peekError := stack.Peek()

	if peek1 != 16 {
		t.Errorf("actual: %d, expected: %d", peek1, 0)
	}

	if pop1 != 16 {
		t.Errorf("actual: %d, expected: %d", pop1, 0)
	}
	if pop2 != 8 {
		t.Errorf("actual: %d, expected: %d", pop2, 0)
	}
	if pop3 != 4 {
		t.Errorf("actual: %d, expected: %d", pop3, 0)
	}
	if pop4 != 2 {
		t.Errorf("actual: %d, expected: %d", pop4, 0)
	}

	if peekError.Error() != "empty" {
		t.Errorf("peek: error should be defined/returned")
	}
	if popError.Error() != "empty" {
		t.Errorf("pop: error should be defined/returned")
	}
}
