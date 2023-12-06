package main

import (
	"bufio"
	"fmt"
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

	ans := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		splitArray := strings.SplitN(line, ":", 2)
		if gameNumber, err := strconv.Atoi(strings.TrimLeft(splitArray[0], "Game ")); err != nil {
			panic(err)
		} else {
			if matchingGame(splitArray[1]) {
				ans += gameNumber
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Solution: %v\n", ans)
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
