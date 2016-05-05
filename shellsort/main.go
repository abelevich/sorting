//https://en.wikipedia.org/wiki/Shellsort
//Best O(n log2 n)
//Worst O(nlog2 2n
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
    n := len(arr)
    
    //Using Marcin Ciura's gap sequence. Good up to 4000 elements
    gaps := []int{701, 301, 132, 57, 23, 10, 4, 1}
    
    for k:=0; k < len(gaps); k++ {
        gap := gaps[k]

        for i := gap; i < n; i++  {
            fmt.Printf("Gap %d\n", gap)
            var j int
            temp := arr[i]
            for j = i; j >= gap && arr[j - gap] > temp; j = j - gap {
                fmt.Printf("i %d j %d\n",i, j)
                fmt.Printf("move %d to %d\n",arr[j - gap], j)
                arr[j] = arr[j - gap]
                fmt.Printf("Unsorted %d\n", unsort)
            }
            
            arr[j] = temp
            fmt.Printf("Unsorted %d\n", unsort)
        }
    }
}

