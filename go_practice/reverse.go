package main

import "fmt"

func main3() {
	var s1 string
	fmt.Scan(&s1)
	fmt.Println(reverse(s1))
}
func reverse(s2 string) string {
	a1 := []rune(s2)
	for i, j := 0, len(a1)-1; i < j; i, j = i+1, j-1 {
		a1[i], a1[j] = a1[j], a1[i]
	}
	return string(a1)
}
