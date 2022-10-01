package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hass), nil
}

func CheckPass(hassPassword, passwordDb string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hassPassword), []byte(passwordDb))
	return err == nil
}
