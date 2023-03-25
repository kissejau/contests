package main

import (
    "bufio"
    "os"
    "fmt"
)

func main() {
    rdr := bufio.NewReader(os.Stdin)
    var n, m, k int
    fmt.Fscan(rdr, &n)
    fmt.Fscan(rdr, &m)
    fmt.Fscan(rdr, &k)

    mins := 0
    cur := n
    for i := 0; i < k; i++ {
        div := cur / m
        if cur % m != 0 {
            div++
        }
        mins += div
        out := cur - m * div
        if (out < 0) {
            out *= -1
        }
        cur = n - out
    }
    fmt.Println(mins)
}
