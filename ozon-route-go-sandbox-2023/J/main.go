package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Word struct {
	Word string
	End  byte
}

func NewWord(word string) *Word {
	return &Word{
		Word: word,
		End:  word[len(word)-1],
	}
}

func main() {
	var (
		dictCount  int
		wordsCount int
		dict       []Word
		dictWord   string
		word       string
		in         *bufio.Reader = bufio.NewReader(os.Stdin)
		out        *bufio.Writer = bufio.NewWriter(os.Stdout)
	)
	defer out.Flush()

	fmt.Fscan(in, &dictCount)
	for i := 0; i < dictCount; i++ {
		fmt.Fscan(in, &dictWord)
		dict = append(dict, *NewWord(dictWord))
	}

	fmt.Fscan(in, &wordsCount)
	for i := 0; i < wordsCount; i++ {
		fmt.Fscan(in, &word)
		maxMatch := -1
		maxMatchIndex := 0
		for j := range dict {
			dictWord := dict[j]
			if dictWord.Word == word {
				continue
			}
			if !isMatchEnd(word[len(word)-1], dict) {
				maxMatchIndex = j
				break
			}
			match := matchCount(dictWord.Word, word)
			if match > maxMatch {
				maxMatch = match
				maxMatchIndex = j
				if match == len(word) {
					break
				}
			}

		}
		fmt.Fprintf(out, "%v\n", dict[maxMatchIndex].Word)
	}
}

func matchCount(target string, word string) int {
	var count int
	minLen := int(math.Min(float64(len(target)), float64(len(word))))
	for i := 1; i <= minLen; i++ {
		if target[len(target)-i] != word[len(word)-i] {
			break
		}
		count++
	}
	return count
}

func isMatchEnd(target byte, dict []Word) bool {
	for _, word := range dict {
		if target == word.End {
			return true
		}
	}
	return false
}
