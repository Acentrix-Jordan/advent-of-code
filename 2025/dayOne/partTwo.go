package dayOne

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func PartTwo() {
	fmt.Println("Day 1, Part 2")

	path, _ := os.Getwd()
	file, err := os.Open(path + "/2025/stubs/dayOne.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count := 0
	currentValue := 50

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		movement := scanner.Text()
		strValue := movement[1:]

		intValue, err := strconv.Atoi(strValue)
		fullCycleSize := 100

		if err != nil {
			log.Fatal(err)
		}

		if strings.HasPrefix(movement, "L") {
			startsOnZero := false

			if currentValue == 0 {
				startsOnZero = true
			}

			// Remove full cycles
			currentValue = currentValue - (intValue % fullCycleSize)

			// If the value starts on zero and doesn't do a full rotation, we were doing duplicating passes
			if startsOnZero == false {
				// Passed zero, increment count
				if currentValue < 0 {
					count++

					// Adjust to valid value
					currentValue += 100
				}
			} else {
				if currentValue < 0 {
					// Adjust to valid value
					currentValue += 100
				}
			}
		} else {
			// Remove full cycles
			currentValue = currentValue + (intValue % fullCycleSize)

			// Passed zero, increment count
			if currentValue > 100 {
				count++

				// Adjust to valid value
				currentValue -= 100
			}

			if currentValue == 100 {
				// 100 = 0
				currentValue = 0
			}
		}

		// Each full rotation passes zero, once increment count
		// R1000 would pass 0, 10x
		fullRotations := intValue / fullCycleSize
		count += fullRotations

		// Check if final value lands on 0
		if currentValue == 0 || currentValue == 100 {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The safe landed on '0' %d times.", count)
}
