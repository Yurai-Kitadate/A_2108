package auth

import (
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/sha3"
)

func VerifyPassword(password, storedHash string) error {
	// Hash生成
	hash := sha3.New512()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return err
	}
	// 検証
	if hex.EncodeToString(hash.Sum(nil)) != storedHash {
		return errors.New("Invalid password")
	}
	// 結果返却
	return nil
}
