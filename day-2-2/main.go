package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var words []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		words = append(words, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	for i := range words {
		for _, v := range words[i+1:] {

			common, ok := compare(words[i], v)
			if ok {
				fmt.Printf("Returned common: %v\n", common)
				return
			}

		}
	}
}

func compare(a, b string) (string, bool) {
	dif := 0
	idx := 0

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}
		dif++
		if dif > 1 {
			return "", false
		}
		idx = i
	}
	if idx == 0 || dif == 0 {
		return "", false

	}

	fmt.Printf("Matching strings with %v differences are:\na: %v\nb: %v\nat idx: %v\n", dif, a, b, idx)
	return a[:idx] + a[idx+1:], true
}
