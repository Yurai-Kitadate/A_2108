package auth

import (
	"fmt"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jphacks/A_2108/src/config"
)

var prvKey []byte = []byte(config.PrivateKey())
var pubKey []byte = []byte(config.PublicKey())

type JwtClaims struct {
	jwt.StandardClaims
}

func GenerateToken(id int) (string, error) {
	ecdsaKey, err := jwt.ParseECPrivateKeyFromPEM(prvKey)
	if err != nil {
		return "", err
	}

	// headerのセット
	token := jwt.New(jwt.SigningMethodES512)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = config.GetIssuer()
	claims["sub"] = strconv.Itoa(id)
	claims["exp"] = time.Now().Add(time.Hour * 168).Unix()
	claims["iat"] = time.Now().Unix()

	// 電子署名
	return token.SignedString(ecdsaKey)
}

func VerifyToken(tokenString string) (JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		ecdsaKey, err := jwt.ParseECPublicKeyFromPEM(pubKey)
		if err != nil {
			return "", err
		}
		return ecdsaKey, nil
	})
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return *claims, nil
	} else {
		return *claims, err
	}
}