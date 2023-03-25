package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "sort"
)

var table map[string]int = map[string]int{}
var keys []string

func sortTable() {
    sort.SliceStable(keys, func(a, b int) bool {
        if table[keys[a]] == table[keys[b]] {
            return keys[a] < keys[b]
        }
        return table[keys[a]] > table[keys[b]]
    })
}

func placeAfterCorrectTable(changes map[string]int, team string) int {
    for key, val := range changes {
        table[key] += val
    }
    defer func(){
        for key, val := range changes {
            table[key] -= val
        }
    }()
    sortTable()
    for i, val := range keys {
        if val == team {
            return i+1
        } 
    }
    return -1
}

func test() {
    fmt.Printf("\n---------------------------------\n")
    for _, val := range keys {
        fmt.Printf("%v - %v\n", val, table[val])
    } 
    fmt.Printf("---------------------------------\n")
}

func main() {
    rdr := bufio.NewReader(os.Stdin)
    var countTeams, score int
    var tmp, team string
    var win, draw, lose int
    
    fmt.Fscan(rdr, &countTeams)
    
    for i := 0; i < countTeams; i++ {
        fmt.Fscan(rdr, &team)
        fmt.Fscan(rdr, &score)
        table[team] = score
    }
    for key := range table {
        keys = append(keys, key)
    }
    
    sortTable()
    
    fmt.Fscan(rdr, &tmp)
    match := strings.Split(tmp, "-")
    
    win = placeAfterCorrectTable(map[string]int{match[0]: 3,}, match[0])
    draw = placeAfterCorrectTable(map[string]int{match[0]: 1, match[1]: 1,}, match[0])
    lose = placeAfterCorrectTable(map[string]int{match[1]: 3,}, match[0]) 
    fmt.Printf("%v %v %v", win, draw, lose)
}
