package utils

import (
	"fmt"
	"regexp"
	"os"
	"bufio"
	"strconv"
)

var red = 12
var green = 13
var blue = 14

func processLines(path string, processLine func(string)) {
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

// Part 1

func PossibleGamesSum(path string) int {
	var sum int

	processLine := func(line string) {
		if gameIsPossible(line, red, blue, green) {
			gameregex := regexp.MustCompile(`Game\s(\d+):`)
			gameNum := gameregex.FindStringSubmatch(line)[1]
			parsed, _ := strconv.ParseInt(gameNum, 10, 64)
			sum += int(parsed)
		}
	}

	processLines(path, processLine)

	return sum
}

func gameIsPossible(game string, red int, blue int, green int) (bool){
	return colorIsPossible(game, "red", red) && colorIsPossible(game, "blue", blue) && colorIsPossible(game, "green", green)
}

func colorIsPossible(game string, color string, amount int) (bool) {
	var isValid = true

	re := regexp.MustCompile(fmt.Sprintf(`(\d+)\s%s`, color))
	for i := range re.FindAllStringSubmatch(game, -1) {
		greenNum := re.FindAllStringSubmatch(game, -1)[i][1]
		parsed, _ := strconv.ParseInt(greenNum, 10, 64)
		if int(parsed) > amount {
			isValid = false
		}
	}

	return isValid
}

// Part 2

func TotalPower(path string) int {
	var totalPower int

	processLine := func(line string) {
		totalPower += setPower(line, "red", red)
	}

	processLines(path, processLine)

	return totalPower
}


func setPower(game string, color string, amount int) (int) {
	return fewestAmountOfColor(game, "red") * fewestAmountOfColor(game, "blue") * fewestAmountOfColor(game, "green")
}

func fewestAmountOfColor(game string, color string) (int) {
	var max = 0

	re := regexp.MustCompile(fmt.Sprintf(`(\d+)\s%s`, color))
	for i := range re.FindAllStringSubmatch(game, -1) {
		greenNum := re.FindAllStringSubmatch(game, -1)[i][1]
		parsed, _ := strconv.ParseInt(greenNum, 10, 64)
		if int(parsed) > max {
			max = int(parsed)
		}
	}

	return max
}