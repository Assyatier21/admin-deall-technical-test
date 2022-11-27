package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"regexp"

	"github.com/golang-jwt/jwt"
)

func JwtGenerator(username string, role_id int, key string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  username,
		"role_id": role_id,
	})

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return err.Error()
	}
	return tokenString
}

func IsValidAlphabet(s string) bool {
	regex, _ := regexp.Compile(`(^[a-zA-Z]+$)`)
	return regex.MatchString(s)
}

func IsValidAlphaNumeric(s string) bool {
	regex, _ := regexp.Compile(`(^[a-zA-Z0-9]*$)`)
	return regex.MatchString(s)
}

func IsValidAlphaNumericHyphen(s string) bool {
	regex, _ := regexp.Compile(`[a-zA-Z0-9-]+`)
	return regex.MatchString(s)
}

func Hash_256(s string) string {
	hasher := sha256.New()
	hasher.Write([]byte(s))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
