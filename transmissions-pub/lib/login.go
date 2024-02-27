package lib

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"transmissions-pub/database"
	"transmissions-pub/src/models"
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {

	var err error

	u := models.AuthUser{}

	err = database.DB.Model(models.AuthUser{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return "", err
	}

	token, err := GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}
