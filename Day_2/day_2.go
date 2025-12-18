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

func invalidIdSumAlgorithmTwo(currentSum, startId, endId int) int {

	for id := startId; id <= endId; id++ {
		stringId := strconv.Itoa(id)

		//if all the numbers are the same
		if idHasTheSameDigits(stringId) {
			fmt.Printf("Value found between %d and %d, it's is %d\n", startId, endId, id)
			currentSum += id
		} else {
			//if it's bigger then 4
			if len(stringId) >= 4 {
				//for the lenght divided by 2
				for size := 2; size <= len(stringId)/2; size++ {
					//if the size is divisible by i
					if len(stringId)%size == 0 {
						baseChar := stringId[0:size]
						ok := true
						for j := size; j < len(stringId); j += size {
							if stringId[j:j+size] != baseChar {
								ok = false
								break
							}
						}
						if ok {
							fmt.Printf("Value found between %d and %d, it's is %d\n", startId, endId, id)
							currentSum += id
						}
					}
				}
			}
		}
	}

	return currentSum
}

func idHasTheSameDigits(stringId string) bool {

	for i := 1; i < len(stringId); i++ {
		if stringId[i] != stringId[0] {
			return false
		}
	}

	return true
}
