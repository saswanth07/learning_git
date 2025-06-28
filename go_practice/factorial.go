package main

import "fmt"

func main6() {
	var num int
	fmt.Scan(&num)
	fmt.Printf("factorial of number %d", factorial(num))
}
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
