package day1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var linePattern = regexp.MustCompile(`\s+`)

func Solve(path string) {
	list1, list2, err := parse(path)
	if err != nil {
		panic(err)
	}
	slices.Sort(list1)
	slices.Sort(list2)

	if len(list1) != len(list2) {
		panic(fmt.Errorf("list 1 has %d items & list 2 has %d items -- cannot compare", len(list1), len(list2)))
	}
	diffs := make([]int, 0, len(list1))
	for i := range len(list1) {
		item1 := list1[i]
		item2 := list2[i]
		if item1 > item2 {
			diffs = append(diffs, item1-item2)
		} else if item1 < item2 {
			diffs = append(diffs, item2-item1)
		} else {
			diffs = append(diffs, 0)
		}
	}
	totalDiff := 0
	for _, d := range diffs {
		totalDiff += d
	}
	fmt.Printf("Total Difference Score: %d\n", totalDiff)
	list2CardMap := cardinalityMap(list2)
	similarityScore := make([]int, 0, len(list1))
	for _, item := range list1 {
		if val, ok := list2CardMap[item]; ok {
			similarityScore = append(similarityScore, item*val)
		} else {
			similarityScore = append(similarityScore, 0)
		}
	}
	totalSimilarity := 0
	for _, s := range similarityScore {
		totalSimilarity += s
	}
	fmt.Printf("Total Similarity Score: %d\n", totalSimilarity)

}

func parse(path string) ([]int, []int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	list1 := make([]int, 0, 10)
	list2 := make([]int, 0, 10)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = linePattern.ReplaceAllString(line, " ")
		parsedLine := strings.Split(line, " ")
		item1, err := strconv.Atoi(parsedLine[0])
		if err != nil {
			return nil, nil, err
		}
		list1 = append(list1, item1)
		item2, err := strconv.Atoi(parsedLine[1])
		if err != nil {
			return nil, nil, err
		}
		list2 = append(list2, item2)
	}
	return list1, list2, nil
}

func cardinalityMap(list []int) map[int]int {
	cardMap := make(map[int]int)
	for _, v := range list {
		val, ok := cardMap[v]
		if !ok {
			cardMap[v] = 1
		} else {
			val++
			cardMap[v] = val

		}
	}
	return cardMap
}
