package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func A (path string) (int) {
	// Specify the path to the file you want to read
	filePath := path

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		// Handle the error
		fmt.Println("Error opening file:", err)
		return 0
	}
	
	// Create a Scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Initialize sum
	var sum = 0

	previousLine := ""
	line := scanner.Text()
	nextLine := scanner.Text()
	sum += findPartNumbers(previousLine, line, nextLine)

	
	// Iterate through each line
	for scanner.Scan() {
		previousLine = line
		line = nextLine
		nextLine = scanner.Text()
		sum += findPartNumbers(previousLine, line, nextLine)
	}

	previousLine = line
	line = nextLine
	nextLine = ""
	sum += findPartNumbers(previousLine, line, nextLine)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Close the file when main() ends
	defer file.Close()

	return sum
}

func findPartNumbers(previousLine string, line string, nextLine string) (int) {
	re := regexp.MustCompile(`(\d+)`)
	index := re.FindAllStringIndex(line, -1)
	var sum = 0

	for i := 0; i < len(index); i++ {

		// Is part number on the left side
		if (index[i][0]-1 >= 0 && line[index[i][0]-1] != '.')  {
			parsed, _ := strconv.ParseInt(line[index[i][0]:index[i][1]], 10, 64)
			sum+=int(parsed)
		}

		// Is part number on the right side
		if  (index[i][1] < len(line) && line[index[i][1]] != '.') {
			parsed, _ := strconv.ParseInt(line[index[i][0]:index[i][1]], 10, 64)
			sum+=int(parsed)
		}

		var start = 0
		if (index[i][0]-1 >= 0){
			start = index[i][0]-1
		}

		var end = index[i][1]-1
		if (index[i][1] < len(line)){
			end = index[i][1]
		}

		for j := start; j <= end; j++ {

			if (previousLine != "" && previousLine[j] != '.' && previousLine[j] != '1' && previousLine[j] != '2' && previousLine[j] != '3' && previousLine[j] != '4' && previousLine[j] != '5' && previousLine[j] != '6' && previousLine[j] != '7' && previousLine[j] != '8' && previousLine[j] != '9' && previousLine[j] != '0') {
				parsed, _ := strconv.ParseInt(line[index[i][0]:index[i][1]], 10, 64)
				sum+=int(parsed)
			}
			if (nextLine != "" && nextLine[j] != '.' && nextLine[j] != '1' && nextLine[j] != '2' && nextLine[j] != '3' && nextLine[j] != '4' && nextLine[j] != '5' && nextLine[j] != '6' && nextLine[j] != '7' && nextLine[j] != '8' && nextLine[j] != '9' && nextLine[j] != '0') {
				parsed, _ := strconv.ParseInt(line[index[i][0]:index[i][1]], 10, 64)
				sum+=int(parsed)
			}
		}

		
	}

	return sum
}
