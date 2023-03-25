package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var sum, maxScore int

func isMax() bool {
	if sum == maxScore {
		return true
	}
	return false
}

func main() {
	rdr := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(rdr, &n)
	fmt.Fscan(rdr, &maxScore)

	table := make([][]int, n)

	for i := range table {
		table[i] = make([]int, 3)
		fmt.Fscan(rdr, &table[i][1])
		fmt.Fscan(rdr, &table[i][2])
		table[i][0] = table[i][2]
		sum += table[i][0]
	}

	sort.Slice(table, func(x, y int) bool {
		return table[x][2] > table[y][2]
	})

	for {
		for i := len(table) - 1; i > len(table)/2; i-- {
			for table[i][0] > table[i][1] {
				table[i][0]--
				sum--
				if isMax() {
					fmt.Println(table[len(table)/2][0])
					return
				}
			}
		}
		for {
			for i := len(table)/2 - 1; i >= 0; i-- {
				if table[i][0] == table[len(table)/2][0] {
					table[len(table)/2][0]--
					sum--
					if isMax() {
						fmt.Println(table[len(table)/2][0])
						return
					}
				}
				table[i][0]--
				sum--
				if isMax() {
					fmt.Println(table[len(table)/2][0])
					return
				}
			}
		}
	}
}
