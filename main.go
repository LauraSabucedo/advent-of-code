package main

import (
	"fmt"
	"utils"
)

func main() {
	fmt.Println("Day 1:")
    fmt.Println("\tPart 1:", utils.Sum("./inputs/day1.txt", `[0-9]`))
    fmt.Println("\tPart 2:", utils.Sum("./inputs/day1.txt", `[0-9]|one|two|three|four|five|six|seven|eight|nine`))

	fmt.Println("\nDay 2:")
	fmt.Println("\tPart 1:", utils.PossibleGamesSum("./inputs/day2.txt"))
	fmt.Println("\tPart 2:", utils.TotalPower("./inputs/day2.txt"))
}