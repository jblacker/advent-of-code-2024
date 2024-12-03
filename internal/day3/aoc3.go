package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var mulPattern = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
var mulAndConditionalPattern = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

const (
	DO   = "do()"
	DONT = "don't()"
)

func Solve(path string, useConditionals bool) {
	var conditionalTitle string
	sum := 0
	rawData, err := parse(path)
	if err != nil {
		panic(err)
	}
	if !useConditionals {
		data := mulPattern.FindAllStringSubmatch(rawData, -1)
		for _, match := range data {
			product, err := processCommand(match)
			if err != nil {
				panic(err)
			}
			sum += product
		}
	} else {
		conditionalTitle = " (with conditionals)"
		data := mulAndConditionalPattern.FindAllStringSubmatch(rawData, -1)
		executeCmd := true
		for _, match := range data {
			switch match[0] {
			case DO:
				executeCmd = true
			case DONT:
				executeCmd = false
			default:
				if !executeCmd {
					continue
				}
				product, err := processCommand(match)
				if err != nil {
					panic(err)
				}
				sum += product
			}
		}
	}

	fmt.Printf("Summation of Products%s: %d\n", conditionalTitle, sum)
}

func parse(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func processCommand(cmd []string) (int, error) {
	if len(cmd) != 3 {
		return 0, fmt.Errorf("unexpected slice length of %d", len(cmd))
	}

	element1, err := strconv.Atoi(cmd[1])
	if err != nil {
		return 0, err
	}
	element2, err := strconv.Atoi(cmd[2])
	if err != nil {
		return 0, err
	}

	return element1 * element2, nil
}
