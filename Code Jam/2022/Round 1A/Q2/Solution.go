package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadInt(s *bufio.Scanner) int {
	//AssumeSplit with Scanwords
	s.Scan()
	res, _ := strconv.Atoi(s.Text())
	CheckExit(res)
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

func CheckExit(r int) {
	if r == -1 {
		os.Exit(3)
	}
}

//Logic goes here
const MaxValue int = 1e9

func Solve(caseNum int, s *bufio.Scanner) {
	var input, answer []int = []int{}, []int{}
	var target int
	N := ReadInt(s)
	temp := (1 << 29)
	for i := 0; i < N; i++ {
		if i < 30 {
			input = append(input, (1 << uint(i)))
		} else {
			input = append(input, (temp + i - 29))
		}
	}
	fmt.Printf("%v\n", arrayToString(input))
	for i := 0; i < N; i++ {
		Ans := ReadInt(s)
		input = append(input, Ans)
		// log.Printf("%v %v\n", i, Ans)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(input)))
	target = sumArray(input) / 2
	for _, val := range input {
		if val < target {
			target -= val
			answer = append(answer, val)
		} else if val == target {
			// log.Printf("%v %v\n", val, target)
			answer = append(answer, val)
			break
		}
	}
	fmt.Printf("%v\n", arrayToString(answer))
}

func sumArray(numbs []int) int {
	result := 0
	for _, numb := range numbs {
		result += numb
	}
	return result
}
