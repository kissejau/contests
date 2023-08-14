package main

import (
	"bufio"
	"fmt"
	"os"
)

const DISCOUNT int = 3

func main() {
	var (
		count         int
		productsCount int
		products      map[int]int
		product       int
		in            *bufio.Reader = bufio.NewReader(os.Stdin)
		out           *bufio.Writer = bufio.NewWriter(os.Stdout)
	)
	defer out.Flush()

	fmt.Fscan(in, &count)
	for i := 0; i < count; i++ {
		var price int
		products = make(map[int]int)
		fmt.Fscan(in, &productsCount)
		for j := 0; j < productsCount; j++ {
			fmt.Fscan(in, &product)
			products[product]++
		}
		for key, val := range products {
			price += (val % DISCOUNT * key) + (val / 3 * 2 * key)
		}
		fmt.Fprintf(out, "%v\n", price)
	}
}
