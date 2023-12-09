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

	fmt.Println("\nDay 3:")
	day3p1, day3p2 := utils.SumPartNumbers("./inputs/day3.txt", `(\d+)`, ".1234567890")
	fmt.Println("\tPart 1:", day3p1)
	fmt.Println("\tPart 2:", day3p2)

	fmt.Println("\nDay 4:")
	fmt.Println("\tPart 1:", utils.TotalPoints("./inputs/day4.txt"))
	fmt.Println("\tPart 2:", utils.TotalScratchCards("./inputs/day4.txt"))

	fmt.Println("\nDay 6:")
	day6p1, day6p2 := utils.WaysToBeatRecord("./inputs/day6.txt")
	fmt.Println("\tPart 1:", day6p1)
	fmt.Println("\tPart 2:", day6p2)

	fmt.Println("\nDay 7:")
	//utils.TotalWinnings("./inputs/day7.txt")

	fmt.Println("\nDay 8:")
	fmt.Println("\tPart 1:", utils.Steps("./inputs/day8.txt"))
	fmt.Println("\tPart 2:", utils.GhostSteps("./inputs/day8.txt"))
}
