package main

import "fmt"

func main() {
	var s []int = []int{}
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	} else {
		fmt.Println("not nil")
	}
}
