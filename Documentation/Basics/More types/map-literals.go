package main

import "fmt"

type Vertex_map_literals struct {
	Lat, Long float64
}

var m_literals = map[string]Vertex_map_literals{
	"Chen": {
		13, 69,
	},
	"Xiao": {
		22, 33,
	},
}

func main() {
	fmt.Println(m_literals)
}
