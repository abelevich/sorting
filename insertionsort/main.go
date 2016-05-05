// Insertion sort O(n^2)
// https://en.wikipedia.org/wiki/Insertion_sort

//Step outer: 1
//Step outer: 2
//Sorted: [2 3 1 8 5 4 0 7 6 9]
//Sorted: [2 1 3 8 5 4 0 7 6 9]
//Sorted: [1 2 3 8 5 4 0 7 6 9]
//Step outer: 3
//Sorted: [1 2 3 5 8 4 0 7 6 9]
//Step outer: 4
// Sorted: [1 2 3 5 4 8 0 7 6 9]
// Sorted: [1 2 3 4 5 8 0 7 6 9]
// Step outer: 5
// Sorted: [1 2 3 4 5 0 8 7 6 9]
// Sorted: [1 2 3 4 0 5 8 7 6 9]
// Sorted: [1 2 3 0 4 5 8 7 6 9]
// Sorted: [1 2 0 3 4 5 8 7 6 9]
// Sorted: [1 0 2 3 4 5 8 7 6 9]
// Sorted: [0 1 2 3 4 5 8 7 6 9]
// Step outer: 6
// Sorted: [0 1 2 3 4 5 7 8 6 9]
// Step outer: 7
// Sorted: [0 1 2 3 4 5 7 6 8 9]
// Sorted: [0 1 2 3 4 5 6 7 8 9]
// Step outer: 8
// Sorted: [0 1 2 3 4 5 6 7 8 9]

package main

import (
	"fmt"
)

func main() {
	unsort := []int{2, 3, 8, 1, 5, 4, 0, 7, 6, 9}
	fmt.Printf("Unsorted %d\n", unsort)
	fmt.Printf("%s\n", "Start Insertion sorting")

	for i := 1; i < len(unsort)-1; i++ {
		j := i
		for j > 0 && unsort[j-1] > unsort[j] {
			swap(&unsort, j-1, j)
			j--
		}
		fmt.Printf("Step outer: %d\n", i)
	}

	fmt.Printf("Sorted: %d\n", unsort)
}

func swap(slice *[]int, l, r int) {
	sliceValue := *slice
	buf := sliceValue[l]
	sliceValue[l] = sliceValue[r]
	sliceValue[r] = buf
	fmt.Printf("Sorted: %d\n", sliceValue)
}

