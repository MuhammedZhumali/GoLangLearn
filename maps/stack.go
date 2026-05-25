package main

import "fmt"

func main() {
	stack := []string{}

	stack = append(stack, "first")
	stack = append(stack, "second")
	stack = append(stack, "third")
	top := stack[len(stack)-1]
	fmt.Println(top)
	stack = stack[:len(stack)-1]
	fmt.Println(stack)

	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid("{[]}"))   // true

	fmt.Println(isValidMap("()"))     // true
	fmt.Println(isValidMap("()[]{}")) // true
	fmt.Println(isValidMap("(]"))     // false
	fmt.Println(isValidMap("([)]"))   // false
	fmt.Println(isValidMap("{[]}"))   // true
	fmt.Println(isValidMap("("))      // false
	fmt.Println(isValidMap("]"))      // false
}

func isValid(s string) bool {
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]

			if ch == ')' && top != '(' {
				return false
			}

			if ch == ']' && top != '[' {
				return false
			}

			if ch == '}' && top != '{' {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func isValidMap(s string) bool {
	stack := []byte{}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		ch := s[i]

		open, isClosing := pairs[ch]

		if isClosing {

			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]

			if top != open {
				return false
			}

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}
