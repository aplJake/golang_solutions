package solutions

import (
	"fmt"
	"sort"
)

/*
The two oldest ages function/method needs to be completed.
It should take an array of numbers as its argument and return
the two highest numbers within the array. The returned value
should be an array in the format [second oldest age, oldest age].
*/

type (
	//Person struct {
	//	Age int
	//}

	ByAge []int
)

func (x ByAge) Len() int {
	return len(x)
}

func (x ByAge) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x ByAge) Less(i, j int) bool  {
	return x[i] < x[j]
}

func ShowTwoOldestAges(ages []int) {
	fmt.Println("Initial array:", ages)

	// initially sort array in ascending order
	sort.Sort(ByAge(ages))

	// return the last two items
	fmt.Println("Two biggest items from array:", ages[len(ages)- 2:])
}
