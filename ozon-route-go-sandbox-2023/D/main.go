package main

import (
	"bufio"
	"fmt"
	"os"
)

func NewMatrix(col, row int) *Matrix {
	arr := make([][]int, col)
	for i := range arr {
		arr[i] = make([]int, row)
	}
	return &Matrix{
		arr:             arr,
		lastModifiedCol: -1,
	}
}

type Matrix struct {
	arr             [][]int
	lastModifiedCol int
}

func (m *Matrix) Sort(col int) {
	fl := true
	for fl {
		fl = false
		for i := 0; i < len(m.arr[col])-1; i++ {
			if m.arr[col][i] > m.arr[col][i+1] {
				for j := 0; j < len(m.arr); j++ {
					m.arr[j][i], m.arr[j][i+1] = m.arr[j][i+1], m.arr[j][i]
				}
				fl = true
			}
		}
	}
}

func (m *Matrix) String() string {
	var out string
	for i := 0; i < len(m.arr[0]); i++ {
		for j := 0; j < len(m.arr); j++ {
			out += fmt.Sprintf("%v ", m.arr[j][i])
		}
		out += fmt.Sprintf("\n")
	}
	return out
}

func main() {
	var (
		n, m         int
		count        int
		countChanges int
		colSort      int
		in           *bufio.Reader = bufio.NewReader(os.Stdin)
		out          *bufio.Writer = bufio.NewWriter(os.Stdout)
	)
	defer out.Flush()

	fmt.Fscan(in, &count)
	for i := 0; i < count; i++ {
		fmt.Fscan(in, &n)
		fmt.Fscan(in, &m)
		matrix := NewMatrix(m, n)
		for row := 0; row < n; row++ {
			for col := 0; col < m; col++ {
				fmt.Fscan(in, &matrix.arr[col][row])
			}
		}
		fmt.Fscan(in, &countChanges)
		for j := 0; j < countChanges; j++ {
			fmt.Fscan(in, &colSort)
			if matrix.lastModifiedCol == colSort {
				continue
			}

			matrix.Sort(colSort - 1)
			matrix.lastModifiedCol = colSort
		}
		fmt.Fprintln(out, matrix)
	}
}
