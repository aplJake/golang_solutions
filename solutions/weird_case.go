package solutions

import (
	"fmt"
	"strings"
)

/*
Write a function toWeirdCase (weirdcase in Ruby) that accepts a string,
and returns the same string with all even indexed characters in each
word upper cased, and all odd indexed characters in each word lower cased.
The indexing just explained is zero based, so the zero-ith index is even,
therefore that character should be upper cased.
 */

func ToWeirdCase(str string) string {
	var resultString string
	arrString :=  []rune(str)

	for i := 0; i < len(str); i++ {
		if i % 2 == 0 {
			resultString += strings.ToUpper(string(arrString[i]))
		} else {
			resultString += strings.ToLower(string(arrString[i]))

		}
	}

	fmt.Println(resultString)
	return string(resultString)
}
