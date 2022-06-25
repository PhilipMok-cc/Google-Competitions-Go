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

const pi float64 = 3.1415927

//Logic goes here
func Solve(caseNum int, s *bufio.Scanner) {
	R := ReadInt(s)
	A := ReadInt(s)
	B := ReadInt(s)
	var CurR, RsqrSum int
	CurR = R
	RsqrSum = 0
	for CurR > 0 {
		RsqrSum += (CurR * CurR) + (CurR * CurR * A * A)
		CurR = CurR * A / B
	}

	// Start Printing
	fmt.Printf("Case #%d: %v\n", caseNum+1, float64(RsqrSum)*pi)
}
