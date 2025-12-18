package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	maxPos   = 100
	startPos = 50
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

	if err := runChristmasKey("Input.txt", part); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func runChristmasKey(filename string, part int) error {
	secret, position := 0, startPos

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open %s: %v", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // skip empty lines
		}
		if len(line) < 2 {
			return fmt.Errorf("invalid line: %q", line)
		}

		direction := line[0]
		distanceChar := line[1:]

		distance, err := strconv.Atoi(distanceChar)
		if err != nil {
			return fmt.Errorf("invalid number in line %q: %v", line, err)
		}

		switch part {
		case 1:
			secret, position = stepPart1(secret, position, direction, distance)
		case 2:
			secret, position = stepPart2(secret, position, direction, distance)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	fmt.Println("The secret is:", secret)
	fmt.Println("Final position:", position)
	return nil
}

// Part 1 logic for one instruction
func stepPart1(secret, position int, direction byte, distance int) (int, int) {
	// normalize distance
	if distance > maxPos {
		distance = distance % maxPos
	}
	if distance == maxPos {
		return secret, position // skip this instruction
	}

	switch direction {
	case 'L':
		position -= distance
		if position == 0 {
			secret++
		} else if position < 0 {
			position += maxPos
		}
	case 'R':
		position += distance
		if position == maxPos {
			secret++
		} else if position > maxPos {
			position -= maxPos
		}
	}
	return secret, position
}

// Part 2 logic for one instruction
func stepPart2(secret, position int, direction byte, distance int) (int, int) {
	// normalize distance and count hundreds
	if distance > maxPos {
		secret += distance / maxPos
		distance = distance % maxPos
	}
	if distance == maxPos {
		secret++
		return secret, position
	}

	switch direction {
	case 'L':
		previousPosition := position
		position -= distance

		if position == 0 {
			secret++
		} else if position < 0 {
			if previousPosition != 0 {
				secret++
			}
			position += maxPos
		}

	case 'R':
		position += distance
		if position == maxPos {
			secret++
			position = 0
		} else if position > maxPos {
			secret++
			position -= maxPos
		}
	}

	return secret, position
}
