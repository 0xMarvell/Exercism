package reverse

func Reverse(input string) string {
	//panic("Please implement the Reverse function")
	var res string
	for _, v := range input {
		res = string(v) + res
	}
	return res
}
