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
	E := ReadInt(s)
	W := ReadInt(s)
	var total_cnt []int
	var weight_cnt, shared_cnt, action_cnt [][]int
	var tmp_shared_cnt [][][]int
	weight_cnt = make([][]int, E)
	shared_cnt = make([][]int, E)
	total_cnt = make([]int, E)
	action_cnt = make([][]int, E)
	tmp_shared_cnt = make([][][]int, E)

	for i := 0; i < E; i++ {
		weight_cnt[i] = make([]int, W)
		for j := 0; j < W; j++ {
			weight_cnt[i][j] = ReadInt(s)
			total_cnt[i] += weight_cnt[i][j]
		}
	}
	// There are E+1 status if there are E execrise
	// Status 0 is no weight in the rod
	// shared_cnt[E][E+1]
	// shared_cnt[i][j] is shared block transform from i to j
	// i=0 is initial status with no weight
	for i := 0; i < E; i++ {
		shared_cnt[i] = make([]int, E+1)
		tmp_shared_cnt[i] = make([][]int, E+1)
		tmp_shared_cnt[i][i] = make([]int, W)
		for j := i; j < E; j++ {
			tmp_shared_cnt[i][j+1] = make([]int, W)
			for k := 0; k < W; k++ {
				if i == j {
					// Consider Execrise start from i instead of 0, first action should share all everything
					tmp_shared_cnt[i][j+1][k] = weight_cnt[i][k]
				} else {
					// When number of execrise started the share cnt will decread base on weight cnt of each action
					tmp_shared_cnt[i][j+1][k] = min(tmp_shared_cnt[i][j][k], weight_cnt[j][k])
				}
				shared_cnt[i][j+1] += tmp_shared_cnt[i][j+1][k]
			}
		}
	}

	// Only count add weight action
	// Remove weight action will be same as add weight action
	// Calculate number of acntion when change of state 1,2,3...E
	// action_cnt[E][E+1]
	for i := 0; i < E; i++ {
		action_cnt[i] = make([]int, E+1)
		//Number of action i to i+1 status
		//Weight of status i+1 => execrise i (status 0 is initial status)
		action_cnt[i][i+1] = total_cnt[i]
	}

	for d := 2; d <= E; d++ {
		for i, j := 0, d; j <= E; i, j = i+1, j+1 {
			// set to a very large value for min operation
			action_cnt[i][j] = 1e12
			for k := i + 1; k <= j-1; k++ {
				// Number of add weight from i to j status
				// take min of each mid point
				// A(0,1) is already calculated
				// A(0,2) = A(0,1)+A(1,2)-shared(0,2)
				// A(0,3) = min(A(0,1)+A(1,3)-shared(0,3),A(0,2)+A(2,3)-shared(0,3))...
				action_cnt[i][j] = min(action_cnt[i][j], action_cnt[i][k]+action_cnt[k][j]-shared_cnt[i][j])
			}
		}
	}
	// log.Printf("E:%v,W:%v", E, W)
	// log.Printf("weight:")
	// for i := 0; i < E; i++ {
	// 	log.Printf("%v,%v", weight_cnt[i], total_cnt[i])
	// }
	// log.Printf("displace:")
	// for i := 0; i < E; i++ {
	// 	log.Printf("%v", action_cnt[i])
	// }
	// log.Printf("share:")
	// for i := 0; i < E; i++ {
	// 	log.Printf("%v", shared_cnt[i])
	// }

	// Print result,
	// Remove weight action will be same as add weight action
	// Action count *2 will be the result
	fmt.Printf("Case #%d: %v\n", caseNum+1, action_cnt[0][E]*2)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
