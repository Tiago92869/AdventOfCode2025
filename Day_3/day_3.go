package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var part int

	fmt.Print("Do you want to run Part 1 or 2? ")
	if _, err := fmt.Scan(&part); err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	if part != 1 && part != 2 {
		log.Fatalf("invalid part: %d (must be 1 or 2)", part)
	}

	if err := runBatteriesCode("Input.txt", part); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func runBatteriesCode(filename string, part int) error {

	batteriesCode := 0

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open %s: %v", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		batterie := strings.TrimSpace(scanner.Text())
		if batterie == "" {
			continue // skip empty lines
		}
		if len(batterie) < 2 {
			return fmt.Errorf("invalid line: %q", batterie)
		}

		switch part {
		case 1:
			batteriesCode = stepPart1(batteriesCode, batterie)
		case 2:
			batteriesCode = stepPart2(batteriesCode, batterie)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	fmt.Println("The final code is:", batteriesCode)
	return nil
}

func stepPart1(batteriesCode int, battery string) int {
	// Guard against too-short input
	if len(battery) < 2 {
		return batteriesCode
	}

	// Helper: turn two digit bytes into an int (e.g. '3','7' -> 37)
	pairValue := func(a, b byte) int {
		return int(a-'0')*10 + int(b-'0')
	}

	first, last := battery[0], battery[1]

	currentSum := pairValue(first, last)
	currentSumString := battery[:2]

	// If there are more than two digits, try to improve the pair
	for i := 2; i < len(battery); i++ {
		d := battery[i]

		// Two possible pairs: (first + d) and (last + d)
		tempSumFirst := pairValue(first, d)
		tempSumSecond := pairValue(last, d)

		// Choose the bigger of the two candidate pairs
		if tempSumFirst > tempSumSecond {
			if tempSumFirst > currentSum {
				currentSum = tempSumFirst
				currentSumString = string([]byte{first, d})
				last = d
			}
		} else {
			if tempSumSecond > currentSum {
				currentSum = tempSumSecond
				currentSumString = string([]byte{last, d})
				first, last = last, d
			}
		}
	}

	fmt.Println("New code found:", currentSumString)
	batteriesCode += currentSum
	return batteriesCode
}

func stepPart2(batteriesCode int, battery string) int {
	// Guard against too-short input
	if len(battery) < 12 {
		fmt.Printf("stepPart2: skipping battery %q (len=%d < 12)\n", battery, len(battery))
		return batteriesCode
	}

	// do a cicle to find the biggest 12 digits, find the first biggest number and then do the same starting from the position of the last big number
	var finalValue [12]byte
	finalValue[0] = battery[0]
	missing, previousPosition := 11, 0

	for turn := 0; turn < 12; turn++ {

		// create a exception when the lenght left is equal to the missing numbers
		if len(battery)-previousPosition+1 == missing {
			for i := 0; i <= missing; i++ {
				finalValue[turn] = battery[previousPosition+i]
				previousPosition++
			}
		}

		finalValue[turn] = battery[previousPosition+1]
		for i := previousPosition + 1; i < len(battery)-missing; i++ {

			if battery[i] > finalValue[turn] {
				finalValue[turn], previousPosition = battery[i], i
			}
		}
		missing--
	}

	finalValueString := string(finalValue[:])
	finalValueInt, _ := strconv.Atoi(finalValueString)
	batteriesCode += finalValueInt
	return batteriesCode
}
