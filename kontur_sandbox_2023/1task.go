package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func sortStr(str string, compare func(a, b byte) bool) string  {
    var fl bool
    arr := []byte(str)
    for !fl {
        fl = true
        for i := 0; i < len(arr)-1; i++ {
            if !compare(arr[i], arr[i+1]) {
                fl = false
                arr[i], arr[i+1] = arr[i+1], arr[i]
            }
        }
        if fl {
            break
        }
    }
    return string(arr)
}

func reverseNum(num string) string {
    var tmp string
    for i := len(num)-1; i >= 0; i-- {
        tmp += string(num[i])
    }
    return tmp
}

func main() {
    rdr := bufio.NewReader(os.Stdin)
    var num string
    fmt.Fscan(rdr, &num)
    mx := sortStr(num, func(a byte, b byte) bool {
        return a >= b
    })
    n1, _ := strconv.Atoi(mx)
    n2, _ := strconv.Atoi(reverseNum(mx))
    fmt.Println(n1 - n2)
}
