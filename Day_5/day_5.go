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

	if err := runFridgeScript("Input.txt", part); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func runFridgeScript(filename string, part int) error {

	freshProducts := 0

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open %s: %v", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fridge := strings.TrimSpace(scanner.Text())
		if fridge == "" {
			continue // skip empty lines
		}
		if len(fridge) < 2 {
			return fmt.Errorf("invalid line: %q", fridge)
		}

		switch part {
		case 1:
			freshProducts = stepPart1()
		case 2:
			freshProducts = stepPart2()
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	fmt.Println("The final amount of fresh food is:", freshProducts)
	return nil
}

func stepPart1() int {
	return 1
}

func stepPart2() int {
	return 1
}
