package leetcode

import "math"

// https://leetcode.com/problems/minimum-cost-walk-in-weighted-graph/

type Edge struct {
	Weight int
	Target int
}

func MinimumCost(n int, edges [][]int, query [][]int) []int {
	graph := map[int][]Edge{}

	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], Edge{Weight: e[2], Target: e[1]})
		graph[e[1]] = append(graph[e[1]], Edge{Weight: e[2], Target: e[0]})
	}
	result := make([]int, len(query))

	for i, q := range query {
		if _, exists := graph[q[0]]; !exists {
			result[i] = -1
			continue
		}
		if _, exists := graph[q[1]]; !exists {
			result[i] = -1
			continue
		}
		visited := make(map[int]bool)
		result[i] = findPath(q[0], q[1], graph, math.MaxInt, visited)
	}

	return result

}

func minInt(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func minIntArr(arr []int) int {
	minVal := math.MaxInt
	for _, n := range arr {
		if minVal > n {
			minVal = n
		}
	}
	return minVal
}

func findPath(cur, target int, graph map[int][]Edge, cost int, visited map[int]bool) int {
	if cur == target {
		return cost
	}
	visited[cur] = true
	minCost := math.MaxInt

	for _, e := range graph[cur] {
		if visited[e.Target] {
			continue
		}
		newCost := findPath(e.Target, target, graph, cost&e.Weight, visited)
		minCost = minInt(minCost, newCost)
	}

	visited[cur] = false
	if minCost == math.MaxInt {
		return -1
	}
	return minCost
}
