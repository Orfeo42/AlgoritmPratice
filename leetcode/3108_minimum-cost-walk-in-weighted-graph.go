package leetcode

//https://leetcode.com/problems/minimum-cost-walk-in-weighted-graph/

type Edge struct {
	Weight int
	Target int
}

func MinimumCost(n int, edges [][]int, query [][]int) []int {
	graph := map[int][]Edge{}

	for _, e := range edges {
		if _, exists := graph[e[0]]; !exists {
			graph[e[0]] = []Edge{}
		}
		graph[e[0]] = append(graph[e[0]], Edge{
			Weight: e[2],
			Target: e[1],
		})

		if _, exists := graph[e[1]]; !exists {
			graph[e[1]] = []Edge{}
		}
		graph[e[1]] = append(graph[e[1]], Edge{
			Weight: e[2],
			Target: e[0],
		})
	}
	var result []int

	for _, q := range query {
		if _, exists := graph[q[0]]; !exists {
			result = append(result, -1)
			continue
		}
		if _, exists := graph[q[1]]; !exists {
			result = append(result, -1)
			continue
		}
		/*links := graph[q[0]]
		minLength := math.MaxInt
		for _, l := range links {
			length := math.MaxInt

		}*/

		result = append(result, 1)
	}

	return result

}
