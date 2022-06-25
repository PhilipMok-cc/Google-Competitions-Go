package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInt(s *bufio.Scanner) int {
	//AssumeSplit with Scanwords
	s.Scan()
	res, _ := strconv.Atoi(s.Text())
	return res
}

func ReadWord(s *bufio.Scanner) string {
	//Assume Split with Scanwords
	s.Scan()
	return s.Text()
}
func arrayToString(a []rune) string {
	return strings.Trim(fmt.Sprint(a), "[]")
}

func main() {
	// Define buffer size
	const bufferSize = 512 * 1024

	// Read from Stdin
	scanner := bufio.NewScanner(os.Stdin)
	// Set Buffer Size
	scanner.Buffer(make([]byte, bufferSize), bufferSize)
	// Config Buffer to read word by word
	scanner.Split(bufio.ScanWords)

	// Read Number of Case
	numOfCase := ReadInt(scanner)
	// Solve each case
	for i := 0; i < numOfCase; i++ {
		Solve(i, scanner)
	}
}

//Logic goes here
func Solve(caseNum int, s *bufio.Scanner) {
	S := ReadWord(s)
	var fin, temp []rune
	var last rune
	for x, i := range S {
		if x == 0 {
			temp = []rune{i}
			last = i
			continue
		}
		if last < i {
			fin = append(fin, temp...)
			fin = append(fin, temp...)
			temp = []rune{i}
			last = i
		} else if last == i {
			temp = append(temp, i)
		} else {
			fin = append(fin, temp...)
			temp = []rune{i}
			last = i
		}

	}
	fin = append(fin, temp...)
	fa := string(fin)

	// Start Printing
	fmt.Printf("Case #%d: %v\n", caseNum+1, fa)
}
