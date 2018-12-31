package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Totals struct {
	TwoLetters   int
	ThreeLetters int
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	t := Totals{}
	for scanner.Scan() {
		m := make(map[rune]int)
		counted2 := false
		counted3 := false
		l := scanner.Text()
		for _, v := range l {
			if _, ok := m[v]; !ok {
				m[v]++

				c := strings.Count(l, string(v))
				if c == 2 && counted2 == false {
					counted2 = true
					t.TwoLetters++
				} else if c == 3 && counted3 == false {
					counted3 = true
					t.ThreeLetters++
				}
			}
		}

	}
	fmt.Printf("%v\n", t.ThreeLetters*t.TwoLetters)

}
