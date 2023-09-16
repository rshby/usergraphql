package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
	"usergraphql/app/user"
)

type AuthHelper struct {
}

var secretKey = os.Getenv("SECRET_KEY")

// method generate token
func (a *AuthHelper) GenerateToken(user *user.User) (string, error) {
	// create claims
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Local().Add(10 * time.Hour),
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenString.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

// method hashPassword
func (a *AuthHelper) GetHashedPassword(password string) (string, error) {
	hasehdPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hasehdPassword), nil
}

// mehod cek password
func (a *AuthHelper) CekPassword(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
