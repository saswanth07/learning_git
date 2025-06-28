package main

import "os"

func main() {
	f, r := os.OpenFile("hello.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if r != nil {
		panic(r)
	}
	defer f.Close()
	if _, r := f.WriteString("cool buddy by someone\n"); r != nil {
		panic(r)
	}
}
