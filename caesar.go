package crypt

import "strings"

// CaesarEncrypt encrypts the string with the given key. (with Alphabet)
// The key is the number of characters to shift.
// For example, if the key is 1, then 'a' will be 'b'.
func CaesarEncrypt(s string, key int) string {
	return CaesarEncryptWithAlphabet(s, key, Alphabet)
}

// CaesarDecrypt decrypts the string with the given key. (with Alphabet)
// The key is the number of characters to shift.
// For example, if the key is 1, then 'b' will be 'a'.
func CaesarDecrypt(s string, key int) string {
	return CaesarDecryptWithAlphabet(s, key, Alphabet)
}

// CaesarEncryptWithAlphabet encrypts the string with the given key.
// The key is the number of characters to shift.
// For example, if the key is 1, then 'a' will be 'b'.
// The alphabet is the alphabet to use.
func CaesarEncryptWithAlphabet(s string, key int, alphabet string) string {
	var result string
	for _, v := range s {
		if strings.Contains(alphabet, string(v)) {
			result += string(alphabet[(strings.Index(alphabet, string(v))+key)%len(alphabet)])
		} else {
			result += string(v)
		}
	}
	return result
}

// CaesarDecryptWithAlphabet decrypts the string with the given key.
// The key is the number of characters to shift.
// For example, if the key is 1, then 'b' will be 'a'.
// The alphabet is the alphabet to use.
func CaesarDecryptWithAlphabet(s string, key int, alphabet string) string {
	var result string
	for _, v := range s {
		if strings.Contains(alphabet, string(v)) {
			result += string(alphabet[(strings.Index(alphabet, string(v))-key)%len(alphabet)])
		} else {
			result += string(v)
		}
	}
	return result
}
