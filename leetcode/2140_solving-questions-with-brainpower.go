package leetcode

//https://leetcode.com/problems/solving-questions-with-brainpower

func mostPoints(questions [][]int) int64 {
	l := len(questions)
	points := make([]int64, l)
	for i := l - 1; i > -1; i-- {
		points[i] = int64(questions[i][0])
		nextT := questions[i][1] + i + 1
		nextS := i + 1
		if nextT < l {
			points[i] += points[nextT]
		}
		if nextS < l {
			points[i] = maxInt64(points[i], points[nextS])
		}
	}

	return points[0]
}

func maxInt64(n1, n2 int64) int64 {
	if n1 > n2 {
		return n1
	}
	return n2
}
