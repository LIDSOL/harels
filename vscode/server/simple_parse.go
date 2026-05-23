package main

import "strings"

type Stack []int

type ranges struct {
	l int
	r int
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func getBlockLines(textDocument string) ([]ranges, int) {
	var rang []ranges
	var st Stack
	lines := strings.Split(textDocument, "\n")

	for lineIdx, line := range lines {
		inString := false
		for charIdx, c := range line {
			if inString {
				if c == '"' {
					inString = false
				}
				continue
			}

			if c == '"' {
				inString = true
				continue
			}

			if c == '{' {
				st.Push(lineIdx)
			}
			if c == '}' {
				r, success := st.Pop()
				if success {
					rang = append(rang, ranges{r, lineIdx})
				} else {
					return rang, charIdx
				}
			}
		}
	}

	if len(st) != 0 {
		return rang, st[0]
	}
	return rang, -1
}
