package main

import (
	"fmt"
	"sort"
)

// how to sort an array in golang
func SortArray(arr []int) {
	sort.Ints(arr)
}

// BinarySearch performs binary search on a sorted array
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// BubbleSort performs bubble sort on an array
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// Stack represents a stack data structure
type Stack struct {
	items []int
}

// Push adds an element to the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1 // Return -1 if stack is empty
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Queue represents a queue data structure
type Queue struct {
	items []int
}

// Enqueue adds an element to the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front element of the queue
func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1 // Return -1 if queue is empty
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func main() {
	// Binary Search
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7
	result := BinarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Element found at index %d\n", result)
	} else {
		fmt.Println("Element not found in the array")
	}

	// Bubble Sort
	arrToSort := []int{64, 34, 25, 12, 22, 11, 90}
	BubbleSort(arrToSort)
	fmt.Println("Sorted array:", arrToSort)

	// Stack operations
	stack := Stack{}
	stack.Push(10)
	stack.Push(20)
	fmt.Println("Stack Pop:", stack.Pop())
	fmt.Println("Stack Pop:", stack.Pop())

	// Queue operations
	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	fmt.Println("Queue Dequeue:", queue.Dequeue())
	fmt.Println("Queue Dequeue:", queue.Dequeue())
}
