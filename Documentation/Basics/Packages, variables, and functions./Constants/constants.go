package main

import "fmt"

//Pi is 3.14
const Pi = 3.14

func main() {
	const world = "世界"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
