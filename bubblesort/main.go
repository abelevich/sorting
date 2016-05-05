//https://en.wikipedia.org/wiki/Bubble_sort
//Best O(n)
//Worst O(n^2)

package main

import (
	"fmt"
)

func main() {
	unsort := []int{2, 3, 8, 1, 5, 4, 10, 7, 9, 6, 0}

	fmt.Printf("Unsorted %d\n", unsort)
	sort(&unsort)
	fmt.Printf("Sorted %d\n", unsort)
}

func sort(unsort *[]int) {
	arr := *unsort
	swapped := true
	n := len(arr)

	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				swap(&arr[i-1], &arr[i])
				swapped = true
			}
		}
		n = n - 1
	}
}

func swap(l, r *int) {
	fmt.Printf("swap elements %d - %d\n", *l, *r)
	buf := *l
	*l = *r
	*r = buf
}
