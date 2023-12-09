package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetExtrapolatedValuesSum(path string) (int, int) {
	filePath := path

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0, 0
	}

	fileContentString := string(fileContent)

	lines := strings.Split(fileContentString, "\n")

	sumForward := 0
	sumBackwards := 0

	for _, line := range lines {
		sumForward += extrapolateForward(line)
		sumBackwards += extrapolateBackwards(line)
	}

	return sumForward, sumBackwards
}

func extrapolate(history string) [][]int {
	h := strings.Split(history, " ")

	var seq [][]int
	var s []int
	var e []int

	for _, v := range h {
		a, _ := strconv.Atoi(v)
		s = append(s, a)
	}
	seq = append(seq, s)

	for count(s, 0) != len(s) {
		for i := 0; i < len(s)-1; i++ {
			a := (s[i])
			b := (s[i+1])
			e = append(e, b-a)
		}
		s = e
		e = nil
		seq = append(seq, s)
	}

	return seq
}

func extrapolateBackwards(history string) int {

	seq := extrapolate(history)

	for i := len(seq) - 1; i > 0; i-- {
		seq[i-1] = prepend(seq[i-1], seq[i-1][0]-seq[i][0])
	}

	return seq[0][0]

}

func extrapolateForward(history string) int {
	seq := extrapolate(history)

	for i := len(seq) - 1; i > 0; i-- {
		seq[i-1] = append(seq[i-1], seq[i][len(seq[i])-1]+seq[i-1][len(seq[i-1])-1])
	}

	return seq[0][len(seq[0])-1]

}

func count(s []int, n int) int {
	count := 0
	for _, v := range s {
		if v == n {
			count++
		}
	}
	return count
}

func prepend(s []int, n int) []int {
	return append([]int{n}, s...)
}
