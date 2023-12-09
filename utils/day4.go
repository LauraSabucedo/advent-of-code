package utils

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strings"
)

func processLinesDay4(path string, processLine func(string)) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		processLine(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func TotalPoints(path string) float64 {
	var sum float64 = 0

	processLine := func(line string) {
		matches := getCardMatchingNumbers(line)
		if matches > 0 {
			sum += math.Pow(2, float64(matches-1))
		}
	}

	processLinesDay4(path, processLine)

	return sum

}

func TotalScratchCards(path string) int {
	filePath := path

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	fileContentString := string(fileContent)

	lines := strings.Split(fileContentString, "\n")

	matches := 0
	sum := 0
	cards := make(map[int]int)

	for index := range lines {
		card := index + 1
		cards[card]++
		matches = getCardMatchingNumbers(lines[index])
		if matches > 0 {
			for i := index + 1; i <= matches+index && i < len(lines); i++ {
				cards[i+1] += cards[card]
			}
		}
	}

	for _, v := range cards {
		sum += v
	}

	return sum
}

func getCardMatchingNumbers(card string) int {
	c := regexp.MustCompile(`Card\s+\d+:`).ReplaceAllString(card, "")
	numbers := strings.Split(c, "|")
	winningNumbers := numbers[0]
	gottenNumbers := numbers[1]

	matches := 0

	a := regexp.MustCompile(`(\d+)`).FindAllString(winningNumbers, -1)
	b := strings.Fields(gottenNumbers)

	for _, winningNumber := range a {

		if slices.Contains(b, winningNumber) {
			matches++
		}
	}

	return matches
}
