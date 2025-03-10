package node

type Node struct {
	Name        string
	SellsMango  bool
	Connections []*Node
}

type WNode struct {
	Name        string
	Connections map[*WNode]int
}
