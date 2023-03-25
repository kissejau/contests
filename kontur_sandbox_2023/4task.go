package main

import (
    "fmt"
    "bufio"
    "os"
    "math"
)

func lenOfSide(arr []string, checked [][]bool, x, y, xOffset, yOffset int) int {
    var l int
    for x < len(arr) && y < len(arr[x]) {
        checked[x][y] = true
        x += xOffset
        y += yOffset
        if (arr[x][y] == '.') {
            l++
        } else {
            return l
        }
    }
    return l
}

func main() {
    var n, m, mx int
    rdr := bufio.NewReader(os.Stdin)
    fmt.Fscan(rdr, &n)
    fmt.Fscan(rdr, &m)

    var checked [][]bool = make([][]bool, n)
    var arr []string = make([]string, n)
    for i, _ := range arr {
        fmt.Fscan(rdr, &arr[i])
        checked[i] = make([]bool, m)
    }
    for i, _ := range arr {
        for j, _ := range arr[i] {
            if arr[i][j] == '.' && !checked[i][j]{
                mx = int( math.Max( float64(mx), (1 + float64( lenOfSide(arr, checked, i, j, 1, 0) ) ) * (1 + float64( lenOfSide(arr, checked, i, j, 0, 1) ) ) ) )
            }
            checked[i][j] = true
        }
    }
    fmt.Println(mx)
}
