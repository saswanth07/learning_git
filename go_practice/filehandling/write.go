package main

import "os"

func main() {
	content := "add this string to the file hello"
	r := os.WriteFile("hello.txt", []byte(content), 0644)
	if r != nil {
		panic(r)
	}
}
