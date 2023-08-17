package crypt

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

type KeyVigenere struct {
	LnBlock  int
	Decal    []int
	Alphabet string
}

func (k KeyVigenere) String() string {
	return fmt.Sprintf("KeyVigenere{LnBlock: %d, Decal: %v, Alphabet: %s}", k.LnBlock, k.Decal, k.Alphabet)
}

// NewKeyVigenere creates a new KeyVigenere (with Alphabet).
// The lnBlock is the length of the block.
// The decal is the decalage for each block.
func NewKeyVigenere(lnBlock int, decal []int) (*KeyVigenere, error) {
	return NewKeyVigenereWithAlphabet(lnBlock, decal, Alphabet)
}

// NewKeyVigenereRandom generate random Key
func NewKeyVigenereRandom() *KeyVigenere {
	return NewKeyVigenereRandomWithAlphabet(Alphabet)
}

// NewKeyVigenereWithAlphabet creates a new KeyVigenere.
// The lnBlock is the length of the block.
// The decal is the decalage for each block.
// The alphabet is the alphabet to use.
func NewKeyVigenereWithAlphabet(lnBlock int, decal []int, alphabet string) (*KeyVigenere, error) {
	ln := len(alphabet)
	for _, i := range decal {
		if i < 0 {
			return nil, errors.New("the decalage should be positive")
		} else if i >= ln {
			return nil, errors.New("the decalage should be less than the length of the alphabet")
		}
	}
	return &KeyVigenere{lnBlock, decal, alphabet}, nil
}

// NewKeyVigenereRandomWithAlphabet generate random Key
func NewKeyVigenereRandomWithAlphabet(alphabet string) *KeyVigenere {
	// Generate random length of block
	lnBlock := rand.Intn(len(alphabet)) + 1
	// Generate random decal
	decal := make([]int, lnBlock)
	for i := 0; i < lnBlock; i++ {
		decal[i] = rand.Intn(len(alphabet))
	}
	return &KeyVigenere{lnBlock, decal, alphabet}
}

// VigenereEncrypt encrypts a message with the vigenere cipher.
// The key is the key to use.
// The message is the message to encrypt.
func VigenereEncrypt(key *KeyVigenere, message string) string {
	encrypted := ""
	for i := 0; i < len(message); i += key.LnBlock {
		for j := 0; j < key.LnBlock; j++ {
			if i+j >= len(message) {
				break
			}
			if strings.Contains(key.Alphabet, string(message[i+j])) {
				encrypted += string(key.Alphabet[(strings.Index(key.Alphabet, string(message[i+j]))+key.Decal[j])%len(key.Alphabet)])
			} else {
				encrypted += string(message[i+j])
			}
		}
	}
	return encrypted
}

// VigenereDecrypt decrypts a message with the vigenere cipher.
// The key is the key to use.
// The message is the message to decrypt.
func VigenereDecrypt(key *KeyVigenere, message string) string {
	decrypted := ""
	for i := 0; i < len(message); i += key.LnBlock {
		for j := 0; j < key.LnBlock; j++ {
			if i+j >= len(message) {
				break
			}
			if strings.Contains(key.Alphabet, string(message[i+j])) {
				decrypted += string(key.Alphabet[(strings.Index(key.Alphabet, string(message[i+j]))-key.Decal[j])%len(key.Alphabet)])
			} else {
				decrypted += string(message[i+j])
			}
		}
	}
	return decrypted
}
