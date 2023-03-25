package main
 
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
    "sort"
)
 
func exists(arr []int, target int) bool {
	for _, val := range arr {
		if val == target {
			return true
		}
	}
	return false
}
 
func splitStreet(street string) (string, int) {
	for i, val := range street {
		if unicode.IsDigit(val) {
			num, _ := strconv.Atoi(street[i:])
			return street[:i], num
		}
	}
	return "", 0
}
 
func main() {
	var n, m int
	completed := map[string][]int{}
	streets := map[string]int{}
	rdr := bufio.NewReader(os.Stdin)
	fmt.Fscan(rdr, &n)
 
	for i := 0; i < n; i++ {
		var street string
		fmt.Fscan(rdr, &street)
		title, num := splitStreet(street)
		completed[title] = append(completed[title], num)
	}

    for i, _ := range completed {
        sort.Ints(completed[i])
    }
 
	fmt.Fscan(rdr, &m)
	for i := 0; i < m; i++ {
		var title string
		fmt.Fscan(rdr, &title)
		if streets[title] == 0 {
			streets[title] = 1
		}
		for exists(completed[title], streets[title]) {
			streets[title]++
            completed[title] = completed[title][1:]
		}
		fmt.Println(streets[title])
		streets[title]++
	}
}
