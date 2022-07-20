package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func HashPasswordBcrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPasswordBcrypt(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

func CreatePasswordSHA(password string, saltSize int) (hashedPassword, saltString string) {
	passwordBytes, _ := base64.URLEncoding.DecodeString(password)

	saltBytes, _ := GenerateRandomSalt(saltSize)
	saltString = base64.URLEncoding.EncodeToString(saltBytes[:])

	hashedPassword = HashPasswordSHA(passwordBytes, saltBytes)

	return hashedPassword, saltString
}

func HashPasswordSHA(passwordBytes, saltBytes []byte) string {
	saltedPasswordBytes := append(passwordBytes, saltBytes...)

	hash := sha256.Sum256(saltedPasswordBytes)

	hashedPassword := base64.URLEncoding.EncodeToString(hash[:])

	return hashedPassword
}

func ComparePasswordSHA(givenPassword, hashedPassword, salt string) bool {
	givenPasswordBytes, _ := base64.URLEncoding.DecodeString(givenPassword)
	saltBytes, _ := base64.URLEncoding.DecodeString(salt)

	givenHashedPassword := HashPasswordSHA(givenPasswordBytes, saltBytes)

	return givenHashedPassword == hashedPassword
}

func GenerateRandomSalt(saltSize int) ([]byte, error) {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])
	if err != nil {
		return nil, err
	}

	return salt, nil
}
