package graph

type Graph struct {
	vertices []Node
}

func (graph Graph) Vertices() []Node {
	return graph.vertices
}

func NewEmptyGraph() *Graph {
	return &Graph{}
}

func (graph Graph) SetVertices(size int) {
	graph.vertices = make([]graph.Node, size)
}

func NewGraph(edges []Node) *Graph {
	return &Graph{edges}
}
