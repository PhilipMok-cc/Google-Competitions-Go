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
	// Defind buffer size
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
	N := ReadInt(s)
	var p []int
	var st, en int
	p = make([]int, N)
	for i := 0; i < N; i++ {
		p[i] = ReadInt(s)
	}
	st, en = 0, N-1
	cnt := 0
	level := 0

	for st <= en {
		if p[st] < p[en] {
			if p[st] >= level {
				cnt++
				level = p[st]
			}
			st++
		} else {
			if p[en] >= level {
				cnt++
				level = p[en]
			}
			en--
		}
	}

	// Start Printing
	fmt.Printf("Case #%d: %v\n", caseNum+1, cnt)
}
