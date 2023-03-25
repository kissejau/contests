package main

import (
    "bufio"
    "os"
    "fmt"
)

func main() {
    rdr := bufio.NewReader(os.Stdin)
    var n1, n2, n3, n4 int
    fmt.Fscan(rdr, &n1)
    fmt.Fscan(rdr, &n2)
    fmt.Fscan(rdr, &n3)
    fmt.Fscan(rdr, &n4)

    if (n1 <= n2 && n2 <= n3 && n3 <= n4) || (n1 >= n2 && n2 >= n3 && n3 >= n4) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
