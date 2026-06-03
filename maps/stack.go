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

	fmt.Println(isAnagram("anagram", "nagaram")) // true
	fmt.Println(isAnagram("rat", "car"))         // false
	fmt.Println(isAnagram("listen", "silent"))   //true
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
	allowed := map[byte]bool{
		'(': true,
		')': true,
		'[': true,
		']': true,
		'{': true,
		'}': true,
	}

	for i := 0; i < len(s); i++ {
		ch := s[i]

		_, exists := allowed[ch]
		if !exists {
			return false
		}

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

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := map[byte]int{}

	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		count[t[i]]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}
