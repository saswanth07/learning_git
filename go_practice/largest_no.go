package main

import (
	"fmt"
	"sort"
)

func main5() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	sort.Ints(arr)
	var n1 int
	for j := n - 1; j > 1; j-- {
		if arr[j] != arr[j-1] {
			n1 = arr[j-1]
			break
		}

	}
	fmt.Println(n1)
}
