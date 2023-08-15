package crypt

import (
	"fmt"
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
	if !equal(CaesarDecrypt(CaesarEncrypt("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nec semp", 20), 20), "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nec semp") {
		t.Error("CaesarDecrypt(CaesarDecrypt(\"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nec semp\",20),20) != \"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nec semp\"")
	}
}

func equal(s, s1 string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s1[i] {
			fmt.Println(string(s[i]), string(s1[i]), i)
			return false
		}
	}
	return true
}
