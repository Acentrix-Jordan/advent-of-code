package dayOne

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func PartOne() {
	fmt.Println("Day 1, Part 1")

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

		if err != nil {
			log.Fatal(err)
		}

		if strings.HasPrefix(movement, "L") {
			currentValue = currentValue - (intValue % 100)

			if currentValue < 0 {
				currentValue += 100
			}
		} else {
			currentValue = currentValue + (intValue % 100)

			if currentValue > 99 {
				currentValue -= 100
			}
		}

		if currentValue == 0 {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The safe landed on '0' %d times.", count)
}
