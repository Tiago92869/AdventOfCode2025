package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	if err := runRollsCode("Input_test.txt", part); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func runRollsCode(filename string, part int) error {

	rollsCode := 0

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
			rollsCode = stepPart1(rollsCode, batterie)
		case 2:
			rollsCode = stepPart2(rollsCode, batterie)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	fmt.Println("The total amount of rolls that can be accessed is:", rollsCode)
	return nil
}

func stepPart1(rollsCode int, roll string) int {
	return 1
}

func stepPart2(rollsCode int, roll string) int {
	return 2
}
