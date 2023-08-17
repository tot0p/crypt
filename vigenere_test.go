package crypt

import (
	"strings"
	"testing"
)

// TestNewKeyVigenere tests the NewKeyVigenere function.
func TestNewKeyVigenere(t *testing.T) {
	Key, err := NewKeyVigenere(3, []int{1, 2, 3})
	if err != nil {
		t.Errorf("The function return an error : %s", err.Error())
	}
	if Key.LnBlock != 3 {
		t.Errorf("The length of the block should be 3 not %d.", Key.LnBlock)
	}
	if Key.Decal[0] != 1 {
		t.Errorf("The decalage should be 1 not %d.", Key.Decal[0])
	}
	if Key.Decal[1] != 2 {
		t.Errorf("The decalage should be 2 not %d.", Key.Decal[1])
	}
	if Key.Decal[2] != 3 {
		t.Errorf("The decalage should be 3 not %d.", Key.Decal[2])
	}
	if strings.Compare(Key.Alphabet, Alphabet) != 0 {
		t.Errorf("The alphabet should be %s not %s.", Alphabet, Key.Alphabet)
	}
}

// TestNewKeyVigenereRandom tests the NewKeyVigenereRandom function.
func TestNewKeyVigenereRandom(t *testing.T) {
	Key := NewKeyVigenereRandom()
	if Key.LnBlock <= 0 {
		t.Errorf("The length of the block should be positive not %d.", Key.LnBlock)
	}
	for _, v := range Key.Decal {
		if v < 0 {
			t.Errorf("The decalage should be positive not %d.", v)
		}
		if v >= len(Key.Alphabet) {
			t.Errorf("The decalage should be less than the length of the alphabet not %d.", v)
		}
	}
	if strings.Compare(Key.Alphabet, Alphabet) != 0 {
		t.Errorf("The alphabet should be %s not %s.", Alphabet, Key.Alphabet)
	}
}

// TestNewKeyVigenereWithAlphabet tests the NewKeyVigenereWithAlphabet function.
func TestNewKeyVigenereWithAlphabet(t *testing.T) {
	Key, err := NewKeyVigenereWithAlphabet(3, []int{1, 2, 3}, AlphabetMaj)
	if err != nil {
		t.Errorf("The function return an error : %s", err.Error())
	}
	if Key.LnBlock != 3 {
		t.Errorf("The length of the block should be 3 not %d.", Key.LnBlock)
	}
	if Key.Decal[0] != 1 {
		t.Errorf("The decalage should be 1 not %d.", Key.Decal[0])
	}
	if Key.Decal[1] != 2 {
		t.Errorf("The decalage should be 2 not %d.", Key.Decal[1])
	}
	if Key.Decal[2] != 3 {
		t.Errorf("The decalage should be 3 not %d.", Key.Decal[2])
	}
	if strings.Compare(Key.Alphabet, AlphabetMaj) != 0 {
		t.Errorf("The alphabet should be %s not %s.", AlphabetMaj, Key.Alphabet)
	}

	_, err = NewKeyVigenereWithAlphabet(3, []int{1, -1, 3}, AlphabetMin)
	if err == nil {
		t.Errorf("The function should return an error.")
	}
	_, err = NewKeyVigenereWithAlphabet(3, []int{1, 30, 3}, AlphabetMin)
	if err == nil {
		t.Errorf("The function should return an error.")
	}
}

// TestNewKeyVigenereRandomWithAlphabet tests the NewKeyVigenereRandomWithAlphabet function.
func TestNewKeyVigenereRandomWithAlphabet(t *testing.T) {
	Key := NewKeyVigenereRandomWithAlphabet(AlphabetMaj)
	if Key.LnBlock <= 0 {
		t.Errorf("The length of the block should be positive not %d.", Key.LnBlock)
	}
	for _, v := range Key.Decal {
		if v < 0 {
			t.Errorf("The decalage should be positive not %d.", v)
		}
		if v >= len(Key.Alphabet) {
			t.Errorf("The decalage should be less than the length of the alphabet not %d.", v)
		}
	}
	if strings.Compare(Key.Alphabet, AlphabetMaj) != 0 {
		t.Errorf("The alphabet should be %s not %s.", AlphabetMaj, Key.Alphabet)
	}
}

// TestKeyVigenereString tests the KeyVigenere.String function.
func TestKeyVigenereString(t *testing.T) {
	Key, _ := NewKeyVigenere(3, []int{1, 2, 3})
	if strings.Compare(Key.String(), "KeyVigenere{LnBlock: 3, Decal: [1 2 3], Alphabet: "+Alphabet+"}") != 0 {
		t.Errorf("The string should be \"KeyVigenere{LnBlock: 3, Decal: [1 2 3], Alphabet: %s}\" not \"%s\".", Alphabet, Key.String())
	}
}

// TestVigenereEncrypt tests the VigenereEncrypt function.
func TestVigenereEncrypt(t *testing.T) {
	Key, _ := NewKeyVigenere(3, []int{1, 2, 3})
	message := "Hello World!  "
	encryptedMessage := VigenereEncrypt(Key, message)

	if strings.Compare(encryptedMessage, "Igomq Xqumf!  ") != 0 {
		t.Errorf("The encrypted message should be \"Igomq Xqumf!  \" not \"%s\".", encryptedMessage)
	}
}

// TestVigenereDecrypt tests the VigenereDecrypt function.
func TestVigenereDecrypt(t *testing.T) {
	Key, _ := NewKeyVigenere(3, []int{1, 2, 3})
	message := "Igomq Xqumf!  "
	decryptedMessage := VigenereDecrypt(Key, message)

	if strings.Compare(decryptedMessage, "Hello World!  ") != 0 {
		t.Errorf("The decrypted message should be \"Hello World!  \" not \"%s\".", decryptedMessage)
	}
}
