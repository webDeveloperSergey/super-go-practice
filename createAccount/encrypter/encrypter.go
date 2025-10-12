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
	key := os.Getenv("ENCRYPTION_KEY")

	if key == "" {
		panic("ENCRYPTION_KEY не задана в .env файле")
	}
	return &Encrypter{
		Key: key,
	}
}

func (enc *Encrypter) Encrypt(plainString []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nones := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nones)
	if err != nil {
		panic(err.Error())
	}

	return aesGCM.Seal(nones, nones, plainString, nil)
}

func (enc *Encrypter) Decrypt(encryptedString []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := encryptedString[:nonceSize], encryptedString[nonceSize:]

	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}

	return plainText
}