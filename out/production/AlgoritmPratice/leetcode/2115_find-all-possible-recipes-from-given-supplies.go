package leetcode

//https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/

func FindAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	recMap := map[string]int{}
	supMap := map[string]bool{}
	visited := make(map[string]int)
	for _, s := range supplies {
		supMap[s] = true
	}
	for i, s := range recipes {
		recMap[s] = i
	}
	var makeRecipe func(ing string) bool
	makeRecipe = func(ing string) bool {
		if val, exists := visited[ing]; exists {
			return val == 1
		}
		if _, exists := supMap[ing]; exists {
			return true
		}
		if _, exists := recMap[ing]; !exists {
			return false
		}
		visited[ing] = 0
		for _, ingredient := range ingredients[recMap[ing]] {
			if !makeRecipe(ingredient) {
				visited[ing] = -1
				return false
			}
		}
		visited[ing] = 1
		supMap[ing] = true
		return true
	}

	var res []string
	for _, rec := range recipes {
		if makeRecipe(rec) {
			res = append(res, rec)
		}
	}
	return res
}
