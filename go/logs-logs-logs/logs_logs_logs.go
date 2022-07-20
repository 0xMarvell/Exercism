package logs

import (
	"fmt"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	var application string

	for _, val := range log {
		appUnicode := fmt.Sprintf("%U", val)
		if appUnicode == "U+2757" {
			application = "recommendation"
			break
		} else if appUnicode == "U+1F50D" {
			application = "search"
			break
		} else if appUnicode == "U+2600" {
			application = "weather"
			break
		} else {
			application = "default"
		}
	}
	return application
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newLog string
	oldRuneUnicode := fmt.Sprintf("%U", oldRune)

	for _, val := range log {
		u := fmt.Sprintf("%U", val)
		if u == oldRuneUnicode {
			newLog += string(newRune)
		} else {
			newLog += string(val)
		}
	}
	return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return limit >= utf8.RuneCountInString(log)
}
