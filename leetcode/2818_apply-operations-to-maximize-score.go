package leetcode

import "sort"

// https://leetcode.com/problems/apply-operations-to-maximize-score

const module = 1000000007

func maximumScore(nums []int, k int) int {
	n := len(nums)
	scores := make([]int, n)

	mapPrime := map[int]int{}
	for i, n := range nums {
		scores[i] = getPrime(n, mapPrime)
	}

	left := make([]int, n)
	var stack []int
	for i := 0; i < len(scores); i++ {
		for len(stack) > 0 && scores[i] > scores[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = 0
		} else {
			left[i] = stack[len(stack)-1] + 1
		}
		stack = append(stack, i)
	}

	right := make([]int, n)
	stack = []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && scores[i] >= scores[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = len(nums) - 1
		} else {
			right[i] = stack[len(stack)-1] - 1
		}
		stack = append(stack, i)
	}

	queue := make([]int, n)
	for i := range queue {
		queue[i] = i
	}
	sort.Slice(queue, func(i, j int) bool {
		return nums[queue[i]] > nums[queue[j]] || (nums[queue[i]] == nums[queue[j]] && queue[i] < queue[j])
	})

	response := 1
	qIdx := 0
	for k > 0 {
		i := queue[qIdx]
		qIdx++
		leftSize := i - left[i]
		rightSize := right[i] - i
		count := min((leftSize+1)*(rightSize+1), k)
		k -= count

		tmpRes := pow(nums[i], count)
		response *= tmpRes
		response %= module
	}
	return response
}

func pow(num, k int) int {
	if k == 1 {
		return num
	}
	if k%2 == 0 {
		response := pow(num, k/2)
		return response * response % module
	} else {
		return num * pow(num, k-1) % module
	}
}

func getPrime(n int, mapPrime map[int]int) int {
	k := n
	response, ok := mapPrime[k]
	if ok {
		return response
	}
	i0 := 2
	for i0 <= n/i0 {
		if n%i0 == 0 {
			response++
		}
		for n%i0 == 0 {
			n /= i0
		}
		i0++
	}
	if n != 1 {
		response++
	}
	mapPrime[k] = response
	return response
}
