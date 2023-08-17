package crypt

import (
	"errors"
	"math/rand"
	"strings"
)

// GenerateKeyForSubstitutionCipher generates a key for the substitution cipher. (with Alphabet)
// The key is a random permutation of the alphabet.
func GenerateKeyForSubstitutionCipher() string {
	return GenerateKeyForSubstitutionCipherWithAlphabet(Alphabet)
}

// GenerateKeyForSC is an alias for GenerateKeyForSubstitutionCipher.
func GenerateKeyForSC() string {
	return GenerateKeyForSubstitutionCipher()
}

// GenerateKeyForSubstitutionCipherWithAlphabet generates a key for the substitution cipher.
// The key is a random permutation of the alphabet.
// The alphabet is the alphabet to use.
func GenerateKeyForSubstitutionCipherWithAlphabet(alphabet string) string {
	tempAlphabet := []rune(alphabet)
	rand.Shuffle(len(tempAlphabet), func(i, j int) { tempAlphabet[i], tempAlphabet[j] = tempAlphabet[j], tempAlphabet[i] })
	return string(tempAlphabet)
}

// GenerateKeyForSCWA is an alias for GenerateKeyForSubstitutionCipherWithAlphabet.
func GenerateKeyForSCWA(alphabet string) string {
	return GenerateKeyForSubstitutionCipherWithAlphabet(alphabet)
}

// SubstitutionCipherEncrypt encrypts a message with the substitution cipher. (with Alphabet)
// The key is the key to use.
// The message is the message to encrypt
func SubstitutionCipherEncrypt(key, message string) (string, error) {
	return SubstitutionCipherEncryptWA(Alphabet, key, message)
}

// SubstitutionCipherEncryptWithAlphabet encrypts a message with the substitution cipher.
// The alphabet is the alphabet to use.
// The key is the key to use.
// The message is the message to encrypt.
func SubstitutionCipherEncryptWithAlphabet(alphabet, key, message string) (string, error) {
	if len(alphabet) != len(key) {
		return "", errors.New("the alphabet and the key should be the same length")
	}
	result := ""
	for _, v := range message {
		i := strings.Index(alphabet, string(v))
		if i == -1 {
			result += string(v)
		} else {
			result += string(key[i])
		}
	}
	return result, nil
}

// SubstitutionCipherEncryptWA is an alias for SubstitutionCipherEncryptWithAlphabetWithKey.
func SubstitutionCipherEncryptWA(alphabet, key, message string) (string, error) {
	return SubstitutionCipherEncryptWithAlphabet(alphabet, key, message)
}

// SubstitutionCipherDecrypt decrypts a message with the substitution cipher. (with Alphabet)
// The key is the key to use.
// The message is the message to decrypt
func SubstitutionCipherDecrypt(key, message string) (string, error) {
	return SubstitutionCipherDecryptWA(Alphabet, key, message)
}

// SubstitutionCipherDecryptWithAlphabet decrypts a message with the substitution cipher.
// The alphabet is the alphabet to use.
// The key is the key to use.
// The message is the message to decrypt.
func SubstitutionCipherDecryptWithAlphabet(alphabet, key, message string) (string, error) {
	if len(alphabet) != len(key) {
		return "", errors.New("the alphabet and the key should be the same length")
	}
	result := ""
	for _, v := range message {
		i := strings.Index(key, string(v))
		if i == -1 {
			result += string(v)
		} else {
			result += string(alphabet[i])
		}
	}
	return result, nil
}

// SubstitutionCipherDecryptWA is an alias for SubstitutionCipherDecryptWithAlphabetWithKey.
func SubstitutionCipherDecryptWA(alphabet, key, message string) (string, error) {
	return SubstitutionCipherDecryptWithAlphabet(alphabet, key, message)
}
