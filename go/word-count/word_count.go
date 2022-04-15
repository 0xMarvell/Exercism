package wordcount

import "strings"

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	// Convert all whitespace chars(\t, \n) to just spaces
	newPhrase := strings.Replace(phrase, "\n", " ", -1)
	newPhrase = strings.Replace(newPhrase, "\t", " ", -1)
	newPhrase = strings.Replace(newPhrase, ",", " ", -1)

	// Now remove all unnecessary symbols
	remove := [8]string{"!", "@", "$", "%", "^", "&", ":", "."}
	for _, i := range remove {
		newPhrase = strings.Replace(newPhrase, i, "", -1)
	}

	// Transfer each word to lowercase
	newPhrase = strings.ToLower(newPhrase)
	// Split the input string
	result := strings.Split(newPhrase, " ")
	// Create new Frequency
	freqmap := make(Frequency)

	// Iterate over remaining list
	for _, val := range result {
		if val == "" {
			continue
		}
		// If first(or last) char is ' having an apostrophe inside the word wouldnt make sense
		if string(val[0]) == "'" {
			val = strings.Replace(val, "'", "", -1)
		}
		// Now iterate counter in map based on encountered word
		freqmap[val] += 1
	}
	return freqmap

}
