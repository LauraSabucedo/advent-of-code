package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func SumPartNumbers (path string, regex string, prohibitedChars string) (int, int64) {
	// Specify the path to the file you want to read
	filePath := path

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		// Handle the error
		fmt.Println("Error opening file:", err)
		return 0, 0
	}
	
	// Create a Scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Initialize sum
	var sum = 0
	var sum1 = int64(0)

	previousLine := ""
	line := scanner.Text()
	nextLine := scanner.Text()
	sum += findPartNumbers(previousLine, line, nextLine, regex, prohibitedChars)
	sum1 += findGearRatios(previousLine, line, nextLine)

	// Iterate through each line
	for scanner.Scan() {
		previousLine = line
		line = nextLine
		nextLine = scanner.Text()
		sum += findPartNumbers(previousLine, line, nextLine, regex, prohibitedChars)
		sum1 += findGearRatios(previousLine, line, nextLine)
	}

	previousLine = line
	line = nextLine
	nextLine = ""
	sum += findPartNumbers(previousLine, line, nextLine, regex, prohibitedChars)
	sum1 += findGearRatios(previousLine, line, nextLine)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Close the file when main() ends
	defer file.Close()

	return sum, sum1
}

func findPartNumbers(previousLine string, line string, nextLine string, regex string, prohibitedChars string) (int) {
	re := regexp.MustCompile(regex)
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
			if (previousLine != "" && !strings.ContainsAny(string(previousLine[j]), prohibitedChars)) {
				parsed, _ := strconv.ParseInt(line[index[i][0]:index[i][1]], 10, 64)
				sum+=int(parsed)
			}
			if (nextLine != "" && !strings.ContainsAny(string(nextLine[j]), prohibitedChars)) {
				parsed, _ := strconv.ParseInt(line[index[i][0]:index[i][1]], 10, 64)
				sum+=int(parsed)
			}
		}		
	}

	return sum
}

func findGearRatios(previousLine string, line string, nextLine string) (int64) {
	re := regexp.MustCompile(`\*`)
	numbers := regexp.MustCompile(`(\.[0-9]\.)|(\d+){2,3}`)
	index := re.FindAllStringIndex(line, -1)
	var sum = int64(0)
	
	for i := 0; i < len(index); i++ {

		var a = []int64{}
		// Is part number on the left side
		if (index[i][0]-1 >= 0 && line[index[i][0]-1] != '.')  {
			parsed, _ := strconv.ParseInt(numbers.FindString(line[index[i][0]-3:]), 10, 64)
			a = append(a, int64(parsed))
		}
		
		// Is part number on the right side
		if  (index[i][1] < len(line) && line[index[i][1]] != '.') {
			parsed, _ := strconv.ParseInt(numbers.FindString(line[index[i][1]:]), 10, 64)
			a = append(a, int64(parsed))
		}

		// Is part number up
		if  (index[i][0]>= 0  && previousLine[index[i][0]] != '.') {
			parsed, _ := strconv.ParseInt(strings.ReplaceAll(numbers.FindString(previousLine[index[i][0]-2:index[i][1]+2]), ".", ""), 10, 64)
			a = append(a, int64(parsed))
		}

		// Is part number up left
		if  (index[i][0]>= 0  && previousLine[index[i][0]] == '.' && previousLine[index[i][0]-1] != '.') {
			parsed, _ := strconv.ParseInt(strings.ReplaceAll(numbers.FindString(previousLine[index[i][0]-3:index[i][0]+1]), ".", ""), 10, 64)
			a = append(a, int64(parsed))
		}

		// Is part number up right
		if  (index[i][0]>= 0  && previousLine[index[i][0]] == '.' && previousLine[index[i][0]+1] != '.') {
			parsed, _ := strconv.ParseInt(strings.ReplaceAll(numbers.FindString(previousLine[index[i][1]-1:index[i][1]+3]), ".", ""), 10, 64)
			a = append(a, int64(parsed))
		}

		// Is part number down
		if  (index[i][0]>= 0  && nextLine[index[i][0]] != '.') {
			parsed, _ := strconv.ParseInt(strings.ReplaceAll(numbers.FindString(nextLine[index[i][0]-2:index[i][1]+2]),".", ""), 10, 64)
			a = append(a, int64(parsed))
		}

		// Is part number down left
		if  (index[i][0]>= 0  && nextLine[index[i][0]] == '.' && nextLine[index[i][0]-1] != '.') {
			parsed, _ := strconv.ParseInt(strings.ReplaceAll(numbers.FindString(nextLine[index[i][0]-3:index[i][0]+1]), ".", ""), 10, 64)
			a = append(a, int64(parsed))
		}

		// Is part number down right
		if  (index[i][0]>= 0  && nextLine[index[i][0]] == '.' && nextLine[index[i][0]+1] != '.') {
			parsed, _ := strconv.ParseInt(strings.ReplaceAll(numbers.FindString(nextLine[index[i][1]-1:index[i][1]+3]), ".", ""), 10, 64)
			a = append(a, int64(parsed))
		}

		if (len(a) == 2){
			sum += a[0] * a[1]
		}
	}
	return sum
}
