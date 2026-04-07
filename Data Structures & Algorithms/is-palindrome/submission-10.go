func isPalindrome(s string) bool {
	start, end := 0, len(s)-1

	for start < end {
		// skip non-alphanumeric
		for start < end && !isAlnum(s[start]) {
			start++
		}
		for start < end && !isAlnum(s[end]) {
			end--
		}

		// compare (case-insensitive)
		if toLower(s[start]) != toLower(s[end]) {
			return false
		}

		start++
		end--
	}

	return true
}

func isAlnum(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}