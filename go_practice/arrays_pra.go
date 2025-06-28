package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main9() {
	fmt.Print("Enter numbers (e.g. [1, 2, 3]): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.Trim(input, "[]")

	numStrings := strings.Split(input, ",")

	for _, s := range numStrings {
		s = strings.TrimSpace(s)

		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Invalid number: %v\n", s)
			continue
		}

		if num%2 == 0 {
			fmt.Printf("The number %v is even\n", num)
		} else {
			fmt.Printf("The number %v is odd\n", num)
		}
	}
}
