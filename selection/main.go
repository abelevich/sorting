// Selection sort O(n^2)
//https://en.wikipedia.org/wiki/Selection_sort

// Start Selection sorting
// Unsorted: [0 3 8 1 5 4 2 7 6 9]
// Unsorted: [0 1 8 3 5 4 2 7 6 9]
// Unsorted: [0 1 2 3 5 4 8 7 6 9]
// Unsorted: [0 1 2 3 5 4 8 7 6 9]
// Unsorted: [0 1 2 3 4 5 8 7 6 9]
// Unsorted: [0 1 2 3 4 5 8 7 6 9]
// Unsorted: [0 1 2 3 4 5 6 7 8 9]
// Unsorted: [0 1 2 3 4 5 6 7 8 9]
// Unsorted: [0 1 2 3 4 5 6 7 8 9]
// Sorted: [0 1 2 3 4 5 6 7 8 9]

package main

import (
	"fmt"
)

func main() {
	unsort := []int{2, 3, 8, 1, 5, 4, 0, 7, 6, 9}
	fmt.Printf("Unsorted %d\n", unsort)
	fmt.Printf("%s\n", "Start Selection sorting")

	for i := 0; i < len(unsort)-1; i++ {
		min := i
		for j := i; j < len(unsort); j++ {
			if unsort[j] < unsort[min] {
				min = j
			}
		}

		if i != min {
			swap(&unsort[i], &unsort[min])
		}

		fmt.Printf("Unsorted: %d\n", unsort)
	}

	fmt.Printf("Sorted: %d\n", unsort)
}

func swap(l, r *int) {
	buf := *l
	*l = *r
	*r = buf
}
