package tests

import (
	"fmt"

	"github.com/andrewbloese-00/gopher_graph/pkg/queue"
)

// tests an integer queue by Pushing 'n' elements then popping and peeking until it is empty.
func TestQueue(n int) {
	fmt.Println("== Test Queue[int] ==")
	int_queue := queue.NewQueue[int]()
	for i := n; i > 0; i -= 1 {
		fmt.Println("Pushing: ", i)
		int_queue.Push(i)
		int_queue.Print()
	}
	int_queue.Print()
	for int_queue.Size > 0 {
		popPtr, err := int_queue.Pop()
		if err != nil {
			panic(err)
		}
		peekPtr, err := int_queue.Peek()
		fmt.Printf("Popped: %d \n", *popPtr)
		if err == nil {
			fmt.Printf("Peeked: %d \n", *peekPtr)
		} else { //only happens when queue is empty (last line in test output)
			fmt.Printf("Peeked: %v", err)
		}

	}

}
