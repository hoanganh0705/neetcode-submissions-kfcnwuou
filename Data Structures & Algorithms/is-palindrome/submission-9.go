func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		// skip non-alphanumeric (left)
		for start < end && !isAlnum(rune(s[start])) {
			start++
		}

		// skip non-alphanumeric (right)
		for start < end && !isAlnum(rune(s[end])) {
			end--
		}

		// compare lowercase
		if toLower(s[start]) != toLower(s[end]) {
			return false
		}

		start++
		end--
	}

	return true
}

func isAlnum(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}
