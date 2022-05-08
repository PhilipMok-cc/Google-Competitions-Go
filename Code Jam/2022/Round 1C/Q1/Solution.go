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
	var t []string
	t = make([]string, 0)
	cnt := true
	for i := 0; i < N; i++ {
		tmp := ReadWord(s)
		if cnt && !checkValid(tmp) {
			cnt = false
		} else {
			if tmp[0] == tmp[len(tmp)-1] {
				t = append([]string{tmp}, t...)
			} else {
				t = append(t, tmp)
			}
		}
	}
	if !cnt {
		fmt.Printf("Case #%d: %v\n", caseNum+1, "IMPOSSIBLE")
		return
	}
	l := len(t)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; {
			if t[i][0] == t[j][len(t[j])-1] {
				t[i] = t[j] + t[i]
				t = Remove(t, j)
				l--
				j = i + 1
			} else if t[i][len(t[i])-1] == t[j][0] {
				t[i] = t[i] + t[j]
				t = Remove(t, j)
				l--
				j = i + 1
			} else {
				j++
			}
		}
	}

	final := strings.Join(t, "")

	// Start Printing
	if checkValid(final) {
		fmt.Printf("Case #%d: %v\n", caseNum+1, final)
	} else {
		fmt.Printf("Case #%d: %v\n", caseNum+1, "IMPOSSIBLE")

	}
}

func checkValid(a string) bool {
	var m map[rune]bool = make(map[rune]bool)
	var last rune
	for i, c := range a {
		if i == 0 {
			m[c] = true
			last = c
		} else if c != last {
			if _, found := m[c]; found {
				return false
			}
			m[c] = true
			last = c
		}
	}
	return true
}

func Remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
