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

	fmt.Print("Do you want to run part 1 or 2? ")
	fmt.Scan(&part)

	if part != 1 && part != 2 {
		log.Fatalf("invalid part: %d (must be 1 or 2)", part)
	}

	if err := runInvalidIdDetection("Input.txt", part); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func runInvalidIdDetection(filename string, part int) error {

	invalidSum := 0

	f, err := os.Open(filename)

	if err != nil {
		return fmt.Errorf("Failed to open file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		return fmt.Errorf("File is empty")
	}
	line := strings.TrimSpace(scanner.Text())

	pairs := strings.Split(line, ",")

	for _, p := range pairs {
		parts := strings.SplitN(strings.TrimSpace(p), "-", 2)
		if len(parts) != 2 {
			return fmt.Errorf("Invalid pair: %s", p)
		}

		startId, err := strconv.Atoi(parts[0])
		if err != nil {
			return fmt.Errorf("invalid number: %s", parts[0])
		}

		endId, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("Invalid number: %s", parts[1])
		}

		switch part {
		case 1:
			invalidSum = invalidIdSumAlgorithmOne(invalidSum, startId, endId)
		case 2:
			invalidSum = invalidIdSumAlgorithmTwo(invalidSum, startId, endId)
		}
	}

	fmt.Printf("The final addition is equal to: %d\n", invalidSum)
	return nil
}

func invalidIdSumAlgorithmOne(currentSum, startId, endId int) int {

	for id := startId; id <= endId; id++ {
		stringId := strconv.Itoa(id)

		if len(stringId)%2 == 0 {

			mid := len(stringId) / 2
			left := stringId[:mid]
			right := stringId[mid:]

			if left == right {
				currentSum += id
			}
		}
	}

	return currentSum
}

func isInvalidPartTwo(id int) bool {
	s := strconv.Itoa(id)
	n := len(s)

	for size := 1; size <= n/2; size++ {

		if n%size != 0 {
			continue
		}

		repeats := n / size
		if repeats < 2 {
			continue
		}

		block := s[:size]
		ok := true
		for i := size; i < n; i += size {
			if s[i:i+size] != block {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}

func invalidIdSumAlgorithmTwo(currentSum, startId, endId int) int {
	for id := startId; id <= endId; id++ {
		if isInvalidPartTwo(id) {
			currentSum += id
		}
	}
	return currentSum
}
