// Package rotationalcipher implements a simple shift cipher
package rotationalcipher

// RotationalCipher encrypts a string using the caesar cipher style of encryption
func RotationalCipher(plain string, shiftKey int) string {
	newStr := make([]rune, 0)

	if shiftKey == 0 || shiftKey == 26 {
		return plain
	}

	rshiftKey := rune(shiftKey)

	for _, r := range plain {
		switch {
		case r >= 'A' && r <= 'Z':
			newStr = append(newStr, 'A'+(r-'A'+rshiftKey)%26)
		case r >= 'a' && r <= 'z':
			newStr = append(newStr, 'a'+(r-'a'+rshiftKey)%26)
		default:
			newStr = append(newStr, r)
		}
	}

	return string(newStr)
}
