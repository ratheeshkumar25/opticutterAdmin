package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pasword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pasword), 14)
	if err != nil {
		return "", err
	}
	Password := string(bytes)
	return Password, nil
}

// CheckPassword compares the hashed password with the plain-text password.
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
