//Merge sort (bottom-up implementation)

//https://en.wikipedia.org/wiki/Merge_sort
//Best O(n log n)
//Worst O(n log n)
package main

import (
	"fmt"
)

func main() {
	unsort := []int{2, 3, 8, 1, 5, 4, 10, 7, 9, 6, 0}

	fmt.Printf("Unsorted %d\n", unsort)
	fmt.Printf("%s\n", "Start Merge sorting")
	sort(unsort)
}

func sort(unsort []int) {
	length := len(unsort) //10
	sort := make([]int, length)

	for width := 1; width <= length; width = 2 * width {
		for i := 0; i < length; i = i + 2*width {
			merge(unsort, i, min(i+width, length), min(i+2*width, length), &sort)
			fmt.Printf("Unsorted: %d\n", sort)

		}
		copy(unsort, sort)
	}

	fmt.Printf("Sorted: %d\n", sort)
}

func merge(unsort []int, begin, middle, end int, sort *[]int) {
	i := begin
	j := middle
	buf := *sort

	for k := begin; k < end; k++ {
		if i < middle && (j >= end || unsort[i] <= unsort[j]) {
			buf[k] = unsort[i]
			i++
		} else {
			buf[k] = unsort[j]
			j++
		}
	}
}

func min(a, b int) int {
	var result int

	if a < b {
		result = a
	} else {
		result = b
	}

	return result
}
