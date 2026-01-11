package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	if err := runRollsCode("Input.txt", part); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func runRollsCode(filename string, part int) error {

	totalAmmount := 0

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open %s: %v", filename, err)
	}
	defer file.Close()

	var stock [][]byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		stock = append(stock, []byte(line))

		if err := scanner.Err(); err != nil {
			log.Fatalf("error reading file: %v", err)
		}
	}

	switch part {
	case 1:
		totalAmmount = stepPart1(stock)
	case 2:
		totalAmmount = stepPart2(stock)
	}

	fmt.Println("The total amount of rolls that can be accessed is:", totalAmmount)
	return nil
}

func stepPart1(stock [][]byte) int {

	finalAmmount := 0
	numberOfRows := len(stock)
	numberOfCols := len(stock[0])

	for r := 0; r < numberOfRows; r++ {
		for c := 0; c < numberOfCols; c++ {

			if stock[r][c] != '@' {
				continue
			}

			//if its the first row
			switch r {
			case 0:
				if firstRowTotal(r, c, stock, numberOfCols) < 4 {
					fmt.Println("1New value found at: row ", r, " column", c)
					finalAmmount++
				}
			case numberOfRows - 1:
				if lastRowTotal(r, c, stock, numberOfCols) < 4 {
					fmt.Println("2New value found at: row ", r, " column", c)
					finalAmmount++
				}
			default:
				if anyOtherRowTotal(r, c, stock, numberOfCols) < 4 {
					fmt.Println("3New value found at: row ", r, " column", c)
					finalAmmount++
				}
			}
		}
	}
	return finalAmmount
}

func firstRowTotal(row int, column int, stock [][]byte, numberOfCols int) int {

	paperRolls := 0

	//if it's first column
	switch column {
	case 0:

		if stock[row][column+1] == '@' || stock[row][column+1] == 'x' {
			paperRolls++
		}

		if stock[row+1][column+1] == '@' || stock[row+1][column+1] == 'x' {
			paperRolls++
		}

		if stock[row+1][column] == '@' || stock[row+1][column] == 'x' {
			paperRolls++
		}
	case numberOfCols - 1:

		if stock[row][column-1] == '@' || stock[row][column-1] == 'x' {
			paperRolls++
		}

		if stock[row+1][column-1] == '@' || stock[row][column-1] == 'x' {
			paperRolls++
		}

		if stock[row+1][column] == '@' || stock[row+1][column] == 'x' {
			paperRolls++
		}
	default:
		if stock[row][column-1] == '@' || stock[row][column-1] == 'x' {
			paperRolls++
		}

		if stock[row+1][column-1] == '@' || stock[row+1][column-1] == 'x' {
			paperRolls++
		}

		if stock[row+1][column] == '@' || stock[row+1][column] == 'x' {
			paperRolls++
		}

		if stock[row][column+1] == '@' || stock[row][column+1] == 'x' {
			paperRolls++
		}

		if stock[row+1][column+1] == '@' || stock[row+1][column+1] == 'x' {
			paperRolls++
		}
	}

	return paperRolls
}

func lastRowTotal(row int, column int, stock [][]byte, numberOfCols int) int {

	paperRolls := 0

	//if it's first column
	switch column {
	case 0:

		if stock[row-1][column] == '@' || stock[row-1][column] == 'x' {
			paperRolls++
		}

		if stock[row-1][column+1] == '@' || stock[row-1][column+1] == 'x' {
			paperRolls++
		}

		if stock[row][column+1] == '@' || stock[row][column+1] == 'x' {
			paperRolls++
		}
	case numberOfCols - 1:

		if stock[row][column-1] == '@' || stock[row][column-1] == 'x' {
			paperRolls++
		}

		if stock[row-1][column-1] == '@' || stock[row-1][column-1] == 'x' {
			paperRolls++
		}

		if stock[row-1][column] == '@' || stock[row-1][column] == 'x' {
			paperRolls++
		}
	default:
		if stock[row][column-1] == '@' || stock[row][column-1] == 'x' {
			paperRolls++
		}

		if stock[row-1][column-1] == '@' || stock[row-1][column-1] == 'x' {
			paperRolls++
		}

		if stock[row-1][column] == '@' || stock[row-1][column] == 'x' {
			paperRolls++
		}

		if stock[row-1][column+1] == '@' || stock[row-1][column+1] == 'x' {
			paperRolls++
		}

		if stock[row][column+1] == '@' || stock[row][column+1] == 'x' {
			paperRolls++
		}
	}

	return paperRolls
}

