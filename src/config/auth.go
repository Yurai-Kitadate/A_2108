package config

import (
	"os"
	"strings"
)

func PrivateKey() string {
	return strings.Replace(os.Getenv("PRIVATE_KEY"), "\\n", "\n", -1)
}

func PublicKey() string {
	return strings.Replace(os.Getenv("PUBLIC_KEY"), "\\n", "\n", -1)
}

func GetIssuer() string {
	return os.Getenv("ISSUER")
}

func GetSalt() []byte {
	return []byte(os.Getenv("SALT"))
}
