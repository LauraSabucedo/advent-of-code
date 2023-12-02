package main

import (
	"bufio"
	"fmt"
	"os"
    "regexp"
    "strconv"
)

func main() {
	// Specify the path to the file you want to read
	filePath := "./day1.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		// Handle the error
		fmt.Println("Error opening file:", err)
		return
	}
    
	// Create a Scanner to read the file line by line
	scanner := bufio.NewScanner(file)

    re := regexp.MustCompile(`[0-9]`)
    var sum = 0
    
	// Iterate through each line
	for scanner.Scan() {
        line := scanner.Text()

        numbers := re.FindAllString(line, -1)

        var number string

        number = numbers[0] + numbers[len(numbers)-1]

        fmt.Println(number)

        parsed, _ := strconv.ParseInt(number, 10, 64)

        sum+=int(parsed)
	}

    fmt.Println("Sum:", sum)
    
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
	}

    // Close the file when main() ends
    defer file.Close()
}
