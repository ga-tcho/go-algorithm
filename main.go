package main

import (
	"fmt"
)

func main() {
	a := []int{2, 4, 5, 1, 3}

	fmt.Println(BubbleSort(a))
	fmt.Println(SelectionSort(a))
}
