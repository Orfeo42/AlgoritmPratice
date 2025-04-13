package leetcode

//https://leetcode.com/problems/count-good-numbers

const mod = 1_000_000_007

func countGoodNumbers(n int64) int {
	evenCount := (n + 1) / 2
	oddCount := n / 2

	evenPart := modPow(5, evenCount, mod)
	oddPart := modPow(4, oddCount, mod)

	return int(evenPart * oddPart % mod)
}

func modPow(base, exp, mod int64) int64 {
	result := int64(1)
	base %= mod

	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp /= 2
	}

	return result
}
