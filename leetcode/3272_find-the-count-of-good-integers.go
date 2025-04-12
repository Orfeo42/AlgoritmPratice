package leetcode

import (
	"math"
	"sort"
	"strconv"
)

//https://leetcode.com/problems/find-the-count-of-good-integers/

func countGoodIntegers(n, k int) int64 {
	halfLen := (n + 1) / 2
	start := int(math.Pow(10, float64(halfLen-1)))
	end := int(math.Pow(10, float64(halfLen)))
	if halfLen == 1 {
		start = 1
	}

	var totale int64 = 0
	permMemo := make(map[string]bool)
	fattMemo := make(map[int]int64)

	for i := start; i < end; i++ {
		firstHalf := strconv.Itoa(i)
		secondHalf := reverseString(firstHalf[:n/2])
		palindromo := firstHalf + secondHalf

		num, _ := strconv.ParseInt(palindromo, 10, 64)
		if num%int64(k) == 0 {
			sorted := []rune(palindromo)
			sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
			sortedStr := string(sorted)

			if permMemo[sortedStr] {
				continue
			}
			permMemo[sortedStr] = true
			totale += permutazioniValidi(palindromo, fattMemo)
		}
	}

	return totale
}

func permutazioniValidi(numero string, fattMemo map[int]int64) int64 {
	freq := make([]int, 10)
	n := len(numero)

	for _, c := range numero {
		freq[c-'0']++
	}

	var totale int64 = 0

	for prima := 1; prima <= 9; prima++ {
		if freq[prima] == 0 {
			continue
		}

		freq[prima]--

		perm := fattoriale(n-1, fattMemo)
		for i := 0; i <= 9; i++ {
			perm /= fattoriale(freq[i], fattMemo)
		}

		totale += perm
		freq[prima]++
	}

	return totale
}

func fattoriale(n int, memo map[int]int64) int64 {
	if n <= 1 {
		return 1
	}
	if val, ok := memo[n]; ok {
		return val
	}
	res := int64(n) * fattoriale(n-1, memo)
	memo[n] = res
	return res
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
