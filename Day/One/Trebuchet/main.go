package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	val := 0
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf("Error encountered: %s\n", err.Error())
		os.Exit(0)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Part One
		val += computeLineValue(scanner.Text())

		// Part two
		// val += computeLineValueWithWords(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error encountered: %s\n", err.Error())
		os.Exit(0)
	}

	fmt.Printf("Answer: %v\n", val)
}

// Assumes valid input
func computeLineValue(line string) int {
	p1 := 0
	p2 := len(line) - 1

	var firstLeft int = -1
	var firstRight int = -1
	var lastFound int = -1

	for i := 0; i < (len(line)/2)+1; i++ {
		if runeVal := rune(line[p1]); unicode.IsDigit(runeVal) {
			numVal := int(runeVal - '0')
			if firstLeft == -1 {
				firstLeft = numVal
			}
			lastFound = numVal
		}
		p1 += 1

		if runeVal := rune(line[p2]); unicode.IsDigit(runeVal) {
			numVal := int(runeVal - '0')
			if firstRight == -1 {
				firstRight = numVal
			}
			lastFound = numVal
		}
		p2 -= 1

		if firstLeft > -1 && firstRight > -1 {
			break
		}
	}

	if firstLeft == -1 {
		firstLeft = lastFound
	} else if firstRight == -1 {
		firstRight = lastFound
	}

	if value := firstLeft*10 + firstRight; value > 0 {
		fmt.Printf("Line: '%s', Value: '%v'\n", line, value)
		return value
	} else {
		fmt.Printf("ERROR: Did not find value for line '%s'\n", line)
		return 0
	}
}

// Assumes valid input
func computeLineValueWithWords(line string) int {
	p1 := 0
	p2 := len(line) - 1

	var firstLeft int = -1
	var firstRight int = -1
	var lastFound int = -1

	var forwardNumCheckers []*numberChecker = make([]*numberChecker, 0)
	var backwardNumCheckers []*numberChecker = make([]*numberChecker, 0)

	for i := 0; i < (len(line)/2)+4; i++ {
		if runeVal := rune(line[p1]); unicode.IsDigit(runeVal) {
			numVal := int(runeVal - '0')
			if firstLeft == -1 {
				firstLeft = numVal
			}
			lastFound = numVal
		} else {
			for _, checker := range forwardNumCheckers {
				if checker.isNext(runeVal) {
					if checker.isFinished() && firstLeft == -1 {
						firstLeft = checker.intValue
					}
				}
			}

			forwardNumCheckers = append(forwardNumCheckers, GetNumCheckers(runeVal, false)...)
		}
		p1 += 1

		if runeVal := rune(line[p2]); unicode.IsDigit(runeVal) {
			numVal := int(runeVal - '0')
			if firstRight == -1 {
				firstRight = numVal
			}
			lastFound = numVal
		} else {
			for _, checker := range backwardNumCheckers {
				if checker.isNext(runeVal) {
					if checker.isFinished() && firstRight == -1 {
						firstRight = checker.intValue
					}
				} else {
					backwardNumCheckers
					delete(backwardNumCheckers, checker)
				}
			}

			backwardNumCheckers = append(backwardNumCheckers, GetNumCheckers(runeVal, true)...)
		}
		p2 -= 1

		if firstLeft > -1 && firstRight > -1 {
			break
		}
	}

	if firstLeft == -1 {
		firstLeft = lastFound
	} else if firstRight == -1 {
		firstRight = lastFound
	}

	if value := firstLeft*10 + firstRight; value > 0 {
		fmt.Printf("Line: '%s', Value: '%v'\n", line, value)
		return value
	} else {
		fmt.Printf("ERROR: Did not find value for line '%s'\n", line)
		return 0
	}
}

type numberChecker struct {
	intValue       int
	backward       bool
	chars          []rune
	currentPointer int
}

func (c *numberChecker) isNext(r rune) bool {
	adjustedPointer := c.currentPointer
	if c.backward {
		adjustedPointer = len(c.chars) - (c.currentPointer + 1)
	}

	if c.currentPointer < len(c.chars) && c.chars[adjustedPointer] == r {
		c.currentPointer++
		return true
	}

	return false
}

func (c *numberChecker) isFinished() bool {
	return c.currentPointer == len(c.chars)+1
}

func (c *numberChecker) reset() {
	c.currentPointer = 0
}

func GetNumCheckers(r rune, backward bool) []*numberChecker {
	numberCheckers := make([]*numberChecker, 0)

	if !backward {
		if r == 'o' {
			numberCheckers = append(numberCheckers, newNumberChecker(1, backward))
		} else if r == 't' {
			numberCheckers = append(numberCheckers, newNumberChecker(2, backward), newNumberChecker(3, backward))
		} else if r == 'f' {
			numberCheckers = append(numberCheckers, newNumberChecker(4, backward), newNumberChecker(5, backward))
		} else if r == 's' {
			numberCheckers = append(numberCheckers, newNumberChecker(6, backward), newNumberChecker(7, backward))
		} else if r == 'e' {
			numberCheckers = append(numberCheckers, newNumberChecker(8, backward))
		} else if r == 'n' {
			numberCheckers = append(numberCheckers, newNumberChecker(9, backward))
		}
	} else {
		if r == 'e' {
			numberCheckers = append(numberCheckers, newNumberChecker(1, backward), newNumberChecker(3, backward), newNumberChecker(5, backward), newNumberChecker(9, backward))
		} else if r == 'o' {
			numberCheckers = append(numberCheckers, newNumberChecker(2, backward))
		} else if r == 'r' {
			numberCheckers = append(numberCheckers, newNumberChecker(4, backward))
		} else if r == 'x' {
			numberCheckers = append(numberCheckers, newNumberChecker(6, backward))
		} else if r == 'n' {
			numberCheckers = append(numberCheckers, newNumberChecker(7, backward))
		} else if r == 't' {
			numberCheckers = append(numberCheckers, newNumberChecker(8, backward))
		}
	}

	return numberCheckers
}

func newNumberChecker(num int, backward bool) *numberChecker {
	numberChecker := &numberChecker{
		intValue:       num,
		currentPointer: 1,
		backward:       backward,
	}
	for _, r := range getStringValue(num) {
		numberChecker.chars = append(numberChecker.chars, r)
	}

	return numberChecker
}

func getStringValue(intValue int) string {
	switch intValue {
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	default:
		return ""
	}
}
