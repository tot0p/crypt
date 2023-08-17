package crypt

// RC4 is a RC4 cipher.
type RC4 [256]byte

// RC4Crypt encrypts or decrypts data with the RC4 algorithm.
func RC4Crypt[T []byte | string](key T, data T) T {
	rc4 := NewRC4([]byte(key))
	return T(rc4.Crypt([]byte(data)))
}

// RC4Encrypt encrypts data with the RC4 algorithm.
func RC4Encrypt[T []byte | string](key T, data T) T {
	return RC4Crypt(key, data)
}

// RC4Decrypt decrypts data with the RC4 algorithm.
func RC4Decrypt[T []byte | string](key T, data T) T {
	return RC4Crypt(key, data)
}

// NewRC4 creates a new RC4.
// The key is the key to use.
func NewRC4(key []byte) *RC4 {
	rc4 := new(RC4)
	for i := 0; i < 256; i++ {
		rc4[i] = byte(i)
	}
	j := 0
	for i := 0; i < 256; i++ {
		j = (j + int(rc4[i]) + int(key[i%len(key)])) % 256
		rc4[i], rc4[j] = rc4[j], rc4[i]
	}
	return rc4
}

// Crypt encrypts or decrypts data with the RC4 algorithm.
func (rc4 RC4) Crypt(data []byte) []byte {
	result := make([]byte, len(data))
	i, j := 0, 0
	for k := 0; k < len(data); k++ {
		i = (i + 1) % 256
		j = (j + int(rc4[i])) % 256
		rc4[i], rc4[j] = rc4[j], rc4[i]
		result[k] = data[k] ^ rc4[(int(rc4[i])+int(rc4[j]))%256]
	}
	return result
}

// Encrypt encrypts data with the RC4 algorithm.
func (rc4 RC4) Encrypt(data []byte) []byte {
	return rc4.Crypt(data)
}

// Decrypt decrypts data with the RC4 algorithm.
func (rc4 RC4) Decrypt(data []byte) []byte {
	return rc4.Crypt(data)
}
