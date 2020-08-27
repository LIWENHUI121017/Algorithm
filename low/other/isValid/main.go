package main

import (
	"fmt"
)

type stack struct {
	index int
	data  []byte
}

func (s *stack) push(value byte) {
	s.data[s.index] = value
	s.index++
}

func (s *stack) empty() bool {
	return s.index == 0
}

func (s *stack) pop() byte {
	if s.empty() {
		panic("stack is empty")
	}
	s.index--
	return s.data[s.index]
}

//func isValid(s string) bool {
//	//用一个栈来存左括号，然后循环遇到右括号的时候，就出栈，如果没匹配的上，就返回false
//	data := make([]byte,len(s))
//	st := &stack{0,data}
//	for i := 0; i < len(s); i++ {
//		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
//			st.push(s[i])
//		} else {
//			if st.empty() {
//				return false
//			}
//			a := st.pop()
//			res := check(s[i])
//			if a!=res {
//				return false
//			}
//		}
//	}
//	if st.index != 0 {
//		return false
//	}
//	return true
//}
func isValid(s string) bool {
	n := len(s)
	//如果是单数，说明肯定是错的
	if n%2 == 1 {
		return false
	}
	//用一个栈来存左括号，然后循环遇到右括号的时候，就出栈，如果没匹配的上，就返回false
	check := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != check[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func check(s byte) byte {
	switch s {
	case ')':
		return '('
	case '}':
		return '{'
	case ']':
		return '['
	default:
		return '0'
	}
}
func main() {
	//15的二进制为1111
	res := isValid("([])")
	fmt.Print(res)
}
