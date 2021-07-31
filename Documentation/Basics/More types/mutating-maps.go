package main

import "fmt"

func main() {
	m_mutate := make(map[string]int)

	m_mutate["Answer"] = 42
	fmt.Println("The value:", m_mutate["Answer"])

	m_mutate["Answer"] = 48
	fmt.Println("The value:", m_mutate["Answer"])

	delete(m_mutate, "Answer")
	fmt.Println("The value:", m_mutate["Answer"])

	v, ok := m_mutate["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
