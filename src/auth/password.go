package auth

import (
	"encoding/hex"
	"errors"

	"github.com/jphacks/A_2108/src/config"
	"golang.org/x/crypto/sha3"
)

func VerifyPassword(password, storedHash string) error {
	hash, err := CreateHash(password)
	if err != nil {
		return err
	}
	// 検証
	if hash != storedHash {
		return errors.New("Invalid password")
	}
	// 結果返却
	return nil
}

func CreateHash(password string) (string, error) {
	hash := sha3.New512()
	p := []byte(password)
	p = append(p, config.GetSalt()...)
	_, err := hash.Write(p)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
