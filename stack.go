package bourke

import (
	"errors"
	"math"
)

type stack[T any] struct {
	array []T
	index int
}

func (stack *stack[T]) Push(value T) error {
	stack.index++
	if stack.index >= math.MaxInt {
		return errors.New("full")
	}
	stack.array[stack.index] = value
	return nil
}

func (stack *stack[T]) Pop() (T, error) {
	var empty T
	if stack.index < 0 {
		stack.index = -1
		return empty, errors.New("empty")
	}
	top := stack.array[stack.index]
	stack.array[stack.index] = empty
	stack.index--
	return top, nil
}

func (stack *stack[T]) Peek() (T, error) {
	if stack.index < 0 {
		var empty T
		return empty, errors.New("empty")
	}
	return stack.array[stack.index], nil
}

func (stack *stack[T]) Size() int {
	return stack.index + 1
}

func (stack *stack[T]) Reset() {
	stack.index = -1
}
