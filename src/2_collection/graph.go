package main

import (
	"bytes"
	"fmt"
)

type Graph struct {
	data map[string]map[string]int
}

func NewGraph() *Graph {
	g := &Graph{}
	g.data = make(map[string]map[string]int)
	return g
}

// оройн жагсаалт
func (g *Graph) V() []string {
	v := make([]string, 0)
	for k, _ := range g.data {
		v = append(v, k)
	}
	return v
}

// ирмэг нэмэх
func (g *Graph) addEdge(source, target string, weight int) {
	map2, ok := g.data[source]
	if !ok {
		map2 = make(map[string]int)
	}
	map2[target] = weight
	g.data[source] = map2
}

func (g *Graph) String() string {
	var strBuf bytes.Buffer
	for k, map2 := range g.data {
		strBuf.WriteString(k + "\n")
		for k2, w := range map2 {
			strBuf.WriteString(fmt.Sprintf("  -> %s -> %d \n", k2, w))
		}
	}
	return strBuf.String()
}

func main() {
	g := NewGraph()
	g.addEdge("A", "B", 2)
	g.addEdge("A", "C", 0)
	g.addEdge("A", "G", 3)
	g.addEdge("A", "H", 5)
	g.addEdge("B", "A", 2)
	g.addEdge("C", "A", 0)
	g.addEdge("G", "A", 3)
	g.addEdge("G", "C", 0)
	g.addEdge("H", "A", 5)

	fmt.Println(g.V())
}
