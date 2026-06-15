func groupAnagrams(strs []string) [][]string {
	group := make(group[[26]int][]string)

	for _, word := range strs {
		var count [26] int
		for char := range word {
			count[char-'a']++
		}

		group[count] = append(group[count], word)
	}

	result = make([][]string, 0, len(group))

	for _, group := range groups {
    	result = append(result, group)
    }

	return result
}
