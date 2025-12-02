package dayTwo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func PartTwo() {
	fmt.Println("Day 2, Part 2")

	path, _ := os.Getwd()
	file, err := os.Open(path + "/2025/stubs/dayTwo.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var invalidIds []int

	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")

		for _, value := range ranges {
			rangeBounds := strings.Split(value, "-")

			start, _ := strconv.Atoi(rangeBounds[0])
			end, _ := strconv.Atoi(rangeBounds[1])

			ids := checkForInvalidIds(start, end)
			invalidIds = append(invalidIds, ids...)
		}
	}

	sum := 0
	for _, id := range invalidIds {
		sum += id
	}

	// Answer: 20077272987
	fmt.Printf("Sum of Invalid IDs: %d", sum)
}

func checkForInvalidIds(rangeStart int, rangeEnd int) []int {
	var invalidIds []int

	for rangeStart <= rangeEnd {
		id := checkId(rangeStart)

		if id != -1 {
			invalidIds = append(invalidIds, id)
		}

		rangeStart++
	}

	return invalidIds
}

/*
An ID is invalid if it is made only of some sequence of digits repeated at least twice.
So, 12341234 (1234 two times), 123123123 (123 three times), 1212121212 (12 five times),
and 1111111 (1 seven times) are all invalid IDs.
*/
func checkId(id int) int {
	idStr := strconv.Itoa(id)
	charCount := utf8.RuneCountInString(idStr)
	// Pattern Size Bounds: 1 -> (StrLen / 2)
	upperBound := charCount / 2

	// Offset by 1 for valid strings
	for i := 1; i < upperBound+1; i++ {
		pattern := idStr[:i]
		patternCount := utf8.RuneCountInString(pattern)

		if (charCount % patternCount) > 0 {
			continue
		}

		repetitionCount := charCount / patternCount
		comparison := strings.Repeat(pattern, repetitionCount)

		if idStr == comparison {
			return id
		}

	}

	return -1
}
