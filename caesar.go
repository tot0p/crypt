package crypt

import "strings"

// CaesarEncrypt encrypts the string with the given key.
// The key is the number of characters to shift.
// For example, if the key is 1, then 'a' will be 'b'.
func CaesarEncrypt(s string, key int) string {
	var result string
	for _, v := range s {
		if strings.Contains(alphatbet, string(v)) {
			result += string(alphatbet[(strings.Index(alphatbet, string(v))+key)%len(alphatbet)])
		} else {
			result += string(v)
		}
	}
	return result
}

// CaesarDecrypt decrypts the string with the given key.
// The key is the number of characters to shift.
// For example, if the key is 1, then 'b' will be 'a'.
func CaesarDecrypt(s string, key int) string {
	var result string
	for _, v := range s {
		if strings.Contains(alphatbet, string(v)) {
			result += string(alphatbet[(strings.Index(alphatbet, string(v))-key)%len(alphatbet)])
		} else {
			result += string(v)
		}
	}
	return result
}
