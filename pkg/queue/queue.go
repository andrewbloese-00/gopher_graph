package queue

import (
	"errors"
	"fmt"
)

type queueNode[T any] struct {
	Data T
	Next *queueNode[T]
}

type Queue[T any] struct {
	Size    int
	Front   *queueNode[T]
	Back    *queueNode[T]
	NONEVAL T
}

type QueueError struct {
	Err error
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Front: nil,
		Back:  nil,
		Size:  0,
	}
}

func (q *Queue[T]) Pop() (*T, error) {
	if q.Size == 0 {
		return nil, errors.New("Queue Underflow Error")
	}
	front := q.Front.Data
	q.Front = q.Front.Next
	q.Size -= 1
	return &front, nil
}

func (q *Queue[T]) Peek() (*T, error) {
	if q.Size == 0 {
		return nil, errors.New("Queue Underflow Error")
	}

	return &q.Front.Data, nil
}

func (q *Queue[T]) Push(data T) {
	if q.Front == nil {
		q.Front = &queueNode[T]{Data: data, Next: nil}
		q.Back = q.Front
	} else {
		node := &queueNode[T]{Data: data, Next: nil}
		q.Back.Next = node
		q.Back = q.Back.Next

	}
	q.Size += 1

}

func (q *Queue[T]) Print() {
	curr := q.Front
	fmt.Print("FRONT [ ")
	for curr != nil {
		fmt.Printf("%v -> ", curr.Data)
		curr = curr.Next
	}
	fmt.Print(" nil ] BACK\n")
}
