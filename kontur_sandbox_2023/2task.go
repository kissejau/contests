package main

import(
    "fmt"
    "bufio"
    "os"
)

func main() {
    rdr := bufio.NewReader(os.Stdin)
    var sum, num, l int 
    fmt.Fscan(rdr, &l)
    
    for i := 0; i < l; i++ {
        fmt.Fscan(rdr, &num)
        sum += num
    }
    
    fmt.Println(sum * -1)
}
