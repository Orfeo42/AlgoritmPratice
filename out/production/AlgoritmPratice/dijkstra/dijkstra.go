package dijkstra

import "fmt"

func Test() {

	graph := map[string]map[string]int{}
	graph["start"] = map[string]int{"A": 5, "B": 2}
	graph["A"] = map[string]int{"C": 4, "D": 2}
	graph["B"] = map[string]int{"A": 8, "D": 7}
	graph["C"] = map[string]int{"fin": 3, "D": 6}
	graph["D"] = map[string]int{"fin": 1}
	graph["fin"] = map[string]int{}

	result := dijkstra(graph, "start", "fin")
	fmt.Println(result)

	graph = map[string]map[string]int{}
	graph["start"] = map[string]int{"A": 10}
	graph["A"] = map[string]int{"B": 20}
	graph["B"] = map[string]int{"fin": 30, "C": 1}
	graph["C"] = map[string]int{"A": 1}
	graph["fin"] = map[string]int{}
	result = dijkstra(graph, "start", "fin")
	fmt.Println(result)

}

func dijkstra(graph map[string]map[string]int, start string, target string) int {
	cost := map[string]int{}
	processed := map[string]bool{}
	parents := map[string]string{}

	for key, value := range graph[start] {
		cost[key] = value
	}

	node := unprocessedCheapestNode(cost, processed)
	for node != "" {
		for key, value := range graph[node] {
			newCost := cost[node] + value
			oldVal, exists := cost[key]
			if !exists || newCost < oldVal {
				cost[key] = newCost
				parents[key] = node
			}
		}
		processed[node] = true
		node = unprocessedCheapestNode(cost, processed)
	}
	return cost[target]
}

func unprocessedCheapestNode(cost map[string]int, processed map[string]bool) string {
	lowest := ""
	for key, value := range cost {
		if _, exists := processed[key]; exists {
			continue
		}
		if value == 0 {
			return key
		}
		if lowest == "" {
			lowest = key
			continue
		}
		if cost[lowest] > value {
			lowest = key
		}
	}
	return lowest
}
