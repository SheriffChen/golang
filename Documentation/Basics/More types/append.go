package main

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1, 2, 3)
	printSlice(s)
}
