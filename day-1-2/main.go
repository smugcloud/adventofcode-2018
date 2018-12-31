package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// total := 0
	// m := make(map[int]int)
	for i := 0; i < 10; i++ {
		f, _ := os.Open("input.txt")

		ScanFile(f)
		f.Close()
		// 	// fmt.Println(i)
		// 	scanner := bufio.NewScanner(f)
		// 	for scanner.Scan() {
		// 		//Get the line in bytes
		// 		l := scanner.Bytes()
		// 		fmt.Println(string(l))
		// 		//Get the number as a string
		// 		n := string(l[1:])
		// 		//Get the int value of the string
		// 		v, _ := strconv.Atoi(n)

		// 		//Negative sign in the string
		// 		if l[0] == 45 {

		// 			total = total - v
		// 		} else {
		// 			total = total + v
		// 		}
		// 		_, ok := m[total]
		// 		if !ok {
		// 			m[total] = total
		// 		} else {
		// 			fmt.Println("Match found.")
		// 			fmt.Printf("First match is: %v\n", m[total])

		// 		}
		// 		// fmt.Println(total)
		// 	}
		// 	if err := scanner.Err(); err != nil {
		// 		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		// 	}
		// 	scanner = bufio.NewScanner(strings.NewReader(""))
		// }
		// fmt.Printf("Does 1 exist in map? %v\n", m[1])
		// fmt.Printf("m is: %v\n", m)

	}
}

func ScanFile(f *os.File) {
	fmt.Println("In ScanFile")
	scanner := bufio.NewScanner(f)
	fmt.Println(scanner)
	for scanner.Scan() {
		//Get the line in bytes
		l := scanner.Bytes()
		fmt.Println(string(l))

	}
	scanner = nil
}
