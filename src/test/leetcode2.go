package main

import "fmt"

func isValid(s string) bool {
	s2word := ""
	for i := len(s); i >= 1; i-- {
		//模拟stack，弹出栈顶
		word := s[i-1 : i]
		flag := false
		size := len(s2word)
		//2个stack，栈顶比较
		if len(s2word) > 0 && isMatch(word, s2word[size-1:size]) {
			flag = true
			//模拟stack，弹出栈顶
			s2word = s2word[0 : size-1]
		}
		if !flag {
			//入栈
			s2word = fmt.Sprintf("%s%s", s2word, word)
		}
	}
	return len(s2word) == 0
}

func isMatch(src string, dest string) bool {
	if src == "{" && dest == "}" {
		return true
	}
	if src == "[" && dest == "]" {
		return true
	}
	if src == "(" && dest == ")" {
		return true
	}
	return false
}

func main() {
	flag := isValid("]")
	fmt.Println(flag)
}
