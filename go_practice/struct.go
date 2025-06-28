package main

import "fmt"

type person struct {
	name   string
	age    int
	height float64
	weight float64
	degree string
	class  string
}

func main1() {
	a := person{
		name:   "praveen",
		age:    40,
		height: 5.5,
		weight: 40,
		degree: "EEE",
		class:  "ThirdYear",
	}
	fmt.Println(a.name)
	fmt.Println(a.age)
	fmt.Println(a.degree)
	fmt.Println(a.class)

}
