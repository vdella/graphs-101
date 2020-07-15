package graph

type Node struct {
	element   *int
	neighbors []Node
}

func (node Node) GetElement() *int {
	return node.element
}

func (node Node) GetNeighbors() []Node {
	return node.neighbors
}