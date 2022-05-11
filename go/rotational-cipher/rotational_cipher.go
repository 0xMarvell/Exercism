package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	//panic("Please implement the RotationalCipher function")
	var newStr string

	if shiftKey == 0 || shiftKey == 26 {
		return plain
	}

	return newStr
}
