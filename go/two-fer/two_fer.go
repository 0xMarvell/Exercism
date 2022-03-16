// Package twofer — Two-fer or 2-fer is short for two for one. One for you and one for me.
package twofer

import (
	"fmt"
)

// ShareWith takes a name as argument and returns a string with this name. However, if the name
// is missing, the function returns the string with the string 'you' in place of the name variable.
func ShareWith(name string) string {
	if name != "" {
		return fmt.Sprintf("One for %s, one for me.", name)
	}
	return fmt.Sprintf("One for you, one for me.")
}
