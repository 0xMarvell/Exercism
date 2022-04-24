// Package anagram checks if a word is an anagram or not
package anagram

import (
	"sort"
	"strings"
)

// Detect filters candidates selecting only anagrams of the word.
func Detect(subject string, candidates []string) []string {
	var anagrams []string

	for _, candidate := range candidates {
		if isAnagram(strings.ToLower(subject), strings.ToLower(candidate)) {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}

// sortedString converts a string to a slice of strings and sorts the elements
// of the slice in ascending order then converts the sorted slice back into a string
func sortedString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// isAnagram checks if the candidate is an anagram of the subject
func isAnagram(subject string, candidate string) bool {
	return subject != candidate && sortedString(subject) == sortedString(candidate)
}
