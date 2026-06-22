type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	result := ""

	for _, str := range strs {
		result += strconv.Itoa(len(str)) + "#" + str
	}

	return result
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}
	i := 0

	for i < len(encoded) {
		j := i

		for encoded[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])

		word := encoded[j+1 : j+1+length]

		res = append(res, word)

		i = j + 1 + length
	}

	return res
}