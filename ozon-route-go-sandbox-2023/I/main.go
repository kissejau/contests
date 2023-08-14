package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var (
		computersCount, tasksCount int
		computers                  [][2]int
		cost                       int
		startTime, execTime        int
		sum                        int
		border                     int
		in                         *bufio.Reader = bufio.NewReader(os.Stdin)
		out                        *bufio.Writer = bufio.NewWriter(os.Stdout)
	)
	defer out.Flush()

	fmt.Fscan(in, &computersCount)
	fmt.Fscan(in, &tasksCount)

	for i := 0; i < computersCount; i++ {
		fmt.Fscan(in, &cost)
		computers = append(computers, [2]int{cost, 1})
	}
	sort.Slice(computers, func(i, j int) bool {
		return computers[i][0] < computers[j][0]
	})

	for i := 0; i < tasksCount; i++ {
		fmt.Fscan(in, &startTime)
		fmt.Fscan(in, &execTime)
		if startTime < border {
			continue
		}
		for i := range computers {
			if computers[i][1] <= startTime {
				sum += computers[i][0] * execTime
				computers[i][1] = startTime + execTime
				recalcBorder(&border, computers)
				break
			}
		}
	}
	fmt.Fprintf(out, "%v\n", sum)
}

func recalcBorder(border *int, arr [][2]int) {
	min := arr[0][1]
	for i := 1; i < len(arr); i++ {
		if arr[i][1] < min {
			min = arr[i][1]
		}
	}
	*border = min
}
