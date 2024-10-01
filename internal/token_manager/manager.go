package manager

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "SuPerR_KeeY#"

func CreateAccessToken(guid string, ipAddress string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"id":  guid,
		"ip":  ipAddress,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})
	return token.SignedString(secretKey)
}

func CreateRefreshToken(guid string, ipAddress string) string {
	token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(
		"%s.%s.%d",
		guid,
		ipAddress,
		time.Now().AddDate(0, 1, 0).Unix(),
	)))
	return token
}

func RTokenExpired(token string) (bool, error) {
	byte_string, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return true, err
	}

	components := bytes.Split(byte_string, []byte{'.'})
	if len(components) != 3 {
		return true, errors.New("Bad token format")
	}

	expTimeStr := string(components[2])
	expTime, err := strconv.ParseInt(expTimeStr, 10, 64)
	if err != nil {
		return true, err
	}

	if time.Now().Unix() > expTime {
		return true, nil
	}
	return false, nil
}

func ATokenExpired(token string) (bool, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return true, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		expTime := claims["exp"].(int64)
		return time.Now().Unix() > expTime, nil
	}
	return true, nil
}
