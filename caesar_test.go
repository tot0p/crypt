package crypt

import (
	"strings"
	"testing"
)

func TestCaesarEncrypt(t *testing.T) {
	if CaesarEncrypt("abc", 1) != "bcd" {
		t.Error("CaesarEncrypt(\"abc\", 1) != \"bcd\"")
	}

	if CaesarEncrypt("abc", 2) != "cde" {
		t.Error("CaesarEncrypt(\"abc\", 2) != \"cde\"")
	}

	if CaesarEncrypt("abc", 3) != "def" {
		t.Error("CaesarEncrypt(\"abc\", 3) != \"def\"")
	}

	if CaesarEncrypt("abc", 4) != "efg" {
		t.Error("CaesarEncrypt(\"abc\", 4) != \"efg\"")
	}

}

func TestCaesarDecrypt(t *testing.T) {
	if CaesarDecrypt("bcd", 1) != "abc" {
		t.Error("CaesarDecrypt(\"bcd\", 1) != \"abc\"")
	}

	if CaesarDecrypt("cde", 2) != "abc" {
		t.Error("CaesarDecrypt(\"cde\", 2) != \"abc\"")
	}

	if CaesarDecrypt("def", 3) != "abc" {
		t.Error("CaesarDecrypt(\"def\", 3) != \"abc\"")
	}
}

func TestCaesarEncryptAndDecrypt(t *testing.T) {
	if CaesarDecrypt(CaesarEncrypt("abc", 1), 1) != "abc" {
		t.Error("CaesarDecrypt(CaesarEncrypt(\"abc\", 1), 1) != \"abc\"")
	}
	if 0 != strings.Compare(CaesarDecrypt(CaesarEncrypt("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nec semp", 20), 20), "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nec semp") {
		t.Error("CaesarDecrypt(CaesarDecrypt(\"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nec semp\",20),20) != \"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nec semp\"")
	}
}

func TestCaesarEncryptWithAlphabet(t *testing.T) {
	if CaesarEncryptWithAlphabet("abc1", 1, AlphabetMin) != "bcd1" {
		t.Error("CaesarEncryptWithAlphabet(\"abc1\", 1, alphabet) != \"bcd1\"")
	}

	if CaesarEncryptWithAlphabet("ABC1", 1, AlphabetMaj) != "BCD1" {
		t.Error("CaesarEncryptWithAlphabet(\"ABC1\", 1, alphabet) != \"BCD1\"")
	}
}

func TestCaesarDecryptWithAlphabet(t *testing.T) {

	if CaesarDecryptWithAlphabet("bcd1", 1, AlphabetMin) != "abc1" {
		t.Error("CaesarDecryptWithAlphabet(\"bcd1\", 1, alphabet) != \"abc1\"")
	}

	if CaesarDecryptWithAlphabet("BCD1", 1, AlphabetMaj) != "ABC1" {
		t.Error("CaesarDecryptWithAlphabet(\"BCD1\", 1, alphabet) != \"ABC1\"")
	}

}
