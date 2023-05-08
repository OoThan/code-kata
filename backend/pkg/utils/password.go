package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"loan-back-services/pkg/logger"
	"strings"

	"golang.org/x/crypto/scrypt"
)

func HashPassword(password string) (string, error) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	shash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}

	// return hex-encoded string with salt appended to password
	hashedPW := fmt.Sprintf("%s.%s", hex.EncodeToString(shash), hex.EncodeToString(salt))

	return hashedPW, nil
}

func ComparePasswords(storedPassword string, suppliedPassword string) (bool, error) {
	pwsalt := strings.Split(storedPassword, ".")
	if len(pwsalt) != 2 {
		return false, errors.New("wrong password")
	}

	// check supplied password salted with hash
	salt, err := hex.DecodeString(pwsalt[1])
	if err != nil {
		return false, errors.New("unable to verify user password")
	}

	shash, err := scrypt.Key([]byte(suppliedPassword), salt, 32768, 8, 1, 32)
	if err != nil {
		return false, err
	}

	logger.Sugar.Debug(hex.EncodeToString(shash) == pwsalt[0])

	return hex.EncodeToString(shash) == pwsalt[0], nil
}
