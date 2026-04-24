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

type Queue[T any] struct {
	item []T
}

func (queue *Queue[T]) enqueue(val T) {
	queue.item = append(queue.item, val)
}

func (queue *Queue[T]) dequeue() (T, error) {
	if len(queue.item) <= 0 {
		var zero T
		return zero, errors.New("Queue is empty")
	}

	itemToRemove := queue.item[0]
	queue.item = queue.item[1:]
	return itemToRemove, nil
}

func (queue *Queue[T]) peekQueue() (T, error) {
	if len(queue.item) <= 0 {
		var zero T
		return zero, errors.New("Queue is empty")
	}

	return queue.item[0], nil
}

func main() {
	s := Stack[int]{}
}
