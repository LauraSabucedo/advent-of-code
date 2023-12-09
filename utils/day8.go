package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Steps(path string) int {
	filePath := path

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	fileContentString := string(fileContent)

	lines := strings.Split(fileContentString, "\n")

	directions := lines[0]

	coordinates := make(map[string][]string)

	for i := 2; i < len(lines); i++ {
		node := strings.Split(string(lines[i]), " =")[0]
		start := strings.Index(lines[i], "(") + 1
		end := strings.Index(lines[i], ")")
		coordinates[node] = strings.Split(lines[i][start:end], ", ")
	}

	coordinate := "AAA"
	count := 0
	c := 0

	for coordinate != "ZZZ" {
		switch directions[c] {
		case 'L':
			newCoordinate := coordinates[coordinate][0]
			coordinate = newCoordinate
			break
		case 'R':
			newCoordinate := coordinates[coordinate][1]
			coordinate = newCoordinate
			break
		}
		count++
		if c == len(directions)-1 {
			c = 0
		} else {
			c++
		}
	}

	return count
}

func GhostSteps(path string) int64 {
	filePath := path

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	fileContentString := string(fileContent)

	lines := strings.Split(fileContentString, "\n")

	directions := lines[0]

	coordinates := make(map[string][]string)

	for i := 2; i < len(lines); i++ {
		node := strings.Split(string(lines[i]), " =")[0]
		start := strings.Index(lines[i], "(") + 1
		end := strings.Index(lines[i], ")")
		coordinates[node] = strings.Split(lines[i][start:end], ", ")
	}

	var currentCoordinate []string

	for key, _ := range coordinates {
		if ok := regexp.MustCompile(`..A`).MatchString(key); ok {
			currentCoordinate = append(currentCoordinate, key)
		}
	}

	var results []int

	for _, coordinate := range currentCoordinate {
		count := 0
		c := 0
		for !regexp.MustCompile(`..Z`).MatchString(coordinate) {
			switch directions[c] {
			case 'L':
				newCoordinate := coordinates[coordinate][0]
				coordinate = newCoordinate
				break
			case 'R':
				newCoordinate := coordinates[coordinate][1]
				coordinate = newCoordinate
				break
			}
			count++
			if c == len(directions)-1 {
				c = 0
			} else {
				c++
			}
		}
		results = append(results, count)
	}

	return arrayLeastCommonMultiple(results)
}

func greatestCommonDivisor(a int64, b int64) int64 {
	if b == 0 {
		return a
	}
	return greatestCommonDivisor(b, a%b)
}

func leastCommonMultiple(a int64, b int64) int64 {
	return (a / greatestCommonDivisor(a, b)) * b
}

func arrayLeastCommonMultiple(numbers []int) int64 {
	result := int64(numbers[0])
	for i := 1; i < len(numbers); i++ {
		result = leastCommonMultiple(result, int64(numbers[i]))
	}
	return result
}
