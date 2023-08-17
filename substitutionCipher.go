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

// SubstitutionCipherEncryptWithAlphabetWithKey encrypts a message with the substitution cipher.
// The alphabet is the alphabet to use.
// The key is the key to use.
// The message is the message to encrypt.
func SubstitutionCipherEncryptWithAlphabetWithKey(alphabet, key, message string) (string, error) {
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

// SubstitutionCipherEncryptWAWK is an alias for SubstitutionCipherEncryptWithAlphabetWithKey.
func SubstitutionCipherEncryptWAWK(alphabet, key, message string) (string, error) {
	return SubstitutionCipherEncryptWithAlphabetWithKey(alphabet, key, message)
}

// SubstitutionCipherDecryptWithAlphabetWithKey decrypts a message with the substitution cipher.
// The alphabet is the alphabet to use.
// The key is the key to use.
// The message is the message to decrypt.
func SubstitutionCipherDecryptWithAlphabetWithKey(alphabet, key, message string) (string, error) {
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

// SubstitutionCipherDecryptWAWK is an alias for SubstitutionCipherDecryptWithAlphabetWithKey.
func SubstitutionCipherDecryptWAWK(alphabet, key, message string) (string, error) {
	return SubstitutionCipherDecryptWithAlphabetWithKey(alphabet, key, message)
}
