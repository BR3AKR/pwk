package cryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/scrypt"
)

// CreateHash takes a string and hashes it
func CreateHash(password string) ([]byte, error) {
	// TODO Generate random salt with init command
	salt := []byte{0xe3, 0x12, 0xc9, 0x8f, 0xc7, 0xcc, 0xab, 0xaf, 0xa3, 0x12, 0xc9, 0x9f, 0xa7, 0xfc, 0xa3, 0xa8}
	dk, err := scrypt.Key([]byte(password), salt, 32768, 16, 1, 32)
	return dk, err
}

func Encrypt(data []byte, hash []byte) []byte {
	block, err := aes.NewCipher(hash)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	return gcm.Seal(nonce, nonce, data, nil)
}

func Decrypt(data []byte, hash []byte) ([]byte, error) {
	block, err := aes.NewCipher(hash)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func EncryptFile(filename string, data []byte, hash []byte) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(Encrypt(data, hash))
}

func DecryptFile(filename string, hash []byte) ([]byte, error) {
	data, _ := ioutil.ReadFile(filename)
	return Decrypt(data, hash)
}