func anyOtherRowTotal(row int, column int, stock [][]byte, numberOfCols int) int {

	paperRolls := 0

	//if it's first column
	switch column {
	case 0:

		if stock[row-1][column] == '@' || stock[row-1][column] == 'x' {
			paperRolls++
		}

		if stock[row-1][column+1] == '@' || stock[row-1][column+1] == 'x' {
			paperRolls++
		}

		if stock[row][column+1] == '@' || stock[row][column+1] == 'x' {
			paperRolls++
		}

		if stock[row+1][column+1] == '@' || stock[row+1][column+1] == 'x' {
			paperRolls++
		}

		if stock[row+1][column] == '@' || stock[row+1][column] == 'x' {
			paperRolls++
		}
	case numberOfCols - 1:

		if stock[row-1][column] == '@' || stock[row-1][column] == 'x' {
			paperRolls++
		}

		if stock[row-1][column-1] == '@' || stock[row-1][column-1] == 'x' {
			paperRolls++
		}

		if stock[row][column-1] == '@' || stock[row][column-1] == 'x' {
			paperRolls++
		}

		if stock[row+1][column-1] == '@' || stock[row+1][column-1] == 'x' {
			paperRolls++
		}

		if stock[row+1][column] == '@' || stock[row+1][column] == 'x' {
			paperRolls++
		}
	default:
		if stock[row-1][column-1] == '@' || stock[row-1][column-1] == 'x' {
			paperRolls++
		}

		if stock[row-1][column] == '@' || stock[row-1][column] == 'x' {
			paperRolls++
		}

		if stock[row-1][column+1] == '@' || stock[row-1][column+1] == 'x' {
			paperRolls++
		}

		if stock[row][column+1] == '@' || stock[row][column+1] == 'x' {
			paperRolls++
		}

		if stock[row+1][column+1] == '@' || stock[row+1][column+1] == 'x' {
			paperRolls++
		}

		if stock[row+1][column] == '@' || stock[row+1][column] == 'x' {
			paperRolls++
		}

		if stock[row+1][column-1] == '@' || stock[row+1][column-1] == 'x' {
			paperRolls++
		}

		if stock[row][column-1] == '@' || stock[row][column-1] == 'x' {
			paperRolls++
		}
	}

	return paperRolls
}

func stepPart2(stock [][]byte) int {

	finalAmmount, rollsRemoved := 0, 1
	numberOfRows := len(stock)
	numberOfCols := len(stock[0])

	for {

		if rollsRemoved == 0 {
			break
		}

		fmt.Println("Starting iteration")
		rollsRemoved = 0

		for r := 0; r < numberOfRows; r++ {
			for c := 0; c < numberOfCols; c++ {

				if stock[r][c] != '@' {
					continue
				}

				//if its the first row
				switch r {
				case 0:
					if firstRowTotal(r, c, stock, numberOfCols) < 4 {
						stock[r][c] = 'x'
						rollsRemoved++
					}
				case numberOfRows - 1:
					if lastRowTotal(r, c, stock, numberOfCols) < 4 {
						stock[r][c] = 'x'
						rollsRemoved++
					}
				default:
					if anyOtherRowTotal(r, c, stock, numberOfCols) < 4 {
						stock[r][c] = 'x'
						rollsRemoved++

					}
				}
			}
		}

		fmt.Println("Rolls removed this iteration:", rollsRemoved)
		finalAmmount += rollsRemoved
		convertMarkedToEmpty(stock, numberOfRows, numberOfCols)
	}

	return finalAmmount

}

func convertMarkedToEmpty(stock [][]byte, numberOfRows int, numberOfCols int) {

	for r := 0; r < numberOfRows; r++ {
		for c := 0; c < numberOfCols; c++ {
			if stock[r][c] == 'x' {
				stock[r][c] = '.'
			}
		}
	}
}
