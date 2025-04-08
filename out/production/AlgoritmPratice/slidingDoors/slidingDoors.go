package slidingDoors

func NumberOfSubstrings(s string) int {
	strLen := len(s)
	result := 0
	for size := 3; size <= strLen; size++ {
		letters := [3]int{}
		for right := 0; right < strLen; right++ {
			chRight := s[right]
			letters[chRight-'a']++
			if right < size-1 {
				continue
			}
			if isValidSubString(letters) {
				result++
			}
			chLeft := s[right+1-size]
			letters[chLeft-'a']--
		}
	}
	return result
}

func isValidSubString(letters [3]int) bool {
	if letters[0] > 0 && letters[1] > 0 && letters[2] > 0 {
		return true
	}
	return false
}
