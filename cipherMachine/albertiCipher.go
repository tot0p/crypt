package cipherMachine

import (
	"errors"
	"strings"
)

const RotorExterior = "ABCDEFGILMNOPQRSTVXZ1234"
const RotorInterior = "USQOMKHFDBACEGILNPRTXZ&Y"
const DecalInit = 0
const Increment = 1
const Period = 4

type AlbertiCipher struct {
	rotorExterior string
	rotorInterior string
	decalInit     int
	increment     int
	period        int
}

// NewAlbertiCipher creates a new AlbertiCipher.
func NewAlbertiCipher() *AlbertiCipher {
	albertiCipher := new(AlbertiCipher)
	albertiCipher.rotorExterior = RotorExterior
	albertiCipher.rotorInterior = RotorInterior
	albertiCipher.decalInit = DecalInit
	albertiCipher.increment = Increment
	albertiCipher.period = Period
	return albertiCipher
}

// SetRotors sets the rotors.
func (albertiCipher *AlbertiCipher) SetRotors(rotorExterior string, rotorInterior string) error {
	if len(rotorExterior) != len(rotorInterior) {
		return errors.New("the rotors should be the same length")
	}
	albertiCipher.rotorExterior = rotorExterior
	albertiCipher.rotorInterior = rotorInterior
	return nil
}

// SetDecalInit sets the decalInit.
func (albertiCipher *AlbertiCipher) SetDecalInit(decalInit int) {
	albertiCipher.decalInit = decalInit
}

// SetIncrement sets the increment.
func (albertiCipher *AlbertiCipher) SetIncrement(increment int) {
	albertiCipher.increment = increment
}

// SetPeriod sets the period.
func (albertiCipher *AlbertiCipher) SetPeriod(period int) error {
	if period <= 0 {
		return errors.New("the period should be positive")
	}
	albertiCipher.period = period
	return nil
}

// Encrypt encrypts data with the Alberti cipher.
// The data is the data to encrypt.
func (albertiCipher AlbertiCipher) Encrypt(data string) (string, error) {
	result := ""
	for i, v := range data {
		j := strings.Index(albertiCipher.rotorExterior, string(v))
		if j == -1 {
			if strings.Contains(albertiCipher.rotorInterior, string(v)) {
				return "", errors.New("the data is not valid because it contains a letter in cipher rotor")
			} else {
				result += string(v)
			}
		} else {
			index := (j + albertiCipher.decalInit + i/albertiCipher.period*albertiCipher.increment) % len(albertiCipher.rotorInterior)
			if index < 0 {
				result += string(albertiCipher.rotorInterior[len(albertiCipher.rotorInterior)+index])
			} else {
				result += string(albertiCipher.rotorInterior[index])
			}
		}
	}
	return result, nil
}

// Decrypt decrypts data with the Alberti cipher.
// The data is the data to decrypt.
func (albertiCipher AlbertiCipher) Decrypt(data string) string {
	result := ""
	for i, v := range data {
		j := strings.Index(albertiCipher.rotorInterior, string(v))
		if j == -1 {
			result += string(v)
		} else {
			index := (j - albertiCipher.decalInit - i/albertiCipher.period*albertiCipher.increment) % len(albertiCipher.rotorExterior)
			if index < 0 {
				result += string(albertiCipher.rotorExterior[len(albertiCipher.rotorExterior)+index])
			} else {
				result += string(albertiCipher.rotorExterior[index])
			}
		}
	}
	return result
}
