package encoding

const (
	iv byte = 0xAB
)

// Encrypt encrypts plaintext with TP-Link's form of an Autokey Cipher.
// The function is simply {ciphertext = key XOR plaintext} with the initial key being  0xAB.
func Encrypt(plaintext []byte) []byte {
	ciphertext := make([]byte, len(plaintext))

	key := iv
	for i, currentByte := range plaintext {
		encByte := key ^ currentByte
		key = encByte
		ciphertext[i] = encByte
	}

	return ciphertext
}

// Decrypt decrypts ciphertext encrypted with TP-Link's form of an Autokey Cipher.
// See Encrypt for details.
func Decrypt(ciphertext []byte) []byte {
	plaintext := make([]byte, len(ciphertext))

	key := iv
	for i, currentByte := range ciphertext {
		decByte := key ^ currentByte
		key = currentByte
		plaintext[i] = decByte
	}

	return plaintext
}
