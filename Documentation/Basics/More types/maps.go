package main

import "fmt"

type Vertex_maps struct {
	Lat, Long float64
}

var m map[string]Vertex_maps

func main() {
	m = make(map[string]Vertex_maps)
	m["Chen XiaoFan"] = Vertex_maps{
		13, 69,
	}
	fmt.Println(m["Chen XiaoFan"])
}
