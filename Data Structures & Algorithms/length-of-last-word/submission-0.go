func lengthOfLastWord(s string) int {
	lenngth:= 0
	for i:= len(s) - 1 ; i >= 0; i -- {
		if lenngth >= 1 && s[i] == ' ' {
			break;
		}

		if (s[i] != ' ') {
			lenngth++
		}
	}

	return lenngth
}
