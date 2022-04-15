package reverse

// Reverse takes a string as argument and
// returns the reversed version of the string.
func Reverse(input string) string {
	var res string
	for _, v := range input {
		res = string(v) + res
	}
	return res
}
