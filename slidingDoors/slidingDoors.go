package slidingDoors

func NumberOfSubstrings(s string) int {
	strLen := len(s)
	result := 0
	for size := 3; size <= strLen; size++ {
		letters := map[byte]int{}
		for right := 0; right < strLen; right++ {
			chRight := s[right]
			if _, exists := letters[chRight]; exists {
				letters[chRight]++
			} else {
				letters[chRight] = 1
			}
			if right < size-1 {
				continue
			}
			if isValidSubString(letters) {
				result++
			}
			chLeft := s[right+1-size]
			letters[chLeft]--
		}
	}
	return result
}

func isValidSubString(letters map[byte]int) bool {
	if val, exists := letters['a']; !exists || val < 1 {
		return false
	}
	if val, exists := letters['b']; !exists || val < 1 {
		return false
	}
	if val, exists := letters['c']; !exists || val < 1 {
		return false
	}
	return true
}
