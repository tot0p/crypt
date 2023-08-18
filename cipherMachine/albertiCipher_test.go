package cipherMachine

import "testing"

func TestNewAlbertiCipher(t *testing.T) {
	ac := NewAlbertiCipher()
	ac.SetIncrement(1)
	err := ac.SetPeriod(-4)
	if err == nil {
		t.Errorf("The function should return an error.")
	}
	err = ac.SetPeriod(4)
	if err != nil {
		t.Errorf("The function return an error : %s", err.Error())
	}
	ac.SetDecalInit(0)
	err = ac.SetRotors("ABC", "abcD")
	if err == nil {
		t.Errorf("The function should return an error.")
	}
	err = ac.SetRotors(RotorExterior, RotorInterior)
	if err != nil {
		t.Errorf("The function return an error : %s", err.Error())
	}
}

func TestAlbertiCipher_Encrypt(t *testing.T) {
	ac := NewAlbertiCipher()
	message := "ABC DE"
	encryptedMessage, err := ac.Encrypt(message)
	if err != nil {
		t.Errorf("The encrypt func return an error : %s", err.Error())
	}
	if encryptedMessage != "USQ MK" {
		t.Errorf("The encrypted message should be \"USQ MK\" not %s.", encryptedMessage)
	}
	encryptedMessage, err = ac.Encrypt("ABCH")
	if err == nil {
		t.Errorf("The encrypt func should return an error.")
	}

	ac.SetIncrement(-11)

	message = "ABCDEFG"
	encryptedMessage, err = ac.Encrypt(message)
	if err != nil {
		t.Errorf("The encrypt func return an error : %s", err.Error())
	}
	if encryptedMessage != "USQOPRT" {
		t.Errorf("The encrypted message should be \"USQOPRT\" not %s.", encryptedMessage)
	}

	ac.SetIncrement(20)

	message = "ABCDEFG"
	encryptedMessage, err = ac.Encrypt(message)
	if err != nil {
		t.Errorf("The encrypt func return an error : %s", err.Error())
	}
	if encryptedMessage != "USQOUSQ" {
		t.Errorf("The encrypted message should be \"USQOUSQ\" not %s.", encryptedMessage)
	}
}

func TestAlbertiCipher_Decrypt(t *testing.T) {
	ac := NewAlbertiCipher()
	message := "USQ MK"
	decryptedMessage := ac.Decrypt(message)
	if decryptedMessage != "ABC DE" {
		t.Errorf("The decrypted message should be \"ABC DE\" not %s.", decryptedMessage)
	}

	ac.SetIncrement(-11)
	message = "USQOPRT"
	decryptedMessage = ac.Decrypt(message)
	if decryptedMessage != "ABCDEFG" {
		t.Errorf("The decrypted message should be \"ABCDEFG\" not %s.", decryptedMessage)
	}

	ac.SetIncrement(20)
	message = "USQOUSQ"
	decryptedMessage = ac.Decrypt(message)
	if decryptedMessage != "ABCDEFG" {
		t.Errorf("The decrypted message should be \"ABCDEFG\" not %s.", decryptedMessage)
	}
}
