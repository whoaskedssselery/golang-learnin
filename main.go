
package main

import (
	"fmt"
	"math"
)

func bubblesort(a []float64) []float64 {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if math.Abs(a[j]) > math.Abs(a[j+1]) {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func main() {
	a := []float64{1, 2, 5, -2, 4, 10, 0}
	sortedArray := bubblesort(a)
	fmt.Println(sortedArray)
}
