package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Take user input for the set of digits
	var input string
	fmt.Println("Enter the digits separated by commas (e.g., 2,0,0,8):")
	fmt.Scanln(&input)

	// Parse the input string into an array of integers
	digitStrings := strings.Split(input, ",")
	digits := make([]int, len(digitStrings))
	for i, ds := range digitStrings {
		digit, err := strconv.Atoi(ds)
		if err != nil {
			fmt.Println("Error: Invalid input")
			return
		}
		digits[i] = digit
	}

	// Calculate the number of unique digits
	uniqueDigits := make(map[int]bool)
	for _, digit := range digits {
		uniqueDigits[digit] = true
	}
	n := len(uniqueDigits)

	// Number of positions (digits in the four-digit number)
	r := 4

	// Calculate permutations with repetition
	permutations := int(math.Pow(float64(n), float64(r)))

	// Create and open a new file for writing
	file, err := os.Create("possible_digits.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for i := 0; i < permutations; i++ {
		number := make([]int, r)
		for j := 0; j < r; j++ {
			digitIndex := (i / int(math.Pow(float64(n), float64(j)))) % n
			number[j] = digits[digitIndex]
		}
		fmt.Fprintf(file, "%d%d%d%d\n", number[0], number[1], number[2], number[3])
	}

	fmt.Printf("Number of possible four-digit numbers: %d\n", permutations)
	fmt.Println("Possible digits have been saved to possible_digits.txt")
}
