package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var total = 0
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Bytes()
		n := string(l[1:])
		v, _ := strconv.Atoi(n)
		// fmt.Printf("n is %v", n)

		if l[0] == 45 {

			total = total - v
		} else {
			total = total + v
		}

	}
	fmt.Printf("Total is: %v\n", total)

}
