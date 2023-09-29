package utill

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"strings"
	"time"
)

type TokenJWT struct {
	signingKey     []byte
	ExpiryInterval time.Duration
}

func NewToken(signingKey string, expiryInterval time.Duration) *TokenJWT {
	return &TokenJWT{signingKey: []byte(signingKey), ExpiryInterval: expiryInterval}
}

func (t *TokenJWT) Generate(sub int) (string, error) {
	return t.generate(strconv.Itoa(sub))
}

func (t *TokenJWT) Parse(tokenStr string) (int, error) {
	tokenStr = removeBearerIfExists(tokenStr)

	subStr, err := t.parse(tokenStr)
	if err != nil {
		return 0, err
	}
	sub, err := strconv.Atoi(subStr)
	if err != nil {
		return 0, err
	}
	return sub, nil
}

func (t *TokenJWT) generate(sub string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(t.ExpiryInterval).Unix(),
		Subject:   sub,
	})
	return token.SignedString(t.signingKey)
}

func (t *TokenJWT) parse(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return t.signingKey, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", errors.New("token claims are not type of  jwt.StandardClaims")
	}
	return claims.Subject, nil
}

func removeBearerIfExists(token string) string {
	arr := strings.Split(token, " ")
	if len(arr) < 2 {
		return token
	}
	if len(arr) == 2 && strings.EqualFold("Bearer", arr[0]) {
		return arr[1]
	}
	return token
}
