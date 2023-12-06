package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type color int

const (
	red   = 0
	green = 1
	blue  = 2
	other = 3
)

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	// ans := partOne(inputFile)
	ans := partTwo(inputFile)

	fmt.Printf("Solution: %v\n", ans)
}

func partOne(inputFile *os.File) int {
	ans := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		splitArray := strings.SplitN(line, ":", 2)
		if gameNumber, err := strconv.Atoi(strings.TrimLeft(splitArray[0], "Game ")); err != nil {
			panic(err)
		} else {
			if matchingGame(splitArray[1]) {
				fmt.Printf("%s\n", splitArray[0])
				ans += gameNumber
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return ans
}

func partTwo(inputFile *os.File) int {
	ans := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		splitArray := strings.SplitN(line, ":", 2)
		if _, err := strconv.Atoi(strings.TrimLeft(splitArray[0], "Game ")); err != nil {
			panic(err)
		} else {
			if power := gamePower(splitArray[1]); power > 0 {
				fmt.Printf("%s, %s --- %v\n", splitArray[0], splitArray[1], power)
				ans += power
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return ans
}

// matchingGame outputs a bool, indicating whether the game would have been possible if the contents of the bag were only 12 red cubes, 13 green cubes, and 14 blue cubes
func matchingGame(input string) bool {
	for _, round := range strings.Split(input, ";") {
		for _, entry := range strings.Split(round, ",") {
			if !checkLegalEntry(parseColorAndNumber(strings.TrimSpace(entry))) {
				return false
			}
		}
	}

	return true
}

func parseColorAndNumber(entry string) (color, int) {
	splitSlice := strings.SplitN(entry, " ", 2)

	if num, err := strconv.Atoi(splitSlice[0]); err != nil {
		panic(err)
	} else {
		return parseColorFromString(strings.ToLower(splitSlice[1])), num
	}
}

func parseColorFromString(colorString string) color {
	switch colorString {
	case "red":
		return red
	case "blue":
		return blue
	case "green":
		return green
	default:
		return other
	}
}

func checkLegalEntry(col color, count int) bool {
	switch col {
	case red:
		return count <= 12
	case green:
		return count <= 13
	case blue:
		return count <= 14
	default:
		return false
	}
}

// gamePower is like 'matchingGame()', however it outputs a power value for each color in place of a single bool
func gamePower(input string) int {
	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	for _, round := range strings.Split(input, ";") {
		for _, entry := range strings.Split(round, ",") {
			c, n := parseColorAndNumber(strings.TrimSpace(entry))
			if c == red {
				maxRed = int(math.Max(float64(maxRed), float64(n)))
			} else if c == green {
				maxGreen = int(math.Max(float64(maxGreen), float64(n)))
			} else if c == blue {
				maxBlue = int(math.Max(float64(maxBlue), float64(n)))
			}
		}
	}

	fmt.Printf("\tRed: %v, Blue: %v, Green: %v\n", maxRed, maxBlue, maxGreen)

	return maxRed * maxBlue * maxGreen
}
