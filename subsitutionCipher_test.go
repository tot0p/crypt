package crypt

import (
	"strings"
	"testing"
)

// TestGenerateKeyForSubstitutionCipher tests the GenerateKeyForSubstitutionCipher function.
func TestGenerateKeyForSubstitutionCipher(t *testing.T) {
	key := GenerateKeyForSubstitutionCipher()
	if len(key) != len(Alphabet) {
		t.Errorf("The key should be the same length as the alphabet.")
	}

	for _, v := range Alphabet {
		if !strings.Contains(key, string(v)) {
			t.Errorf("The key should contain all the letters of the alphabet.")
		}
	}
}

// TestGenerateKeyForSC tests the GenerateKeyForSC function.
func TestGenerateKeyForSC(t *testing.T) {
	key := GenerateKeyForSC()
	if len(key) != len(Alphabet) {
		t.Errorf("The key should be the same length as the alphabet.")
	}

	for _, v := range Alphabet {
		if !strings.Contains(key, string(v)) {
			t.Errorf("The key should contain all the letters of the alphabet.")
		}
	}
}

// TestGenerateKeyForSubstitutionCipherWithAlphabet tests the GenerateKeyForSubstitutionCipherWithAlphabet function.
func TestGenerateKeyForSubstitutionCipherWithAlphabet(t *testing.T) {
	key := GenerateKeyForSubstitutionCipherWithAlphabet(AlphabetMin)
	if len(key) != len(AlphabetMin) {
		t.Errorf("The key should be the same length as the alphabet.")
	}

	for _, v := range AlphabetMin {
		if !strings.Contains(key, string(v)) {
			t.Errorf("The key should contain all the letters of the alphabet.")
		}
	}
}

// TestGenerateKeyForSCWA tests the GenerateKeyForSCWA function.
func TestGenerateKeyForSCWA(t *testing.T) {
	key := GenerateKeyForSCWA(AlphabetMin)
	if len(key) != len(AlphabetMin) {
		t.Errorf("The key should be the same length as the alphabet.")
	}

	for _, v := range AlphabetMin {
		if !strings.Contains(key, string(v)) {
			t.Errorf("The key should contain all the letters of the alphabet.")
		}
	}
}

// TestSubstitutionCipherEncrypt tests the SubstitutionCipherEncrypt function.
func TestSubstitutionCipherEncrypt(t *testing.T) {
	key := func() string {
		key := ""
		for i := len(Alphabet) - 1; i >= 0; i-- {
			key += string(Alphabet[i])
		}
		return key
	}()
	message := "Hello World!"
	encryptedMessage, err := SubstitutionCipherEncrypt(key, message)
	if err != nil {
		t.Errorf("The encrypt func return an error : %s", err.Error())
	}
	if encryptedMessage != "C5YYV nVSY6!" {
		t.Errorf("The encrypted message should be \"C5YYV nVSY6!\" not %s.", encryptedMessage)
	}
}

// TestSubstitutionCipherEncryptWithAlphabet tests the SubstitutionCipherEncryptWithAlphabet function.
func TestSubstitutionCipherEncryptWithAlphabet(t *testing.T) {
	key := func() string {
		key := ""
		for i := len(Alphabet) - 1; i >= 0; i-- {
			key += string(Alphabet[i])
		}
		return key
	}()
	message := "Hello World!"
	encryptedMessage, err := SubstitutionCipherEncryptWithAlphabet(Alphabet, key, message)
	if err != nil {
		t.Errorf("The encrypt func return an error : %s", err.Error())
	}
	if encryptedMessage != "C5YYV nVSY6!" {
		t.Errorf("The encrypted message should be \"C5YYV nVSY6!\" not %s.", encryptedMessage)
	}
}

// TestSubstitutionCipherEncryptWA tests the SubstitutionCipherEncryptWA function.
func TestSubstitutionCipherEncryptWA(t *testing.T) {
	key := func() string {
		key := ""
		for i := len(Alphabet) - 1; i >= 0; i-- {
			key += string(Alphabet[i])
		}
		return key
	}()
	message := "Hello World!"
	encryptedMessage, err := SubstitutionCipherEncryptWA(Alphabet, key, message)
	if err != nil {
		t.Errorf("The encrypt func return an error : %s", err.Error())
	}
	if encryptedMessage != "C5YYV nVSY6!" {
		t.Errorf("The encrypted message should be \"C5YYV nVSY6!\" not %s.", encryptedMessage)
	}
}

// TestSubstitutionCipherDecrypt tests the SubstitutionCipherDecrypt function.
func TestSubstitutionCipherDecrypt(t *testing.T) {
	key := func() string {
		key := ""
		for i := len(Alphabet) - 1; i >= 0; i-- {
			key += string(Alphabet[i])
		}
		return key
	}()
	message := "C5YYV nVSY6!"
	decryptedMessage, err := SubstitutionCipherDecrypt(key, message)
	if err != nil {
		t.Errorf("The decrypt func return an error : %s", err.Error())
	}
	if decryptedMessage != "Hello World!" {
		t.Errorf("The decrypted message should be \"Hello World!\" not %s.", decryptedMessage)
	}
}

// TestSubstitutionCipherDecryptWithAlphabet tests the SubstitutionCipherDecryptWithAlphabet function.
func TestSubstitutionCipherDecryptWithAlphabet(t *testing.T) {
	key := func() string {
		key := ""
		for i := len(Alphabet) - 1; i >= 0; i-- {
			key += string(Alphabet[i])
		}
		return key
	}()
	message := "C5YYV nVSY6!"
	decryptedMessage, err := SubstitutionCipherDecryptWithAlphabet(Alphabet, key, message)
	if err != nil {
		t.Errorf("The decrypt func return an error : %s", err.Error())
	}
	if decryptedMessage != "Hello World!" {
		t.Errorf("The decrypted message should be \"Hello World!\" not %s.", decryptedMessage)
	}
}

// TestSubstitutionCipherDecryptWA tests the SubstitutionCipherDecryptWA function.
func TestSubstitutionCipherDecryptWA(t *testing.T) {
	key := func() string {
		key := ""
		for i := len(Alphabet) - 1; i >= 0; i-- {
			key += string(Alphabet[i])
		}
		return key
	}()
	message := "C5YYV nVSY6!"
	decryptedMessage, err := SubstitutionCipherDecryptWA(Alphabet, key, message)
	if err != nil {
		t.Errorf("The decrypt func return an error : %s", err.Error())
	}
	if decryptedMessage != "Hello World!" {
		t.Errorf("The decrypted message should be \"Hello World!\" not %s.", decryptedMessage)
	}
}
