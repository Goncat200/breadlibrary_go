package testMain

import (
	"fmt"
)

func TestMain() {
	var len_array int
	fmt.Scan(&len_array)
	var array []int = make([]int, len_array)
	var a int
	for i := 0; i < len(array); i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	for i := 0; i < len(array); i++ {
		if (i % 2) == 0 {
			fmt.Printf("%d ", array[i])
		}
	}
}