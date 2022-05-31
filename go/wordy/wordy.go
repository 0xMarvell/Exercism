// Package wordy helps convert word problems into solvable math problems.
package wordy

import (
	"strconv"
	"strings"
)

// Answer parses question into a format that can be used calculated.
func Answer(question string) (int, bool) {
	var word string
	var num int
	var err error
	// get rid of the question mark
	question = strings.TrimSuffix(question, "?")

	i, j := 2, 0
	// convert question into a slice of strings
	quesArray := strings.Split(question, " ")
	if i >= len(quesArray) {
		return 0, false
	}

	word = quesArray[i]
	// convert word from string to int
	if num, err = strconv.Atoi(word); err != nil {
		return 0, false
	}
	j = num

	for {
		i++
		if i == len(quesArray) { // if quesArray has the same numer of words as i's incremented value
			return j, true
		}
		word = quesArray[i]

		switch word {
		case "plus":
			i++
			if i == len(quesArray) {
				return 0, false
			}
			word = quesArray[i]
			if num, err = strconv.Atoi(word); err != nil {
				return 0, false
			}
			j += num
		case "minus":
			i++
			if i == len(quesArray) {
				return 0, false
			}
			word = quesArray[i]
			if num, err = strconv.Atoi(word); err != nil {
				return 0, false
			}
			j -= num
		case "multiplied":
			i += 2
			if i >= len(quesArray) {
				return 0, false
			}
			word = quesArray[i]
			if num, err = strconv.Atoi(word); err != nil {
				return 0, false
			}
			j *= num
		case "divided":
			i += 2
			if i >= len(quesArray) {
				return 0, false
			}
			word = quesArray[i]
			if num, err = strconv.Atoi(word); err != nil {
				return 0, false
			}
			j /= num
		default:
			return 0, false
		}
	}

}
