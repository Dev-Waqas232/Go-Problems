package main

import (
	"errors"
)

type Stack[T any] struct {
	item []T
}

func (stack *Stack[T]) push(val T) {
	stack.item = append(stack.item, val)
}

func (stack *Stack[T]) pop() (T, error) {
	if len(stack.item) <= 0 {
		var zero T
		return zero, errors.New("Slice is empty")
	}

	lastItem := stack.item[len(stack.item)-1]
	stack.item = stack.item[0 : len(stack.item)-1]
	return lastItem, nil
}

func (stack *Stack[T]) peek() (T, error) {
	if len(stack.item) <= 0 {
		var zero T
		return zero, errors.New("Slice is empty")
	}

	return stack.item[len(stack.item)-1], nil
}

func main() {
	s := Stack[int]{}
}
