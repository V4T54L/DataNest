package utils

import (
	"backend/internals/schemas"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"strings"
)

func Hash(value string) string {
	newValue := value + "hashSecret"

	hasher := sha256.New()
	hasher.Write([]byte(newValue))
	hash := hasher.Sum(nil)
	hashHex := hex.EncodeToString(hash)
	return hashHex
}

func VerifyHash(plain, hashed string) bool {
	return Hash(plain) == hashed
}

func GenerateToken(user schemas.UserDetails) (string, error) {
	data, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	dataStr := string(data)

	token := dataStr + "+" + Hash(dataStr+"tokenSecret")

	encoded := base64.StdEncoding.EncodeToString([]byte(token))

	return encoded, nil
}

func VerifyToken(tokenStr string) (*schemas.UserDetails, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(tokenStr)
	if err != nil {
		log.Println("Error decoding Base64:", err)
		return nil, err
	}
	decoded := string(decodedBytes)

	arr := strings.Split(decoded, "+")
	if len(arr) != 2 {
		return nil, errors.New("invalid token : contains multiple special charcters")
	}
	if Hash(arr[0]+"tokenSecret") != arr[1] {
		return nil, errors.New("invalid token : token has been pampered")
	}

	tokenInfo := schemas.UserDetails{}
	if err := json.Unmarshal([]byte(arr[0]), &tokenInfo); err != nil {
		return nil, err
	}
	return &tokenInfo, nil
}
