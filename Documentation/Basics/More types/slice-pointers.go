package main

import "fmt"

func main() {
	names := [4]string{
		"Zhao",
		"Qian",
		"Sun",
		"Li",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "Chen"
	fmt.Println(a, b)
	fmt.Println(names)
}
