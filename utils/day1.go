package utils

import (
	"bufio"
	"fmt"
	"os"
    "regexp"
    "strconv"
)

func Sum(path string, regex string) (int) {
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
    
	// Iterate through each line
	for scanner.Scan() {
        // Get the line
        line := scanner.Text()

        // Find all matches in the line and store them in a slice
        numbers := findAllMatches(line, regex)
        
        // Concat first and last number
        var number string
        number = translate(numbers[0]) + translate(numbers[len(numbers)-1])

        // Parse string to int and add to sum
        parsed, _ := strconv.ParseInt(number, 10, 64)
        sum+=int(parsed)
	}
    
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
	}

    // Close the file when main() ends
    defer file.Close()

    return sum
}

func findAllMatches(line, regex string) []string {
	re := regexp.MustCompile(regex)
	var matches []string

	for d := range line {
		// Find the first match in the remaining part of the line
		match := re.FindString(line[d:])

        // If match is found, add it to the matches slice
        if len(match) > 0 {
            matches = append(matches, match)
        }
	}

	return matches
}



func translate(number string) (string) {
    translations := map[string]string {
        "zero": "0",
        "one": "1",
        "two": "2",
        "three": "3",
        "four": "4",
        "five": "5",
        "six": "6",
        "seven": "7",
        "eight": "8",
        "nine": "9",
    }
    
    if translated, ok := translations[number]; ok {
		return translated
	} else {
		return number
	}
}