package breadthFirst

import (
	"fmt"
	"test-go/node"
	"test-go/queue"
)

func breadthFirstMango(start node.Node) (node.Node, bool) {
	var q queue.Queue[node.Node]
	q.Push(start)
	for !q.IsEmpty() {
		current, _ := q.Pop()
		if current.SellsMango {
			return current, true
		}
		for _, n := range current.Connections {
			q.Push(*n)
		}
	}

	var void node.Node
	return void, false
}

func Test() {
	paperino := node.Node{
		Name:       "paperino",
		SellsMango: false,
	}
	braccoBaldo := node.Node{
		Name:       "braccoBaldo",
		SellsMango: true,
	}
	pluto := node.Node{
		Name:        "pluto",
		SellsMango:  true,
		Connections: []*node.Node{&braccoBaldo},
	}
	pippo := node.Node{
		Name:        "pippo",
		SellsMango:  false,
		Connections: []*node.Node{&pluto, &paperino},
	}
	result, _ := breadthFirstMango(pippo)
	fmt.Printf("%+v\n", result)
}
