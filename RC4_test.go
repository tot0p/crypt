package crypt

import "testing"

func TestRC4_Crypt(t *testing.T) {
	key := "key"
	data := "data"
	encrypt := RC4Crypt(key, data)
	decrypt := RC4Crypt(key, encrypt)
	if data != decrypt {
		t.Errorf("The data should be the same as the decrypt data.")
	}
}

func TestRC4_Encrypt_And_RC4_Decrypt(t *testing.T) {
	key := "key"
	data := "data"
	encrypt := RC4Encrypt(key, data)
	decrypt := RC4Decrypt(key, encrypt)
	if data != decrypt {
		t.Errorf("The data should be the same as the decrypt data.")
	}

	rc4 := NewRC4([]byte(key))
	encrypt2 := rc4.Encrypt([]byte(data))
	decrypt2 := rc4.Decrypt(encrypt2)
	if data != string(decrypt2) {
		t.Errorf("The data should be the same as the decrypt data.")
	}
}
