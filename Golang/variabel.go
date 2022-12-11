package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Println("My Age Is", age)
	age = 29
	fmt.Println("My Age Is", age)
	age = 50
	fmt.Println("My Age Is", age)

	var width, height int
	fmt.Println("Width is", width, "Height Is", height)
	width = 300
	height = 200
	fmt.Println("New Width Is", width, "New Height Is", height)

	name := "Sean Michael"
	fmt.Println("My name is", name)
}
