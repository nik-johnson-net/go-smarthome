package encoding

import (
	"reflect"
	"testing"
)

var (
	plaintext  = []byte{'t', 'e', 's', 't'}
	ciphertext = []byte{0xdf, 0xba, 0xc9, 0xbd}
)

func assertEqual(t *testing.T, expected interface{}, result interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, result) {
		t.Error("Expected", expected, "got", result)
	}
}

func TestEncrypt(t *testing.T) {
	result := Encrypt(plaintext)
	assertEqual(t, ciphertext, result)
}

func TestDecrypt(t *testing.T) {
	result := Decrypt(ciphertext)
	assertEqual(t, plaintext, result)
}
