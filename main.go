package main

func isValid(s string) bool {
	// initialize stack
	stack := make([]byte, 0)

	mappings := make(map[byte]byte, 3)
	mappings[')'] = '('
	mappings['}'] = '{'
	mappings[']'] = '['

	for i := 0; i < len(s); i++ {
		// character
		char := s[i]
		// is it a closing character?
		if mappings[char] != 0 {

			// if stack is empty
			// mistmatch
			if len(stack) == 0 {
				return false
			}

			// Get the top
			// Pop() it off
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// does the top match, the character?
			if top != mappings[char] {
				return false
			}
		} else {
			stack = append(stack, char)
		}

	}
	return len(stack) == 0
}
