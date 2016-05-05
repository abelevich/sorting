//heapsort implementation (bottom-up)
//Best O(n log n)
//Worst O(n log n)
//https://en.wikipedia.org/wiki/Heapsort

package main

import (
	"fmt"
	"math"
)

func main() {
	unsort := []int{2, 3, 8, 1, 5, 4, 0, 7, 6, 9} //10

	fmt.Printf("Unsort %d\n", unsort)
	fmt.Println("Start sorting")

	heapify(&unsort)

	fmt.Printf("heapify  %d\n\n", unsort)

	for end := len(unsort) - 1; end > 0; {
		swap(&unsort[end], &unsort[0])

		end--

		buildHeap(&unsort, 0, end)
		fmt.Printf("rebuildHeap  %d\n\n", unsort)
	}

	fmt.Printf("Sorted %d\n", unsort)
}

func heapify(unsort *[]int) {
	array := *unsort
	count := len(array) - 1

	//start is a parent of the last element in the tree
	for start := iParent(count); start >= 0; start-- {
		buildHeap(unsort, start, count)
	}
}

func buildHeap(unsort *[]int, start, end int) {
	array := *unsort
	root := start

	fmt.Printf("Parent %d\n", root)

	for iLeftChild(root) <= end {
		child := iLeftChild(root)
		iSwap := root

		if array[iSwap] < array[child] {
			iSwap = child
		}

		if child+1 <= end && array[iSwap] < array[child+1] {
			iSwap = child + 1
		}

		if iSwap != root {
			swap(&array[root], &array[iSwap])
			root = iSwap
		} else {
			return
		}
	}
}

func swap(l, r *int) {
	fmt.Printf("swap elements %d - %d\n", *l, *r)
	buf := *l
	*l = *r
	*r = buf
}

func iParent(i int) int {
	return int(math.Floor(float64((i - 1) / 2)))
}

func iLeftChild(i int) int {
	return 2*i + 1
}
