package common

import "golang.org/x/crypto/bcrypt"

type Encryptor struct {
}

func (encryptor Encryptor) Encrypt(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

func (encryptor Encryptor) CompareHashAndPassword(password string, hashedPassword string) bool {
	incoming := []byte(password)
	existing := []byte(hashedPassword)
	return bcrypt.CompareHashAndPassword(existing, incoming) == nil
}
