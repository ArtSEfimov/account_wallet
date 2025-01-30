package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("ENCRYPT_KEY")
	if key == "" {
		panic("encrypt key not found")
	}

	return &Encrypter{
		Key: key,
	}
}

func (e *Encrypter) Encrypt(plainString []byte) []byte {
	block, err := aes.NewCipher([]byte(e.Key))
	if err != nil {
		panic(err)
	}
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	nonce := make([]byte, aesGSM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err)
	}

	return aesGSM.Seal(nonce, nonce, plainString, nil)
}

func (e *Encrypter) Decrypt(encryptedString []byte) []byte {
	block, err := aes.NewCipher([]byte(e.Key))
	if err != nil {
		panic(err)
	}
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	nonceSize := aesGSM.NonceSize()
	nonce, cipherText := encryptedString[:nonceSize], encryptedString[nonceSize:]

	plainText, err := aesGSM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err)
	}
	return plainText
}
