func isValid(s string) bool {
    stack := []byte{}

    // mapping closing -> opening
    pairs := map[byte]byte{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for i := 0; i < len(s); i++ {
        char := s[i]

        // nếu là closing bracket
        if open, ok := pairs[char]; ok {
            if len(stack) == 0 {
                return false
            }

            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            if top != open {
                return false
            }
        } else {
            // opening bracket → push
            stack = append(stack, char)
        }
    }

    return len(stack) == 0
}