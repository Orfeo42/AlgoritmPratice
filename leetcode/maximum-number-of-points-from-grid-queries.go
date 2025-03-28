pakage leetcode

type Cell struct {
	val, x, y int
}

type MinHeap []Cell

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Cell))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxPoints(grid [][]int, queries []int) []int {
	m, n := len(grid), len(grid[0])
	result := make([]int, len(queries))

	qIdx := make([][2]int, len(queries))
	for i, q := range queries {
		qIdx[i] = [2]int{q, i}
	}
	sort.Slice(qIdx, func(i, j int) bool { return qIdx[i][0] < qIdx[j][0] })

	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, Cell{grid[0][0], 0, 0})

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	visited[0][0] = true

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	count, idx := 0, 0

	for idx < len(qIdx) {
		query, resIdx := qIdx[idx][0], qIdx[idx][1]

		for h.Len() > 0 && (*h)[0].val < query {
			cell := heap.Pop(h).(Cell)
			count++

			for _, d := range dirs {
				x, y := cell.x+d[0], cell.y+d[1]
				if x >= 0 && x < m && y >= 0 && y < n && !visited[x][y] {
					heap.Push(h, Cell{grid[x][y], x, y})
					visited[x][y] = true
				}
			}
		}
		result[resIdx] = count
		idx++
	}

	return result
}
