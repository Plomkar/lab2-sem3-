package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"project/stack"
)

func prec(c byte) int {
	switch c {
	case '^':
		return 3
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

func isOp(c byte) bool {
	return c == '+' || c == '-' || c == '*' || c == '/' || c == '^'
}

func isAlnum(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

// InfixToPostfix — алгоритм сортировочной станции.
func InfixToPostfix(s string, result *string) bool {
	var st stack.Stack
	st.Init()
	*result = ""
	expectOperand := true
	ok := true

	for i := 0; i < len(s) && ok; {
		c := s[i]
		if isSpace(c) {
			i++
			continue
		}
		if isAlnum(c) || c == '.' {
			if !expectOperand {
				ok = false
				break
			}
			var tok strings.Builder
			for i < len(s) && (isAlnum(s[i]) || s[i] == '.') {
				tok.WriteByte(s[i])
				i++
			}
			*result += tok.String() + " "
			expectOperand = false
		} else if c == '(' {
			if !expectOperand {
				ok = false
				break
			}
			st.Push(int(c))
			i++
		} else if c == ')' {
			if expectOperand {
				ok = false
				break
			}
			for !st.IsEmpty() && st.Peek() != '(' {
				*result += string(byte(st.Pop())) + " "
			}
			if st.IsEmpty() {
				ok = false
				break
			}
			st.Pop()
			i++
		} else if isOp(c) {
			if expectOperand {
				ok = false
				break
			}
			right := c == '^'
			for !st.IsEmpty() && st.Peek() != '(' &&
				(prec(byte(st.Peek())) > prec(c) ||
					(prec(byte(st.Peek())) == prec(c) && !right)) {
				*result += string(byte(st.Pop())) + " "
			}
			st.Push(int(c))
			expectOperand = true
			i++
		} else {
			ok = false
		}
	}

	if ok && expectOperand {
		ok = false
	}
	for ok && !st.IsEmpty() {
		c := byte(st.Pop())
		if c == '(' {
			ok = false
			break
		}
		*result += string(c) + " "
	}
	for !st.IsEmpty() {
		st.Pop()
	}
	if ok && len(*result) > 0 {
		*result = (*result)[:len(*result)-1]
	}
	return ok
}

func handle(line string) {
	var res string
	if InfixToPostfix(line, &res) {
		fmt.Println(res)
	} else {
		fmt.Println("Ошибка в выражении")
	}
}

func main() {
	if len(os.Args) > 1 {
		handle(strings.Join(os.Args[1:], " ") + " ")
		return
	}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		if line := sc.Text(); line != "" {
			handle(line)
		}
	}
}
