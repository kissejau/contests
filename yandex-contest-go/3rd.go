package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
)

func minIndex(arr []int) int {
    index := 0
    for i, _ := range arr {
        if arr[i] < arr[index] {
            index = i
        }
    }
    return index
}

func main() {
    rdr := bufio.NewReader(os.Stdin)
    var n int
    var s string
    fmt.Fscan(rdr, &n)
    fmt.Fscan(rdr, &s)
    chs := []string {"a", "b", "c", "d"}
    chsCount := []int {
        strings.Count(s, chs[0]),
        strings.Count(s, chs[1]),
        strings.Count(s, chs[2]),
        strings.Count(s, chs[3])} 

    if chsCount[0] == 0 || chsCount[1] == 0 || chsCount[2] == 0 || chsCount[3] == 0 {
        fmt.Println(-1)
        return
    }

    mn := minIndex(chsCount)
    if chsCount[mn] == 1 {
        pos := strings.Index(s, chs[mn])
        for i := 3; i < len(s); i++ {
            if pos - i >= 0 {
                if (strings.Contains(s[pos-i:pos+1], chs[0]) &&
                    strings.Contains(s[pos-i:pos+1], chs[1]) &&
                    strings.Contains(s[pos-i:pos+1], chs[2]) &&
                    strings.Contains(s[pos-i:pos+1], chs[3])) {
                    fmt.Println(i+1)
                    return
                }
            }
            if pos + i <= len(s) {
                if (strings.Contains(s[pos:pos+i+1], chs[0]) &&
                    strings.Contains(s[pos:pos+i+1], chs[1]) &&
                    strings.Contains(s[pos:pos+i+1], chs[2]) &&
                    strings.Contains(s[pos:pos+i+1], chs[3])) {
                    fmt.Println(i+1)
                    return
                }
            }
        }
    }

    for i := 4; i <= len(s); i++ {
        for j := 0; j <= len(s) - i; j++ {
            if (strings.Contains(s[j:j+i], chs[0]) &&
                strings.Contains(s[j:j+i], chs[1]) &&
                strings.Contains(s[j:j+i], chs[2]) &&
                strings.Contains(s[j:j+i], chs[3])) {
                fmt.Println(i)
                return
            }
        }
    }
    fmt.Println(-1)
}
