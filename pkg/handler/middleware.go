package handler

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	signingKey = "adfa6464aE"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	GUID string `json:"guid"`
}

func generateAccessToken(guid string) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		guid,
	})
	return accessToken.SignedString([]byte(signingKey))

}

func generateRefreshToken() ([]byte, error) {
	password := make([]byte, 16)

	_, err := rand.Read(password)
	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hash, err
}

func decodeBase64(value string) (string, error) {
	decodedValue, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	return string(decodedValue), nil
}
