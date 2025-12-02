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

func PartOne() {
	fmt.Println("Day 2, Part 1")

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

			var rangeBoundInts [2]int

			for index, value := range rangeBounds {
				rangeBoundInts[index], _ = strconv.Atoi(value)
			}

			for rangeBoundInts[0] < rangeBoundInts[1] {
				id := rangeBoundInts[0]
				// Check for double sequence
				idStr := strconv.Itoa(id)
				charCount := utf8.RuneCountInString(idStr)

				if charCount%2 == 0 {
					mid := charCount / 2

					left := idStr[0:mid]
					right := idStr[mid:]

					if left == right {
						invalidIds = append(invalidIds, id)
					}
				}

				rangeBoundInts[0]++
			}
		}
	}

	sum := 0
	for _, id := range invalidIds {
		sum += id
		fmt.Println(id)
	}

	fmt.Printf("Sum of Invalid IDs: %d", sum)
}
