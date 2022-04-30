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
func arrayToString(a []int) string {
	return strings.Trim(fmt.Sprint(a), "[]")
}

func main() {
	// Defind buffer size
	const bufferSize = 512 * 1024

	// Read from Stdin
	scanner := bufio.NewScanner(os.Stdin)
	// Config Buffer
	scanner.Buffer(make([]byte, bufferSize), bufferSize)
	// Config Buffer to read word by word
	scanner.Split(bufio.ScanWords)

	numOfCase := ReadInt(scanner)
	// prcess input block by block
	for i := 0; i < numOfCase; i++ {
		Solve(i, scanner)
	}
}

func Solve(caseNum int, s *bufio.Scanner) {
	N := ReadInt(s)
	P := ReadInt(s)
	var a [][]int = make([][]int, N)
	for i := 0; i < N; i++ {
		a[i] = make([]int, 2)
		for j := 0; j < P; j++ {
			t := ReadInt(s)
			if j == 0 {
				a[i][0] = t
				a[i][1] = t
			} else if t > a[i][1] {
				a[i][1] = t
			} else if a[i][0] > t {
				a[i][0] = t
			}
		}
	}

	base_click := a[0][1]
	top, bot := 0, a[0][1]-a[0][0]

	for i := 1; i < N; i++ {
		base_click += a[i][1] - a[i][0]
		nbot := min(top+diff(a[i-1][1], a[i][1]), bot+diff(a[i-1][0], a[i][1]))
		ntop := min(top+diff(a[i-1][1], a[i][0]), bot+diff(a[i-1][0], a[i][0]))
		top = ntop
		bot = nbot
	}

	// Start Printing
	fmt.Printf("Case #%d: %v\n", caseNum+1, base_click+min(top, bot))
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
