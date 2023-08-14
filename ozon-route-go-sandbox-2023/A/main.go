package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		count int
		n1    int
		n2    int
		in    = bufio.NewReader(os.Stdin)
		out   = bufio.NewWriter(os.Stdout)
	)
	defer out.Flush()

	fmt.Fscan(in, &count)

	for i := 0; i < count; i++ {
		fmt.Fscan(in, &n1)
		fmt.Fscan(in, &n2)
		fmt.Fprintf(out, "%v\n", n1+n2)
	}
}
