package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func WaysToBeatRecord(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	time := strings.Fields(strings.Split(scanner.Text(), ":")[1])

	scanner.Scan()
	distance := strings.Fields(strings.Split(scanner.Text(), ":")[1])

	part1 := WaysToBeatRecordPart1(time, distance)

	part2 := WaysToBeatRecordPart1([]string{strings.Join(time, "")}, []string{strings.Join(distance, "")})

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return part1, part2
}

func WaysToBeatRecordPart1(time []string, distance []string) int {
	result := 1
	for i := 0; i < len(time); i++ {

		count := 0
		x, _ := strconv.Atoi(time[i])
		d, _ := strconv.Atoi(distance[i])

		for y := 0; y < x; y++ {
			if y*(x-y) > d {
				count++
			}
		}

		result *= count
	}

	return result
}
