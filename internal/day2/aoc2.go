package day2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	SAFE   = "safe"
	UNSAFE = "unsafe"
	INC    = "inc"
	DEC    = "dec"
)

func Solve(path string, useDampener bool) {
	reports, err := parse(path)
	if err != nil {
		panic(err)
	}

	safeCount := 0
	for _, r := range reports {
		if r.checkSafety() == SAFE {
			safeCount++
			continue
		}

		if useDampener && r.safetyWithDampener() == SAFE {
			safeCount++
		}
	}

	var dampenerString string
	if useDampener {
		dampenerString = " (with dampener)"
	}
	fmt.Printf("Safe Report Count%s: %d\n", dampenerString, safeCount)
}

type Report struct {
	levels []int
	status string
}

func (r *Report) safetyWithDampener() string {
	fmt.Println("Dampener Check!")
	for i := 0; i < len(r.levels); i++ {
		levelCopy := slices.Clone(r.levels)
		levelCopy = append(levelCopy[:i], levelCopy[i+1:]...)
		if checkSafety(levelCopy) == SAFE {
			fmt.Printf("Dampener indicates safety removing %d!\n", i)
			return SAFE
		}
	}

	fmt.Println("Still unsafe with dampener")
	return UNSAFE
}

func (r *Report) print() {
	stringifiedLevels := make([]string, 0, len(r.levels))
	for _, l := range r.levels {
		stringifiedLevels = append(stringifiedLevels, strconv.FormatInt(int64(l), 10))
	}
	fmt.Printf("%s - %s\n", strings.Join(stringifiedLevels, ", "), r.status)
}

func checkSafety(levels []int) string {
	levelCount := len(levels)
	changes := make([]int, 0, levelCount-1)
	for i := 1; i < levelCount; i++ {
		diff := levels[i] - levels[i-1]
		if diff != 0 {
			changes = append(changes, diff)
		} else {
			return UNSAFE
		}
	}

	var direction string
	for i, d := range changes {
		if i == 0 {
			direction = getDirection(d)
		} else if i > 0 && direction != getDirection(d) {
			return UNSAFE
		}

		if math.Abs(float64(d)) > 3 {
			return UNSAFE
		}
	}

	return SAFE
}

func (r *Report) checkSafety() string {
	safety := checkSafety(r.levels)
	r.status = safety
	return safety
}

func getDirection(diff int) string {
	if diff < 0 {
		return DEC
	} else {
		return INC
	}
}

func parse(path string) ([]*Report, error) {
	reports := make([]*Report, 0, 10)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parsedLine := strings.Split(line, " ")
		report := make([]int, 0, len(parsedLine))
		for _, v := range parsedLine {
			val, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			report = append(report, val)
		}
		reports = append(reports, &Report{levels: report})
	}
	return reports, nil
}
