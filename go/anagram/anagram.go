package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	//panic("Please implement the Detect function")
	var anagrams []string

	for _, candidate := range candidates {
		if isAnagram(strings.ToLower(subject), strings.ToLower(candidate)) {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}

func sortedString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func isAnagram(subject string, candidate string) bool {
	return subject != candidate && sortedString(subject) == sortedString(candidate)
}
