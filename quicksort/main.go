//https://en.wikipedia.org/wiki/Quicksort
//Best O(n log n)
//Worst O(n^2)

// Unsorted [2 3 8 1 5 4 0 7 6 9]
// swap elements 2 - 0
// swap elements 3 - 1
// swap elements 8 - 6
// swap elements 6 - 2
// swap elements 5 - 4
// Sorted [0 1 2 3 4 5 6 7 8 9]
package main

import (
	"fmt"
)

func main() {
	unsort := []int{2, 3, 8, 1, 5, 4, 0, 7, 6, 9}
	fmt.Printf("Unsorted %d\n", unsort)

	quicksort(&unsort, 0, len(unsort)-1)
	fmt.Printf("Sorted %d\n", unsort)
}

func quicksort(arr *[]int, lo int, hi int) {
	if lo < hi {
		p := partion(arr, lo, hi)
		quicksort(arr, lo, p)
		quicksort(arr, p+1, hi)
	}
}

//Hoare partition scheme
func partion(arr *[]int, lo int, hi int) int {
	A := *arr
	pivot := A[lo]
	i := lo - 1
	j := hi + 1

	fmt.Printf("Start for \n")

	for {
		//do ... while
		for ok := true; ok; {
			i = i + 1
			fmt.Printf("%d < %d\n", A[i], pivot)
			ok = A[i] < pivot
		}

		for ok := true; ok; {
			j = j - 1
			fmt.Printf("%d > %d\n", A[j], pivot)
			ok = A[j] > pivot
		}

		fmt.Printf("i=%d, j=%d\n", i, j)

		if i >= j {
			return j
		}

		swap(&A[i], &A[j])
		fmt.Printf("Unsorted %d\n", A)
	}
}

func swap(l, r *int) {
	fmt.Printf("swap elements %d - %d\n", *l, *r)
	buf := *l
	*l = *r
	*r = buf
}
