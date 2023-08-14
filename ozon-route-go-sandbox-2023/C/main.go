package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var (
		count           int
		developersCount int
		developers      []int
		in              *bufio.Reader = bufio.NewReader(os.Stdin)
		out             *bufio.Writer = bufio.NewWriter(os.Stdout)
	)
	defer out.Flush()
	fmt.Fscan(in, &count)

	for i := 0; i < count; i++ {
		fmt.Fscan(in, &developersCount)
		developers = make([]int, developersCount)
		for j := 0; j < developersCount; j++ {
			fmt.Fscan(in, &developers[j])
		}
		groups := getGroupedDevelopers(developers)
		for _, group := range groups {
			fmt.Fprintf(out, "%v %v\n", group[0], group[1])
		}
		fmt.Fprintln(out)
	}
}

func getGroupedDevelopers(devs []int) [][2]int {
	var indexOfMin int
	var min int
	var groupNum int
	groups := make([][2]int, len(devs)/2)
	for i := range devs {
		dev := devs[i]
		if dev == -1 { // i`ll asign -1 for grouped deveopers
			continue
		}

		min = 101 // max skill level

		for j := i + 1; j < len(devs); j++ {
			if devs[j] == -1 {
				continue
			}
			diff := int(math.Abs(float64(dev - devs[j])))
			if diff < min {
				min = diff
				indexOfMin = j
			}
		}
		groups[groupNum] = [2]int{i + 1, indexOfMin + 1}
		groupNum++
		devs[i], devs[indexOfMin] = -1, -1
	}
	return groups
}
