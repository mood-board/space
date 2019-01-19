package account

import (
	"crypto/rand"
	"crypto/subtle"
	"fmt"
	"strings"

	"github.com/ofonimefrancis/spaceship/common/must"
	"golang.org/x/crypto/scrypt"
)

const (
	SaltLen        = 32
	HashLen        = 64
	MinPasswordLen = 8
)

type WrongPasswordError string

func (obj WrongPasswordError) Error() string {
	return string(obj)
}

type PasswordHash struct {
	Hash []byte `json:"hash"`
	Salt []byte `json:"salt"`
}

func NewPasswordHash(password string) (*PasswordHash, error) {
	salt := generateSalt()
	hash, err := createPasswordHash(password, salt)
	if err != nil {
		return nil, err
	}
	return &PasswordHash{Hash: hash, Salt: salt}, nil
}

func (self *PasswordHash) IsEqualTo(password string) bool {
	return VerifyPassword(password, self.Hash, self.Salt)
}

// Generate a random salt of suitable length
func generateSalt() []byte {
	salt := make([]byte, SaltLen)
	must.DoF(func() error {
		_, err := rand.Read(salt)
		return err
	})
	return salt
}

func createPasswordHash(password string, salt []byte) ([]byte, error) {
	password = strings.TrimSpace(password)

	if len(password) < MinPasswordLen {
		return nil, fmt.Errorf("This password is too short")
	}

	hash, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, HashLen)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

// VerifyPassword checks that a password matches a stored hash and salt
func VerifyPassword(password string, hash []byte, salt []byte) bool {
	newhash, err := createPasswordHash(password, salt)
	if err != nil {
		return false
	}
	return subtle.ConstantTimeCompare(newhash, hash) == 1
}
