package main

import (
    "bufio"
    "os"
    "fmt"
    "sort"
)

func main() {
    rdr := bufio.NewReader(os.Stdin)
    var n int 
    reps := map[int]int{}
    fmt.Fscan(rdr, &n)
    arr := make([]int, n)
    
    for i := 0; i < n; i++ {
        fmt.Fscan(rdr, &arr[i])
        reps[arr[i]]++
    }
    keys := make([]int, len(reps))
    i := 0
    for key := range reps {
        keys[i] = key
        i++
    }

    for i := len(arr) - 1; i >= 0; i-- {
        sort.Slice(keys, func(x, y int) bool {
            return reps[keys[x]] > reps[keys[y]]
        })
        var s, f, target int
        if reps[keys[0]] == reps[keys[1]] {
            s, f = 0, len(keys) - 1
            target = reps[keys[0]]
        } else {
            s, f = 1, len(keys)
            target = reps[keys[0]] - 1
        }
        var fl bool = true
        for i := s; i < f; i++ {
            if reps[keys[i]] != target && reps[keys[i]] != 0 {
                fl = false
                break
            }
        }
        if fl {
            sum := 0
            for _, v := range keys {
                sum += reps[v]
            }
            fmt.Println(sum)
            return
        }
        reps[arr[i]]--
        if reps[arr[i]] == 0 {
           keys = append(keys[:len(keys)-1]) 
        }
    }
}
