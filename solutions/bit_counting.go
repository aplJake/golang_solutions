package solutions

import (
	"fmt"
	"strconv"
)

/*
Write a function that takes an integer as input,
and returns the number of bits that are equal to one
in the binary representation of that number. You can
guarantee that input is non-negative.
 */

func CountBits(number uint) int {
	// transform integer number to binary form
	sumOfBits := 0
	binary := 0
	counter := 1
	remainder := 0

	for number != 0 {
		remainder = int(number % 2)
		number = number / 2
		binary += remainder * counter
		counter *= 10

	}
	fmt.Println(binary)

	// get the sum of bits
	stringBinary := strconv.Itoa(binary)
	for _, b := range stringBinary {
		if b == '1' {
			sumOfBits += 1
		}
	}
	fmt.Println(sumOfBits)

	return sumOfBits
}
