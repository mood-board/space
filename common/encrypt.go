package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
)

type EncryptedData struct {
	CipherText []byte `bson:"cipher_text"`
	MAC        []byte `bson:"mac"`
}

func Encrypt(obj interface{}, encryptKey []byte, authKey []byte) (EncryptedData, error) {
	/* The struct that we are encrypting needs to support JSON serialization */
	body, err := json.Marshal(obj)
	if err != nil {
		return EncryptedData{nil, nil}, err
	}

	block, err := aes.NewCipher(encryptKey)
	if err != nil {
		return EncryptedData{nil, nil}, err
	}

	/* We'll store the initialization vector at the start of the
	 * func (this InstrumentDataCard) cipher text */
	cipherText := make([]byte, aes.BlockSize+len(body))
	iv := cipherText[:aes.BlockSize]

	/* Fill initialization vector */
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return EncryptedData{nil, nil}, err
	}

	/* Do encryption using CTR mode of operation */
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], body)

	/* Generate MAC to be used to authenticate the encrypted data */
	hasher := hmac.New(sha256.New, authKey)
	_, err = hasher.Write(cipherText[aes.BlockSize:])
	if err != nil {
		return EncryptedData{nil, nil}, err
	}
	calculatedMac := hasher.Sum(nil)

	return EncryptedData{cipherText, calculatedMac}, nil
}

func (self EncryptedData) DecryptTo(target interface{}, encryptKey []byte, authKey []byte) error {
	/* Authenticate the encrypted data by calculating the MAC and
	 * func and comparing it with the given MAC */
	hasher := hmac.New(sha256.New, authKey)
	_, err := hasher.Write(self.CipherText[aes.BlockSize:])
	if err != nil {
		return err
	}
	calculatedMac := hasher.Sum(nil)

	if !hmac.Equal(self.MAC, calculatedMac) {
		return fmt.Errorf("Could not authenticate encrypted data")
	}

	block, err := aes.NewCipher(encryptKey)
	if err != nil {
		return err
	}

	body := make([]byte, len(self.CipherText)-aes.BlockSize)

	/* IV is stored in the first aes.BlockSize bytes of data */
	stream := cipher.NewCTR(block, self.CipherText[:aes.BlockSize])
	stream.XORKeyStream(body, self.CipherText[aes.BlockSize:])

	err = json.Unmarshal(body, target)

	return err
}
